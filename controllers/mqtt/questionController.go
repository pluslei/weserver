package mqtt

import (
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

type questionMessage struct {
	infochan chan *QuestionInfo
}

type QuestionController struct {
	controllers.PublicController
}

var (
	question *questionMessage
)

func init() {
	question = &questionMessage{
		infochan: make(chan *QuestionInfo, 20480),
	}
	question.runWriteDb()
}

// get online teacher list send private msg
func (this *QuestionController) GetQuestionTeacher() {
	if this.IsAjax() {
		Id, err := this.GetInt64("CompanyId")
		if err != nil {
			beego.Debug("Get CompanyId Fail", err)
			return
		}
		roomId := this.GetString("RoomId")
		var infoMsg []m.Regist
		teacher, _, err := m.GetRegistInfoByRole(Id, int64(ROLE_TEACHER), roomId)
		if err != nil {
			beego.Debug("Get CompanyInfo Error", err)
			return
		}
		for _, v := range teacher {
			var info m.Regist
			info.CompanyId = v.CompanyId
			info.Room = v.Room
			info.Username = v.Username
			info.UserId = v.UserId
			info.Nickname = v.Nickname
			info.UserIcon = v.UserIcon
			v.Titlename, err = m.GetTitleName(v.Title.Id)
			info.Titlename = v.Titlename
			infoMsg = append(infoMsg, info)
		}
		data := make(map[string]interface{})
		data["TeacherInfo"] = infoMsg
		this.Data["json"] = &data
		this.ServeJSON()
	}
	this.Ctx.WriteString("")
}

func (this *QuestionController) GetQuestionToSend() {
	if this.IsAjax() {
		chatmsg := this.GetString("str")
		status := parseQuestMsg(chatmsg)
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

func parseQuestMsg(msg string) int {
	msginfo := new(QuestionInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("simplejson error", err)
		return POST_STATUS_FALSE
	}

	info.MsgType = MSG_TYPE_QUESTION_ADD //消息类型

	topic := info.Room

	beego.Debug("info", info)

	// v, err := ToJSON(info)
	// if err != nil {
	// 	beego.Error("json error", err)
	// 	return POST_STATUS_FALSE
	// }
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
			beego.Debug("interface{} type is no define")
		}
	}

	// mq.SendMessage(topic, v) //发消息

	// 消息入库
	SaveQuestionMsgdata(info)
	return POST_STATUS_TRUE
}

func (this *QuestionController) GetQuestionTeacherRsp() {
	if this.IsAjax() {
		chatmsg := this.GetString("str")
		parseRspMsg(chatmsg)
	}
	this.Ctx.WriteString("")
}

func parseRspMsg(msg string) int {
	msginfo := new(QuestionInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("simplejson error", err)
		return POST_STATUS_FALSE
	}

	info.MsgType = MSG_TYPE_QUESTION_ADD //消息类型

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
			beego.Debug("interface{} type is no define")
		}
	}

	mq.SendMessage(topic, v) //发消息

	// 消息入库
	SaveQuestionMsgdata(info)
	return POST_STATUS_TRUE
}

//question List
func (this *QuestionController) GetQuestionHistoryList() {
	if this.IsAjax() {
		strId := this.GetString("Id")
		beego.Debug("id", strId)
		nId, _ := strconv.ParseInt(strId, 10, 64)
		roomId := this.GetString("room")
		username := this.GetString("username")
		RoleId, err := this.GetInt64("RoleId")
		if err != nil {
			beego.Debug("QuestionList get RoleId error", err)
			return
		}
		beego.Debug("Get Question List info  RoomId, Id ", nId, roomId, username, RoleId)

		data := make(map[string]interface{})
		sysconfig, _ := m.GetAllSysConfig()
		sysCount := sysconfig.QuestionCount
		var infoMsg []m.Question
		switch sysconfig.HistoryMsg { //是否显示历史消息 0显示  1 不显示
		case 0:
			historyMsg, totalCount, _ := m.GetAllQuestionMsg(roomId, username, RoleId)
			if nId == 0 {
				var i int64
				if totalCount < sysCount {
					beego.Debug("nCount sysCont", totalCount, sysCount)
					for i = 0; i < totalCount; i++ {
						var info m.Question
						info.Id = historyMsg[i].Id
						info.CompanyId = historyMsg[i].CompanyId
						info.Room = historyMsg[i].Room
						info.Uname = historyMsg[i].Uname
						info.Nickname = historyMsg[i].Nickname
						info.UserIcon = historyMsg[i].UserIcon
						info.RoleName = historyMsg[i].RoleName
						info.RoleTitle = historyMsg[i].RoleTitle
						info.Sendtype = historyMsg[i].Sendtype
						info.RoleTitleCss = historyMsg[i].RoleTitleCss
						info.RoleTitleBack = historyMsg[i].RoleTitleBack
						info.Content = historyMsg[i].Content
						info.Uuid = historyMsg[i].Uuid
						info.DatatimeStr = historyMsg[i].DatatimeStr
						info.AcceptNickname = historyMsg[i].AcceptNickname
						info.AcceptTitle = historyMsg[i].AcceptTitle
						historyquestion, count, err := m.GetMoreRspQuestion(info.Id)
						if count != 0 {
							info.RspNickname = historyquestion[0].Nickname
							info.RspTitle = historyquestion[0].RoleTitle
							info.RspIcon = historyquestion[0].UserIcon
							info.RspContent = historyquestion[0].Content
							info.RspTimestr = historyquestion[0].DatatimeStr
						}
						if err != nil && count != 0 {
							beego.Debug("Get RspQuestion info error", err)
							return
						}
						infoMsg = append(infoMsg, info)
					}
				} else {
					for i = 0; i < sysCount; i++ {
						var info m.Question
						info.Id = historyMsg[i].Id
						info.Room = historyMsg[i].Room
						info.Uname = historyMsg[i].Uname
						info.Nickname = historyMsg[i].Nickname
						info.UserIcon = historyMsg[i].UserIcon
						info.RoleName = historyMsg[i].RoleName
						info.RoleTitle = historyMsg[i].RoleTitle
						info.Sendtype = historyMsg[i].Sendtype
						info.RoleTitleCss = historyMsg[i].RoleTitleCss
						info.RoleTitleBack = historyMsg[i].RoleTitleBack
						info.Content = historyMsg[i].Content
						info.Uuid = historyMsg[i].Uuid
						info.DatatimeStr = historyMsg[i].DatatimeStr
						info.AcceptNickname = historyMsg[i].AcceptNickname
						info.AcceptTitle = historyMsg[i].AcceptTitle
						historyquestion, count, err := m.GetMoreRspQuestion(info.Id)
						if count != 0 {
							info.RspNickname = historyquestion[0].Nickname
							info.RspTitle = historyquestion[0].RoleTitle
							info.RspIcon = historyquestion[0].UserIcon
							info.RspContent = historyquestion[0].Content
							info.RspTimestr = historyquestion[0].DatatimeStr
						}
						if err != nil && count != 0 {
							beego.Debug("Get RspQuestion info error", err)
							return
						}
						infoMsg = append(infoMsg, info)
					}
				}
				data["historyQuestion"] = infoMsg
				this.Data["json"] = &data
				this.ServeJSON()
			} else {
				var index int64
				for nindex, value := range historyMsg {
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
						var info m.Question
						info.Id = historyMsg[i].Id
						info.Room = historyMsg[i].Room
						info.Uname = historyMsg[i].Uname
						info.Nickname = historyMsg[i].Nickname
						info.UserIcon = historyMsg[i].UserIcon
						info.RoleName = historyMsg[i].RoleName
						info.RoleTitle = historyMsg[i].RoleTitle
						info.Sendtype = historyMsg[i].Sendtype
						info.RoleTitleCss = historyMsg[i].RoleTitleCss
						info.RoleTitleBack = historyMsg[i].RoleTitleBack
						info.Content = historyMsg[i].Content
						info.Uuid = historyMsg[i].Uuid
						info.DatatimeStr = historyMsg[i].DatatimeStr
						info.AcceptNickname = historyMsg[i].AcceptNickname
						info.AcceptTitle = historyMsg[i].AcceptTitle
						historyquestion, count, err := m.GetMoreRspQuestion(info.Id)
						if count != 0 {
							info.RspNickname = historyquestion[0].Nickname
							info.RspTitle = historyquestion[0].RoleTitle
							info.RspIcon = historyquestion[0].UserIcon
							info.RspContent = historyquestion[0].Content
							info.RspTimestr = historyquestion[0].DatatimeStr
						}
						if err != nil && count != 0 {
							beego.Debug("Get RspQuestion info error", err)
							return
						}
						infoMsg = append(infoMsg, info)
					}
				} else {
					for i := index; i < totalCount; i++ {
						var info m.Question
						info.Id = historyMsg[i].Id
						info.Room = historyMsg[i].Room
						info.Uname = historyMsg[i].Uname
						info.Nickname = historyMsg[i].Nickname
						info.UserIcon = historyMsg[i].UserIcon
						info.RoleName = historyMsg[i].RoleName
						info.RoleTitle = historyMsg[i].RoleTitle
						info.Sendtype = historyMsg[i].Sendtype
						info.RoleTitleCss = historyMsg[i].RoleTitleCss
						info.RoleTitleBack = historyMsg[i].RoleTitleBack
						info.Content = historyMsg[i].Content
						info.Uuid = historyMsg[i].Uuid
						info.DatatimeStr = historyMsg[i].DatatimeStr
						info.AcceptNickname = historyMsg[i].AcceptNickname
						info.AcceptTitle = historyMsg[i].AcceptTitle
						historyquestion, count, err := m.GetMoreRspQuestion(info.Id)
						if count != 0 {
							info.RspNickname = historyquestion[0].Nickname
							info.RspTitle = historyquestion[0].RoleTitle
							info.RspIcon = historyquestion[0].UserIcon
							info.RspContent = historyquestion[0].Content
							info.RspTimestr = historyquestion[0].DatatimeStr
						}
						if err != nil && count != 0 {
							beego.Debug("Get RspQuestion info error", err)
							return
						}
						infoMsg = append(infoMsg, info)
					}
				}
				data["historyQuestion"] = infoMsg
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

func SaveQuestionMsgdata(info QuestionInfo) {
	jsondata := &info
	select {
	case question.infochan <- jsondata:
		break
	default:
		beego.Error("write db error!!!")
		break
	}
}

// 写数据
func (w *questionMessage) runWriteDb() {
	go func() {
		for {
			infoMsg, ok := <-w.infochan
			if ok {
				operateQuestionData(infoMsg)
			}
		}
	}()
}

func operateQuestionData(info *QuestionInfo) {
	beego.Debug("PositionOperate", info)
	var op m.Question
	op.Id = info.Id
	op.Room = info.Room
	OPERTYPE := info.OperateType
	beego.Debug("operate type", OPERTYPE)
	switch OPERTYPE {
	case OPERATE_ASK_QUESTION:
		if op.Id == 0 {
			status := addQuestionData(info)
			if !status {
				beego.Debug("Add Question Fail")
				return
			}
		}
		break
	case OPERATE_RSP_QUESTION:
		if op.Id != 0 {
			status := addRspQuestionData(info)
			if !status {
				beego.Debug("Add RspQuestion Fail")
				return
			}
		}
		break
	case OPERATE_IGN_QUESTION:
		status := OpereateIgnQuestionData(info)
		if !status {
			beego.Debug("Upadate Ignore RspQuestion Fail")
			return
		}
		break
	default:
	}
}

func addQuestionData(info *QuestionInfo) bool {
	var question m.Question
	question.Uuid = info.Uuid //uuid
	question.CompanyId = info.CompanyId
	question.Room = info.Room                 //房间号
	question.Uname = info.Uname               //用户名
	question.Nickname = info.Nickname         //用户昵称
	question.UserIcon = info.UserIcon         //用户logo
	question.RoleName = info.RoleName         //用户角色[vip,silver,gold,jewel]
	question.RoleTitle = info.RoleTitle       //用户角色名[会员,白银会员,黄金会员,钻石会员]
	question.Sendtype = info.Sendtype         //用户发送消息类型('TXT','IMG','VOICE')
	question.RoleTitleCss = info.RoleTitleCss //头衔颜色
	question.RoleTitleBack = info.RoleTitleBack
	question.Content = info.Content //消息内容
	question.IsIgnore = 1           // 默认显示
	question.Time = time.Now()
	question.DatatimeStr = question.Time.Format("2006-01-02 15:04:05")
	question.AcceptNickname = info.AcceptNickname
	question.AcceptTitle = info.AcceptTitle
	_, err := m.AddQuestion(&question)
	if err != nil {
		beego.Debug(err)
		return false
	} else {
		// 推送管理页面
		pushWebAdmin(question)
		return true
	}
}

func addRspQuestionData(info *QuestionInfo) bool {
	var Rsp m.RspQuestion
	Rsp.Uuid = info.Uuid //uuid
	Rsp.CompanyId = info.CompanyId
	Rsp.Room = info.Room                 //房间号
	Rsp.Uname = info.Uname               //用户名
	Rsp.Nickname = info.Nickname         //用户昵称
	Rsp.UserIcon = info.UserIcon         //用户logo
	Rsp.RoleName = info.RoleName         //用户角色[vip,silver,gold,jewel]
	Rsp.RoleTitle = info.RoleTitle       //用户角色名[会员,白银会员,黄金会员,钻石会员]
	Rsp.Sendtype = info.Sendtype         //用户发送消息类型('TXT','IMG','VOICE')
	Rsp.RoleTitleCss = info.RoleTitleCss //头衔颜色
	Rsp.RoleTitleBack = info.RoleTitleBack
	Rsp.Content = info.Content //消息内容
	Rsp.Time = time.Now()
	Rsp.DatatimeStr = Rsp.Time.Format("2006-01-02 15:04:05")
	Rsp.Question = &m.Question{Id: info.Id}
	_, err := m.AddRspQuestion(&Rsp)
	if err != nil {
		beego.Debug(err)
		return false
	} else {
		// 推送管理页面
		// pushWebAdmin(question)
		return true
	}
}

func OpereateIgnQuestionData(info *QuestionInfo) bool {
	_, err := m.OpeateIgnore(info.Id)
	if err != nil {
		beego.Debug("Opeate Ignore Question error", err)
		return false
	}
	return true
}

//rpc 推送 给管理页面
func pushWebAdmin(chat m.Question) {
	// chat.DatatimeStr = chat.Datatime.Format("2006-01-02 15:04:05")
	rpc.Broadcast("chat", chat, func(result []string) { beego.Debug("result", result) })
}
