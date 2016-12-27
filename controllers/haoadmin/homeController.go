package haoadmin

import (
	"github.com/astaxie/beego"
	m "weserver/models"
)

type HomeController struct {
	CommonController
}

func (this *HomeController) AboutUs() {
	action := this.GetString("action")
	if action == "edit" {
		content := this.GetString("content")
		c := new(m.Company)
		c.Id = 1
		c.Content = content
		err := c.UpdateCompanyFields("Content")
		if err != nil {
			beego.Error(err)
			this.AlertBack("更新失败")
			return
		} else {
			this.Alert("修改成功", "./aboutme")
			return
		}
	} else {
		this.CommonMenu()
		info := m.GetCompanyInfo(1)
		this.Data["info"] = info
		this.TplName = "haoadmin/home/aboutme.html"
	}
}

func (this *HomeController) ContactUs() {
	action := this.GetString("action")
	if action == "edit" {
		content := this.GetString("content")
		c := new(m.Company)
		c.Id = 2
		c.Content = content
		err := c.UpdateCompanyFields("Content")
		if err != nil {
			beego.Error(err)
			this.AlertBack("更新失败")
			return
		} else {
			this.Alert("修改成功", "./contact")
			return
		}
	} else {
		this.CommonMenu()
		info := m.GetCompanyInfo(2)
		this.Data["info"] = info
		this.TplName = "haoadmin/home/contact.html"
	}
}
