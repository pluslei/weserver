package models

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "strconv"
	"time"
)

type ValidataCode struct {
	Id      int64
	Phone   int64  //手机号码
	Code    string //验证码
	Times   int64  //时间戳
	Timeday string //日期
	Ip      string //客户端IP
	Issure  int    `orm:"default(0)"` //状态 [1、校验成功 0、未校验]
}

func (v *ValidataCode) TableName() string {
	return "validata_code"
}

func init() {
	orm.RegisterModel(new(ValidataCode))
}

// 插入
func InsertCode(v *ValidataCode) (int64, error) {
	model := orm.NewOrm()
	insertid, err := model.Insert(v)
	return insertid, err
}

// 查询电话同一天是否三次
func CheckPhoneDay(phone int64) (int64, error) {
	model := orm.NewOrm()
	timeday := time.Now().Format("2006-01-02")
	count, err := model.QueryTable("validata_code").Filter("Phone", phone).Filter("Timeday", timeday).Count()
	return count, err
}

// 查询IP同一天是否三个电话
func CheckIpDay(ip string, phone string) int {
	model := orm.NewOrm()
	var varlidatacount []orm.Params
	timeday := time.Now().Format("2006-01-02")
	sql := "SELECT * FROM validata_code WHERE timeday = ? AND ip = ? AND phone !=? GROUP BY phone"
	model.Raw(sql, timeday, ip, phone).Values(&varlidatacount)
	if varlidatacount == nil {
		return 0
	} else {
		countip := len(varlidatacount)
		return countip
	}
}

// 检验验证码是是否正确
func CheckCode(phone int64, code string) (codevalidata ValidataCode, err error) {
	model := orm.NewOrm()
	timeday := time.Now().Format("2006-01-02")
	err = model.QueryTable("validata_code").Filter("Phone", phone).Filter("Timeday", timeday).Filter("Code", code).One(&codevalidata)
	return codevalidata, err
}

// 更新验证码验证状态
func UpdateValidataCode(id int64) (int64, error) {
	model := orm.NewOrm()
	num, err := model.QueryTable("validata_code").Filter("Id", id).Update(orm.Params{"issure": 1})
	return num, err
}
