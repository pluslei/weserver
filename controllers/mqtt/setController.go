package mqtt

import (
	"strconv"
	"weserver/controllers"
	m "weserver/models"
	. "weserver/src/cache"
	. "weserver/src/msg"
	. "weserver/src/tools"

	"github.com/astaxie/beego"
	// for json get
)

type SetController struct {
	controllers.PublicController
}

type SetMessage struct {
	set chan *SetInfo
}

var (
	msg *SetMessage
)

func init() {
	msg = &SetMessage{
		set: make(chan *SetInfo, 20480),
	}
	msg.runWriteDb()
}

func (this *SetController) SetIdentiCode() {
	if this.IsAjax() {
		Id, err := this.GetInt64("CompanyId")
		if err != nil {
			beego.Debug("get companyId error", err)
			return
		}
		var info m.Company
		strId := strconv.FormatInt(Id, 10)
		inter, ok := MapCache[strId]
		if !ok {
			info, err = m.GetCompanyById(Id)
			if err != nil {
				beego.Debug("get login companyinfo error")
				return
			}
		} else {
			info, _ = inter.(m.Company)
			beego.Debug("memcache find")
		}
		username := this.GetString("Username")
		phoneNum := this.GetString("phoneNum")
		num, err := strconv.ParseInt(phoneNum, 10, 64)
		code := RandomInt64(1000, 9999)
		_, err = m.UpdateUserAuthCode(username, num, code)
		if err != nil {
			beego.Debug("update phoneNum code error", err)
			return
		}
		SendIdentifyCode(phoneNum, info.Sign, code)
	}
	this.Ctx.WriteString("")
}

func (this *SetController) VerifyCode() {
	if this.IsAjax() {
		username := this.GetString("Username")
		phoneNum := this.GetString("phoneNum")
		code := this.GetString("AuthCode")
		num, err := strconv.ParseInt(phoneNum, 10, 64)
		authCode, err := strconv.ParseInt(code, 10, 64)
		_, count, err := m.GetUserAuthCode(username, num, authCode)
		if err != nil && count == 1 {
			this.Data["json"] = true
			this.ServeJSON()
			return
		} else {
			this.Data["json"] = false
			this.ServeJSON()
			return
		}
	}
	this.Ctx.WriteString("")
}

func (this *SetController) Setperson() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseSetperson(msg)
		if b {
			this.Rsp(true, "修改图标", "")
			return
		} else {
			this.Rsp(false, "修改图标失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

func (this *SetController) SetPhoneNum() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parsePhoneNumMsg(msg)
		if b {
			this.Rsp(true, "修改手机号", "")
			return
		} else {
			this.Rsp(false, "修改手机号发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

func (this *SetController) SetPushWechat() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parsePushWechatMsg(msg)
		if b {
			this.Rsp(true, "修改手机号", "")
			return
		} else {
			this.Rsp(false, "修改手机号发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

func (this *SetController) SetPushSMS() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parsePushSMS(msg)
		if b {
			this.Rsp(true, "修改手机号", "")
			return
		} else {
			this.Rsp(false, "修改手机号发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

func parseSetperson(msg string) bool {
	msginfo := new(SetInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("parseSetperson simplejson error", err)
		return false
	}
	update(info)
	return true
}

func parsePushWechatMsg(msg string) bool {
	msginfo := new(SetInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("parsePushWechatMsg simplejson error", err)
		return false
	}
	update(info)
	return true
}

func parsePhoneNumMsg(msg string) bool {
	msginfo := new(SetInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("parsePhoneNumMsg simplejson error", err)
		return false
	}
	update(info)
	return true
}

func parsePushSMS(msg string) bool {
	msginfo := new(SetInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("parsePushSMS simplejson error", err)
		return false
	}
	update(info)
	return true
}

func (s *SetMessage) runWriteDb() {
	go func() {
		for {
			infoMsg, ok := <-s.set
			if ok {
				updateInfo(infoMsg)
			}
		}
	}()
}

func update(info SetInfo) {
	jsondata := &info
	select {
	case msg.set <- jsondata:
		break
	default:
		beego.Error("update person settting db error!!!")
		break
	}
}

func updateInfo(info *SetInfo) {

	_, err := m.UpdateRegistNickname(info.Uname, info.CompanyId, info.Nickname, info.Icon)
	if err != nil {
		beego.Debug("update Regist nickname error", err)
		return
	}
	_, err = m.UpdateUserNickname(info.Uname, info.Nickname, info.Icon)
	if err != nil {
		beego.Debug("update user nickname error", err)
		return
	}

	if info.Nickname == "" && info.Icon == "" && info.Phonenum == 0 && info.PushSMS == "" {
		flag := strconv.ParseInt(info.PushWechat, 64, 10)
		_, err := m.UpdateRegistPushWechat(info.RoomId, info.Uname, flag)
		if err != nil {
			beego.Debug("update user PushWechat error", err)
			return
		}
	}

	if info.Nickname == "" && info.Icon == "" && info.Phonenum == 0 && info.PushWechat == "" {
		flag := strconv.ParseInt(info.PushSMS, 64, 10)
		_, err := m.UpdateRegistPushSMS(info.RoomId, info.Uname, flag)
		if err != nil {
			beego.Debug("update user PushSMS error", err)
			return
		}
	}

	if info.Phonenum != 0 && info.Nickname == "" && info.Icon == "" && info.PushSMS == "" && info.PushWechat == "" {
		_, err := m.UpdateUserPhoneNum(info.Uname, info.Phonenum)
		if err != nil {
			beego.Debug("update user Phonenum error", err)
			return
		}
	}
}
