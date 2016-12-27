package models

import (
	"github.com/astaxie/beego/orm"
)

//分组表
type Course struct {
	Id        int64
	Name      string `orm:"size(128)" valid:"Required"`
	StartTime string `orm:"size(128)" valid:"Required"`
	EndTime   string `orm:"size(128)" valid:"Required"`
	Uname     string `orm:"size(128)" valid:"Required"`
	Monday    int    `orm:"default(0)"`
	Tuesday   int    `orm:"default(0)"`
	Wednesday int    `orm:"default(0)"`
	Thursday  int    `orm:"default(0)"`
	Friday    int    `orm:"default(0)"`
	Saturday  int    `orm:"default(0)"`
	Sunday    int    `orm:"default(0)"`
}

func (g *Course) TableName() string {
	return "course"
}

func init() {
	orm.RegisterModel(new(Course))
}

//get title list
func GetCourseList(page int64, page_size int64, sort string) (courses []orm.Params, count int64) {
	o := orm.NewOrm()
	course := new(Course)
	qs := o.QueryTable(course)
	qs.Limit(page_size, page).OrderBy(sort).Values(&courses)
	count, _ = qs.Count()
	return courses, count
}

func AddCourse(c *Course) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(c)
	return id, err
}

func UpdateCourse(c *Course) (int64, error) {
	o := orm.NewOrm()
	var table Course
	id, err := o.QueryTable(table).Filter("Id", c.Id).Update(orm.Params{
		"Name":      c.Name,
		"StartTime": c.StartTime,
		"EndTime":   c.EndTime,
		"Monday":    c.Monday,
		"Tuesday":   c.Tuesday,
		"Wednesday": c.Wednesday,
		"Thursday":  c.Thursday,
		"Friday":    c.Friday,
		"Saturday":  c.Saturday,
		"Sunday":    c.Sunday,
		"Uname":     c.Uname})
	return id, err
}

func DelCourseById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Course{Id: Id})
	return status, err
}

func CourseList() (courses []orm.Params) {
	o := orm.NewOrm()
	course := new(Course)
	qs := o.QueryTable(course)
	qs.Values(&courses)
	return courses
}

func ReadCourseById(id int64) (Course, error) {
	o := orm.NewOrm()
	course := Course{Id: id}
	err := o.Read(&course)
	if err != nil {
		return course, err
	}
	return course, nil
}
