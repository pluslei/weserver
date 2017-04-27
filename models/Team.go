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

// 讲师分页
func GetTeacherInfoList(page int64, page_size int64, sort string) (t []orm.Params, count int64) {
	o := orm.NewOrm()
	teacher := new(Teacher)
	query := o.QueryTable(teacher)
	query.Limit(page_size, page).OrderBy(sort).Values(&t)
	count, _ = query.Count()
	return t, count
}

// 根据id查询
func GetTeacherInfoById(id int64) (t Teacher, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(Teacher)).Filter("Id", id).One(&t)
	return t, err
}

// 更新
func UpdateTeacherInfo(id int64, teacher map[string]interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable(new(Teacher)).Filter("Id", id).Update(teacher)
}
