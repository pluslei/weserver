package models

import (
	//"errors"
	"github.com/astaxie/beego/orm"
)

// 节点分组表
type Group struct {
	Id     int64
	Name   string  `orm:"size(128)" form:"Name"  valid:"Required"`
	Title  string  `orm:"size(128)" form:"Title"  valid:"Required"`
	Status int     `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //状态(1、正常，2、关闭)
	Sort   int     `orm:"default(50)" form:"Sort"`                     //默认sort
	Nodes  []*Node `orm:"reverse(many)"`
}

func (g *Group) TableName() string {
	return "group"
}

func init() {
	orm.RegisterModel(new(Group))
}

//get group list
func GetGrouplist(page int64, page_size int64, sort string) (groups []orm.Params, count int64) {
	o := orm.NewOrm()
	group := new(Group)
	qs := o.QueryTable(group)
	qs.Limit(page_size, page).OrderBy(sort).Values(&groups)
	count, _ = qs.Count()
	return groups, count
}

func AddGroup(g *Group) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(g)
	return id, err
}

func (this *Group) UpdateGroup(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func DelGroupById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Group{Id: Id})
	return status, err
}

func GroupList() (groups []orm.Params) {
	o := orm.NewOrm()
	group := new(Group)
	qs := o.QueryTable(group)
	qs.Values(&groups, "id", "title")
	return groups
}

func ReadGroupById(gid int64) (Group, error) {
	o := orm.NewOrm()
	group := Group{Id: gid}
	err := o.Read(&group)
	if err != nil {
		return group, err
	}
	return group, nil
}

func ReadRoleGroup(nid int64) (Node, error) {
	o := orm.NewOrm()
	node := Node{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}
