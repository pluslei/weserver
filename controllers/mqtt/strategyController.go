package mqtt

import (
	"strconv"
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

//策略操作 置顶/取消置顶/点赞/取消点赞/删除
func (this *StrategyController) OperateStrategy() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseOPStrategyMsg(msg)
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

//Strategy List
func (this *StrategyController) GetStrategyList() {
	if this.IsAjax() {
		count := this.GetString("count")
		nEnd, _ := strconv.ParseInt(count, 10, 64)
		roomId := this.GetString("room")
		data := make(map[string]interface{})
		sysconfig, _ := m.GetAllSysConfig()
		sysCount := sysconfig.StrategyCount
		var Strinfo []m.Strategy
		historyStrategy, nCount, _ := m.GetStrategyList(roomId)
		if nCount < sysCount {
			beego.Debug("nCount sysCont", nCount, sysCount)
			var i int64
			for i = 0; i < nCount; i++ {
				var info m.Strategy
				info.Id = historyStrategy[i].Id
				info.Room = historyStrategy[i].Room
				info.Icon = historyStrategy[i].Icon
				info.Titel = historyStrategy[i].Titel
				info.Data = historyStrategy[i].Data
				info.IsTop = historyStrategy[i].IsTop
				info.IsDelete = historyStrategy[i].IsDelete
				info.ThumbNum = historyStrategy[i].ThumbNum
				info.Time = historyStrategy[i].Time
				Strinfo = append(Strinfo, info)
			}
			data["historyStrategy"] = Strinfo
			this.Data["json"] = &data
			this.ServeJSON()
			return
		}
		mod := (nEnd - nCount) % sysCount
		beego.Debug("mod", mod)
		if nEnd > nCount && mod == 0 {
			beego.Debug("mod = 0")
			data["historyStrategy"] = ""
			this.Data["json"] = &data
			this.ServeJSON()
			return
		}
		var nstart int64
		nstart = nEnd - sysCount
		if nEnd > nCount {
			nEnd = nCount
			mod = nEnd % sysCount
			nstart = nEnd - mod
			beego.Debug("mod", mod)
		}
		for i := nstart; i < nEnd; i++ {
			var info m.Strategy
			info.Id = historyStrategy[i].Id
			info.Room = historyStrategy[i].Room
			info.Icon = historyStrategy[i].Icon
			info.Titel = historyStrategy[i].Titel
			info.Data = historyStrategy[i].Data
			info.IsTop = historyStrategy[i].IsTop
			info.IsDelete = historyStrategy[i].IsDelete
			info.ThumbNum = historyStrategy[i].ThumbNum
			info.Time = historyStrategy[i].Time
			Strinfo = append(Strinfo, info)
		}
		data["historyStrategy"] = Strinfo
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
		beego.Error("Strategy: simplejson error", err)
		return false
	}
	info.MsgType = MSG_TYPE_STRATEGY_ADD
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

func parseOPStrategyMsg(msg string) bool {
	msginfo := new(StrategyOperate)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("StrategyOperate: simplejson error", err)
		return false
	}
	info.MsgType = MSG_TYPE_STRATEGY_OPE
	room := info.Room
	beego.Debug("Operate Strategy info", info)

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
		_, err := m.StickOption(strategy.Id)
		if err != nil {
			beego.Debug("Oper Strategy Top Fail", err)
		}
		break
	case OPERATE_UNTOP:
		_, err := m.UnStickOption(strategy.Id)
		if err != nil {
			beego.Debug("Oper Strategy UnTop Fail", err)
		}
		break
	case OPERATE_THUMB:
		_, err := m.ThumbOptionAdd(strategy.Id)
		if err != nil {
			beego.Debug("Oper Strategy Thumb Fail", err)
		}
		break
	case OPERATE_UNTHUMB:
		_, err := m.ThumbOptionDel(strategy.Id)
		if err != nil {
			beego.Debug("Oper Strategy Thumb Fail", err)
		}
		break
	case OPERATE_DEL:
		_, err := m.DelStrategyById(strategy.Id)
		if err != nil {
			beego.Debug("Oper Strategy Delete Fail:", err)
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
