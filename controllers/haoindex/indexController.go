package haoindex

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
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
	Codeid        string //房间号公司代码加密
	Uname         string //用户名
	Nickname      string //用户昵称
	UserIcon      string //logo
	RoleName      string //用户角色[vip,silver,gold,jewel]
	RoleTitle     string //用户角色名[会员,白银会员,黄金会员,钻石会员]
	RoleTitleCss  string //用户角色样式
	RoleTitleBack bool   //角色聊天背景
	RoleIcon      string //用户角色默认头像
	Insider       int64  //1内部人员或0外部人员
	IsLogin       bool   //是否登入
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
	beego.Debug("wx tokenaccess", wx)
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
		beego.Debug("user info:", userInfo)
		this.Redirect("/index", 302)
	}
}

//
func (this *IndexController) Index() {
	Info := this.GetSession("indexUserInfo")
	if Info != nil {
		// userInfo := new(m.User)
		userInfo := Info.(*m.User)
		sysconfig, _ := m.GetAllSysConfig() //系统设置
		userLoad, err := m.LoadRelatedUser(userInfo, "Username")
		if err != nil {
			beego.Error("load retalteduser error", err)
		}
		beego.Debug("get the userload:", userLoad)
		beego.Debug("get the userload:", userLoad.Role)
		beego.Debug("get the userload:", userLoad.Title)
		user := new(Userinfor)
		prevalue := beego.AppConfig.String("company") + "_" + beego.AppConfig.String("room")
		codeid := tools.MainEncrypt(prevalue)
		user.Codeid = codeid
		user.Uname = userInfo.Username
		user.UserIcon = userInfo.Headimgurl
		user.RoleName = userLoad.Role.Name

		// 设置昵称使用设置的
		if len(userInfo.Remark) <= 0 {
			user.Nickname = userInfo.Nickname
		} else {
			user.Nickname = userInfo.Remark
		}

		// 用户为禁用和未审核状态不准登录
		if userLoad.Status == 2 && userLoad.RegStatus == 2 {
			user.IsLogin = true
		} else {
			user.IsLogin = false
		}

		if userLoad.Title.Id > 0 {
			user.RoleTitle = userLoad.Title.Name //用户类型
		} else {
			user.RoleTitle = "游客" //用户类型
		}
		user.RoleIcon = "/upload/usertitle/" + userLoad.Title.Css

		// 设置头衔颜色
		if len(userLoad.Title.Css) <= 0 {
			user.RoleTitleCss = "#000000"
		} else {
			user.RoleTitleCss = userLoad.Title.Css
		}

		// RoleTitleBack
		if userLoad.Title.Background == 1 {
			user.RoleTitleBack = true
		} else {
			user.RoleTitleBack = false
		}

		user.Insider = 1                          //1内部人员或0外部人员
		this.Data["title"] = sysconfig.WelcomeMsg //公告
		this.Data["user"] = user
		url := "http://" + this.Ctx.Request.Host + this.Ctx.Input.URI()

		jssdk := wx.GetJs(this.Ctx.Request, this.Ctx.ResponseWriter)
		jsapi, err := jssdk.GetConfig(url)
		if err != nil {
			beego.Error("get the jsapi config error", err)
			this.Ctx.WriteString("请联系客服 代码：404")
		}
		this.Data["appId"] = appId
		this.Data["timestamp"] = jsapi.TimeStamp //jsapi.Timestamp
		this.Data["nonceStr"] = jsapi.NonceStr   //jsapi.NonceStr
		this.Data["signature"] = jsapi.Signature //jsapi.Signature

		system, _ := m.GetSysConfig() //获取配置表数据
		this.Data["system"] = system
		//this.TplName = "dist/index.html"
		this.TplName = "index.html"

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

func (this *IndexController) GetMediaURL() {
	// type medio struct {
	// 	Status   bool   `json:"status"`
	// 	Mediourl string `json:"mediourl"`
	// }

	media := this.GetString("media")
	material := wx.GetMaterial()
	mediaURL, err := material.GetMediaURL(media)
	srcfile := beego.AppConfig.String("imagesrc") + "/static/images/nono.jpg"
	if err == nil {
		filepath := "./static/down/" + media + ".jpg"
		if _, err := os.Stat(filepath); err != nil {
			req := httplib.Get(mediaURL)
			req.ToFile(filepath)
		}

		if _, err := os.Stat(filepath); err != nil {
			srcfile = beego.AppConfig.String("imagesrc") + "/static/images/nono.jpg"
		} else {
			srcfile = beego.AppConfig.String("imagesrc") + "/static/down/" + media + ".jpg"
		}
	}

	// m := new(medio)
	// m.Status = true
	// m.Mediourl = srcfile
	// this.Data["json"] = m
	// this.ServeJSON()
	this.Redirect(srcfile, 301)
}

func (this *IndexController) saveUser(userInfo oauth.UserInfo) bool {
	config, _ := m.GetSysConfig()
	configRole := config.Registerrole
	configTitle := config.Registertitle
	configVerify := config.Verify
	u := new(m.User)
	u.Username = userInfo.OpenID
	if configVerify == 0 { //是否开启验证  0开启 1不开启
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
	u.Status = 2
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

func (this *IndexController) SetNickname() {
	id, _ := this.GetInt64("id")
	username := this.GetString("username")
	remark := this.GetString("nickname")

	usr := new(m.User)
	usr.Id = id
	usr.Username = username
	user, err := m.ReadFieldUser(usr, "Id", "Username")
	if err != nil {
		beego.Error("get the user error:", err)
	} else {
		user.Remark = remark
		err := user.UpdateUserFields("Remark")
		if err != nil {
			this.Rsp(false, "修改昵称失败", "")
		} else {
			this.Rsp(true, "昵称修改成功", "")
		}
	}
}
