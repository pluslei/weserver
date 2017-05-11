package haoadmin

import (
	"weserver/models"

	"time"

	"github.com/astaxie/beego"
)

type TeacherController struct {
	CommonController
}

func (this *TeacherController) Index() {
	if this.IsAjax() {
		user := this.GetSession("userinfo").(*models.User)
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
		teacher, count := models.GetTeacherInfoList(iStart, iLength, "-Id", user.CompanyId)
		for _, item := range teacher {
			roomInfo, err := models.GetRoomInfoByRoomID(item["Room"].(string))
			if err != nil {
				item["Room"] = "未知房间"
			} else {
				item["Room"] = roomInfo.RoomTitle
			}
			Info, err := models.GetCompanyById(item["CompanyId"].(int64))
			if err != nil {
				item["CompanyName"] = "未知公司"
			} else {
				item["CompanyName"] = Info.Company
			}
		}

		// json
		data := make(map[string]interface{})
		data["aaData"] = teacher
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonMenu()

		roonInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Error("Get the Roominfo error", err)
			return
		}
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/teacher/list.html"
	}
}

func (this *TeacherController) Add() {
	action := this.GetString("action")
	if action == "add" {
		teacher := new(models.Teacher)
		teacher.Room = this.GetString("Room")
		teacher.Name = this.GetString("Name")
		teacher.Icon = this.GetString("Icon")
		teacher.Data = this.GetString("Data")
		teacher.Datatime = time.Now()
		time := time.Now()
		tm := time.Format("2006-01-02 15:04:05")
		teacher.Time = tm
		_, err := models.AddTeacher(teacher)
		if err != nil {
			this.AlertBack("添加失败")
			return
		}
		this.Alert("添加成功", "teacher_index")
	} else {
		this.CommonMenu()

		roonInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Error("Get the Roominfo error", err)
			return
		}
		beego.Debug("roonInfo", roonInfo)
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/teacher/add.html"
	}
}

func (this *TeacherController) Edit() {
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		this.AlertBack("数据错误")
		return
	}
	if action == "edit" {
		teacher := make(map[string]interface{})
		teacher["Room"] = this.GetString("Room")
		teacher["Name"] = this.GetString("Name")
		teacher["Icon"] = this.GetString("Icon")
		teacher["Data"] = this.GetString("Data")
		teacher["Datatime"] = time.Now()
		_, err := models.UpdateTeacherInfoById(id, teacher)
		if err != nil {
			this.AlertBack("修改失败")
		}
		this.Alert("修改成功", "teacher_index")
	} else {
		this.CommonMenu()

		roonInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Error("Get the Roominfo error", err)
			return
		}
		teacher, err := models.GetTeacherInfoById(id)
		if err != nil {
			this.AlertBack("数据错误")
		}
		this.Data["teacher"] = teacher
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/teacher/edit.html"
	}
}

func (this *TeacherController) Del() {
	id, err := this.GetInt64("id")
	if err != nil {
		this.Rsp(false, "获取错误", "")
		return
	}
	_, err = models.DelTeacherById(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
		return
	}
	this.Rsp(true, "删除成功", "")
}

func (this *TeacherController) GetTeacher() {
	room := this.GetString("room")
	t, err := models.GetTeacherListByRoom(room)
	if err != nil {
		beego.Error("error", err)
	}
	this.Data["json"] = t
	this.ServeJSON()
}
