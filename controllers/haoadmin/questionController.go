package haoadmin

import (
	"weserver/models"

	"time"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
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
		for _, item := range questionList {
			rspInfo, err := models.GetRspByQuestionId(item["Id"].(int64))
			if err == nil {
				rspContent := rspInfo.Content
				rspId := rspInfo.Id
				item["AcceptContent"] = rspContent //回复的内容
				item["rspId"] = rspId
			} else {
				item["AcceptContent"] = ""
			}
		}
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

func (this *QuestionController) QuestionReply() {

	user := this.GetSession("userinfo").(*models.User)
	if user == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
		return
	}
	qid, _ := this.GetInt64("questionId")
	action := this.GetString("action")
	if action == "reply" {
		AcceptContent := this.GetString("AcceptContent") //回复内容
		QuestionId := qid
		if len(AcceptContent) <= 0 {
			beego.Debug("AcceptContent不能为空")
		}
		if QuestionId <= 0 {
			beego.Debug("QuestionId不能为空")
		}
		//获取回复者的信息
		nowTime := time.Now()
		question, _ := models.GetQuestionIdData(QuestionId)
		rspQuestion := new(models.RspQuestion)
		rspQuestion.Content = AcceptContent
		rspQuestion.Uuid = question.Uuid
		rspQuestion.Time = nowTime
		rspQuestion.CompanyId = user.CompanyId
		rspQuestion.Uname = user.Openid
		rspQuestion.Nickname = user.Nickname
		rspQuestion.UserIcon = user.UserIcon
		rspQuestion.Room = question.Room
		rspQuestion.RoleName = user.Rolename
		rspQuestion.RoleTitle = user.Titlename
		//获取title_css, title_background
		title := user.Title
		rspQuestion.RoleTitleCss = title.Css
		rspQuestion.RoleTitleBack = title.Background
		rspQuestion.DatatimeStr = time.Unix(nowTime.Unix(), 0).Format("2006-01-02 15:04:05")
		rspQuestion.Question = &question
		rspId, err := models.AddRspQuestion(rspQuestion)
		if err != nil {
			beego.Error("inser faild", err)
			this.AlertBack("回复失败")
			return
		} else {
			this.Alert("回复成功", "/weserver/data/question")
			questionInfo, _ := models.GetQuestionIdData(QuestionId)  //获取纸条提问信息
			repquestionInfo, _ := models.GetRspQuestionIdData(rspId) //获回复信息
			beego.Info(questionInfo, repquestionInfo)

		}
	}
	this.CommonMenu()
	this.Data["qid"] = qid
	this.TplName = "haoadmin/data/question/reply.html"
}

//删除回复
func (this *QuestionController) QuestionDel() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Debug("get id error", err)
		this.Rsp(false, "获取失败", "")
		return
	}

	_, err = models.DeleteRspQuestionById(id)
	if err != nil {
		this.Rsp(false, "删除失败", "")
	}
	this.Rsp(true, "删除成功", "")
}
