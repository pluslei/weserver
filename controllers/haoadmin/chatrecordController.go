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

type ChatMessage struct {
	Delchan chan *MessageDEL
}

var (
	delMsg *ChatMessage
)

func init() {
	delMsg = &ChatMessage{
		Delchan: make(chan *MessageDEL, 20480),
	}
	delMsg.runWriteDb()
}

// 聊天记录
func (this *ChatRecordController) ChatRecordList() {
	if this.IsAjax() {
		user := this.GetSession("userinfo").(*m.User)
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		nickname := this.GetString("sSearch_0")

		chatrecord, count := m.GetChatRecordList(iStart, iLength, "-datatime", nickname, user.CompanyId)
		for _, v := range chatrecord {
			roomInfo, err := m.GetRoomInfoByRoomID(v["Room"].(string))
			if err != nil {
				v["Room"] = "未知房间"
			} else {
				v["Room"] = roomInfo.RoomTitle
			}
			Info, err := m.GetCompanyById(v["CompanyId"].(int64))
			if err != nil {
				v["CompanyName"] = "未知公司"
			} else {
				v["CompanyName"] = Info.Company
			}
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
		beego.Error("CheckMessge json error", err)
		return false
	}
	mq.SendMessage(topic, v) //发消息
	return true
}

//删除消息
func DelMsg(topic, uuid string) bool {
	info := new(MessageDEL)
	info.Uuid = uuid
	info.Room = topic
	info.MsgType = MSG_TYPE_CHAT_DEL

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("DeMsg json error", err)
		return false
	}
	mq.SendMessage(topic, string(v)) //发消息
	deleteMsg(info)
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
		this.Rsp(false, "Get UUid error", "")
		return
	}
	DelMsg(topic, uuid)
}

// 写数据
func (m *ChatMessage) runWriteDb() {
	go func() {
		for {
			infoMsg, ok := <-m.Delchan
			if ok {
				deleteConten(infoMsg)
			}
		}
	}()
}

func deleteMsg(info *MessageDEL) {
	jsondata := info
	select {
	case delMsg.Delchan <- jsondata:
		break
	default:
		beego.Error("DELETE ChatRecord db error!!!")
		break
	}
}

func deleteConten(info *MessageDEL) {
	beego.Debug("Delete ChatMSG:", info)
	// room := info.Room
	Uuid := info.Uuid
	_, err := m.DelChatById(Uuid)
	if err != nil {
		beego.Debug("Delete ChatMsg error:", err)
	}
}
