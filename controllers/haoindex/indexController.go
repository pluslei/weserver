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
	. "weserver/src/cache"
	"weserver/src/tools"

	m "weserver/models"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/oauth"
)

type IndexController struct {
	CommonController
}

type Userinfor struct {
	Uname         string //用户名
	CompanyId     int64
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

func GetWxObj(id int64) (*wechat.Wechat, string, string) {
	var info m.Company
	var err error
	strId := strconv.FormatInt(id, 10)
	inter, ok := MapCache[strId]
	if !ok {
		info, err = m.GetCompanyById(id)
		if err != nil {
			beego.Debug("get login companyinfo error")
			return nil, "", ""
		}
	} else {
		info, _ = inter.(m.Company)
		beego.Debug("memcache find")
	}
	macache := cache.NewMemcache()
	cfg := &wechat.Config{
		AppID:          info.AppId,
		AppSecret:      info.AppSecret,
		Token:          "Token",
		EncodingAESKey: "EncodingAESKey",
		Cache:          macache,
	}
	return wechat.NewWechat(cfg), info.AppId, info.Url
}

func (this *IndexController) Redirectr() {
	var user = new(m.User)
	Id := this.GetString("id")
	this.SetSession("LoginInfo", user)
	this.Redirect("/wechat?state="+Id, 302)
}

func (this *IndexController) WeChatLogin() {
	companyId := this.GetString("id")
	beego.Debug("url id", companyId)
	this.TplName = "haoindex/login.html"
}

func (this *IndexController) LoginHandle() {
	var user = new(m.User)
	username := this.GetString("username")
	password := this.GetString("password")
	if len(username) <= 0 || len(password) <= 0 {
		this.AlertBack("请填写账户信息")
		return
	}

	user, err := m.ReadFieldUser(&m.User{Account: username}, "Account")
	if user == nil || err != nil {
		this.AlertBack("用户名异常 401")
		return
	}

	if user.Password != tools.EncodeUserPwd(username, password) {
		this.AlertBack("用户名和密码错误 402")
		beego.Debug("PassWord Error")
		return
	}

	Id := strconv.FormatInt(user.CompanyId, 10)
	this.SetSession("LoginInfo", user)
	this.Redirect("/wechat?state="+Id, 302)
}

func (this *IndexController) PCLogin() {
	if this.IsAjax() {
		var user = new(m.User)
		username := this.GetString("username")
		password := this.GetString("password")

		user, err := m.ReadFieldUser(&m.User{Account: username}, "Account")
		if user == nil || err != nil {
			this.AlertBack("用户名异常 401")
			return
		}

		if user.Password != tools.EncodeUserPwd(username, password) {
			this.AlertBack("用户名和密码错误 402")
			beego.Debug("PassWord Error")
			return
		}

		info, err := m.GetCompanyById(user.CompanyId)
		if err != nil {
			beego.Debug("get login company id error")
			return
		}
		beego.Debug("info", info)
	}
}

// 获取userinfo
func (this *IndexController) GetWeChatInfo() {
	Id := this.GetString("state")
	nId, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		beego.Debug("get company id error", err)
		return
	}

	Wx, AppId, Url := GetWxObj(nId)
	if Wx == nil {
		beego.Debug("Get Wx Object Fail")
		return
	}

	code := this.GetString("code")
	beego.Debug("code", code)
	if code == "" {
		oauthAccess := Wx.GetOauth(this.Ctx.Request, this.Ctx.ResponseWriter)
		err := oauthAccess.Redirect(Url, "snsapi_userinfo", Id)
		if err != nil {
			beego.Error("oauthAccess error", err)
			this.Redirect("/login", 302)
			return
		}
	} else {
		oauthAccess := Wx.GetOauth(this.Ctx.Request, this.Ctx.ResponseWriter)
		resToken, err := oauthAccess.GetUserAccessToken(code)
		if err != nil {
			beego.Error("get the user token error", err)
			this.Redirect("/login", 302)
			return
		}

		_, err = oauthAccess.CheckAccessToken(resToken.AccessToken, AppId)
		if err != nil {
			beego.Error("CheckAccessToken error", err)
		}

		userInfo, err := oauthAccess.GetUserInfo(resToken.AccessToken, resToken.OpenID)
		if err != nil {
			beego.Error("get the userinfo error", err)
			this.Redirect("/login", 302)
			return
		}

		info, err := m.GetUserByUsername(userInfo.OpenID)

		if err != nil || info.Id <= 0 {
			this.saveUser(userInfo)
		} else {
			this.updateUser(info.Id, userInfo)
		}
		// if len(info.Account) > 0 {
		sessionUser, _ := m.GetUserByUsername(userInfo.OpenID)
		this.SetSession("indexUserInfo", &sessionUser)
		this.Redirect("/index", 302)
		// } else {
		// 	this.Redirect("/login", 302)
		// }
		beego.Debug("userInfo", userInfo)
	}
	this.Ctx.WriteString("")
}

func (this *IndexController) Index() {
	indexUserInfo := this.GetSession("indexUserInfo")
	if indexUserInfo != nil {
		userInfo := new(m.User)
		userInfo.Account = indexUserInfo.(*m.User).Account
		userLoad, err := m.LoadRelatedUser(userInfo, "Account")
		beego.Debug("userInfo and load", userInfo, userLoad)
		if err != nil {
			beego.Error("load retalteduser error", err)
			return
		}
		user := new(Userinfor)
		user.Uname = userInfo.Username
		user.UserIcon = userInfo.Headimgurl
		user.RoleName = userLoad.Role.Name
		user.CompanyId = userLoad.CompanyId

		// 设置昵称使用设置的
		if len(userInfo.Remark) <= 0 {
			user.Nickname = userInfo.Nickname
		} else {
			user.Nickname = userInfo.Remark
		}
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

		var info m.Company
		strId := strconv.FormatInt(user.CompanyId, 10)
		inter, ok := MapCache[strId]
		if !ok {
			info, err = m.GetCompanyById(user.CompanyId)
			if err != nil {
				beego.Debug("get login companyinfo error")
			}
		} else {
			info, _ = inter.(m.Company)
		}

		// 消息审核(0 开启 1 关闭(默认))
		// 是否隶属公司内部角色[0、否 1、是]
		if info.AuditMsg == 1 {
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

		user.Insider = 1                     //1内部人员或0外部人员
		this.Data["title"] = info.WelcomeMsg //公告
		this.Data["user"] = user

		Wx, AppId, _ := GetWxObj(userInfo.CompanyId)

		url := "http://" + this.Ctx.Request.Host + this.Ctx.Input.URI()
		beego.Debug("url", url)

		jssdk := Wx.GetJs(this.Ctx.Request, this.Ctx.ResponseWriter)
		jsapi, err := jssdk.GetConfig(url)
		if err != nil {
			beego.Error("get the jsapi config error", err)
		}
		this.Data["appId"] = AppId
		this.Data["timestamp"] = jsapi.TimeStamp //jsapi.Timestamp
		this.Data["nonceStr"] = jsapi.NonceStr   //jsapi.NonceStr
		this.Data["signature"] = jsapi.Signature //jsapi.Signature

		this.Data["system"] = info.WelcomeMsg
		this.Data["serverurl"] = beego.AppConfig.String("localServerAdress") //链接
		this.Data["serviceimg"] = beego.AppConfig.String("serviceimg")       //客服图片
		this.Data["loadingimg"] = beego.AppConfig.String("loadingimg")       //公司logo
		this.Data["servicephone"] = beego.AppConfig.String("servicephone")   //服务电话
		this.TplName = "dist/index.html"
		// this.TplName = "index.html"
	} else {
		this.Redirect("/login", 302)
	}
}

// 后台获取声音文件流转换
func (this *IndexController) Voice() {
	Id := this.GetString("Id") //companyId
	nId, err := strconv.ParseInt(Id, 64, 10)
	if err != nil {
		beego.Debug("GetMediaUrl CompanyId Error", err)
		return
	}
	Wx, _, _ := GetWxObj(nId)
	if Wx == nil {
		beego.Debug("GetMediaUrl Wx object Error")
		return
	}
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
	Id := this.GetString("Id") //companyId
	nId, err := strconv.ParseInt(Id, 64, 10)
	if err != nil {
		beego.Debug("GetMediaUrl CompanyId Error", err)
		return
	}
	Wx, _, Url := GetWxObj(nId)
	if Wx == nil {
		beego.Debug("GetMediaUrl Wx object Error")
		return
	}
	media := this.GetString("media")
	material := Wx.GetMaterial()
	mediaURL, err := material.GetMediaURL(media)
	beego.Info("mediaURL is", mediaURL)
	srcfile := Url + "/static/images/nono.jpg"
	if err == nil {
		resp, err := http.Get(mediaURL)
		beego.Debug("resp", resp)
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

func (this *IndexController) saveUser(userInfo oauth.UserInfo) bool {
	info := this.GetSession("LoginInfo").(*m.User)
	if info.Account == "" {
		this.Redirect("/login", 302)
	}

	if len(info.Username) <= 0 {
		_, err := m.BindWechatIcon(info.Id, &userInfo)
		if err != nil {
			beego.Debug("Bind User Account Error", err)
			return false
		}
		_, err1 := m.UpdateRegistName(info.Id, userInfo.OpenID, userInfo.Nickname, userInfo.HeadImgURL)
		if err1 != nil {
			beego.Debug("Update Regist UserName Error", err1)
			return false
		}
	}
	return true
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
	Id := this.GetString("Id") //companyId
	media := this.GetString("img")
	imgpath := GetWxServerImg(media, Id)
	beego.Debug("firlena", imgpath)
}

// 获取图片保存至本地
func GetWxServerImg(media, Id string) (imgpath string) {

	nId, err := strconv.ParseInt(Id, 64, 10)
	if err != nil {
		beego.Debug("GetMediaUrl CompanyId Error", err)
		return
	}
	Wx, _, _ := GetWxObj(nId)
	if Wx == nil {
		beego.Debug("GetMediaUrl Wx object Error")
		return
	}

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

//收藏
func (this *IndexController) WechatFree() {
	this.Rsp(true, "", "")
}
