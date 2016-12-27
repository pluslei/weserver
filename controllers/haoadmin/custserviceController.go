package haoadmin

import (
	"github.com/astaxie/beego"
	m "weserver/models"
)

type CustserviceController struct {
	CommonController
}

// 组别显示页面
func (this *CustserviceController) Index() {
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
		qslist, count := m.GetCustservicelist(iStart, iLength, "Id")

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
		this.TplName = "haoadmin/data/custservice/list.html"
	}

}

// 增加组别
func (this *CustserviceController) AddCust() {
	CustName := this.GetString("CustName")
	CustNumber := this.GetString("CustNumber")
	Status, _ := this.GetInt("Status")
	Order, _ := this.GetInt("Order")
	count, _ := m.CustStatusCount()
	if len(CustName) > 0 && len(CustNumber) > 0 && Status > 0 {
		c := new(m.Custservice)
		c.CustName = CustName
		c.CustNumber = CustNumber
		c.Status = Status
		c.Order = Order
		// if Status == 1 {
		if count > 5 {
			this.AlertBack("新增客服失败,客服数量超过6个。")
			return
		} else {
			id, err := m.AddCustservice(c)
			if err != nil && id <= 0 {
				beego.Error(err)
				this.AlertBack("添加失败")
				return
			}
			this.Alert("添加成功", "custservice_index")
		}
		// } else {
		// 	id, err := m.AddCustservice(c)
		// 	if err != nil && id <= 0 {
		// 		beego.Error(err)
		// 		this.AlertBack("添加失败")
		// 		return
		// 	}
		// 	this.Alert("添加成功", "custservice_index")
		// }

		// id, err := m.AddCustservice(c)
		// if err != nil && id <= 0 {
		// 	beego.Error(err)
		// 	this.AlertBack("添加失败")
		// 	return
		// }
		// this.Alert("添加成功", "index")
	} else {
		this.CommonMenu()
		this.TplName = "haoadmin/data/custservice/add.html"
	}

}

func (this *CustserviceController) UpdateCust() {
	CustName := this.GetString("CustName")
	CustNumber := this.GetString("CustNumber")
	Status, _ := this.GetInt("Status")
	Order, _ := this.GetInt("Order")
	Id, _ := this.GetInt64("Id")
	c := new(m.Custservice)
	c.Id = Id
	c.CustName = CustName
	c.CustNumber = CustNumber
	c.Status = Status
	c.Order = Order
	if len(CustName) > 0 && len(CustNumber) > 0 && Status > 0 && Id > 0 {
		err := c.UpdateCustservice("CustName", "CustNumber", "Status", "Order")
		if err != nil {
			beego.Error(err)
			this.AlertBack("修改失败")
			return
		}
		this.Alert("修改成功", "custservice_index")
	} else {
		this.CommonMenu()
		id, err := this.GetInt64("Id")
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取信息错误")
			return
		}
		qsList, err := m.ReadCustserviceById(id)
		if err != nil {
			beego.Error(err)
			this.AlertBack("获取信息错误")
			return
		}
		this.Data["qsList"] = qsList
		this.TplName = "haoadmin/data/custservice/edit.html"
	}

}

func (this *CustserviceController) DelCust() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelCustserviceById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "删除成功", "")
		return
	} else {
		this.Rsp(false, err.Error(), "")
		return
	}
}
