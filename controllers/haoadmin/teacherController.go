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

type TeacherController struct {
	CommonController
}

func (this *TeacherController) Index() {
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
		teacherlist, count := m.GetTeacherList(iStart, iLength, "Id")
		// json
		data := make(map[string]interface{})
		data["aaData"] = teacherlist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/teacher/list.html"
	}

}

func (this *TeacherController) AddTeacher() {
	Name := this.GetString("Name")
	Title := this.GetString("Uname")
	Detail := this.GetString("Detail")
	Display, _ := this.GetInt("Display")
	Image := this.GetString("fname")
	if len(Name) > 0 && len(Detail) > 0 {
		t := new(m.Teacher)
		t.Name = Name
		t.Title = Title
		t.Detail = Detail
		t.Display = Display
		if len(Image) > 0 {
			prevalue := strings.Split(Image, "/upload/teacher/")
			t.Image = prevalue[1]
		}
		num, _ := m.QueryDisplayCount()
		if Display == 1 {
			if num > 5 {
				this.AlertBack("新增讲师失败,超过最大显示数!")
				return
			} else {
				id, err := m.AddTeacher(t)
				if err != nil && id <= 0 {
					beego.Error(err)
					this.AlertBack("讲师添加失败")
					return
				}
				this.Alert("添加成功", "teacher_index")
			}
		} else {
			id, err := m.AddTeacher(t)
			if err != nil && id <= 0 {
				beego.Error(err)
				this.AlertBack("讲师添加失败")
				return
			}
			this.Alert("添加成功", "teacher_index")
		}
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/teacher/add.html"
	}
}

// 更新头衔
func (this *TeacherController) UpdateTeacher() {
	Name := this.GetString("Name")
	Detail := this.GetString("Detail")
	Title := this.GetString("Uname")
	Display, _ := this.GetInt("Display")
	Image := this.GetString("fname")
	Id, _ := this.GetInt64("Id")

	if len(Name) > 0 && len(Detail) > 0 && Id != 0 {
		var err error
		c := new(m.Teacher)
		c.Id = Id
		c.Name = Name
		c.Title = Title
		c.Detail = Detail
		c.Display = Display
		if len(Image) > 0 {
			prevalue := strings.Split(Image, "/upload/teacher/")
			c.Image = prevalue[1]
			err = c.UpdateTeacher("Name", "Title", "Detail", "Display", "Image")
		} else {
			err = c.UpdateTeacher("Name", "Title", "Detail", "Display")
		}
		if err != nil {
			beego.Error(err)
			this.AlertBack("讲师修改失败")
			return
		}
		this.Alert("修改成功", "teacher_index")

	} else {
		this.CommonMenu()
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.AlertBack("课程讲师信息失败")
			return
		}
		teacherList, err := m.ReadTeacherById(id)
		if err != nil {
			beego.Error(err)
			this.AlertBack("课程讲师信息错误")
			return
		}
		this.Data["teacherList"] = teacherList
		this.TplName = "haoadmin/data/teacher/edit.html"
	}

}

func (this *TeacherController) DelTeacher() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelTeacherById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

func (this *TeacherController) GetAllTeacher() {
	teachers, _ := m.IndexTeacherList()
	this.Data["json"] = teachers
	this.ServeJSON()
}

func (this *TeacherController) Upload() string {
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

	dir := path.Join("..", "upload", "teacher")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		beego.Error(err)
		return FileName
	}
	// 设置保存文件名

	nowtime := time.Now().Unix()
	FileName = h.Filename
	FileName = fmt.Sprintf("%d", nowtime) + ".jpg"
	dirPath := path.Join("..", "upload", "teacher", FileName)
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

func (this *TeacherController) UploadTeacher() {
	_, _, err := this.GetFile("Filedata")
	if err == nil {
		var FileName string
		FileName = this.Upload()
		FileName = path.Join("/upload", "teacher", FileName)
		this.Rsp(true, "修改成功", FileName)
	}
}
