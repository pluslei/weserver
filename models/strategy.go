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
	Id        int64  `orm:"pk;auto"`
	Room      string //房间号 topic
	Icon      string //头像
	Name      string `orm:"size(128)" form:"Uname" valid:"Required"` //操作者的用户名
	Titel     string
	Data      string    `orm:"type(text)"` //策略内容
	FileName  string    //图片
	TxtColour string    //文字颜色
	IsTop     bool      //是否置顶 置顶1 否 0
	IsDelete  bool      //是否删除,删除 1 否 0
	ThumbNum  int64     //点赞次数
	Time      string    //前台给的时间
	Datatime  time.Time `orm:"type(datetime)"` //添加时间

	DatatimeStr string `orm:"-"`
}

func init() {
	orm.RegisterModel(new(Strategy))
}

func (b *Strategy) TableName() string {
	return "strategy"
}

// 获取指定房间的最低条数
func GetStrategyCount(room string, count int64) ([]Strategy, int64, error) {
	o := orm.NewOrm()
	var info []Strategy
	num, err := o.QueryTable("Strategy").Filter("Room", room).OrderBy("-Id").OrderBy("IsTop").Limit(count).All(&info)
	var infoSort []Strategy
	for i := 0; i < len(info); i++ {
		infoSort = append(infoSort, info[len(info)-1-i])
	}
	return infoSort, num, err
}

// 获取指定房间的策略列表
func GetStrategyList(room string) ([]Strategy, int64, error) {
	o := orm.NewOrm()
	var info []Strategy
	num, err := o.QueryTable("Strategy").Filter("Room", room).OrderBy("-Id").OrderBy("IsTop").All(&info)
	var infoSort []Strategy
	for i := 0; i < len(info); i++ {
		infoSort = append(infoSort, info[len(info)-1-i])
	}
	return infoSort, num, err
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
func ThumbOptionAdd(id int64) (int64, error) {
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

//取消点赞
func ThumbOptionDel(id int64) (int64, error) {
	num, err := GetThumbNum(id)
	if err != nil {
		beego.Debug("Get ThumbNum fail", err)
		return 0, nil
	}
	o := orm.NewOrm()
	var info Strategy
	status, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"ThumbNum": num - 1})
	return status, err
}

//获取原来点赞次数
func GetThumbNum(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Strategy
	err := o.QueryTable(info).Filter("id", id).One(&info, "ThumbNum")
	return info.ThumbNum, err
}

// 策略分页
func GetStrategyInfoList(page int64, page_size int64, sort string) (ms []orm.Params, count int64) {
	o := orm.NewOrm()
	strategy := new(Strategy)
	query := o.QueryTable(strategy)
	query.Limit(page_size, page).OrderBy(sort).Values(&ms)
	count, _ = query.Count()
	return ms, count
}

// 获取单个信息
func GetStrategyInfoById(id int64) (info Strategy, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(info).Filter("id", id).One(&info)
	return info, err
}

// 更新
func UpdateStrategy(id int64, strate map[string]interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable(new(Strategy)).Filter("Id", id).Update(strate)
}
