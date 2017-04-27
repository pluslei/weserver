package haoadmin

import (
	"weserver/models"

	"github.com/astaxie/beego"
)

type SuggestController struct {
	CommonController
}

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
			roomInfo, err := models.GetRoomInfoByRoomID(item["Room"].(string))
			if err != nil {
				item["Room"] = "未知房间"
			} else {
				item["Room"] = roomInfo.RoomTitle
			}
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

func (this *SuggestController) Add() {
	action := this.GetString("action")
	if action == "add" {
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

func (this *SuggestController) Del() {
	id, _ := this.GetInt64("id")
	_, err := models.DelStrategyById(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
	} else {
		this.Rsp(true, "删除成功", "")
	}
}
