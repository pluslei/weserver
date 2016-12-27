package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//角色表
type BlackList struct {
	Id        int64
	Coderoom  int       //房间号
	Uname     string    `orm:"size(128)" form:"Uname" valid:"Required"`      //操作的用户名
	Objname   string    `orm:"size(128)" form:"Objname"  valid:"Required"`   //被拉黑用户
	Status    int       `orm:"default(0)" form:"Status" valid:"Range(0,1)"`  //状态 [1、正常 0、异常]
	Ipaddress string    `orm:"size(128)" form:"IpAddress"  valid:"Required"` //IP地址
	Procities string    `orm:"size(128)" form:"Procities"  valid:"Required"` //省市
	Datatime  time.Time `orm:"type(datetime)"`
}

func (r *BlackList) TableName() string {
	return "black_list"
}

func init() {
	orm.RegisterModel(new(BlackList))
}

//添加黑名单信息
func AddBlackList(b *BlackList) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(b)
	return id, err
}

//删除黑名单信息
func DelBlackListById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&BlackList{Id: Id})
	return status, err
}

//获取内容
func UpdateBlackList(b BlackList) (int64, error) {
	o := orm.NewOrm()
	var table BlackList
	id, err := o.QueryTable(table).Filter("Id", b.Id).Update(orm.Params{
		"Uname":    b.Uname,
		"Status":   b.Status,
		"Coderoom": b.Coderoom,
		"Datatime": b.Datatime,
	})
	return id, err
}

func SelectBlackList(name string) (b BlackList, err error) {
	o := orm.NewOrm()
	b = BlackList{Objname: name}
	err = o.Read(&b, "Objname")
	if err != nil {
		return b, err
	}
	return b, nil
}
