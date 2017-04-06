package haoadmin

import (
	"time"
	m "weserver/models"
	. "weserver/src/tools"

	mq "weserver/src/mqtt"

	"github.com/astaxie/beego"
)

type ChatRecordController struct {
	CommonController
}

// 聊天记录
func (this *ChatRecordController) ChatRecordList() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		chatrecord, count := m.GetChatRecordList(iStart, iLength, "-datatime")
		for _, v := range chatrecord {
			v["Datatime"] = v["Datatime"].(time.Time).Format("2006-01-02 15:04:05")
		}
		// json
		data := make(map[string]interface{})
		data["aaData"] = chatrecord
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["localserveraddress"] = beego.AppConfig.String("wslocalServerAdress") + "/rpc"
		this.CommonMenu()
		prevalue := beego.AppConfig.String("company") + "_" + beego.AppConfig.String("room")
		codeid := MainEncrypt(prevalue)
		this.Data["codeid"] = codeid
		this.TplName = "haoadmin/data/chat/list.html"
	}
}

// 审核发消息
func CheckMessage(topic string, msg m.ChatRecord) bool {
	msg.MsgType = MSG_TYPE_CHAT_ADD
	beego.Debug("msg", msg)

	v, err := ToJSON(msg)
	if err != nil {
		beego.Error("json error", err)
		return false
	}
	mq.SendMessage(topic, v) //发消息
	return true
}

//删除消息
func DelMsg(topic, uuid string) bool {
	info := new(MessageDEL)
	info.MsgType = MSG_TYPE_CHAT_DEL
	info.Uuid = uuid

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}
	mq.SendMessage(topic, string(v)) //发消息
	return true
}

// 消息审核
func (this *ChatRecordController) CheckRecord() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Error("get id error", err)
		this.Rsp(false, "审核失败", "")
		return
	}

	chatinfo, err := m.GetChatIdData(id)
	if err != nil {
		beego.Error("get chat data error", err)
		this.Rsp(false, "审核失败", "")
		return
	}

	if CheckMessage(chatinfo.Room, chatinfo) {
		m.UpdateChatStatus(id)
		this.Rsp(true, "审核成功", "")
		return
	}
	this.Rsp(false, "审核失败", "")
	beego.Info("chatinfo", chatinfo)
}

// 删除消息
func (this *ChatRecordController) DelRecord() {
	uuid := this.GetString("uuid")
	topic := this.GetString("topic") //管理页面ajax获取的

	if len(uuid) <= 0 {
		this.Rsp(false, "删除失败", "")
		return
	}

	if DelMsg(topic, uuid) {
		m.DelChatById(uuid)
		this.Rsp(true, "审核成功", "")
		return
	}
	this.Rsp(false, "审核失败", "")
}
