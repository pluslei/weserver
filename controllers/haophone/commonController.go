package haophone

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"weserver/controllers"
	m "weserver/models"
)

type CommonController struct {
	controllers.PublicController
}

// 公共方法
func (this *CommonController) CommonMenu() {
	userInfo := this.GetSession("indexUserInfo")
	if userInfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("auth_gateway"))
	} else {

		var uname string
		var unameid int64
		var roleid int64
		uname = fmt.Sprintf("%s", userInfo.(*m.User).Username)
		unameid, _ = strconv.ParseInt(fmt.Sprintf("%d", userInfo.(*m.User).Id), 10, 64)
		role, _ := m.GetRoleByUserId(unameid)
		roleid = role.Id
		functree := this.GetFuncList(uname, roleid)
		this.Data["json"] = &functree
	}
}

func (this *CommonController) GetFuncList(uname string, Id int64) []orm.Params {
	//var length int = 0
	var resources []orm.Params
	// adminUser := p.Adminuser
	if uname == "admin" {
		resources, _ = m.GetNodeGroupWebTree()
	} else {
		resources, _ = m.GetResourcesByRoleId(Id)
	}
	return resources

}

// 跳转第一个房间
func (this *CommonController) RedirectRoom() {
	roomInfo, err := m.GetFristerRoom()
	if err != nil {
		this.Data["content"] = "正在努力创建中..."
		this.Data["err"] = "404 Error"
		this.TplName = "haoindex/404.html"
		return
	}
	this.Ctx.Redirect(301, "/"+fmt.Sprintf(`%d`, roomInfo.RommNumber))
}

//检查用户是否有权限
func init() {
	var Check = func(ctx *context.Context) {
		user_auth_type, _ := strconv.Atoi(beego.AppConfig.String("user_auth_type"))
		rbac_auth_gateway := beego.AppConfig.String("rbac_auth_gateway")
		var accesslist map[string]bool
		if user_auth_type > 0 {
			params := strings.Split(strings.ToLower(ctx.Request.URL.Path), "/")
			if m.CheckAccess(params) {
				uinfo := ctx.Input.Session("indexUserInfo")
				if uinfo == nil {
					ctx.Redirect(302, rbac_auth_gateway)
				}
				//admin用户不用认证权限
				adminuser := beego.AppConfig.String("rbac_admin_user")
				if uinfo.(*m.User).Username == adminuser {
					return
				}
				//
				if user_auth_type == 1 {
					listbysession := ctx.Input.Session("accesslist")
					if listbysession != nil {
						accesslist = listbysession.(map[string]bool)
					}
				} else if user_auth_type == 2 {
					accesslist, _ = m.GetAccessList(uinfo.(*m.User).Id)
				}

				ret := m.AccessDecision(params, accesslist)
				if !ret {
					ctx.Redirect(302, rbac_auth_gateway)
					ctx.Output.JSON(&map[string]interface{}{"status": false, "info": "权限不足"}, true, false)
				}
			}

		}
	}
	beego.InsertFilter("/socket/*", beego.BeforeRouter, Check)
}
