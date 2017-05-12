package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
* Company table
 */
type Company struct {
	Id      int64 `orm:"pk;auto"`
	Company string
	CompanyIntro string
	CompanyIcon string
	CompanyBanner string
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

// 获取公司列表分页
func GetCompanys(page int64, page_size int64) (ms []orm.Params, count int64) {
	o := orm.NewOrm()
	query := o.QueryTable("company")
	query.Limit(page_size, page).OrderBy("id").Values(&ms)
	count, _ = query.Count()
	return ms, count
}

// 更新公司信息
func UpdateCompanyInfo(id int64, companyInfo orm.Params) (int64, error) {
	beego.Debug("companyInfo", companyInfo, id)
	o := orm.NewOrm()
	return o.QueryTable("company").Filter("Id", id).Update(companyInfo)
}

// 获取公司信息
func GetCompanyInfoById(id int64) (info Company, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("company").Filter("Id", id).Limit(1).One(&info)
	return info, err
}