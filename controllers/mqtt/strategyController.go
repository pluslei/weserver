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
	operchan chan *StrategyOperate
}

var (
	strategy *strategyMessage
)

func init() {
	strategy = &strategyMessage{
		infochan: make(chan *StrategyInfo, 20480),
		operchan: make(chan *StrategyOperate, 20480),
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

//策略操作 置顶/取消置顶/点赞/删除
func (this *StrategyController) OperateStrategy() {
	if this.IsAjax() {
		room := this.GetString("Room")
		id, _ := this.GetInt64("Id")
		op, _ := this.GetInt64("OperType")
		b := OPStrategy(room, id, op)
		if b {
			this.Rsp(true, "策略操作成功", "")
			return
		} else {
			this.Rsp(false, "策略操作失败,请重新发送", "")
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

func OPStrategy(room string, id, op int64) bool {
	var info StrategyOperate
	info.Id = id
	info.Room = room
	info.OperType = op

	info.MsgType = MSG_TYPE_STRATEGY_OPE

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("OPERATE Strategy JSON ERROR", err)
		return false
	}
	mq.SendMessage(room, v) //发消息
	OperateStrategyMsg(info)
	return true
}

// 写数据
func (n *strategyMessage) runWriteDb() {
	go func() {
		for {
			select {
			case infoMsg, ok := <-n.infochan:
				if ok {
					addStrategyContent(infoMsg)
				}
			case infoOper, ok1 := <-n.operchan:
				if ok1 {
					OperateStrategyContent(infoOper)
				}
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

func OperateStrategyMsg(info StrategyOperate) {
	jsondata := &info
	select {
	case strategy.operchan <- jsondata:
		break
	default:
		beego.Error("OPER NOTICE db error!!!")
		break
	}
}

func OperateStrategyContent(info *StrategyOperate) {
	beego.Debug("StrategyOper", info)
	var strategy m.Strategy
	strategy.Id = info.Id
	strategy.Room = info.Room
	OPERTYPE := info.OperType
	switch OPERTYPE {
	case OPERATE_TOP:
		beego.Debug("top")
		break
	case OPERATE_UNTOP:
		beego.Debug("untop")
		break
	case OPERATE_THUMB:
		beego.Debug("thumb")
		break
	case OPERATE_DEL:
		_, err := m.DelStrategyById(strategy.Id)
		if err != nil {
			beego.Debug("Oper Strategy Fail:", err)
		}
		break
	default:
	}
}

func addStrategyContent(info *StrategyInfo) {
	beego.Debug("Add StrategyInfo", info)
	var strategy m.Strategy
	strategy.Room = info.Room
	strategy.Icon = info.Icon
	strategy.Name = info.Name
	strategy.Titel = info.Titel
	strategy.Data = info.Data
	strategy.IsTop = info.IsTop
	strategy.IsDelete = info.IsDelete
	strategy.ThumbNum = info.ThumbNum
	strategy.Time = info.Time
	strategy.Datatime = time.Now()

	_, err := m.AddStrategy(&strategy)
	if err != nil {
		beego.Debug("Add Strategy Fail:", err)
	}
}
