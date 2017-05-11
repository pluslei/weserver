package haoadmin

import (
	"time"
	"weserver/models"
	. "weserver/src/tools"

	"github.com/astaxie/beego"
	// "github.com/denverdino/aliyungo/mq"

	mq "weserver/src/mqtt"
)

type SuggestController struct {
	CommonController
}

// 建仓列表
func (this *SuggestController) Index() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")

		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		operposition, count := models.GetOperPositionList(iStart, iLength, "-Id")
		for _, item := range operposition {
			roomInfo, err := models.GetRoomInfoByRoomID(item["RoomId"].(string))
			if err != nil {
				item["RoomId"] = "未知房间"
			} else {
				item["RoomId"] = roomInfo.RoomTitle
			}
			item["Timestr"] = item["Time"].(time.Time).Format("2006-01-02 15:04:05")
		}

		// json
		data := make(map[string]interface{})
		data["aaData"] = operposition
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/suggest/list.html"
	}
}

// 添加建仓
func (this *SuggestController) Add() {
	action := this.GetString("action")
	if action == "add" {
		userInfo := this.GetSession("userinfo").(*models.User)
		if userInfo == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		beego.Debug("userInfo", userInfo.Title.Id)
		titleInfo, err := models.ReadTitleById(userInfo.Title.Id)
		if err != nil {
			beego.Debug("title info error", err)
			return
		}
		beego.Debug("titleInfo", titleInfo)

		oper := new(models.OperPosition)
		oper.RoomId = this.GetString("RoomId")

		if userInfo.CompanyId == 0 {
			id, err := models.GetRoomCompany(oper.RoomId)
			if err != nil {
				beego.Debug("Get Room Company Error", err)
				return
			}
			oper.CompanyId = id
		} else {
			oper.CompanyId = userInfo.CompanyId
		}
		oper.RoomTeacher = titleInfo.Name
		oper.Time = time.Now()
		oper.Type = this.GetString("Type")
		BuySell, _ := this.GetInt("BuySell")
		oper.BuySell = BuySell
		oper.Entrust = this.GetString("Entrust")
		oper.Index = this.GetString("Index")
		oper.Position = this.GetString("Position")
		oper.ProfitPoint = this.GetString("ProfitPoint")
		oper.LossPoint = this.GetString("LossPoint")
		oper.Notes = this.GetString("Notes")
		time := time.Now()
		tm := time.Format("2006-01-02 15:04:05")
		oper.Timestr = tm

		_, err = models.AddPosition(oper)
		if err != nil {
			this.AlertBack("添加失败")
			return
		}
		msg := new(PositionInfo)
		msg.CompanyId = oper.CompanyId
		msg.RoomId = oper.RoomId
		msg.RoomTeacher = oper.RoomTeacher
		msg.Type = oper.Type
		msg.BuySell = oper.BuySell
		msg.Entrust = oper.Entrust
		msg.Index = oper.Index
		msg.Position = oper.Position
		msg.ProfitPoint = oper.ProfitPoint
		msg.LossPoint = oper.LossPoint
		msg.Notes = oper.Notes
		msg.Liquidation = 0
		msg.MsgType = MSG_TYPE_POSITION_ADD
		topic := msg.RoomId

		beego.Debug("msginfo", msg)

		v, err := ToJSON(msg)
		if err != nil {
			beego.Error(" Suggest add position json error", err)
			return
		}
		mq.SendMessage(topic, v)
		this.Alert("添加成功", "suggest_index")

	} else {
		this.CommonMenu()
		roonInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/suggest/add.html"
	}
}

// 编辑建仓
func (this *SuggestController) Edit() {
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		this.AlertBack("编辑失败")
		return
	}
	if action == "edit" {
		oper := make(map[string]interface{})
		oper["RoomId"] = this.GetString("RoomId")
		oper["RoomTeacher"] = this.GetString("RoomTeacher")
		oper["Time"] = time.Now()
		oper["Type"] = this.GetString("Type")
		BuySell, _ := this.GetInt("BuySell")
		oper["BuySell"] = BuySell
		oper["Entrust"] = this.GetString("Entrust")
		oper["Index"] = this.GetString("Index")
		oper["Position"] = this.GetString("Position")
		oper["ProfitPoint"] = this.GetString("ProfitPoint")
		oper["LossPoint"] = this.GetString("LossPoint")
		oper["Notes"] = this.GetString("Notes")

		_, err := models.UpdatePosition(id, oper)
		if err != nil {
			this.AlertBack("修改失败")
			return
		}
		this.Alert("修改成功", "suggest_index")
	} else {
		this.CommonMenu()
		roomInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		operInfo, err := models.GetOpersitionInfoById(id)
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		roomteacher, err := models.GetTeacherListByRoom(operInfo.RoomId)
		if err != nil {
			beego.Error("error", err)
			return
		}
		this.Data["roomteacherInfo"] = roomteacher
		this.Data["operInfo"] = operInfo
		this.Data["roonInfo"] = roomInfo
		this.TplName = "haoadmin/data/suggest/edit.html"
	}
}

// 删除建仓
func (this *SuggestController) Del() {
	id, _ := this.GetInt64("id")
	_, err := models.DelPositionById(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
	} else {
		this.Rsp(true, "删除成功", "")
	}
	models.DelClosePositionByOperId(id)
}

// 添加平仓
func (this *SuggestController) AddClose() {
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		this.AlertBack("数据查找失败")
		return
	}
	if action == "add" {
		userInfo := this.GetSession("userinfo").(*models.User)
		if userInfo == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		oper := new(models.ClosePosition)
		oper.RoomId = this.GetString("RoomId")
		if userInfo.CompanyId == 0 {
			id, err := models.GetRoomCompany(oper.RoomId)
			if err != nil {
				beego.Debug("Get Room Company Error", err)
				return
			}
			oper.CompanyId = id
		} else {
			oper.CompanyId = userInfo.CompanyId
		}
		oper.RoomTeacher = this.GetString("RoomTeacher")
		oper.Time = time.Now()
		oper.Type = this.GetString("Type")
		BuySell, _ := this.GetInt("BuySell")
		oper.BuySell = BuySell
		oper.Entrust = this.GetString("Entrust")
		oper.Index = this.GetString("Index")
		oper.Position = this.GetString("Position")
		oper.ProfitPoint = this.GetString("ProfitPoint")
		oper.LossPoint = this.GetString("LossPoint")
		oper.Notes = this.GetString("Notes")
		oper.OperPosition = &models.OperPosition{Id: id}
		time := time.Now()
		tm := time.Format("2006-01-02 15:04:05")
		oper.Timestr = tm

		_, err = models.AddClosePosition(oper)
		if err != nil {
			this.AlertBack("添加失败")
			return
		}
		models.UpdatePositonLq(id)
		this.Alert("添加成功", "suggest_index")

	} else {
		this.CommonMenu()
		roonInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Error("Get the Roominfo error", err)
			return
		}
		this.Data["roonInfo"] = roonInfo

		operInfo, err := models.GetOpersitionInfoById(id)
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		this.Data["timestr"] = operInfo.Time.Format("2006-01-02 15:04:05")
		this.Data["operInfo"] = operInfo
		room, err := models.GetRoomInfoByRoomID(operInfo.RoomId)
		if err != nil {
			this.Data["Room"] = "未知房间"
		} else {
			this.Data["Room"] = room.Title
		}
		this.TplName = "haoadmin/data/suggest/addclose.html"
	}
}

// 获取平仓
func (this *SuggestController) GetClose() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Error("error", err)
		return
	}
	closeOper, _, err := models.GetMoreClosePosition(id)
	if err != nil {
		beego.Error("error", err)
		return
	}
	for _, item := range closeOper {
		roomInfo, err := models.GetRoomInfoByRoomID(item.RoomId)
		if err != nil {
			item.RoomId = "未知房间"
		} else {
			item.RoomId = roomInfo.RoomTitle
		}
		item.Timestr = item.Time.Format("2006-01-02 15:04:05")
	}
	this.Data["json"] = closeOper
	this.ServeJSON()
}

// 编辑平仓
func (this *SuggestController) EditClose() {
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		this.AlertBack("编辑失败")
		return
	}
	if action == "edit" {
		close := make(map[string]interface{})
		close["Index"] = this.GetString("Index")
		close["Notes"] = this.GetString("Notes")

		_, err = models.UpdateClosePosition(id, close)
		if err != nil {
			this.AlertBack("修改失败")
			return
		}
		models.UpdatePositonLq(id)
		this.Alert("修改成功", "suggest_index")
	} else {
		this.CommonMenu()
		closeInfo, err := models.GetClosePositionInfo(id)
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		this.Data["closeInfo"] = closeInfo

		this.TplName = "haoadmin/data/suggest/addclose.html"
	}
}

// 删除平仓
func (this *SuggestController) DelClose() {
	id, err := this.GetInt64("id")
	if err != nil {
		this.Rsp(false, "删除失败", "")
	} else {
		info, err := models.GetClosePositionInfo(id)
		if err != nil {
			this.Rsp(false, "获取建仓id失败", "")
		}
		id := info.OperPosition.Id
		_, err = models.DelClosePositionById(id)
		if err != nil {
			this.Rsp(false, "删除失败", "")
		}
		err1 := models.UpdatePositonUnLq(id)
		if err1 != nil {
			this.Rsp(false, "设置平仓信息失败", "")
		}

		this.Rsp(true, "删除成功", "")
	}
}
