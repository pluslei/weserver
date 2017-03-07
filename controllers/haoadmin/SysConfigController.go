package haoadmin

import (
	"github.com/astaxie/beego"
	m "weserver/models"
)

type SysConfigController struct {
	CommonController
}

func (this *SysConfigController) Index() {

	action := this.GetString("action")
	if action == "edit" {
		systemname := this.GetString("systemname")         //系统标题
		chartinterval, _ := this.GetInt64("chatinterval")  //游客聊天间隔时间
		registertitle, _ := this.GetInt64("registertitle") //默认注册用户头衔
		registerrole, _ := this.GetInt64("registerrole")   //默认注册角色
		historymsg, _ := this.GetInt64("historymsg")       //是否显示历史消息
		historycount, _ := this.GetInt64("historycount")   //显示历史记录条数
		welcomemsg := this.GetString("welcomemsg")         //欢迎语
		verify, _ := this.GetInt64("verify")               //是否开启注册验证
		loginsys, _ := this.GetInt64("loginsys")           //是否允许登陆后台
		auditmsg, _ := this.GetInt64("auditmsg")           //是否消息审核
		virtualuser, _ := this.GetInt64("virtualuser")     //增加虚拟用户
		sys := new(m.SysConfig)
		sys.Systemname = systemname
		sys.ChatInterval = chartinterval
		sys.Registerrole = registerrole
		sys.Registertitle = registertitle
		sys.HistoryMsg = historymsg
		sys.HistoryCount = historycount
		sys.WelcomeMsg = welcomemsg
		sys.Verify = verify
		sys.LoginSys = loginsys
		sys.AuditMsg = auditmsg
		sys.VirtualUser = virtualuser
		count := m.GetSysConfigCount()
		if count == 0 {
			id, err := m.AddSysConfig(sys)
			if err != nil && id <= 0 {
				beego.Error(err)
				this.AlertBack("配置修改失败")
				return
			}
			this.Alert("添加成功", "index")
		} else {
			sysid, _ := this.GetInt64("Id")
			sys.Id = sysid
			err := sys.UpdateSysConfig("Systemname", "ChatInterval", "Registerrole", "Registertitle", "HistoryMsg", "HistoryCount", "WelcomeMsg", "Verify", "LoginSys", "AuditMsg", "VirtualUser")
			if err != nil {
				beego.Error(err)
				this.AlertBack("配置修改失败")
				return
			}
			this.Alert("修改成功", "index")
		}
	} else {
		this.CommonMenu()
		configinfo, err := m.ReadSysConfigById(1)
		if err != nil {
			beego.Error(err)
		}
		title := m.TitleList()
		role, err := m.GetAllUserRole()
		if err != nil {
			beego.Error(err)
		}
		this.Data["title"] = title
		this.Data["role"] = role
		this.Data["configinfo"] = configinfo
		beego.Debug("configinfo", configinfo.Verify, configinfo.AuditMsg)
		this.TplName = "haoadmin/rbac/sysconfig/edit.html"
	}
}
