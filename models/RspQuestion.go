package models

import (
	"errors"
	"time"

	"weserver/src/tools"

	"github.com/astaxie/beego/orm"
)

/*
*  RspQuestion table
 */
type RspQuestion struct {
	Id            int64 `orm:"pk;auto"`
	CompanyId     int64
	Room          string //房间号 topic
	Uname         string //用户名  openid
	Nickname      string //用户昵称
	UserIcon      string
	RoleName      string //用户角色[vip,silver,gold,jewel]
	RoleTitle     string //用户角色名[会员,白银会员,黄金会员,钻石会员]
	Sendtype      string //用户发送消息类型('TXT','IMG','VOICE')
	RoleTitleCss  string //头衔颜色
	RoleTitleBack int    `orm:"default(0)"` //角色聊天背景
	Content       string `orm:"type(text)"` //消息内容
	DatatimeStr   string //添加时间
	Time          time.Time
	Uuid          string // uuid

	Question *Question `orm:"rel(fk)"`

	MsgType int `orm:"-"`
}

func init() {
	orm.RegisterModel(new(RspQuestion))
}

func (c *RspQuestion) TableName() string {
	return "rspquestion"
}

func AddRspQuestion(c *RspQuestion) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(c)
	return id, err
}

func GetMoreRspQuestion(id int64) ([]*RspQuestion, int64, error) {
	o := orm.NewOrm()
	var rsp []*RspQuestion
	num, err := o.QueryTable(new(RspQuestion)).Filter("question", id).RelatedSel().All(&rsp)
	return rsp, num, err
}

//事务添加数据
func AddRspQuestiondata(chat []RspQuestion, length int) error {
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
func DelRspQuestionByUUId(uuid string) (int64, error) {
	o := orm.NewOrm()
	var chat RspQuestion
	status, err := o.QueryTable(chat).Filter("Uuid", uuid).Delete()
	return status, err
}

func DeleteRspQuestionById(id int64) (int64, error) {
	o := orm.NewOrm()
	var chat RspQuestion
	status, err := o.QueryTable(chat).Filter("Id", id).Delete()
	return status, err
}

//获取数据库表中的总条数
func GetRspQuestionCount() (int64, error) {
	o := orm.NewOrm()
	var table RspQuestion
	count, err := o.QueryTable(table).Count()
	return count, err
}

func GetLastRspQuestionInfo(count int64) (chat []Question, countline int64) {
	o := orm.NewOrm()
	var table RspQuestion
	countline, _ = o.QueryTable(table).OrderBy("-Id").Limit(count).All(&chat)
	return chat, countline
}

//根据表名和数量 获取信息
func GetRspQuestionInfo(count int64, roomId string) ([]RspQuestion, int64, error) {
	o := orm.NewOrm()
	var chat []RspQuestion
	num, err := o.QueryTable("rspquestion").Filter("Room", roomId).OrderBy("Id").Limit(count).All(&chat)
	return chat, num, err
}

//获取指定的聊天记录
func GetAllRspQuestionMsg(roomId, username string, RoleId int64) ([]RspQuestion, int64, error) {
	o := orm.NewOrm()
	var chat []RspQuestion
	var table RspQuestion
	if RoleId == int64(tools.ROLE_MANAGER) || RoleId == int64(tools.ROLE_TEACHER) || RoleId == int64(tools.ROLE_ASSISTANT) {
		num, err := o.QueryTable(table).Filter("Room", roomId).OrderBy("-Id").All(&chat)
		return chat, num, err
	}
	num, err := o.QueryTable(table).Filter("Room", roomId).Filter("Uname", username).OrderBy("-Id").All(&chat)
	return chat, num, err
}

// 根据id查询聊天内容
func GetRspQuestionIdData(id int64) (chat RspQuestion, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("rspquestion").Filter("Id", id).One(&chat)
	return chat, err
}

// 获取消息列表
func GetRspQuestionRecordList(page int64, page_size int64, sort, Nickname string, companyId int64) (ms []orm.Params, count int64) {
	o := orm.NewOrm()
	chatrecord := new(RspQuestion)
	query := o.QueryTable(chatrecord)
	if companyId != 0 {
		query.Limit(page_size, page).Filter("CompanyId", companyId).Filter("nickname__contains", Nickname).OrderBy(sort).Values(&ms)
		count, _ = query.Count()
		return ms, count
	}
	query.Limit(page_size, page).Filter("nickname__contains", Nickname).OrderBy(sort).Values(&ms)
	count, _ = query.Count()
	return ms, count
}

//回复消息
func ReplyQuestion(id int64, replyMsg string) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE rspquestion SET content = ? WHERE id = ?", replyMsg, id).Exec()
	var resnum int64
	if err != nil {
		num, _ := res.RowsAffected()
		resnum = int64(num)
	}
	return resnum, err
}

func GetRspByQuestionId(id int64) (RspQuestion, error) {
	o := orm.NewOrm()
	var chat RspQuestion
	err := o.Raw("SELECT id,content,nickname,user_icon,role_title FROM rspquestion WHERE question_id = ?", id).QueryRow(&chat)
	return chat, err
}

// 根据id查询聊天内容
func GetRspQuestionById(id int64) (RspQuestion, error) {
	o := orm.NewOrm()
	var chat RspQuestion
	err := o.QueryTable(chat).Filter("Id", id).One(&chat)
	return chat, err
}
