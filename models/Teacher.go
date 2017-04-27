package models

import (
	//"github.com/astaxie/beego"
	"time"

	"github.com/astaxie/beego/orm"
)

/*
* 专家团队
 */
type Teacher struct {
	Id       int64  `orm:"pk;auto"`
	Room     string //房间号 topic
	Name     string
	Icon     string //头像
	Title    string
	Data     string    `orm:"type(text)"` //专家介绍
	Time     string    //前台给的时间
	Datatime time.Time `orm:"type(datetime)"` //添加时间
}

func init() {
	orm.RegisterModel(new(Teacher))
}

func (b *Teacher) TableName() string {
	return "teacher"
}

// 获取指定房间的策略列表
func GetTeacherList(room string) ([]Teacher, int64, error) {
	o := orm.NewOrm()
	var info []Teacher
	num, err := o.QueryTable("teacher").Filter("Room", room).OrderBy("-Id").All(&info)
	return info, num, err
}

/*
*增加老师
 */
func AddTeacher(t *Teacher) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(t)
	return id, err
}

//更新
func UpdateTeacherInfo(t *Teacher) (int64, error) {
	o := orm.NewOrm()
	id, err := o.QueryTable("teacher").Filter("Id", t.Id).Update(orm.Params{
		"Name":  t.Name,
		"Room":  t.Room,
		"Icon":  t.Icon,
		"Title": t.Title,
		"Data":  t.Data,
		"Time":  t.Time,
	})
	return id, err
}

//删除老师
func DelTeacherById(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Teacher
	status, err := o.QueryTable(info).Filter("Id", id).Delete()
	return status, err
}

//更新老师名字
func UpdateTeacherName(id int64, strName string) (int64, error) {
	o := orm.NewOrm()
	var info Teacher
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"Name": strName})
	return id, err
}

//更新老师头衔
func UpdateTeacherTitle(id int64, strTitle string) (int64, error) {
	o := orm.NewOrm()
	var info Teacher
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"Title": strTitle})
	return id, err
}

//更新老师头像
func UpdateTeacherIcon(id int64, strFilePath string) (int64, error) {
	o := orm.NewOrm()
	var info Teacher
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"Icon": strFilePath})
	return id, err
}

//更新内容
func UpdateContent(id int64, strContent string) (int64, error) {
	o := orm.NewOrm()
	var info Teacher
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"Data": strContent})
	return id, err
}
