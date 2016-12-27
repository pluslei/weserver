package models

import (
	// "github.com/astaxie/beego"
	// "errors"
	"github.com/astaxie/beego/orm"
)

type TelBanner struct {
	Id     int64
	Detail string `orm:"type(text) "` //描述
	Url    string `orm:"type(text) "` //图片路径
}

func (n *TelBanner) TableName() string {
	return "telbanner"
}

func init() {
	orm.RegisterModel(new(TelBanner))
}

func (this *TelBanner) UpdateTelBanner(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func GetTelBannerInfo(id int64) (c TelBanner) {
	model := orm.NewOrm()
	model.QueryTable("telbanner").Filter("Id", id).All(&c)
	return
}
