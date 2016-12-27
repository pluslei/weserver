package haoindex

import (
	"github.com/astaxie/beego"
	m "weserver/models"
	"weserver/src/tools"
	// "github.com/berkaroad/weixinapi"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/oauth"
)

type IndexController struct {
	CommonController
}

var (
	appId     = "wx26ed6ed15f2a7b17"
	appSecret = "1ac297e601224d5ab6bafd6ceacb1228"

	redirect_uri = "http://live.780.com.cn"
	wx           *wechat.Wechat
	oauthAccess  *oauth.Oauth
)

type Userinfor struct {
	Codeid    string //房间号公司代码加密
	Uname     string //用户名
	Nickname  string //用户昵称
	UserIcon  string //logo
	RoleName  string //用户角色
	Titlerole string //用户类型
	Authorcss string //头衔
	Insider   int64  //1内部人员或0外部人员
	IsLogin   bool   //是否登入
}

func init() {
	rediscache, err := cache.NewRediscache("127.0.0.1", uint(6379))
	if err != nil {
		beego.Error("connect rediscache error", err)
	}
	cfg := &wechat.Config{
		AppID:          appId,
		AppSecret:      appSecret,
		Token:          "Token",
		EncodingAESKey: "EncodingAESKey",
		Cache:          rediscache,
	}
	wx = wechat.NewWechat(cfg)
}

func (this *IndexController) Get() {
	// if this.CheckUserIsAuth() {
	// 	this.Redirect("/index", 302)
	// }

	code := this.GetString("code")
	if code == "" {
		oauthAccess = wx.GetOauth(this.Ctx.Request, this.Ctx.ResponseWriter)
		err := oauthAccess.Redirect(redirect_uri, "snsapi_userinfo", "ihaoyue")
		if err != nil {
			this.Redirect("/", 302)
			beego.Error("error", err)
			return
		}
	} else {
		resToken, err := oauthAccess.GetUserAccessToken(code)
		if err != nil {
			beego.Error("get the user token error", err)
			this.Redirect("/", 302)
			return
		}
		userInfo, err := oauthAccess.GetUserInfo(resToken.AccessToken, resToken.OpenID)
		if err != nil {
			beego.Error("get the userinfo error", err)
			this.Redirect("/", 302)
			return
		}

		info, err := m.GetUserByUsername(userInfo.OpenID)
		if err != nil || info.Id <= 0 {
			this.saveUser(userInfo)
		} else {
			this.updateUser(info.Id, userInfo)
		}

		sessionUser, _ := m.GetUserByUsername(userInfo.OpenID)
		this.SetSession("indexUserInfo", &sessionUser)

		beego.Info("user info:", userInfo)
		this.Data["userInfo"] = userInfo
		this.TplName = "dist/index.html"
		this.Redirect("/index", 302)
	}
}

func (this *IndexController) Index() {
	Info := this.GetSession("indexUserInfo")
	if Info != nil {
		// userInfo := new(m.User)
		userInfo := Info.(*m.User)
		beego.Debug("userInfo--==", userInfo.Username)
		userLoad, err := m.LoadRelatedUser(userInfo, "Username")
		if err != nil {
			beego.Error("load retalteduser error", err)
		}
		beego.Debug("get the userload:", userLoad)
		user := new(Userinfor)
		prevalue := "100" + "_" + "10000"
		codeid := tools.MainEncrypt(prevalue)
		user.Codeid = codeid
		user.Uname = userInfo.Username
		user.Nickname = userInfo.Nickname
		user.IsLogin = true
		user.RoleName = userLoad.Role.Name

		if userLoad.Role.Id > 0 {
			user.Titlerole = userLoad.Role.Title //用户类型
		} else {
			user.Titlerole = "游客" //用户类型
		}
		user.Authorcss = "/upload/usertitle/visitor.png"
		user.UserIcon = userInfo.Headimgurl
		user.Insider = 0 //1内部人员或0外部人员
		this.Data["user"] = user

		url := "http://" + this.Ctx.Request.Host + this.Ctx.Input.URI()

		jssdk := wx.GetJs(this.Ctx.Request, this.Ctx.ResponseWriter)
		jsapi, err := jssdk.GetConfig(url)
		if err != nil {
			beego.Error("get the jsapi config error", err)
		}
		this.Data["appId"] = appId
		this.Data["timestamp"] = jsapi.TimeStamp //jsapi.Timestamp
		this.Data["nonceStr"] = jsapi.NonceStr   //jsapi.NonceStr
		this.Data["signature"] = jsapi.Signature //jsapi.Signature

		this.TplName = "dist/index.html"
	} else {
		this.Redirect("/", 302)
	}

}

func (this *IndexController) Login() {
	this.TplName = "login.html"
}

func (this *IndexController) Voice() {
	url := "http://" + this.Ctx.Request.Host + this.Ctx.Input.URI()

	jssdk := wx.GetJs(this.Ctx.Request, this.Ctx.ResponseWriter)
	jsapi, err := jssdk.GetConfig(url)
	if err != nil {
		beego.Error("get the jsapi config error", err)
	}
	this.Data["appId"] = appId
	this.Data["timestamp"] = jsapi.TimeStamp //jsapi.Timestamp
	this.Data["nonceStr"] = jsapi.NonceStr   //jsapi.NonceStr
	this.Data["signature"] = jsapi.Signature //jsapi.Signature
	this.TplName = "voice.html"
}

func (this *IndexController) saveUser(userInfo oauth.UserInfo) bool {
	config, _ := m.GetSysConfig()
	configRole := config.Registerrole
	configTitle := config.Registertitle
	configVerify := config.Verify
	u := new(m.User)
	u.Username = userInfo.OpenID
	u.Status = 1
	if configVerify == 0 {
		u.RegStatus = 1
	} else {
		u.RegStatus = 2
	}
	u.UserIcon = userInfo.HeadImgURL
	u.Role = &m.Role{Id: configRole}
	u.Title = &m.Title{Id: configTitle}
	u.Nickname = userInfo.Nickname
	u.Openid = userInfo.OpenID
	u.Sex = userInfo.Sex
	u.Province = userInfo.Province
	u.City = userInfo.City
	u.Country = userInfo.Country
	u.Headimgurl = userInfo.HeadImgURL
	u.Unionid = userInfo.Unionid
	userid, err := m.AddUser(u)
	if err == nil && userid > 0 {
		return true
	} else {
		beego.Error(err)
		return false
	}
	return false
}

func (this *IndexController) updateUser(id int64, userInfo oauth.UserInfo) error {
	u := new(m.User)
	u.Id = id
	u.UserIcon = userInfo.HeadImgURL
	u.Nickname = userInfo.Nickname
	u.Sex = userInfo.Sex
	u.Province = userInfo.Province
	u.City = userInfo.City
	u.Country = userInfo.Country
	u.Headimgurl = userInfo.HeadImgURL
	u.Unionid = userInfo.Unionid
	return u.UpdateUserFields("UserIcon", "Nickname", "Sex", "Province", "City", "Country", "Headimgurl", "Unionid")
}
