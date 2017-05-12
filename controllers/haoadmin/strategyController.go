package haoadmin

import (
	"time"
	"weserver/models"
	. "weserver/src/tools"

	mq "weserver/src/mqtt"

	"github.com/astaxie/beego"
)

type StrategyController struct {
	CommonController
}

func (this *StrategyController) Index() {
	if this.IsAjax() {
		user := this.GetSession("userinfo").(*models.User)
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")

		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		stratelist, count := models.GetStrategyInfoList(iStart, iLength, "-Id", user.CompanyId)
		for _, item := range stratelist {
			roomInfo, err := models.GetRoomInfoByRoomID(item["Room"].(string))
			if err != nil {
				item["Room"] = "未知房间"
			} else {
				item["Room"] = roomInfo.RoomTitle
			}
			Info, err := models.GetCompanyById(item["CompanyId"].(int64))
			if err != nil {
				item["CompanyName"] = "未知公司"
			} else {
				item["CompanyName"] = Info.Company
			}
			item["DatatimeStr"] = item["Datatime"].(time.Time).Format("2006-01-02 15:04:05")
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
	userInfo := this.GetSession("userinfo").(*models.User)
	if userInfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
		return
	}
	if action == "add" {
		roomId := this.GetStrings("RoomId")
		if len(roomId) == 0 {
			this.Alert("请选择房间号", "/weserver/data/strategy_index")
			return
		}
		for _, val := range roomId {

			strategy := new(models.Strategy)
			companyId, err := this.GetInt64("company")
			beego.Debug("companyId", companyId)
			if err != nil {
				beego.Error(err)
				this.Alert("获取公司id出错", "/weserver/data/strategy_index")
				return
			}
			strategy.CompanyId = companyId
			strategy.Room = val
			strategy.Icon = userInfo.Headimgurl
			strategy.Name = userInfo.Nickname
			userTitleInfo, err := models.ReadTitleById(userInfo.Title.Id)
			if err != nil {
				strategy.Titel = userInfo.Nickname
			} else {
				strategy.Titel = userTitleInfo.Name
			}
			strategy.Data = this.GetString("Data")
			strategy.FileName = this.GetString("FileNameFile")
			Top, _ := this.GetBool("Top")
			strategy.IsTop = Top
			ThumbNum, _ := this.GetInt64("ThumbNum")
			strategy.ThumbNum = ThumbNum
			strategy.Datatime = time.Now()
			strategy.TxtColour = this.GetString("TxtColour")
			strategy.Time = time.Now().Format("2006-01-02 15:04:05")
			_, err = models.AddStrategy(strategy)
			if err != nil {
				this.AlertBack("添加失败")
				continue
			}
			SendStrage(strategy)
			this.Alert("添加成功", "/weserver/data/strategy_index")
		}

	} else {
		this.CommonMenu()
		roonInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Error("get the roominfo error", err)
			return
		}
		companyList, _, err := models.GetCompanyList(userInfo.CompanyId)
		if err != nil {
			beego.Error("get the companyList error", err)
			return
		}
		this.Data["CompanyInfo"] = companyList
		this.Data["roonInfo"] = roonInfo
		this.TplName = "haoadmin/data/strategy/add.html"
	}
}

func SendStrage(info *models.Strategy) {
	msg := new(StrategyInfo)
	msg.CompanyId = info.CompanyId
	msg.Room = info.Room
	msg.Icon = info.Icon
	msg.Name = info.Name
	msg.Titel = info.Titel
	msg.Data = info.Data
	msg.FileName = info.FileName
	msg.TxtColour = info.TxtColour
	msg.IsTop = info.IsTop
	msg.IsDelete = info.IsDelete
	msg.ThumbNum = info.ThumbNum
	msg.Time = info.Time
	msg.MsgType = MSG_TYPE_STRATEGY_ADD

	topic := msg.Room

	beego.Debug("msginfo", msg)

	v, err := ToJSON(msg)
	if err != nil {
		beego.Error(" Suggest add position json error", err)
		return
	}
	mq.SendMessage(topic, v)
}

func (this *StrategyController) Edit() {
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		this.AlertBack("编辑失败")
		return
	}
	if action == "edit" {
		strategy := make(map[string]interface{})
		strategy["Room"] = this.GetString("Room")
		strategy["Data"] = this.GetString("Data")
		strategy["TxtColour"] = this.GetString("TxtColour")
		strategy["FileName"] = this.GetString("FileNameFile")
		_, err = models.UpdateStrategy(id, strategy)
		if err != nil {
			this.AlertBack("修改失败")
		}
		this.Alert("修改成功", "strategy_index")
	} else {
		this.CommonMenu()
		info, err := models.GetStrategyInfoById(id)
		if err != nil {
			this.AlertBack("编辑失败")
			return
		}
		beego.Debug("info", info)
		this.Data["Info"] = info

		roomInfo, err := this.GetRoomInfo()
		if err != nil {
			beego.Debug("Get RoomInfo error", err)
			return
		}
		this.Data["roonInfo"] = roomInfo
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
