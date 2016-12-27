package haoadmin

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	m "weserver/models"
)

type FaceController struct {
	CommonController
}

var FaceImg = beego.AppConfig.String("faceImg")

// 获取所有分组
func (this *FaceController) List() {
	// 获取所有的分组
	groups, _ := m.GetGroupList()
	length := len(groups)
	for i := 0; i < length; i++ {
		groups[i].GroupFace = FaceImg + groups[i].GroupFace
	}
	this.Data["group"] = groups
	this.CommonMenu()
	this.TplName = "haoadmin/data/face/index.html"
}

// 根据分组查询表情
func (this *FaceController) GetFaces() {
	group, _ := this.GetInt64("group")
	// 根据分组查询每组的表情
	faces, err := m.GetFaceByGroup(group)
	for _, v := range faces {
		v["Url"] = FaceImg + v["Url"].(string)
		v["GroupFace"] = FaceImg + v["GroupFace"].(string)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = faces
	this.ServeJSON()
}

// 删除表情信息
func (this *FaceController) DeleteFace() {
	faceIds := this.GetString("ids")
	ids := strings.Split(faceIds, ",")
	for i := 0; i < len(ids); i++ {
		id, err := strconv.ParseInt(ids[i], 10, 64)
		if err != nil {
			beego.Error(err)
		}
		if id == 0 {
			beego.Error("The face is not exists.")
		}
		face, _ := m.GetFaceById(id)
		// dirPath := path.Join(FaceImg, face.Url)
		err = os.Remove(FaceImg + face.Url)
		if err != nil {
			beego.Error(err)
			this.Rsp(false, "删除表情失败！", "")
		}
		_, err = m.DelFace(id)
		if err != nil {
			beego.Error(err)
			this.Rsp(false, "删除表情失败！", "")
		}
		this.Rsp(true, "删除表情成功！", "")
	}
}

// 获取表情信息
func (this *FaceController) GetFace() {
	Id, _ := this.GetInt64("Id")
	face, err := m.GetFaceById(Id)
	if err != nil {
		beego.Error(err)
	}
	face.GroupFace = FaceImg + face.GroupFace
	this.Data["json"] = face
	this.ServeJSON()
}

// 上传表情
func (this *FaceController) Upload() {
	_, h, err := this.GetFile("Filedata")
	if err != nil {
		beego.Error("get file error", err)
		// 获取错误则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err}
		this.ServeJSON()
		return
	}
	dir := "/upload"
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		beego.Error("mkdir upload file error", err)
		return
	}
	nowtime := time.Now().Unix()
	// 设置保存表情库
	FileName := h.Filename
	filenamebase := path.Base(FileName)
	filesuffix := path.Ext(filenamebase)
	FileName = fmt.Sprintf("%d", nowtime) + filesuffix
	dirPath := path.Join("..", FaceImg, FileName)
	err = this.SaveToFile("Filedata", dirPath)
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err}
		this.ServeJSON()
		return
	} else {
		FileName = path.Join(FaceImg, FileName)
		this.Rsp(true, "上传成功", FileName)
	}

}

// 添加或修改表情
func (this *FaceController) AddOrEdit() {
	Id := this.GetString("Id")
	fileName := this.GetString("fileName")
	groupData := this.GetString("groupData")
	group, _ := this.GetInt64("group")
	faceTitle := this.GetString("faceTitle")
	splitname := strings.Split(fileName, "/")
	splitgroup := strings.Split(groupData, "/")
	// 判断 Id 是否为空
	if Id == "" {
		face := new(m.Face)
		face.Title = faceTitle
		face.Group = group
		face.Url = splitname[len(splitname)-1]
		face.GroupFace = splitgroup[len(splitgroup)-1]
		_, err := m.AddFace(face)
		if err != nil {
			beego.Error(err)
			this.Rsp(false, "添加表情失败！", "")
		}
		this.Rsp(true, "添加表情成功！", "")
	} else {
		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		face, _ := m.GetFaceById(id)
		face.Title = faceTitle
		face.Group = group
		face.Url = splitname[len(splitname)-1]
		face.GroupFace = splitgroup[len(splitgroup)-1]
		_, err = m.EditFace(&face)
		if err != nil {
			beego.Error(err)
			this.Rsp(false, "修改表情失败！", "")
		}
		this.Rsp(true, "修改表情成功！", "")
	}
}

// 获取分组最大值
func (this *FaceController) GetMaxGroupValue() {
	maxs, _ := m.GetMaxGroup()
	max := maxs[0]["Max"].(string)
	maxvalue, err := strconv.Atoi(max)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = maxvalue
	this.ServeJSON()
}
