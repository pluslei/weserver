package haoadmin

import (
	"github.com/astaxie/beego"
	"time"
	"weserver/controllers/haoindex"
	m "weserver/models"
)

type ChatRecordController struct {
	CommonController
}

func (this *ChatRecordController) ChatRecordList() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		chatrecord, count := m.GetChatRecordList(iStart, iLength, "Id")
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
		url := "http://" + this.Ctx.Request.Host + this.Ctx.Input.URI()

		jssdk := haoindex.Wx.GetJs(this.Ctx.Request, this.Ctx.ResponseWriter)
		jsapi, err := jssdk.GetConfig(url)
		if err != nil {
			beego.Error("get the jsapi config error", err)
		}
		this.Data["appId"] = haoindex.APPID
		this.Data["timestamp"] = jsapi.TimeStamp //jsapi.Timestamp
		this.Data["nonceStr"] = jsapi.NonceStr   //jsapi.NonceStr
		this.Data["signature"] = jsapi.Signature //jsapi.Signature

		this.CommonMenu()
		this.TplName = "haoadmin/data/chat/list.html"
	}
}
