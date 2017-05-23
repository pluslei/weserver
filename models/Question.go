package models

import (
	"errors"
	"time"

	"weserver/src/tools"

	"github.com/astaxie/beego/orm"
)

/*
*  Question table
 */
type Question struct {
	Id            int64 `orm:"pk;auto"`
	CompanyId     int64
	Room          string //房间号 topic
	Uname         string //用户名  openid
	Nickname      string //用户昵称
	UserIcon      string //用户logo
	RoleName      string //用户角色[vip,silver,gold,jewel]
	RoleTitle     string //用户角色名[会员,白银会员,黄金会员,钻石会员]
	Sendtype      string //用户发送消息类型('TXT','IMG','VOICE')
	RoleTitleCss  string //头衔颜色
	RoleTitleBack int    `orm:"default(0)"` //角色聊天背景
	Content       string `orm:"type(text)"` //消息内容
	DatatimeStr   string //添加时间
	IsIgnore      int64  //是否忽略 0 显示 1 忽略
	Uuid          string // uuid
	Time          time.Time
	RspQuestion   []*RspQuestion `orm:"reverse(many)"` //一对多

	//回复信息
	RspNickname string `orm:"-"`
	RspTitle    string `orm:"-"`
	RspIcon     string `orm:"-"`
	RspTimestr  string `orm:"-"`
	RspContent  string `orm:"-"`
}

func init() {
	orm.RegisterModel(new(Question))
}

func (c *Question) TableName() string {
	return "question"
}

func AddQuestion(c *Question) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(c)
	return id, err
}

//事务添加数据
func AddQuestiondata(chat []Question, length int) error {
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
func DelQuestionById(uuid string) (int64, error) {
	o := orm.NewOrm()
	var chat Question
	status, err := o.QueryTable(chat).Filter("Uuid", uuid).Delete()
	return status, err
}

//获取数据库表中的总条数
func GetQuestionCount() (int64, error) {
	o := orm.NewOrm()
	var table Question
	count, err := o.QueryTable(table).Count()
	return count, err
}

// 获取最后一个id
func GetLastQuestionId(count int64) (chat []Question, countline int64) {
	o := orm.NewOrm()
	var table Question
	countline, _ = o.QueryTable(table).OrderBy("-Id").Limit(count).All(&chat)
	return chat, countline
}

//根据表名和数量 获取信息
func GetQuestionMsgData(count int64, roomId string) ([]Question, int64, error) {
	o := orm.NewOrm()
	var chat []Question
	var table Question
	num, err := o.QueryTable(table).Filter("Room", roomId).OrderBy("Id").Limit(count).All(&chat)
	return chat, num, err
}

//获取指定的聊天记录
func GetAllQuestionMsg(roomId, username string, RoleId int64) ([]Question, int64, error) {
	o := orm.NewOrm()
	var chat []Question
	var table Question
	if RoleId == tools.ROLE_MANAGER || RoleId == tools.ROLE_TEACHER || RoleId == tools.ROLE_ASSISTANT {
		num, err := o.QueryTable(table).Filter("Room", roomId).OrderBy("-Id").All(&chat)
		return chat, num, err
	}
	num, err := o.QueryTable(table).Filter("Room", roomId).Filter("Uname", username).OrderBy("-Id").All(&chat)
	return chat, num, err
}

// 根据id查询聊天内容
func GetQuestionIdData(id int64) (Question, error) {
	o := orm.NewOrm()
	var chat Question
	err := o.QueryTable(chat).Filter("Id", id).One(&chat)
	return chat, err
}

// 获取消息列表
func GetQuestionRecordList(page int64, page_size int64, sort, Nickname string, companyId int64) (ms []orm.Params, count int64) {
	o := orm.NewOrm()
	chatrecord := new(Question)
	query := o.QueryTable(chatrecord)
	if companyId != 0 {
		query.Limit(page_size, page).Filter("CompanyId", companyId).Filter("nickname__contains", Nickname).RelatedSel().OrderBy(sort).Values(&ms)
		count, _ = query.Count()
		return ms, count
	}
	query.Limit(page_size, page).Filter("nickname__contains", Nickname).RelatedSel().OrderBy(sort).Values(&ms)
	count, _ = query.Count()
	return ms, count
}

//回复消息
func QuestionReply(id int64, replyMsg string) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE question SET accept_content = ? WHERE id = ?", replyMsg, id).Exec()
	var resnum int64
	if err != nil {
		num, _ := res.RowsAffected()
		resnum = int64(num)
	}
	return resnum, err
}

//删除数据库中表中ID对应的行信息
func DeleteQueById(id int64) (int64, error) {
	o := orm.NewOrm()
	var chat Question
	status, err := o.QueryTable(chat).Filter("Id", id).Delete()
	return status, err
}

func OpeateIgnore(id int64) (int64, error) {
	o := orm.NewOrm()
	var info Question
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"IsIgnore": 1})
	return id, err
}
