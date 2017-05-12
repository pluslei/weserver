package haoadmin

import (
	"weserver/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CompanyController struct {
	CommonController
}

//公司列表
func (this *CompanyController) Index() {
	if this.IsAjax() {
		user := this.GetSession("userinfo").(*models.User)
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")

		if err != nil {
			beego.Error(err)
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Error(err)
		}
		companys, count := models.GetCompanys(iStart, iLength)
		// json
		data := make(map[string]interface{})
		data["aaData"] = companys
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		beego.Info("companys: ", companys[0])
		this.ServeJSON()
	} else {
		this.CommonMenu()
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
		company.CompanyIntro = companyIntro
		
		_, err := models.AddCompany(company)
		if err != nil {
			this.AlertBack("添加失败")
			return
		}
		this.Alert("添加成功", "company")
	} else {
		this.CommonMenu()
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
	this.Rsp(true, "删除成功", "")
}

func (this *CompanyController) EditCompany() {
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Debug("get id error", err)
		this.AlertBack("获取公司信息失败")
		return
	}
	if action == "edit" {
		var company = make(orm.Params)
		company["Company"] = this.GetString("companyName")
		company["CompanyIntro"] = this.GetString("companyIntro")
		company["CompanyIcon"] = this.GetString("CompanyIconFile")
		company["CompanyBanner"] = this.GetString("CompanyBannerFile")

		_, err = models.UpdateCompanyInfo(id, company)
		if err != nil {
			beego.Error("inser faild", err)
			this.AlertBack("修改失败")
			return
		} else {
			this.Alert("修改成功", "/weserver/data/company")
		}
	}
	companyInfo, _ := models.GetCompanyInfoById(id)
	this.CommonMenu()
	this.Data["companyInfo"] = companyInfo
	this.TplName = "haoadmin/data/company/edit.html"
}