package haoadmin

import (
	"github.com/astaxie/beego"
	"time"
	"weserver/controllers/mqtt"
	m "weserver/models"
	. "weserver/src/tools"
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

	msgtype := mqtt.NewMessageType(mqtt.MSG_TYPE_CHAT)
	if msgtype.CheckMessage(chatinfo) {
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

	if len(uuid) <= 0 {
		this.Rsp(false, "删除失败", "")
		return
	}

	msgtype := mqtt.NewMessageType(mqtt.MSG_TYPE_DEL)
	if msgtype.DelMessage(uuid) {
		m.DelChatById(uuid)
		this.Rsp(true, "审核成功", "")
		return
	}
	this.Rsp(false, "审核失败", "")
}
