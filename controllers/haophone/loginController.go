package haophone

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"regexp"
	"strconv"
	"strings"
	"time"
	m "weserver/models"
	"weserver/src/tools"
)

type LoginController struct {
	CommonController
}

const (
	regularusername = `^[a-zA-Z]\w{3,15}$`                                 //长度在6-18之间，只能包含字符、数字和下划线
	regularphone    = `^(13[0-9]|14[57]|15[0-35-9]|18[0-9]|17[0-9])\d{8}$` //11位的纯数字
	regularemail    = `^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`                    //邮箱验证
	regularqq       = `^[1-9]{1}[0-9]{4,15}$`                              //qq
)

type OnUserMsg struct {
	Exist     bool   //用户名是否存在
	Uname     string //用户名
	Phone     string //手机号码
	AecPhone  string //加密的手机号码
	PhoneCode string //验证码
	ShowMsg   string //返回的信息
}

var FaceImg = beego.AppConfig.String("faceImg")

func (this *LoginController) PhoneIndex() {
	if this.IsAjax() {
		data := make(map[string]interface{})
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		roomdata := this.GetString("roomid")
		roomid, _ := strconv.ParseInt(roomdata, 10, 64)
		roomInfo, err := m.GetRoomNumber(roomid)
		if err != nil {
			beego.Error(err)
			this.RedirectRoom()
			return
		} else {
			type Userinfor struct {
				Codeid    string //房间号公司代码加密
				Uname     string //用户名
				UserIcon  string //logo
				RoleName  string //用户角色
				Titlerole string //用户类型
				Authorcss string //头衔
				Insider   int64  //1内部人员或0外部人员
				IsLogin   bool   //是否登入
			}
			companyid := roomInfo.CompanyCode
			temp := fmt.Sprintf("%d", roomid)
			prevalue := companyid + "_" + temp
			codeid := tools.MainEncrypt(prevalue)
			var user Userinfor
			user.Codeid = codeid
			user.Uname = ""
			user.IsLogin = false
			user.RoleName = "guest" //角色名称
			//查询数据库获取用户类型
			fieldrole := new(m.Role)
			fieldrole.Name = "guest"
			fieldrole, err = m.ReadFieldRole(fieldrole, "Name")
			if err == nil {
				user.Titlerole = fieldrole.Title //用户类型
			} else {
				user.Titlerole = "游客" //用户类型
			}
			user.Authorcss = "/upload/usertitle/visitor.png"
			user.UserIcon = "/upload/usericon/icon.png"
			user.Insider = 0 //1内部人员或0外部人员
			if this.GetSession("indexUserInfo") != nil {
				Info := this.GetSession("indexUserInfo")
				userinfor := new(m.User)
				userinfor.Username = Info.(*m.User).Username
				userUsername, err := m.LoadRelatedUser(userinfor, "Username")
				if err == nil {
					body, err := json.Marshal(userUsername)
					if err == nil {
						js, _ := simplejson.NewJson(body)
						user.Uname = userinfor.Username
						Rolemap := js.Get("Role").MustMap()
						rolename := fmt.Sprintf("%s", Rolemap["Name"])      //角色名称
						rolename = strings.Replace(rolename, " ", "", -1)   //去空格
						utype := fmt.Sprintf("%s", Rolemap["Title"])        //用户类型
						uinsider := fmt.Sprintf("%s", Rolemap["IsInsider"]) //1内部人员或0外部人员
						insider, _ := strconv.ParseInt(uinsider, 10, 64)
						icon := js.Get("UserIcon").MustString()
						//改动的
						Titlemap := js.Get("Title").MustMap()
						ucss := "/upload/usertitle/" + fmt.Sprintf("%s", Titlemap["Css"])
						user.RoleName = rolename
						user.Titlerole = utype
						user.UserIcon = "/upload/usericon/" + icon
						user.Authorcss = ucss
						user.Insider = insider
						user.IsLogin = true
					}
				}
			}
			this.Data["user"] = user
			this.Data["actionid"] = roomInfo.ActivityId
		}
		this.TplName = "dist/index.html"
	}
}

func (this *LoginController) PhoneLogin() {
	Username := this.GetString("username")
	Password := this.GetString("password")
	if len(Username) <= 0 || len(Password) <= 0 {
		this.Rsp(false, "请确认数据填写完成", "")
		return
	}
	user := new(m.User)
	if regphone := regexp.MustCompile(regularphone); regphone.MatchString(Username) { //手机号登录
		PhoneInt, err := strconv.ParseInt(Username, 10, 64)
		if err != nil {
			this.Rsp(false, "用户名登录异常", "")
			return
		}
		user.Phone = PhoneInt
		userUsername, err := m.ReadFieldUser(user, "Phone")
		if userUsername == nil || err != nil {
			this.Rsp(false, "用户名或者密码错误", "")
			return
		} else {
			Cpassword := tools.EncodeUserPwd(userUsername.Username, Password)
			user.Password = Cpassword
			userUsername, err := m.LoadRelatedUser(user, "Phone", "Password")
			if userUsername == nil || err != nil {
				this.Rsp(false, "用户名或者密码错误", "")
				return
			} else if userUsername.Status > 1 {
				this.Rsp(false, "用户状态异常", "")
				return
			} else if userUsername.RegStatus != 2 {
				this.Rsp(false, "审核未通过,请联系在线客服.", "")
			} else if userUsername.Id > 0 {
				roleId := strconv.FormatInt(userUsername.Role.Id, 10)
				this.Rsp(true, userUsername.Username, roleId)
				this.SetSession("indexUserInfo", userUsername)
				this.DelSession("xsrf_token")
				return
			} else {
				this.Rsp(false, "用户登录失败", "")
			}
		}
	} else { //用户登录
		user.Username = Username
		user.Password = tools.EncodeUserPwd(Username, Password)
		userUsername, err := m.LoadRelatedUser(user, "Username", "Password")
		if userUsername == nil || err != nil {
			this.Rsp(false, "用户名或者密码错误", "")
			return
		} else if userUsername.Status > 1 {
			this.Rsp(false, "用户状态异常", "")
		} else if userUsername.RegStatus != 2 {
			this.Rsp(false, "审核未通过,请联系在线客服.", "")
		} else if userUsername.Id > 0 {
			roleId := strconv.FormatInt(userUsername.Role.Id, 10)
			this.Rsp(true, userUsername.Username, roleId)
			this.SetSession("indexUserInfo", userUsername)
			this.DelSession("xsrf_token")
			return
		} else {
			this.Rsp(false, "用户登录失败", "")
		}
	}
}

// 注销登录
func (this *LoginController) PhoneLogout() {
	this.DelSession("indexUserInfo")
	this.Ctx.Redirect(302, "/")
}

//发送验证码
func (this *LoginController) SendCode() {
	if this.IsAjax() {
		var usermsg OnUserMsg
		uname := this.GetString("username")
		usermsg.Uname = uname
		uphone := this.GetString("userphone")
		userphonereg, _ := regexp.Match(regularphone, []byte(uphone))
		if !userphonereg {
			usermsg.Exist = false
			usermsg.ShowMsg = "手机号格式不正确"
		} else {
			usermsg.AecPhone = uphone
			phone, _ := strconv.ParseInt(uphone, 10, 64)
			user, _ := m.GetUserNameByPhone(uname, phone)
			if uname == "" {
				usermsg.Exist = false
				usermsg.ShowMsg = "用户名不能为空"
			}
			if uphone == "" {
				usermsg.Exist = false
				usermsg.ShowMsg = "手机号不能为空"
			}
			if user.Id == 0 {
				countPhone, _ := m.CheckPhoneDay(phone)
				if countPhone >= 3 {
					usermsg.ShowMsg = "当前手机验证超过限制"
				} else {
					Ip := this.GetClientip()
					validatacode := tools.RandomAlphaOrNumeric(6, false, true)
					if tools.GetPhoneCode(phone, validatacode) {
						validata := new(m.ValidataCode)
						validata.Code = validatacode
						validata.Phone = phone
						validata.Ip = Ip
						validata.Times = time.Now().Unix()
						validata.Timeday = time.Now().Format("2006-01-02")
						_, err := m.InsertCode(validata)
						if err != nil {
							usermsg.ShowMsg = "验证码发送失败"
						} else {
							usermsg.PhoneCode = tools.MainEncrypt(validatacode)
							usermsg.ShowMsg = "验证码发送成功"
						}
					} else {
						usermsg.ShowMsg = "验证码发送失败"
					}
				}
			} else {
				usermsg.ShowMsg = "验证码发送失败"
			}
		}
		data := make(map[string]interface{})
		data["msg"] = usermsg
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

//获取客户的真是IP地址
func (this *LoginController) GetClientip() string {
	var addrArr []string
	if len(this.Ctx.Request.Header.Get("X-Forwarded-For")) > 0 {
		addr := this.Ctx.Request.Header.Get("X-Forwarded-For")
		addrArr = strings.Split(addr, ":")
	} else if len(this.Ctx.Request.RemoteAddr) > 0 {
		addr := this.Ctx.Request.RemoteAddr
		addrArr = strings.Split(addr, ":")
	} else {
		addrArr[0] = "127.0.0.1"
	}
	beego.Debug(addrArr)
	return addrArr[0]
}

// 注册
func (this *LoginController) PhoneRegister() {
	username := this.GetString("username")
	userphone := this.GetString("userphone")
	password := this.GetString("password")
	repeatpassword := this.GetString("repeatpassword")
	code := this.GetString("validatacode")
	userphonereg, _ := regexp.Match(regularphone, []byte(userphone))
	if !userphonereg {
		this.Rsp(false, "手机号格式错误", "")
		return
	}
	PhoneInt, err := strconv.ParseInt(userphone, 10, 64)
	if err != nil {
		this.Rsp(false, "手机号异常", "")
		return
	}
	validata, err := m.CheckCode(PhoneInt, code)
	if err != nil {
		this.Rsp(false, "手机验证码有误", "")
		return
	}
	nowTime := time.Now().Unix()
	if validata.Times < nowTime-15*60 {
		this.Rsp(false, "验证码过期", "")
		return
	}
	num, err := m.UpdateValidataCode(validata.Id)
	if err != nil || num <= 0 {
		this.Rsp(false, "验证码失效", "")
		return
	}
	if password != repeatpassword {
		this.Rsp(false, "两次输入密码不一致", "")
		return
	}
	uUsername := new(m.User)
	uUsername.Username = username
	userUsername, err := m.ReadFieldUser(uUsername, "Username")
	if userUsername != nil && err == nil {
		this.Rsp(false, "用户名已使用", "")
		return
	}
	config, _ := m.GetSysConfig()
	configRole := config.Registerrole
	configTitle := config.Registertitle
	configVerify := config.Verify
	u := new(m.User)
	u.Username = username
	u.Password = tools.EncodeUserPwd(username, password)
	u.Status = 1
	if configVerify == 0 {
		u.RegStatus = 1
	} else {
		u.RegStatus = 2
	}

	u.Phone = PhoneInt
	u.UserIcon = "haoyue.png"
	u.Role = &m.Role{Id: configRole}
	u.Title = &m.Title{Id: configTitle}
	userid, err := m.AddUser(u)
	if err == nil && userid > 0 {
		this.Rsp(true, "注册成功,请登录！", "/")
	} else {
		beego.Error(err)
		this.Rsp(false, "用户注册失败", "")
		return
	}
}

// 获取所有分组
func (this *LoginController) GetPhoneGroups() {
	if this.IsAjax() {
		// 获取所有的分组
		groups, _ := m.GetGroupList()
		length := len(groups)
		for i := 0; i < length; i++ {
			groups[i].GroupFace = FaceImg + groups[i].GroupFace
		}
		this.Data["json"] = groups
		this.ServeJSON()
	} else {
		this.TplName = "dist/index.html"
	}
}

// 根据分组查询表情
func (this *LoginController) GetPhoneFaces() {
	if this.IsAjax() {
		group, _ := this.GetInt64("group")
		// 根据分组查询每组的表情
		faces, err := m.GetFaceByGroup(group)
		for _, v := range faces {
			v["Url"] = FaceImg + v["Url"].(string)
			v["GroupFace"] = FaceImg + v["GroupFace"].(string)
		}
		if err != nil {
			beego.Error(err)
		}
		this.Data["json"] = faces
		this.ServeJSON()
	} else {
		this.TplName = "dist/index.html"
	}
}
