package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"weserver/controllers/haoindex"
	m "weserver/models"
	_ "weserver/routers"
	"weserver/src/socket"
	// "github.com/astaxie/beego/plugins/cors"
)

func main() {
	// 链接数据库
	m.Connect()
	// 创建数据库
	orm.RunSyncdb("default", false, true)
	beego.ErrorController(&haoindex.ErrorController{}) //注册错误处理的函数

	// 允许跨域访问
	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowAllOrigins: true,
	// }))
	//socket
	socket.Chatprogram()
	beego.Run()
}
