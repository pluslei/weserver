package main

import (
	"weserver/controllers/haoindex"
	m "weserver/models"
	_ "weserver/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	// 链接数据库
	m.Connect()
	// 创建数据库
	orm.RunSyncdb("default", false, true)
	beego.ErrorController(&haoindex.ErrorController{}) //注册错误处理的函数

	// 允许跨域访问
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
	}))
	// 获取配置
	// Info := m.GetMqttConfig()
	// var isreconnect bool
	// if Info.MqIsreconnect == 1 {
	// 	isreconnect = true
	// } else {
	// 	isreconnect = false
	// }
	// var isclean bool
	// if Info.MqIsCleansession == 1 {
	// 	isclean = true
	// } else {
	// 	isclean = false
	// }

	// opts := mq.SetConnectOptions(
	// 	Info.MqAddress,
	// 	Info.MqUserName,
	// 	Info.MqPwd,
	// 	Info.MqClientID,
	// 	(uint)(Info.MqVersion),
	// 	isreconnect,
	// 	isclean,
	// 	mq.FU)

	// client := mq.ConnectSubScribe(opts, Info.MqTopic)
	// beego.Debug("client", client)
	m.GetMqttClient()

	beego.Run()
}
