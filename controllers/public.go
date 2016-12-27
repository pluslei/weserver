package controllers

import (
	"github.com/astaxie/beego"

	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

type PublicController struct {
	beego.Controller
}

func (this *PublicController) Rsp(status bool, str string, url string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str, "url": url}
	this.ServeJSON()
}

func (this *PublicController) Alert(info string, url string) {
	alert := fmt.Sprintf("<script>alert('" + info + "');location.href='" + url + "';</script>")
	this.Ctx.WriteString(alert)
}

func (this *PublicController) AlertBack(info string) {
	alert := fmt.Sprintf("<script>alert('" + info + "');history.go(-1);</script>")
	this.Ctx.WriteString(alert)
}

func (this *PublicController) HttpnRmageRequest(uri string, params map[string]string) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	content_type := writer.FormDataContentType()

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", uri, body)

	req.Header.Set("Content-Type", content_type)
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}
