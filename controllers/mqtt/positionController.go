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

type PositionController struct {
	controllers.PublicController
}

type positionMessage struct {
	infochan chan *PositionInfo
}

var (
	pos *positionMessage
)

func init() {
	pos = &positionMessage{
		infochan: make(chan *PositionInfo, 20480),
	}
	pos.runWriteDb()
}

//Add Position
func (this *PositionController) OperatePosition() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parsePositionMsg(msg)
		if b {
			this.Rsp(true, "操作信息发送成功", "")
			return
		} else {
			this.Rsp(false, "操作信息发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

//All Position List
func (this *PositionController) GetAllPositionList() {
	if this.IsAjax() {
		roomId := this.GetString("room")
		data := make(map[string]interface{})
		historyPosition, _, err := m.GetAllPositionList(roomId)
		if err != nil {
			beego.Debug("Get All Position List error", err)
			this.Rsp(false, "Get All Position List Error", "")
			return
		}
		data["historyPosition"] = historyPosition
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

func (this *PositionController) GetPositionNearRecord() {
	if this.IsAjax() {
		roomId := this.GetString("room")
		data := make(map[string]interface{})
		historyNear, err := m.GetNearRecord(roomId)
		if err != nil {
			beego.Debug("Get All Position List error", err)
			this.Rsp(false, "Get All Position List Error", "")
			return
		}
		data["historyNear"] = historyNear
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

func (this *PositionController) GetClosePositionRecord() {
	if this.IsAjax() {
		strId := this.GetString("Id")
		nId, _ := strconv.ParseInt(strId, 10, 64)
		data := make(map[string]interface{})
		historyClose, _, err := m.GetMoreClosePosition(nId)
		if err != nil {
			beego.Debug("Get All Position List error", err)
			this.Rsp(false, "Get All Position List Error", "")
			return
		}
		data["historyClose"] = historyClose
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

//Position List
func (this *PositionController) GetPositionList() {
	if this.IsAjax() {
		strId := this.GetString("Id")
		beego.Debug("id", strId)
		nId, _ := strconv.ParseInt(strId, 10, 64)
		roomId := this.GetString("room")
		beego.Debug("Position list ", nId, roomId)
		data := make(map[string]interface{})
		sysconfig, _ := m.GetAllSysConfig()
		sysCount := sysconfig.PositionCount
		var positionInfo []m.OperPosition
		historyPosition, totalCount, err := m.GetAllPositionList(roomId)
		if err != nil {
			beego.Debug("Get Position List error:", err)
			this.Rsp(false, "Get Position List error", "")
			return
		}
		if nId == 0 {
			var i int64
			if totalCount < sysCount {
				beego.Debug("nCount sysCont", totalCount, sysCount)
				for i = 0; i < totalCount; i++ {
					var info m.OperPosition
					info.Id = historyPosition[i].Id
					info.RoomId = historyPosition[i].RoomId
					info.RoomTeacher = historyPosition[i].RoomTeacher
					info.Type = historyPosition[i].Type               //种类
					info.BuySell = historyPosition[i].BuySell         //买卖 0 1
					info.Entrust = historyPosition[i].Entrust         //委托类型
					info.Index = historyPosition[i].Index             //点位
					info.Position = historyPosition[i].Position       //仓位
					info.ProfitPoint = historyPosition[i].ProfitPoint //止盈点
					info.LossPoint = historyPosition[i].LossPoint     //止损点
					info.Notes = historyPosition[i].Notes             // 备注
					info.Timestr = historyPosition[i].Timestr
					info.Icon = historyPosition[i].Icon
					info.Liquidation = historyPosition[i].Liquidation //平仓详情 (0:未平仓 1:平仓)
					if info.Liquidation == 1 {
						historyClose, _, err := m.GetMoreClosePosition(info.Id)
						if err != nil {
							beego.Debug("Get historyClosePosition info error", err)
						}
						info.CloseType = historyClose[0].Type    //平仓种类
						info.CloseIndex = historyClose[0].Index  //平仓点位
						info.CloseNotes = historyClose[0].Notes  //平仓备注
						info.CloseTime = historyClose[0].Timestr //平仓时间
					}
					positionInfo = append(positionInfo, info)
				}
			} else {
				for i = 0; i < sysCount; i++ {
					var info m.OperPosition
					info.Id = historyPosition[i].Id
					info.RoomId = historyPosition[i].RoomId
					info.RoomTeacher = historyPosition[i].RoomTeacher
					info.Type = historyPosition[i].Type               //种类
					info.BuySell = historyPosition[i].BuySell         //买卖 0 1
					info.Entrust = historyPosition[i].Entrust         //委托类型
					info.Index = historyPosition[i].Index             //点位
					info.Position = historyPosition[i].Position       //仓位
					info.ProfitPoint = historyPosition[i].ProfitPoint //止盈点
					info.LossPoint = historyPosition[i].LossPoint     //止损点
					info.Notes = historyPosition[i].Notes             // 备注
					info.Liquidation = historyPosition[i].Liquidation //平仓详情 (0:未平仓 1:平仓)
					info.Timestr = historyPosition[i].Timestr
					info.Icon = historyPosition[i].Icon
					if info.Liquidation == 1 {
						historyClose, _, err := m.GetMoreClosePosition(info.Id)
						if err != nil {
							beego.Debug("Get historyClosePosition info error", err)
						}
						info.CloseType = historyClose[0].Type    //平仓种类
						info.CloseIndex = historyClose[0].Index  //平仓点位
						info.CloseNotes = historyClose[0].Notes  //平仓备注
						info.CloseTime = historyClose[0].Timestr //平仓时间
					}
					positionInfo = append(positionInfo, info)
				}
			}
			data["historyPosition"] = positionInfo
			this.Data["json"] = &data
			this.ServeJSON()
		} else {
			var index int64
			for nindex, value := range historyPosition {
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
				data["historyPosition"] = ""
				this.Data["json"] = &data
				this.ServeJSON()
				return
			}
			if nCount < totalCount {
				for i := index; i < nCount; i++ {
					var info m.OperPosition
					info.Id = historyPosition[i].Id
					info.RoomId = historyPosition[i].RoomId
					info.RoomTeacher = historyPosition[i].RoomTeacher
					info.Type = historyPosition[i].Type               //种类
					info.BuySell = historyPosition[i].BuySell         //买卖 0 1
					info.Entrust = historyPosition[i].Entrust         //委托类型
					info.Index = historyPosition[i].Index             //点位
					info.Position = historyPosition[i].Position       //仓位
					info.ProfitPoint = historyPosition[i].ProfitPoint //止盈点
					info.LossPoint = historyPosition[i].LossPoint     //止损点
					info.Notes = historyPosition[i].Notes             // 备注
					info.Liquidation = historyPosition[i].Liquidation //平仓详情 (0:未平仓 1:平仓)
					info.Timestr = historyPosition[i].Timestr
					info.Icon = historyPosition[i].Icon
					if info.Liquidation == 1 {
						historyClose, _, err := m.GetMoreClosePosition(info.Id)
						if err != nil {
							beego.Debug("Get historyClosePosition info error", err)
						}
						info.CloseType = historyClose[0].Type    //平仓种类
						info.CloseIndex = historyClose[0].Index  //平仓点位
						info.CloseNotes = historyClose[0].Notes  //平仓备注
						info.CloseTime = historyClose[0].Timestr //平仓时间
					}

					positionInfo = append(positionInfo, info)
				}
			} else {
				for i := index; i < totalCount; i++ {
					var info m.OperPosition
					info.Id = historyPosition[i].Id
					info.RoomId = historyPosition[i].RoomId
					info.RoomTeacher = historyPosition[i].RoomTeacher
					info.Type = historyPosition[i].Type               //种类
					info.BuySell = historyPosition[i].BuySell         //买卖 0 1
					info.Entrust = historyPosition[i].Entrust         //委托类型
					info.Index = historyPosition[i].Index             //点位
					info.Position = historyPosition[i].Position       //仓位
					info.ProfitPoint = historyPosition[i].ProfitPoint //止盈点
					info.LossPoint = historyPosition[i].LossPoint     //止损点
					info.Notes = historyPosition[i].Notes             // 备注
					info.Liquidation = historyPosition[i].Liquidation //平仓详情 (0:未平仓 1:平仓)
					info.Timestr = historyPosition[i].Timestr
					info.Icon = historyPosition[i].Icon
					if info.Liquidation == 1 {
						historyClose, _, err := m.GetMoreClosePosition(info.Id)
						if err != nil {
							beego.Debug("Get historyClosePosition info error", err)
						}
						info.CloseType = historyClose[0].Type    //平仓种类
						info.CloseIndex = historyClose[0].Index  //平仓点位
						info.CloseNotes = historyClose[0].Notes  //平仓备注
						info.CloseTime = historyClose[0].Timestr //平仓时间
					}

					positionInfo = append(positionInfo, info)
				}
			}
			data["historyPosition"] = positionInfo
			this.Data["json"] = &data
			this.ServeJSON()
		}
	} else {
		this.Ctx.Redirect(302, "/")
	}
	this.Ctx.WriteString("")
}

func parsePositionMsg(msg string) bool {
	msginfo := new(PositionInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("PositionInfo: simplejson error", err)
		return false
	}
	if info.OperType == OPERATE_POSITION_ADD {
		info.MsgType = MSG_TYPE_POSITION_ADD
		topic := info.RoomId

		beego.Debug("info", info)

		v, err := ToJSON(info)
		if err != nil {
			beego.Error("json error", err)
			return false
		}
		mq.SendMessage(topic, v)
	}

	// 消息入库
	operatePositiondata(info)
	return true
}

// 写数据
func (n *positionMessage) runWriteDb() {
	go func() {
		for {
			select {
			case infoMsg, ok := <-n.infochan:
				if ok {
					operatePosData(infoMsg)
				}
			}
		}
	}()
}

func operatePositiondata(info PositionInfo) {
	jsondata := &info
	select {
	case pos.infochan <- jsondata:
		break
	default:
		beego.Error("WRITE PositionInfo db error!!!")
		break
	}
}

func operatePosData(info *PositionInfo) {
	beego.Debug("PositionOperate", info)
	var op m.OperPosition
	op.Id = info.Id
	op.RoomId = info.RoomId
	OPERTYPE := info.OperType
	switch OPERTYPE {
	case OPERATE_POSITION_ADD:
		if op.Id == 0 {
			id, err := addPositionConten(info)
			if err != nil {
				beego.Debug("Oper Position Add Fail", err)
				return
			}
			//增加平仓信息
			if info.CloseIndex != "" {
				err := addClosePositionData(info, id)
				if err != nil {
					beego.Debug("Oper ClosePosition Add Fail", err)
					return
				}
			}
		}
		break
	case OPERATE_POSITION_DEL:
		_, err := m.DelPositionById(op.Id)
		if err != nil {
			beego.Debug("Oper Position Del Fail", err)
			return
		}
		_, errClose := m.DelClosePositionByOperId(op.Id)
		if errClose != nil {
			beego.Debug("Oper ClosePosition Del Fail", err)
			return
		}
		break
	case OPERATE_POSITION_UPDATE:
		err := updatePositionConten(info)
		if err != nil {
			beego.Debug("Oper Position update Fail", err)
			return
		}
		errClose := updateClosePositionData(info)
		if errClose != nil {
			beego.Debug("Oper ClosePosition Update Fail", errClose)
			return
		}
		break
	default:
	}
}

func addPositionConten(info *PositionInfo) (int64, error) {
	beego.Debug("Add PositionInfo", info)
	var op m.OperPosition
	op.RoomId = info.RoomId
	op.RoomTeacher = info.RoomTeacher
	op.Type = info.Type
	op.BuySell = info.BuySell
	op.Entrust = info.Entrust
	op.Index = info.Index
	op.Position = info.Position
	op.ProfitPoint = info.ProfitPoint
	op.LossPoint = info.LossPoint
	op.Notes = info.Notes
	op.Liquidation = info.Liquidation
	op.Icon = info.Icon
	op.Time = time.Now()
	op.Timestr = op.Time.Format("2006-01-02 03:04:05")
	id, err := m.AddPosition(&op)
	if err != nil {
		beego.Debug("Add Teacher Fail:", err)
		return 0, err
	}
	return id, nil
}

func updatePositionConten(info *PositionInfo) error {
	beego.Debug("Update Position Info", info)
	var pos m.OperPosition
	pos.RoomId = info.RoomId
	pos.RoomTeacher = info.RoomTeacher
	pos.Type = info.Type
	pos.BuySell = info.BuySell
	pos.Entrust = info.Entrust
	pos.Index = info.Index
	pos.Position = info.Position
	pos.ProfitPoint = info.ProfitPoint
	pos.LossPoint = info.LossPoint
	pos.Notes = info.Notes
	pos.Liquidation = info.Liquidation
	pos.Icon = info.Icon
	pos.Time = time.Now()
	pos.Timestr = pos.Time.Format("2006-01-02 03:04:05")

	_, err := m.UpdatePositionInfo(&pos)
	if err != nil {
		beego.Debug("Add Teacher Fail:", err)
		return err
	}
	return nil
}

func addClosePositionData(info *PositionInfo, Id int64) error {
	beego.Debug("Add ClosePosition Info", info)
	var pos m.ClosePosition
	pos.RoomId = info.RoomId
	pos.RoomTeacher = info.RoomTeacher
	pos.Type = info.Type
	pos.BuySell = info.CloseBuySell
	pos.Entrust = info.Entrust
	pos.Index = info.CloseIndex
	pos.Position = "--"
	pos.ProfitPoint = "--"
	pos.LossPoint = "--"
	pos.Notes = info.CloseNotes
	pos.Time = time.Now()
	pos.Timestr = pos.Time.Format("2006-01-02 03:04:05")
	pos.OperPosition = &m.OperPosition{Id: Id}
	_, err := m.AddClosePosition(&pos)
	if err != nil {
		beego.Debug("Add ClosePosition Fail:", err)
		return err
	}
	return nil
}

func updateClosePositionData(info *PositionInfo) error {
	beego.Debug("Update ClosePosition Info", info, info.Id)
	close, err := m.GetIdByOperPositionId(info.Id)
	if err != nil {
		addClosePositionData(info, info.Id)
	} else {
		var pos m.ClosePosition
		pos.RoomId = info.RoomId
		pos.RoomTeacher = info.RoomTeacher
		pos.Type = info.Type
		pos.BuySell = info.CloseBuySell
		pos.Entrust = info.Entrust
		pos.Index = info.CloseIndex
		pos.Position = "--"
		pos.ProfitPoint = "--"
		pos.LossPoint = "--"
		pos.Notes = info.CloseNotes
		pos.Time = time.Now()
		pos.Timestr = pos.Time.Format("2006-01-02 03:04:05")
		_, err = m.UpdateClosePositionInfo(close.Id, &pos)
		if err != nil {
			beego.Debug("Add ClosePosition Fail:", err)
			return err
		}
	}
	return nil
}
