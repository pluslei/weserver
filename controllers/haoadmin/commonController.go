package haoadmin

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"weserver/controllers"
	m "weserver/models"
)

type CommonController struct {
	controllers.PublicController
}

func (this *CommonController) GetResList(uname string, Id int64) []Tree {
	var cnt, length int = 0, 0
	var nodes []orm.Params
	adminUser := beego.AppConfig.String("rbac_admin_user")
	if uname == adminUser {
		_, nodes = m.GetAllNode()
	} else {
		nodes, _ = m.GetNodeByRoleId(Id)
	}

	// 计算数组的最大长度
	for _, v := range nodes {
		if v["Pid"].(int64) == 0 {
			length = length + 1
		}
	}
	tree := make([]Tree, length)

	for k, v := range nodes {
		if v["Pid"].(int64) == 0 {
			k = cnt
			cnt = cnt + 1
			tree[k].Id = v["Id"].(int64)
			tree[k].Name = v["Name"].(string)
			tree[k].Index = cnt
			tree[k].Text = v["Title"].(string)
			tree[k].Url = v["Url"].(string)
			tree[k].Ico = v["Ico"].(string)
			// 1代表菜单（目录下面的所有资源）没有把一些不需要的权限去掉

			var childCnt int = 0
			children := make([]map[string]interface{}, 8)
			for _, v3 := range nodes {
				if v3["Pid"].(int64) == v["Id"].(int64) {
					children[childCnt] = v3
					childCnt++
				}
			}

			tree[k].Children = make([]Tree, childCnt)
			for k1, v1 := range children {
				if v1 == nil {

				} else {
					if v1["Pid"].(int64) == v["Id"].(int64) {
						tree[k].Children[k1].Id = v1["Id"].(int64)
						tree[k].Children[k1].Name = v1["Name"].(string)
						tree[k].Children[k1].Text = v1["Title"].(string)
						tree[k].Children[k1].Ico = v1["Ico"].(string)
						tree[k].Children[k1].Url = v1["Url"].(string)
					}
				}
			}
		}

	}
	return tree
}

func (this *CommonController) GetTree() []Tree {
	nodes, _ := m.GetNodeTree(0, 1)
	tree := make([]Tree, len(nodes))
	for k, v := range nodes {
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Title"].(string)
		children, _ := m.GetNodeTree(v["Id"].(int64), 2)
		tree[k].Children = make([]Tree, len(children))
		for k1, v1 := range children {
			tree[k].Children[k1].Id = v1["Id"].(int64)
			tree[k].Children[k1].Text = v1["Title"].(string)
			// tree[k].Children[k1].Attributes.Url = "/" + v["Name"].(string) + "/" + v1["Name"].(string)
		}
	}
	return tree
}

// 按组输出节点树
func (this *CommonController) GetGroupTree(gid int64) []Tree {
	nodes, _ := m.GetNodeGroupTree(0, 1, gid)
	tree := make([]Tree, len(nodes))
	for k, v := range nodes {
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Title"].(string)
		children, _ := m.GetNodeGroupTree(v["Id"].(int64), 2, gid)
		tree[k].Children = make([]Tree, len(children))
		for k1, v1 := range children {
			tree[k].Children[k1].Id = v1["Id"].(int64)
			tree[k].Children[k1].Text = v1["Title"].(string)
			// tree[k].Children[k1].Attributes.Url = "/" + v["Name"].(string) + "/" + v1["Name"].(string)
		}
	}
	return tree
}

func (this *CommonController) CommonMenu() {
	userInfo := this.GetSession("userinfo")
	if userInfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("auth_gateway"))
		return
	} else {
		role, _ := m.GetRoleByUserId(userInfo.(*m.User).Id)
		tree := this.GetResList(userInfo.(*m.User).Username, role.Id)
		treearr := strings.Split(this.Ctx.Input.URI(), "/")
		this.Data["treeurl"] = treearr[2]
		this.Data["tree"] = &tree
	}
	this.Data["serverurl"] = beego.AppConfig.String("localServerAdress")
	this.Layout = "haoadmin/layout/base.html"
}

func init() {
	//验证权限
	m.AccessRegister()
}
