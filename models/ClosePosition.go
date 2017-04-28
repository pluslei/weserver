package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*
*  	平仓操作
 */
type ClosePosition struct {
	Id          int64  `orm:"pk;auto"`
	RoomId      string //topic
	RoomTeacher string //老师
	Time        time.Time
	Type        string //种类
	BuySell     int    //买卖 0 1
	Entrust     string //委托类型
	Index       string //点位
	Position    string //仓位
	ProfitPoint string //止盈点
	LossPoint   string //止损点
	Notes       string // 备注
	Timestr     string //时间字符

	OperPosition *OperPosition `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(ClosePosition))
}

func (c *ClosePosition) TableName() string {
	return "closeposition"
}

/*
* 新增加平仓操作
 */
func AddClosePosition(c *ClosePosition) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(c)
	return id, err
}

// 根据建仓id 获取多个平仓操作
func GetMoreClosePosition(id int64) ([]*ClosePosition, int64, error) {
	o := orm.NewOrm()
	var close []*ClosePosition
	num, err := o.QueryTable(new(ClosePosition)).Filter("OperPosition", id).RelatedSel().All(&close)
	return close, num, err
}

// 根据id获取信息
func GetClosePositionInfo(id int64) (close ClosePosition, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(ClosePosition)).Filter("Id", id).One(&close)
	return close, err
}

// 根据operpositon删除
func DelClosePositionByOperId(id int64) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable(new(ClosePosition)).Filter("OperPosition", id).Delete()
}

// 根据id删除
func DelClosePositionById(id int64) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable(new(ClosePosition)).Filter("Id", id).Delete()
}

// 更新
func UpdateClosePosition(id int64, close map[string]interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable(new(ClosePosition)).Filter("Id", id).Update(close)
}

//更新
func UpdateClosePositionInfo(t *ClosePosition) (int64, error) {
	o := orm.NewOrm()
	id, err := o.QueryTable("teacher").Filter("Id", t.Id).Update(orm.Params{
		"RoomTeacher": t.RoomTeacher,
		"Type":        t.Type,
		"BuySell":     t.BuySell,
		"Entrust":     t.Entrust,
		"Index":       t.Index,
		"Position":    t.Position,
		"ProfitPoint": t.Position,
		"LossPoint":   t.LossPoint,
		"Notes":       t.Notes,
		"Time":        t.Time,
		"Timestr":     t.Timestr,
	})
	return id, err
}
