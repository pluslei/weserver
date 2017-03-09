package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//虚拟用户表
type VirtualUser struct {
	Id       int64
	Username string
	Nickname string `orm:"size(255)"`
	UserIcon string `orm:"null;size(255)"`
}

func (v *VirtualUser) TableName() string {
	return "virtual_user"
}

func init() {
	orm.RegisterModel(new(VirtualUser))
}

func GetNumberVirtualUser(count int64) (virtual []VirtualUser, err error) {
	_, err = NewVirtualUser().Limit(count).All(&virtual)
	return virtual, err
}

func NewVirtualUser() orm.QuerySeter {
	o := orm.NewOrm()
	return o.QueryTable(new(VirtualUser))
}

//最近 X 天 人员总列表信息
func VirtualUserList(nDay int64) (userlist []VirtualUser, count int) {
	onlineuser, err := GetAllUser(nDay)
	if err != nil {
		beego.Error("get the user error", err)
	} else {
		for _, item := range onlineuser {
			var user VirtualUser
			user.Username = item.Username
			user.Nickname = item.Nickname
			user.UserIcon = item.UserIcon
			userlist = append(userlist, user)
		}
	}
	sysconfig, _ := GetAllSysConfig()
	if sysconfig.VirtualUser > 0 {
		virtualUser, err := GetNumberVirtualUser(sysconfig.VirtualUser) // 获取虚拟表中的数据 由sysconfig中vitual 指定
		if err != nil {
			beego.Error("user count error", err)
		} else {
			for _, item := range virtualUser {
				userlist = append(userlist, item)
			}
		}
	}
	count = len(userlist)
	return userlist, count
}
