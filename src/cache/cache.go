package cache

import (
	"strconv"
	m "weserver/models"

	"github.com/astaxie/beego"
)

var MapCache map[string]interface{}

func GetShutMapInfo() {
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

func GetCompanyInfo() {
	companyInfo, _, err := m.GetAllCompanyInfo()
	if err != nil {
		beego.Error("get the companyInfo error", err)
	}
	for _, info := range companyInfo {
		companyaId := info.Id
		strId := strconv.FormatInt(companyaId, 10)
		_, ok := MapCache[strId]
		if !ok {
			MapCache[strId] = info
		}
	}
}

func InitCache() {
	MapCache = make(map[string]interface{})
	GetShutMapInfo()
	GetCompanyInfo()
	// beego.Debug(MapCache)
}
