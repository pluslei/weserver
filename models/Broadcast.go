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
type Broadcast struct {
	Id        int64     `orm:"pk;auto"`
	Coderoom  int       //房间号
	Uname     string    `orm:"size(128)" form:"Uname" valid:"Required"` //操作者的用户名
	Data      string    `orm:"type(text)"`                              //消息内容
	Ipaddress string    //IP地址
	Procities string    //省市
	Datatime  time.Time `orm:"type(datetime)"` //添加时间
}

func init() {
	orm.RegisterModel(new(Broadcast))
}

func (b *Broadcast) TableName() string {
	return "broadcast"
}

//get user list
func GetBroadcastlist(page int64, page_size int64, sort string) (broad []orm.Params, count int64) {
	o := orm.NewOrm()
	obj := new(Broadcast)
	qs := o.QueryTable(obj)
	qs.Limit(page_size, page).OrderBy(sort).Values(&broad)
	count, _ = qs.Count()
	return broad, count
}

/*
*添加聊天消息
 */
func AddBroadcast(b *Broadcast) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(b)
	return id, err
}

//获取内容
func GetBroadcastData(codeid int) (string, error) {
	o := orm.NewOrm()
	var broad Broadcast
	err := o.QueryTable(broad).Filter("Coderoom", codeid).OrderBy("-Id").Limit(1).One(&broad, "Id", "Data")
	return broad.Data, err
}
