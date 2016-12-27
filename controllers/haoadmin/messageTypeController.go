package haoadmin

import (
	"github.com/astaxie/beego"
	"time"
	m "weserver/models"
)

type MessageTypeController struct {
	CommonController
}

// 消息分类列表
func (this *MessageTypeController) MessageTypeList() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		mts, count := m.GetMessageTypeListByPager(iStart, iLength, "Id")
		for _, v := range mts {
			v["CreateTimeFormat"] = v["CreateTime"].(time.Time).Format("2006-01-02 15:04:05")
		}
		// json
		data := make(map[string]interface{})
		data["aaData"] = mts
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/messagetype/list.html"
	}
}

// 删除消息分类
func (this *MessageTypeController) DeleteMessageType() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelMessageType(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除消息分类成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

// 增加消息分类
func (this *MessageTypeController) AddMessageType() {
	user := this.GetSession("userinfo")
	username := user.(*m.User).Username
	name := this.GetString("Name")
	level, _ := this.GetInt64("Level")
	mt := new(m.MessageType)
	mt.Name = name
	mt.Level = level
	mt.CreateTime = time.Now()
	mt.CreateMan = username
	id, err := m.AddMessageType(mt)
	if err != nil && id <= 0 {
		beego.Error(err)
		this.Rsp(false, "添加消息分类失败", "")
		return
	}
	this.Rsp(true, "添加消息分类成功", "")
}

// 根据 Id 获取消息分类
func (this *MessageTypeController) GetMessageTypeById() {
	Id, _ := this.GetInt64("Id")
	mt, err := m.GetMessageTypeById(Id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = mt
	this.ServeJSON()
}

// 修改消息分类
func (this *MessageTypeController) EditMessageType() {
	user := this.GetSession("userinfo")
	username := user.(*m.User).Username
	Id, _ := this.GetInt64("Id")
	name := this.GetString("Name")
	level, _ := this.GetInt64("Level")
	mt, err := m.GetMessageTypeById(Id)
	if err != nil {
		beego.Error(err)
	}
	mt.Name = name
	mt.Level = level
	mt.CreateTime = time.Now()
	mt.CreateMan = username
	id, err := m.EditMessageType(&mt)
	if err != nil && id <= 0 {
		beego.Error(err)
		this.Rsp(false, "修改消息分类失败", "")
		return
	}
	this.Rsp(true, "修改消息分类成功", "")
}
