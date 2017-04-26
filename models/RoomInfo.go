package models

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
*  	房间信息
 */
type RoomInfo struct {
	Id          int64  `orm:"pk;auto"`
	RoomId      string //topic
	Qos         byte   `orm:"default(0)"` // mqtt 协议订阅等级
	RoomTitle   string //房间名
	RoomTeacher string //老师
	RoomNum     string //关注人数
	GroupId     string //组id
	Url         string //客户端
	Port        int
	Tls         bool
	Access      string
	SecretKey   string
	RoomIcon    string //房间图标
	RoomIntro   string `orm:"size(512)"` //简介
	RoomBanner  string //图片
	Title       string //标题
	MidPage     int64  //0 不显示 1 显示
}

func init() {
	orm.RegisterModel(new(RoomInfo))
}

func (c *RoomInfo) TableName() string {
	return "roominfo"
}

/*
* 新增加聊天室
 */
func AddRoom(r *RoomInfo) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(r)
	return id, err
}

//更新房间名
func UpdateRoomName(id int64, str string) (int64, error) {
	o := orm.NewOrm()
	var chat RoomInfo
	id, err := o.QueryTable(chat).Filter("Id", id).Update(orm.Params{"RoomTitle": str})
	return id, err
}

//更新房间小图标
func UpdateRoomIcon(id int64, str string) (int64, error) {
	o := orm.NewOrm()
	var chat RoomInfo
	id, err := o.QueryTable(chat).Filter("Id", id).Update(orm.Params{"RoomIcon": str})
	return id, err
}

//更新房间banner
func UpdateRoomBanner(id int64, str string) (int64, error) {
	o := orm.NewOrm()
	var chat RoomInfo
	id, err := o.QueryTable(chat).Filter("Id", id).Update(orm.Params{"RoomBanner": str})
	return id, err
}

//更新房间标题
func UpdateRoomTitle(id int64, str string) (int64, error) {
	o := orm.NewOrm()
	var chat RoomInfo
	id, err := o.QueryTable(chat).Filter("Id", id).Update(orm.Params{"Title": str})
	return id, err
}

//更新房间简介
func UpdateRoomIntro(id int64, str string) (int64, error) {
	o := orm.NewOrm()
	var chat RoomInfo
	id, err := o.QueryTable(chat).Filter("Id", id).Update(orm.Params{"RoomIntro": str})
	return id, err
}

//根据roomid 删除某个聊天室
func DelRoomById(roomid string) (int64, error) {
	o := orm.NewOrm()
	var chat RoomInfo
	status, err := o.QueryTable(chat).Filter("RoomId", roomid).Delete()
	return status, err
}

//事务添加多个聊天室
func AddMulRoom(room []RoomInfo, length int) error {
	model := orm.NewOrm()
	err := model.Begin()
	SuccessNum := 0
	if err == nil {
		for i := 0; i < length; i++ {
			id, err := model.Insert(&room[i])
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

//获取聊天室个数和聊天室名
func GetRoomName() (map[string]interface{}, int64, error) {
	o := orm.NewOrm()
	res := make(orm.Params)
	nums, err := o.Raw("SELECT room_id, qos FROM roominfo Order By Id").RowsToMap(&res, "room_id", "qos")
	return res, nums, err
}

//获取聊天室信息
func GetRoomInfo() ([]RoomInfo, int64, error) {
	o := orm.NewOrm()
	var info []RoomInfo
	num, err := o.QueryTable("roominfo").OrderBy("Id").All(&info)
	beego.Debug("num", num)
	return info, num, err
}

// 获取房间信息
func GetRoomInfoByRoomID(RoomId string) (info RoomInfo, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("roominfo").Filter("RoomId", RoomId).All(&info)
	return info, err
}
