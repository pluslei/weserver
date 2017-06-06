package haoadmin

import (
	"strconv"
	"strings"
	"time"
	m "weserver/models"
	"weserver/src/tools"

	"github.com/astaxie/beego"
)

type UserController struct {
	CommonController
}

// 用户管理
func (this *UserController) Index() {
	if this.IsAjax() {
		user := this.GetSession("userinfo").(*m.User)
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")
		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		nickname := this.GetString("sSearch_0")
		roleIdStr := this.GetString("sSearch_5")
		titleIdStr := this.GetString("sSearch_6")
		var roleId int64 = -1
		var titleId int64 = -1
		if len(roleIdStr) > 0 {
			roleId, _ = strconv.ParseInt(roleIdStr, 10, 64)
		}
		if len(titleIdStr) > 0 {
			titleId, _ = strconv.ParseInt(titleIdStr, 10, 64)
		}
		companyId := user.CompanyId
		userlist, count := m.GetWechatUserList(iStart, iLength, "-Id", nickname, companyId, roleId, titleId)
		for _, item := range userlist {
			item["Lastlogintime"] = item["Lastlogintime"].(time.Time).Format("2006-01-02 15:04:05")

			if item["Title"] == 0 {
				item["Titlename"] = "未知头衔"
			} else {
				titleinfo, _ := m.ReadTitleById(item["Title"].(int64))
				if err != nil {
					beego.Error(err)
					item["Titlename"] = "未知头衔"
				} else {
					item["Titlename"] = titleinfo.Name
				}
			}
			roleInfo, err := m.GetRoleInfoById(item["Role"].(int64))
			if err != nil {
				item["RoleName"] = "未知角色"
			} else {
				item["RoleName"] = roleInfo.Title
			}
			roomInfo, err := m.GetRoomInfoByRoomID(item["Room"].(string))
			if err != nil {
				item["RoomName"] = "未知房间"
			} else {
				item["RoomName"] = roomInfo.RoomTitle
			}
			Info, err := m.GetCompanyById(item["CompanyId"].(int64))
			if err != nil {
				item["CompanyName"] = "未知公司"
			} else {
				item["CompanyName"] = Info.Company
			}
			itemUserName := item["Username"].(string)
			if len(itemUserName) <= 0 {
				item["UserAccount"] = ""
			} else {
				theUserInfo, err := m.GetUserByUsername(itemUserName)
				if err != nil {
					item["UserAccount"] = ""
					beego.Info("err:", err)
				} else {
					if len(theUserInfo.Account) <= 0 {
						item["UserAccount"] = ""
					} else {
						item["UserAccount"] = theUserInfo.Account
					}
				}

			}

		}

		// json
		data := make(map[string]interface{})
		data["aaData"] = userlist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		user := this.GetSession("userinfo")
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		username := user.(*m.User).Username
		this.Data["username"] = username
		this.CommonController.CommonMenu()
		roles, _ := m.GetAllUserRole()
		this.Data["roles"] = roles
		titles, _ := m.GetAllUserTitle()
		this.Data["titles"] = titles
		this.TplName = "haoadmin/rbac/user/reglist.html"
	}
}

// 用户设置列表
func (this *UserController) UserList() {
	if this.IsAjax() {
		user := this.GetSession("userinfo").(*m.User)
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")
		if err != nil {
			beego.Debug("userlist get", err)
			return
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Debug("userlist Error", err)
			return
		}
		nickname := this.GetString("sSearch_0")

		companyId := user.CompanyId
		userlist, count := m.Getuserlist(iStart, iLength, "-Id", nickname, companyId)
		for _, item := range userlist {
			if item["Lastlogintime"] == nil {
				item["Lastlogintime"] = "未获取到时间"
			} else {
				item["Lastlogintime"] = item["Lastlogintime"].(time.Time).Format("2006-01-02 15:04:05")
			}
			if item["Title"] == 0 {
				item["Titlename"] = "未知头衔"
			} else {
				titleinfo, _ := m.ReadTitleById(item["Title"].(int64))
				if err != nil {
					beego.Error(err)
					item["Titlename"] = "未知头衔"
				} else {
					item["Titlename"] = titleinfo.Name
				}
			}
			roleInfo, err := m.GetRoleInfoById(item["Role"].(int64))
			if err != nil {
				item["Rolename"] = "未知角色"
			} else {
				item["Rolename"] = roleInfo.Title
			}
			Info, err := m.GetCompanyById(item["CompanyId"].(int64))
			if err != nil {
				item["CompanyName"] = "未知公司"
			} else {
				item["CompanyName"] = Info.Company
			}
		}
		// json
		data := make(map[string]interface{})
		data["aaData"] = userlist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		user := this.GetSession("userinfo")
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		username := user.(*m.User).Username
		this.Data["username"] = username
		this.CommonController.CommonMenu()
		roles, _ := m.GetAllUserRole()
		this.Data["roles"] = roles
		this.TplName = "haoadmin/rbac/user/list.html"
	}
}

// 设置用户名
func (this *UserController) SetUsername() {
	action := this.GetString("action")
	id, err := this.GetInt64("Id")
	if err != nil {
		beego.Error("get the id", err)
		return
	}
	if action == "set" {
		userInfo := new(m.User)
		userInfo.Id = id
		userLoad, err := m.LoadRelatedUser(userInfo, "Id")
		if err != nil {
			beego.Error("load retalteduser error", err)
		}

		Password := this.GetString("Password")
		role, _ := this.GetInt64("role")
		title, _ := this.GetInt64("title")
		Nickname := this.GetString("nickname")

		u := new(m.User)
		u.Id = id
		u.Role = &m.Role{Id: role}
		u.Title = &m.Title{Id: title}
		u.Nickname = Nickname
		if len(Password) > 0 {
			beego.Debug("user===", userLoad.Account, Password)
			u.Password = tools.EncodeUserPwd(userLoad.Account, Password)
			err = u.UpdateUserFields("Role", "Title", "Password", "Nickname")
		}
		err = u.UpdateUserFields("Role", "Title", "Nickname")
		if err != nil {
			beego.Error(err)
			this.AlertBack("密码修改失败")
			this.Rsp(false, "修改失败", "")
			return
		}
		this.Alert("用户修改成功", "usersetlist")

		roomId := this.GetStrings("RoomId") //添加进入房间权限
		for _, val := range roomId {
			//判断用户是否申请房间
			bl := m.CheckRegistApply(val, userLoad.Username)
			if !bl {
				reg := new(m.Regist)
				reg.Room = val
				reg.UserId = id
				reg.Nickname = Nickname
				reg.RegStatus = 2
				reg.CompanyId = userLoad.CompanyId
				reg.Username = userLoad.Username
				reg.Role = &m.Role{Id: role}
				reg.Title = &m.Title{Id: title}
				reg.Lastlogintime = time.Now()
				reg.UserIcon = userLoad.UserIcon
				_, err := m.AddRegist(reg)
				if err != nil {
					beego.Error("add regist error", err)
					this.AlertBack("授权进入房间失败")
					this.Rsp(false, "授权进入房间失败", "")
					return
				}
			} else {
				_, err := m.UpdateRegistStatus(val, userLoad.Username, 2)
				if err != nil {
					beego.Debug("udpate regist status error")
					this.AlertBack("更新房间状态失败")
					this.Rsp(false, "更新房间状态失败", "")
					return
				}
			}

		}
		return
	} else {
		userInfo := new(m.User)
		userInfo.Id = id
		userLoad, err := m.LoadRelatedUser(userInfo, "Id")
		if err != nil {
			beego.Error("load retalteduser error", err)
		}

		roles, _ := m.GetAllUserRole()
		titles, _ := m.GetAllUserTitle()
		this.CommonMenu()
		this.Data["userList"] = userLoad
		this.Data["RoleList"] = roles
		this.Data["TitleList"] = titles
		this.TplName = "haoadmin/rbac/user/setusername.html"
	}
}

//解除禁言
func (this *UserController) SetUnShutUp() {

}

// 添加用户
func (this *UserController) AddUser() {
	action := this.GetString("action")
	if action == "add" {
		account := this.GetString("account")
		if m.CheckAccountIsExist(account) {
			this.AlertBack("用户名已存在")
			return
		}
		email := this.GetString("email")
		phone, _ := this.GetInt64("phone")
		nickname := this.GetString("nickname")
		password := this.GetString("password")
		remark := this.GetString("remark")
		status, err := this.GetInt("status")
		userIcon := this.GetString("userIconFile")
		if err != nil {
			beego.Error(err)
			return
		}
		companyId, err := this.GetInt64("company")
		if err != nil {
			beego.Error(err)
			return
		}
		role, err := this.GetInt64("role")
		if err != nil {
			beego.Error(err)
			return
		}
		title, err := this.GetInt64("title")
		if err != nil {
			beego.Error(err)
			return
		}

		//urlImage := beego.AppConfig.String("httplocalServerAdress") + "/static/img/defaultIco.png"
		uuid := tools.GetGuid()
		u := new(m.User)
		u.Account = account
		u.Username = uuid
		u.Email = email
		u.Phone = phone
		u.Nickname = nickname
		u.Password = tools.EncodeUserPwd(account, password)
		u.Remark = remark
		u.Status = status
		u.Headimgurl = userIcon
		u.RegStatus = 2
		u.CompanyId = companyId
		u.Role = &m.Role{Id: role}
		u.Title = &m.Title{Id: title}
		u.Lastlogintime = time.Now()
		u.UserIcon = userIcon
		id, err := m.AddUser(u)
		if err == nil && id > 0 {
			this.Alert("用户添加成功", "index")
			roomId := this.GetStrings("RoomId")
			for _, val := range roomId {
				reg := new(m.Regist)
				reg.Room = val
				reg.UserId = id
				reg.Username = uuid
				reg.Nickname = u.Nickname
				reg.RegStatus = 2
				reg.CompanyId = companyId
				reg.Role = &m.Role{Id: role}
				reg.Title = &m.Title{Id: title}
				reg.Lastlogintime = time.Now()
				reg.UserIcon = userIcon
				_, err := m.AddRegist(reg)
				if err != nil {
					beego.Error("add regist error", err)
				}
			}
			return
		} else {
			beego.Error("add user error", err)
			this.AlertBack("用户添加失败")
			return
		}
	} else {
		this.CommonMenu()
		user := this.GetSession("userinfo").(*m.User)
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		companyId := user.CompanyId

		companyList, _, err := m.GetCompanyList(companyId)
		if err != nil {
			beego.Error("get the companyList error", err)
			return
		}

		roonInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Error("Get the Roominfo error", err)
			return
		}

		this.Data["CompanyList"] = companyList
		this.Data["roonInfo"] = roonInfo
		roles, _ := m.GetAllUserRole()
		titles, _ := m.GetAllUserTitle()
		this.Data["RoleList"] = roles
		this.Data["TitleList"] = titles
		this.TplName = "haoadmin/rbac/user/add.html"
	}
}

// 更新用户
func (this *UserController) UpdateUser() {
	action := this.GetString("action")
	if action == "edit" {
		id, err := this.GetInt64("id")
		if err != nil {
			beego.Error(err)
			return
		}
		role, _ := this.GetInt64("role")
		title, _ := this.GetInt64("title")
		regstatus, err := this.GetInt("regstatus")
		if err != nil {
			beego.Error(err)
			return
		}
		_, err = m.UpdateWechatUserInfo(id, role, title, regstatus)
		if err == nil {
			this.Alert("用户更新成功", "index")
			return
		} else {
			this.AlertBack("用户更新失败")
			return
		}
	} else {
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			return
		}

		userList, err := m.GetWechatUserInfoById(id)
		if err != nil {
			this.Alert("获取用户失败", "index")
			return
		}
		roles, _ := m.GetAllUserRole()
		titles, _ := m.GetAllUserTitle()
		this.CommonMenu()
		if userList.Username != "" {
			this.Data["username"] = userList.Username
		} else {
			userInfo, err := m.GetUserInfoById(userList.UserId)
			if err != nil {
				this.Data["username"] = userList.Username
			} else {
				this.Data["username"] = userInfo.Username
			}
		}
		this.Data["userList"] = userList
		this.Data["RoleList"] = roles
		this.Data["TitleList"] = titles
		this.TplName = "haoadmin/rbac/user/regedit.html"
	}
}

// 删除房间用户
func (this *UserController) DelRegistUser() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelRegistUserById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

// 批量删除房间用户
func (this *UserController) PrepareDelRegistUser() {
	IdArray := this.GetString("Id")
	var idarr []int64
	if len(IdArray) > 0 {
		preValue := strings.Split(IdArray, ",")
		for _, v := range preValue {
			id, _ := strconv.ParseInt(v, 10, 64)
			idarr = append(idarr, id)

		}
	}
	status, err := m.PrepareDelReisterUser(idarr)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

// 删除用户
func (this *UserController) DelUser() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelUserById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

// 批量删除用户
func (this *UserController) PrepareDelUser() {
	IdArray := this.GetString("Id")
	var idarr []int64
	if len(IdArray) > 0 {
		preValue := strings.Split(IdArray, ",")
		for _, v := range preValue {
			id, _ := strconv.ParseInt(v, 10, 64)
			idarr = append(idarr, id)

		}
	}
	status, err := m.PrepareDelUser(idarr)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

// 用户赋予角色
func (this *UserController) UserToRole() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Error(err)
	}
	value, err := this.GetInt64("value")
	if err != nil {
		beego.Error(err)
	}
	user := new(m.User)
	user.Id = id
	user.Role = &m.Role{Id: value}
	err = user.UpdateUserFields("Role")
	if err != nil {
		this.Rsp(false, "更新失败", "")
	} else {
		RoleList, _ := m.GetRoleInfoById(value)
		this.Ctx.WriteString(RoleList.Title)
	}
}

// 用户赋予头衔
func (this *UserController) SetUserTitle() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Error(err)
	}
	value, err := this.GetInt64("value")
	if err != nil {
		beego.Error(err)
	}

	_, err = m.UpdateWechatUserTitle(id, value)
	if err != nil {
		beego.Error(err.Error())
		this.Rsp(false, "更新失败", "")
	} else {
		titleInfo, _ := m.ReadTitleById(value)
		this.Ctx.WriteString(titleInfo.Name)
	}
}

// 用户赋予头衔
func (this *UserController) UserToTitle() {
	action := this.GetString("action")
	if action == "add" {
		titleid, _ := this.GetInt64("titleid")
		userid, err := this.GetInt64("userid")
		if err != nil {
			beego.Error(err)
		}

		// err = m.DelUserTitle(userid)
		// if err != nil {
		// 	this.AlertBack("删除用户权限错误")
		// 	return
		// }
		// if len(titleid) > 0 {
		// 	for _, v := range titleid {
		// 		id, _ := strconv.Atoi(v)
		// 		_, err := m.AddUserTitle(userid, int64(id))
		// 		if err != nil {
		// 			this.AlertBack("添加错误")
		// 			return
		// 		}
		// 	}
		// }
		userUser := new(m.User)
		userUser.Id = userid
		userUser.Title = &m.Title{Id: titleid}
		err1 := userUser.UpdateUserFields("Title")
		if err1 != nil {
			beego.Error(err1.Error())
		}
		this.Alert("用户头衔添加成功", "index")
	} else {
		this.CommonMenu()
		userid, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
		}
		usr := new(m.User)
		usr.Id = userid
		userlist, err := m.ReadFieldUser(usr, "Id")
		isRoleList, _ := m.GetTitleByUserId(userid)
		if err != nil {
			beego.Error(err)
		}
		titleList := m.TitleList()
		for _, item := range titleList {
			for _, kitem := range isRoleList {
				if item["Id"] == kitem["Id"] {
					item["Checked"] = "checked"
				}
			}
		}
		this.Data["isRoleList"] = isRoleList
		this.Data["titleList"] = titleList
		this.Data["userlist"] = userlist
		this.TplName = "haoadmin/rbac/user/usertorole.html"
	}

}

/*
//在线用户
func (this *UserController) Onlineuser() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		var (
			count int
		)
		// json
		data := make(map[string]interface{})
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonController.CommonMenu()
		//获取所有的房间号
		roominfo, _, _ := m.GetAllRoomDate()
		this.Data["roomnum"] = roominfo
		this.TplName = "haoadmin/rbac/user/online.html"
	}
}
*/

// 用户审核
func (this *UserController) UpdateRegStatus() {
	id, _ := this.GetInt64("id")
	status, _ := this.GetInt("status")
	_, err := m.UpdateWechtUserStatus(id, status)
	if err != nil {
		beego.Debug("udpate status error", err, "id=", id, "status=", status)
		this.Rsp(false, "状态改变失败", "")
	} else {
		this.Rsp(true, "修改成功", "")
	}
}

// 更改用户状态
func (this *UserController) UpdateStatus() {
	id, _ := this.GetInt64("id")
	status, _ := this.GetInt("status")
	usr := new(m.User)
	usr.Id = id
	user, _ := m.ReadFieldUser(usr, "Id")
	user.Status = status
	if this.changeuserstatus(user) {
		this.Rsp(true, "修改成功", "")
	} else {
		beego.Debug("udpate status error id=", id, "status=", status)
		this.Rsp(false, "状态改变失败", "")

	}
}

// 踢出房间
// 房间踢出失败原因可能人不再map里面
func (this *UserController) KictUser() {
	userid, _ := this.GetInt64("Id")
	usr := new(m.User)
	usr.Id = userid
	user, err := m.ReadFieldUser(usr, "Id")
	if err != nil {
		this.Rsp(false, "踢出失败", "")
		return
	}
	if user.Status == 1 {
		this.Rsp(false, "用户已未禁用状态", "")
		return
	}
	if user.RegStatus == 1 {
		this.Rsp(false, "用户暂未审核", "")
		return
	}
	if this.changeuserstatus(user) {
		this.Rsp(true, "踢出成功", "")
	} else {
		this.Rsp(false, "踢出失败", "")
	}
}

func (this *UserController) changeuserstatus(user *m.User) bool {
	if user.Status == 1 {
		user.Status = 2
	} else {
		user.Status = 1
	}
	err := user.UpdateUserFields("Status")
	if err != nil {
		beego.Error("update the status is error:", err)
		return false
	} else {
		return true
	}
	return false
}
