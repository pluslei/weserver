package haoadmin

import (
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	m "weserver/models"
)

type CourseController struct {
	CommonController
}

func (this *CourseController) Index() {
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
		courselist, count := m.GetCourseList(iStart, iLength, "Id")
		// json
		data := make(map[string]interface{})
		data["aaData"] = courselist
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data

		this.ServeJSON()

	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/course/list.html"
	}

}

func (this *CourseController) AddCourse() {
	CourseData := this.GetString("CourseData")
	if len(CourseData) > 0 {
		// fmt.Println("%qn",)
		c := new(m.Course)
		array := strings.Split(CourseData, "-")
		for i := 0; i < len(array)-1; i++ {
			course := array[i]
			coursearr := strings.Split(course, "|")
			//for k := 0; k < len(coursearr); k++ {
			Id, _ := strconv.ParseInt(coursearr[0], 10, 64)
			StartTime := coursearr[1]
			EndTime := coursearr[2]
			CourseName := coursearr[3]
			Monday, _ := strconv.ParseInt(coursearr[4], 10, 32)
			Tuesday, _ := strconv.ParseInt(coursearr[5], 10, 32)
			Wednesday, _ := strconv.ParseInt(coursearr[6], 10, 32)
			Thursday, _ := strconv.ParseInt(coursearr[7], 10, 32)
			Friday, _ := strconv.ParseInt(coursearr[8], 10, 32)
			Saturday, _ := strconv.ParseInt(coursearr[9], 10, 32)
			Sunday, _ := strconv.ParseInt(coursearr[10], 10, 32)
			Uname1 := coursearr[11]
			Uname2 := coursearr[12]
			Uname3 := coursearr[13]
			Uname4 := coursearr[14]
			Uname5 := coursearr[15]
			Uname6 := coursearr[16]
			Uname7 := coursearr[17]
			Uname := Uname1 + "," + Uname2 + "," + Uname3 + "," + Uname4 + "," + Uname5 + "," + Uname6 + "," + Uname7
			c.Name = CourseName
			c.StartTime = StartTime
			c.EndTime = EndTime
			c.Monday = int(Monday)
			c.Tuesday = int(Tuesday)
			c.Wednesday = int(Wednesday)
			c.Thursday = int(Thursday)
			c.Friday = int(Friday)
			c.Saturday = int(Saturday)
			c.Sunday = int(Sunday)
			c.Uname = Uname
			if len(CourseName) > 0 && len(StartTime) > 0 && len(EndTime) > 0 && len(Uname) > 0 {
				if Id == 0 {
					id, err := m.AddCourse(c)
					if err != nil && id <= 0 {
						beego.Error(err)
						this.Rsp(false, "课程保存失败", "")
						return
					}
					this.Rsp(true, "课程保存成功", "")
				} else {
					c.Id = Id
					id, err := m.UpdateCourse(c)
					if err != nil && id <= 0 {
						beego.Error(err)
						this.Rsp(false, "课程保存失败", "")
						return
					}
					this.Rsp(true, "课程保存成功", "")

				}
			} else {
				this.CommonMenu()
				this.TplName = "haoadmin/data/course/list.html"
			}
			//}
		}
	}
	// beego.Debug(Name, StartTime, EndTime, Monday, Tuesday, Wednesday, Thursday, Friday, Uname1, Uname2, Uname3, Uname4, Uname5)
	// if len(Name) > 0 && len(StartTime) > 0 && len(EndTime) > 0 && Monday > 0 && Tuesday > 0 && Wednesday > 0 && Thursday > 0 && Friday > 0 && len(Uname1) > 0 && len(Uname2) > 0 && len(Uname4) > 0 && len(Uname4) > 0 && len(Uname5) > 0 {
	// 	c := new(m.Course)
	// 	c.Name = Name
	// 	c.StartTime = StartTime
	// 	c.EndTime = EndTime
	// 	c.Monday = Monday
	// 	c.Tuesday = Tuesday
	// 	c.Wednesday = Wednesday
	// 	c.Thursday = Thursday
	// 	c.Friday = Friday
	// 	c.Uname = Uname1 + "," + Uname2 + "," + Uname3 + "," + Uname4 + "," + Uname5
	// 	id, err := m.AddCourse(c)
	// 	if err != nil && id <= 0 {
	// 		beego.Error(err)
	// 		this.Rsp(false, "课程添加失败", "")
	// 		return
	// 	}
	// 	this.Rsp(true, "课程添加成功", "")
	// } else {
	// 	this.CommonMenu()
	// 	this.TplName = "haoadmin/data/course/list.html"
	// }

}

// 更新头衔

func (this *CourseController) UpdateCourse() {
	action := this.GetString("action")
	if action == "edit" {
		Name := this.GetString("Name")
		StartTime := this.GetString("StartTime")
		EndTime := this.GetString("EndTime")
		Id, _ := this.GetInt64("Id")
		Monday, _ := this.GetInt("Monday")
		Tuesday, _ := this.GetInt("Tuesday")
		Wednesday, _ := this.GetInt("Wednesday")
		Thursday, _ := this.GetInt("Thursday")
		Saturday, _ := this.GetInt("Saturday")
		Sunday, _ := this.GetInt("Sunday")
		Friday, _ := this.GetInt("Friday")
		Uname1 := this.GetString("Uname1")
		Uname2 := this.GetString("Uname2")
		Uname3 := this.GetString("Uname3")
		Uname4 := this.GetString("Uname4")
		Uname5 := this.GetString("Uname5")
		Uname6 := this.GetString("Uname6")
		Uname7 := this.GetString("Uname7")
		c := new(m.Course)
		c.Name = Name
		c.Id = Id
		c.StartTime = StartTime
		c.EndTime = EndTime
		c.Monday = Monday
		c.Tuesday = Tuesday
		c.Wednesday = Wednesday
		c.Thursday = Thursday
		c.Friday = Friday
		c.Saturday = Saturday
		c.Sunday = Sunday
		c.Uname = Uname1 + "," + Uname2 + "," + Uname3 + "," + Uname4 + "," + Uname5 + "," + Uname6 + "," + Uname7
		id, err := m.UpdateCourse(c)
		if err != nil && id <= 0 {
			beego.Error(err)
			this.AlertBack("课程修改失败")
			return
		}
		this.Alert("修改成功", "index")
	} else {
		this.CommonMenu()
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.AlertBack("课程获取失败")
			return
		}
		courseList, err := m.ReadCourseById(id)
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取课程信息错误")
			return
		}
		this.Data["courseList"] = courseList
		this.TplName = "haoadmin/data/course/edit.html"
	}

}

func (this *CourseController) DelCourse() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelCourseById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}

func (this *CourseController) GetCourseJson() {
	courselist := m.CourseList()
	teacherlist, _ := m.TeacherList()
	data := make(map[string]interface{})
	data["teacherlist"] = teacherlist
	data["courselist"] = courselist
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *CourseController) IndexCourseJson() {
	courselist := m.CourseList()
	this.Data["json"] = courselist
	this.ServeJSON()
}
