package mqtt

import (
	"strconv"
	"time"
	m "weserver/models"
	mq "weserver/src/mqtt"

	"weserver/controllers"
	. "weserver/src/tools"

	"github.com/astaxie/beego"
	// for json get
)

type ManagerController struct {
	controllers.PublicController
}

type kickMessage struct {
	Delchan chan *KickOutInfo
}

var (
	kick *kickMessage
)

func init() {
	kick = &kickMessage{
		Delchan: make(chan *KickOutInfo, 20480),
	}
	kick.runWriteDb()
}

/*
var slice []string

func (this *ManagerController) Init() {
	slice = make([]string, 0, 100)
	shutInfo, err := m.GetShutUpInfoToday()
	if err != nil {
		beego.Error("get the shutup error", err)
	}
	for _, info := range shutInfo {
		if len(info.UserIcon) > 0 {
			var Uname string
			Uname = info.Username
			slice = append(slice, Uname)
		}
	}
}
*/

// 当前在线
func (this *ManagerController) GetUserOnline() {
	if this.IsAjax() {
		roomId := this.GetString("Room")
		onlineuser, err := m.GetUserInfoToday(roomId)
		if err != nil {
			beego.Error("get the user error", err)
		}
		var userInfo []OnLineInfo
		for _, user := range onlineuser {
			if len(user.UserIcon) > 0 {
				var info OnLineInfo
				info.Uname = EncodeB64(user.Username)
				info.Nickname = EncodeB64(user.Nickname)
				info.UserIcon = EncodeB64(user.UserIcon)
				str := strconv.FormatBool(user.IsShutup)
				info.ShutUp = EncodeB64(str)
				userInfo = append(userInfo, info)
			}
		}
		data := make(map[string]interface{})
		data["userlist"] = userInfo
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

//踢人
func (this *ManagerController) GetKickOutInfo() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseKickMsg(msg)
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

//禁言
func (this *ManagerController) GetShutUpInfo() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseShutUpMsg(msg)
		if b {
			this.Rsp(true, "禁言消息发送成功", "")
			return
		} else {
			this.Rsp(false, "禁言消息发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

//禁言消息
func parseShutUpMsg(msg string) bool {
	var msginfo ShutUpInfo
	// msginfo := new(ShutUpInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("Shutup simplejson error", err)
		return false
	}
	for i := 0; i < len(info); i++ {
		var msg ShutUpInfo
		msg.Room = info[i].Room
		msg.Uname = info[i].Uname
		msg.IsShutUp = info[i].IsShutUp
		msg.MsgType = MSG_TYPE_SHUTUP

		beego.Debug("info", msg)
		topic := msg.Room
		v, err := ToJSON(msg)
		if err != nil {
			beego.Error("json error", err)
			return false
		}
		mq.SendMessage(topic, v) //发消息

		// 更新user 字段
		UpdateUserInfo(msg)
	}
	return true
}

func parseKickMsg(msg string) bool {
	msginfo := new(KickOutInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("KickOut simplejson error", err)
		return false
	}
	for i := 0; i < len(info); i++ {
		var msg KickOutInfo
		msg.Room = info[i].Room
		msg.OperUid = info[i].OperUid
		msg.OperName = info[i].OperName
		msg.ObjUid = info[i].ObjUid
		msg.ObjName = info[i].ObjName

		msg.MsgType = MSG_TYPE_KICKOUT

		beego.Debug("info", info)
		topic := msg.Room
		v, err := ToJSON(msg)
		if err != nil {
			beego.Error("json error", err)
			return false
		}

		mq.SendMessage(topic, v) //发消息

		// 删除此用户
		delKickout(msg)
	}
	return true
}

// 写数据
func (k *kickMessage) runWriteDb() {
	go func() {
		for {
			infoMsg, ok := <-k.Delchan
			if ok {
				delKickContent(infoMsg)
				addKickContent(infoMsg)
			}
		}
	}()
}

func delKickout(info KickOutInfo) {
	jsondata := &info
	select {
	case kick.Delchan <- jsondata:
		break
	default:
		beego.Error("Kick Msg db error!!!")
		break
	}
}

//删除用户
func delKickContent(info *KickOutInfo) {
	beego.Debug("KickOut DELETE", info)
	var user m.User
	user.Room = info.Room
	user.Username = info.ObjUid
	_, err := m.DelUserByUame(user.Room, user.Username)
	if err != nil {
		beego.Debug(" DELETE KickOut Record Fail:", err)
	}
}

//踢人记录
func addKickContent(info *KickOutInfo) {
	beego.Debug("ADD KICK INFO", info)
	//写数据库
	var kick m.KickOut
	kick.Coderoom = info.Room
	kick.Operuid = info.OperUid
	kick.Opername = info.OperName
	kick.Objuid = info.ObjUid
	kick.Objname = info.ObjName
	kick.Status = OPERATE_KICKOUT
	kick.Opertime = time.Now()

	_, err := m.AddKickOut(&kick)
	if err != nil {
		beego.Debug("Add KickOut Record Fail:", err)
	}
}

//更新user表禁言字段
func UpdateUserInfo(info ShutUpInfo) {
	var u m.User
	u.Room = info.Room
	u.Username = info.Uname
	u.IsShutup = info.IsShutUp
	// if info.IsShutUp == 1 {
	// 	u.IsShutup = true
	// } else {
	// 	u.IsShutup = false
	// }
	_, err := m.UpdateShutUp(u.Room, u.Username, u.IsShutup)
	if err != nil {
		beego.Debug("Update Shutup Field fail", err)
	}
}
