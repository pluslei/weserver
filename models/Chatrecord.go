package models

import (
	//"github.com/astaxie/beego"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
*用户表
* beego 中会把名为Id的字段自动设置文自增加的主键
 */
type ChatRecord struct {
	Id          int64     `orm:"pk;auto"`
	Coderoom    int       //房间号
	Uname       string    `orm:"size(128)" form:"Uname" valid:"Required"`                   //聊天的用户名
	Objname     string    `orm:"size(128)" form:"Objname"  valid:"Required"`                //被聊天的用户名
	Sendtype    int       `orm:"default(0)" form:"Sendtype" valid:"Required;Range(0,1)"`    //聊天状态(0：公聊，1：私聊)
	AuditStatus int       `orm:"default(0)" form:"AuditStatus" valid:"Required;Range(0,2)"` //信息是否审核,0：不用审核，1：审核通过，2：未审核(游客，会员需要审核)
	Data        string    `orm:"type(text)"`                                                //消息内容
	Ipaddress   string    //IP地址
	Procities   string    //省市
	Datatime    time.Time `orm:"type(datetime)"` //添加时间
	Timestmp    int64     //时间戳
}

func init() {
	orm.RegisterModel(new(ChatRecord))
}

func (c *ChatRecord) TableName() string {
	return "chat_record"
}

/*
*添加聊天消息
 */
func AddChat(c *ChatRecord) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(c)
	return id, err
}

//事务添加数据
func AddChatdata(chat []ChatRecord, length int) error {
	model := orm.NewOrm()
	err := model.Begin()
	SuccessNum := 0
	if err == nil {
		for i := 0; i < length; i++ {
			id, err := model.Insert(&chat[i])
			if err == nil && id > 0 {
				SuccessNum++
			}
		}
	} else {
		err = errors.New("事务申请失败!")
	}
	if SuccessNum == length {
		err = model.Commit()
	} else {
		err = errors.New("事务提交失败!")
	}
	return err
}

//删除数据库中表中ID对应的行信息
func DelChatById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&ChatRecord{Id: Id})
	return status, err
}

//获取数据库表中的总条数
func GetChatCount() (int64, error) {
	o := orm.NewOrm()
	var table ChatRecord
	count, err := o.QueryTable(table).Count()
	return count, err
}

//获取内容
func GetChatData(uname string, datanum int) ([]ChatRecord, int64, error) {
	o := orm.NewOrm()
	var chat []ChatRecord
	var table ChatRecord
	num, err := o.QueryTable(table).Filter("Uname", uname).OrderBy("-Id").Limit(datanum).All(&chat)
	return chat, num, err
}

//获取内容
func UpdateChatData(chat ChatRecord) (int64, error) {
	o := orm.NewOrm()
	var table ChatRecord
	id, err := o.QueryTable(table).Filter("Id", chat.Id).Update(orm.Params{"Data": chat.Data, "Coderoom": chat.Coderoom, "Datatime": chat.Datatime})
	return id, err
}
