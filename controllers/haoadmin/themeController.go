package haoadmin

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"path"
	"time"
	m "weserver/models"
	tools "weserver/src/tools"
)

type ThemeController struct {
	CommonController
}

func (this *ThemeController) Index() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")
		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		qslist, count := m.GetThemelist(iStart, iLength, "Id")

		// json
		data := make(map[string]interface{})
		data["aaData"] = qslist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/theme/list.html"
	}

}

func (this *ThemeController) AddTheme() {
	Name := this.GetString("Name")
	if len(Name) > 0 {
		t := new(m.Theme)
		t.Name = Name
		t.Img = this.Upload()
		t.Status = 1
		id, err := m.AddTheme(t)
		if err != nil && id <= 0 {
			beego.Error(err)
			this.AlertBack("主题添加失败")
			return
		}
		this.Alert("添加成功", "theme_index")
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/theme/add.html"
	}

}

func (this *ThemeController) UpdateTheme() {
	oldId, _ := this.GetInt64("oldId")
	Id, _ := this.GetInt64("Id")
	if Id > 0 {
		q := new(m.Theme)
		q.Id = oldId
		q.Status = 1
		err := q.UpdateTheme("Status")
		q.Id = Id
		q.Status = 2
		err = q.UpdateTheme("Status")
		if err != nil {
			beego.Error(err)
			this.Rsp(false, "主题修改失败", "")
			return
		}
		this.Rsp(true, "修改成功", "index")
	} else {
		this.CommonMenu()
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.Rsp(false, "主题修改失败", "")
			return
		}
		themeList, err := m.ReadThemeById(id)
		if err != nil {
			beego.Error(err)
			this.Rsp(false, "获取主题信息错误", "")
			return
		}
		this.Data["themeList"] = themeList
		this.TplName = "haoadmin/data/theme/edit.html"
	}

}

func (this *ThemeController) DelTheme() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelThemeById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

func (this *ThemeController) Upload() string {
	var FileName string
	f, h, err := this.GetFile("Img")
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

	dir := path.Join("..", "upload", "theme")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		beego.Error(err)
		return FileName
	}
	// 设置保存文件名
	nowtime := time.Now().Unix()
	FileName = h.Filename
	FileName = fmt.Sprintf("%d", nowtime) + ".png"
	dirPath := path.Join("..", "upload", "theme", FileName)
	// 将文件保存到服务器中
	err = this.SaveToFile("Img", dirPath)
	tools.Imagepro(dirPath, dirPath, 150, 250)
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err}
		this.ServeJSON()
		return FileName
	}
	return FileName
}
