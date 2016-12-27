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
			if item["Role"] == 0 {
				item["UserRole"] = "未知角色1"
			} else {
				rolelist, _ := m.GetRoleInfoById(item["Role"].(int64))
				if err != nil {
					beego.Error(err)
					item["UserRole"] = "未知角色2"
				} else {
					item["UserRole"] = rolelist.Title
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
		sroomnumber := this.GetString("sSearch_0")
		roomnumber, _ := strconv.ParseInt(sroomnumber, 10, 64) //房间号
		var (
			onlinelist []tools.OnlineUserList
			online     tools.OnlineUserList
			urolename  []string
			objstart   int
			objend     int
			count      int
		)
		if roomnumber == 0 {
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
				for i := objstart; i < objend; i++ {
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
				count = ulength
			}
		} else {
			//用户列表信息
			userroom := make(map[string]tools.Usertitle) //房间对应的用户信息
			jobroom := "coderoom_" + p.Code + "_" + sroomnumber
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
			for i := objstart; i < objend; i++ {
				var onlineval string
				online.Roomid = userroom[urolename[i]].Roomid                             //房间号
				onlineval = userroom[urolename[i]].Uname                                  //用户名
				online.Uname = onlineval                                                  //用户名
				online.Procities = userroom[urolename[i]].Procities                       //省市
				onlineval = userroom[urolename[i]].Datatime.Format("2006-01-02 15:04:05") //发言时间
				online.Logintime = onlineval                                              //登入时间
				onlinetimevar := time.Now().Unix() - userroom[urolename[i]].Datatime.Unix()
				timehours := onlinetimevar / 3600
				if timehours < 99 {
					onlineval = fmt.Sprintf("%02d时%02d分%02d秒", timehours, time.Unix(onlinetimevar, 0).Minute(), time.Unix(onlinetimevar, 0).Second()) //在线时长
				} else {
					onlineval = fmt.Sprintf("%d时%02d分%02d秒", timehours, time.Unix(onlinetimevar, 0).Minute(), time.Unix(onlinetimevar, 0).Second()) //在线时长
				}
				online.Onlinetime = onlineval                       //在线时长
				online.Ipaddress = userroom[urolename[i]].Ipaddress //ip地址
				onlinelist = append(onlinelist, online)
			}
			count = ulength
		}

		// json
		data := make(map[string]interface{})
		data["aaData"] = onlinelist
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
