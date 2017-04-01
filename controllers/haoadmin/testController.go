package haoadmin

import (
	// tool "weserver/src/tool"
	// "fmt"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"weserver/controllers/mqtt"
	"weserver/src/tools"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Index() {
	// this.Ctx.WriteString("TEST")
	this.TplName = "haoadmin/test/index.html"
}

func (this *TestController) PostApi() {
	data := this.GetString("data")
	jsonStr := tools.MainEncrypt(strings.Trim(data, " "))
	testurl := "http://localhost:" + beego.AppConfig.String("httpport") + "/api"
	u, _ := url.Parse(testurl)
	q := u.Query()
	q.Set("data", jsonStr)
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		beego.Debug("get error")
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", result)
	this.Ctx.WriteString(string(result))
}

func (this *TestController) Test() {
	msgtype := mqtt.NewMessageType(mqtt.MSG_TYPE_BROCAST)
	topic := this.GetString("topic") // 从管理页面获取topic
	msgtype.SendBrocast(topic, "12d1d2s1d12")
}
