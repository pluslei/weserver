package mqtt

import (
	m "weserver/models"

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
	MqTopic          map[string]byte
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
	// mqTopic := beego.AppConfig.String("mqServerTopic")
	// slice := strings.Split(",", mqTopic)
	// for index := range slice {
	// 	conf.MqTopic[slice[index]] = 0
	// }
	room, count, err := m.GetRoomName()
	if err != nil {
		beego.Error("Get Room Topic Fail", err)
		return nil
	}
	beego.Debug("Get Room Topic Num: ", count)
	conf.MqTopic = make(map[string]byte)
	for k, v := range room {
		value, ok := v.(byte)
		if !ok {
			beego.Debug("type is no match")
		}
		conf.MqTopic[k] = value
	}
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
func SendMessage(topic, message string) {
	mq.sendMessage(topic, message)
}
