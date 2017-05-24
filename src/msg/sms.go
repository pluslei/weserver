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
	USER_IDENTI_Url  string
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
	info.USER_IDENTI_Url = beego.AppConfig.String("SMS_IDENTI_URL")
	return &info
}

func SMSRun() {
	info := getParam()
	msg = Start(info)
	msg.RunSMSing()
	msg.RunCodeing()
	beego.Debug("SMS Init ok !")
}

func SendSMSMsg(phoneNum, sign, sms string) error {
	err := msg.sendSMSmsg(phoneNum, sign, sms)
	if err != nil {
		beego.Debug("SendTxTMsg() error:", err)
		return err
	}
	return nil
}

func SendIdentifyCode(phoneNum, sign string, code int64) error {
	err := msg.sendSMSCode(phoneNum, sign, code)
	if err != nil {
		beego.Debug("SendTxTMsg() error:", err)
		return err
	}
	return nil
}
