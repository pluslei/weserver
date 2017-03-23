package mqtt

import (
	"github.com/astaxie/beego"
	"weserver/src/tools"
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
	MqCheckTopic     string
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
	conf.MqCheckTopic = beego.AppConfig.String("mqCheckTopic")
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
func SendMessage(message tools.MessageInfo) {
	// 开启消息审核
	if Config.check {
		mq.sendMessage(Config.MqCheckTopic, message)
	} else {
		mq.sendMessage(Config.MqTopic, message)
	}
}
