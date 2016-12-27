package models

import (
	// "errors"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//分组表
type Theme struct {
	Id     int64  `orm:"pk;auto"`
	Name   string `orm:"size(128)" form:"Name"  valid:"Required"`
	Status int    `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //状态
	Img    string `orm:"size(128)" form:"Name"  valid:"Required"`     //主题预览
}

func (g *Theme) TableName() string {
	return "theme"
}

func init() {
	orm.RegisterModel(new(Theme))
}

//get title list
func GetThemelist(page int64, page_size int64, sort string) (q []orm.Params, count int64) {
	o := orm.NewOrm()
	t := new(Theme)
	query := o.QueryTable(t)
	query.Limit(page_size, page).OrderBy(sort).Values(&q)
	count, _ = query.Count()
	return q, count
}

func AddTheme(t *Theme) (int64, error) {
	o := orm.NewOrm()
	theme := new(Theme)
	theme.Name = t.Name
	theme.Status = 1
	theme.Img = t.Img

	id, err := o.Insert(theme)
	return id, err
}

func DelThemeById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Theme{Id: Id})
	return status, err
}

func GetThemeList() (q []orm.Params) {
	o := orm.NewOrm()
	t := new(Theme)
	query := o.QueryTable(t)
	query.Values(&q)
	return q
}

func ReadThemeById(Id int64) (Theme, error) {
	o := orm.NewOrm()
	t := Theme{Id: Id}
	err := o.Read(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func ReadThemeByStatus() (themes []orm.Params, count int64) {
	o := orm.NewOrm()
	theme := new(Theme)
	qs := o.QueryTable(theme)
	count, _ = qs.Filter("Status", 2).Values(&themes)
	return themes, count
}

func (this *Theme) UpdateTheme(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}
