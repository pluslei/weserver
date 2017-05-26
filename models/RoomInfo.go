package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
*  	房间信息
 */
type RoomInfo struct {
	Id          int64 `orm:"pk;auto"`
	CompanyId   int64
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
	PcRoomText  string // 仅pc端使用 免责声明
	PcRoomad    string // pc 端 广告
	MidPage     int64  //0 不显示 1 显示

	//公司信息
	CompanyName   string `orm:"-"`
	CompanyIntro  string `orm:"-"`
	CompanyIcon   string `orm:"-"`
	CompanyBanner string `orm:"-"`
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

// 根据id 删除某个聊天室
func DelRoomInfoId(id int64) (int64, error) {
	o := orm.NewOrm()
	var chat RoomInfo
	status, err := o.QueryTable(chat).Filter("Id", id).Delete()
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
func GetRoomInfo(id int64) ([]RoomInfo, int64, error) {
	o := orm.NewOrm()
	var info []RoomInfo
	num, err := o.QueryTable("roominfo").Filter("CompanyId", id).Filter("MidPage", 1).OrderBy("Id").All(&info)
	return info, num, err
}

func GetAllRoomInfo() ([]RoomInfo, int64, error) {
	o := orm.NewOrm()
	var info []RoomInfo
	num, err := o.QueryTable("roominfo").OrderBy("Id").All(&info)
	beego.Debug("num", num)
	return info, num, err
}

//获取聊天室个数和聊天室名
func GetRoomCompany(room string) (int64, error) {
	var info RoomInfo
	o := orm.NewOrm()
	err := o.QueryTable("roominfo").Filter("RoomId", room).Limit(1).One(&info)
	return info.CompanyId, err
}

// 获取房间信息
func GetRoomInfoByRoomID(RoomId string) (info RoomInfo, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("roominfo").Filter("RoomId", RoomId).All(&info)
	return info, err
}

// 判断是否存在
func IsRoomInfo(roomid string) bool {
	o := orm.NewOrm()
	return o.QueryTable("roominfo").Filter("RoomId", roomid).Exist()
}

// 获取房间信息
func GetRoomInfoById(id int64) (info RoomInfo, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("roominfo").Filter("Id", id).Limit(1).One(&info)
	return info, err
}

// 获取消息列表
func GetRoomInfoList(page int64, page_size int64, companyId int64, SearchId string) (ms []orm.Params, count int64) {
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
	roominfo := new(RoomInfo)

	if SearchId != "" {
		query := o.QueryTable(roominfo)
		query.Limit(page_size, page).Filter("CompanyId", sId).OrderBy("-Id").Values(&ms)
		count, _ = query.Count()
		return ms, count
	}
	if companyId != 0 {
		query := o.QueryTable(roominfo)
		query.Limit(page_size, page).Filter("CompanyId", companyId).OrderBy("-Id").Values(&ms)
		count, _ = query.Count()
		return ms, count
	}
	query := o.QueryTable(roominfo)
	query.Limit(page_size, page).OrderBy("-Id").Values(&ms)
	count, _ = query.Count()
	return ms, count
}

// 获取房间信息
func GetRoomInfoOne() (info RoomInfo, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("roominfo").Limit(1).One(&info)
	return info, err
}

// 更新房间信息
func UpdateRoomInfo(id int64, roominfo orm.Params) (int64, error) {
	beego.Debug("roominfo", roominfo, id)
	o := orm.NewOrm()
	return o.QueryTable("roominfo").Filter("Id", id).Update(roominfo)
}

// 获取最大数量
func GetRoomInfoCount() (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable(new(RoomInfo)).Count()
}

// 隐藏or显示房间
func UpdateRoomShowed(id, newShowed int64) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE roominfo SET `mid_page` = ? WHERE id = ?", newShowed, id).Exec()
	var resnum int64
	if err != nil {
		num, _ := res.RowsAffected()
		resnum = int64(num)
	}
	return resnum, err
}
