package haoadmin

import (
	"weserver/models"

	"time"

	"github.com/astaxie/beego"
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
		oper := new(models.OperPosition)
		oper.RoomId = this.GetString("RoomId")
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

		_, err := models.AddPosition(oper)
		if err != nil {
			this.AlertBack("添加失败")
			return
		}
		this.Alert("添加成功", "suggest_index")

	} else {
		this.CommonMenu()
		roonInfo, _, err := models.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		beego.Debug("roonInfo", roonInfo)
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/suggest/add.html"
	}
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
		oper := new(models.ClosePosition)
		oper.RoomId = this.GetString("RoomId")
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

		_, err = models.AddClosePosition(oper)
		if err != nil {
			this.AlertBack("添加失败")
			return
		}
		models.UpdatePositonLq(id)
		this.Alert("添加成功", "suggest_index")

	} else {
		this.CommonMenu()
		roonInfo, _, err := models.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
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
			this.Data["RoomId"] = "未知房间"
		} else {
			this.Data["RoomId"] = room.Title
		}
		this.TplName = "haoadmin/data/suggest/addclose.html"
	}
}

// 编辑建仓
func (this *SuggestController) Edit() {
	action := this.GetString("action")
	_, err := this.GetInt64("id")
	if err != nil {
		this.AlertBack("编辑失败")
		return
	}
	if action == "edit" {

	} else {
		this.CommonMenu()

		roonInfo, _, err := models.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/strategy/edit.html"
	}
}

// 删除建仓
func (this *SuggestController) Del() {
	id, _ := this.GetInt64("id")
	_, err := models.DelStrategyById(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
	} else {
		this.Rsp(true, "删除成功", "")
	}
}
