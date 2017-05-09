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

func GetCompanyList() ([]Company, int64, error) {
	o := orm.NewOrm()
	var info []Company
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
