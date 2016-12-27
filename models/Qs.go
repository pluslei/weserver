package models

import (
	//"errors"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//分组表
type Qs struct {
	Id       int64
	Weight   int    //权重
	Question string //问题
	Answer   string `orm:"type(text)"` //解答
}

func (g *Qs) TableName() string {
	return "qs"
}

func init() {
	orm.RegisterModel(new(Qs))
}

//get title list
func GetQslist(page int64, page_size int64, sort string) (q []orm.Params, count int64) {
	o := orm.NewOrm()
	qs := new(Qs)
	query := o.QueryTable(qs)
	query.Limit(page_size, page).OrderBy(sort).Values(&q)
	count, _ = query.Count()
	return q, count
}

//get title list
func GetQsListByWeight() (q []orm.Params) {
	o := orm.NewOrm()
	qs := new(Qs)
	query := o.QueryTable(qs)
	query.OrderBy("-Weight").Values(&q)
	return q
}

func AddQs(t *Qs) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(t)
	return id, err
}

func (this *Qs) UpdateQs(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func DelQsById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Qs{Id: Id})
	return status, err
}

func GetQsList() (q []orm.Params) {
	o := orm.NewOrm()
	qs := new(Qs)
	query := o.QueryTable(qs)
	query.Values(&q)
	return q
}

func ReadQsById(Id int64) (Qs, error) {
	o := orm.NewOrm()
	qs := Qs{Id: Id}
	err := o.Read(&qs)
	if err != nil {
		return qs, err
	}
	return qs, nil
}
