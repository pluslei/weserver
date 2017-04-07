package mqtt

import (
	"time"
	m "weserver/models"
	mq "weserver/src/mqtt"

	"github.com/astaxie/beego"

	"weserver/controllers"
	. "weserver/src/tools"
	// for json get
)

type NoticeController struct {
	controllers.PublicController
}

type noticeMessage struct {
	infochan chan *NoticeInfo
	Delchan  chan *NoticeDEL
}

var (
	notice *noticeMessage
)

func init() {
	notice = &noticeMessage{
		infochan: make(chan *NoticeInfo, 20480),
		Delchan:  make(chan *NoticeDEL, 20480),
	}
	notice.runWriteDb()
}

//发布公告
func (this *NoticeController) GetPublishNotice() {
	if this.IsAjax() {
		noticMsg := this.GetString("str")
		b := parseNoticeMsg(noticMsg)
		if b {
			this.Rsp(true, "消息发送成功", "")
			return
		} else {
			this.Rsp(false, "消息发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

//删除公告
func (this *NoticeController) DeleteNotice() {
	if this.IsAjax() {
		room := this.GetString("Room")
		id, _ := this.GetInt64("Id")
		b := DelNotice(room, id)
		if b {
			this.Rsp(true, "消息发送成功", "")
			return
		} else {
			this.Rsp(false, "消息发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

//历史公告
func (this *NoticeController) GetRoomNoticeList() {
	if this.IsAjax() {
		roomId := this.GetString("room") //公司房间标识符
		sysconfig, _ := m.GetAllSysConfig()
		noticeCount := sysconfig.NoticeCount //公告消息条数
		var historyNotice []m.Notice
		historyNotice, _, _ = m.GetNoticeList(noticeCount, roomId)
		data := make(map[string]interface{})
		data["historyNotice"] = historyNotice //公告的历史信息
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

func parseNoticeMsg(msg string) bool {
	msginfo := new(NoticeInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("simplejson error", err)
		return false
	}
	info.MsgType = MSG_TYPE_NOTICE_ADD //公告
	topic := info.Room

	beego.Debug("info", info)

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}

	mq.SendMessage(topic, v) //发消息

	// 消息入库
	insertMsgdata(info)
	return true
}

func DelNotice(room string, id int64) bool {
	var info NoticeDEL
	info.Id = id
	info.Room = room
	info.MsgType = MSG_TYPE_NOTICE_DEL

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("DELETE Notice JSON ERROR", err)
		return false
	}
	mq.SendMessage(room, v) //发消息
	DeleteMsg(info)
	return true
}

// 写数据
func (n *noticeMessage) runWriteDb() {
	go func() {
		for {
			infoMsg, ok := <-n.infochan
			if ok {
				addContent(infoMsg)
			}
			infoDel, ok1 := <-n.Delchan
			if ok1 {
				delContent(infoDel)
			}
		}
	}()
}

func insertMsgdata(info NoticeInfo) {
	jsondata := &info
	select {
	case notice.infochan <- jsondata:
		break
	default:
		beego.Error("WRITE NOTICE db error!!!")
		break
	}
}

func DeleteMsg(info NoticeDEL) {
	jsondata := &info
	select {
	case notice.Delchan <- jsondata:
		break
	default:
		beego.Error("DELETE NOTICE db error!!!")
		break
	}
}

func delContent(info *NoticeDEL) {
	beego.Debug("NoticeDEL", info)
	//写数据库
	var notice m.Notice
	notice.Id = info.Id
	notice.Room = info.Room
	_, err := m.DelNoticeById(notice.Id)
	if err != nil {
		beego.Debug("AddNotice Fail:", err)
	}
}

func addContent(info *NoticeInfo) {
	beego.Debug("NoticeInfo", info)
	//写数据库
	var notice m.Notice
	notice.Room = info.Room
	notice.Uname = info.Uname
	notice.Nickname = info.Nickname
	notice.Data = info.Content
	notice.Time = info.Time
	notice.Datatime = time.Now()

	_, err := m.AddNoticeMsg(&notice)
	if err != nil {
		beego.Debug("AddNotice Fail:", err)
	}
}
