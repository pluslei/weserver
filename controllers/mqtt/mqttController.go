package mqtt

import (
	"strings"
	"sync"
	"time"
	m "weserver/models"
	. "weserver/src/cache"
	mq "weserver/src/mqtt"
	rpc "weserver/src/rpcserver"

	"github.com/astaxie/beego"

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

// get company info and room info
func (this *MqttController) GetCompanyInfo() {
	if this.IsAjax() {
		Id, err := this.GetInt64("CompanyId")
		if err != nil {
			beego.Debug("Get CompanyId Fail", err)
			return
		}
		companyInfo, err := m.GetCompanyById(Id)
		if err != nil {
			beego.Debug("Get CompanyInfo Error", err)
			return
		}
		roomInfo, _, err := m.GetRoomInfo(Id)
		if err != nil {
			beego.Debug("GetRoomInfo fail", err)
			return
		}
		data := make(map[string]interface{})
		data["roomInfo"] = roomInfo //聊天室信息
		data["CompanyInfo"] = companyInfo
		this.Data["json"] = &data
		this.ServeJSON()
	}
	this.Ctx.WriteString("")
}

// 发送聊天消息
func (this *MqttController) GetMessageToSend() {
	if this.IsAjax() {
		chatmsg := this.GetString("str")
		status := parseMsg(chatmsg)
		switch status {
		case POST_STATUS_TRUE:
			this.Rsp(true, "POST_STATUS_TRUE", "")
			return
		case POST_STATUS_FALSE:
			this.Rsp(false, "POST_STATUS_FALSE", "")
			return
		case POST_STATUS_SHUTUP:
			this.Rsp(false, "POST_STATUS_SHUTUP", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

func parseMsg(msg string) int {
	msginfo := new(MessageInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("simplejson error", err)
		return POST_STATUS_FALSE
	}

	info.MsgType = MSG_TYPE_CHAT_ADD //消息类型

	topic := info.Room

	beego.Debug("info", info)

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("json error", err)
		return POST_STATUS_FALSE
	}
	inter, ok := MapCache[topic]
	if ok {
		arr, ok := inter.([]string)
		if ok {
			for _, v := range arr {
				if v == info.Uname {
					return POST_STATUS_SHUTUP
				}
			}
		} else {
			beego.Debug("interface{} no define")
		}
	}

	var company m.Company
	strId := strconv.FormatInt(info.CompanyId, 10)
	inter1, ok := MapCache[strId]
	if !ok {
		company, err = m.GetCompanyById(info.CompanyId)
		if err != nil {
			beego.Debug("get login companyinfo error")
		}
	} else {
		company, _ = inter1.(m.Company)
	}

	if company.AuditMsg == 0 {
		mq.SendMessage(topic, v) //发消息
	}
	beego.Debug("IsFilter", info.IsFilter)
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

//chat List
func (this *MqttController) GetChatHistoryList() {
	if this.IsAjax() {
		strId := this.GetString("Id")
		beego.Debug("id", strId)
		nId, _ := strconv.ParseInt(strId, 10, 64)
		roomId := this.GetString("room")
		beego.Debug("Get Chat List info  RoomId, Id ", nId, roomId)

		data := make(map[string]interface{})
		sysconfig, _ := m.GetAllSysConfig()
		sysCount := sysconfig.HistoryCount
		var infoChat []m.ChatRecord
		switch sysconfig.HistoryMsg { //是否显示历史消息 0显示  1 不显示
		case 0:
			historychat, totalCount, _ := m.GetAllChatMsgData(roomId, "chat_record")
			if nId == 0 {
				var i int64
				if totalCount < sysCount {
					beego.Debug("nCount sysCont", totalCount, sysCount)
					for i = 0; i < totalCount; i++ {
						var info m.ChatRecord
						info.Id = historychat[i].Id
						info.Room = historychat[i].Room
						info.Uname = historychat[i].Uname
						info.Nickname = historychat[i].Nickname
						info.UserIcon = historychat[i].UserIcon
						info.RoleName = historychat[i].RoleName
						info.RoleTitle = historychat[i].RoleTitle
						info.Sendtype = historychat[i].Sendtype
						info.RoleTitleCss = historychat[i].RoleTitleCss
						info.RoleTitleBack = historychat[i].RoleTitleBack
						info.Insider = historychat[i].Insider
						info.IsLogin = historychat[i].IsLogin
						info.Content = historychat[i].Content
						info.Status = historychat[i].Status
						info.Uuid = historychat[i].Uuid
						info.AcceptUuid = historychat[i].AcceptUuid
						info.AcceptTitle = historychat[i].AcceptTitle
						info.AcceptContent = historychat[i].AcceptContent
						infoChat = append(infoChat, info)
					}
				} else {
					for i = 0; i < sysCount; i++ {
						var info m.ChatRecord
						info.Id = historychat[i].Id
						info.Room = historychat[i].Room
						info.Uname = historychat[i].Uname
						info.Nickname = historychat[i].Nickname
						info.UserIcon = historychat[i].UserIcon
						info.RoleName = historychat[i].RoleName
						info.RoleTitle = historychat[i].RoleTitle
						info.Sendtype = historychat[i].Sendtype
						info.RoleTitleCss = historychat[i].RoleTitleCss
						info.RoleTitleBack = historychat[i].RoleTitleBack
						info.Insider = historychat[i].Insider
						info.IsLogin = historychat[i].IsLogin
						info.Content = historychat[i].Content
						info.Status = historychat[i].Status
						info.Uuid = historychat[i].Uuid
						info.AcceptUuid = historychat[i].AcceptUuid
						info.AcceptTitle = historychat[i].AcceptTitle
						info.AcceptContent = historychat[i].AcceptContent
						infoChat = append(infoChat, info)
					}
				}
				data["historyChat"] = infoChat
				this.Data["json"] = &data
				this.ServeJSON()
			} else {
				var index int64
				for nindex, value := range historychat {
					if value.Id == nId {
						index = int64(nindex) + 1
					}
				}
				beego.Debug("index", index)
				nCount := index + sysCount
				mod := (totalCount - nCount) % sysCount
				beego.Debug("mod", mod)
				if nCount > totalCount && mod == 0 {
					beego.Debug("mod = 0")
					data["historyChat"] = ""
					this.Data["json"] = &data
					this.ServeJSON()
					return
				}
				if nCount < totalCount {
					for i := index; i < nCount; i++ {
						var info m.ChatRecord
						info.Id = historychat[i].Id
						info.Room = historychat[i].Room
						info.Uname = historychat[i].Uname
						info.Nickname = historychat[i].Nickname
						info.UserIcon = historychat[i].UserIcon
						info.RoleName = historychat[i].RoleName
						info.RoleTitle = historychat[i].RoleTitle
						info.Sendtype = historychat[i].Sendtype
						info.RoleTitleCss = historychat[i].RoleTitleCss
						info.RoleTitleBack = historychat[i].RoleTitleBack
						info.Insider = historychat[i].Insider
						info.IsLogin = historychat[i].IsLogin
						info.Content = historychat[i].Content
						info.Status = historychat[i].Status
						info.Uuid = historychat[i].Uuid
						info.AcceptUuid = historychat[i].AcceptUuid
						info.AcceptTitle = historychat[i].AcceptTitle
						info.AcceptContent = historychat[i].AcceptContent
						infoChat = append(infoChat, info)
					}
				} else {
					for i := index; i < totalCount; i++ {
						var info m.ChatRecord
						info.Id = historychat[i].Id
						info.Room = historychat[i].Room
						info.Uname = historychat[i].Uname
						info.Nickname = historychat[i].Nickname
						info.UserIcon = historychat[i].UserIcon
						info.RoleName = historychat[i].RoleName
						info.RoleTitle = historychat[i].RoleTitle
						info.Sendtype = historychat[i].Sendtype
						info.RoleTitleCss = historychat[i].RoleTitleCss
						info.RoleTitleBack = historychat[i].RoleTitleBack
						info.Insider = historychat[i].Insider
						info.IsLogin = historychat[i].IsLogin
						info.Content = historychat[i].Content
						info.Status = historychat[i].Status
						info.Uuid = historychat[i].Uuid
						info.AcceptUuid = historychat[i].AcceptUuid
						info.AcceptTitle = historychat[i].AcceptTitle
						info.AcceptContent = historychat[i].AcceptContent
						infoChat = append(infoChat, info)
					}
				}
				data["historyChat"] = infoChat
				this.Data["json"] = &data
				this.ServeJSON()
			}
		default:
		}
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
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
		chatrecord.Uuid = info.Uuid //uuid
		chatrecord.CompanyId = info.CompanyId
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
		chatrecord.Insider = info.Insider //1内部人员或0外部人员
		chatrecord.IsLogin = 1            //状态 [1、登录 0、未登录]
		chatrecord.Content = info.Content //消息内容
		chatrecord.Datatime = time.Now()  //添加时间
		chatrecord.AcceptUuid = info.AcceptUuid
		chatrecord.AcceptTitle = info.AcceptTitle
		chatrecord.AcceptContent = info.AcceptContent
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
