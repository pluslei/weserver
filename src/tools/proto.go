package tools

import (
	"encoding/json"
	"math/rand"

	"github.com/astaxie/beego"
)

// post status code
const (
	POST_STATUS_TRUE int = iota
	POST_STATUS_FALSE
	POST_STATUS_SHUTUP
)

// Role type match database table "Role"
const (
	ROLE_RESERVE int = iota
	ROLE_MANAGER
	ROLE_CUSTOMER
	ROLE_ASSISTANT
	ROLE_TEACHER
	ROLE_NORMAL
	ROLE_TOURIST
)

// msg type
const (
	MSG_TYPE_CHAT_ADD int = iota //聊天消息
	MSG_TYPE_CHAT_DEL
	MSG_TYPE_QUESTION_ADD
	MSG_TYPE_QUESTION_DEL
	MSG_TYPE_NOTICE_ADD //公告消息
	MSG_TYPE_NOTICE_DEL
	MSG_TYPE_STRATEGY_ADD //策略消息
	MSG_TYPE_STRATEGY_OPE
	MSG_TYPE_STRATEGY_UPDATE
	MSG_TYPE_KICKOUT     //踢人
	MSG_TYPE_SHUTUP      // 禁言
	MSG_TYPE_TEACHER_ADD //teacher
	MSG_TYPE_TEACHER_DEL
	MSG_TYPE_POSITION_ADD
	MSG_TYPE_POSITION_DEL
	MSG_TYPE_CLOSEPOSITION_ADD
	MSG_TYPE_CLOSEPOSITION_DEL
)

// operate type
const (
	OPERATE_KICKOUT = iota
	OPERATE_SHUTUP
	OPERATE_UNSHUTUP //取消禁言
)

//房间信息
type RoomInfo struct {
	RoomIcon    string //房间图标
	RommTitle   string //房间名
	RoomTeacher string //老师
	RoomNum     string //关注人数
}

//用户信息
type Membermsg struct {
	Totalmembers  int64  //会员人数
	Totalonline   int64  //在线人数
	Totalroom     int64  //房间数
	Totallinetime string //在线总时长
}

//OnlineUserList内容发送信息
type OnlineUserList struct {
	Roomid     string //房间号
	Uname      string //用户名
	Logintime  string //登入时间
	Onlinetime string //在线时长
	Ipaddress  string //ip地址
	Procities  string //省市
}

//用户列表信息
type Usertitle struct {
	Uname     string //微信唯一标识
	Nickname  string //微信名
	UserIcon  string //微信头像
	RoleName  string //头衔名称
	RoleTitle string //用户头衔名称
	InSider   int    //人员类别内部人员或外部人员[1: 内部人员，0：]
	IsLogin   bool   //状态 [true、登录 false、未登录]
}

//######################################################################################

type RoleInfo struct {
	RoleId        int64
	RoleName      string //用户角色[vip,silver,gold,jewel]
	RoleTitle     string //用户角色名[会员,白银会员,黄金会员,钻石会员]
	RoleTitleCss  string //用户角色样式
	RoleTitleBack bool   //角色聊天背景
}

//######################################################################################

// 在线人数信息
type OnLineInfo struct {
	Room     string
	Uname    string //微信唯一标识
	Nickname string //微信名
	UserIcon string //微信头像
	ShutUp   string //禁言状态
}

//######################################################################################

//mqtt发送聊天信息
type MessageInfo struct {
	Id            int64  //数据库中id
	CompanyId     int64  //公司id
	Room          string //房间号 topic
	Uname         string //用户名 openid
	Nickname      string //用户昵称
	UserIcon      string //用户logo
	RoleName      string //用户角色[vip,silver,gold,jewel]
	RoleTitle     string //用户角色名[会员,白银会员,黄金会员,钻石会员]
	Sendtype      string //用户发送消息类型('TXT','IMG','VOICE')
	RoleTitleCss  string //头衔颜色
	RoleTitleBack bool   //角色聊天背景
	Insider       int    //1内部人员或0外部人员
	IsLogin       bool   //状态 [1、登录 0、未登录]
	Content       string //消息内容
	IsFilter      bool   //消息是否过滤[true: 过滤, false: 不过滤]
	Status        int    //审核状态(0：未审核，1：审核)
	Uuid          string //uuid

	AcceptUuid    string
	AcceptTitle   string
	AcceptContent string

	MsgType int //消息类型
}

type MessageDEL struct {
	Uuid    string //消息uuid
	Room    string //房间号
	MsgType int    //消息类型
}

func (m *MessageInfo) ParseJSON(msg []byte) (s MessageInfo, err error) {
	var result MessageInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################
// notice msg

type NoticeInfo struct {
	CompanyId int64
	Room      string //房间号
	Uname     string //操作者的用户名
	Nickname  string
	Content   string //广播内容
	Time      string //发送公告时间

	MsgType int //消息类型
}
type NoticeDEL struct {
	Id      int64  //消息id 唯一
	Room    string //房间号
	MsgType int    //消息类型
}

func (n *NoticeInfo) ParseJSON(msg []byte) (s NoticeInfo, err error) {
	var result NoticeInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (n *NoticeDEL) ParseJSON(msg []byte) (s NoticeDEL, err error) {
	var result NoticeDEL
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################
//question

// Operate type
const (
	OPERATE_ASK_QUESTION = iota
	OPERATE_RSP_QUESTION
	OPERATE_IGN_QUESTION
)

type QuestionInfo struct {
	Id            int64
	CompanyId     int64
	Room          string //房间号 topic
	Uname         string //用户名  openid
	Nickname      string //用户昵称
	UserIcon      string //用户logo
	RoleName      string //用户角色[vip,silver,gold,jewel]
	RoleTitle     string //用户角色名[会员,白银会员,黄金会员,钻石会员]
	Sendtype      string //用户发送消息类型('TXT','IMG','VOICE')
	RoleTitleCss  string //头衔颜色
	RoleTitleBack int    //角色聊天背景
	Content       string //消息内容
	IsIgnore      int64  //是否忽略 0 显示 1 忽略
	Uuid          string // uuid

	AcceptNickname string
	AcceptTitle    string
	AcceptUuid     string
	OperateType    int64
	MsgType        int //消息类型
}

type QuestionDEL struct {
	Uuid    string //消息uuid
	Room    string //房间号
	MsgType int    //消息类型
}

func (m *QuestionInfo) ParseJSON(msg []byte) (s QuestionInfo, err error) {
	var result QuestionInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################
// Operate type
const (
	OPERATE_TOP = iota
	OPERATE_UNTOP
	OPERATE_THUMB
	OPERATE_UNTHUMB
	OPERATE_DEL
	OPERATE_ADD
	OPERATE_UPDATE
)

// 策略消息
type StrategyInfo struct {
	Id        int64
	CompanyId int64
	Room      string //房间号 topic
	Icon      string //头像
	Name      string //操作者的用户名
	Titel     string
	Data      string //策略内容
	FileName  string //图片
	TxtColour string //颜色字段
	IsPush    bool   //微信推送
	IsTop     bool   //是否置顶 置顶1 否 0
	IsDelete  bool   //是否删除,删除 1 否 0
	ThumbNum  int64  //点赞次数
	Time      string

	OperType int64
	MsgType  int //消息类型
}

// 置顶 /取消置顶 /点赞/ 删除
type StrategyOperate struct {
	Id       int64  //消息id 唯一
	Room     string //房间号
	OperType int64  // 1 /0	/2 /3

	MsgType int //消息类型
}

func (t *StrategyInfo) ParseJSON(msg []byte) (s StrategyInfo, err error) {
	var result StrategyInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (t *StrategyOperate) ParseJSON(msg []byte) (s StrategyOperate, err error) {
	var result StrategyOperate
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################

// Operate type
const (
	OPERATE_POSITION_ADD = iota
	OPERATE_POSITION_DEL
	OPERATE_POSITION_UPDATE
)

// 建仓信息
type PositionInfo struct {
	Id          int64
	CompanyId   int64
	RoomId      string //topic
	RoomTeacher string //老师
	Type        string //种类
	BuySell     int    //买卖 0 1
	Entrust     string //委托类型
	Index       string //点位
	Position    string //仓位
	ProfitPoint string //止盈点
	LossPoint   string //止损点
	Notes       string // 备注
	Liquidation int    //平仓详情 (0:未平仓 1:平仓)
	Icon        string //头像
	IsPush      bool   //是否推送到微信
	OperType    int64
	MsgType     int //消息类型

	//平仓信息
	/********************/
	CloseBuySell int    //平仓操作
	CloseIndex   string //平仓点位
	CloseNotes   string //平仓备注
}

type PositionOperate struct {
	Id       int64  //消息id 唯一
	Room     string //房间号
	OperType int64  // 1 /0	/2 /3

	MsgType int //消息类型
}

func (t *PositionInfo) ParseJSON(msg []byte) (s PositionInfo, err error) {
	var result PositionInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (t *PositionOperate) ParseJSON(msg []byte) (s PositionOperate, err error) {
	var result PositionOperate
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################
// Operate type
const (
	OPERATE_CLOSEPOSITION_ADD = iota
	OPERATE_CLOSEPOSITION_DEL
	OPERATE_CLOSEPOSITION_UPDATE
)

// 平仓信息
type ClosePositionInfo struct {
	Id          int64  //开仓信息的id
	RoomId      string //topic
	RoomTeacher string //老师
	Type        string //种类
	BuySell     int    //买卖 0 1
	Entrust     string //委托类型
	Index       string //点位
	Position    string //仓位
	ProfitPoint string //止盈点
	LossPoint   string //止损点
	Notes       string // 备注

	OperType int64
	MsgType  int //消息类型
}

type ClosePositionOperate struct {
	Id       int64  //消息id 唯一
	Room     string //房间号
	OperType int64  // 1 /0	/2 /3

	MsgType int //消息类型
}

func (t *ClosePositionInfo) ParseJSON(msg []byte) (s ClosePositionInfo, err error) {
	var result ClosePositionInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (t *ClosePositionOperate) ParseJSON(msg []byte) (s ClosePositionInfo, err error) {
	var result ClosePositionInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################

// Operate type
const (
	OPERATE_TEACHER_ADD = iota
	OPERATE_TEACHER_DEL
	OPERATE_TEACHER_UPDATE
	OPERATE_TEACHER_TOP
	OPERATE_TEACHER_UNTOP
	OPERATE_TEACHER_THUMB
	OPERATE_TEACHER_UNTHUMB
)

// 老师信息
type TeacherInfo struct {
	Id        int64
	CompanyId int64
	Room      string //房间号 topic
	Name      string //teacher name
	Icon      string //头像
	Title     string
	IsTop     bool   //是否置顶 置顶1 否 0
	ThumbNum  int64  //点赞次数
	Data      string //老师简介
	Time      string
	OperType  int64

	MsgType int //消息类型
}

type TeacherOperate struct {
	Id        int64 //消息id 唯一
	CompanyId int64
	Room      string //房间号
	Nickname  string // op name
	Username  string

	OperType int64

	MsgType int //消息类型
}

func (t *TeacherInfo) ParseJSON(msg []byte) (s TeacherInfo, err error) {
	var result TeacherInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (t *TeacherOperate) ParseJSON(msg []byte) (s TeacherOperate, err error) {
	var result TeacherOperate
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################

//踢人
type KickOutInfo struct {
	CompanyId int64
	Room      string //房间号 topic
	OperUid   string //踢人uuid
	OperName  string //踢人的用户名
	ObjUid    string //被踢的uuid
	ObjName   string //被踢的用户名

	MsgType int //消息类型
}

func (k *KickOutInfo) ParseJSON(msg []byte) (s []KickOutInfo, err error) {
	var result []KickOutInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################

//禁言
type ShutUpInfo struct {
	CompanyId int64
	Room      string //房间号 topic
	Uname     string
	IsShutUp  bool //是否禁言 1 否 0

	MsgType int //消息类型
}

func (k *ShutUpInfo) ParseJSON(msg []byte) (s []ShutUpInfo, err error) {
	var result []ShutUpInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################
//person setting

type SetInfo struct {
	Uname     string
	CompanyId int64
	Icon      string
	Nickname  string
}

//######################################################################################

func ParseJSONArray(msg []byte) (s []interface{}, err error) {
	var result []interface{}
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

func ParseJSON(msg []byte) (s interface{}, err error) {
	var result interface{}
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

func ToJSON(v interface{}) (string, error) {
	value, err := json.Marshal(v)
	if err != nil {
		beego.Error("json marshal error", err)
		return "", err
	}
	return string(value), nil
}

//在线人数信息
type OnlineUserMsg struct {
	Nickname string //用户昵称
	UserIcon string //用户logo
}

//生成一个新的验证码
func NewRandName(randlen int, s rand.Source) string {
	var strrandname string
	Letternumber := []byte(`ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678`) // 默认去掉了容易混淆的字符oOLl,9gq,Vv,Uu,I1
	length := len(Letternumber)
	r := rand.New(s)
	for i := 0; i < randlen; i++ {
		n := r.Intn(length)
		strrandname += string(Letternumber[n])
	}
	return strrandname
}
