package tools

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
)

const (
	MSG_TYPE_CHAT_ADD int = iota //聊天消息
	MSG_TYPE_CHAT_DEL
	MSG_TYPE_NOTICE_ADD   //公告消息
	MSG_TYPE_NOTICE_DEL   //公告消息
	MSG_TYPE_STRATEGY_ADD //策略消息
	MSG_TYPE_STRATEGY_OPE
	MSG_TYPE_KICKOUT //踢人
	MSG_TYPE_SHUTUP  // 禁言
)

//策略
const (
	OPERATE_TOP = iota
	OPERATE_UNTOP
	OPERATE_THUMB
	OPERATE_DEL
)

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

//mqtt发送聊天信息
type MessageInfo struct {
	Id            int64     //数据库中id
	Room          string    //房间号 topic
	Uname         string    //用户名 openid
	Nickname      string    //用户昵称
	UserIcon      string    //用户logo
	RoleName      string    //用户角色[vip,silver,gold,jewel]
	RoleTitle     string    //用户角色名[会员,白银会员,黄金会员,钻石会员]
	Sendtype      string    //用户发送消息类型('TXT','IMG','VOICE')
	RoleTitleCss  string    //头衔颜色
	RoleTitleBack bool      //角色聊天背景
	Insider       int       //1内部人员或0外部人员
	IsLogin       bool      //状态 [1、登录 0、未登录]
	Content       string    //消息内容
	IsFilter      bool      //消息是否过滤[true: 过滤, false: 不过滤]
	Datatime      time.Time //添加时间
	Status        int       //审核状态(0：未审核，1：审核)
	Uuid          string    //uuid

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

// 公告消息
type NoticeInfo struct {
	Room     string //房间号
	Uname    string //操作者的用户名
	Nickname string
	Content  string //广播内容
	Time     string //发送公告时间

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

//######################################################################################

// 策略消息
type StrategyInfo struct {
	Room     string //房间号 topic
	Icon     string //头像
	Name     string //操作者的用户名
	Titel    string
	Data     string //策略内容
	IsTop    bool   //是否置顶 置顶1 否 0
	IsDelete bool   //是否删除,删除 1 否 0
	ThumbNum int64  //点赞次数
	Time     string

	MsgType int //消息类型
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

//######################################################################################

//踢人
type KickOutInfo struct {
	Room     string //房间号 topic
	OperUid  string //踢人uuid
	OperName string //踢人的用户名
	ObjUid   string //被踢的uuid
	ObjName  string //被踢的用户名

	MsgType int //消息类型
}

func (k *KickOutInfo) ParseJSON(msg []byte) (s KickOutInfo, err error) {
	var result KickOutInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################

//禁言
type ShutUpInfo struct {
	Room     string //房间号 topic
	Uname    string
	IsShutUp bool //是否禁言 1 否 0

	MsgType int //消息类型
}

func (k *ShutUpInfo) ParseJSON(msg []byte) (s ShutUpInfo, err error) {
	var result ShutUpInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

//######################################################################################
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

var Resultuser []Usertitle  //模拟的用户数据
var Copyresuser []Usertitle //拷贝数据

func Jsontosocket(req string) (s []MessageInfo, err error) {
	var result []MessageInfo
	if err := json.Unmarshal([]byte(req), &result); err != nil {
		result = make([]MessageInfo, 0)
		return result, err
	}
	return result, nil
}

func Jsontoroommap(req string) (s map[string]Usertitle, err error) {
	var result map[string]Usertitle
	if err := json.Unmarshal([]byte(req), &result); err != nil {
		result = make(map[string]Usertitle)
		return result, err
	}
	return result, nil
}

func Jsontoroomcode(req string) (s map[string][]int, err error) {
	var result map[string][]int
	if err := json.Unmarshal([]byte(req), &result); err != nil {
		result = make(map[string][]int)
		return result, err
	}
	return result, nil
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
