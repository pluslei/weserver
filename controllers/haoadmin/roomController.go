package haoadmin

import (
	"fmt"
	"os"
	"path"
	"time"
	"weserver/models"
	"weserver/src/tools"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RoomController struct {
	CommonController
}

func (this *RoomController) Index() {
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
		SearchId := this.GetString("sSearch_0")

		companyId := user.CompanyId
		roolist, count := models.GetRoomInfoList(iStart, iLength, companyId, SearchId)

		for _, item := range roolist {
			Info, err := models.GetCompanyById(item["CompanyId"].(int64))
			if err != nil {
				item["CompanyName"] = "未知公司"
			} else {
				item["CompanyName"] = Info.Company
			}
		}
		// json
		data := make(map[string]interface{})
		data["aaData"] = roolist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/room/list.html"
	}
}

func (this *RoomController) Add() {
	action := this.GetString("action")
	UserInfo := this.GetSession("userinfo").(*models.User)
	if UserInfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
		return
	}
	if action == "add" {
		roomInfo, err := models.GetRoomInfoOne()
		if err != nil {
			this.AlertBack("请先添加一个房间")
			return
		}
		room := new(models.RoomInfo)
		room.RoomId = beego.AppConfig.String("mqServerTopic") + "/" + getRoomId()

		room.RoomTitle = this.GetString("RoomTitle")
		room.RoomTeacher = this.GetString("RoomTeacher")
		room.RoomNum = this.GetString("RoomNum")
		companyId, err := this.GetInt64("company")
		if err != nil {
			beego.Error(err)
			return
		}
		room.CompanyId = companyId
		if this.GetString("GroupId") == "" {
			room.GroupId = roomInfo.GroupId
		} else {
			room.GroupId = this.GetString("GroupId")
		}
		if this.GetString("Url") == "" {
			room.Url = roomInfo.Url
		} else {
			room.Url = this.GetString("Url")
		}
		port, err := this.GetInt("Port")
		if err != nil {
			room.Port = roomInfo.Port
		} else {
			room.Port = port
		}
		tls, err := this.GetBool("Tls")
		if err != nil {
			room.Tls = roomInfo.Tls
		} else {
			room.Tls = tls
		}
		if this.GetString("Access") == "" {
			room.Access = roomInfo.Access
		} else {
			room.Access = this.GetString("Access")
		}
		if this.GetString("SecretKey") == "" {
			room.SecretKey = roomInfo.SecretKey
		} else {
			room.SecretKey = this.GetString("SecretKey")
		}
		room.RoomIcon = this.GetString("RoomIcoFile")
		room.RoomIntro = this.GetString("RoomIntro")
		room.RoomBanner = this.GetString("RoomBannerFile")
		room.PcBanner = this.GetString("RoomBannerPCFile")
		room.PcRoomText = this.GetString("PcRoomText")
		room.PcRoomad = this.GetString("PcRoomadFile")
		room.Title = this.GetString("Title")
		_, err = models.AddRoom(room)
		if err != nil {
			this.AlertBack("插入失败")
			return
		}
		this.Alert("房间添加成功", "/weserver/data/room_index")
	} else {
		this.CommonController.CommonMenu()

		companyInfo, _, err := models.GetCompanyList(UserInfo.CompanyId)
		if err != nil {
			beego.Debug("get companylist error", err)
			return
		}
		this.Data["CompanyInfo"] = companyInfo
		this.TplName = "haoadmin/data/room/add.html"
	}
}

func (this *RoomController) Edit() {
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Debug("get id error", err)
		this.AlertBack("获取房间信息失败")
		return
	}
	if action == "edit" {
		var room = make(orm.Params)
		room["RoomTitle"] = this.GetString("RoomTitle")
		room["RoomTeacher"] = this.GetString("RoomTeacher")
		room["RoomNum"] = this.GetString("RoomNum")
		if this.GetString("GroupId") != "" {
			room["GroupId"] = this.GetString("GroupId")
		}
		if this.GetString("Url") != "" {
			room["Url"] = this.GetString("Url")
		}
		port, err := this.GetInt("Port")
		if err == nil {
			room["Port"] = port
		}
		tls, err := this.GetBool("Tls")
		if err == nil {
			room["Tls"] = tls
		}
		if this.GetString("Access") != "" {
			room["Access"] = this.GetString("Access")
		}
		if this.GetString("SecretKey") != "" {
			room["SecretKey"] = this.GetString("SecretKey")
		}
		room["RoomIcon"] = this.GetString("RoomIcoFile")
		room["RoomIntro"] = this.GetString("RoomIntro")
		room["RoomBanner"] = this.GetString("RoomBannerFile")
		room["PcBanner"] = this.GetString("RoomBannerPCFile")
		room["PcRoomad"] = this.GetString("PcRoomadFile")
		room["PcRoomText"] = this.GetString("PcRoomText")
		room["Title"] = this.GetString("Title")
		_, err = models.UpdateRoomInfo(id, room)
		if err != nil {
			beego.Error("inser faild", err)
			this.AlertBack("修改失败")
			return
		} else {
			this.Alert("修改成功", "/weserver/data/room_index")
		}
	}
	roomInfo, _ := models.GetRoomInfoById(id)
	this.CommonMenu()
	this.Data["roomInfo"] = roomInfo
	this.TplName = "haoadmin/data/room/edit.html"
}

func (this *RoomController) Del() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Debug("get id error", err)
		this.Rsp(false, "获取失败", "")
		return
	}
	_, err = models.DelRoomInfoId(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
	}
	this.Rsp(true, "删除成功", "")
}

func (this *RoomController) Upload() {
	uploadtype := this.GetString("uploadtype")

	_, h, err := this.GetFile("Filedata")
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
	err = this.SaveToFile("Filedata", dirPath)
	if err != nil {
		beego.Error(err)
		this.Rsp(false, uploadtype, "")
	}
	filepath := path.Join("/upload", "room", FileName)
	this.Rsp(true, uploadtype, filepath)
}

func getRoomId() string {
	random := tools.RandomAlphanumeric(6)
	if models.IsRoomInfo(random) {
		getRoomId()
	}
	return random
}

//隐藏/显示 房间
func (this *RoomController) IsShow() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Debug("get id error", err)
		this.AlertBack("获取房间信息失败")
		return
	}
	if this.IsAjax() {
		showed, _ := this.GetInt64("midPage")
		var newShowed int64
		var replyCount string
		if showed == 0 {
			newShowed = 1
			replyCount = "显示房间"
		} else {
			newShowed = 0
			replyCount = "隐藏房间"
		}
		_, err = models.UpdateRoomShowed(id, newShowed)
		if err != nil {
			replyCount = replyCount + "失败"
			beego.Error("inser faild", err)
			this.Rsp(false, replyCount, "")
		} else {
			replyCount = replyCount + "成功"
			this.Rsp(true, replyCount, "")
		}
		this.Rsp(false, replyCount, "")
	}
}
