package haoadmin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "weserver/models"

	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	tools "weserver/src/tools"
)

type RoleController struct {
	CommonController
}

// 列表显示
func (this *RoleController) Index() {
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
		rolelist, count := m.GetRolelist(iStart, iLength, "Id")
		// json
		data := make(map[string]interface{})
		data["aaData"] = rolelist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/rbac/role/list.html"
	}

}

// 增加角色
func (this *RoleController) AddRole() {
	action := this.GetString("action")
	if action == "add" {

		// Name := this.GetString("Name")
		Randnum, _ := this.GetInt("Randnum")
		Title := this.GetString("Title")
		Name := tools.Strtomd5(Title)
		Status, _ := this.GetInt("Status")
		Remark := this.GetString("Remark")
		Weight, _ := this.GetInt("Weight")
		Delay, _ := this.GetInt("Delay")
		RandTitle, _ := this.GetInt64("RandTitle")
		IsInsider, _ := this.GetInt("IsInsider")
		fname := this.GetString("fname")
		splitname := strings.Split(fname, "/")

		r := new(m.Role)
		r.Name = Name
		r.Title = Title
		r.Status = Status
		r.Remark = Remark
		r.Weight = Weight
		r.Delay = Delay
		r.Randnum = Randnum
		r.Ico = splitname[len(splitname)-1]
		r.IsInsider = IsInsider
		r.RandTitle = &m.Title{Id: RandTitle}
		id, err := m.AddRole(r)
		if err != nil && id <= 0 {
			beego.Error(err)
			this.AlertBack("角色添加失败")
			return
		}
		this.Alert("添加成功", "index")
	} else {
		this.Data["TitleList"] = m.TitleList()
		this.CommonMenu()
		this.TplName = "haoadmin/rbac/role/add.html"
	}
}

// 更新角色
func (this *RoleController) UpdataRole() {
	action := this.GetString("action")
	if action == "edit" {
		// Name := this.GetString("Name")
		Randnum, _ := this.GetInt("Randnum")
		Title := this.GetString("Title")
		Status, _ := this.GetInt("Status")
		Remark := this.GetString("Remark")
		Id, _ := this.GetInt64("Id")
		Weight, _ := this.GetInt("Weight")
		Delay, _ := this.GetInt("Delay")
		RandTitle, _ := this.GetInt64("RandTitle")
		IsInsider, _ := this.GetInt("IsInsider")
		fname := this.GetString("fname")
		splitname := strings.Split(fname, "/")
		r := new(m.Role)
		r.Id = Id
		// r.Name = Name
		r.Title = Title
		r.Status = Status
		r.Remark = Remark
		r.Weight = Weight
		r.Delay = Delay
		r.RandTitle = &m.Title{Id: RandTitle}
		r.IsInsider = IsInsider
		r.Randnum = Randnum
		r.Ico = splitname[len(splitname)-1]
		err := r.UpdateRoleFields("Title", "Status", "Remark", "Weight", "Delay", "IsInsider", "Randnum", "Ico", "RandTitle")
		if err == nil {
			this.Alert("更新成功", "index")
			return
		} else {
			this.AlertBack("更新失败")
			return
		}
	} else {
		this.CommonMenu()
		roleid, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取ID失败")
			return
		}
		roleList, _ := m.GetRoleInfoById(roleid)
		this.Data["TitleList"] = m.TitleList()
		this.Data["roleList"] = roleList
		this.Data["imgUrl"] = "/upload/usertitle/" + roleList.Ico
		this.TplName = "haoadmin/rbac/role/edit.html"
	}
}

// 删除节点
func (this *RoleController) DelRole() {
	Id, _ := this.GetInt64("Id")
	beego.Debug(Id)
	status, err := m.DelRoleById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

// 赋予角色访问权限
func (this *RoleController) AddAccess() {
	action := this.GetString("action")
	if action == "add" {
		roleid, _ := this.GetInt64("roleid")
		group_id, _ := this.GetInt64("group_id")
		err := m.DelGroupNode(roleid, group_id)
		if err != nil {
			beego.Error(err)
			return
		}
		nodeids := this.GetStrings("ids")
		for _, v := range nodeids {
			id, _ := strconv.Atoi(v)
			_, err := m.AddRoleNode(roleid, int64(id))
			if err != nil {
				beego.Error(err)
				this.Rsp(false, err.Error(), "")
			}
		}
		this.Alert("权限成功", "index")
	} else {
		roleid, err := this.GetInt64("roleid")
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取ID错误")
			return
		}
		this.CommonMenu()
		groupList := m.GroupList()
		this.Data["roleId"] = roleid
		this.Data["groupList"] = groupList
		this.TplName = "haoadmin/rbac/role/addaccess.html"
	}
}

// AJAX获取节点列表
func (this *RoleController) AccessToNode() {
	roleid, err := this.GetInt64("roleid")
	if err != nil {
		beego.Error(err)
		this.AlertBack("获取ID错误")
		return
	}

	groupid, _ := this.GetInt64("group_id")
	nodes, count := m.GetNodelistByGroupid(groupid)
	list, _ := m.GetNodelistByRoleId(roleid)
	for i := 0; i < len(nodes); i++ {
		if nodes[i]["Pid"] != 0 {
			nodes[i]["_parentId"] = nodes[i]["Pid"]
		} else {
			nodes[i]["state"] = "closed"
		}
		for x := 0; x < len(list); x++ {
			if nodes[i]["Id"] == list[x]["Id"] {
				nodes[i]["Checked"] = "checked"
			}
		}
	}
	if len(nodes) < 1 {
		nodes = []orm.Params{}
	}
	this.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
	this.ServeJSON()

}

func (this *RoleController) Getlist() {
	roles, _ := m.GetRolelist(1, 1000, "Id")
	if len(roles) < 1 {
		roles = []orm.Params{}
	}
	this.Data["json"] = &roles
	this.ServeJSON()
	return
}

// 获取所有角色json
func (this *RoleController) GetAllRole() {
	roles, _ := m.GetAllUserRole()
	var roleJson = "{"
	for _, item := range roles {
		itemjson := fmt.Sprintf(`%d:"%s"`, item.Id, item.Title)
		roleJson = roleJson + itemjson + ","
	}
	roleJson = strings.TrimRight(roleJson, ",")
	roleJson = roleJson + "}"
	this.Ctx.WriteString(roleJson)
}

func (this *RoleController) Upload() {
	_, h, err := this.GetFile("Filedata")
	if err != nil {
		beego.Error("get file error", err)
		// 获取错误则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err}
		this.ServeJSON()
		return
	}

	dir := "upload"
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		beego.Error("mkdir upload file error", err)
		return
	}
	// 设置保存文件名
	nowtime := time.Now().Unix()
	FileName := h.Filename
	FileName = fmt.Sprintf("%d", nowtime) + ".jpg"
	dirPath := path.Join("..", "upload", "usertitle", FileName)
	// 将文件保存到服务器中
	err = this.SaveToFile("Filedata", dirPath)
	beego.Debug(dirPath)
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err}
		this.ServeJSON()
		return
	} else {
		FileName = path.Join("/upload", "usertitle", FileName)
		this.Rsp(true, "修改成功", FileName)
	}
}
