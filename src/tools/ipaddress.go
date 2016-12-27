package tools

import (
	// "bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/axgle/mahonia"
	"github.com/bitly/go-simplejson"
	"strings"
)

func GetIpProvinceCity(ip string) string {
	var result string
	if len(ip) <= 0 {
		return result
	}
	url := "http://ip.taobao.com/service/getIpInfo.php?ip="
	requestUrl := url + ip
	req := httplib.Get(requestUrl)
	str, err := req.String()
	if err != nil {
		beego.Debug(err)
		return result
	}
	key := []byte(str)
	js, err := simplejson.NewJson(key)
	if err != nil {
		beego.Error(err)
		return result
	}
	province := js.Get("data").Get("region").MustString()
	city := js.Get("data").Get("city").MustString()
	result = province + city
	return result
}

func GetPhoneCode(phone int64, code string) bool {
	var (
		uphone string
		ucode  string
		state  bool
	)

	uphone = fmt.Sprintf("%d", phone)
	ucode = fmt.Sprintf("验证码是:%s,15分钟输入有效,请完成注册。", code)
	requestUrl := "http://www.139000.com/send/gsend.asp?name=yccfgjs&pwd=Jzz123456&dst="
	requestUrl += uphone
	requestUrl += "&sender=&time=&txt=ccdx&msg="
	input := []byte(ucode)
	out := make([]byte, len(input)*2)
	// iconv.Convert(input, out, "utf-8", "gb2312")
	// out = bytes.TrimRight(out, "\x00")
	ucode = string(out)
	requestUrl += ucode
	req := httplib.Get(requestUrl)
	result, _ := req.String()
	enc := mahonia.NewDecoder("GB18030")
	result = enc.ConvertString(result)
	resultstate01 := strings.Index(result, "err=发送成功")
	resultstate02 := strings.Index(result, "errid=0")
	if -1 == resultstate01 || -1 == resultstate02 {
		state = false
	} else {
		state = true
	}
	return state
}
