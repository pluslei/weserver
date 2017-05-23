package msg

import (
	"fmt"

	"github.com/astaxie/beego"
	// for json get
)

type Config struct {
	Url      string
	USER_Url string
}

var msg *SMS

func getParam() *Config {
	var info Config
	Account := beego.AppConfig.String("SMS_NAME")
	Pwd := beego.AppConfig.String("SMS_PWD")
	Url := beego.AppConfig.String("SMS_URL")
	info.USER_Url = beego.AppConfig.String("SMS_USER_URL")
	info.Url = fmt.Sprintf(Url, Account, Pwd)
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
