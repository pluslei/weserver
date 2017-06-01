package haoadmin

import (
	"strconv"
	"strings"
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
		nickname := this.GetString("sSearch_3")
		SearchId := this.GetString("sSearch_0")
		Room := this.GetString("sSearch_1")
		questionList, count := models.GetQuestionRecordList(iStart, iLength, "-Id", nickname, user.CompanyId, SearchId, Room)
		for _, item := range questionList {
			rspInfo, err := models.GetRspByQuestionId(item["Id"].(int64))
			if err == nil {
				rspContent := rspInfo.Content
				rspId := rspInfo.Id
				item["AcceptContent"] = rspContent //回复的内容
				item["AcceptNickName"] = rspInfo.Nickname
				item["AcceptUserIcon"] = rspInfo.UserIcon
				item["AcceptTitleName"] = rspInfo.RoleTitle
				item["rspId"] = rspId
			} else {
				item["AcceptContent"] = ""
				item["AcceptNickName"] = ""
				item["AcceptUserIcon"] = ""
			}
			//获取公司个房间信息
			roomInfo, err := models.GetRoomInfoByRoomID(item["Room"].(string))
			if err != nil {
				item["RoomName"] = "未知房间"
			} else {
				item["RoomName"] = roomInfo.RoomTitle
			}

			companyInfo, err := models.GetCompanyById(item["CompanyId"].(int64))
			if err != nil {
				item["CompanyName"] = "未知公司"
			} else {
				item["CompanyName"] = companyInfo.Company
			}
			beego.Info("RoomName:", item["RoomName"])
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
		TeacherInfo := this.GetString("TeacherInfo")
		var UserIcon string
		var TitleName string
		var NickName string
		if len(TeacherInfo) <= 0 {
			UserIcon = user.UserIcon
			TitleName = user.Title.Name
			NickName = user.Nickname
		} else {
			TeacherInfos := strings.Split(TeacherInfo, ",")
			if len(TeacherInfos) >= 3 {
				UserIcon = TeacherInfos[0]
				TitleName = TeacherInfos[1]
				NickName = TeacherInfos[2]
			} else {
				UserIcon = user.UserIcon
				TitleName = user.Title.Name
				NickName = user.Nickname
			}
		}

		//获取回复者的信息
		nowTime := time.Now()
		question, _ := models.GetQuestionIdData(QuestionId)
		rspQuestion := new(models.RspQuestion)
		rspQuestion.Content = AcceptContent
		rspQuestion.Uuid = question.Uuid
		rspQuestion.Time = nowTime
		rspQuestion.CompanyId = user.CompanyId
		rspQuestion.Uname = user.Username
		rspQuestion.Nickname = NickName
		rspQuestion.UserIcon = UserIcon
		rspQuestion.Room = question.Room
		//获取title_css, title_background
		roleInfo := user.Role
		title := user.Title
		rspQuestion.Sendtype = "TXT"
		rspQuestion.RoleTitle = TitleName
		rspQuestion.RoleName = roleInfo.Name
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
	Question, err := models.GetQuestionIdData(qid)
	if err != nil {
		beego.Error("inser faild", err)
		this.AlertBack("回复失败")
		return
	}
	//获取选择回复的头衔
	roomId := this.GetString("room")
	var infoMsg []models.Regist
	teacher, _, err := models.GetRegistInfoByRole(Question.CompanyId, roomId)
	if err != nil {
		beego.Debug("Get CompanyInfo Error", err)
		return
	}
	for _, v := range teacher {
		var info models.Regist
		info.UserIcon = v.UserIcon
		v.Titlename, err = models.GetTitleName(v.Title.Id)
		info.Titlename = v.Titlename
		info.Nickname = v.Nickname
		infoMsg = append(infoMsg, info)
	}
	this.CommonMenu()
	this.Data["qid"] = qid
	this.Data["TeacherInfo"] = infoMsg
	this.Data["Question"] = Question
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

// 批量删除纸条提问
func (this *QuestionController) QuestionDels() {
	IdArray := this.GetString("Id")
	beego.Info("IdArray:", IdArray)
	var idarr []int64
	if len(IdArray) > 0 {
		preValue := strings.Split(IdArray, ",")
		for _, v := range preValue {
			id, _ := strconv.ParseInt(v, 10, 64)
			idarr = append(idarr, id)

		}
	}
	status, err := models.PrepareDelQuestion(idarr)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}
