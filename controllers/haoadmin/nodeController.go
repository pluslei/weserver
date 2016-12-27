package haoadmin

import (
	"github.com/astaxie/beego"

	m "weserver/models"
)

type NodeController struct {
	CommonController
}

// 节点列表
func (this *NodeController) Index() {
	pid, err := this.GetInt64("pid")
	if err != nil {
		beego.Error(err)
	}
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
		nodelist, count := m.GetNodelist(iStart, iLength, "Sort", pid) //page ,pagesize
		for _, item := range nodelist {
			if item["Pid"].(int64) == 0 {
				item["PidName"] = "无"
			} else {
				node, err := m.ReadNode(item["Pid"].(int64))
				if err != nil {
					beego.Error(err)
					item["PidName"] = "查询失败"
				} else {
					item["PidName"] = node.Title
				}
			}
			group, err := m.ReadGroupById(item["Group"].(int64))
			if err != nil {
				item["GroupName"] = "无分组"
			} else {
				item["GroupName"] = group.Title
			}
		}

		// json
		data := make(map[string]interface{})
		data["aaData"] = nodelist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonMenu()
		this.Data["pid"] = pid
		group := m.GroupList()
		this.Data["group"] = group
		this.TplName = "haoadmin/rbac/node/list.html"
	}

}

// 增加一个节点
func (this *NodeController) AddNode() {
	action := this.GetString("action")
	if action == "add" {
		Name := this.GetString("Name")
		Title := this.GetString("Title")
		Group, _ := this.GetInt64("Group")
		Pid, _ := this.GetInt64("Pid")
		Url := this.GetString("Url")
		Sort, _ := this.GetInt("Sort")
		Status, _ := this.GetInt("Status")
		Hide, _ := this.GetInt("Hide")
		Remark := this.GetString("Remark")

		n := new(m.Node)
		group := new(m.Group)
		group.Id = Group
		n.Group = group
		n.Name = Name
		n.Title = Title
		n.Pid = Pid
		n.Url = Url
		n.Sort = Sort
		n.Status = Status
		n.Hide = Hide
		n.Remark = Remark
		if n.Pid != 0 {
			n1, _ := m.ReadNode(n.Pid)
			n.Level = n1.Level + 1
		} else {
			n.Level = 1
		}
		id, err := m.AddNode(n)
		if err != nil && id <= 0 {
			beego.Error(err)
			this.AlertBack("添加错误")
			return
		}
		this.Alert("添加成功", "index")
	} else {
		this.CommonMenu()
		group := m.GroupList()
		this.Data["group"] = group
		this.TplName = "haoadmin/rbac/node/add.html"
	}
}

// 获取分组节点树
func (this *NodeController) GetNodeTree() {
	groupId, err := this.GetInt64("Gid")
	if err != nil {
		beego.Error(err)
		return
	}
	tree := this.GetGroupTree(groupId)
	if len(tree) <= 0 {
		this.Data["json"] = &map[string]interface{}{"status": false, "tree": ""}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": true, "tree": tree}
	}
	this.ServeJSON()
}

func (this *NodeController) UpdateNode() {
	Id, _ := this.GetInt64("Id")
	nodeName := this.GetString("nodeName")
	nodeTitle := this.GetString("nodeTitle")
	nodeGroup, _ := this.GetInt64("nodeGroup")
	sort, _ := this.GetInt("sort")
	parentNode, _ := this.GetInt64("parentNode")
	status, _ := this.GetInt("status")
	hide, _ := this.GetInt("hide")
	node, err := m.ReadNode(Id)
	if err != nil && node.Id <= 0 {
		beego.Error(err)
		this.Rsp(false, "查询本条数据错误", "")
	}
	node.Name = nodeName
	node.Title = nodeTitle
	node.Pid = parentNode
	node.Hide = hide
	node.Status = status
	node.Sort = sort
	node.Group = &m.Group{Id: nodeGroup}
	if node.Pid != 0 {
		n1, _ := m.ReadNode(node.Pid)
		node.Level = n1.Level + 1
	} else {
		node.Level = 1
	}
	id, err := m.EditNode(&node)
	if err != nil && id <= 0 {
		this.Rsp(false, "修改数据失败！", "")
	}
	this.Rsp(true, "修改数据成功！", "")
}

// 删除节点
func (this *NodeController) DelNode() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelNodeById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

func (this *NodeController) GetNode() {
	Id, _ := this.GetInt64("Id")
	node, err := m.ReadNode(Id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = node
	this.ServeJSON()
}
