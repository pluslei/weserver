package models

import (
	//"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//节点表
type Node struct {
	Id     int64
	Title  string  `orm:"size(128)" form:"Title"  valid:"Required"`          //显示名称
	Name   string  `orm:"size(128)" form:"Name"  valid:"Required"`           //Name 应用名称
	Level  int     `orm:"default(1)" form:"Level"  valid:"Required"`         //层次
	Pid    int64   `form:"Pid"  valid:"Required"`                            //上级ID
	Remark string  `orm:"null;size(255)" form:"Remark" valid:"MaxSize(255)"` //备注
	Status int     `orm:"default(1)" form:"Status" valid:"Range(1,2)"`       //状态 [1、开启 2、关闭]
	Url    string  `orm:"size(128)"`                                         //url地址
	Hide   int     `orm:"default(1)"`                                        //是否显示 [1、显示，2、隐藏]
	Sort   int     `orm:"default(50)"`                                       //排序
	Ico    string  //ico图表
	Group  *Group  `orm:"rel(fk)"` //分组
	Role   []*Role `orm:"rel(m2m)"`
}

func (n *Node) TableName() string {
	return "node"
}

func init() {
	orm.RegisterModel(new(Node))
}

//get node list
func GetNodelist(page int64, page_size int64, sort string, pid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(Node)
	qs := o.QueryTable(node)
	count, _ = qs.Limit(page_size, page).OrderBy(sort).Filter("Pid", pid).Values(&nodes) //, "Id", "Title", "Group_id", "Url", "Sort", "Hide"
	return nodes, count
}

func ReadNode(nid int64) (Node, error) {
	o := orm.NewOrm()
	node := Node{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddNode(n *Node) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(n)
	return id, err
}

// 修改节点
func EditNode(n *Node) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(n)
	return id, err
}

//更新用户
func (this *Node) UpdateNode01(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func DelNodeById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Node{Id: Id})
	return status, err
}

func GetNodelistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(Node)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Filter("Hide", 1).Values(&nodes)
	return nodes, count
}

func GetNodeTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(Node)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 1).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

func GetNodeGroupTree(pid int64, level int64, gid int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(Node)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 1).Filter("Group", gid).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

func GetAllNode() (count int64, nodes []orm.Params) {
	o := orm.NewOrm()
	node := Node{}
	qs := o.QueryTable(node).Filter("group__id", "1").Filter("Level__in", []int{1, 2}).Filter("Hide", 1)
	count, err := qs.OrderBy("level", "sort").Values(&nodes)
	if err != nil {
		beego.Error(err)
	}
	return count, nodes
}

func GetNodeByRoleId(Id int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(Node)
	count, err := o.QueryTable(node).Filter("group__id", "1").Filter("Role__Role__Id", Id).Filter("Hide", 1).Values(&nodes)
	if err != nil {
		beego.Error(err)
	}
	return nodes, count
}

func GetNodeGroupWebTree() ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(Node)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Status", 1).Filter("Group", 4).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

func GetNodeGroupWebTree1() ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(Node)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Status", 1).Filter("Group", 4).Filter("Pid__gt", 0).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

func GetResourcesByRoleId(Nid int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(Node)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Status", 1).Filter("Group", 4).Filter("Role__Role__Id", Nid).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

func GetNodeCount() (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable(new(Node)).Count()
}
