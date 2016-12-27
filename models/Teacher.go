package models

import (
	// "errors"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//分组表
type Teacher struct {
	Id      int64
	Name    string `orm:"size(128)" form:"Name"  valid:"Required"`
	Title   string `orm:"size(128)" form:"Title"  valid:"Required"`
	Detail  string `orm:"size(511)" form:"Detail"  valid:"Required"`
	Display int    `orm:"default(0)" form:"Display" valid:"Required;Range(0,1)"` //显示状态(0：不显示，1：显示)
	Image   string `orm:"size(511)" form:"Image"  valid:"Required"`
}

func (g *Teacher) TableName() string {
	return "teacher"
}

func init() {
	orm.RegisterModel(new(Teacher))
}

//get title list
func GetTeacherList(page int64, page_size int64, sort string) (teachers []orm.Params, count int64) {
	o := orm.NewOrm()
	teacher := new(Teacher)
	qs := o.QueryTable(teacher)
	qs.Limit(page_size, page).OrderBy(sort).Values(&teachers)
	count, _ = qs.Count()
	return teachers, count
}

func AddTeacher(t *Teacher) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(t)
	return id, err
}

func (this *Teacher) UpdateTeacher(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func DelTeacherById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Teacher{Id: Id})
	return status, err
}

func TeacherList() (teachers []Teacher, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("teacher").All(&teachers)
	return teachers, err
}

func IndexTeacherList() (teachers []Teacher, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("teacher").Filter("Display", 1).All(&teachers)
	return teachers, err
}

func ReadTeacherById(id int64) (Teacher, error) {
	o := orm.NewOrm()
	teacher := Teacher{Id: id}
	err := o.Read(&teacher)
	if err != nil {
		return teacher, err
	}
	return teacher, nil
}

func QueryDisplayCount() (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("teacher").Filter("Display", 1).Count()
	if err != nil {
		return num, err
	}
	return num, nil
}

// func GetTeacherByUserId(id int64) (teachers []orm.Params, count int64) {
// 	o := orm.NewOrm()
// 	teacher := new(Teacher)
// 	count, _ = o.QueryTable(teacher).Filter("User__User__id", userId).Values(&titles)
// 	return titles, count
// }

// // 关联查询所有用户的头衔
// func GetRelationTitle(userid int64) (titles []Title, err error) {
// 	o := orm.NewOrm()
// 	_, err = o.QueryTable("title").Filter("Id", userid).All(&titles, "Id")
// 	return titles, err
// }

// // 增加头衔给用户
// func AddUserTitle(userid int64, titleid int64) (int64, error) {
// 	o := orm.NewOrm()
// 	title := Title{Id: titleid}
// 	user := User{Id: userid}
// 	m2m := o.QueryM2M(&user, "Title")
// 	num, err := m2m.Add(&title)
// 	return num, err
// }

// // 删除用户的角色
// func DelUserTitle(userid int64) error {
// 	o := orm.NewOrm()
// 	_, err := o.QueryTable("user_titles").Filter("user_id", userid).Delete()
// 	return err
// }
