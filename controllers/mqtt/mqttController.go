package mqtt

import (
	"crypto/rand"
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

// 获取聊天消息
func (this *MqttController) GetMessageToSend() {
	if this.IsAjax() {
		chatmsg := this.GetString("str")
		status := parseMsg(chatmsg)
		switch status {
		case POST_STATUS_TRUE:
			this.Rsp(true, "消息发送成功", "")
			return
		case POST_STATUS_FALSE:
			this.Rsp(false, "消息发送失败,请重新发送", "")
			return
		case POST_STATUS_SHUTUP:
			this.Rsp(false, "您已被禁言", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

// 聊天消息
func parseMsg(msg string) int {
	msginfo := new(MessageInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("simplejson error", err)
		return POST_STATUS_FALSE
	}
	info.Datatime = time.Now()       //添加时间
	info.MsgType = MSG_TYPE_CHAT_ADD //消息类型
	topic := info.Room

	beego.Debug("info", info)

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("json error", err)
		return POST_STATUS_FALSE
	}
	for _, v := range mq.Slice {
		if info.Uname == v {
			return POST_STATUS_SHUTUP
		}
	}
	// 内部人员
	if info.IsFilter == false {
		mq.SendMessage(topic, v) //发消息
	}
	beego.Debug("isfilter", info.IsFilter)
	// 消息入库
	SaveChatMsgdata(info)
	return POST_STATUS_TRUE
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
		roomid := this.GetString("codeid")    //房间id
		sysconfig, _ := m.GetAllSysConfig()   //系统设置
		recordcount := sysconfig.HistoryCount //显示历史记录条数
		var historychat []m.ChatRecord
		switch sysconfig.HistoryMsg { //是否显示历史消息 0显示  1 不显示
		case 0:
			historychat, _, _ = m.GetChatMsgData(recordcount, roomid, "chat_record")
		default:
		}

		data := make(map[string]interface{})
		data["historydata"] = historychat //聊天的历史信息
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

//聊天消息入库
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
					updateData(infoMsg)
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
		chatrecord.Uuid = info.Uuid                 //uuid
		chatrecord.Room = info.Room                 //房间号
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
			// 推送管理页面
			broadcastChat(chatrecord)
		}
	}
}

func updateData(info *MessageInfo) {
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
