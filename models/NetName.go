package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 网名
type NetName struct {
	Id   int64  // 网名标识
	Name string // 网名
}

var table NetName

// 指定数据库表名
func (n *NetName) TableName() string {
	return "netname"
}

// 初始化网名表
func init() {
	orm.RegisterModel(new(NetName))
}

// 添加网名
func AddNetName(n *NetName) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(n)
	return id, err
}

// 判断网名是否存在
func IsExitNetName(netname string) (error, NetName) {
	o := orm.NewOrm()
	netName := NetName{Name: netname}
	err := o.Read(&netName, "Name")
	return err, netName
}

//检查电话号码是否存在
func CheckIsNetName(netname string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable(table).Filter("Name", netname).Exist()
	return exist
}

//取数据
func GetNetNameData() (netname []NetName, count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable(table).Limit(-1).All(&netname, "Name")
	return netname, count, err
}

//事务添加数据
func AddBeginNetName(net []NetName, length int) error {
	model := orm.NewOrm()
	err := model.Begin()
	SuccessNum := 0
	if err == nil {
		for i := 0; i < length; i++ {
			id, err := model.Insert(&net[i])
			if err == nil && id > 0 {
				SuccessNum++
			}
		}
	} else {
		err = errors.New("NetName commit error")
	}
	if SuccessNum == length {
		err = model.Commit()
		beego.Debug("NetName Commit Complete.")
	} else {
		err = errors.New("NetName Transaction commit failed!")
	}
	return err
}
