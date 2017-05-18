package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
* Company table
 */
type Company struct {
	Id            int64  `orm:"pk;auto"`
	Company       string //公司名称
	CompanyIntro  string //公司简介
	CompanyIcon   string //公司icon图
	CompanyBanner string //公司banner图
	HistoryMsg    int64  //是否显示历史消息 0显示  1 不显示
	Registerrole  int64  //默认注册用户角色
	WelcomeMsg    string //欢迎语
	AuditMsg      int64  //是否开启消息审核  0开启 1关闭
	Verify        int64  //是否开启用户审核  0开启 1不开启
	AppId         string //appid
	AppSecret     string //密钥
	Url           string //跳转url

	Rolename string `orm:"-"` //头衔名称
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
func GetCompanys(page int64, page_size int64, companyId int64) (ms []orm.Params, count int64) {
	o := orm.NewOrm()
	query := o.QueryTable("company")
	if companyId != 0 {
		query.Limit(page_size, page).Filter("Id", companyId).Values(&ms)
		count, _ = query.Count()
		return ms, count
	}
	query.Limit(page_size, page).OrderBy("id").Values(&ms)
	count, _ = query.Count()
	return ms, count
}

// 更新公司信息
func UpdateCompanyInfo(id int64, companyInfo Company, companyId int64) (int64, error) {
	beego.Debug("companyInfo", companyInfo, id)
	o := orm.NewOrm()
	if companyId != 0 {
		return o.QueryTable("company").Filter("Id", id).Update(orm.Params{
			"Company":       companyInfo.Company,
			"CompanyIntro":  companyInfo.CompanyIntro,
			"CompanyIcon":   companyInfo.CompanyIcon,
			"CompanyBanner": companyInfo.CompanyBanner,
			"HistoryMsg":    companyInfo.HistoryMsg,
			"Registerrole":  companyInfo.Registerrole,
			"WelcomeMsg":    companyInfo.WelcomeMsg,
			"AuditMsg":      companyInfo.AuditMsg,
			"Verify":        companyInfo.Verify,
		})
	}
	return o.QueryTable("company").Filter("Id", id).Update(orm.Params{
		"Company":       companyInfo.Company,
		"CompanyIntro":  companyInfo.CompanyIntro,
		"CompanyIcon":   companyInfo.CompanyIcon,
		"CompanyBanner": companyInfo.CompanyBanner,
		"HistoryMsg":    companyInfo.HistoryMsg,
		"Registerrole":  companyInfo.Registerrole,
		"WelcomeMsg":    companyInfo.WelcomeMsg,
		"AuditMsg":      companyInfo.AuditMsg,
		"Verify":        companyInfo.Verify,
	    "AppId":         companyInfo.AppId,
	    "AppSecret":     companyInfo.AppSecret,
	    "Url":           companyInfo.Url,
	})
}

// 获取公司信息
func GetCompanyInfoById(id int64) (info Company, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("company").Filter("Id", id).Limit(1).One(&info)
	return info, err
}
