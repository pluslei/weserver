package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"weserver/src/tools"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//
func AccessRegister() {
	var Check = func(ctx *context.Context) {
		user_auth_type, _ := strconv.Atoi(beego.AppConfig.String("user_auth_type"))
		rbac_auth_gateway := beego.AppConfig.String("rbac_auth_gateway")
		var accesslist map[string]bool
		if user_auth_type > 0 {
			pathUrl := strings.ToLower(ctx.Request.URL.Path)
			beego.Debug("pathUrl", len(pathUrl))
			if len(pathUrl) < 3 {
				ctx.Redirect(302, rbac_auth_gateway)
			}
			params := strings.Split(pathUrl[1:], "/")
			uinfo := ctx.Input.Session("userinfo")
			if uinfo == nil {
				ctx.Redirect(302, rbac_auth_gateway)
				return
			}
			//admin不需要验证
			adminuser := beego.AppConfig.String("rbac_admin_user")
			if uinfo.(*User).Username == adminuser {
				return
			}
			accesslist, _ = GetAccessList(uinfo.(*User).Id)
			ret := AccessDecision(params, accesslist)

			if ret == false {
				ctx.Redirect(302, rbac_auth_gateway)
				ctx.Output.JSON(&map[string]interface{}{"status": false, "info": "权限不足"}, true, false)
			} else {
				return
			}
		}
	}
	beego.InsertFilter("/weserver/*", beego.BeforeRouter, Check)
}

//Determine whether need to verify
func CheckAccess(params []string) bool {
	for _, nap := range strings.Split(beego.AppConfig.String("not_auth_package"), ",") {
		if params[1] == nap {
			return false
		}
	}
	return true
}

//To test whether permissions
func AccessDecision(params []string, accesslist map[string]bool) bool {
	if len(params) < 3 {
		return false
	}
	beego.Debug("params", params, CheckAccess(params))
	if CheckAccess(params) {
		s := fmt.Sprintf("/%s/%s/%s", params[0], params[1], params[2])
		if len(accesslist) < 1 {
			return false
		}
		_, ok := accesslist[s]
		if ok != false {
			return true
		}
	} else {
		return true
	}
	return false
}

type AccessNode struct {
	Id        int64
	Name      string
	Childrens []*AccessNode
}

//Access permissions list
func GetAccessList(uid int64) (map[string]bool, error) {
	list, err := AccessList(uid)
	if err != nil {
		return nil, err
	}
	alist := make([]*AccessNode, 0)
	for _, l := range list {
		if l["Pid"].(int64) == 0 && l["Level"].(int64) == 1 {
			anode := new(AccessNode)
			anode.Id = l["Id"].(int64)
			anode.Name = l["Url"].(string)
			alist = append(alist, anode)
		}
	}
	for _, l1 := range list {
		if l1["Level"].(int64) == 2 {
			for _, an := range alist {
				if an.Id == l1["Pid"].(int64) {
					anode := new(AccessNode)
					anode.Id = l1["Id"].(int64)
					anode.Name = l1["Url"].(string)
					an.Childrens = append(an.Childrens, anode)
				}
			}
		}
	}
	for _, l2 := range list {
		if l2["Level"].(int64) == 3 {
			for _, an := range alist {
				for _, an1 := range an.Childrens {
					if an1.Id == l2["Pid"].(int64) {
						anode := new(AccessNode)
						anode.Id = l2["Id"].(int64)
						anode.Name = l2["Url"].(string)
						an1.Childrens = append(an1.Childrens, anode)
					}
				}
			}
		}
	}

	accesslist := make(map[string]bool)
	var str string = ""
	for _, v := range alist {
		for _, v1 := range v.Childrens {
			str = "/" + v1.Name
			accesslist[str] = true
			for _, v2 := range v1.Childrens {
				//	for _, v2 := range v1.Childrens {
				// vname := strings.Split(v.Name, "/")
				// v1name := strings.Split(v1.Name, "/")
				//	v2name := strings.Split(v2.Name, "/")
				//	str := fmt.Sprintf("%s/%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[0]), strings.ToLower(v2name[0]))
				// str := fmt.Sprintf("%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[0]))
				// accesslist[str] = true
				//	}
				str = "/" + v2.Name
				accesslist[str] = true
			}
		}
	}
	return accesslist, nil

}

//check login
func CheckLogin(username string, password string) (user *User, err error) {
	var u User
	u.Username = username
	user, err = ReadFieldUser(&u, "Username")
	if user == nil || err != nil {
		beego.Error(err)
		return &u, err
	}
	if user.Id == 0 {
		return user, errors.New("The user is not exits.")
	}
	if user.Password != tools.EncodeUserPwd(username, password) {
		return user, errors.New("The password is error.")
	}
	return user, nil
}
