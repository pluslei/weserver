package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Company struct {
	Id      int64  // 1 关于我们   2 联系我们
	Content string `orm:"type(text) "`
}

func (n *Company) TableName() string {
	return "company"
}

func init() {
	orm.RegisterModel(new(Company))
}

func (this *Company) UpdateCompanyFields(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func GetCompanyInfo(id int64) (c Company) {
	model := orm.NewOrm()
	model.QueryTable("company").Filter("Id", id).All(&c)
	return
}
