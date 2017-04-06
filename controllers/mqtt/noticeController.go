package mqtt

import (
	"time"
	m "weserver/models"
	mq "weserver/src/mqtt"

	"github.com/astaxie/beego"

	"strconv"
	"weserver/controllers"
	. "weserver/src/tools"
	// for json get
)

type NoticeController struct {
	controllers.PublicController
}

type noticeMessage struct {
	infochan chan *NoticeInfo
	Delchan  chan *NoticeDEL
}

var (
	notice *noticeMessage
)

func init() {
	notice = &noticeMessage{
		infochan: make(chan *NoticeInfo, 20480),
		Delchan:  make(chan *NoticeDEL, 20480),
	}
	notice.runWriteDb()
}

//发布公告
func (this *NoticeController) GetPublishNotice() {
	if this.IsAjax() {
		noticMsg := this.GetString("str")
		b := parseNoticeMsg(noticMsg)
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

//删除公告
func (this *NoticeController) DeleteNotice() {
	if this.IsAjax() {
		room := this.GetString("Room")
		id, _ := this.GetInt64("Id")
		b := DelNotice(room, id)
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

func parseNoticeMsg(msg string) bool {
	msginfo := new(NoticeInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("simplejson error", err)
		return false
	}
	info.MsgType = MSG_TYPE_NOTICE_ADD //公告
	topic := info.Room

	beego.Debug("info", info)

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}

	mq.SendMessage(topic, v) //发消息

	// 消息入库
	insertMsgdata(info)
	return true
}

func DelNotice(room string, id int64) bool {
	var info NoticeDEL
	info.Id = id
	info.Room = room
	info.MsgType = MSG_TYPE_NOTICE_DEL

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("DELETE Notice JSON ERROR", err)
		return false
	}
	mq.SendMessage(room, v) //发消息
	DeleteMsg(info)
	return true
}

//获取历史公告
func (this *NoticeController) GetNoticeList() {
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
		broaddata, _ := m.GetNoticeData(int(roomid))
		data["notice"] = broaddata //公告
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

//根据消息id 从数据库获取相应的消息
func (this *NoticeController) GetMsgInfoFromDatabase(id int64) MessageInfo {
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

// 获取在线用户信息列表
func (this *NoticeController) GetOnlineUseInfo() {
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

// 写数据
func (n *noticeMessage) runWriteDb() {
	go func() {
		for {
			infoMsg, ok := <-n.infochan
			if ok {
				addContent(infoMsg)
			}
			infoDel, ok1 := <-n.Delchan
			if ok1 {
				delContent(infoDel)
			}
		}
	}()
}

func insertMsgdata(info NoticeInfo) {
	jsondata := &info
	select {
	case notice.infochan <- jsondata:
		break
	default:
		beego.Error("WRITE NOTICE db error!!!")
		break
	}
}

func DeleteMsg(info NoticeDEL) {
	jsondata := &info
	select {
	case notice.Delchan <- jsondata:
		break
	default:
		beego.Error("DELETE NOTICE db error!!!")
		break
	}
}

func delContent(info *NoticeDEL) {
	beego.Debug("NoticeDEL", info)
	//写数据库
	var notice m.Notice
	notice.Id = info.Id
	notice.Room = info.Room
	_, err := m.DelNoticeById(notice.Id)
	if err != nil {
		beego.Debug("AddNotice Fail:", err)
	}
}

func addContent(info *NoticeInfo) {
	beego.Debug("NoticeInfo", info)
	//写数据库
	var notice m.Notice
	notice.Room = info.Room
	notice.Uname = info.Uname
	notice.Nickname = info.Nickname
	notice.Data = info.Content
	notice.Datatime = time.Now()

	_, err := m.AddNoticeMsg(&notice)
	if err != nil {
		beego.Debug("AddNotice Fail:", err)
	}
}
