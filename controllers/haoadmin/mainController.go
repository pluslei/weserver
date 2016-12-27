package haoadmin

import (
	"github.com/astaxie/beego"
	m "weserver/models"
	p "weserver/src/parameter"
	. "weserver/src/tools"
	//"strconv"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"weserver/src/tools"
)

type MainController struct {
	CommonController
}

type Tree struct {
	Id       int64
	Index    int    `json:"index"`
	Text     string `json:"text"`
	IconCls  string `json:"iconCls"`
	Checked  string `json:"checked"`
	Name     string `json:"name"`
	State    string `json:"state"`
	Children []Tree `json:"menu"`
	Ico      string `json:"ico"`
	Url      string `json:"href"`
}

//首页
func (this *MainController) Index() {
	if this.IsAjax() {
		var (
			online     tools.OnlineUserList
			onlinelist []tools.OnlineUserList
			urolename  []string
		)
		roomlist := make(map[string]tools.Usertitle) //房间对应的用户信息
		//获取所有的房间号
		roominfo, num, _ := m.GetAllRoomDate()
		if num > 0 {
			length := int(num)
			for i := 0; i < length; i++ {
				//用户列表信息
				userroom := make(map[string]tools.Usertitle) //房间对应的用户信息
				jobroom := "coderoom_" + p.Code + "_" + fmt.Sprintf("%d", roominfo[i].RommNumber)
				roomdata, _ := p.Client.Get(jobroom)
				if len(roomdata) > 0 {
					userroom, _ = tools.Jsontoroommap(roomdata)
				}
				for rolval, userId := range userroom {
					if len(userId.Uname) > 0 {
						urolename = append(urolename, rolval)
						roomlist[rolval] = userId
					}
				}
			}
			sort.Strings(urolename)
			length = len(urolename)
			for i := 0; i < length; i++ {
				var onlineval string
				online.Roomid = roomlist[urolename[i]].Roomid                             //房间号
				onlineval = roomlist[urolename[i]].Uname                                  //用户名
				online.Uname = onlineval                                                  //用户名
				online.Procities = roomlist[urolename[i]].Procities                       //省市
				onlineval = roomlist[urolename[i]].Datatime.Format("2006-01-02 15:04:05") //发言时间
				online.Logintime = onlineval                                              //登入时间
				onlinetimevar := time.Now().Unix() - roomlist[urolename[i]].Datatime.Unix()
				timehours := onlinetimevar / 3600
				if timehours < 99 {
					onlineval = fmt.Sprintf("%02d时%02d分%02d秒", timehours, time.Unix(onlinetimevar, 0).Minute(), time.Unix(onlinetimevar, 0).Second()) //在线时长
				} else {
					onlineval = fmt.Sprintf("%d时%02d分%02d秒", timehours, time.Unix(onlinetimevar, 0).Minute(), time.Unix(onlinetimevar, 0).Second()) //在线时长
				}
				online.Onlinetime = onlineval                       //在线时长
				online.Ipaddress = roomlist[urolename[i]].Ipaddress //ip地址
				onlinelist = append(onlinelist, online)
			}
		}
		// json
		data := make(map[string]interface{})
		data["onlinelist"] = onlinelist
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		userinfo := this.GetSession("userinfo")
		if userinfo == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
		}
		this.CommonMenu()
		// verifyuser, _ := m.GetRegStatusUser(1)

		onlineuser, _ := m.GetUserByOnlineDesc()
		countonline := m.CountOnline()
		totaltime, _ := strconv.ParseInt(countonline, 10, 64)
		weeklist := m.CountWeekRegist()
		datearr := make([]string, 0)
		countarr := make([]int64, 0)
		if len(weeklist) > 0 {
			for i := 0; i < len(weeklist); i++ {
				prevalue, _ := strconv.ParseInt(weeklist[i][1].(string), 10, 64)
				datearr = append(datearr, weeklist[i][0].(string))
				countarr = append(countarr, prevalue)
			}
		}
		var showmsg Membermsg
		showmsg = GetMembermsg()
		timehours := totaltime / 3600
		if timehours < 99 {
			showmsg.Totallinetime = fmt.Sprintf("%02d时%02d分%02d秒", timehours, time.Unix(totaltime, 0).Minute(), time.Unix(totaltime, 0).Second()) //在线时长
		} else {
			showmsg.Totallinetime = fmt.Sprintf("%d时%02d分%02d秒", timehours, time.Unix(totaltime, 0).Minute(), time.Unix(totaltime, 0).Second()) //在线时长
		}
		this.Data["datearr"] = datearr
		this.Data["countarr"] = countarr
		this.Data["onlineuser"] = onlineuser
		this.Data["companyinfo"] = showmsg
		this.TplName = "haoadmin/index.html"
	}
}

func (this *MainController) OnlineIndex() {
	if this.IsAjax() {
		var (
			online     tools.OnlineUserList
			onlinelist []tools.OnlineUserList
			urolename  []string
		)
		roomlist := make(map[string]tools.Usertitle) //房间对应的用户信息
		//获取所有的房间号
		roominfo, num, _ := m.GetAllRoomDate()
		if num > 0 {
			length := int(num)
			for i := 0; i < length; i++ {
				//用户列表信息
				userroom := make(map[string]tools.Usertitle) //房间对应的用户信息
				jobroom := "coderoom_" + p.Code + "_" + fmt.Sprintf("%d", roominfo[i].RommNumber)
				roomdata, _ := p.Client.Get(jobroom)
				if len(roomdata) > 0 {
					userroom, _ = tools.Jsontoroommap(roomdata)
				}
				for rolval, userId := range userroom {
					if len(userId.Uname) > 0 {
						urolename = append(urolename, rolval)
						roomlist[rolval] = userId
					}
				}
			}
			sort.Strings(urolename)
			length = len(urolename)
			for i := 0; i < length; i++ {
				var onlineval string
				online.Roomid = roomlist[urolename[i]].Roomid                             //房间号
				onlineval = roomlist[urolename[i]].Uname                                  //用户名
				online.Uname = onlineval                                                  //用户名
				online.Procities = roomlist[urolename[i]].Procities                       //省市
				onlineval = roomlist[urolename[i]].Datatime.Format("2006-01-02 15:04:05") //发言时间
				online.Logintime = onlineval                                              //登入时间
				onlinetimevar := time.Now().Unix() - roomlist[urolename[i]].Datatime.Unix()
				timehours := onlinetimevar / 3600
				if timehours < 99 {
					onlineval = fmt.Sprintf("%02d时%02d分%02d秒", timehours, time.Unix(onlinetimevar, 0).Minute(), time.Unix(onlinetimevar, 0).Second()) //在线时长
				} else {
					onlineval = fmt.Sprintf("%d时%02d分%02d秒", timehours, time.Unix(onlinetimevar, 0).Minute(), time.Unix(onlinetimevar, 0).Second()) //在线时长
				}
				online.Onlinetime = onlineval                       //在线时长
				online.Ipaddress = roomlist[urolename[i]].Ipaddress //ip地址
				onlinelist = append(onlinelist, online)
			}
		}
		// json
		data := make(map[string]interface{})
		data["onlinelist"] = onlinelist
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/public/index")
	}
}

func GetMembermsg() (msg Membermsg) {
	//获取所有的房间号
	roominfo, num, _ := m.GetAllRoomDate()
	msg.Totalroom = num //房间数
	if p.Rediserr == nil && num > 0 {
		length := int(num)
		for i := 0; i < length; i++ {
			//用户列表信息
			userroom := make(map[string]Usertitle) //房间对应的用户信息
			jobroom := "coderoom_" + p.Code + "_" + fmt.Sprintf("%d", roominfo[i].RommNumber)
			roomdata, _ := p.Client.Get(jobroom)
			if len(roomdata) > 0 {
				userroom, _ = Jsontoroommap(roomdata)
				for _, userId := range userroom {
					if len(userId.Uname) > 0 {
						msg.Totalonline++
					}
				}
			}
		}
		count, _ := m.GetUserNumber()
		msg.Totalmembers = count
	}
	return msg
}

//登录
func (this *MainController) Login() {
	isajax := this.GetString("isajax")

	if isajax == "1" {
		username := this.GetString("username")
		password := this.GetString("password")
		user, err := m.CheckLogin(username, password)
		adminUser := beego.AppConfig.String("rbac_admin_user")
		sysconfig, _ := m.GetSysConfig()
		loginsys := sysconfig.LoginSys
		if loginsys == 0 {
			if err == nil {
				this.SetSession("userinfo", user)
				accesslist, _ := m.GetAccessList(user.Id)
				this.SetSession("accesslist", accesslist)
				this.Rsp(true, "登录成功", "/public/index")

				return
			} else {
				this.Rsp(false, err.Error(), "")
				return
			}
		} else {
			if username == adminUser {
				if err == nil {
					this.SetSession("userinfo", user)
					accesslist, _ := m.GetAccessList(user.Id)
					this.SetSession("accesslist", accesslist)
					this.Rsp(true, "登录成功", "/public/index")

					return
				} else {
					this.Rsp(false, err.Error(), "")
					return
				}
			} else {
				this.Ctx.Redirect(302, "/public/index")
			}
		}

	}
	// userinfo := this.GetSession("userinfo")
	// if userinfo != nil {
	// 	this.Ctx.Redirect(302, "/public/index")
	// }
	this.TplName = "haoadmin/login.html"
}

//退出
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/public/login")
}

func (this *MainController) UpdateAdminIndex() {
	this.TplName = "haoadmin/updatepwd.html"
}

//改密码
func (this *MainController) UpdateAdminPwd() {
	var unameid int64
	var uname string
	userInfo := this.GetSession("userinfo")
	if userInfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}
	OldPwd := strings.Trim(this.GetString("oldpassword"), " ")
	NewPwd := strings.Trim(this.GetString("newpassword"), " ")
	ChkNewPwd := strings.Trim(this.GetString("repeatpassword"), " ")
	if len(OldPwd) <= 0 || len(NewPwd) <= 0 || len(ChkNewPwd) <= 0 {
		this.Rsp(false, "填写不完整", "")
		return
	}
	if NewPwd != ChkNewPwd {
		this.Rsp(false, "新密码不一致", "")
		return
	}
	uname = fmt.Sprintf("%s", userInfo.(*m.User).Username)
	unameid = userInfo.(*m.User).Id
	uPassword := new(m.User)
	DBPwd := tools.EncodeUserPwd(uname, OldPwd)
	uPassword.Password = DBPwd
	userpwd, err := m.ReadFieldUser(uPassword, "Password")
	if userpwd == nil {
		this.Rsp(false, "密码不正确", "")
		return
	}
	DBNewPwd := tools.EncodeUserPwd(uname, ChkNewPwd)
	u := new(m.User)
	u.Id = unameid
	u.Password = DBNewPwd

	err = u.UpdateUserFields("Password")
	if err != nil {
		beego.Error(err)
		this.AlertBack("密码修改失败")
		this.Rsp(false, "修改失败", "")
		return
	}

	this.Rsp(true, "修改成功", "")
	this.DelSession("userinfo")
}

//修改密码
func (this *MainController) Changepwd() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}
	oldpassword := this.GetString("oldpassword")
	newpassword := this.GetString("newpassword")
	repeatpassword := this.GetString("repeatpassword")
	if newpassword != repeatpassword {
		this.Rsp(false, "两次输入密码不一致", "")
	}
	user, err := m.CheckLogin(userinfo.(m.User).Username, oldpassword)
	if err == nil {
		userUser := new(m.User)
		userUser.Id = user.Id
		userUser.Password = newpassword
		err := userUser.UpdateUserFields("Password")
		if err == nil {
			this.Rsp(true, "密码修改成功", "")
			return
		} else {
			this.Rsp(false, err.Error(), "")
			return
		}
	}
	this.Rsp(false, "密码有误", "")
}
