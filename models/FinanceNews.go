package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//分组表
type FinanceNews struct {
	Id        int64  `orm:"pk;auto"`
	Pulltime  string //显示时间
	Contents  string `orm:"type(text)"` //内容
	Cmd5      string //内容md5
	Timestamp int64  //时间戳标识
	Style     int    //样式 [1 加红色]
}

func (s *FinanceNews) TableName() string {
	return "finance_news"
}

func init() {
	orm.RegisterModel(new(FinanceNews))
}

// 插入数据
func AddFinanceNews(c *FinanceNews) (int64, error) {
	o := orm.NewOrm()
	finance := new(FinanceNews)
	finance.Pulltime = c.Pulltime
	finance.Contents = c.Contents
	finance.Cmd5 = c.Cmd5
	finance.Timestamp = time.Now().Unix()
	finance.Style = c.Style
	id, err := o.Insert(finance)
	return id, err
}

// 删除旧数据
func DelFinanceNewsData() (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("finance_news").Filter("timestamp__lte", time.Now().Unix()-24*60*60).Delete()
	return num, err
}

//get title list
func GetFinanceNewslist(page int64, page_size int64, sort string) (configs []orm.Params, count int64) {
	o := orm.NewOrm()
	config := new(FinanceNews)
	qs := o.QueryTable(config)
	qs.Limit(1).OrderBy(sort).Values(&configs)
	count, _ = qs.Count()
	return configs, count
}

func GetFinanceNewsCount() int64 {
	o := orm.NewOrm()
	finance := new(FinanceNews)
	sys := o.QueryTable(finance)
	count, _ := sys.Count()
	return count
}

func (this *FinanceNews) UpdateFinanceNews01(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func ReadFinanceNewsById(id int64) (FinanceNews, error) {
	o := orm.NewOrm()
	finance := FinanceNews{Id: id}
	err := o.Read(&finance)
	if err != nil {
		return finance, err
	}
	return finance, nil
}

func GetFinanceNews() (finance []orm.Params, err error) {
	model := orm.NewOrm()
	_, err = model.QueryTable("finance_news").OrderBy("-Id").Limit(20).Values(&finance)
	return finance, err
}

func GetFinanceNewsOne() (finance FinanceNews, err error) {
	model := orm.NewOrm()
	err = model.QueryTable("finance_news").OrderBy("-Id").Limit(1).One(&finance)
	return finance, err
}
