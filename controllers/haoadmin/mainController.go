package haoadmin

import (
	m "weserver/models"
	"weserver/src/tools"
	. "weserver/src/tools"

	"github.com/astaxie/beego"
	//"strconv"
	"fmt"
	"strconv"
	"strings"
	//"time"
)

type MainController struct {
	CommonController
}

//首页
func (this *MainController) Index() {
	beego.Debug("===index")
	if this.IsAjax() {
		// json
		data := make(map[string]interface{})
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		userinfo := this.GetSession("userinfo")
		if userinfo == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		this.CommonMenu()
		beego.Debug("userinfo", userinfo)
		// verifyuser, _ := m.GetRegStatusUser(1)
		onlineuser, _ := m.GetUserByOnlineDesc()
		//countonline := m.CountOnline()
		//totaltime, _ := strconv.ParseInt(countonline, 10, 64)
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
		this.Data["datearr"] = datearr
		this.Data["countarr"] = countarr
		this.Data["onlineuser"] = onlineuser
		this.Data["companyinfo"] = showmsg
		this.TplName = "haoadmin/index.html"
	}
}

func (this *MainController) OnlineIndex() {
	if this.IsAjax() {
		// json
		data := make(map[string]interface{})
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/public/index")
	}
}

//登录
func (this *MainController) Login() {
	isajax := this.GetString("isajax")
	if isajax == "1" {
		var admin m.User
		username := this.GetString("username")
		password := this.GetString("password")
		md5password := tools.EncodeUserPwd(username, password)
		admin, err := m.GetUserInfoByUsername(username, md5password)
		if err != nil {
			admin, err = m.GetUserInfoByAccount(username, md5password)
			if err != nil {
				this.Rsp(false, "用户名或密码不正确", "")
				return
			}
		}
		beego.Debug("user", admin)

		Superadmin := beego.AppConfig.String("rbac_admin_user")

		sysconfig, _ := m.GetSysConfig()
		loginsys := sysconfig.LoginSys
		beego.Debug("user", loginsys, username, Superadmin, admin.Role.IsInsider)

		if admin.Role.Id != 1 || admin.Title.Id != 1 {
			this.Rsp(true, "无登录权限", "/weserver/public/login")
			return
		}
		if loginsys == 0 {
			this.SetSession("userinfo", &admin)
			accesslist, _ := m.GetAccessList(admin.Id)
			this.SetSession("accesslist", accesslist)
			this.Rsp(true, "登录成功", "/weserver/public/index")
			return
		} else {
			if Superadmin == admin.Username && admin.Role.IsInsider == 1 {
				this.SetSession("userinfo", &admin)
				accesslist, _ := m.GetAccessList(admin.Id)
				this.SetSession("accesslist", accesslist)
				this.Rsp(true, "登录成功", "/weserver/public/index")
				return
			} else {
				this.Ctx.Redirect(302, "/weserver/public/index")
			}
		}

	}
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
		return
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
	DBPwd := EncodeUserPwd(uname, OldPwd)
	uPassword.Password = DBPwd
	userpwd, err := m.ReadFieldUser(uPassword, "Password")
	if userpwd == nil {
		this.Rsp(false, "密码不正确", "")
		return
	}
	DBNewPwd := EncodeUserPwd(uname, ChkNewPwd)
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
		return
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
