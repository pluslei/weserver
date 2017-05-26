package cache

import (
	"strconv"
	m "weserver/models"

	"github.com/astaxie/beego"
)

var Token_Url string

var MapCache map[string]interface{}

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
			mapinfo.CompanyBanner = info.CompanyBanner
			mapinfo.CompanyIcon = info.CompanyIcon
			mapinfo.CompanyBanner = info.CompanyBanner
			mapinfo.HistoryMsg = info.HistoryMsg
			mapinfo.AuditMsg = info.AuditMsg
			mapinfo.Verify = info.Verify
			mapinfo.AppId = info.AppId
			mapinfo.AppSecret = info.AppSecret
			mapinfo.Url = info.Url
			mapinfo.Sign = info.Sign
			mapinfo.LoginIcon = info.LoginIcon
			mapinfo.LoginBackicon = info.LoginBackicon
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

func InitCache() {
	MapCache = make(map[string]interface{})
	GetShutMapCache()
	GetCompanyCache()
	beego.Debug(MapCache)
}
