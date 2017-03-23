package haoadmin

import (
	"github.com/astaxie/beego"
	"strings"
	"time"
	"weserver/controllers/mqtt"
	m "weserver/models"
	. "weserver/src/tools"
)

type QsController struct {
	CommonController
}

// 发送广播列表
func (this *QsController) BroadList() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")
		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		Broadlist, count := m.GetBroadcastlist(iStart, iLength, "Room")
		for _, item := range Broadlist {
			item["Datatime"] = item["Datatime"].(time.Time).Format("2006-01-02 15:04:05")
		}
		// json
		data := make(map[string]interface{})
		data["aaData"] = Broadlist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonController.CommonMenu()
		this.TplName = "haoadmin/data/qs/databroad.html"
	}
}

// 发送广播
func (this *QsController) SendBroad() {
	prevalue := beego.AppConfig.String("company") + "_" + beego.AppConfig.String("room")
	codeid := MainEncrypt(prevalue)
	this.Data["codeid"] = codeid
	if this.GetSession("userinfo") != nil {
		UserInfo := this.GetSession("userinfo")
		this.Data["uname"] = UserInfo.(*m.User).Username
	}
	this.Data["ipaddress"] = this.GetClientip()
	this.Data["serverurl"] = beego.AppConfig.String("localServerAdress")
	this.TplName = "haoadmin/data/qs/sendbroad.html"
}

// 发送广播
func (this *QsController) SendBroadHandle() {
	code, _ := beego.AppConfig.Int("company")
	room, _ := beego.AppConfig.Int("room")

	UserInfo := this.GetSession("userinfo")
	uname := UserInfo.(*m.User).Username
	data := this.GetString("content")

	broad := new(m.Broadcast)
	broad.Code = code
	broad.Room = room
	broad.Uname = uname
	broad.Datatime = time.Now()
	broad.Data = data
	_, err := m.AddBroadcast(broad)
	if err != nil {
		this.Rsp(false, "广播发送失败", "")
		beego.Debug(err)
	} else {
		msgtype := mqtt.NewMessageType(mqtt.MSG_TYPE_BROCAST)
		b := msgtype.SendBrocast(data)
		if b {
			this.Rsp(true, "广播发送成功", "")
			return
		}
		this.Rsp(false, "广播发送失败", "")
	}
}

//获取客户的真是IP地址
func (this *QsController) GetClientip() string {
	var addrArr []string
	if len(this.Ctx.Request.Header.Get("X-Forwarded-For")) > 0 {
		addr := this.Ctx.Request.Header.Get("X-Forwarded-For")
		addrArr = strings.Split(addr, ":")
	} else if len(this.Ctx.Request.RemoteAddr) > 0 {
		addr := this.Ctx.Request.RemoteAddr
		addrArr = strings.Split(addr, ":")
	} else {
		addrArr[0] = "127.0.0.1"
	}
	return addrArr[0]
}
