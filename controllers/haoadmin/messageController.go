package haoadmin

import (
	"github.com/astaxie/beego"
	"time"
	m "weserver/models"
)

type MessageController struct {
	CommonController
}

// 分页查询消息列表
func (this *MessageController) MessageList() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		ms, count := m.GetMessageListByPager(iStart, iLength, "Id")
		for _, v := range ms {
			v["CreateTimeFormat"] = v["CreateTime"].(time.Time).Format("2006-01-02 15:04:05")
			mt, err := m.GetMessageTypeById(v["MessageType"].(int64))
			if err != nil {
				beego.Error(err)
			}
			v["TypeName"] = mt.Name
		}
		// json
		data := make(map[string]interface{})
		data["aaData"] = ms
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		mt, err := m.GetMessageTypeList()
		if err != nil {
			beego.Error(err)
		}
		groups, _ := m.GetGroupList()
		length := len(groups)
		for i := 0; i < length; i++ {
			groups[i].GroupFace = FaceImg + groups[i].GroupFace
		}
		this.Data["group"] = groups
		this.Data["messagetype"] = mt
		this.CommonMenu()
		this.TplName = "haoadmin/data/message/list.html"
	}
}

// 增加消息
func (this *MessageController) AddMessage() {
	user := this.GetSession("userinfo")
	username := user.(*m.User).Username
	content := this.GetString("Content")
	typeId, _ := this.GetInt64("MessageType")
	message := new(m.Message)
	message.Content = content
	message.MessageType = &m.MessageType{Id: typeId}
	message.CreateTime = time.Now()
	message.CreateMan = username
	id, err := m.AddMessage(message)
	if err != nil && id <= 0 {
		beego.Error(err)
		this.Rsp(false, "添加消息失败", "")
		return
	}
	this.Rsp(true, "添加消息成功", "")
}

// 根据 Id 获取消息
func (this *MessageController) GetMessageById() {
	Id, _ := this.GetInt64("Id")
	message, err := m.GetMessageById(Id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = message
	this.ServeJSON()
}

// 修改消息分类
func (this *MessageController) EditMessage() {
	user := this.GetSession("userinfo")
	username := user.(*m.User).Username
	Id, _ := this.GetInt64("Id")
	content := this.GetString("Content")
	messagetype, _ := this.GetInt64("MessageType")
	message, err := m.GetMessageById(Id)
	if err != nil {
		beego.Error(err)
	}
	message.Content = content
	message.MessageType = &m.MessageType{Id: messagetype}
	message.CreateTime = time.Now()
	message.CreateMan = username
	id, err := m.EditMessage(&message)
	if err != nil && id <= 0 {
		beego.Error(err)
		this.Rsp(false, "修改消息失败", "")
		return
	}
	this.Rsp(true, "修改消息成功", "")
}

// 删除消息
func (this *MessageController) DeleteMessage() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelMessage(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除消息成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}
