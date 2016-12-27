package models

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//分组表
type SysConfig struct {
	Id            int64
	Systemname    string //直播室名称
	Redirect      string //首页跳转
	Imagesize     int64  //图片大小 单位Kb
	Imagetype     string //图片类型
	Guestchat     int64  //0 禁止聊天 1 允许聊天
	ChatInterval  int64  //0 无间隔  其它数字为间隔时间（秒）
	Registerrole  int64  //默认注册用户角色
	Registertitle int64  //默认注册用户头衔
	HistoryMsg    int64  //是否显示历史消息 0显示  1 不显示
	HistoryCount  int64  //显示历史记录条数
	WelcomeMsg    string //欢迎语
	Verify        int64  //是否开启验证  0开启 1不开启
	LoginSys      int64  //是否允许登陆后台  0允许 1禁止
	AuditStatus   int    `orm:"default(2)" form:"AuditStatus" valid:"Required;Range(1,2)"` //1：不审核，2：需要审核
}

func (s *SysConfig) TableName() string {
	return "sys_config"
}

func init() {
	orm.RegisterModel(new(SysConfig))
}

//get title list
func GetSysConfiglist(page int64, page_size int64, sort string) (configs []orm.Params, count int64) {
	o := orm.NewOrm()
	config := new(SysConfig)
	qs := o.QueryTable(config)
	qs.Limit(1).OrderBy(sort).Values(&configs)
	count, _ = qs.Count()
	return configs, count
}

func GetSysConfigCount() int64 {
	o := orm.NewOrm()
	config := new(SysConfig)
	sys := o.QueryTable(config)
	count, _ := sys.Count()
	return count
}

func GetAllSysConfig() (sys SysConfig, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(sys).One(&sys, "Imagesize", "Imagetype", "Guestchat", "ChatInterval", "HistoryMsg", "HistoryCount", "WelcomeMsg", "AuditStatus")
	return sys, err
}

func AddSysConfig(c *SysConfig) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(c)
	return id, err
}

func (this *SysConfig) UpdateSysConfig(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func ReadSysConfigById(id int64) (SysConfig, error) {
	o := orm.NewOrm()
	config := SysConfig{Id: id}
	err := o.Read(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func GetSysConfig() (sysconfig SysConfig, err error) {
	model := orm.NewOrm()
	err = model.QueryTable("sys_config").One(&sysconfig)
	return sysconfig, err
}
