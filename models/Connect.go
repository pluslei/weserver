package models

import (
	"fmt"
	mq "weserver/src/mqttConn"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
)

var MqClient MQTT.Client

type AccountInfo struct {
	MqAddress        string
	MqUserName       string
	MqPwd            string
	MqClientID       string
	MqIsreconnect    int //是否重连
	MqIsCleansession int
	MqVersion        int
	MqTopic          string
}

// 链接数据库
func Connect() {
	dns, _ := getConfig()
	beego.Info("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		beego.Error("数据库连接失败")
	} else {
		beego.Info("数据库连接成功")
		// writeSiteConf()
	}
}

func getConfig() (string, string) {
	db_host := beego.AppConfig.String("host")
	db_port := beego.AppConfig.String("port")
	db_user := beego.AppConfig.String("username")
	db_pass := beego.AppConfig.String("password")
	db_name := beego.AppConfig.String("dbname")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	return dns, db_name
}

func GetMqttConfig() AccountInfo {
	var info AccountInfo
	info.MqAddress = beego.AppConfig.String("mqhost")
	info.MqUserName = beego.AppConfig.String("mqaccessKey")
	info.MqPwd = beego.AppConfig.String("mqsecretKey")
	info.MqClientID = beego.AppConfig.String("mqclientId")
	info.MqIsreconnect, _ = beego.AppConfig.Int("mqIsreconnect")
	info.MqIsCleansession, _ = beego.AppConfig.Int("mqcleansession")
	info.MqVersion, _ = beego.AppConfig.Int("mqVersion")
	info.MqTopic = beego.AppConfig.String("mqtopic")
	beego.Debug("...........", info)
	return info

}

func GetMqttClient() {
	// 获取配置
	Info := GetMqttConfig()
	var isreconnect bool
	if Info.MqIsreconnect == 1 {
		isreconnect = true
	} else {
		isreconnect = false
	}
	var isclean bool
	if Info.MqIsCleansession == 1 {
		isclean = true
	} else {
		isclean = false
	}

	opts := mq.SetConnectOptions(
		Info.MqAddress,
		Info.MqUserName,
		Info.MqPwd,
		Info.MqClientID,
		(uint)(Info.MqVersion),
		isreconnect,
		isclean,
		mq.FU)
	MqClient = mq.ConnectSubScribe(opts, Info.MqTopic)
}
