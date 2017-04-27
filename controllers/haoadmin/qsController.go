package haoadmin

import (
	"strings"
	"time"
	m "weserver/models"
	. "weserver/src/tools"

	mq "weserver/src/mqtt"

	"github.com/astaxie/beego"
)

type QsController struct {
	CommonController
}

// 发送公告列表
func (this *QsController) SendNoticeList() {
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
		Noticelist, count := m.GetAllNoticeList(iStart, iLength, "-Id")
		for _, item := range Noticelist {
			item["Datatime"] = item["Datatime"].(time.Time).Format("2006-01-02 15:04:05")
			roomInfo, err := m.GetRoomInfoByRoomID(item["Room"].(string))
			if err != nil {
				item["Room"] = "未知房间"
			} else {
				item["Room"] = roomInfo.RoomTitle
			}
		}
		// json
		data := make(map[string]interface{})
		data["aaData"] = Noticelist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonController.CommonMenu()
		this.TplName = "haoadmin/data/qs/list.html"
	}
}

// 发送公告
func (this *QsController) SendBroad() {
	action := this.GetString("action")
	if action == "add" {
		UserInfo := this.GetSession("userinfo").(*m.User)
		data := this.GetString("Content")
		room := this.GetString("Room")
		filename := this.GetString("FileNameFile")

		broad := new(m.Notice)
		broad.Nickname = UserInfo.Nickname
		broad.Room = room
		broad.Uname = UserInfo.Username
		broad.Datatime = time.Now()
		broad.Data = data
		broad.FileName = filename
		_, err := m.AddNoticeMsg(broad)
		if err != nil {
			this.AlertBack("公告写入数据库失败")
			return
		} else {
			b := SendBrocast(room, data)
			if b {
				this.Alert("公告发送成功", "/weserver/data/qs_broad")
				return
			}
			this.AlertBack("公告添加失败")
		}
	} else {
		this.CommonController.CommonMenu()

		roonInfo, _, err := m.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/qs/add.html"
	}
}

// 编辑
func (this *QsController) Edit() {
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		this.AlertBack("数据错误")
		return
	}
	if action == "edit" {
		notice := make(map[string]interface{})
		notice["Data"] = this.GetString("Content")
		notice["FileName"] = this.GetString("FileNameFile")

		_, err := m.UpdateNoticeData(id, notice)
		if err != nil {
			this.AlertBack("更新失败")
			return
		}
		this.Alert("更新成功", "qs_broad")
	} else {
		this.CommonController.CommonMenu()

		roonInfo, _, err := m.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		notice, err := m.GetNoticeInfoByID(id)
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		this.Data["notice"] = notice
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/qs/add.html"
	}
}

// 删除
func (this *QsController) Del() {
	id, err := this.GetInt64("id")
	if err != nil {
		this.Rsp(false, "删除失败", "")
	} else {
		_, err := m.DelNoticeById(id)
		if err != nil {
			this.Rsp(false, "删除失败", "")
			return
		}
		this.Rsp(true, "删除成功", "")
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

// 发送公告消息
func SendBrocast(topic, content string) bool {
	info := new(NoticeInfo)
	info.Content = content
	info.MsgType = MSG_TYPE_NOTICE_ADD
	v, err := ToJSON(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}
	mq.SendMessage(topic, string(v)) //发消息
	return true
}
