package models

import (
	// "github.com/astaxie/beego"
	"time"

	"github.com/astaxie/beego/orm"
)

// 踢人和禁言记录表
type KickOut struct {
	Id        int64
	CompanyId int64
	Coderoom  string    //房间号
	Operuid   string    //踢人uuid
	Opername  string    `orm:"size(128)" form:"OperName" valid:"Required"` //操作的用户名
	Objuid    string    //被踢的uuid
	Objname   string    `orm:"size(128)" form:"Objname"  valid:"Required"`  //被踢出的用户名
	Status    int       `orm:"default(0)" form:"Status" valid:"Range(0,1)"` //状态 [0、踢人 1、禁言 2 取消禁言]
	Opertime  time.Time `orm:"type(datetime)"`                              //操作的时间
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

// 根据Uid查找
func GetUserByUname(Uid string) (kickout KickOut) {
	kickout = KickOut{Operuid: Uid}
	o := orm.NewOrm()
	o.Read(&kickout, "Username")
	return kickout
}

//查找被踢的人
func SelectKickOut(Uid string) (k KickOut, err error) {
	o := orm.NewOrm()
	k = KickOut{Objuid: Uid}
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
		"Coderoom": k.Coderoom,
		"Operuid":  k.Operuid,
		"Opername": k.Opername,
		"Objuid":   k.Objuid,
		"Objname":  k.Objname,
		"Status":   k.Objname,
		"Opertime": k.Opertime})
	return id, err
}
