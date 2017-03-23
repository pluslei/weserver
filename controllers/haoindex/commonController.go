package haoindex

import (
	"os"

	"github.com/astaxie/beego"
	"weserver/controllers"
)

// 获取直播地址
var LiveUrl = beego.AppConfig.String("liveurl")

type CommonController struct {
	controllers.PublicController
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
