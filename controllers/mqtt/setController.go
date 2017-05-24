package mqtt

import (
	"weserver/controllers"
	m "weserver/models"
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

func (this *SetController) SetIcon() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseSetIcon(msg)
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

func (this *SetController) SetNickname() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseSetNicknameMsg(msg)
		if b {
			this.Rsp(true, "修改昵称", "")
			return
		} else {
			this.Rsp(false, "修改昵称发送失败,请重新发送", "")
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

func parseSetIcon(msg string) bool {
	info, err := ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("parseSetIcon simplejson error", err)
		return false
	}
	info, ok := info.(SetInfo)
	if ok {
		update(info.(SetInfo))
		return true
	}
	return false
}

func parseSetNicknameMsg(msg string) bool {
	info, err := ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("parseSetNickName simplejson error", err)
		return false
	}
	info, ok := info.(SetInfo)
	if ok {
		update(info.(SetInfo))
		return true
	}
	return false
}

func parsePhoneNumMsg(msg string) bool {
	info, err := ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("parseSetNickName simplejson error", err)
		return false
	}
	info, ok := info.(SetInfo)
	if ok {
		update(info.(SetInfo))
		return true
	}
	return false
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
	if info.Nickname != "" && info.Icon == "" && info.Phonenum == 0 {
		_, err := m.UpdateRegistNickname(info.Uname, info.CompanyId, info.Nickname)
		if err != nil {
			beego.Debug("update Regist nickname error", err)
			return
		}
		_, err = m.UpdateUserNickname(info.Uname, info.Nickname)
		if err != nil {
			beego.Debug("update user nickname error", err)
			return
		}
	}

	if info.Icon != "" && info.Nickname == "" && info.Phonenum == 0 {
		_, err := m.UpdateRegistIcon(info.Uname, info.CompanyId, info.Icon)
		if err != nil {
			beego.Debug("update Regist nickname error", err)
			return
		}
		_, err = m.UpdateUserIcon(info.Uname, info.Icon)
		if err != nil {
			beego.Debug("update user nickname error", err)
			return
		}
	}

	if info.Phonenum != 0 && info.Nickname == "" && info.Icon == "" {
		_, err := m.UpdateUserPhoneNum(info.Uname, info.Phonenum)
		if err != nil {
			beego.Debug("update user Phonenum error", err)
			return
		}
	}
}
