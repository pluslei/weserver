package mqtt

import (
	"time"
	m "weserver/models"
	mq "weserver/src/mqtt"

	"github.com/astaxie/beego"

	"weserver/controllers"
	. "weserver/src/tools"
	// for json get
)

type StrategyController struct {
	controllers.PublicController
}

type strategyMessage struct {
	infochan chan *StrategyInfo
	Delchan  chan *StrategyDEL
}

var (
	strategy *strategyMessage
)

func init() {
	strategy = &strategyMessage{
		infochan: make(chan *StrategyInfo, 20480),
		Delchan:  make(chan *StrategyDEL, 20480),
	}
	strategy.runWriteDb()
}

//发策略
func (this *StrategyController) GetStrategyInfo() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseStrategyMsg(msg)
		if b {
			this.Rsp(true, "策略发送成功", "")
			return
		} else {
			this.Rsp(false, "策略发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

//删除策略
func (this *StrategyController) DeleteStrategy() {
	if this.IsAjax() {
		room := this.GetString("Room")
		id, _ := this.GetInt64("Id")
		b := DelStrategy(room, id)
		if b {
			this.Rsp(true, "策略删除成功", "")
			return
		} else {
			this.Rsp(false, "策略删除失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

//策略列表
func (this *StrategyController) GetStrategyList() {
	if this.IsAjax() {
		roomId := this.GetString("room") //公司房间标识符
		sysconfig, _ := m.GetAllSysConfig()
		Count := sysconfig.StrategyCount
		var info []m.Strategy
		info, _, _ = m.GetStrategyList(roomId, Count)
		data := make(map[string]interface{})
		data["historyStrategy"] = info //公告的历史信息
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

func parseStrategyMsg(msg string) bool {
	msginfo := new(StrategyInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("simplejson error", err)
		return false
	}
	info.MsgType = MSG_TYPE_STRATEGY_ADD //公告
	topic := info.Room

	beego.Debug("info", info)

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}

	mq.SendMessage(topic, v) //发消息

	// 消息入库
	insertStrageydata(info)
	return true
}

func DelStrategy(room string, id int64) bool {
	var info StrategyDEL
	info.Id = id
	info.Room = room
	info.MsgType = MSG_TYPE_STRATEGY_DEL

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("DELETE Strategy JSON ERROR", err)
		return false
	}
	mq.SendMessage(room, v) //发消息
	DeleteStrategyMsg(info)
	return true
}

// 写数据
func (n *strategyMessage) runWriteDb() {
	go func() {
		for {
			infoMsg, ok := <-n.infochan
			if ok {
				addStrategyContent(infoMsg)
			}
			infoDel, ok1 := <-n.Delchan
			if ok1 {
				delStrategyContent(infoDel)
			}
		}
	}()
}

func insertStrageydata(info StrategyInfo) {
	jsondata := &info
	select {
	case strategy.infochan <- jsondata:
		break
	default:
		beego.Error("WRITE NOTICE db error!!!")
		break
	}
}

func DeleteStrategyMsg(info StrategyDEL) {
	jsondata := &info
	select {
	case strategy.Delchan <- jsondata:
		break
	default:
		beego.Error("DELETE NOTICE db error!!!")
		break
	}
}

func delStrategyContent(info *StrategyDEL) {
	beego.Debug("StrategyDEL", info)
	//写数据库
	var strategy m.Strategy
	strategy.Id = info.Id
	strategy.Room = info.Room
	_, err := m.DelStrategyById(strategy.Id)
	if err != nil {
		beego.Debug("Delete Strategy Fail:", err)
	}
}

func addStrategyContent(info *StrategyInfo) {
	beego.Debug("Add StrategyInfo", info)
	//写数据库
	var strategy m.Strategy
	strategy.Room = info.Room
	strategy.Icon = info.Icon
	strategy.Name = info.Name
	strategy.Titel = info.Titel
	strategy.Data = info.Data
	strategy.IsTop = info.IsTop
	strategy.IsDelete = info.IsDelete
	strategy.ThumbNum = info.ThumbNum
	strategy.Datatime = time.Now()

	_, err := m.AddStrategy(&strategy)
	if err != nil {
		beego.Debug("Add Strategy Fail:", err)
	}
}
