package models

import (
	//"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

/*
* 点赞
 */
type ThumbInfo struct {
	Id       int64  `orm:"pk;auto"`
	Room     string //房间号 topic
	Nickname string
	Username string
	Timestr  string
	IsThumb  bool

	Teacher *Teacher `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(ThumbInfo))
}

func (b *ThumbInfo) TableName() string {
	return "thumbinfo"
}

/*
*增加老师
 */
func AddThumbInfo(t *ThumbInfo) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(t)
	return id, err
}

func GetMoreThumbInfo(Room string, id int64) ([]*ThumbInfo, int64, error) {
	o := orm.NewOrm()
	var info []*ThumbInfo
	num, err := o.QueryTable(new(ThumbInfo)).Filter("Room", Room).Filter("Teacher", id).RelatedSel().All(&info)
	return info, num, err
}

func GetThumbInfo(Uname, Room string, id int64) (ThumbInfo, error) {
	o := orm.NewOrm()
	var info ThumbInfo
	err := o.QueryTable("thumbinfo").Filter("Room", Room).Filter("Username", Uname).Filter("Teacher", id).One(&info)
	return info, err
}

func UpdateThumb(id int64) (int64, error) {
	o := orm.NewOrm()
	var info ThumbInfo
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"IsThumb": true})
	return id, err
}

func UpdateUnThumb(Uname, Room string, id int64) (int64, error) {
	o := orm.NewOrm()
	var info ThumbInfo
	id, err := o.QueryTable(info).Filter("Room", Room).Filter("Username", Uname).Filter("Teacher", id).Update(orm.Params{"IsThumb": false})
	return id, err
}
