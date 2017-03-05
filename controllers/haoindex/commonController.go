package haoindex

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"strings"
	"weserver/controllers"
	m "weserver/models"
)

// 获取直播地址
var LiveUrl = beego.AppConfig.String("liveurl")

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
		// fmt.Println(&functree)
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

	// for _, v := range resources {
	// 	if v["Pid"].(int64) == 0 {
	// 		length = length + 1

	// 	}
	// }
	// tree := make([]string, length)
	// for _, v := range resources {
	// 	if v["Pid"].(int64) == 4 {
	// 		tree = append(tree, v["Title"].(string))

	// 	}
	// }
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
					return
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

// 播放参数
type Result struct {
	Retval  string
	Reterr  string
	Retinfo map[string]interface{}
}

// 获取播放地址公共方法
func (this *CommonController) GetLive() string {
	var result Result
	request := httplib.Get(LiveUrl)
	str, err := request.String()
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		fmt.Println("json error")
	}
	var strurl string
	if result.Retinfo["play_url"] != nil {
		strurl = result.Retinfo["play_url"].(string)
	} else {
		strurl = ""
	}
	return strurl
}

func (this *CommonController) GetPhoneLive() string {
	var phonestr string
	str := this.GetLive()
	if len(str) > 0 {
		phone := strings.Split(str, ":1935")
		phoneurl := phone[0] + phone[1]
		phonestr = strings.Replace(phoneurl, "rtmp://wsrtmp", "http://wshls", -1) + "/playlist.m3u8"
		beego.Debug("phonestr ======> ", phonestr)
	}
	return phonestr
}

func (this *CommonController) CheckUserIsAuth() bool {
	exist := this.GetSession("indexUserInfo")
	if exist == nil {
		return true
	}
	return false
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
