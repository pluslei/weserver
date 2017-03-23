package mqtt

import (
	"strings"
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

var w *WriteData

type WriteData struct {
	infochan chan *MessageInfo
}

type MqttController struct {
	controllers.PublicController
}

var (
	recordcount int = 10 //历史消息显示数量
)

func init() {
	w = &WriteData{
		infochan: make(chan *MessageInfo, 20480),
	}
	w.runWriteDb()
}

// 获取发送消息
func (this *MqttController) GetMessageToSend() {
	if this.IsAjax() {
		msg := this.GetString("str")
		this.ParseMsg(msg)
	}
	this.Ctx.WriteString("")
}

//解析消息
func (this *MqttController) ParseMsg(msg string) {
	var info MessageInfo
	msg = DecodeB64(msg)
	key := []byte(msg)
	js, err := simplejson.NewJson(key)
	if err != nil {
		beego.Error(err)
	}
	codeid := js.Get("Codeid").MustString()
	codeid = Transformname(codeid, "", -1) //解码公司代码和房间号
	code, _ := strconv.ParseInt(beego.AppConfig.String("company"), 10, 64)
	info.Code = int(code)                                               //公司代码
	room, _ := strconv.ParseInt(beego.AppConfig.String("room"), 10, 64) //房间号
	info.Room = int(room)
	info.Uuid = js.Get("Uuid").MustString()                 //uuid
	info.Uname = js.Get("Uname").MustString()               //用户名
	info.Nickname = js.Get("Nickname").MustString()         //用户昵称
	info.UserIcon = js.Get("UserIcon").MustString()         //用户logo
	info.RoleName = js.Get("RoleName").MustString()         //用户角色[vip,silver,gold,jewel]
	info.RoleTitle = js.Get("RoleTitle").MustString()       //用户角色名[会员,白银会员,黄金会员,钻石会员]
	info.Sendtype = js.Get("Sendtype").MustString()         //用户发送消息类型('TXT','IMG','VOICE')
	info.RoleTitleCss = js.Get("RoleTitleCss").MustString() //头衔颜色
	info.RoleTitleBack = js.Get("RoleTitleBack").MustBool() //角色聊天背景
	info.Insider = js.Get("Insider").MustInt()              //1内部人员或0外部人员
	info.IsLogin = js.Get("IsLogin").MustBool()             //状态 [1、登录 0、未登录]
	info.Content = js.Get("Content").MustString()           //消息内容
	info.IsFilter = js.Get("IsFilter").MustBool()           //消息是否过滤[true: 过滤, false: 不过滤]
	info.Status = js.Get("Status").MustInt()                //审核状态(0：未审核，1：审核)
	info.Datatime = time.Now()                              //添加时间
	// 消息入库
	beego.Debug("insider message:", info.Insider)
	mq.SendMessage(info) //发消息

	SaveChatMsgdata(info)
	beego.Debug("aaaaa", info)
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
			historychat, _, _ = m.GetChatMsgData(recordcount)
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
		// topic := mq.Config.MqTopic // this.GetTopic()
		// mq.SendMessage(topic, msgInfo) //发消息
	}
}

//获取在线人数
func (this *MqttController) GetOnlineUseCount() {
	if this.IsAjax() {
		_, usercount := m.VirtualUserList(30)
		// data := make(map[string]interface{})
		// data["count"] = usercount
		// this.Data["json"] = &data
		// this.ServeJSON()
		this.Data["json"] = &map[string]interface{}{"status": true, "count": usercount}
		this.ServeJSON()
	} else {
		this.Ctx.WriteString("")
	}
}

// 获取在线用户信息列表
func (this *MqttController) GetOnlineUseInfo() {
	if this.IsAjax() {
		count := this.GetString("count")
		num, error := strconv.Atoi(count)
		if num <= 0 || error != nil {
			beego.Error("GetOnlineUseInfo Fail!!")
		}
		defult_Rsp, _ := strconv.ParseInt(beego.AppConfig.String("Defult_OnLine_Rsp"), 10, 64) // 默认发送的列表条数
		userlist, userlen := m.VirtualUserList(30)                                             //人员总列表信息
		end := num
		if end > userlen {
			end = userlen
		}
		var userinfor []m.VirtualUser
		start := num - int(defult_Rsp)
		for i := start; i < end; i++ {
			if len(userlist[i].UserIcon) > 0 {
				var msg m.VirtualUser
				msg.Id = userlist[i].Id
				msg.Username = EncodeB64(userlist[i].Username)
				msg.Nickname = EncodeB64(userlist[i].Nickname)
				msg.UserIcon = EncodeB64(userlist[i].UserIcon)
				userinfor = append(userinfor, msg)
			}
		}
		data := make(map[string]interface{})
		data["listinfo"] = userinfor
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
	case w.infochan <- jsondata:
		break
	default:
		beego.Error("write db error!!!")
		break
	}
}

// 写数据
func (w *WriteData) runWriteDb() {
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
		chatrecord.Uuid = info.Uuid                 //uuid
		chatrecord.Code = info.Code                 //公司代码
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
			// 插入成功广播
			broadcastChat(chatrecord)
		}
	}
}

//rpc 推送 给管理页面
func broadcastChat(chat m.ChatRecord) {
	chat.DatatimeStr = chat.Datatime.Format("2006-01-02 15:04:05")
	rpc.Broadcast("chat", chat, func(result []string) { beego.Debug("result", result) })
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
