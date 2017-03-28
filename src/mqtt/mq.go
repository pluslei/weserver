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
	check            bool
}

func init() {

}

var mq *MQ
var Config *Configer

func GetMqttConfig() *Configer {
	var conf Configer
	conf.MqAddress = beego.AppConfig.String("mqServerHost")
	conf.MqUserName = beego.AppConfig.String("mqServerAccess")
	conf.MqPwd = beego.AppConfig.String("mqServerKey")
	conf.MqClientID = beego.AppConfig.String("mqServerClientId")
	conf.MqIsreconnect, _ = beego.AppConfig.Bool("mqServerIsreconnect")
	conf.MqIsCleansession, _ = beego.AppConfig.Bool("mqSeverCleanSession")
	conf.MqVersion, _ = beego.AppConfig.Int("mqServerVersion")
	conf.MqTopic = beego.AppConfig.String("mqServerTopic")
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
