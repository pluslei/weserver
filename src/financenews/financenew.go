package financenews

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"strings"
	m "weserver/models"
	"weserver/src/tools"
	// "time"
)

func GetFinanceNews() {
	defer func() {
		if err := recover(); err != nil {
			beego.Error("GetFinanceNews: ", err)
		}
	}()

	var url = "http://www.jin10.com/jin10.com.html"
	p, err := goquery.NewDocument(url)
	if err != nil {
		beego.Error("FinanceNewController:get the url error", err)
	}

	listline := p.Find("div #listarea div.listline")
	financeOne, err := m.GetFinanceNewsOne()

	// 初始化
	if err != nil {
		beego.Error("FinanceNewController:get the database error", err)
		for i := listline.Length() - 1; i >= 0; i-- {
			bo := InsetFinanceNews(financeOne.Cmd5, listline, i)
			if !bo {
				break
			}
		}
	} else {
		for i := 0; i < listline.Length(); i++ {
			bo := InsetFinanceNews(financeOne.Cmd5, listline, i)
			if !bo {
				break
			}
		}
	}
}

func InsetFinanceNews(Cmd5 string, listline *goquery.Selection, i int) bool {
	var (
		style int
		html  string
	)

	times := strings.TrimSpace(listline.Eq(i).Find("div.left div.time").Text())
	content, _ := listline.Eq(i).Find("div.right p.text").Html()
	cmd5 := tools.EncodeUserPwd(content, times)
	attrValue, _ := listline.Eq(i).Attr("class")
	if attrValue == "listline important" {
		style = 1
	}

	if Cmd5 != cmd5 {
		_, hashref := listline.Eq(i).Find("div.right p.text a").Attr("href")
		if !hashref {
			html2str := beego.HTML2str(content)
			if html2str == "" {
				htmltemp, _ := listline.Eq(i).Find("div.right div.content").Html()
				html = beego.HTML2str(htmltemp)
			} else {
				html = html2str
			}

			finance := new(m.FinanceNews)
			finance.Pulltime = times
			finance.Contents = strings.Trim(html, " \n")
			finance.Cmd5 = cmd5
			finance.Style = style
			id, err := m.AddFinanceNews(finance)
			if err != nil && id <= 0 {
				beego.Error(err)
			}
		}
		return true
	} else {
		_, err := m.DelFinanceNewsData()
		if err != nil {
			beego.Error("FinanceNewController:Delect data error", err)
		}
		return false
	}
	return true
}
