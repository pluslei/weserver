package cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	m "weserver/models"

	"github.com/astaxie/beego"
)

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
			mapinfo.HistoryMsg = info.HistoryMsg
			mapinfo.AuditMsg = info.AuditMsg
			mapinfo.Verify = info.Verify
			mapinfo.AppId = info.AppId
			mapinfo.AppSecret = info.AppSecret
			mapinfo.Url = info.Url
			MapCache[strId] = mapinfo
		}
	}
}

func getAccessToken() error {

	requestLine := fmt.Sprintf(w.accessTokenFetchUrl, w.appID, w.appSecret)
	beego.Debug("GetTokenString", requestLine)

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	beego.Debug("AccessToken:", string(body))
	if bytes.Contains(body, []byte("access_token")) {
		beego.Debug("Request ok!!!!!!!")
		err = json.Unmarshal(body, w)
		if err != nil {
			return err
		}
		w.TextUrl = fmt.Sprintf(w.customServicePostUrl, w.AccessToken)
		w.TemplateUrl = fmt.Sprintf(w.templatePostUrl, w.AccessToken)
	} else {
		beego.Debug("Request error!!!!!!!")
		err = json.Unmarshal(body, w)
		if err != nil {
			return err
		}
	}
	// beego.Debug(w.AccessToken, w.ExpiresIn, w.TextUrl)
	return nil
}

func InitCache() {
	MapCache = make(map[string]interface{})
	GetShutMapCache()
	GetCompanyCache()
	beego.Debug(MapCache)
}
