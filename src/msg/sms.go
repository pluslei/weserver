package msg

import (
	"fmt"

	"github.com/astaxie/beego"
	// for json get
)

type Config struct {
	Url string
}

func getParam() *Config {
	var info Config
	Account := beego.AppConfig.String("SMS_NAME")
	Pwd := beego.AppConfig.String("SMS_PWD")
	Url := beego.AppConfig.String("SMS_URL")
	info.Url = fmt.Sprintf(Url, Account, Pwd)
	return &info
}

func WechatRun() {
	info := getParam()
	chat = Start(info)
	chat.Running()
	beego.Debug("WeChat Init ok !")
}

func SendTxTMsg(openId, msg string) error {
	err := chat.sendCustomTxTMsg(openId, msg)
	if err != nil {
		beego.Debug("SendTxTMsg() error:", err)
		return err
	}
	beego.Debug("SendTxTMsg() ok!!")
	return nil
}
