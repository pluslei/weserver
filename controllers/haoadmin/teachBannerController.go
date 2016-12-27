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

type TeachBannerController struct {
	CommonController
}

// 组别显示页面
func (this *TeachBannerController) Index() {
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
		titlelist, count := m.GetTeachBannerList(iStart, iLength, "Id")

		// json
		data := make(map[string]interface{})
		data["aaData"] = titlelist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/teachbanner/list.html"
	}

}

func (this *TeachBannerController) AddBanner() {
	Name := this.GetString("Name")
	Display, _ := this.GetInt("Display")
	Img := this.GetString("fname")
	Url := this.GetString("Url")
	Order, _ := this.GetInt("Order")
	if len(Name) > 0 && Display > 0 && len(Img) > 0 && len(Url) > 0 && Order > 0 {
		prevalue := strings.Split(Img, "/upload/teachbanner/")
		t := new(m.TeachBanner)
		t.Name = Name
		t.Display = Display
		t.Img = prevalue[1]
		t.Url = Url
		t.Order = Order
		id, err := m.AddTeachBanner(t)
		if err != nil && id <= 0 {
			beego.Error(err)
			this.AlertBack("添加失败")
			return
		}
		this.Alert("添加成功", "teachbanner_index")
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/teachbanner/add.html"
	}
}

func (this *TeachBannerController) UpdateBanner() {
	Name := this.GetString("Name")
	Display, _ := this.GetInt("Display")
	Id, _ := this.GetInt64("Id")
	fname := this.GetString("fname")
	Url := this.GetString("Url")
	Order, _ := this.GetInt("Order")
	if len(Name) > 0 && Id > 0 && Display > 0 {
		var err error
		t := new(m.TeachBanner)
		t.Id = Id
		t.Name = Name
		t.Display = Display
		t.Url = Url
		t.Order = Order
		if len(fname) > 0 {
			prevalue := strings.Split(fname, "/upload/teachbanner/")
			t.Img = prevalue[1]
			err = t.UpdateTeachBanner("Img", "Name", "Display", "Url", "Order")
		} else {
			err = t.UpdateTeachBanner("Name", "Display", "Url", "Order")
		}
		if err != nil {
			beego.Error(err)
			this.AlertBack("修改失败")
			return
		}
		this.Alert("修改成功", "teachbanner_index")
	} else {
		this.CommonMenu()
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取失败")
			return
		}
		techList, err := m.ReadTeachBannerById(id)
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取信息错误")
			return
		}
		this.Data["techList"] = techList
		this.TplName = "haoadmin/data/teachbanner/edit.html"
	}

}

func (this *TeachBannerController) DelBanner() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelTeachBannerById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

func (this *TeachBannerController) Upload() string {
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

	dir := path.Join("..", "upload", "teachbanner")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		beego.Error(err)
		return FileName
	}
	// 设置保存文件名

	nowtime := time.Now().Unix()
	FileName = h.Filename
	FileName = fmt.Sprintf("%d", nowtime) + ".jpg"
	dirPath := path.Join("..", "upload", "teachbanner", FileName)
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

func (this *TeachBannerController) UploadBanner() {
	_, _, err := this.GetFile("Filedata")
	var FileName string
	if err == nil {
		FileName = this.Upload()
		FileName = path.Join("/upload", "teachbanner", FileName)
		this.Rsp(true, "修改成功", FileName)
	}
}
