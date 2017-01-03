package models

import (
	// "errors"
	// "github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"strings"
	. "weserver/src/tools"
	//"time"
)

//分组表
type Title struct {
	Id         int64
	Name       string `orm:"size(128)" form:"Name"  valid:"Required"`
	Css        string `orm:"size(128)" form:"Css"  valid:"Required"`
	Background string `orm:"size(128)" form:"Css"  valid:"Required"`
	Weight     int    `orm:"default(1)" form:"Weight" valid:"Range(1,2)"`       //权重
	Remark     string `orm:"null;size(255)" form:"Remark" valid:"MaxSize(255)"` //备注
	User       *User  `orm:"reverse(one)"`
}

func (g *Title) TableName() string {
	return "title"
}

func init() {
	orm.RegisterModel(new(Title))
}

//get title list
func GetTitlelist(page int64, page_size int64, sort string) (titles []orm.Params, count int64) {
	o := orm.NewOrm()
	title := new(Title)
	qs := o.QueryTable(title)
	qs.Limit(page_size, page).OrderBy(sort).Values(&titles)
	count, _ = qs.Count()
	return titles, count
}

func AddTitle(t *Title) (int64, error) {
	o := orm.NewOrm()
	title := new(Title)
	title.Name = t.Name
	title.Css = t.Css
	title.Background = t.Background
	title.Weight = t.Weight
	title.Remark = t.Remark
	id, err := o.Insert(title)
	return id, err
}

func (this *Title) UpdateTitle(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func DelTitleById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Title{Id: Id})
	return status, err
}

func TitleList() (titles []orm.Params) {
	o := orm.NewOrm()
	title := new(Title)
	qs := o.QueryTable(title)
	qs.Values(&titles)
	return titles
}

func ReadTitleById(gid int64) (Title, error) {
	o := orm.NewOrm()
	title := Title{Id: gid}
	err := o.Read(&title)
	if err != nil {
		return title, err
	}
	return title, nil
}

func GetTitleByUserId(userId int64) (titles []orm.Params, count int64) {
	o := orm.NewOrm()
	title := new(Title)
	count, _ = o.QueryTable(title).Filter("User__id", userId).Values(&titles)
	return titles, count
}

// 关联查询所有用户的头衔
func GetRelationTitle(userid int64) (titles []Title, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("title").Filter("Id", userid).All(&titles, "Id")
	return titles, err
}

// 获取所有的角色
func GetAllUserTitle() (titles []Title, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("title").All(&titles)
	return titles, err
}

// 增加头衔给用户
func AddUserTitle(userid int64, titleid int64) (int64, error) {
	o := orm.NewOrm()
	title := Title{Id: titleid}
	user := User{Id: userid}
	m2m := o.QueryM2M(&user, "Title")
	num, err := m2m.Add(&title)
	return num, err
}

// 删除用户的头衔
func DelUserTitle(userid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user_titles").Filter("user_id", userid).Delete()
	return err
}

//模拟的数据
func AddSimulatedSata(usertype string, count int, titleid int64, roleico string, netname []NetName, randnano rand.Source) (user []Usertitle) {
	// 模拟数据
	var (
		uname     string
		utitle    string
		uicon     string
		authorcss string
		Insort    int
	)
	usertype = strings.Replace(usertype, " ", "", -1) //去空格
	switch usertype {
	case "manager":
		utitle = "管理员"
		Insort = 10000
	case "assistant":
		utitle = "助理"
		Insort = 9000
	case "customer":
		utitle = "客服"
		Insort = 8000
	case "teacher":
		utitle = "讲师"
		Insort = 7000
	case "nl_supreme":
		utitle = "至尊会员"
		Insort = 6000
	case "nl_vip":
		utitle = "VIP会员"
		Insort = 5000
	case "nl_platinum":
		utitle = "铂金会员"
		Insort = 4000
	case "nl_gold":
		utitle = "黄金会员"
		Insort = 3000
	case "nl_silver":
		utitle = "白银会员"
		Insort = 2000
	case "nl_ordinary":
		utitle = "普通会员"
		Insort = 1000
	case "guest":
		utitle = "游客"
		Insort = 100
	default:
		Insort = 0
	}
	for i := 0; i < count; i++ {
		var userinfor Usertitle
		switch usertype {
		case "guest":
			{
				//后台获取随机名
				uname = NewRandName(6, randnano) //随机的用户名
				uicon = "/upload/usericon/icon.png"
				authorcss = "/upload/usertitle/visitor.png"
			}
		default:
			{
				//后台获取随机名
				randlen := len(netname)
				r := rand.New(randnano)
				n := r.Intn(randlen)
				//randname := GetRandName(randnano)
				uname = netname[n].Name //随机的用户名
				n = r.Intn(6) + 1
				if len(roleico) > 0 {
					uicon = fmt.Sprintf("upload/usertitle/%s", roleico)
				} else {
					uicon = fmt.Sprintf("/upload/usericon/%d.png", n)
				}
				if titleid > 0 {
					titlelist, _ := ReadTitleById(titleid)
					authorcss = fmt.Sprintf("/upload/usertitle/%s", titlelist.Css)
				} else {
					titlelis, _ := GetAllUserTitle()
					k := len(titlelis)
					if k > 0 {
						n := r.Intn(k)
						authorcss = fmt.Sprintf("/upload/usertitle/%s", titlelis[n].Css)
					}
				}
			}
		}
		/*
			userinfor.Uname = EncodeB64(uname)         //用户名
			userinfor.RoleName = EncodeB64(usertype)   //用户角色
			userinfor.Titlerole = EncodeB64(utitle)    //用户类型
			userinfor.UserIcon = EncodeB64(uicon)      //用户Icon
			userinfor.Authorcss = EncodeB64(authorcss) //用户头衔
		*/
		userinfor.Uname = uname         //用户名
		userinfor.RoleName = usertype   //用户角色
		userinfor.Titlerole = utitle    //用户类型
		userinfor.UserIcon = uicon      //用户Icon
		userinfor.Authorcss = authorcss //用户头衔
		userinfor.Insort = Insort       //排序
		checkeddata := true
		urolelen := len(user)
		for k := 0; k < urolelen; k++ {
			if user[k].Uname == userinfor.Uname {
				checkeddata = false
				break
			}
		}
		if checkeddata {
			user = append(user, userinfor)
		}
		//time.Sleep(time.Millisecond * 1)
	}
	return user
}
