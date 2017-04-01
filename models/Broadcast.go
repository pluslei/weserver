package models

import (
	//"github.com/astaxie/beego"
	"time"

	"github.com/astaxie/beego/orm"
)

/*
* 发送广播表
 */
type Broadcast struct {
	Id       int64     `orm:"pk;auto"`
	Code     int       //公司代码
	Room     string    //房间号 topic
	Uname    string    `orm:"size(128)" form:"Uname" valid:"Required"` //操作者的用户名
	Data     string    `orm:"type(text)"`                              //消息内容
	Datatime time.Time `orm:"type(datetime)"`                          //添加时间
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
	err := o.QueryTable(broad).Filter("Room", codeid).OrderBy("-Id").Limit(1).One(&broad, "Id", "Data")
	return broad.Data, err
}
