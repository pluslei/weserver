package wechat

import (
	m "weserver/models"

	"github.com/astaxie/beego"
	// for json get
)

var chat *Wechat
var MapUname map[string][]string

type Config struct {
	appID                string
	appSecret            string
	accessTokenFetchUrl  string
	customServicePostUrl string
}

func getParam() *Config {
	var info Config
	info.appID = beego.AppConfig.String("APPID")
	info.appSecret = beego.AppConfig.String("APPSECRET")
	info.accessTokenFetchUrl = beego.AppConfig.String("TOKEN_URL")
	info.customServicePostUrl = beego.AppConfig.String("CUSOMSER_POST_URL")
	return &info
}

func Init() {
	MapUname = make(map[string][]string)
	Info, err := m.GetWechatUser(2)
	if err != nil {
		beego.Error("wechat:get the userinfo error", err)
	}
	for _, info := range Info {
		Room := info.Room
		Uname := info.Username
		v, ok := MapUname[Room]
		if !ok {
			MapUname[Room] = []string{Uname}
		} else {
			v = append(v, Uname)
			MapUname[Room] = v
		}
	}
	beego.Debug(MapUname)
}

func WechatRun() {
	Init()
	info := getParam()
	chat = Start(info)
	chat.Work()
	beego.Debug("WeChat Init ok !")
}

func SendTxTMsg(openId, msg string) error {
	err := chat.sendCustomTxTMsg(openId, msg)
	if err != nil {
		beego.Debug("SendTxTMsg error:", err)
	}
	return nil
}
