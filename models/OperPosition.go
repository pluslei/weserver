package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*
*  	建仓操作
 */
type OperPosition struct {
	Id            int64  `orm:"pk;auto"`
	RoomId        string //topic
	RoomTeacher   string //老师
	Time          time.Time
	Type          string           //种类
	BuySell       int64            //买卖 0 1
	Entrust       string           //委托类型
	Index         string           //点位
	Position      string           //仓位
	ProfitPoint   string           //止盈点
	LossPoint     string           //止损点
	Notes         string           // 备注
	Liquidation   int              //平仓详情 (0:未平仓 1:平仓)
	ClosePosition []*ClosePosition `orm:"reverse(many)"` //一对多
}

func init() {
	orm.RegisterModel(new(OperPosition))
}

func (o *OperPosition) TableName() string {
	return "operposition"
}

/*
* 建仓操作
 */
func AddPosition(o *OperPosition) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(o)
	return id, err
}

//删除建仓信息
func DelPositionById(id int64) (int64, error) {
	o := orm.NewOrm()
	var info OperPosition
	status, err := o.QueryTable(info).Filter("Id", id).Delete()
	return status, err
}

//更新
func UpdatePositionInfo(t *OperPosition) (int64, error) {
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
		"Liquidation": t.Liquidation,
		"Time":        t.Time,
	})
	return id, err
}

// 分页
func GetOperPositionList(page int64, page_size int64, sort string) (ms []orm.Params, count int64) {
	o := orm.NewOrm()
	poer := new(OperPosition)
	query := o.QueryTable(poer)
	query.Limit(page_size, page).OrderBy(sort).Values(&ms)
	count, _ = query.Count()
	return ms, count
}

//获取最近的一条记录
func GetNearRecord(Room string) (OperPosition, error) {
	o := orm.NewOrm()
	var oper OperPosition
	err := o.QueryTable("operposition").Filter("Room", Room).OrderBy("-Id").Limit(1).One(&oper)
	return oper, err
}

//获取所有记录
func GetAllPositionList(Room string) ([]OperPosition, int64, error) {
	o := orm.NewOrm()
	var info []OperPosition
	num, err := o.QueryTable("operposition").Filter("Room", Room).OrderBy("-Id").All(&info)
	return info, num, err
}
