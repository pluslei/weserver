package haoadmin

import (
	"github.com/astaxie/beego"
	m "weserver/models"
	"weserver/src/tools"
)

type NetNameController struct {
	CommonController
}

// 添加网名
func (this *NetNameController) AddNetName() {
	for i := 0; i <= 50000; i++ {
		netName := tools.GetNetName()
		_, netname := m.IsExitNetName(netName)
		if netname.Id != 0 {
			beego.Error("The NetName is exits.")
		} else {
			internetName := new(m.NetName)
			internetName.Name = netName
			id, err := m.AddNetName(internetName)
			if err != nil && id <= 0 {
				beego.Error(err)
			}
		}
		if i == 50000 {
			this.Alert("50000次已经循环完成。正在跳转到 消息库管理", "./message_index")
		}
	}
}
