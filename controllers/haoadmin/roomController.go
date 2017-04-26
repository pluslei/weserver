package haoadmin

import (
	"github.com/astaxie/beego"
)

type RoomController struct {
	CommonController
}

func (this *RoomController) Index(){
		beego.Debug("=")
		this.CommonMenu()
		this.TplName = "haoadmin/data/room/list.html"
}