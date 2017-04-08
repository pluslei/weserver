package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

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

func InitDb() {

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
