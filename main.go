package main

import (
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
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

	ClearDownImage()
	beego.Run()
}

// 清空目录
func ClearDownImage() {
	tk := toolbox.NewTask("tk", "0 0 23 * * 1", func() error {
		walkDir := "./static/down/"
		err := filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
			} else {
				os.Remove(path)
			}
			return err
		})
		if err != nil {
			beego.Error("error", err)
		}
		return err
	})
	toolbox.AddTask("tk", tk)
	toolbox.StartTask()
}
