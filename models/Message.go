package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 消息
type Message struct {
	Id               int64        // 消息标识
	Content          string       `orm:"type(text)"` // 消息内容
	CreateTime       time.Time    // 创建时间
	CreateMan        string       // 创建人
	MessageType      *MessageType `orm:"rel(fk)"` // 分类
	CreateTimeFormat string       `orm:"-"`       // 格式化时间，无实际意义
	TypeName         string       `orm:"-"`       // 分类名称，无实际意义
}

// 定义数据库表名称
func (m *Message) TableName() string {
	return "message"
}

// 初始化表
func init() {
	orm.RegisterModel(new(Message))
}

// 获取消息列表
func GetMessageListByPager(page int64, page_size int64, sort string) (ms []orm.Params, count int64) {
	o := orm.NewOrm()
	mt := new(Message)
	query := o.QueryTable(mt)
	query.Limit(page_size, page).OrderBy(sort).Values(&ms)
	count, _ = query.Count()
	return ms, count
}

// 获取消息集合
func GetMessageList() (ms []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("message").OrderBy("-CreateTime").Values(&ms)
	return ms, err
}

// 添加消息
func AddMessage(m *Message) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	return id, err
}

// 删除消息
func DelMessage(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Message{Id: Id})
	return status, err
}

// 修改消息
func EditMessage(m *Message) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(m)
	return id, err
}

// 根据消息 Id 获取消息
func GetMessageById(Id int64) (Message, error) {
	o := orm.NewOrm()
	m := Message{Id: Id}
	err := o.Read(&m)
	if err != nil {
		return m, err
	}
	return m, nil
}

// 根据消息分类id 查询所有消息
func GetMessageByTypeId(id int64) (messages []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("message").Filter("MessageType__Id__exact", id).Values(&messages)
	return messages, err
}
