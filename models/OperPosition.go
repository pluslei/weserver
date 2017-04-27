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
	BuySell       int              //买卖 0 1
	Entrust       string           //委托类型
	Index         string           //点位
	Position      string           //仓位
	ProfitPoint   string           //止盈点
	LossPoint     string           //止损点
	Notes         string           // 备注
	Liquidation   int              //平仓详情 (0:未平仓 1:平仓)
	ClosePosition []*ClosePosition `orm:"reverse(many)"` //一对多

	Timestr string `orm:"-"` //时间字符
}

func init() {
	orm.RegisterModel(new(OperPosition))
}

func (o *OperPosition) TableName() string {
	return "operposition"
}

/*
* 新增加建仓操作
 */
func AddPosition(o *OperPosition) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(o)
	return id, err
}

// 修改平仓详情
func UpdatePositonLq(id int64) {
	o := orm.NewOrm()
	o.QueryTable(new(OperPosition)).Filter("Id", id).Update(orm.Params{
		"Liquidation": 1,
	})
}

// 根据id查询
func GetOpersitionInfoById(id int64) (oper OperPosition, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(OperPosition)).Filter("Id", id).One(&oper)
	return oper, err
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
