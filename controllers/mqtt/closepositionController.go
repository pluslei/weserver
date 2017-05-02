package mqtt

import (
	"time"
	m "weserver/models"

	"github.com/astaxie/beego"

	"weserver/controllers"
	. "weserver/src/tools"
	// for json get
)

type ClosepositionController struct {
	controllers.PublicController
}

type closepositionMessage struct {
	infochan chan *ClosePositionInfo
}

var (
	closepos *closepositionMessage
)

func init() {
	closepos = &closepositionMessage{
		infochan: make(chan *ClosePositionInfo, 20480),
	}
	closepos.runWriteDb()
}

// ClosePosition
func (this *PositionController) OperateClosePosition() {
	if this.IsAjax() {
		msg := this.GetString("str")
		b := parseClosePositionMsg(msg)
		if b {
			this.Rsp(true, "平仓信息发送成功", "")
			return
		} else {
			this.Rsp(false, "平仓信息发送失败,请重新发送", "")
			return
		}
	}
	this.Ctx.WriteString("")
}

func parseClosePositionMsg(msg string) bool {
	msginfo := new(ClosePositionInfo)
	info, err := msginfo.ParseJSON(DecodeBase64Byte(msg))
	if err != nil {
		beego.Error("ClosePositionInfo: simplejson error", err)
		return false
	}
	/*
		info.OperType = OPERATE_CLOSEPOSITION_ADD
			info.MsgType = MSG_TYPE_CLOSEPOSITION_ADD
			topic := info.RoomId

			beego.Debug("info", info)

			v, err := ToJSON(info)
			if err != nil {
				beego.Error("json error", err)
				return false
			}
			mq.SendMessage(topic, v)
			// 消息入库
	*/
	operateClosePositiondata(info)
	return true
}

// 写数据
func (n *closepositionMessage) runWriteDb() {
	go func() {
		for {
			select {
			case infoMsg, ok := <-n.infochan:
				if ok {
					operateClosePosData(infoMsg)
				}
			}
		}
	}()
}

func operateClosePositiondata(info ClosePositionInfo) {
	jsondata := &info
	select {
	case closepos.infochan <- jsondata:
		break
	default:
		beego.Error("WRITE PositionInfo db error!!!")
		break
	}
}

func operateClosePosData(info *ClosePositionInfo) {
	beego.Debug("ClosePositionInfo", info)
	var op m.OperPosition
	op.Id = info.Id
	op.RoomId = info.RoomId
	OPERTYPE := info.OperType
	switch OPERTYPE {
	case OPERATE_CLOSEPOSITION_ADD:
		if op.Id == 0 {
			err := addClosePositionConten(info)
			if err != nil {
				beego.Debug("Oper ClosePosition Add Fail", err)
				return
			}
		}
		break
	case OPERATE_CLOSEPOSITION_DEL:
		_, err := m.DelClosePositionById(op.Id)
		if err != nil {
			beego.Debug("Oper ClosePosition Del Fail", err)
			return
		}
		break
	case OPERATE_CLOSEPOSITION_UPDATE:
		err := updateClosePositionConten(info)
		if err != nil {
			beego.Debug("Oper ClosePosition update Fail", err)
			return
		}
		break
	default:
	}
}

func addClosePositionConten(info *ClosePositionInfo) error {
	beego.Debug("Add ClosePosition Info", info)
	var pos m.ClosePosition
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
	pos.Time = time.Now()
	pos.Timestr = pos.Time.Format("2006-01-02 03:04:05")
	pos.OperPosition = &m.OperPosition{Id: info.Id}
	_, err := m.AddClosePosition(&pos)
	if err != nil {
		beego.Debug("Add ClosePosition Fail:", err)
		return err
	}
	return nil
}

func updateClosePositionConten(info *ClosePositionInfo) error {
	beego.Debug("Update ClosePosition Info", info, info.Id)
	close, err := m.GetIdByOperPositionId(info.Id)
	if err != nil {
		addClosePositionConten(info)
	} else {
		var pos m.ClosePosition
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
