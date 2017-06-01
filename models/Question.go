package models

import (
	"errors"
	"strconv"
	"time"

	"weserver/src/tools"

	"github.com/astaxie/beego"
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
	IsIgnore      int64  //是否忽略 0 忽略 1 显示
	Uuid          string // uuid
	Time          time.Time
	RspQuestion   []*RspQuestion `orm:"reverse(many)"` //一对多

	AcceptNickname string
	AcceptTitle    string
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
	if RoleId == int64(tools.ROLE_MANAGER) || RoleId == int64(tools.ROLE_TEACHER) || RoleId == int64(tools.ROLE_ASSISTANT) {
		num, err := o.QueryTable(table).Filter("Room", roomId).Filter("IsIgnore", 1).OrderBy("-Id").All(&chat)
		return chat, num, err
	}
	num, err := o.QueryTable(table).Filter("Room", roomId).Filter("Uname", username).Filter("IsIgnore", 1).OrderBy("-Id").All(&chat)
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
func GetQuestionRecordList(page int64, page_size int64, sort, Nickname string, companyId int64, SearchId, RoomId string) (ms []orm.Params, count int64) {
	var sId int64
	var err error
	if SearchId != "" {
		sId, err = strconv.ParseInt(SearchId, 10, 10)
		if err != nil {
			beego.Info("000000")
			beego.Debug("get Search 0 Fail", err)
			return
		}
	}

	o := orm.NewOrm()
	obj := new(Question)
	query := o.QueryTable(obj)
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
	id, err := o.QueryTable(info).Filter("Id", id).Update(orm.Params{"IsIgnore": 0})
	return id, err
}

//批量纸条提问
func PrepareDelQuestion(IdArray []int64) (int64, error) {
	o := orm.NewOrm()
	err := o.Begin()
	var status int64
	beego.Info("IdArrayBBBBB:", IdArray)
	for i := 0; i < len(IdArray); i++ {
		status, err = o.Delete(&Question{Id: IdArray[i]})
		_, err2 := o.Raw("delete from rspquestion where question_id = ?", IdArray[i]).Exec()
		if err2 != nil {
			err = o.Rollback()
			return status, err
		}
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

//根据questionId删除rspquestion
func DelRspByQuestionId(QuestionId int64) (int64, error) {
	beego.Info("QuestionId:", QuestionId)
	o := orm.NewOrm()
	res, err := o.Raw("delete from rspquestion where `question_id` = ?", QuestionId).Exec()
	num, _ := res.RowsAffected()
	return num, err
}
