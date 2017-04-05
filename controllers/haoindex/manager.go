package haoindex

import (
	"weserver/controllers"
	m "weserver/models"

	"github.com/astaxie/beego"
)

type ManagerController struct {
	controllers.PublicController
}

//收藏
func (this *ManagerController) GetCollectInfo() {
	if this.IsAjax() {
		collect := new(m.Collect)
		collect.Uname = this.GetString("Uname")
		collect.Nickname = this.GetString("Nickname")
		collect.RoomIcon = this.GetString("RoomIcon")
		collect.RoomTitle = this.GetString("RoomTitle")
		collect.RoomTeacher = this.GetString("RoomTeacher")
		collect.IsCollect, _ = this.GetBool("IsCollect")
		collect.IsAttention, _ = this.GetBool("IsAttention")
		_, err := m.AddCollect(collect)
		if err != nil {
			this.Rsp(false, "收藏写入数据库失败", "")
			beego.Debug(err)
		} else {
			this.Rsp(true, "收藏写入数据库成功", "")
		}
	}
}

// 禁言
func (this *ManagerController) GetShutUpInfo() {

}

// 移除
func (this *ManagerController) GetKickOutInfo() {

}

//发布公告
func (this *ManagerController) PublishNotice() {
	if this.IsAjax() {
		notice := new(m.Notice)
		//strconv.Atoi(code)
		notice.Code, _ = this.GetInt("Code")
		notice.Room = this.GetString("Room")
		notice.Uname = this.GetString("Uname")
		notice.Nickname = this.GetString("Nickname")
		notice.Data = this.GetString("Data")
		time := this.GetString("Datatime")
		beego.Debug("get notice", time)
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
*/
