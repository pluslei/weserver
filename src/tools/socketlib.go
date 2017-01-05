package tools

import (
	"encoding/json"
	"math/rand"
	"time"
)

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
	//InSider int    //人员类别内部人员或外部人员
	Roomval [2][]string //房间
}

//用户列表信息
type Usertitle struct {
	Id        int64
	Uname     string //用户名
	RoleName  string //角色名称
	Titlerole string //用户类型
	UserIcon  string //用户Icon
	Roomid    string //房间号
	//Authorcss []string  //用户头衔
	Authorcss string    //用户头衔
	Logintime string    //登入时间
	Datatime  time.Time //添加时间
	Ipaddress string    //ip地址
	Procities string    //省市
	InSider   int       //人员类别内部人员或外部人员
	Insort    int       //用户排序
}

//socket内容发送信息
type Socketjson struct {
	Id          int64
	Author      string //用户名
	AuthorRole  string //角色名称
	Authortype  string //用户类别
	AuditStatus int    //信息是否审核,0：不用审核，1：审核通过，2：未审核(游客，会员需要审核)
	//Authorcss  []string  //用户头衔
	Authorcss     string //用户头衔
	AuthorDelay   int    //用户禁言时间
	AuthorInSider int    //人员类别内部人员或外部人员
	IsPrivateChat bool   //是否私聊
	IsBroadCast   bool   //是否广播
	Sendtype      string //用户发送消息类型，TXT, IMG

	Username string //私聊的用户名
	UserRole string //角色名称
	Usertype string //用户类别
	//Usercss   []string //用户头衔
	Usercss       string    //用户头衔
	UserInSider   int       //人员类别内部人员或外部人员
	Objname       string    //操作人员用户名
	Chat          string    //公聊，对他说，私聊
	Content       string    //发送的内容
	Coderoom      int       //房间号
	Codeid        string    //解码公司代码和房间号
	Time          string    //时间
	Newtime       string    //时间
	Datatime      time.Time //添加时间
	Ipaddress     string    //ip地址
	Procities     string    //省市
	IsEmitBroad   bool      //消息发送模式
	RoleTitleCss  string    // 头衔颜色
	RoleTitleBack bool      // 聊天背景颜色
}

var Resultuser []Usertitle  //模拟的用户数据
var Copyresuser []Usertitle //拷贝数据

func Jsontosocket(req string) (s []Socketjson, err error) {
	var result []Socketjson
	if err := json.Unmarshal([]byte(req), &result); err != nil {
		result = make([]Socketjson, 0)
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
