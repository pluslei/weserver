package models

import (
	//"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

/*
* Company table
 */
type Company struct {
	Id      int64 `orm:"pk;auto"`
	Company string
}

func init() {
	orm.RegisterModel(new(Company))
}

func (c *Company) TableName() string {
	return "company"
}

func GetCompanyList(id int64) ([]Company, int64, error) {
	o := orm.NewOrm()
	var info []Company
	if id != 0 {
		num, err := o.QueryTable("company").Filter("Id", id).All(&info)
		return info, num, err
	}
	num, err := o.QueryTable("company").OrderBy("Id").All(&info)
	return info, num, err
}

func AddCompany(c *Company) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(c)
	return id, err
}

func DelCompanyById(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Company
	status, err := o.QueryTable(info).Filter("Id", id).Delete()
	return status, err
}

func GetCompanyById(id int64) (Company, error) {
	o := orm.NewOrm()
	var info Company
	_, err := o.QueryTable(info).Filter("Id", id).All(&info)
	return info, err
}
