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

// SMS_NAME = 13552389469
// SMS_PWD = 25FE2D232B6A0ABA40508CFC1E0F
// SMS_URL = "http://web.wasun.cn/asmx/smsservice.aspx?"
// SMS_ACCOUNT = "name=%s&pwd=%s&"
// SMS_USER_URL = "mobile=%s&content=%s&sign=%s&stime=&type=pt&extno="

var msg *SMS

func getParam() *Config {
	var info Config
	info.Url = beego.AppConfig.String("SMS_URL")
	Name := beego.AppConfig.String("SMS_NAME")
	Pwd := beego.AppConfig.String("SMS_PWD")
	info.USER_ACCOUNT_URL = fmt.Sprintf(info.Url, Name, Pwd)
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
