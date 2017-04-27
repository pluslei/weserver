package mqtt

import (
	"strconv"
	m "weserver/models"

	"github.com/astaxie/beego"

	"weserver/controllers"
	// for json get
)

type TeacherController struct {
	controllers.PublicController
}

//Teacher List
func (this *TeacherController) GetTeacherList() {
	if this.IsAjax() {
		strId := this.GetString("Id")
		beego.Debug("id", strId)
		nId, _ := strconv.ParseInt(strId, 10, 64)
		roomId := this.GetString("room")
		beego.Debug("teacher list ", nId, roomId)
		data := make(map[string]interface{})
		sysconfig, _ := m.GetAllSysConfig()
		sysCount := sysconfig.TeacherCount
		var teacherinfo []m.Teacher
		historyTeacher, totalCount, err := m.GetTeacherList(roomId)
		if err != nil {
			beego.Debug("Get TeacherList error:", err)
			this.Rsp(false, "Get TeacherList error", "")
			return
		}
		if nId == 0 {
			var i int64
			if totalCount < sysCount {
				beego.Debug("nCount sysCont", totalCount, sysCount)
				for i = 0; i < totalCount; i++ {
					var info m.Teacher
					info.Id = historyTeacher[i].Id
					info.Room = historyTeacher[i].Room
					info.Icon = historyTeacher[i].Icon
					info.Name = historyTeacher[i].Name
					info.Title = historyTeacher[i].Title
					info.Data = historyTeacher[i].Data
					info.Time = historyTeacher[i].Time
					teacherinfo = append(teacherinfo, info)
				}
			} else {
				for i = 0; i < sysCount; i++ {
					var info m.Teacher
					info.Id = historyTeacher[i].Id
					info.Room = historyTeacher[i].Room
					info.Icon = historyTeacher[i].Icon
					info.Name = historyTeacher[i].Name
					info.Title = historyTeacher[i].Title
					info.Data = historyTeacher[i].Data
					info.Time = historyTeacher[i].Time
					teacherinfo = append(teacherinfo, info)
				}
			}
			data["historyTeacher"] = teacherinfo
			this.Data["json"] = &data
			this.ServeJSON()
		} else {
			var index int64
			for nindex, value := range historyTeacher {
				if value.Id == nId {
					index = int64(nindex) + 1
				}
			}
			beego.Debug("index", index)
			nCount := index + sysCount
			mod := (totalCount - nCount) % sysCount
			beego.Debug("mod", mod)
			if nCount > totalCount && mod == 0 {
				beego.Debug("mod = 0")
				data["historyTeacher"] = ""
				this.Data["json"] = &data
				this.ServeJSON()
				return
			}
			if nCount < totalCount {
				for i := index; i < nCount; i++ {
					var info m.Teacher
					info.Id = historyTeacher[i].Id
					info.Room = historyTeacher[i].Room
					info.Icon = historyTeacher[i].Icon
					info.Name = historyTeacher[i].Name
					info.Title = historyTeacher[i].Title
					info.Data = historyTeacher[i].Data
					info.Time = historyTeacher[i].Time
					teacherinfo = append(teacherinfo, info)
				}
			} else {
				for i := index; i < totalCount; i++ {
					var info m.Teacher
					info.Id = historyTeacher[i].Id
					info.Room = historyTeacher[i].Room
					info.Icon = historyTeacher[i].Icon
					info.Name = historyTeacher[i].Name
					info.Title = historyTeacher[i].Title
					info.Data = historyTeacher[i].Data
					info.Time = historyTeacher[i].Time
					teacherinfo = append(teacherinfo, info)
				}
			}
			data["historyTeacher"] = teacherinfo
			this.Data["json"] = &data
			this.ServeJSON()
		}
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

//AllTeacher List
func (this *TeacherController) GetAllTeahcerList() {
	if this.IsAjax() {
		roomId := this.GetString("room")
		data := make(map[string]interface{})
		historyTeacher, _, err := m.GetTeacherList(roomId)
		if err != nil {
			beego.Debug("GetAllTeacherList error", err)
			this.Rsp(false, "Get AllTeacherList Error", "")
			return
		}
		data["historyTeacher"] = historyTeacher
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}
