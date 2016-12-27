package haoadmin

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"path"
	"strings"
	"time"
	m "weserver/models"
)

type TelBannerController struct {
	CommonController
}

func (this *TelBannerController) TelBannerIndex() {
	action := this.GetString("action")
	if action == "edit" {
		var err error
		url := this.GetString("fname")
		detail := this.GetString("detail")
		c := new(m.TelBanner)
		c.Id = 1
		c.Detail = detail
		if len(url) > 0 {
			prevalue := strings.Split(url, "/upload/telbanner/")
			c.Url = prevalue[1]
			err = c.UpdateTelBanner("Url", "Detail")
		} else {
			err = c.UpdateTelBanner("Detail")
		}
		if err != nil {
			beego.Error(err)
			this.AlertBack("更新失败")
			return
		} else {
			this.Alert("修改成功", "./telbanner_index")
			return
		}
	} else {
		this.CommonMenu()
		info := m.GetTelBannerInfo(1)
		this.Data["info"] = info
		this.TplName = "haoadmin/data/telbanner/edit.html"
	}
}

func (this *TelBannerController) Upload() string {
	var FileName string
	f, h, err := this.GetFile("Filedata")
	if err == nil {
		// 关闭文件
		f.Close()
	}
	if err != nil {
		// 获取错误则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err}
		this.ServeJSON()
		return FileName
	}

	dir := path.Join("..", "upload", "telbanner")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		beego.Error(err)
		return FileName
	}
	// 设置保存文件名

	nowtime := time.Now().Unix()
	FileName = h.Filename
	FileName = fmt.Sprintf("%d", nowtime) + ".jpg"
	dirPath := path.Join("..", "upload", "telbanner", FileName)
	// 将文件保存到服务器中
	err = this.SaveToFile("Filedata", dirPath)
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err}
		this.ServeJSON()
		return FileName
	}
	return FileName
}

func (this *TelBannerController) UploadTelBanner() {
	_, _, err := this.GetFile("Filedata")
	var FileName string
	if err == nil {
		FileName = this.Upload()
		FileName = path.Join("/upload", "telbanner", FileName)
		this.Rsp(true, "修改成功", FileName)
	}
}
