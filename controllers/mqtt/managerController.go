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

// 当前在线
func (this *ManagerController) GetUserOnline() {
	if this.IsAjax() {
		roomId := this.GetString("Room")
		onlineuser, err := m.GetLoginInfoToday(roomId)
		if err != nil {
			beego.Error("get the user error", err)
		}
		var userInfo []OnLineInfo
		for _, user := range onlineuser {
			if len(user.Username) > 0 {
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

func (this *ManagerController) GetUserLogin() {
	if this.IsAjax() {
		roomId := this.GetString("Room")
		Uname := this.GetString("Username")
		_, count, err := m.GetRegistPermiss(roomId, Uname)
		if count == 1 && err == nil {
			this.Rsp(true, "验证权限通过", "")
			return
		} else {
			this.Rsp(false, "验证权限失败", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

func (this *ManagerController) GetUserApply() {
	if this.IsAjax() {
		roomId := this.GetString("Room")
		Username := this.GetString("Username")
		Icon := this.GetString("Icon")
		Nickname := this.GetString("Nickname")

		config, _ := m.GetSysConfig()
		configRole := config.Registerrole
		configTitle := config.Registertitle
		configVerify := config.Verify
		u := new(m.Regist)
		u.Room = roomId
		u.Username = Username
		if configVerify == 0 { //是否开启验证  0开启 1不开启
			u.RegStatus = 1
		} else {
			u.RegStatus = 2
		}
		u.UserIcon = Icon
		u.Role = &m.Role{Id: configRole}
		u.Title = &m.Title{Id: configTitle}
		u.IsShutup = false //默认0
		u.Nickname = Nickname
		u.Lastlogintime = time.Now()
		userid, err := m.AddRegistUser(u)
		if err == nil && userid > 0 {
			this.Rsp(true, "", "")
			return
		} else {
			beego.Error(err)
			this.Rsp(false, "", "")
			return
		}
	}
	this.Ctx.WriteString("")
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

//禁言操作
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

		arr, ok := mq.MapShutUp[msg.Room]
		if ok {
			for _, v := range arr {
				if v == msg.Uname {
					break
				} else {
					arr = append(arr, msg.Uname)
					mq.MapShutUp[msg.Room] = arr
				}
			}
		} else {
			mq.MapShutUp[msg.Room] = []string{msg.Uname}
		}
		beego.Debug("Shut up", mq.MapShutUp)
		// beego.Debug("info", msg)
		// topic := msg.Room
		// v, err := ToJSON(msg)
		// if err != nil {
		// 	beego.Error("json error", err)
		// 	return false
		// }
		// mq.SendMessage(topic, v) //发消息

		// 更新user 字段
		UpdateUserInfo(msg)
	}
	return true
}

//解除禁言
func (this *ManagerController) GetUnShutUpInfo() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseUnShutUpMsg(msg)
		if b {
			this.Rsp(true, "解除禁言消息发送成功", "")
			return
		} else {
			this.Rsp(false, "解除禁言消息发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

//解除禁言
func parseUnShutUpMsg(msg string) bool {
	var msginfo ShutUpInfo
	// msginfo := new(ShutUpInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("UnShutup simplejson error", err)
		return false
	}
	for i := 0; i < len(info); i++ {
		var msg ShutUpInfo
		msg.Room = info[i].Room
		msg.Uname = info[i].Uname
		msg.IsShutUp = info[i].IsShutUp
		// msg.MsgType = MSG_TYPE_UNSHUTUP

		arr, ok := mq.MapShutUp[msg.Room]
		if ok {
			for i, v := range arr {
				if v == msg.Uname {
					index := i + 1
					arr = append(arr[:i], arr[index:]...) //删除
					mq.MapShutUp[msg.Room] = arr
					break
				} else {
					break
				}
			}
		} else {
			beego.Debug("UnShutUp no Find element")
		}
		beego.Debug("UnShut up", mq.MapShutUp)
		// beego.Debug("info", msg)
		// topic := msg.Room
		// v, err := ToJSON(msg)
		// if err != nil {
		// 	beego.Error("json error", err)
		// 	return false
		// }
		// mq.SendMessage(topic, v) //发消息

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
	var user m.Regist
	user.Room = info.Room
	user.Username = info.ObjUid
	_, err := m.DelRegistUame(user.Room, user.Username)
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
	var u m.Regist
	u.Room = info.Room
	u.Username = info.Uname
	u.IsShutup = info.IsShutUp
	// if info.IsShutUp == 1 {
	// 	u.IsShutup = true
	// } else {
	// 	u.IsShutup = false
	// }
	_, err := m.UpdateRegistIsShut(u.Room, u.Username, u.IsShutup)
	if err != nil {
		beego.Debug("Update Shutup Field fail", err)
	}
}
