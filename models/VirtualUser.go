package models

import (
	"github.com/astaxie/beego/orm"
)

//虚拟用户表
type VirtualUser struct {
	Id       int64
	Username string
	Nickname string `orm:"size(255)"`
	UserIcon string `orm:"null;size(255)"`
}

func (v *VirtualUser) TableName() string {
	return "virtual_user"
}

func init() {
	orm.RegisterModel(new(VirtualUser))
}

func GetNumberVirtualUser(count int64) (virtual []VirtualUser, err error) {
	_, err = NewVirtualUser().Limit(count).All(&virtual)
	return virtual, err
}

func NewVirtualUser() orm.QuerySeter {
	o := orm.NewOrm()
	return o.QueryTable(new(VirtualUser))
}
