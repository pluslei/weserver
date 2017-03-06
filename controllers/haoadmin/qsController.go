package haoadmin

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"time"
	m "weserver/models"
	p "weserver/src/parameter"
	. "weserver/src/tools"
)

type QsController struct {
	CommonController
}

func (this *QsController) Index() {
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
		qslist, count := m.GetQslist(iStart, iLength, "Id")

		// json
		data := make(map[string]interface{})
		data["aaData"] = qslist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/qs/list.html"
	}

}

// 增加组别
func (this *QsController) AddQs() {
	Question := this.GetString("Question")
	Answer := this.GetString("Answer")
	Weight, _ := this.GetInt("Weight")
	if len(Question) > 0 && len(Answer) > 0 && Weight > 0 {
		q := new(m.Qs)
		q.Question = Question
		q.Answer = Answer
		q.Weight = Weight
		id, err := m.AddQs(q)
		if err != nil && id <= 0 {
			beego.Error(err)
			this.AlertBack("一问一答添加失败")
			return
		}
		this.Alert("添加成功", "qs_index")
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/qs/add.html"
	}

}

func (this *QsController) UpdateQs() {
	Question := this.GetString("Question")
	Answer := this.GetString("Answer")
	Weight, _ := this.GetInt("Weight")
	Id, _ := this.GetInt64("Id")
	if len(Question) > 0 && len(Answer) > 0 && Weight > 0 && Id > 0 {
		q := new(m.Qs)
		q.Id = Id
		q.Question = Question
		q.Answer = Answer
		q.Weight = Weight
		err := q.UpdateQs("Question", "Answer", "Weight")
		if err != nil {
			beego.Error(err)
			this.AlertBack("问答修改失败")
			return
		}
		this.Alert("修改成功", "qs_index")
	} else {
		this.CommonMenu()
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.AlertBack("问答获取失败")
			return
		}
		qsList, err := m.ReadQsById(id)
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取头衔信息错误")
			return
		}
		this.Data["qsList"] = qsList
		this.TplName = "haoadmin/data/qs/edit.html"
	}

}

func (this *QsController) DelQs() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelQsById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

func (this *QsController) DataBroadQs() {
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
		Broadlist, count := m.GetBroadcastlist(iStart, iLength, "Room")
		for _, item := range Broadlist {
			item["Datatime"] = item["Datatime"].(time.Time).Format("2006-01-02 15:04:05")
		}
		// json
		data := make(map[string]interface{})
		data["aaData"] = Broadlist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.CommonController.CommonMenu()
		this.TplName = "haoadmin/data/qs/databroad.html"
	}
}

// 发送广播
func (this *QsController) SendBroad() {
	prevalue := beego.AppConfig.String("company") + "_" + beego.AppConfig.String("room")
	codeid := MainEncrypt(prevalue)
	this.Data["codeid"] = codeid
	if this.GetSession("userinfo") != nil {
		UserInfo := this.GetSession("userinfo")
		this.Data["uname"] = UserInfo.(*m.User).Username
	}
	this.Data["ipaddress"] = this.GetClientip()
	this.Data["serverurl"] = beego.AppConfig.String("localServerAdress")
	this.TplName = "haoadmin/data/qs/sendbroad.html"
}

//机器人发言
func (this *QsController) DataRobotSpeak() {
	if this.GetSession("userinfo") != nil {
		if this.IsAjax() {
			var (
				robotlist []Usertitle //模拟的用户数据
				urole     []Usertitle //获取所有的房间号
				roomcode  []string    //获取所有房间号及公司代码
			)
			//在线用户数据
			roominfo, num, _ := m.GetAllRoomDate()
			roomlen := int(num)
			//roominfo, _, _ := m.GetAllRoomDate()
			for i := 0; i < roomlen; i++ {
				//用户列表信息
				userroom := make(map[string]Usertitle) //房间对应的用户信息
				jobroom := "coderoom_" + p.Code + "_" + fmt.Sprintf("%d", roominfo[i].RommNumber)
				roomdata, _ := p.Client.Get(jobroom)
				if len(roomdata) > 0 {
					userroom, _ = Jsontoroommap(roomdata)
				}
				for _, userId := range userroom {
					if len(userId.Uname) > 0 {
						urole = append(urole, userId)
					}
				}
				//获取所有房间号及公司代码
				companyid := roominfo[i].CompanyCode
				temp := fmt.Sprintf("%d", roominfo[i].RommNumber)
				prevalue := companyid + "_" + temp
				codeid := MainEncrypt(prevalue)
				roomcode = append(roomcode, codeid)
			}
			//水军数据
			robotlist = Resultuser
			// json
			data := make(map[string]interface{})
			data["shuijun"] = robotlist
			data["roomcode"] = roomcode
			data["urole"] = urole
			this.Data["json"] = &data
			this.ServeJSON()
		} else {
			UserInfo := this.GetSession("userinfo")
			// 获取所有的分组
			groups, _ := m.GetGroupList()
			length := len(groups)
			for i := 0; i < length; i++ {
				groups[i].GroupFace = FaceImg + groups[i].GroupFace
			}
			this.CommonMenu()
			this.Data["uname"] = UserInfo.(*m.User).Username
			this.Data["group"] = groups
			this.TplName = "haoadmin/shuijun.html"
		}
	} else {
		this.Rsp(false, "非法操作", "")
	}
}

//获取客户的真是IP地址
func (this *QsController) GetClientip() string {
	var addrArr []string
	if len(this.Ctx.Request.Header.Get("X-Forwarded-For")) > 0 {
		addr := this.Ctx.Request.Header.Get("X-Forwarded-For")
		addrArr = strings.Split(addr, ":")
	} else if len(this.Ctx.Request.RemoteAddr) > 0 {
		addr := this.Ctx.Request.RemoteAddr
		addrArr = strings.Split(addr, ":")
	} else {
		addrArr[0] = "127.0.0.1"
	}
	return addrArr[0]
}
