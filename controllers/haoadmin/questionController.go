package haoadmin

import (
	"weserver/models"

	"github.com/astaxie/beego"
	 "github.com/astaxie/beego/orm"
)

type QuestionController struct {
	CommonController
}

//纸条列表
func (this *QuestionController) QuestionList() {
	if this.IsAjax() {
		user := this.GetSession("userinfo").(*models.User)
		if user == nil {
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
		sEcho := this.GetString("sEcho")
		iStart, err := this.GetInt64("iDisplayStart")
		if err != nil {
			beego.Debug("userlist get", err)
			return
		}
		iLength, err := this.GetInt64("iDisplayLength")
		if err != nil {
			beego.Debug("userlist Error", err)
			return
		}
		nickname := this.GetString("sSearch_0")
		companyId := user.CompanyId

		questionList, count := models.GetQuestionRecordList(iStart, iLength, "-Id", nickname, companyId)
		// json
		data := make(map[string]interface{})
		data["aaData"] = questionList
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/question/index.html"
	}
}

//回复
func (this *QuestionController) QuestionReply() {
	user := this.GetSession("userinfo").(*models.User)
	if user == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
		return
	}
	action := this.GetString("action")
	id, err := this.GetInt64("id")
	beego.Info("id", id)
	if err != nil {
		beego.Debug("get id error", err)
		this.AlertBack("获取信息失败")
		return
	}
	if action == "reply" {
		var question = make(orm.Params)
		msgContent := this.GetString("AcceptContent")
		question["AcceptContent"] = msgContent
		question["CompanyIntro"] = this.GetString("companyIntro") //回复者id 这里是老师的id
		_, err = models.QuestionReply(id, msgContent)
		if err != nil {
			beego.Error("inser faild", err)
			this.AlertBack("回复失败")
			return
		} else {
			this.Alert("回复成功", "/weserver/data/question")
			questionInfo, _ := models.GetQuestionIdData(id) //获取纸条信息
			beego.Info("AcceptUuid: ", questionInfo.AcceptUuid)
		}
	}
	questionInfo, _ := models.GetQuestionIdData(id)
	//获取所有角色
	role, err := models.GetAllUserRole()
	if err != nil {
		beego.Error(err)
	}
	title := models.TitleList()
	this.CommonMenu()

	if user.CompanyId == 0 {
		this.Data["show"] = true
	} else {
		this.Data["show"] = false
	}

	this.Data["questionInfo"] = questionInfo
	this.Data["role"] = role
	this.Data["title"] = title
	this.TplName = "haoadmin/data/question/reply.html"
}

//删除纸条提问
func (this *QuestionController) QuestionDel() {
		id, err := this.GetInt64("id")
	if err != nil {
		beego.Debug("get id error", err)
		this.Rsp(false, "获取失败", "")
		return
	}
	_, err = models.DeleteById(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
	}
	this.Rsp(true, "删除成功", "")
}