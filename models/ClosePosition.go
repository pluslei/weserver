package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

/*
*  	平仓操作
 */
type ClosePosition struct {
	Id           int64  `orm:"pk;auto"`
	RoomId       string //topic
	RoomTeacher  string //老师
	Time         time.Time
	Type         string        //种类
	BuySell      string        //买卖
	Entrust      string        //委托类型
	Index        string        //点位
	Position     string        //仓位
	ProfitPoint  string        //止盈点
	LossPoint    string        //止损点
	Notes        string        // 备注
	OperPosition *OperPosition `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(ClosePosition))
}

func (c *ClosePosition) TableName() string {
	return "closeposition"
}

/*
* 新增加平仓操作
 */
func AddClosePosition(c *ClosePosition) (int64, error) {
	omodel := orm.NewOrm()
	id, err := omodel.Insert(c)
	return id, err
}

// 根据建仓id 获取多个平仓操作
func GetMoreClosePosition(id int64) ([]*ClosePosition, int64, error) {
	o := orm.NewOrm()
	var close []*ClosePosition
	num, err := o.QueryTable("ClosePosition").Filter("OperPosition", id).RelatedSel().All(&close)
	return close, num, err
}

/*
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
*/
