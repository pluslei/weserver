package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

/*
*  消息记录表
 */
type ChatRecord struct {
	Id            int64     `orm:"pk;auto"`
	Room          string    //房间号 topic
	Uname         string    //用户名  openid
	Nickname      string    //用户昵称
	UserIcon      string    //用户logo
	RoleName      string    //用户角色[vip,silver,gold,jewel]
	RoleTitle     string    //用户角色名[会员,白银会员,黄金会员,钻石会员]
	Sendtype      string    //用户发送消息类型('TXT','IMG','VOICE')
	RoleTitleCss  string    //头衔颜色
	RoleTitleBack int       `orm:"default(0)"`                                           //角色聊天背景
	Insider       int       `orm:"default(1)"`                                           //1内部人员或0外部人员
	IsLogin       int       `orm:"default(0)"`                                           //状态 [1、登录 0、未登录]
	Content       string    `orm:"type(text)"`                                           //消息内容
	Datatime      time.Time `orm:"type(datetime)"`                                       //添加时间
	Status        int       `orm:"default(0)" form:"Status" valid:"Required;Range(0,1)"` //消息审核[1 通过 0 未通过]
	Uuid          string    // uuid

	DatatimeStr string `orm:"-"`
	MsgType     int    `orm:"-"`
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
func DelChatById(uuid string) (int64, error) {
	o := orm.NewOrm()
	var chat ChatRecord
	status, err := o.QueryTable(chat).Filter("Uuid", uuid).Delete()
	return status, err
}

//获取数据库表中的总条数
func GetChatCount() (int64, error) {
	o := orm.NewOrm()
	var table ChatRecord
	count, err := o.QueryTable(table).Count()
	return count, err
}

// 获取最后一个id
func GetLastChatId(count int64) (chat []ChatRecord, countline int64) {
	o := orm.NewOrm()
	var table ChatRecord
	countline, _ = o.QueryTable(table).Filter("Status", 1).OrderBy("-Id").Limit(count).All(&chat)
	return chat, countline
}

//根据表名和数量 获取信息
func GetChatMsgData(count int64, roomId, tableName string) ([]ChatRecord, int64, error) {
	o := orm.NewOrm()

	// var satrt int64
	// lastdata, fromcount := GetLastChatId(count)
	// if fromcount == 0 {
	// 	satrt = 0
	// } else {
	// 	satrt = lastdata[fromcount-1].Id
	// }
	//num, err := o.QueryTable(tableName).OrderBy("Id").Filter("Status", 1).Filter("Room", roomId).Filter("Id__gt", satrt).All(&chat)
	var chat []ChatRecord
	num, err := o.QueryTable(tableName).Filter("Status", 1).Filter("Room", roomId).OrderBy("Id").Limit(count).All(&chat)
	return chat, num, err
}

//获取聊天记录
func GetAllChatMsgData(roomId, tableName string) ([]ChatRecord, int64, error) {
	o := orm.NewOrm()
	var chat []ChatRecord
	num, err := o.QueryTable(tableName).Filter("Status", 1).Filter("Room", roomId).OrderBy("-Id").All(&chat)
	return chat, num, err
}

// 根据id查询聊天内容
func GetChatIdData(id int64) (ChatRecord, error) {
	o := orm.NewOrm()
	var chat ChatRecord
	err := o.QueryTable(chat).Filter("Id", id).One(&chat)
	return chat, err
}

//更新内容
func UpdateChatStatus(id int64) (int64, error) {
	o := orm.NewOrm()
	var table ChatRecord
	id, err := o.QueryTable(table).Filter("Id", id).Update(orm.Params{"Status": 1})
	return id, err
}

// 获取消息列表
func GetChatRecordList(page int64, page_size int64, sort string) (ms []orm.Params, count int64) {
	o := orm.NewOrm()
	chatrecord := new(ChatRecord)
	query := o.QueryTable(chatrecord)
	query.Limit(page_size, page).OrderBy(sort).Values(&ms)
	count, _ = query.Count()
	return ms, count
}
