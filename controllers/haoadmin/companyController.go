package haoadmin

import (
	"strconv"
	"weserver/models"
	. "weserver/src/cache"

	"github.com/astaxie/beego"
)

type CompanyController struct {
	CommonController
}

//公司列表
func (this *CompanyController) Index() {
	user := this.GetSession("userinfo").(*models.User)
	if user == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
		return
	}
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
		companys, count := models.GetCompanys(iStart, iLength, user.CompanyId)
		for _, item := range companys {
			roleInfo, err := models.GetRoleInfoById(item["Registerrole"].(int64))
			if err != nil {
				item["Rolename"] = "未知角色"
			} else {
				item["Rolename"] = roleInfo.Title
			}
			if user.CompanyId != 0 {
				item["AppId"] = ""
				item["AppSecret"] = ""
				item["Url"] = ""
			}
		}

		// json
		data := make(map[string]interface{})
		data["aaData"] = companys
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonMenu()
		if user.CompanyId == 0 {
			this.Data["show"] = true
		} else {
			this.Data["show"] = false
		}
		this.TplName = "haoadmin/data/company/index.html"
	}
}

// 添加公司
func (this *CompanyController) AddCompany() {
	action := this.GetString("action")
	if action == "add" {
		companyName := this.GetString("companyName")
		companyIntro := this.GetString("companyIntro")
		companyIcon := this.GetString("companyIconFile")
		companyBanner := this.GetString("companyBannerFile")
		companyLoginIcon := "./static/img/login/default-logo.png"
		companyLoginBackicon := "./static/img/login/default-login-bg.png"
		Registerrole, _ := this.GetInt64("Registerrole")
		WelcomeMsg := this.GetString("welcomemsg")
		HistoryMsg, _ := this.GetInt64("historymsg")
		AuditMsg, _ := this.GetInt64("auditmsg")
		Verify, _ := this.GetInt64("verify")
		AppId := this.GetString("AppId")
		AppSecret := this.GetString("AppSecret")
		Url := this.GetString("Url")
		if len(AppId) <= 0 {
			beego.Debug("AppId不能为空")
		}
		if len(AppSecret) <= 0 {
			beego.Debug("App密钥不能为空")
		}
		if len(Url) <= 0 {
			beego.Debug("Url不能为空")
		}

		if len(companyName) <= 0 {
			beego.Debug("companyName不能为空！")
		}
		if len(companyIntro) <= 0 {
			beego.Debug("companyIntro不能为空！")
		}
		if len(companyIcon) <= 0 {
			beego.Debug("companyIcon不能为空！")
		}
		if len(companyBanner) <= 0 {
			beego.Debug("companyBanner不能为空！")
		}
		company := new(models.Company)
		company.Company = companyName
		company.CompanyBanner = companyBanner
		company.CompanyIcon = companyIcon
		company.LoginIcon = companyLoginIcon
		company.LoginBackicon = companyLoginBackicon
		company.CompanyIntro = companyIntro
		company.Registerrole = Registerrole
		company.WelcomeMsg = WelcomeMsg
		company.HistoryMsg = HistoryMsg
		company.AuditMsg = AuditMsg
		company.Verify = Verify
		company.AppId = AppId
		company.AppSecret = AppSecret
		company.Url = Url

		_, err := models.AddCompany(company)
		if err != nil {
			this.AlertBack("添加失败")
			return
		}
		GetCompanyCache()
		this.Alert("添加成功", "company")
	} else {
		//获取所有角色
		role, err := models.GetAllUserRole()
		if err != nil {
			beego.Error(err)
		}
		title := models.TitleList()
		this.CommonMenu()
		this.Data["role"] = role
		this.Data["title"] = title
		this.TplName = "haoadmin/data/company/add.html"
	}

}

//删除公司
func (this *CompanyController) DelCompany() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Debug("get id error", err)
		this.Rsp(false, "获取失败", "")
		return
	}
	_, err = models.DelCompanyById(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
	}
	strId := strconv.FormatInt(id, 10)
	_, ok := MapCache[strId]
	if ok {
		delete(MapCache, strId)
	}

	this.Rsp(true, "删除成功", "")
}

func (this *CompanyController) EditCompany() {
	user := this.GetSession("userinfo").(*models.User)
	if user == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
		return
	}
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Debug("get id error", err)
		this.AlertBack("获取公司信息失败")
		return
	}
	if action == "edit" {
		var company models.Company
		company.Company = this.GetString("companyName")
		company.CompanyIntro = this.GetString("companyIntro")
		company.CompanyIcon = this.GetString("CompanyIconFile")
		company.CompanyBanner = this.GetString("CompanyBannerFile")
		company.LoginIcon = this.GetString("LoginIconFile")
		company.LoginBackicon = this.GetString("LoginBackiconFile")
		beego.Info("LoginIcon:", company.LoginIcon, "LoginBackicon:", company.LoginBackicon)
		company.Registerrole, _ = this.GetInt64("Registerrole")
		company.WelcomeMsg = this.GetString("welcomemsg")
		company.HistoryMsg, _ = this.GetInt64("historymsg")
		company.AuditMsg, _ = this.GetInt64("auditmsg")
		company.Verify, _ = this.GetInt64("verify")
		company.Sign = this.GetString("Sign")
		company.AppId = this.GetString("AppId")
		company.AppSecret = this.GetString("AppSecret")
		company.Url = this.GetString("Url")
		_, err = models.UpdateCompanyInfo(id, company, user.CompanyId)
		if err != nil {
			beego.Error("inser faild", err)
			this.AlertBack("修改失败")
			return
		} else {
			GetCompanyCache()
			this.Alert("修改成功", "/weserver/company/company")
		}
	}
	this.CommonMenu()
	companyInfo, _ := models.GetCompanyInfoById(id)
	//获取所有角色
	role, err := models.GetAllUserRole()
	if err != nil {
		beego.Error(err)
	}
	title := models.TitleList()

	if user.CompanyId != 0 {
		companyInfo.AppId = ""
		companyInfo.AppSecret = ""
		companyInfo.Url = ""
	}
	this.Data["companyInfo"] = companyInfo
	this.Data["role"] = role
	this.Data["title"] = title
	this.TplName = "haoadmin/data/company/edit.html"
}
