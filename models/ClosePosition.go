package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*
*  	平仓操作
 */
type ClosePosition struct {
	Id           int64  `orm:"pk;auto"`
	RoomId       string //topic
	RoomTeacher  string //老师
	Time         time.Time
	Type         string        //种类
	BuySell      string        //买卖
	Entrust      string        //委托类型
	Index        string        //点位
	Position     string        //仓位
	ProfitPoint  string        //止盈点
	LossPoint    string        //止损点
	Notes        string        // 备注
	Liquidation  string        //平仓详情
	OperPosition *OperPosition `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(ClosePosition))
}

func (c *ClosePosition) TableName() string {
	return "ClosePosition"
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
	num, err := o.QueryTable("ClosePosition").Filter("OperPosition", id).RelatedSel().All(&close)
	return close, num, err
}
