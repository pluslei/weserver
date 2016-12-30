package haoadmin

import (
	"fmt"
	"github.com/astaxie/beego"
	"sort"
	"strconv"
	"strings"
	"time"
	m "weserver/models"
	p "weserver/src/parameter"
	"weserver/src/socket"
	"weserver/src/tools"
)

type UserController struct {
	CommonController
}

// 用户管理
func (this *UserController) Index() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")
		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		userlist, count := m.Getuserlist(iStart, iLength, "-Id")
		for _, item := range userlist {
			item["Createtime"] = item["Createtime"].(time.Time).Format("2006-01-02 15:04:05")

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
		this.CommonController.CommonMenu()
		roles, _ := m.GetAllUserRole()
		this.Data["roles"] = roles
		this.TplName = "haoadmin/rbac/user/list.html"
	}
}

// 添加用户
func (this *UserController) AddUser() {
	action := this.GetString("action")
	if action == "add" {
		username := this.GetString("username")
		email := this.GetString("email")
		phone, _ := this.GetInt64("phone")
		nickname := this.GetString("nickname")
		password := this.GetString("password")
		remark := this.GetString("remark")
		status, err := this.GetInt("status")
		if err != nil {
			beego.Error(err)
			return
		}
		regstatus, err := this.GetInt("regstatus")
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
		u := new(m.User)
		u.Username = username
		u.Email = email
		u.Phone = phone
		u.Nickname = nickname
		u.Password = tools.EncodeUserPwd(username, password)
		u.Remark = remark
		u.Status = status
		u.RegStatus = regstatus
		u.Role = &m.Role{Id: role}
		u.Title = &m.Title{Id: title}
		id, err := m.AddUser(u)
		if err == nil && id > 0 {
			this.Alert("用户添加成功", "index")
			return
		} else {
			this.AlertBack("用户添加失败")
			return
		}
	} else {
		this.CommonMenu()
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
		username := this.GetString("username")
		email := this.GetString("email")
		phone, _ := this.GetInt64("phone")
		nickname := this.GetString("nickname")
		password := this.GetString("password")
		remark := this.GetString("remark")
		status, err := this.GetInt("status")
		if err != nil {
			beego.Error(err)
		}
		regstatus, err := this.GetInt("regstatus")
		if err != nil {
			beego.Error(err)
			return
		}
		id, err := this.GetInt64("id")
		if err != nil {
			beego.Error(err)
		}
		role, err := this.GetInt64("role")
		if err != nil {
			beego.Error(err)
		}
		u := new(m.User)
		u.Id = id
		u.Username = username
		u.Email = email
		u.Phone = phone
		u.Nickname = nickname
		u.Password = password
		u.Remark = remark
		u.Status = status
		u.RegStatus = regstatus
		u.Role = &m.Role{Id: role}
		if len(u.Password) > 0 {
			u.Password = tools.EncodeUserPwd(u.Username, u.Password)
			err = u.UpdateUserFields("Username", "Email", "Phone", "Nickname", "Password", "Remark", "Status", "RegStatus", "Role")
		} else {
			err = u.UpdateUserFields("Username", "Email", "Phone", "Nickname", "Remark", "Status", "RegStatus", "Role")
		}
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
		roles, _ := m.GetAllUserRole()
		u := new(m.User)
		u.Id = id
		userList, err := m.ReadFieldUser(u, "Id")
		if userList == nil || err != nil {
			this.Alert("获取用户失败", "index")
			return
		}
		this.CommonMenu()
		this.Data["userList"] = userList
		this.Data["RoleList"] = roles
		this.TplName = "haoadmin/rbac/user/edit.html"
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

	userUser := new(m.User)
	userUser.Id = id
	userUser.Title = &m.Title{Id: value}
	err = userUser.UpdateUserFields("Title")
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
func (this *UserController) Rerifyuser() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")
		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		titlelist, count := m.GetRegStatusUser(1, iStart, iLength, "Id")
		// json
		data := make(map[string]interface{})
		data["aaData"] = titlelist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/rbac/user/verify.html"
	}
}

//在线用户
func (this *UserController) Onlineuser() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")
		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		var (
			urolename []string
			objstart  int
			objend    int
			count     int
		)

		//用户列表信息
		userroom := make(map[string]tools.Usertitle) //房间对应的用户信息
		jobroom := "coderoom_" + beego.AppConfig.String("company") + "_" + beego.AppConfig.String("room")
		roomdata, _ := p.Client.Get(jobroom)
		if len(roomdata) > 0 {
			userroom, _ = tools.Jsontoroommap(roomdata)
		}

		for rolval, userId := range userroom {
			if len(userId.Uname) > 0 {
				urolename = append(urolename, rolval)
			}
		}
		sort.Strings(urolename)
		ulength := len(urolename)
		if iLength == -1 {
			iLength = int64(ulength)
		}
		ipagetotal := ulength / int(iLength)
		if 0 != ulength%int(iLength) {
			ipagetotal = ipagetotal + 1
		}
		if ipagetotal == 1 {
			objstart = 0
			objend = ulength
		} else {
			objstart = int(iStart)
			objend = int(iStart + iLength)
			if objend > ulength {
				objend = ulength
			}
		}

		OnlineUser := make([]*m.User, 0)
		for i := objstart; i < objend; i++ {
			username := userroom[urolename[i]].Uname //用户名
			//user := m.GetUserByUsername(username)
			u := new(m.User)
			u.Username = username

			if userInfo, err := m.ReadFieldUser(u, "Username"); err == nil {
				userInfo.LogintimeStr = userroom[urolename[i]].Datatime.Format("2006-01-02 15:04:05")

				onlinetimevar := time.Now().Unix() - userroom[urolename[i]].Datatime.Unix()
				timehours := onlinetimevar / 3600
				if timehours < 99 {
					userInfo.OnlinetimeStr = fmt.Sprintf("%02d时%02d分%02d秒", timehours, time.Unix(onlinetimevar, 0).Minute(), time.Unix(onlinetimevar, 0).Second()) //在线时长
				} else {
					userInfo.OnlinetimeStr = fmt.Sprintf("%d时%02d分%02d秒", timehours, time.Unix(onlinetimevar, 0).Minute(), time.Unix(onlinetimevar, 0).Second()) //在线时长
				}
				userInfo.Ipaddress = userroom[urolename[i]].Ipaddress //ip地址

				if userInfo.Title.Id == 0 {
					userInfo.Titlename = "未知头衔"
				} else {
					titleinfo, _ := m.ReadTitleById(userInfo.Title.Id)
					if err != nil {
						beego.Error(err)
						userInfo.Titlename = "未知头衔"
					} else {
						userInfo.Titlename = titleinfo.Name
					}
				}

				OnlineUser = append(OnlineUser, userInfo)
			}
		}

		beego.Debug("OnlineUser the ", OnlineUser, "online statuce:", userroom)
		count = ulength

		// json
		data := make(map[string]interface{})
		data["aaData"] = OnlineUser
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

// 用户审核
func (this *UserController) UpdateRegStatus() {
	Id, _ := this.GetInt64("Id")
	u := new(m.User)
	u.Id = Id
	u.RegStatus = 2
	err := u.UpdateUserFields("RegStatus")
	if err != nil {
		beego.Error(err)
		this.Rsp(false, "审核失败", "")
		return
	} else {
		this.Rsp(true, "审核成功", "")
	}
}

// 用户状态
func (this *UserController) UpdateUserStatus() {
	userid, _ := this.GetInt64("Id")
	usr := new(m.User)
	usr.Id = userid
	user, err := m.ReadFieldUser(usr, "Id")
	if err != nil {
		this.Rsp(false, "状态改变失败", "")
	} else {
		if this.changeuserstatus(user) {
			this.Rsp(true, "修改成功", "")
		} else {
			this.Rsp(false, "修改失败", "")
		}
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
		if socket.KictUser(user.Username) {
			this.Rsp(true, "踢出成功", "")
		} else {
			this.Rsp(false, "踢出失败", "")
		}
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
		beego.Error(err)
		return false
	} else {
		return true
	}
	return false
}
