package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*
* 公告信息表
 */
type Notice struct {
	Id       int64  `orm:"pk;auto"`
	Room     string //房间号 topic
	Uname    string `orm:"size(128)" form:"Uname" valid:"Required"` //操作者的用户名
	Nickname string
	Data     string    `orm:"type(text)"` //消息内容
	Time     string    //前端发送时间
	Datatime time.Time `orm:"type(datetime)"` //服务器写入时间
}

func init() {
	orm.RegisterModel(new(Notice))
}

func (b *Notice) TableName() string {
	return "notice"
}

func GetNoticeListCount(count int64, room string) ([]Notice, int64, error) {
	o := orm.NewOrm()
	var info []Notice
	num, err := o.QueryTable("notice").Filter("Room", room).OrderBy("-Id").Limit(count).All(&info)
	return info, num, err
}

// 获取指定房间公告列表
func GetNoticeList(room string) ([]Notice, int64, error) {
	o := orm.NewOrm()
	var info []Notice
	num, err := o.QueryTable("notice").Filter("Room", room).OrderBy("-Id").All(&info)
	return info, num, err
}

//获取所有的公告列表
func GetAllNoticeList(page int64, page_size int64, sort string) (broad []orm.Params, count int64) {
	o := orm.NewOrm()
	obj := new(Notice)
	qs := o.QueryTable(obj)
	qs.Limit(page_size, page).OrderBy(sort).Values(&broad)
	count, _ = qs.Count()
	return broad, count
}

/*
*添加公告
 */
func AddNoticeMsg(n *Notice) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(n)
	return id, err
}

//删除公告
func DelNoticeById(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Notice
	status, err := o.QueryTable(info).Filter("Id", id).Delete()
	return status, err
}

//获取内容
func GetNoticeData(codeid int) (string, error) {
	o := orm.NewOrm()
	var notice Notice
	err := o.QueryTable(notice).Filter("Room", codeid).OrderBy("-Id").Limit(1).One(&notice, "Id", "Data")
	return notice.Data, err
}
