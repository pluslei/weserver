package haoadmin

import (
	"github.com/astaxie/beego"
	m "weserver/models"
)

type GroupController struct {
	CommonController
}

// 组别显示页面
func (this *GroupController) Index() {
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
		grouplist, count := m.GetGrouplist(iStart, iLength, "Id")

		// json
		data := make(map[string]interface{})
		data["aaData"] = grouplist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/rbac/group/list.html"
	}

}

// 增加组别
func (this *GroupController) AddGroup() {
	action := this.GetString("action")
	if action == "add" {
		g := m.Group{}
		if err := this.ParseForm(&g); err != nil {
			this.AlertBack("验证失败")
			return
		}
		id, err := m.AddGroup(&g)
		if err == nil && id > 0 {
			this.Alert("添加成功", "index")
			return
		} else {
			this.AlertBack("添加失败")
			return
		}
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/rbac/group/add.html"
	}

}

// 更新组别
func (this *GroupController) UpdateGroup() {
	action := this.GetString("action")
	if action == "edit" {
		g := new(m.Group)
		if err := this.ParseForm(g); err != nil {
			this.AlertBack("验证信息失败")
			return
		}
		err := g.UpdateGroup("Name", "Title", "Status", "Sort")
		if err == nil {
			this.Alert("更新成功", "index")
			return
		} else {
			this.AlertBack("更新失败")
			return
		}
	} else {
		this.CommonMenu()
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.AlertBack("Id获取失败")
			return
		}
		groupList, err := m.ReadGroupById(id)
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取组别信息错误")
			return
		}
		this.Data["groupList"] = groupList
		this.TplName = "haoadmin/rbac/group/edit.html"
	}

}

// 删除组别
func (this *GroupController) DelGroup() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelGroupById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}
