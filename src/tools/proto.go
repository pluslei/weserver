package tools

import (
	"encoding/json"
	"math/rand"
	"time"
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

//在线用户ip和省市
type OnlineIpPro struct {
	Ip  string //ip地址
	Pro string //省市
}

//用户角色对应的房间
type RoleRoom struct {
	Roomval [2][]string //房间
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

//mqtt发送信息
type MessageInfo struct {
	Id            int64     //数据库中id
	Code          int       //公司代码
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
	MsgType       int       //消息类型
	Datatime      time.Time //添加时间
	Status        int       //审核状态(0：未审核，1：审核)
	Uuid          string    //uuid
}

// func (msg *MessageInfo) serialization() string {

// }

// func (msg *MessageInfo) Parse(string) error {

// }
// 广播消息
type BrocastInfo struct {
	Code    int    //公司代码
	Room    int    //房间号
	MsgType int    //消息类型
	Content string //广播内容
}

// 删除消息
type DelMessage struct {
	Code    int    //公司代码
	Room    int    //房间号
	MsgType int    //消息类型
	Uuid    string //消息uuid
}

//在线人数信息
type OnlineUserMsg struct {
	Nickname string //用户昵称
	UserIcon string //用户logo
}

var Resultuser []Usertitle  //模拟的用户数据
var Copyresuser []Usertitle //拷贝数据

func (m *MessageInfo) MashJson(msg []byte) (s MessageInfo, err error) {
	var result MessageInfo
	if err := json.Unmarshal(msg, &result); err != nil {
		return result, err
	}
	return result, nil
}

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
