package models

import (
	//"errors"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//分组表
type TeachBanner struct {
	Id      int64
	Name    string `orm:"size(128)" form:"Name" valid:"Required"`
	Img     string `orm:"size(511)" form:"Img"  valid:"Required"`
	Url     string `orm:"size(511)" form:"Url"  valid:"Required"`
	Order   int
	Display int `orm:"default(0)" form:"Display" valid:"Required;Range(0,1)"` //显示状态(0：不显示，1：显示)
}

func (g *TeachBanner) TableName() string {
	return "teachbanner"
}

func init() {
	orm.RegisterModel(new(TeachBanner))
}

//get title list
func GetTeachBannerList(page int64, page_size int64, sort string) (teachs []orm.Params, count int64) {
	o := orm.NewOrm()
	tech := new(TeachBanner)
	qs := o.QueryTable(tech)
	qs.Limit(page_size, page).OrderBy(sort).Values(&teachs)
	count, _ = qs.Count()
	return teachs, count
}

func AddTeachBanner(t *TeachBanner) (int64, error) {
	o := orm.NewOrm()
	teach := new(TeachBanner)
	teach.Name = t.Name
	teach.Img = t.Img
	teach.Display = t.Display
	teach.Url = t.Url
	teach.Order = t.Order
	id, err := o.Insert(teach)
	return id, err
}

func (this *TeachBanner) UpdateTeachBanner(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func DelTeachBannerById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&TeachBanner{Id: Id})
	return status, err
}

func TeachBannerList() (teachs []TeachBanner, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("teachbanner").All(&teachs)
	return teachs, err
}

func IndexTeachBannerList() (tech []orm.Params) {
	// o := orm.NewOrm()
	// _, err = o.QueryTable("teachbanner").Filter("Display", 1).OrderBy("-Order").All(&teachs)
	// return teachs, err

	o := orm.NewOrm()
	banner := new(TeachBanner)
	query := o.QueryTable(banner)
	query.Filter("Display", 1).OrderBy("-Order").Values(&tech)
	return tech
}

func ReadTeachBannerById(id int64) (TeachBanner, error) {
	o := orm.NewOrm()
	teach := TeachBanner{Id: id}
	err := o.Read(&teach)
	if err != nil {
		return teach, err
	}
	return teach, nil
}
