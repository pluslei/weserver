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
	operchan chan *TeacherOperate
}

var (
	teacher *teacherMessage
)

func init() {
	teacher = &teacherMessage{
		infochan: make(chan *TeacherInfo, 20480),
		operchan: make(chan *TeacherOperate, 20480),
	}
	teacher.runWriteDb()
}

//Add teacher
func (this *TeacherController) AddTeacher() {
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

//OP teacher
func (this *TeacherController) OperateTeacher() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseOPTeacherMsg(msg)
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
		var Uname []m.ThumbInfo
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
					info.IsTop = historyTeacher[i].IsTop
					info.ThumbNum = historyTeacher[i].ThumbNum
					info.Data = historyTeacher[i].Data
					info.Time = historyTeacher[i].Time

					historyThumbInfo, _, err := m.GetMoreThumbInfo(info.Room, info.Id)
					if err != nil {
						beego.Debug("GetMoreThumbInfo() Fail")
						return
					}
					for i := 0; i < len(historyThumbInfo); i++ {
						var thumb m.ThumbInfo
						thumb.Id = historyThumbInfo[i].Teacher.Id
						thumb.Username = historyThumbInfo[i].Username
						thumb.IsThumb = historyThumbInfo[i].IsThumb
						Uname = append(Uname, thumb)
					}
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
					info.IsTop = historyTeacher[i].IsTop
					info.ThumbNum = historyTeacher[i].ThumbNum
					info.Data = historyTeacher[i].Data
					info.Time = historyTeacher[i].Time

					historyThumbInfo, _, err := m.GetMoreThumbInfo(info.Room, info.Id)
					if err != nil {
						beego.Debug("GetMoreThumbInfo() Fail")
						return
					}
					for i := 0; i < len(historyThumbInfo); i++ {
						var thumb m.ThumbInfo
						thumb.Id = historyThumbInfo[i].Teacher.Id
						thumb.Username = historyThumbInfo[i].Username
						thumb.IsThumb = historyThumbInfo[i].IsThumb
						Uname = append(Uname, thumb)
					}
					teacherinfo = append(teacherinfo, info)
				}
			}
			data["historyTeacher"] = teacherinfo
			data["Uname"] = Uname
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
					info.IsTop = historyTeacher[i].IsTop
					info.ThumbNum = historyTeacher[i].ThumbNum
					info.Data = historyTeacher[i].Data
					info.Time = historyTeacher[i].Time

					historyThumbInfo, _, err := m.GetMoreThumbInfo(info.Room, info.Id)
					if err != nil {
						beego.Debug("GetMoreThumbInfo() Fail")
						return
					}
					for i := 0; i < len(historyThumbInfo); i++ {
						var thumb m.ThumbInfo
						thumb.Id = historyThumbInfo[i].Teacher.Id
						thumb.Username = historyThumbInfo[i].Username
						thumb.IsThumb = historyThumbInfo[i].IsThumb
						Uname = append(Uname, thumb)
					}
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
					info.IsTop = historyTeacher[i].IsTop
					info.ThumbNum = historyTeacher[i].ThumbNum
					info.Data = historyTeacher[i].Data
					info.Time = historyTeacher[i].Time

					historyThumbInfo, _, err := m.GetMoreThumbInfo(info.Room, info.Id)
					if err != nil {
						beego.Debug("GetMoreThumbInfo() Fail")
						return
					}
					for i := 0; i < len(historyThumbInfo); i++ {
						var thumb m.ThumbInfo
						thumb.Id = historyThumbInfo[i].Teacher.Id
						thumb.Username = historyThumbInfo[i].Username
						thumb.IsThumb = historyThumbInfo[i].IsThumb
						Uname = append(Uname, thumb)
					}
					teacherinfo = append(teacherinfo, info)
				}
			}
			data["historyTeacher"] = teacherinfo
			data["Uname"] = Uname
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

func parseOPTeacherMsg(msg string) bool {
	msginfo := new(TeacherOperate)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("TeacherOperate: simplejson error", err)
		return false
	}
	info.MsgType = MSG_TYPE_STRATEGY_OPE

	beego.Debug("Operate Teacher info", info)
	/*
		room := info.Room
		v, err := ToJSON(info)
		if err != nil {
			beego.Error("OPERATE Strategy JSON ERROR", err)
			return false
		}
		mq.SendMessage(room, v) //发消息
	*/
	OperateTeacherMsg(info)
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
			case infoOper, ok1 := <-n.operchan:
				if ok1 {
					OperateTeacherContent(infoOper)
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

func OperateTeacherMsg(info TeacherOperate) {
	jsondata := &info
	select {
	case teacher.operchan <- jsondata:
		break
	default:
		beego.Error("OPER NOTICE db error!!!")
		break
	}
}

func OperateTeacherContent(info *TeacherOperate) {
	beego.Debug("TeacherOper", info)
	var teacher m.Teacher
	teacher.Id = info.Id
	teacher.Room = info.Room
	OPERTYPE := info.OperType
	switch OPERTYPE {
	case OPERATE_TEACHER_TOP:
		_, err := m.TopOption(teacher.Id)
		if err != nil {
			beego.Debug("Oper Teacher Top Fail", err)
		}
		break
	case OPERATE_TEACHER_UNTOP:
		_, err := m.UnTopOption(teacher.Id)
		if err != nil {
			beego.Debug("Oper Teacher UnTop Fail", err)
		}
		break
	case OPERATE_TEACHER_THUMB:
		_, err := m.ThumbTeacherAdd(teacher.Id)
		if err != nil {
			beego.Debug("Oper Teacher Thumb Fail", err)
			return
		}
		err = addThumbInfo(info, teacher.Id)
		if err != nil {
			beego.Debug("Oper ThumbInfo Add Fail", err)
			return
		}
		break
	case OPERATE_TEACHER_UNTHUMB:
		_, err := m.ThumbTeacherDel(teacher.Id)
		if err != nil {
			beego.Debug("Oper Teacher Thumb Fail", err)
			return
		}
		err = UnThumbInfo(info, teacher.Id)
		if err != nil {
			beego.Debug("Oper UnThumbInfo Fail", err)
			return
		}
		break
	case OPERATE_TEACHER_DEL:
		_, err := m.DelTeacherById(teacher.Id)
		if err != nil {
			beego.Debug("Oper Teacher Del Fail", err)
		}
		break
	default:
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
			_, err := addTeacherConten(info)
			if err != nil {
				beego.Debug("Oper Teacher Add Fail", err)
				return
			}
		}
		break
	case OPERATE_TEACHER_UPDATE:
		err := updateTeacherConten(info)
		if err != nil {
			beego.Debug("Oper Teacher update Fail", err)
			return
		}
		break
	default:
	}
}

func addThumbInfo(info *TeacherOperate, Id int64) error {
	var thumb m.ThumbInfo
	data, err := m.GetThumbInfo(info.Username, info.Room, Id)
	if data.Id != 0 && err == nil {
		_, err := m.UpdateThumb(data.Id)
		if err != nil {
			beego.Debug("Update Thumb status Fail")
			return nil
		}
	} else {
		thumb.CompanyId = info.CompanyId
		thumb.Room = info.Room
		thumb.Nickname = info.Nickname
		thumb.Username = info.Username
		thumb.Timestr = time.Now().Format("2006-01-02 15:04:05")
		thumb.IsThumb = true
		thumb.Teacher = &m.Teacher{Id: Id}
		beego.Debug("Add ThumbInfo", thumb)
		_, err = m.AddThumbInfo(&thumb)
		if err != nil {
			beego.Debug("Add ThumbInfo Fail:", err)
			return err
		}
	}
	return nil
}

func UnThumbInfo(info *TeacherOperate, Id int64) error {
	beego.Debug("Update UnThumbInfo", info)
	_, err := m.UpdateUnThumb(info.Username, info.Room, Id)
	if err != nil {
		beego.Debug("Update UnThumbInfo Fail:", err)
		return err
	}
	return nil
}

func addTeacherConten(info *TeacherInfo) (int64, error) {
	beego.Debug("Add TeacherInfo", info)
	var teacher m.Teacher
	teacher.Room = info.Room
	teacher.Icon = info.Icon
	teacher.Name = info.Name
	teacher.Title = info.Title
	teacher.IsTop = info.IsTop
	teacher.ThumbNum = info.ThumbNum
	teacher.Data = info.Data
	teacher.Time = info.Time
	teacher.Datatime = time.Now()

	Id, err := m.AddTeacher(&teacher)
	if err != nil {
		beego.Debug("Add Teacher Fail:", err)
		return 0, err
	}
	return Id, nil
}

func updateTeacherConten(info *TeacherInfo) error {
	beego.Debug("Update TeacherInfo", info)
	var teacher m.Teacher
	teacher.Id = info.Id
	teacher.Room = info.Room
	teacher.Icon = info.Icon
	teacher.Name = info.Name
	teacher.Title = info.Title
	teacher.IsTop = info.IsTop
	teacher.ThumbNum = info.ThumbNum
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
