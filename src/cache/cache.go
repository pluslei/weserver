package cache

import (
	"strconv"
	m "weserver/models"

	"github.com/astaxie/beego"
)

var Token_Url string

var MapCache map[string]interface{}
var MapPhone map[string][]string

func GetShutMapCache() {
	var status bool = true
	shutInfo, err := m.GetAllShutUpInfo()
	if err != nil {
		beego.Error("get the shutup error", err)
	}
	for _, info := range shutInfo {
		Room := info.Room
		Uname := info.Username
		inter, ok := MapCache[Room]
		if !ok {
			MapCache[Room] = []string{Uname}
		} else {
			// arr, ok := inter.([]string)
			// if ok {
			// 	for _, v := range arr {
			// 		if v == Uname {
			// 			status = false
			// 			break
			// 		}
			// 	}
			// 	if status {
			// 		arr = append(arr, Uname)
			// 		MapCache[Room] = arr
			// 	}
			// }
			switch t := inter.(type) {
			case []string:
				for _, v := range t {
					if v == Uname {
						status = false
						break
					}
				}
				if status {
					t = append(t, Uname)
					MapCache[Room] = t
				}
			default:
				beego.Debug("interface type is not found", t)
			}
		}
	}
}

func GetCompanyCache() {
	companyInfo, _, err := m.GetAllCompanyInfo()
	if err != nil {
		beego.Error("get the companyInfo error", err)
	}
	for _, info := range companyInfo {
		companyaId := info.Id
		strId := strconv.FormatInt(companyaId, 10)
		inter, ok := MapCache[strId]
		if !ok {
			MapCache[strId] = info
		} else {
			mapinfo, _ := inter.(m.Company)
			mapinfo.Company = info.Company
			mapinfo.CompanyIntro = info.CompanyIntro
			mapinfo.LoginIcon = info.LoginIcon
			mapinfo.LoginBackicon = info.LoginBackicon
			mapinfo.CompanyIcon = info.CompanyIcon
			mapinfo.CompanyBanner = info.CompanyBanner
			mapinfo.HistoryMsg = info.HistoryMsg
			mapinfo.Registerrole = info.Registerrole
			mapinfo.Registertitle = info.Registertitle
			mapinfo.WelcomeMsg = info.WelcomeMsg
			mapinfo.AuditMsg = info.AuditMsg
			mapinfo.Verify = info.Verify
			mapinfo.AppId = info.AppId
			mapinfo.AppSecret = info.AppSecret
			mapinfo.Url = info.Url
			mapinfo.Sign = info.Sign
			MapCache[strId] = mapinfo
		}
	}
}

func GetCompanyInfo(strId string) (info m.Company) {
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		beego.Debug("ParseInt error", err)
		return
	}
	inter, ok := MapCache[strId]
	if !ok {
		info, err = m.GetCompanyById(id)
		if err != nil {
			beego.Debug("get companyinfo error")
			return
		}
	} else {
		info, _ = inter.(m.Company)
	}
	return info
}

func GetPhoneNumInfo() {
	var status bool = true
	Info, err := m.GetAllPhoneNum()
	if err != nil {
		beego.Error("wechat:get all phoneNum error", err)
		return
	}
	for _, info := range Info {
		Room := info.Room
		phoneNum := info.Phonenum
		arr, ok := MapPhone[Room]
		if !ok {
			MapPhone[Room] = []string{phoneNum}
		} else {
			for _, v := range arr {
				if phoneNum == v {
					status = false
					break
				}
			}
			if status {
				arr = append(arr, phoneNum)
				MapPhone[Room] = arr
			}
		}
	}
	beego.Debug("phone num", MapPhone)
}

func GetRoomPhone(RoomId string) (info []string) {
	info, ok := MapPhone[RoomId]
	if !ok {
		arr, err := m.GetRoomPhoneNum(RoomId)
		if err != nil {
			beego.Debug("get Room Phone num Error", err)
			return nil
		}
		var arrNum []string
		for _, v := range arr {
			phoneNum := v.Phonenum
			arrNum = append(arrNum, phoneNum)
		}
		MapPhone[RoomId] = arrNum
		return arrNum
	}
	return info
}

func UpdateNewPhoneNum(oldPhoneNum, newPhoneNum string) {
	if oldPhoneNum == newPhoneNum {
		return
	}
	roomInfo, _, err := m.GetAllRoomInfo()
	if err != nil {
		beego.Debug("Get All RoomInfo error", err)
		return
	}
	for _, v := range roomInfo {
		RoomId := v.RoomId
		info, ok := MapPhone[RoomId]
		if !ok {
			arr, err := m.GetRoomPhoneNum(RoomId)
			if err != nil {
				beego.Debug("get Room Phone num Error", err)
				return
			}
			var arrNum []string
			for _, v := range arr {
				phoneNum := v.Phonenum
				arrNum = append(arrNum, phoneNum)
			}
			MapPhone[RoomId] = arrNum
		}
		for i, v := range info {
			if v == oldPhoneNum {
				index := i + 1
				info = append(info[:i], info[index:]...)
				info = append(info, newPhoneNum)
				MapPhone[RoomId] = info
				break
			}
		}
	}
	beego.Debug("Update MapPhone", MapPhone)
}

func InitCache() {
	MapCache = make(map[string]interface{})
	MapPhone = make(map[string][]string)
	GetShutMapCache()
	GetCompanyCache()
	GetPhoneNumInfo()
	beego.Debug(MapCache, MapPhone)
}
