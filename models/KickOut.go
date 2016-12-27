package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//角色表
type KickOut struct {
	Id        int64
	Coderoom  int       //房间号
	Uname     string    `orm:"size(128)" form:"Uname" valid:"Required"`    //操作的用户名
	Objname   string    `orm:"size(128)" form:"Objname"  valid:"Required"` //被提出的用户名
	Kicktime  int64     //禁言的时间戳
	Status    int       `orm:"default(0)" form:"Status" valid:"Range(0,1)"` //状态 [1、正常 0、异常]
	Ipaddress string    //IP地址
	Procities string    //省市
	Datatime  time.Time `orm:"type(datetime)"` //添加时间
}

func (r *KickOut) TableName() string {
	return "kick_out"
}

func init() {
	orm.RegisterModel(new(KickOut))
}

//添加黑名单信息
func AddKickOut(k *KickOut) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(k)
	return id, err
}

// 根据用户名查找
func GetUserByUname(username string) (kickout KickOut) {
	kickout = KickOut{Uname: username}
	o := orm.NewOrm()
	o.Read(&kickout, "Username")
	return kickout
}

func SelectKickOut(name string) (k KickOut, err error) {
	o := orm.NewOrm()
	k = KickOut{Objname: name}
	err = o.Read(&k, "Objname")
	if err != nil {
		return k, err
	}
	return k, nil
}

//获取内容
func UpdateKickControl(k KickOut) (int64, error) {
	o := orm.NewOrm()
	var table KickOut
	id, err := o.QueryTable(table).Filter("Id", k.Id).Update(orm.Params{
		"Uname":     k.Uname,
		"Kicktime":  k.Kicktime,
		"Procities": k.Procities,
		"Ipaddress": k.Ipaddress,
		"Status":    k.Status,
		"Coderoom":  k.Coderoom,
		"Datatime":  k.Datatime})
	return id, err
}
