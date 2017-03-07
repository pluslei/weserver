package models

import (
	//"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	//"weserver/src/tools"
)

//用户表
type User struct {
	Id            int64
	Username      string `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(32);MinSize(6)"`
	Password      string `orm:"size(32)" form:"Password" valid:"Required;MaxSize(32);MinSize(6)"`
	Repassword    string `orm:"-" form:"Repassword" valid:"Required"`
	Nickname      string `orm:"size(255)" form:"Nickname" valid:"Required;MaxSize(255);MinSize(2)"`
	Email         string `orm:"unique;size(32)" form:"Email" valid:"Email"`
	Phone         int64  `orm:"unique;size(11)" form:"Phone" valid:"MaxSize(11);MinSize(1)"`
	Qq            int64
	Remark        string    `orm:"null;size(255)" form:"Remark" valid:"MaxSize(255)"`
	Status        int       `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //用户注册状态 1为未审核 2为审核通过
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	UserIcon      string    `orm:"null;size(255)" form:"UserIcon" valid:"MaxSize(255)"`
	RegStatus     int       `orm:"default(1)" form:"Status" valid:"Range(1,2)"` //用户注册状态 1为未审核 2为审核通过
	OnlineTime    int64     //用户在线时长
	Openid        string    //用户的唯一标识
	Sex           int32     //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province      string    //用户个人资料填写的省份
	City          string    //普通用户个人资料填写的城市
	Country       string    //	国家，如中国为CN
	Headimgurl    string    //用户头像最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Unionid       string    //只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：
	Role          *Role     `orm:"rel(one)"`
	Title         *Title    `orm:"rel(one)"`

	LogintimeStr  string `orm:"-"` //登录时间
	OnlinetimeStr string `orm:"-"` //在线时长
	Ipaddress     string `orm:"-"` //ip地址
	Titlename     string `orm:"-"` //头衔名称
}

func (u *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

//get user list
func Getuserlist(page int64, page_size int64, sort, nickname string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user).Exclude("Username", "admin")
	qs.Limit(page_size, page).Filter("nickname__contains", nickname).OrderBy(sort).RelatedSel().Values(&users)
	count, _ = qs.Count()
	return users, count
}

//添加用户
func AddUser(u *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	return id, err
}

// 删除用户
func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

//批量删除用户
func PrepareDelUser(IdArray []int64) (int64, error) {
	o := orm.NewOrm()
	err := o.Begin()
	var status int64
	for i := 0; i < len(IdArray); i++ {
		status, err = o.Delete(&User{Id: IdArray[i]})
	}
	// 此过程中的所有使用 o Ormer 对象的查询都在事务处理范围内
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return status, err
}

// 根据用户名查找
func GetUserByUsername(username string) (user User, err error) {
	user = User{Username: username}
	o := orm.NewOrm()
	err = o.Read(&user, "Username")
	return user, err
}

// 查询用户名和手机号是否存在
func GetUserNameByPhone(name string, phone int64) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", name).Filter("Phone", phone).One(&user)
	return user, err
}

// 获取用户一对一关系
func LoadRelatedUser(u *User, fields ...string) (*User, error) {
	o := orm.NewOrm()
	beego.Debug("userInfo", u)
	err := o.Read(u, fields...)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	_, err = o.LoadRelated(u, "Role")
	_, err = o.LoadRelated(u, "Title")
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return u, nil
}

// 获取用户
func ReadFieldUser(u *User, fields ...string) (*User, error) {
	o := orm.NewOrm()
	err := o.Read(u, fields...)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return u, nil
}

func (this *User) UpdateUserFields(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

// 查询会员人数
func GetUserNumber() (int64, error) {
	model := orm.NewOrm()
	var table User
	number, err := model.QueryTable(table).Count()
	return number, err
}

func GetRegStatusUser(statu int, page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	qs.Limit(page_size, page).OrderBy(sort).Filter("RegStatus", 1).Values(&users)
	count, _ = qs.Count()
	return users, count
}

func GetUserByOnlineDesc() (users []User, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("user").OrderBy("-OnlineTime").Filter("Role", 5).Limit(5).All(&users)
	return users, err
}

func CountOnline() string {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("select SUM(online_time) as onlineTime from user").Values(&maps)
	if err == nil && num > 0 {
		//	maps[0]["online_time"] // slene
	}
	count := maps[0]["onlineTime"]
	online := count.(string)
	return online

}

func CountWeekRegist() (week []orm.ParamsList) {
	o := orm.NewOrm()
	var lists []orm.ParamsList
	num, err := o.Raw("SELECT date_format(createtime, '%Y-%m-%d') AS createtime, count(*) as count FROM USER WHERE DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= date(createtime) GROUP BY date_format(createtime, '%Y-%m-%d') ORDER BY createtime desc").ValuesList(&lists)
	if err == nil && num > 0 {

	}
	return lists
}

func GetAllUser() (users []User, err error) {
	o := orm.NewOrm()
	nowtime := time.Now().Unix() - 3*24*60*60
	_, err = o.QueryTable("user").Filter("Lastlogintime__gte", time.Unix(nowtime, 0).Format("2006-01-02 15:04:05")).All(&users)
	return users, err
}
