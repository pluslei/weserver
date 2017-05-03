package haoindex

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
	"weserver/src/tools"

	m "weserver/models"

	"github.com/astaxie/beego"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/oauth"
)

type IndexController struct {
	CommonController
}

var (
	APPID     = beego.AppConfig.String("APPID")
	APPSECRET = beego.AppConfig.String("APPSECRET")

	redirect_uri = beego.AppConfig.String("httplocalServerAdress")
	Wx           *wechat.Wechat
	oauthAccess  *oauth.Oauth
)

type Userinfor struct {
	Uname         string //用户名
	Nickname      string //用户昵称
	UserIcon      string //logo
	RoleName      string //用户角色[vip,silver,gold,jewel]
	RoleTitle     string //用户角色名[会员,白银会员,黄金会员,钻石会员]
	RoleTitleCss  string //用户角色样式
	RoleTitleBack bool   //角色聊天背景
	RoleIcon      string //用户角色默认头像
	RoleId        int64
	Insider       int64 //1内部人员或0外部人员
	IsLogin       bool  //是否登入
	IsFilter      bool  //是否检查
}

type VoiceResponse struct {
	Staus   bool
	Wavfile string
	Info    string
}

func init() {
	macache := cache.NewMemcache()
	cfg := &wechat.Config{
		AppID:          APPID,
		AppSecret:      APPSECRET,
		Token:          "Token",
		EncodingAESKey: "EncodingAESKey",
		Cache:          macache,
	}
	Wx = wechat.NewWechat(cfg)
}

// 获取userinfo
func (this *IndexController) Get() {
	// if this.CheckUserIsAuth() {
	// 	this.Redirect("/index", 302)
	// }
	code := this.GetString("code")
	if code == "" {
		oauthAccess = Wx.GetOauth(this.Ctx.Request, this.Ctx.ResponseWriter)
		err := oauthAccess.Redirect(redirect_uri, "snsapi_userinfo", "ihaoyue")
		if err != nil {
			beego.Error("oauthAccess error", err)
			this.Redirect("/", 302)
			return
		}
	} else {
		resToken, err := oauthAccess.GetUserAccessToken(code)
		if err != nil {
			beego.Error("get the user token error", err)
			this.Redirect("/", 302)
			return
		}

		_, err = oauthAccess.CheckAccessToken(resToken.AccessToken, APPID)
		if err != nil {
			beego.Error("CheckAccessToken error", err)
		}
		beego.Debug("AccessToken", resToken.AccessToken)

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
		beego.Debug("userInfo", userInfo)

		if len(info.Account) > 0 {
			sessionUser, _ := m.GetUserByUsername(userInfo.OpenID)
			this.SetSession("indexUserInfo", &sessionUser)
			this.Redirect("/index", 302)
		} else {
			this.Redirect("/login?openid="+userInfo.OpenID, 302)
		}
	}
	this.Ctx.WriteString("")
}

func (this *IndexController) Login() {
	openid := this.GetString("openid")
	if this.IsAjax() {
		if len(openid) <= 0 {
			this.Rsp(false, "请退出重新进入", "")
			return
		}
		username := this.GetString("username")
		password := this.GetString("password")
		beego.Debug("userinfo", username, password)
		if len(username) <= 0 || len(password) <= 0 {
			this.Rsp(false, "请填写账户信息", "")
			return
		}

		beego.Debug("openid", openid)
		var u m.User
		u.Account = username
		user, err := m.ReadFieldUser(&u, "Account")
		beego.Debug("ReadFieldUser", user, err)
		if user == nil || err != nil {
			this.Rsp(false, "用户名和密码错误 401", "")
			return
		}
		beego.Debug("userinfo", user.Password, tools.EncodeUserPwd(username, password), err)
		if user.Password != tools.EncodeUserPwd(username, password) {
			this.Rsp(false, "用户名和密码错误 402", "")
			beego.Debug("PassWord Error")
			return
		}

		_, err = m.BindUserAccount(openid, user)
		if err != nil {
			this.Rsp(false, "用户名和密码错误 404", "")
			beego.Debug("Bind User Account Error", err)
			return
		}
		if user.Username == "" {
			_, err := m.DelUserById(user.Id)
			if err != nil {
				beego.Debug("DELETE User ID Error", err)
				return
			}
		}
		sessionUser, err := m.GetUserByUsername(openid)
		if err != nil {
			this.Rsp(false, "用户名和密码错误 405", "")
			beego.Debug("Get UseInfo Error", err)
			return
		}
		beego.Debug("sssssssssssssssss")
		_, err1 := m.UpdateRegistName(user.Id, sessionUser.Username, sessionUser.UserIcon)
		if err1 != nil {
			this.Rsp(false, "用户名和密码错误 406", "")
			beego.Debug("Update Regist UserName Error", err1)
			return
		}
		this.SetSession("indexUserInfo", &sessionUser)
		this.Redirect("/index", 302)
	}
	this.Data["openid"] = openid
	this.TplName = "haoindex/login.html"
	// this.Ctx.WriteString("")
}

//从数据库获取信息
func (this *IndexController) Index() {
	Info := this.GetSession("indexUserInfo")
	beego.Debug("info", Info)
	if Info != nil {
		// userInfo := new(m.User)
		userInfo := Info.(*m.User)
		sysconfig, _ := m.GetAllSysConfig() //系统设置
		userLoad, err := m.LoadRelatedUser(userInfo, "Username")
		if err != nil {
			beego.Error("load retalteduser error", err)
		}
		user := new(Userinfor)
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
		// if userLoad.Status == 2 && userLoad.RegStatus == 2 {
		// 	user.IsLogin = true
		// } else {
		// 	user.IsLogin = false
		// }
		user.IsLogin = true
		if userLoad.Role.Id > 0 {
			user.RoleId = userLoad.Role.Id
		}
		if userLoad.Title.Id > 0 {
			user.RoleTitle = userLoad.Title.Name //用户类型
		} else {
			user.RoleTitle = "游客" //用户类型
		}
		user.RoleIcon = "/upload/usertitle/" + userLoad.Title.Css

		// 消息审核(0 开启 1 关闭(默认))
		// 是否隶属公司内部角色[0、否 1、是]
		beego.Debug("userload", userLoad.Role.IsInsider, sysconfig, sysconfig.AuditMsg)
		if sysconfig.AuditMsg == 1 {
			user.IsFilter = false
		} else {
			user.IsFilter = true
		}

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

		jssdk := Wx.GetJs(this.Ctx.Request, this.Ctx.ResponseWriter)
		jsapi, err := jssdk.GetConfig(url)
		if err != nil {
			beego.Error("get the jsapi config error", err)
		}
		this.Data["appId"] = APPID
		this.Data["timestamp"] = jsapi.TimeStamp //jsapi.Timestamp
		this.Data["nonceStr"] = jsapi.NonceStr   //jsapi.NonceStr
		this.Data["signature"] = jsapi.Signature //jsapi.Signature

		system, _ := m.GetSysConfig() //获取配置表数据
		this.Data["system"] = system
		this.Data["serverurl"] = beego.AppConfig.String("localServerAdress") //链接
		this.Data["serviceimg"] = beego.AppConfig.String("serviceimg")       //客服图片
		this.Data["loadingimg"] = beego.AppConfig.String("loadingimg")       //公司logo
		this.Data["servicephone"] = beego.AppConfig.String("servicephone")   //服务电话
		this.TplName = "dist/index.html"
		// this.TplName = "index.html"
	} else {
		this.Redirect("/", 302)
	}
}

// 后台获取声音文件流转换
func (this *IndexController) Voice() {
	media := this.GetString("media")

	var filename string

	savepath := fmt.Sprintf("../upload/temp/%s/", time.Now().Format("2006-01-02"))
	wavfilename := savepath + media + ".wav"

	voice := new(VoiceResponse)
	if Exist(wavfilename) {
		voice.Staus = true
		voice.Wavfile = wavfilename
	} else {
		material := Wx.GetMaterial()
		mediaURL, err := material.GetMediaURL(media)

		if !Exist(savepath) {
			os.MkdirAll(savepath, 0755)
		}
		filename = media + ".amr"
		savefile := savepath + filename
		resp, err := http.Get(mediaURL)
		if err != nil {
			beego.Error("http get media url error", err)
			voice.Info = err.Error()
		} else {
			file, err := os.Create(savefile)
			defer file.Close()
			if err != nil {
				beego.Error("create file error", err)
				voice.Info = err.Error()
			} else {
				io.Copy(file, resp.Body)
				defer resp.Body.Close()

				wavfile, err := this.AmrToWav(savepath, filename)
				if err == nil || wavfile != "" {
					voice.Staus = true
					voice.Wavfile = wavfile
					os.Remove(savefile)
				} else {
					voice.Info = err.Error()
				}
			}
		}
	}
	this.Data["json"] = voice
	this.ServeJSON()

}

// 获取图片媒体的文档流
func (this *IndexController) GetMediaURL() {
	media := this.GetString("media")
	material := Wx.GetMaterial()
	mediaURL, err := material.GetMediaURL(media)
	beego.Info("mediaURL is", mediaURL)
	srcfile := redirect_uri + "/static/images/nono.jpg"
	if err == nil {
		resp, err := http.Get(mediaURL)
		if err != nil {
			file, _ := os.Open(srcfile)
			defer file.Close()
			io.Copy(this.Ctx.ResponseWriter, file)
		} else {
			if resp.Header.Get("Content-Type") != "text/plain" {
				defer resp.Body.Close()
				io.Copy(this.Ctx.ResponseWriter, resp.Body)
			}
		}
	} else {
		beego.Error("get the media url error", err)
		file, _ := os.Open(srcfile)
		defer file.Close()
		io.Copy(this.Ctx.ResponseWriter, file)
	}
	this.Ctx.WriteString("")
}

// 保存用户至数据库
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
	u.Lastlogintime = time.Now()
	userid, err := m.AddUser(u)
	beego.Debug("user", u)
	if err == nil && userid > 0 {
		return true
	} else {
		beego.Error(err)
		return false
	}
	return false
}

// 更新用户数据
func (this *IndexController) updateUser(id int64, userInfo oauth.UserInfo) error {
	u := new(m.User)
	u.Id = id
	u.UserIcon = userInfo.HeadImgURL
	u.Sex = userInfo.Sex
	u.Province = userInfo.Province
	u.City = userInfo.City
	u.Country = userInfo.Country
	u.Headimgurl = userInfo.HeadImgURL
	u.Unionid = userInfo.Unionid
	u.Lastlogintime = time.Now()
	return u.UpdateUserFields("UserIcon", "Sex", "Province", "City", "Country", "Headimgurl", "Unionid", "Lastlogintime")
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

// 声音转换 Amr=>Wav
func (this *IndexController) AmrToWav(filedir, filename string) (string, error) {
	newfilename := filename[0:strings.LastIndex(filename, ".")]

	oldpathfilename := filedir + filename
	savepathfilename := filedir + newfilename + ".wav"
	if Exist(savepathfilename) {
		return savepathfilename, nil
	}

	toolName := "static/wmv_tools.static" //转换工具路径
	if Exist(toolName) == false {
		beego.Error("don't find the tools")
		return "", errors.New("don't find the tools")
	}

	cmdStr := toolName + " " + oldpathfilename + " " + savepathfilename //执行转换格式工具命令
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	err := cmd.Run()
	if err != nil {
		chmod := "chmod 755 " + toolName
		exec.Command("/bin/sh", "-c", chmod)
		beego.Error("chmod error", err)
		return "", err
	}
	return savepathfilename, nil
}

func (this *IndexController) WxServerImg() {
	media := this.GetString("img")
	imgpath := GetWxServerImg(media)
	beego.Debug("firlena", imgpath)
}

// 获取图片保存至本地
func GetWxServerImg(media string) (imgpath string) {
	material := Wx.GetMaterial()
	mediaURL, err := material.GetMediaURL(media)
	beego.Info("mediaURL is", mediaURL)

	notFile := "/static/images/nono.jpg"
	if err != nil {
		beego.Debug("get url error", err)
		return notFile
	} else {
		resp, err := http.Get(mediaURL)
		defer resp.Body.Close()
		if err != nil {
			beego.Error("get images error:", err)
			return notFile
		}
		dir := path.Join("..", "upload", "room")
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			beego.Error("mkdir images dir error:", err)
			return notFile
		}
		nowtime := time.Now().UnixNano()
		extimg := "jpg"
		if ext, ok := tools.ContentTypeToExt[resp.Header.Get("Content-Type")]; ok {
			extimg = ext
		}
		FileName := fmt.Sprintf("%d%s%s%s", nowtime, tools.RandomNumeric(4), ".", extimg)
		dirPath := path.Join("..", "upload", "room", FileName)

		f, err := os.Create(dirPath)
		defer f.Close()
		if err != nil {
			beego.Error("create images error:", err)
			return notFile
		}
		_, err = io.Copy(f, resp.Body)
		if err != nil {
			beego.Error("ioread error", err)
			return notFile
		}
		return dirPath
	}
}
