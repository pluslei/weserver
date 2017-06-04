package main

import (
	"os"

	"weserver/controllers/haoindex"
	m "weserver/models"
	_ "weserver/routers"
	. "weserver/src/cache"
	"weserver/src/mqtt"
	"weserver/src/msg"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	if _, err := os.Stat("log"); err != nil {
		os.Mkdir("log", 0755)
	}
	beego.SetLogger("file", `{"filename":"log/htrans.log"}`)
	level, err := beego.AppConfig.Int("log_level")
	if err == nil {
		beego.SetLevel(level)
	}

	// 链接数据库
	m.Connect()
	// 创建数据库
	orm.RunSyncdb("default", false, true)

	mqtt.Run()
	// wechat.WechatRun()
	// msg := "策略消息3"
	// wechat.SendTxTMsg("oWrhuv7EjuWJs6d3K3xTJ1YOlkUc", msg)

	msg.SMSRun()
	// phoneNum := "13554460187"
	// sign := "红木咨询"
	// sms := "这是一条新策略"
	// msg.SendSMSMsg(phoneNum, sign, sms)
	// msg.SendIdentifyCode(phoneNum, sign, code)

	InitCache()

	beego.ErrorController(&haoindex.ErrorController{}) //注册错误处理的函数

	// 允许跨域访问
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
	}))
	beego.Run()
}
