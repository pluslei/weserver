package mqtt

import (
	"github.com/astaxie/beego"
)

type Configer struct {
	MqAddress        string
	MqUserName       string
	MqPwd            string
	MqClientID       string
	MqIsreconnect    bool //是否重连
	MqIsCleansession bool
	MqVersion        int
	MqTopic          string
	MqusernameAccess string
	MqpasswordSecret string
	Mqport           int
	MquseTLS         int
	MqgroupId        string
	Mqurl            string
	check            bool
}

func init() {

}

var mq *MQ
var Config *Configer

func GetMqttConfig() *Configer {
	var conf Configer
	conf.MqAddress = beego.AppConfig.String("mqhost")
	conf.MqUserName = beego.AppConfig.String("mqaccessKey")
	conf.MqPwd = beego.AppConfig.String("mqsecretKey")
	conf.MqClientID = beego.AppConfig.String("mqclientId")
	conf.MqIsreconnect, _ = beego.AppConfig.Bool("mqIsreconnect")
	conf.MqIsCleansession, _ = beego.AppConfig.Bool("mqcleansession")
	conf.MqVersion, _ = beego.AppConfig.Int("mqVersion")
	conf.MqTopic = beego.AppConfig.String("mqtopic")

	conf.Mqurl = beego.AppConfig.String("Mqurl")
	conf.MqusernameAccess = beego.AppConfig.String("mqusernameAccess")
	conf.MqpasswordSecret = beego.AppConfig.String("mqpasswordSecret")
	conf.Mqport, _ = beego.AppConfig.Int("mqport")
	conf.MquseTLS, _ = beego.AppConfig.Int("mquseTLS")
	conf.MqgroupId = beego.AppConfig.String("MqgroupId")
	return &conf
}

func Run() {
	// 获取配置
	Config = GetMqttConfig()
	mq = NewMQ(Config)
	mq.Runing()
	beego.Debug("MQTT Client Init OK.")
}

//发消息

func SendMessage(message string) {
	mq.sendMessage(Config.MqTopic, message)
}
