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
	UserId        int64     `orm:"index"`
	Nickname      string    `orm:"size(255)" form:"Nickname" valid:"Required;MaxSize(255);MinSize(2)"`
	UserIcon      string    `orm:"null;size(255)" form:"UserIcon" valid:"MaxSize(255)"`
	RegStatus     int       `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //用户注册状态 1为未审核 2为审核通过
	Role          *Role     `orm:"rel(one)"`
	Title         *Title    `orm:"rel(one)"`
	IsShutup      bool      //是否禁言
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `

	RoomName string `orm:"-"` //房间名字
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

//跟新登录时间
func UpdateLoginTime(room, username string) (int64, error) {
	o := orm.NewOrm()
	time := time.Now()
	var table Regist
	id, err := o.QueryTable(table).Filter("Room", room).Filter("Username", username).Update(orm.Params{"Lastlogintime": time})
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

// //最近 X 天 人员总列表信息
// func GetAllRegistList(RoomId string, nDay int64) (users []Regist, count int) {
// 	o := orm.NewOrm()
// 	nowtime := time.Now().Unix() - nDay*24*60*60
// 	_, err = o.QueryTable("regist").Exclude("Username", "admin").Filter("Lastlogintime__gte", time.Unix(nowtime, 0).Format("2006-01-02 15:04:05")).Limit(-1).All(&users)
// 	return users, err
// }

//get user list
func GetWechatUserList(page int64, page_size int64, sort, nickname string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(Regist)
	qs := o.QueryTable(user).Exclude("Username", "admin")
	qs.Limit(page_size, page).Filter("nickname__contains", nickname).OrderBy(sort).RelatedSel().Values(&users)
	count, _ = qs.Count()
	return users, count
}

// 更新用户进入房间状态
func UpdateWechtUserStatus(id int64, status int) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("regist").Filter("Id", id).Update(orm.Params{
		"RegStatus": status,
	})
}

// 更新用户头衔
func UpdateWechatUserTitle(id int64, titleid int64) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("regist").Filter("Id", id).Update(orm.Params{
		"title_id": titleid,
	})
}

// 获取用户信息
func GetWechatUserInfoById(id int64) (user Regist, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("regist").Filter("Id", id).One(&user)
	return user, err
}

// 更新用户
func UpdateWechatUserInfo(id, roleId, titleId int64, regstatus int) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("regist").Filter("Id", id).Update(orm.Params{
		"title_id":  titleId,
		"role_id":   roleId,
		"RegStatus": regstatus,
	})
}

// 更新指定账户的username
func UpdateRegistName(userid int64, username, icon string) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("regist").Filter("UserId", userid).Update(orm.Params{
		"Username": username,
		"UserIcon": icon,
	})
}
