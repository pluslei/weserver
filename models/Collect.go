package models

import (
	//"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

/*
* 收藏和浏览表
 */
type Collect struct {
	Id          int64  `orm:"pk;auto"`
	Uname       string //openid
	Nickname    string
	RoomIcon    string //房间图标
	RoomTitle   string //房间名
	RoomTeacher string //老师
	IsCollect   bool   //是否收藏 收藏：1 浏览：0
	IsAttention bool   //是否关注 收藏：1 浏览：0

}

func init() {
	orm.RegisterModel(new(Collect))
}

func (c *Collect) TableName() string {
	return "Collect"
}

// 获取收藏列表
func GetCollectList(Uname string) ([]Collect, int64, error) {
	o := orm.NewOrm()
	var info []Collect
	num, err := o.QueryTable("Collect").Filter("Uname", Uname).Filter("IsCollect", 1).Filter("IsAttention", 1).OrderBy("Id").All(&info)
	return info, num, err
}

// 获取浏览列表
func GetScanList(Uname string) ([]Collect, int64, error) {
	o := orm.NewOrm()
	var info []Collect
	num, err := o.QueryTable("Collect").Filter("Uname", Uname).Filter("IsCollect", 0).Filter("IsAttention", 0).OrderBy("Id").All(&info)
	return info, num, err
}

/*
*增加记录
 */
func AddCollect(c *Collect) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(c)
	return id, err
}

//删除记录
func DelCollectById(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Collect
	status, err := o.QueryTable(info).Filter("Id", id).Delete()
	return status, err
}
