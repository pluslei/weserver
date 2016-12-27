package haoadmin

import (
	"github.com/astaxie/beego"
	simplejson "github.com/bitly/go-simplejson"
	m "weserver/models"
	p "weserver/src/parameter"
	tools "weserver/src/tools"

	"fmt"
	"io/ioutil"
	"strconv"
)

type RoomController struct {
	CommonController
}

var serveraddress = beego.AppConfig.String("ServerAddress")

func (this *RoomController) Index() {
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
		roomlist, count := m.GetRoomList(iStart, iLength, "Id")

		// json
		data := make(map[string]interface{})
		data["aaData"] = roomlist
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

// 刷新房间
func (this *RoomController) RefreshRoom() {
	data := fmt.Sprintf(`{"action":"roomlist","company":"%s"}`, p.Code)
	parameter := map[string]string{"data": tools.MainEncrypt(data)}
	res, resperr := this.HttpnRmageRequest(serveraddress, parameter)
	if resperr != nil {
		beego.Error(resperr)
		this.Rsp(false, "数据请求错误", "")
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		this.Rsp(false, "数据读取错误", "")
		beego.Error(err)
		return
	}
	res.Body.Close()
	jsonStr, err := tools.MainDecrypt(string(result))
	if err != nil {
		beego.Error(fmt.Sprintf("decrypt data error:%v", err))
		this.Rsp(false, "数据通讯错误", "")
		return
	}
	var roomjson *simplejson.Json
	js, err := simplejson.NewJson(jsonStr)
	code := js.Get("code").MustInt64()
	msg := js.Get("msg").MustString("")
	roomjson = js.Get("roomlist")
	if code < 0 {
		this.Rsp(false, msg, "")
	} else {
		var success int
		for i := 0; i < len(roomjson.MustArray()); i++ {
			Streams := roomjson.GetIndex(i).Get("streams").MustString()
			ActivityId := roomjson.GetIndex(i).Get("activityid").MustString()
			num := m.GetIsRoom(Streams, ActivityId)
			if num <= 0 {
				RoomNumber, _ := m.RoomNumber()
				room := new(m.Room)
				room.Nickname = roomjson.GetIndex(i).Get("nickname").MustString()
				room.Rtmpurl = roomjson.GetIndex(i).Get("rtmpurl").MustString()
				room.Streams = Streams
				room.ActivityId = ActivityId
				room.RoomDescription = roomjson.GetIndex(i).Get("roomdescription").MustString()
				room.RoomStatus = roomjson.GetIndex(i).Get("roomstatus").MustInt()
				room.Centerid = roomjson.GetIndex(i).Get("id").MustInt64()
				room.CompanyCode = roomjson.GetIndex(i).Get("companycode").MustString()
				room.RommNumber = RoomNumber + 10000
				num, err := m.InsertRoom(room)
				if err == nil && num > 0 {
					success++
				}
			}
		}
		if err != nil {
			this.Rsp(false, "数据插入失败", "")
			beego.Error(err)
		} else {
			if success > 0 {
				inser := strconv.Itoa(success)
				this.Rsp(true, "成功获取到 "+inser+" 数据", "")
			} else {
				this.Rsp(true, "没有更新的数据", "")
			}

		}
	}
}

// 修改房间
func (this *RoomController) UpdateRoom() {
	roomid, err := this.GetInt64("Id")
	if err != nil {
		beego.Error(err)
		this.Rsp(false, "参数错误", "")
		return
	}
	action := this.GetString("action")
	roominfo, err := m.GetRoomInfo(roomid)
	if err != nil {
		beego.Error(err)
		this.Rsp(false, "参数错误", "")
		return
	}
	if action == "1" {
		Nickname := this.GetString("Nickname")
		RoomDescription := this.GetString("RoomDescription")

		data := fmt.Sprintf(`{"action":"updateroominfo","company":"%s","id":%d,"nickname":"%s","roomdescription":"%s"}`, p.Code, roominfo.Centerid, Nickname, RoomDescription)
		parameter := map[string]string{"data": tools.MainEncrypt(data)}
		res, resperr := this.HttpnRmageRequest(serveraddress, parameter)
		if resperr != nil {
			beego.Error(resperr)
			this.Rsp(false, "数据请求错误", "")
			return
		}
		result, err := ioutil.ReadAll(res.Body)
		if err != nil {
			this.Rsp(false, "数据读取错误", "")
			beego.Error(err)
			return
		}
		res.Body.Close()
		jsonStr, err := tools.MainDecrypt(string(result))
		if err != nil {
			beego.Error(fmt.Sprintf("decrypt data error:%v", err))
			this.Rsp(false, "数据通讯错误", "")
			return
		}
		js, _ := simplejson.NewJson(jsonStr)
		code := js.Get("code").MustInt64()
		msg := js.Get("msg").MustString("")
		if code == 0 {
			_, err = m.UpdateRoom(roomid, Nickname, RoomDescription)
			if err != nil {
				this.Rsp(false, "更新失败", "")
				beego.Error(err)
			} else {
				this.Rsp(true, "更新成功", "")
			}
		} else {
			this.Rsp(false, msg, "")
		}

	} else {
		this.Data["roominfo"] = roominfo
		this.TplName = "haoadmin/data/room/edit.html"
	}

}
