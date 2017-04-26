package haoadmin

import (
	"time"
	"weserver/models"

	"github.com/astaxie/beego"
)

type StrategyController struct {
	CommonController
}

func (this *StrategyController) Index() {
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
		stratelist, count := models.GetStrategyInfoList(iStart, iLength, "-Id")
		for _, item := range stratelist {
			roomInfo, err := models.GetRoomInfoByRoomID(item["Room"].(string))
			if err != nil {
				item["Room"] = "未知房间"
			} else {
				item["Room"] = roomInfo.RoomTitle
			}
		}

		// json
		data := make(map[string]interface{})
		data["aaData"] = stratelist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/strategy/list.html"
	}
}

func (this *StrategyController) Add() {
	action := this.GetString("action")
	if action == "add" {
		userInfo := this.GetSession("userinfo").(*models.User)
		beego.Debug("userInfo", userInfo)
		strategy := new(models.Strategy)
		strategy.Room = this.GetString("Room")
		strategy.Icon = userInfo.Headimgurl
		strategy.Name = userInfo.Nickname
		strategy.Titel = userInfo.Title.Name
		strategy.Data = this.GetString("Data")
		strategy.FileName = this.GetString("FileNameFile")
		Top, _ := this.GetBool("Top")
		strategy.IsTop = Top
		ThumbNum, _ := this.GetInt64("ThumbNum")
		strategy.ThumbNum = ThumbNum
		strategy.Datatime = time.Now()

		_, err := models.AddStrategy(strategy)
		if err != nil {
			this.AlertBack("添加失败")
		}
		this.Alert("添加成功", "/weserver/data/strategy_index")
	} else {
		this.CommonMenu()
		roonInfo, _, err := models.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		beego.Debug("roonInfo", roonInfo)
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/strategy/add.html"
	}
}

func (this *StrategyController) Edit() {
	action := this.GetString("action")
	if action == "edit" {

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/strategy/edit.html"
	}
}

func (this *StrategyController) Del() {
	id, _ := this.GetInt64("id")
	_, err := models.DelStrategyById(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
	} else {
		this.Rsp(true, "删除成功", "")
	}
}
