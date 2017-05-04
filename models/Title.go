package models

import (
	// "errors"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"time"
)

//分组表
type Title struct {
	Id         int64
	Name       string `orm:"size(128)" form:"Name"  valid:"Required"`
	Css        string `orm:"size(128)" form:"Css"  valid:"Required"`
	Background int    `orm:"default(0)" form:"Background"  valid:"Required"`    //背景(0 普通，1 高亮)
	Weight     int    `orm:"default(1)" form:"Weight" valid:"Range(1,2)"`       //权重
	Remark     string `orm:"null;size(255)" form:"Remark" valid:"MaxSize(255)"` //备注
	User       *User  `orm:"reverse(one)"`
}

func (g *Title) TableName() string {
	return "title"
}

func init() {
	orm.RegisterModel(new(Title))
}

//get title list
func GetTitlelist(page int64, page_size int64, sort string) (titles []orm.Params, count int64) {
	o := orm.NewOrm()
	title := new(Title)
	qs := o.QueryTable(title)
	qs.Limit(page_size, page).OrderBy(sort).Values(&titles)
	count, _ = qs.Count()
	return titles, count
}

func AddTitle(t *Title) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(t)
	return id, err
}

func GetTitleCount() (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable(new(Title)).Count()
}

func (this *Title) UpdateTitle(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func DelTitleById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Title{Id: Id})
	return status, err
}

func TitleList() (titles []orm.Params) {
	o := orm.NewOrm()
	title := new(Title)
	qs := o.QueryTable(title)
	qs.Values(&titles)
	return titles
}

func ReadTitleById(gid int64) (Title, error) {
	o := orm.NewOrm()
	title := Title{Id: gid}
	err := o.Read(&title)
	if err != nil {
		return title, err
	}
	return title, nil
}

func GetTitleByUserId(userId int64) (titles []orm.Params, count int64) {
	o := orm.NewOrm()
	title := new(Title)
	count, _ = o.QueryTable(title).Filter("User__id", userId).Values(&titles)
	return titles, count
}

// 关联查询所有用户的头衔
func GetRelationTitle(userid int64) (titles []Title, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("title").Filter("Id", userid).All(&titles, "Id")
	return titles, err
}

// 获取所有的角色
func GetAllUserTitle() (titles []Title, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("title").All(&titles)
	return titles, err
}

// 增加头衔给用户
func AddUserTitle(userid int64, titleid int64) (int64, error) {
	o := orm.NewOrm()
	title := Title{Id: titleid}
	user := User{Id: userid}
	m2m := o.QueryM2M(&user, "Title")
	num, err := m2m.Add(&title)
	return num, err
}

// 删除用户的头衔
func DelUserTitle(userid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user_titles").Filter("user_id", userid).Delete()
	return err
}
