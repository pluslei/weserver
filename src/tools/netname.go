package tools

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"

	"math/rand"
	"strconv"
	"time"
)

func GetNetName() string {
	dec := mahonia.NewDecoder("gbk")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randpage := r.Intn(924) + 1
	url := "http://www.yimanwu.com/daquan/list_7_" + strconv.Itoa(randpage) + ".html"
	p, err := goquery.NewDocument(url)
	if err != nil {
		beego.Error("FinanceNewController:get the url error", err)
		return ""
	} else {
		listline := p.Find("div .list ul li")
		randname := r.Intn(listline.Length())
		return dec.ConvertString(listline.Eq(randname).Find("p").Text())
	}
}
