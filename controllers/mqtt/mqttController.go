package mqtt

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"strings"
	"sync"
	"time"
	m "weserver/models"
	mq "weserver/src/mqtt"
	rpc "weserver/src/rpcserver"

	"github.com/astaxie/beego"
	simplejson "github.com/bitly/go-simplejson"

	"strconv"
	"weserver/controllers"
	. "weserver/src/tools"
	// for json get
)

const (
	MSG_TYPE_CHAT    int = iota //聊天消息
	MSG_TYPE_BROCAST            //广播消息
	MSG_TYPE_DEL                //删除消息
)

type MessageType struct {
	Code    int //公司代码
	Room    int //房间号码
	Msgtype int //消息类型
}

type historyMessage struct {
	infochan chan *MessageInfo
}

type MqttController struct {
	controllers.PublicController
}

var (
	history     *historyMessage
	totalLock   sync.Mutex
	total       int64      //在线人数
	recordcount int   = 10 //历史消息显示数量
)

func init() {
	history = &historyMessage{
		infochan: make(chan *MessageInfo, 20480),
	}
	history.runWriteDb()
	go userTotal()
}

func NewMessageType(msgtype int) *MessageType {
	code, _ := strconv.Atoi(beego.AppConfig.String("company"))
	room, _ := strconv.Atoi(beego.AppConfig.String("room"))
	return &MessageType{Code: code, Room: room, Msgtype: msgtype}
}

// 获取聊天室信息
func (this *MqttController) GetRoomInfo() {
	if this.IsAjax() {
		roomInfo, _, err := m.GetRoomInfo()
		if err != nil {
			beego.Debug("GetRoomInfo fail", err)
			return
		}
		data := make(map[string]interface{})
		data["roomInfo"] = roomInfo //聊天室信息
		this.Data["json"] = &data
		this.ServeJSON()
	}
	this.Ctx.WriteString("")
}

// 获取发送聊天消息
func (this *MqttController) GetMessageToSend() {
	if this.IsAjax() {
		chatmsg := this.GetString("str")
		msgtype := NewMessageType(MSG_TYPE_CHAT)
		b := msgtype.ParseMsg(chatmsg)
		if b {
			this.Rsp(true, "消息发送成功", "")
			return
		} else {
			this.Rsp(false, "消息发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

// 聊天消息
func (c *MessageType) ParseMsg(msg string) bool {
	msginfo := new(MessageInfo)
	info, err := msginfo.MashJson(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("simplejson error", err)
		return false
	}
	info.Code = c.Code //公司代码
	/*
		info.Room = c.Room           //房间号码
	*/
	info.Datatime = time.Now()   //添加时间
	info.MsgType = MSG_TYPE_CHAT //0 普通消息 1 广播

	beego.Debug("info", info)

	v, err := c.Json(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}

	// 内部人员
	if info.IsFilter == false {
		mq.SendMessage(v) //发消息
	}
	beego.Debug("isfilter", info.IsFilter)
	// 消息入库
	SaveChatMsgdata(info)
	return true
}

// 发送广播消息
func (c *MessageType) SendBrocast(content string) bool {
	info := new(BrocastInfo)
	info.Content = content
	info.Code = c.Code
	info.Room = c.Room
	info.MsgType = MSG_TYPE_BROCAST
	v, err := c.Json(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}
	mq.SendMessage(string(v)) //发消息
	return true
}

func (c *MessageType) DelMessage(uuid string) bool {
	info := new(DelMessage)
	info.Code = c.Code
	info.Room = c.Room
	info.MsgType = MSG_TYPE_DEL
	info.Uuid = uuid

	v, err := c.Json(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}
	mq.SendMessage(string(v)) //发消息
	return true
}

// 后台审核消息
func (c *MessageType) CheckMessage(msg m.ChatRecord) bool {
	msg.MsgType = MSG_TYPE_CHAT //0 普通消息 1 广播
	beego.Debug("msg", msg)

	v, err := c.Json(msg)
	if err != nil {
		beego.Error("json error", err)
		return false
	}
	mq.SendMessage(v) //发消息
	return true
}

func (c *MessageType) Json(v interface{}) (string, error) {
	value, err := json.Marshal(v)
	if err != nil {
		beego.Error("json marshal error", err)
		return "", err
	}
	return string(value), nil
}

//获取客户的真实IP地址
func (this *MqttController) GetClientip() string {
	var addrArr []string
	if len(this.Ctx.Request.Header.Get("X-Forwarded-For")) > 0 {
		addr := this.Ctx.Request.Header.Get("X-Forwarded-For")
		addrArr = strings.Split(addr, ":")
	} else if len(this.Ctx.Request.RemoteAddr) > 0 {
		addr := this.Ctx.Request.RemoteAddr
		addrArr = strings.Split(addr, ":")
	} else {
		addrArr[0] = "127.0.0.1"
	}
	return addrArr[0]
}

//获取聊天历史信息
func (this *MqttController) GetChatHistoryList() {
	if this.IsAjax() {
		codeid := this.GetString("codeid")              //公司房间标识符
		codeid = Transformname(codeid, "", -1)          //解码公司代码和房间号
		coderoom := Transformname(codeid, "", 2)        //房间号
		roomid, _ := strconv.ParseInt(coderoom, 10, 64) //房间号
		sysconfig, _ := m.GetAllSysConfig()             //系统设置
		recordcount := sysconfig.HistoryCount           //显示历史记录条数
		var historychat []m.ChatRecord
		switch sysconfig.HistoryMsg { //是否显示历史消息 0显示  1 不显示
		case 0:
			historychat, _, _ = m.GetChatMsgData(recordcount, "chat_record")
		default:
		}
		data := make(map[string]interface{})
		data["historydata"] = historychat //聊天的历史信息
		//从数据库中获取公告中的最后一条内容
		broaddata, _ := m.GetBroadcastData(int(roomid))
		data["notice"] = broaddata //公告
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

//根据消息id 从数据库获取相应的消息
func (this *MqttController) GetMsgInfoFromDatabase(id int64) MessageInfo {
	var info MessageInfo
	if id > 0 {
		chat, _ := m.GetChatIdData(id)
		if chat.Status == 1 {
			return info
		}
		info.Uname = chat.Uname               //用户名
		info.Nickname = chat.Nickname         //用户昵称
		info.UserIcon = chat.UserIcon         //用户logo
		info.RoleName = chat.RoleName         //用户角色[vip,silver,gold,jewel]
		info.RoleTitle = chat.RoleTitle       //用户角色名[会员,白银会员,黄金会员,钻石会员]
		info.Sendtype = chat.Sendtype         //用户发送消息类型('TXT','IMG','VOICE')
		info.RoleTitleCss = chat.RoleTitleCss //头衔颜色
		if chat.RoleTitleBack == 1 {
			info.RoleTitleBack = true //角色聊天背景
		} else {
			info.RoleTitleBack = false //角色聊天背景
		}
		if chat.IsLogin == 1 {
			info.IsLogin = true //状态 [1、登录 0、未登录]
		} else {
			info.IsLogin = false //状态 [1、登录 0、未登录]
		}
		info.Insider = chat.Insider //1内部人员或0外部人员
		info.Content = chat.Content //消息内容
		info.Uuid = chat.Uuid       //uuid
		info.IsFilter = true        //消息是否过滤[true: 过滤, false: 不过滤]
		info.Status = 1
		info.Datatime = chat.Datatime //添加时间
	}
	return info
}

// 获取后台审核的消息id
func (this *MqttController) GetPassId() {
	if this.IsAjax() {
		str := this.GetString("sendstr")
		msg := DecodeB64(str)
		key := []byte(msg)
		js, err := simplejson.NewJson(key)
		if err != nil {
			beego.Error(err)
		}
		id := js.Get("id").MustInt64()
		msgInfo := this.GetMsgInfoFromDatabase(id)
		beego.Debug("ddddddddddddddd", msgInfo)
		// 发消息
		//	mq.SendMessage(msgInfo) //发消息
		// topic := mq.Config.MqTopic // this.GetTopic()
		// mq.SendMessage(topic, msgInfo) //发消息
	}
}

//获取在线人数
func (this *MqttController) GetOnlineUseCount() {
	usercount := getToal()
	this.Data["json"] = &map[string]interface{}{"status": true, "count": usercount}
	this.ServeJSON()
}

// 获取在线用户信息列表
func (this *MqttController) GetOnlineUseInfo() {
	if this.IsAjax() {
		count := this.GetString("count")                //请求的数据总数
		listindex, _ := strconv.ParseInt(count, 10, 64) //客户端请求的列表个数
		data := make(map[string]interface{})
		if listindex > 0 {
			defult_Rsp, _ := strconv.ParseInt(beego.AppConfig.String("Defult_OnLine_Rsp"), 10, 64) // 默认发送的列表条数
			userlist, userlen := m.VirtualUserList(30)                                             //人员总列表信息
			listend := int(listindex)
			if listend > userlen {
				listend = userlen
			}
			var userinfor []m.VirtualUser
			liststart := int(listindex) - int(defult_Rsp)
			for i := liststart; i < listend; i++ {
				if len(userlist[i].UserIcon) > 0 {
					var msg m.VirtualUser
					msg.Id = userlist[i].Id
					msg.Username = EncodeB64(userlist[i].Username)
					msg.Nickname = EncodeB64(userlist[i].Nickname)
					msg.UserIcon = EncodeB64(userlist[i].UserIcon)
					userinfor = append(userinfor, msg)
				}
			}
			data["userlist"] = userinfor
		}
		_, onlinecount := m.VirtualUserList(30)
		data["onlinecount"] = onlinecount //在线人数
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

func (this *MqttController) ChatUpload() {
	this.Ctx.WriteString("")
}

func (this *MqttController) ChatKickOut() {
	this.Ctx.WriteString("")
}

func (this *MqttController) ChatModifyIcon() {
	if this.GetSession("indexUserInfo") != nil && this.IsAjax() {
		data := make(map[string]interface{})
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

/*
//广播入库
func SaveBroadCastdata(info MessageInfo) {
	//写数据库
	var broad m.Broadcast
	broad.Code = info.Code
	broad.Room = info.Room
	broad.Uname = DecodeB64(info.Uname)
	broad.Data = DecodeB64(info.Content)
	broad.Datatime = time.Now()

	_, err := m.AddBroadcast(&broad)
	if err != nil {
		beego.Debug(err)
	}

}
*/

//时时消息入库
func SaveChatMsgdata(info MessageInfo) {
	jsondata := &info
	select {
	case history.infochan <- jsondata:
		break
	default:
		beego.Error("write db error!!!")
		break
	}
}

// 写数据
func (w *historyMessage) runWriteDb() {
	go func() {
		for {
			infoMsg, ok := <-w.infochan
			if ok {
				if infoMsg.Status == 0 {
					addData(infoMsg)
				} else {
					UpdateData(infoMsg)
				}
			}
		}
	}()
}

func addData(info *MessageInfo) {
	beego.Debug("im here", info.IsFilter, info.RoleTitleBack)
	if info.IsLogin && info.Insider == 1 {
		//写数据库
		var chatrecord m.ChatRecord
		chatrecord.Uuid = info.Uuid //uuid
		chatrecord.Code = info.Code //公司代码
		/*
			chatrecord.Room = info.Room                 //房间号
		*/
		chatrecord.Uname = info.Uname               //用户名
		chatrecord.Nickname = info.Nickname         //用户昵称
		chatrecord.UserIcon = info.UserIcon         //用户logo
		chatrecord.RoleName = info.RoleName         //用户角色[vip,silver,gold,jewel]
		chatrecord.RoleTitle = info.RoleTitle       //用户角色名[会员,白银会员,黄金会员,钻石会员]
		chatrecord.Sendtype = info.Sendtype         //用户发送消息类型('TXT','IMG','VOICE')
		chatrecord.RoleTitleCss = info.RoleTitleCss //头衔颜色
		if info.RoleTitleBack {
			chatrecord.RoleTitleBack = 1 //角色聊天背景
		} else {
			chatrecord.RoleTitleBack = 0 //角色聊天背景
		}
		chatrecord.Insider = info.Insider   //1内部人员或0外部人员
		chatrecord.IsLogin = 1              //状态 [1、登录 0、未登录]
		chatrecord.Content = info.Content   //消息内容
		chatrecord.Datatime = info.Datatime //添加时间
		if !info.IsFilter {
			chatrecord.Status = 1 //审核状态(0：未审核，1：审核)
		} else {
			chatrecord.Status = info.Status //审核状态(0：未审核，1：审核)
		}

		_, err := m.AddChat(&chatrecord)
		if err != nil {
			beego.Debug(err)
		} else {
			// 插入成功广播
			broadcastChat(chatrecord)
		}
	}
}

func UpdateData(info *MessageInfo) {
	beego.Debug("im here", info, info.RoleTitleBack)
	if info.IsLogin && info.Insider == 1 {

		//更新数据库
		_, err := m.UpdateChatStatus(info.Id)
		if err != nil {
			beego.Debug(err)
		}
	}
}

//rpc 推送 给管理页面
func broadcastChat(chat m.ChatRecord) {
	chat.DatatimeStr = chat.Datatime.Format("2006-01-02 15:04:05")
	rpc.Broadcast("chat", chat, func(result []string) { beego.Debug("result", result) })
}

// 总人数
func userTotal() {
	t := time.Tick(time.Second * 5)
	for {
		<-t
		dayonlineuser, err := m.GetAllUserCount(30)
		if err != nil {
			beego.Error("get the usercount error", err)
		}
		sysconfig, _ := m.GetAllSysConfig()
		totalLock.Lock()

		total = dayonlineuser + sysconfig.VirtualUser + RandomInt64(-10, 10)
		totalLock.Unlock()
	}
}

func getToal() int64 {
	totalLock.Lock()
	defer totalLock.Unlock()
	return total
}

//随机数
func RandomInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	iInt64 := i.Int64()
	if iInt64 < min {
		iInt64 = RandomInt64(min, max)
	}
	return iInt64
}
