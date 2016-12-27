package models

import (
	//"errors"
	"github.com/astaxie/beego/orm"
)

//分组表
type Custservice struct {
	Id         int64
	CustName   string `orm:"size(128)" form:"CustName"  valid:"Required"`
	CustNumber string `orm:"size(255)" form:"CustNumber"  valid:"Required"`
	Status     int    `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //状态(1、启用，2、停用)
	Order      int    `orm:"default(1)" form:"Order"`
}

func (g *Custservice) TableName() string {
	return "custservice"
}

func init() {
	orm.RegisterModel(new(Custservice))
}

//get group list
func GetCustservicelist(page int64, page_size int64, sort string) (custs []orm.Params, count int64) {
	o := orm.NewOrm()
	cust := new(Custservice)
	qs := o.QueryTable(cust)
	qs.Limit(page_size, page).OrderBy(sort).Values(&custs)
	count, _ = qs.Count()
	return custs, count
}

func AddCustservice(c *Custservice) (int64, error) {
	o := orm.NewOrm()
	cust := new(Custservice)
	cust.CustName = c.CustName
	cust.CustNumber = c.CustNumber
	cust.Status = c.Status
	cust.Order = c.Order
	id, err := o.Insert(cust)
	return id, err
}

func (this *Custservice) UpdateCustservice(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func DelCustserviceById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Custservice{Id: Id})
	return status, err
}

func CustserviceList() (custs []orm.Params) {
	o := orm.NewOrm()
	cust := new(Custservice)
	qs := o.QueryTable(cust)
	qs.Values(&custs)
	return custs
}

func ReadCustserviceById(cid int64) (Custservice, error) {
	o := orm.NewOrm()
	cust := Custservice{Id: cid}
	err := o.Read(&cust)
	if err != nil {
		return cust, err
	}
	return cust, nil
}

func CustStatusCount() (int64, error) {
	o := orm.NewOrm()
	var table Custservice
	num, err := o.QueryTable(table).Count()
	if err != nil {
		return num, err
	}
	return num, nil
}

func IndexCustList() (custs []orm.Params) {
	o := orm.NewOrm()
	cust := new(Custservice)
	query := o.QueryTable(cust)
	query.OrderBy("-Order").Filter("Status", 1).Values(&custs)
	return custs
}
