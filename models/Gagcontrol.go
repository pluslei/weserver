package models

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
*用户表
* beego 中会把名为Id的字段自动设置文自增加的主键
 */
type GagControl struct {
	Id        int64     `orm:"pk;auto"`
	Coderoom  int       //房间号
	Uname     string    `orm:"size(128)" form:"Uname" valid:"Required"`    //操作者的用户名
	Objname   string    `orm:"size(128)" form:"Objname"  valid:"Required"` //被禁言的用户名
	Gagtime   int64     //禁言的时间戳
	Gagmode   int       //禁言的方式
	Gatcount  int64     //禁言的次数
	Gatstate  int       `orm:"default(0)" form:"Gatstate" valid:"Required;Range(0,1)"` //聊天状态(0：禁言，1：没有禁言)
	Ipaddress string    //IP地址
	Procities string    //省市
	Datatime  time.Time `orm:"type(datetime)"` //添加时间
}

func init() {
	orm.RegisterModel(new(GagControl))
}

func (g *GagControl) TableName() string {
	return "gag_control"
}

/*
*添加聊天消息
 */
func AddGagControl(g *GagControl) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(g)
	return id, err
}

func SelectGagControl(name string) (gag GagControl, err error) {
	o := orm.NewOrm()
	gag = GagControl{Objname: name}
	err = o.Read(&gag, "Objname")
	if err != nil {
		return gag, err
	}
	return gag, nil
}

//获取内容
func UpdateGagControl(g GagControl) (int64, error) {
	o := orm.NewOrm()
	var table GagControl
	id, err := o.QueryTable(table).Filter("Id", g.Id).Update(orm.Params{
		"Uname":    g.Uname,
		"Gagtime":  g.Gagtime,
		"Gagmode":  g.Gagmode,
		"Gatcount": g.Gatcount,
		"Gatstate": g.Gatstate,
		"Coderoom": g.Coderoom,
		"Datatime": g.Datatime,
	})
	return id, err
}

//删除数据库中表中ID对应的行信息
func DelGagControl(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&GagControl{Id: Id})
	return status, err
}
