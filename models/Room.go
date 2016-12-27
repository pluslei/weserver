package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Room struct {
	Id              int64
	Centerid        int64  //中心服务器Id
	RommNumber      int64  //房间号
	Nickname        string //房间别名
	Rtmpurl         string //推流地址
	Streams         string //流名称
	ActivityId      string //活动id
	CompanyCode     string //公司代码
	RoomStatus      int    //房间状态 status:0 暂停推流 status:1 开启推流
	RoomDescription string //房间描述
}

func (r *Room) TableName() string {
	return "room"
}

func init() {
	orm.RegisterModel(new(Room))
}

// 显示
func GetRoomList(page int64, page_size int64, sort string) (rooms []orm.Params, count int64) {
	o := orm.NewOrm()
	room := new(Room)
	qs := o.QueryTable(room)
	count, _ = qs.Limit(page_size, page).OrderBy(sort).Values(&rooms) //, "Id", "Title", "Group_id", "Url", "Sort", "Hide"
	return rooms, count
}

// 插入数据
func InsertRoom(r *Room) (int64, error) {
	model := orm.NewOrm()
	inserid, err := model.Insert(r)
	return inserid, err
}

// 查询房间的数量
func RoomNumber() (int64, error) {
	model := orm.NewOrm()
	number, err := model.QueryTable("room").Count()
	return number, err
}

// 查询第一个房间信息
func GetFristerRoom() (roominfo Room, err error) {
	model := orm.NewOrm()
	err = model.QueryTable("room").OrderBy("RommNumber").Limit(1).One(&roominfo)
	return roominfo, err
}

// 查询房间是否存在
func GetIsRoom(Streams string, ActivityId string) int64 {
	model := orm.NewOrm()
	num, _ := model.QueryTable("room").Filter("Streams", Streams).Filter("ActivityId", ActivityId).Count()
	return num
}

// 查询房间信息
func GetRoomInfo(id int64) (roominfo Room, err error) {
	model := orm.NewOrm()
	err = model.QueryTable("room").Filter("Id", id).One(&roominfo)
	return roominfo, err
}

// 根据房间号查询
func GetRoomNumber(roomid int64) (roominfo Room, err error) {
	model := orm.NewOrm()
	err = model.QueryTable("room").Filter("RommNumber", roomid).One(&roominfo)
	return roominfo, err
}

// 更新房间
func UpdateRoom(id int64, nickname string, desction string) (int64, error) {
	model := orm.NewOrm()
	num, err := model.QueryTable("room").Filter("Id", id).Update(orm.Params{
		"Nickname":        nickname,
		"RoomDescription": desction,
	})
	return num, err
}

//获取所有的房间号
func GetAllRoomDate() ([]*Room, int64, error) {
	var (
		roominfo []*Room
		num      int64 = 0
		err      error
	)
	omodel := orm.NewOrm()
	num, err = omodel.QueryTable("Room").Limit(-1).All(&roominfo)
	return roominfo, num, err
}
