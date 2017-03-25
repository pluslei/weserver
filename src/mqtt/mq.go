package mqtt

import (
	. "weserver/src/tools"

	"github.com/astaxie/beego"
)

type Config struct {
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
var config *Config

func GetMqttConfig() *Config {
	var conf Config
	conf.MqAddress = beego.AppConfig.String("mqhost")
	conf.MqUserName = beego.AppConfig.String("mqaccessKey")
	conf.MqPwd = beego.AppConfig.String("mqsecretKey")
	conf.MqClientID = beego.AppConfig.String("mqclientId")
	conf.MqIsreconnect, _ = beego.AppConfig.Bool("mqIsreconnect")
	conf.MqIsCleansession, _ = beego.AppConfig.Bool("mqcleansession")
	conf.MqVersion, _ = beego.AppConfig.Int("mqVersion")
	conf.MqTopic = beego.AppConfig.String("mqtopic")
	return &conf
}

func Run() {
	// 获取配置
	config = GetMqttConfig()
	mq = NewMQ(config)
	mq.Runing()
	beego.Debug("MQTT Client Init OK.")
}

//发消息
func SendMessage(message MessageInfo) {
	mq.sendMessage(message., message)
}
