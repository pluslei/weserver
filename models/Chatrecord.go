package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
*  消息记录表
 */
type ChatRecord struct {
	Id            int64 `orm:"pk;auto"`
	CompanyId     int64
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

	AcceptUuid    string
	AcceptTitle   string
	AcceptContent string

	CompanyName string `orm:"-"`
	RoomName    string `orm:"-"`
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

//获取指定的聊天记录
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
func GetChatRecordList(page int64, page_size int64, sort, Nickname string, companyId int64, SearchId, RoomId string) (ms []orm.Params, count int64) {
	var sId int64
	var err error
	if SearchId != "" {
		sId, err = strconv.ParseInt(SearchId, 10, 10)
		if err != nil {
			beego.Debug("get Search 0 Fail", err)
			return
		}
	}
	o := orm.NewOrm()
	obj := new(ChatRecord)
	if SearchId != "" && RoomId != "" {
		qs := o.QueryTable(obj)
		qs.Limit(page_size, page).Filter("CompanyId", sId).Filter("Room", RoomId).OrderBy("-Id").Values(&ms)
		count, _ = qs.Count()
		return ms, count
	}
	if SearchId != "" && RoomId == "" {
		qs := o.QueryTable(obj)
		qs.Limit(page_size, page).Filter("CompanyId", sId).OrderBy("-Id").Values(&ms)
		count, _ = qs.Count()
		return ms, count
	}
	if companyId != 0 {
		if SearchId == "" && RoomId == "" {
			qs := o.QueryTable(obj)
			qs.Limit(page_size, page).Filter("CompanyId", companyId).OrderBy("-Id").Filter("nickname__contains", Nickname).Values(&ms)
			count, _ = qs.Count()
			return ms, count
		}
	}
	query := o.QueryTable(obj)
	query.Limit(page_size, page).Filter("nickname__contains", Nickname).OrderBy(sort).Values(&ms)
	count, _ = query.Count()
	return ms, count
}

//批量删除聊天记录
func PrepareDelChatrecords(IdArray []int64) (int64, error) {
	o := orm.NewOrm()
	err := o.Begin()
	var status int64
	for i := 0; i < len(IdArray); i++ {
		status, err = o.Delete(&ChatRecord{Id: IdArray[i]})
		beego.Info()
	}
	// 此过程中的所有使用 o Ormer 对象的查询都在事务处理范围内
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return status, err
}
