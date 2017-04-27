package mqtt

import (
	"strconv"
	"time"
	m "weserver/models"

	"github.com/astaxie/beego"

	"weserver/controllers"
	. "weserver/src/tools"
	// for json get
)

type TeacherController struct {
	controllers.PublicController
}

type teacherMessage struct {
	infochan chan *TeacherInfo
}

var (
	teacher *teacherMessage
)

func init() {
	teacher = &teacherMessage{
		infochan: make(chan *TeacherInfo, 20480),
	}
	teacher.runWriteDb()
}

//Add teacher
func (this *TeacherController) OperateTeacher() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseTeacherMsg(msg)
		if b {
			this.Rsp(true, "老师信息发送成功", "")
			return
		} else {
			this.Rsp(false, "老师信息发送失败,请重新发送", "")
			return
		}
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

func parseTeacherMsg(msg string) bool {
	msginfo := new(TeacherInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("Teacher: simplejson error", err)
		return false
	}
	//可扩展为实时发送
	/*
		info.MsgType = MSG_TYPE_TEACHER_ADD
		topic := info.Room

		beego.Debug("info", info)

		v, err := ToJSON(info)
		if err != nil {
			beego.Error("json error", err)
			return false
		}
		mq.SendMessage(topic, v)
	*/
	// 消息入库
	operateTeacherdata(info)
	return true
}

// 写数据
func (n *teacherMessage) runWriteDb() {
	go func() {
		for {
			select {
			case infoMsg, ok := <-n.infochan:
				if ok {
					operateData(infoMsg)
				}
			}
		}
	}()
}

func operateTeacherdata(info TeacherInfo) {
	jsondata := &info
	select {
	case teacher.infochan <- jsondata:
		break
	default:
		beego.Error("WRITE Teacher db error!!!")
		break
	}
}

func operateData(info *TeacherInfo) {
	beego.Debug("TeacherOperate", info)
	var teacher m.Teacher
	teacher.Id = info.Id
	teacher.Room = info.Room
	OPERTYPE := info.OperType
	switch OPERTYPE {
	case OPERATE_TEACHER_ADD:
		if teacher.Id == 0 {
			err := addTeacherConten(info)
			if err != nil {
				beego.Debug("Oper Teacher Add Fail", err)
			}
		}
		break
	case OPERATE_TEACHER_DEL:
		_, err := m.DelTeacherById(teacher.Id)
		if err != nil {
			beego.Debug("Oper Teacher Del Fail", err)
		}
		break
	case OPERATE_TEACHER_UPDATE:
		err := updateTeacherConten(info)
		if err != nil {
			beego.Debug("Oper Teacher update Fail", err)
		}
		break
	default:
	}
}

func addTeacherConten(info *TeacherInfo) error {
	beego.Debug("Add TeacherInfo", info)
	var teacher m.Teacher
	teacher.Room = info.Room
	teacher.Icon = info.Icon
	teacher.Name = info.Name
	teacher.Title = info.Title
	teacher.Data = info.Data
	teacher.Time = info.Time
	teacher.Datatime = time.Now()

	_, err := m.AddTeacher(&teacher)
	if err != nil {
		beego.Debug("Add Teacher Fail:", err)
		return err
	}
	return nil
}

func updateTeacherConten(info *TeacherInfo) error {
	beego.Debug("Add TeacherInfo", info)
	var teacher m.Teacher
	teacher.Id = info.Id
	teacher.Room = info.Room
	teacher.Icon = info.Icon
	teacher.Name = info.Name
	teacher.Title = info.Title
	teacher.Data = info.Data
	teacher.Time = info.Time
	teacher.Datatime = time.Now()

	_, err := m.UpdateTeacherInfo(&teacher)
	if err != nil {
		beego.Debug("Add Teacher Fail:", err)
		return err
	}
	return nil
}
