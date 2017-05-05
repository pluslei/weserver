package mqtt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	m "weserver/models"

	"github.com/astaxie/beego"
)

type Configer struct {
	MqAddress        string
	MqUserName       string
	MqPwd            string
	MqClientID       string
	MqGroupID        string
	MqIsreconnect    bool //是否重连
	MqIsCleansession bool
	MqVersion        int
	MqTopic          string //单级订阅
	check            bool
	// MqTopic          map[string]byte //多级订阅
}

var mq *MQ
var Config *Configer
var MapShutUp map[string][]string

//get pwd
func getSecretKey(secretkey, groupId string) string {
	key := []byte(secretkey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(groupId))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func GetMqttConfig() *Configer {
	var conf Configer
	conf.MqAddress = beego.AppConfig.String("mqServerHost")
	conf.MqUserName = beego.AppConfig.String("mqServerAccess")
	conf.MqClientID = beego.AppConfig.String("mqServerClientId")
	conf.MqIsreconnect, _ = beego.AppConfig.Bool("mqServerIsreconnect")
	conf.MqIsCleansession, _ = beego.AppConfig.Bool("mqSeverCleanSession")
	conf.MqVersion, _ = beego.AppConfig.Int("mqServerVersion")
	conf.MqTopic = beego.AppConfig.String("mqServerTopic")
	key := beego.AppConfig.String("mqServerKey")
	groupId := beego.AppConfig.String("mqServerGroupId")
	conf.MqPwd = getSecretKey(key, groupId)
	beego.Debug(conf)
	/*
		// 多级订阅
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
	*/
	return &conf
}

func GetShutMapInfo() {
	var status bool = true
	shutInfo, err := m.GetShutUpInfoToday()
	if err != nil {
		beego.Error("get the shutup error", err)
	}
	for _, info := range shutInfo {
		Room := info.Room
		Uname := info.Username
		arr, ok := MapShutUp[Room]
		if !ok {
			MapShutUp[Room] = []string{Uname}
		} else {
			for _, v := range arr {
				if v == Uname {
					status = false
					break
				}
			}
			if status {
				arr = append(arr, Uname)
				MapShutUp[Room] = arr
			}
		}
	}
}

func Init() {
	MapShutUp = make(map[string][]string)
	GetShutMapInfo()
	beego.Debug(MapShutUp)
}

func Run() {
	// 获取配置
	Init()
	Config = GetMqttConfig()
	mq = NewMQ(Config)
	mq.Runing()
	beego.Debug("MQTT Client Init OK.")
}

//发消息
func SendMessage(topic, message string) {
	mq.sendMessage(topic, message)
}
