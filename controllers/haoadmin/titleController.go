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

type TitleController struct {
	CommonController
}

// 组别显示页面
func (this *TitleController) Index() {
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
		titlelist, count := m.GetTitlelist(iStart, iLength, "Id")

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
		this.TplName = "haoadmin/rbac/title/list.html"
	}

}

func (this *TitleController) AddTitle() {
	Name := this.GetString("Name")
	Remark := this.GetString("Remark")
	Weight, _ := this.GetInt("Weight")
	Css := this.GetString("Css")
	Background := this.GetString("BackGround")
	if len(Name) > 0 && len(Css) > 0 {
		t := new(m.Title)
		t.Name = Name
		t.Remark = Remark
		t.Css = Css
		t.Weight = Weight
		t.Background = Background
		id, err := m.AddTitle(t)
		if err != nil && id <= 0 {
			beego.Error(err)
			this.AlertBack("头衔添加失败")
			return
		}
		this.Alert("添加成功", "index")
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/rbac/title/add.html"
	}
}

// 更新头衔
func (this *TitleController) UpdateTitle() {
	Name := this.GetString("Name")
	Remark := this.GetString("Remark")
	Weight, _ := this.GetInt("Weight")
	Id, _ := this.GetInt64("Id")
	Css := this.GetString("Css")
	Background := this.GetString("BackGround")
	if len(Name) > 0 {
		var err error
		t := new(m.Title)
		t.Id = Id
		t.Name = Name
		t.Remark = Remark
		t.Weight = Weight
		t.Css = Css
		t.Background = Background
		err = t.UpdateTitle("Name", "Remark", "Css", "Weight", "Background")
		if err != nil {
			beego.Error(err)
			this.AlertBack("头衔修改失败")
			return
		}
		this.Alert("修改成功", "index")
	} else {
		this.CommonMenu()
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.AlertBack("头衔获取失败")
			return
		}
		titleList, err := m.ReadTitleById(id)
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取头衔信息错误")
			return
		}
		this.Data["titleList"] = titleList
		this.TplName = "haoadmin/rbac/title/edit.html"
	}

}

func (this *TitleController) DelTitle() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelTitleById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

func (this *TitleController) Upload() string {
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

	dir := path.Join("..", "upload", "usertitle")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		beego.Error(err)
		return FileName
	}
	// 设置保存文件名
	nowtime := time.Now().Unix()
	FileName = h.Filename
	FileName = fmt.Sprintf("%d", nowtime) + ".jpg"
	dirPath := path.Join("..", "upload", "usertitle", FileName)
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

func (this *TitleController) UploadTitle() {
	_, _, err := this.GetFile("Filedata")
	var FileName string
	if err == nil {
		FileName = this.Upload()
		FileName = path.Join("/upload", "usertitle", FileName)
		this.Rsp(true, "修改成功", FileName)
	}
}

func (this *TitleController) GetAllTitle() {
	title, _ := m.GetAllUserTitle()
	var titleJson = "{"
	for _, item := range title {
		itemjson := fmt.Sprintf(`%d:"%s"`, item.Id, item.Name)
		titleJson = titleJson + itemjson + ","
	}
	titleJson = strings.TrimRight(titleJson, ",")
	titleJson = titleJson + "}"
	this.Ctx.WriteString(titleJson)
}
