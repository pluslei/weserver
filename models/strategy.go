package models

import (
	//"github.com/astaxie/beego"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
* 策略表
 */
type Strategy struct {
	Id       int64  `orm:"pk;auto"`
	Room     string //房间号 topic
	Icon     string //头像
	Name     string `orm:"size(128)" form:"Uname" valid:"Required"` //操作者的用户名
	Titel    string
	Data     string    `orm:"type(text)"` //策略内容
	IsTop    bool      //是否置顶 置顶1 否 0
	IsDelete bool      //是否删除,删除 1 否 0
	ThumbNum int64     //点赞次数
	Time     string    //前台给的时间
	Datatime time.Time `orm:"type(datetime)"` //添加时间
}

func init() {
	orm.RegisterModel(new(Strategy))
}

func (b *Strategy) TableName() string {
	return "strategy"
}

// 获取指定房间的策略列表
func GetStrategyList(room string, count int64) ([]Strategy, int64, error) {
	o := orm.NewOrm()
	var info []Strategy
	num, err := o.QueryTable("Strategy").Filter("Room", room).OrderBy("-Id").Limit(count).All(&info)
	return info, num, err
}

/*
*增加策略
 */
func AddStrategy(s *Strategy) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(s)
	return id, err
}

//删除策略
func DelStrategyById(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Strategy
	status, err := o.QueryTable(info).Filter("Id", id).Delete()
	return status, err
}

//置顶操作
func StickOption(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Strategy
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"IsTop": 1})
	return id, err
}

//取消置顶
func UnStickOption(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Strategy
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"IsTop": 0})
	return id, err
}

//点赞操作
func ThumbOption(id int64) (int64, error) {
	num, err := GetThumbNum(id)
	if err != nil {
		beego.Debug("Get ThumbNum fail", err)
		return 0, nil
	}
	o := orm.NewOrm()
	var info Strategy
	status, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"ThumbNum": num + 1})
	return status, err
}

//获取原来点赞次数
func GetThumbNum(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Strategy
	err := o.QueryTable(info).OrderBy("id").One(&info, "ThumbNum")
	return info.ThumbNum, err
}
