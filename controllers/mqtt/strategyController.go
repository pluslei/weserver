package mqtt

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
	m "weserver/models"
	mq "weserver/src/mqtt"
	"weserver/src/wechat"

	"github.com/astaxie/beego"

	"weserver/controllers"
	"weserver/controllers/haoindex"
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

//发策略
func (this *StrategyController) GetEditStrategyInfo() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseEditStrategyMsg(msg)
		if b {
			this.Rsp(true, "编辑策略发送成功", "")
			return
		} else {
			this.Rsp(false, "编辑策略发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

func (this *StrategyController) GetUnameMapInfo() {
	if this.IsAjax() {
		wechat.GetUnameMapInfo()
		beego.Debug("Uname Map", wechat.MapUname)
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
		strId := this.GetString("Id")
		beego.Debug("id", strId)
		nId, _ := strconv.ParseInt(strId, 10, 64)
		roomId := this.GetString("room")
		beego.Debug("Stragety list ", nId, roomId)
		data := make(map[string]interface{})
		sysconfig, _ := m.GetAllSysConfig()
		sysCount := sysconfig.StrategyCount
		var Strinfo []m.Strategy
		historyStrategy, totalCount, _ := m.GetStrategyList(roomId)
		if nId == 0 {
			var i int64
			if totalCount < sysCount {
				beego.Debug("nCount sysCont", totalCount, sysCount)
				for i = 0; i < totalCount; i++ {
					var info m.Strategy
					info.Id = historyStrategy[i].Id
					info.Room = historyStrategy[i].Room
					info.Icon = historyStrategy[i].Icon
					info.Name = historyStrategy[i].Name
					info.Titel = historyStrategy[i].Titel
					info.Data = historyStrategy[i].Data
					info.FileName = historyStrategy[i].FileName
					info.TxtColour = historyStrategy[i].TxtColour
					info.IsTop = historyStrategy[i].IsTop
					info.IsDelete = historyStrategy[i].IsDelete
					info.ThumbNum = historyStrategy[i].ThumbNum
					info.Time = historyStrategy[i].Time
					Strinfo = append(Strinfo, info)
				}
			} else {
				for i = 0; i < sysCount; i++ {
					var info m.Strategy
					info.Id = historyStrategy[i].Id
					info.Room = historyStrategy[i].Room
					info.Icon = historyStrategy[i].Icon
					info.Name = historyStrategy[i].Name
					info.Titel = historyStrategy[i].Titel
					info.Data = historyStrategy[i].Data
					info.FileName = historyStrategy[i].FileName
					info.TxtColour = historyStrategy[i].TxtColour
					info.IsTop = historyStrategy[i].IsTop
					info.IsDelete = historyStrategy[i].IsDelete
					info.ThumbNum = historyStrategy[i].ThumbNum
					info.Time = historyStrategy[i].Time
					Strinfo = append(Strinfo, info)
				}
			}
			data["historyStrategy"] = Strinfo
			this.Data["json"] = &data
			this.ServeJSON()
		} else {
			var index int64
			for nindex, value := range historyStrategy {
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
				data["historyStrategy"] = ""
				this.Data["json"] = &data
				this.ServeJSON()
				return
			}
			if nCount < totalCount {
				for i := index; i < nCount; i++ {
					var info m.Strategy
					info.Id = historyStrategy[i].Id
					info.Room = historyStrategy[i].Room
					info.Icon = historyStrategy[i].Icon
					info.Name = historyStrategy[i].Name
					info.Titel = historyStrategy[i].Titel
					info.Data = historyStrategy[i].Data
					info.FileName = historyStrategy[i].FileName
					info.TxtColour = historyStrategy[i].TxtColour
					info.IsTop = historyStrategy[i].IsTop
					info.IsDelete = historyStrategy[i].IsDelete
					info.ThumbNum = historyStrategy[i].ThumbNum
					info.Time = historyStrategy[i].Time
					Strinfo = append(Strinfo, info)
				}
			} else {
				for i := index; i < totalCount; i++ {
					var info m.Strategy
					info.Id = historyStrategy[i].Id
					info.Room = historyStrategy[i].Room
					info.Icon = historyStrategy[i].Icon
					info.Name = historyStrategy[i].Name
					info.Titel = historyStrategy[i].Titel
					info.Data = historyStrategy[i].Data
					info.FileName = historyStrategy[i].FileName
					info.TxtColour = historyStrategy[i].TxtColour
					info.IsTop = historyStrategy[i].IsTop
					info.IsDelete = historyStrategy[i].IsDelete
					info.ThumbNum = historyStrategy[i].ThumbNum
					info.Time = historyStrategy[i].Time
					Strinfo = append(Strinfo, info)
				}
			}
			data["historyStrategy"] = Strinfo
			this.Data["json"] = &data
			this.ServeJSON()
		}
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

/*
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
				info.Name = historyStrategy[i].Name
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
			info.Name = historyStrategy[i].Name
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
*/

func (this *StrategyController) Upload() {
	uploadtype := this.GetString("uploadtype")

	_, h, err := this.GetFile("file")
	if err != nil {
		beego.Error("getfile error", err)
		this.Rsp(false, uploadtype, "")
	}

	dir := path.Join("..", "upload", "room")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		beego.Error(err)
		this.Rsp(false, uploadtype, "")
	}
	// 设置保存文件名

	nowtime := time.Now().Unix()
	FileName := fmt.Sprintf("%d%s", nowtime, h.Filename)
	dirPath := path.Join("..", "upload", "room", FileName)
	// 将文件保存到服务器中
	err = this.SaveToFile("file", dirPath)
	if err != nil {
		beego.Error(err)
		this.Rsp(false, uploadtype, "")
	}
	filepath := path.Join("/upload", "room", FileName)
	this.Rsp(true, uploadtype, filepath)
}

func parseStrategyMsg(msg string) bool {
	msginfo := new(StrategyInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("Strategy: simplejson error", err)
		return false
	}
	info.OperType = OPERATE_ADD
	info.MsgType = MSG_TYPE_STRATEGY_ADD
	topic := info.Room
	sendmsg := info.Data

	beego.Debug("info", info)

	v, err := ToJSON(info)
	if err != nil {
		beego.Error("json error", err)
		return false
	}

	mq.SendMessage(topic, v)
	SendWeChatStrategy(topic, sendmsg) // send to wechat
	// 消息入库
	editStrageydata(info)
	return true
}

func parseEditStrategyMsg(msg string) bool {
	msginfo := new(StrategyInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("EditStrategy: simplejson error", err)
		return false
	}
	info.OperType = OPERATE_UPDATE
	/*
		info.MsgType = MSG_TYPE_STRATEGY_UPDATE
		topic := info.Room
		//sendmsg := info.Data

		beego.Debug("info", info)

		v, err := ToJSON(info)
		if err != nil {
			beego.Error("json error", err)
			return false
		}

		mq.SendMessage(topic, v)
	*/
	// SendWeChatStrategy(topic, sendmsg) // send to wechat
	// 消息入库
	editStrageydata(info)
	return true
}

func SendWeChatStrategy(room, msg string) {
	beego.Debug("send wechat aaaaaaaaaaaaa")
	arr, ok := wechat.MapUname[room]
	if ok {
		for _, v := range arr {
			wechat.SendTxTMsg(v, msg)
		}
	}
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
					editStrategyContent(infoMsg)
				}
			case infoOper, ok1 := <-n.operchan:
				if ok1 {
					OperateStrategyContent(infoOper)
				}
			}
		}
	}()
}

func editStrageydata(info StrategyInfo) {
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

func editStrategyContent(info *StrategyInfo) {
	beego.Debug("edit StrategyInfo", info)
	OPERATETYPE := info.OperType
	switch OPERATETYPE {
	case OPERATE_ADD:
		var strategy m.Strategy
		if info.FileName != "" {
			fileName := haoindex.GetWxServerImg(info.FileName)
			strategy.FileName = fileName
		} else {
			strategy.FileName = info.FileName
		}
		strategy.Room = info.Room
		strategy.Icon = info.Icon
		strategy.Name = info.Name
		strategy.Titel = info.Titel
		strategy.Data = info.Data
		strategy.TxtColour = info.TxtColour
		strategy.IsTop = info.IsTop
		strategy.IsDelete = info.IsDelete
		strategy.ThumbNum = info.ThumbNum
		strategy.Time = info.Time
		strategy.WxServerImgid = info.FileName
		strategy.Datatime = time.Now()
		_, err := m.AddStrategy(&strategy)
		if err != nil {
			beego.Debug("Add Strategy Fail:", err)
		}
		break
	case OPERATE_UPDATE:
		strategyInfo, err := m.GetStrategyInfoById(info.Id)
		if err != nil {
			beego.Debug("get Strategy id error", err)
		}

		var strategy m.Strategy
		strategy.Id = info.Id
		strategy.Room = info.Room
		strategy.Icon = info.Icon
		strategy.Name = info.Name
		strategy.Titel = info.Titel
		strategy.Data = info.Data
		if strategyInfo.WxServerImgid != info.FileName && info.FileName != "" {
			fileName := haoindex.GetWxServerImg(info.FileName)
			strategy.FileName = fileName
			strategy.WxServerImgid = info.FileName
		} else {
			strategy.FileName = strategyInfo.FileName
			strategy.WxServerImgid = strategyInfo.WxServerImgid
		}
		strategy.TxtColour = info.TxtColour
		strategy.IsTop = info.IsTop
		strategy.IsDelete = info.IsDelete
		strategy.ThumbNum = info.ThumbNum
		strategy.Time = info.Time
		strategy.Datatime = time.Now()
		_, err = m.UpdateStrategyById(&strategy)
		if err != nil {
			beego.Debug("Update Strategy Fail:", err)
		}
		break
	default:
	}

}
