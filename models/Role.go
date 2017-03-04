package models

import (
	//"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//角色表
type Role struct {
	Id        int64
	Title     string `orm:"size(128)" form:"Title"  valid:"Required"`
	Name      string `orm:"size(128)" form:"Name"  valid:"Required"`
	Remark    string `orm:"null;size(255)" form:"Remark" valid:"MaxSize(255)"`
	Status    int    `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //状态 [1、开启 2、关闭]
	Weight    int    `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //权重 [1、开启 2、关闭]
	Delay     int    `orm:"default(0)" form:"Status" valid:"Range(1,2)"` //发言间隔 [1、开启 2、关闭]
	IsInsider int    `orm:"default(0)" form:"Status" valid:"Range(0,1)"` //是否隶属公司内部角色[0、否 1、是]
	// Randnum   int     `orm:"default(0)"`                                  //随机人数
	Ico string //角色的头像
	// RandTitle *Title  `orm:"rel(one)"` //随机头衔
	User *User   `orm:"reverse(one)"`
	Node []*Node `orm:"reverse(many)"`
}

func (r *Role) TableName() string {
	return "role"
}

func init() {
	orm.RegisterModel(new(Role))
}

//get role list
func GetRolelist(page int64, page_size int64, sort string) (roles []orm.Params, count int64) {
	o := orm.NewOrm()
	role := new(Role)
	qs := o.QueryTable(role)
	qs.Limit(page_size, page).OrderBy(sort).Values(&roles)
	count, _ = qs.Count()

	return roles, count
}

func AddRole(r *Role) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(r)
	return id, err
}

func (r *Role) UpdateRoleFields(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}

func DelRoleById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Role{Id: Id})
	return status, err
}

func GetNodelistByRoleId(Id int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(Node)
	count, _ = o.QueryTable(node).Filter("Role__Role__Id", Id).Values(&nodes)
	return nodes, count
}

func DelGroupNode(roleid int64, groupid int64) error {
	var nodes []*Node
	var node Node
	role := Role{Id: roleid}
	o := orm.NewOrm()
	num, err := o.QueryTable(node).Filter("Group", groupid).RelatedSel().All(&nodes)
	if err != nil {
		return err
	}
	if num < 1 {
		return nil
	}
	for _, n := range nodes {
		m2m := o.QueryM2M(n, "Role")
		_, err1 := m2m.Remove(&role)
		if err1 != nil {
			return err1
		}
	}
	return nil
}

func AddRoleNode(roleid int64, nodeid int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: roleid}
	node := Node{Id: nodeid}
	m2m := o.QueryM2M(&node, "Role")
	num, err := m2m.Add(&role)
	return num, err
}

// 删除用户上所有的角色
func DelUserRole(userid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user_roles").Filter("user_id", userid).Delete()
	return err
}

func GetUserByRoleId(roleid int64) (nodelist []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(Node)
	count, _ = o.QueryTable(node).Filter("Role__Role__Id", roleid).Values(&nodelist)
	return nodelist, count
}

func AccessList(uid int64) (list []orm.Params, err error) {
	var roles []orm.Params
	o := orm.NewOrm()
	role := new(Role)
	_, err = o.QueryTable(role).Filter("User", uid).Values(&roles)
	if err != nil {
		return nil, err
	}
	var nodes []orm.Params
	node := new(Node)
	for _, r := range roles {
		_, err := o.QueryTable(node).Filter("Role__Role__Id", r["Id"]).Values(&nodes)
		if err != nil {
			return nil, err
		}
		for _, n := range nodes {
			list = append(list, n)
		}
	}
	return list, nil
}

func GetRoleByUserId(userId int64) (roles Role, err error) {
	o := orm.NewOrm()
	role := new(Role)
	err = o.QueryTable(role).Filter("User__Id", userId).One(&roles)
	return roles, err

}

// 根据Id获取
func GetRoleInfoById(roleId int64) (roles Role, err error) {
	o := orm.NewOrm()
	role := new(Role)
	err = o.QueryTable(role).Filter("Id", roleId).One(&roles)
	if err != nil {
		beego.Error(err)
	}
	return roles, err
}

// 获取用户
func ReadFieldRole(r *Role, fields ...string) (*Role, error) {
	o := orm.NewOrm()
	err := o.Read(r, fields...)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return r, nil
}

// 获取所有的角色
func GetAllUserRole() (roles []Role, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("role").All(&roles)
	return roles, err
}

// 获取所有的角色
func GetRoleId(name string) (id int64, err error) {
	o := orm.NewOrm()
	var roles Role
	err = o.QueryTable("role").Filter("Name", name).One(&roles, "Id")
	return roles.Id, err
}
