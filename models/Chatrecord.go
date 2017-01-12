package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
*用户表
* beego 中会把名为Id的字段自动设置文自增加的主键
 */
type ChatRecord struct {
	Id            int64     `orm:"pk;auto"`
	Code          int       //公司代码
	Room          int       //房间号
	Uname         string    //用户名
	Nickname      string    //用户昵称
	UserIcon      string    //用户logo
	RoleName      string    //用户角色[vip,silver,gold,jewel]
	RoleTitle     string    //用户角色名[会员,白银会员,黄金会员,钻石会员]
	Sendtype      string    //用户发送消息类型('TXT','IMG','VOICE')
	RoleTitleCss  string    //头衔颜色
	RoleTitleBack int       `orm:"default(0)"`     //角色聊天背景
	Insider       int       `orm:"default(1)"`     //1内部人员或0外部人员
	IsLogin       int       `orm:"default(0)"`     //状态 [1、登录 0、未登录]
	Content       string    `orm:"type(text)"`     //消息内容
	Datatime      time.Time `orm:"type(datetime)"` //添加时间
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
func GetChatMsgData(count int64) ([]ChatRecord, int64, error) {
	o := orm.NewOrm()
	var chat []ChatRecord
	chatcount, _ := GetChatCount()
	startpos := chatcount - count
	if startpos < 0 {
		startpos = 0
	}
	num, err := o.QueryTable("ChatRecord").OrderBy("Id").Limit(chatcount, startpos).All(&chat)
	return chat, num, err
}

//获取内容
func UpdateChatData(chat ChatRecord) (int64, error) {
	o := orm.NewOrm()
	var table ChatRecord
	id, err := o.QueryTable(table).Filter("Id", chat.Id).Update(orm.Params{"Content": chat.Content, "Room": chat.Room, "Datatime": chat.Datatime})
	return id, err
}
