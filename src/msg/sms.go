package msg

import (
	"fmt"

	"github.com/astaxie/beego"
	// for json get
)

type Config struct {
	Url              string
	USER_ACCOUNT_URL string
	USER_POST_Url    string
}

var msg *SMS

func getParam() *Config {
	var info Config
	info.Url = beego.AppConfig.String("SMS_URL")
	Name := beego.AppConfig.String("SMS_NAME")
	Pwd := beego.AppConfig.String("SMS_PWD")
	ACCOUNT := beego.AppConfig.String("SMS_ACCOUNT")
	info.USER_ACCOUNT_URL = fmt.Sprintf(ACCOUNT, Name, Pwd)
	info.USER_POST_Url = beego.AppConfig.String("SMS_USER_URL")
	return &info
}

func SMSRun() {
	info := getParam()
	msg = Start(info)
	msg.Running()
	beego.Debug("SMS Init ok !")
}

func SendSMSMsg(phoneNum, sms, sign string) error {
	err := msg.sendSMSmsg(phoneNum, sms, sign)
	if err != nil {
		beego.Debug("SendTxTMsg() error:", err)
		return err
	}
	return nil
}
