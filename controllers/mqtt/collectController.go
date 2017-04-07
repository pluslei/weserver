package mqtt

import (
	m "weserver/models"

	"github.com/astaxie/beego"

	"weserver/controllers"
	// for json get
)

type CollectController struct {
	controllers.PublicController
}

//收藏
func (this *CollectController) GetCollectInfo() {
	if this.IsAjax() {
		collect := new(m.Collect)
		collect.Uname = this.GetString("Uname")
		collect.Nickname = this.GetString("Nickname")
		collect.RoomIcon = this.GetString("RoomIcon")
		collect.RoomTitle = this.GetString("RoomTitle")
		collect.RoomTeacher = this.GetString("RoomTeacher")
		collect.IsCollect, _ = this.GetBool("IsCollect")
		collect.IsAttention, _ = this.GetBool("IsAttention")
		_, err := m.AddCollect(collect)
		if err != nil {
			this.Rsp(false, "收藏写入数据库失败", "")
			beego.Debug(err)
		} else {
			this.Rsp(true, "收藏写入数据库成功", "")
		}
	}
}
