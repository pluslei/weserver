package models

import (
	//"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"weserver/src/tools"
)

//用户表
type Regist struct {
	Id            int64
	Room          string
	Username      string    `orm:"size(32)" form:"Username"  valid:"Required;MaxSize(32);MinSize(6)"`
	Nickname      string    `orm:"size(255)" form:"Nickname" valid:"Required;MaxSize(255);MinSize(2)"`
	UserIcon      string    `orm:"null;size(255)" form:"UserIcon" valid:"MaxSize(255)"`
	RegStatus     int       `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //用户注册状态 1为未审核 2为审核通过
	Role          *Role     `orm:"rel(one)"`
	Title         *Title    `orm:"rel(one)"`
	IsShutup      bool      //是否禁言
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
}

func init() {
	orm.RegisterModel(new(Regist))
}

func (r *Regist) TableName() string {
	return "regist"
}

//添加用户
func AddRegistUser(r *Regist) (int64, error) {
	o := orm.NewOrm()
	var info []Regist
	num, err := o.QueryTable("regist").Filter("Room", r.Room).Filter("Username", r.Username).All(&info)
	if num == 0 && err == nil {
		id, err := o.Insert(r)
		return id, err
	}
	return 0, err
}

// 获取用户一对一关系
func LoadRegist(r *Regist, fields ...string) (*Regist, error) {
	o := orm.NewOrm()
	err := o.Read(r, fields...)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	_, err = o.LoadRelated(r, "Role")
	_, err = o.LoadRelated(r, "Title")
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return r, nil
}

// 删除指定用户
func DelRegistUame(Room, Uname string) (int64, error) {
	o := orm.NewOrm()
	beego.Debug("Room, Uname", Room, Uname)
	num, err := o.QueryTable("regist").Filter("Room", Room).Filter("Username", Uname).Delete()
	return num, err
}

//更新禁言字段
func UpdateRegistIsShut(room, username string, b bool) (int64, error) {
	o := orm.NewOrm()
	var table Regist
	id, err := o.QueryTable(table).Filter("Room", room).Filter("Username", username).Update(orm.Params{"IsShutup": b})
	return id, err
}

//获取用户权限
func GetRegistPermiss(room, username string) ([]Regist, int64, error) {
	o := orm.NewOrm()
	var info []Regist
	num, err := o.QueryTable("regist").Filter("Room", room).Filter("Username", username).Filter("RegStatus", 2).All(&info)
	return info, num, err
}

//获取Regist表中当天所有禁言人数信息
func GetShutUpInfoToday() (users []Regist, err error) {
	o := orm.NewOrm()
	nowtime := time.Now().Unix() - 24*60*60
	_, err = o.QueryTable("regist").Exclude("Username", "admin").Exclude("Username", "").Filter("IsShutUp", 1).Filter("Lastlogintime__gte", time.Unix(nowtime, 0).Format("2006-01-02 15:04:05")).All(&users)
	return users, err
}

//获取user表中最近当天登录列表信息
func GetLoginInfoToday(roomId string) (users []Regist, err error) {
	o := orm.NewOrm()
	nowtime := time.Now().Unix() - 24*60*60
	_, err = o.QueryTable("regist").Exclude("Username", "admin").Filter("Room", roomId).Filter("Lastlogintime__gte", time.Unix(nowtime, 0).Format("2006-01-02 15:04:05")).All(&users)
	return users, err
}

func GetWechatUser(nDay int64) (users []Regist, err error) {
	o := orm.NewOrm()
	nowtime := time.Now().Unix() - nDay*24*60*60
	_, err = o.QueryTable("regist").Exclude("Username", "admin").Exclude("Username", "").Exclude("Room", "").Filter("Lastlogintime__gte", time.Unix(nowtime, 0).Format("2006-01-02 15:04:05")).All(&users)
	return users, err
}
