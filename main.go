package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"weserver/controllers/haoindex"
	m "weserver/models"
	_ "weserver/routers"
	"weserver/src/socket"
)

func main() {
	// 链接数据库
	m.Connect()
	// 创建数据库
	orm.RunSyncdb("default", false, false)
	beego.ErrorController(&haoindex.ErrorController{}) //注册错误处理的函数

	//socket
	socket.Chatprogram()
	beego.Run()
}
