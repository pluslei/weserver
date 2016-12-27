package haoadmin

import (
	// m "hnnaserver/models"
	// tool "hnnaserver/src/tool"
	// "fmt"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"weserver/src/tools"
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
