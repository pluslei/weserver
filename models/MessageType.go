package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 消息分类
type MessageType struct {
	Id               int64      // 分类标识
	Name             string     // 分类名称
	Level            int64      // 分类等级
	CreateTime       time.Time  // 创建时间
	CreateMan        string     // 创建人
	Message          []*Message `orm:"reverse(many)"`
	CreateTimeFormat string     `orm:"-"` // 格式化时间，无实际意义
}

// 定义数据库表名称
func (m *MessageType) TableName() string {
	return "message_type"
}

// 初始化表
func init() {
	orm.RegisterModel(new(MessageType))
}

// 获取消息分类列表
func GetMessageTypeListByPager(page int64, page_size int64, sort string) (mts []orm.Params, count int64) {
	o := orm.NewOrm()
	mt := new(MessageType)
	query := o.QueryTable(mt)
	query.Limit(page_size, page).OrderBy(sort).Values(&mts)
	count, _ = query.Count()
	return mts, count
}

// 获取消息分类集合
func GetMessageTypeList() (mts []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("message_type").OrderBy("-CreateTime").Values(&mts)
	return mts, err
}

// 添加消息分类
func AddMessageType(m *MessageType) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	return id, err
}

// 修改消息分类
func EditMessageType(m *MessageType) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(m)
	return id, err
}

// 删除消息分类
func DelMessageType(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&MessageType{Id: Id})
	return status, err
}

// 根据消息分类 Id 获取消息分类
func GetMessageTypeById(Id int64) (MessageType, error) {
	o := orm.NewOrm()
	mt := MessageType{Id: Id}
	err := o.Read(&mt)
	if err != nil {
		return mt, err
	}
	return mt, nil
}
