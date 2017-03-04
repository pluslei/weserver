package haoadmin

import (
	m "weserver/models"
	// tool "weserver/src/tool"
	// "fmt"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	rpc "weserver/src/rpcserver"
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

func (this *TestController) Test() {
	//写数据库
	var chatrecord m.ChatRecord
	chatrecord.Id = 71
	chatrecord.Nickname = "handler"                                                                                                                          //用户昵称
	chatrecord.UserIcon = "http://wx.qlogo.cn/mmopen/KetjXWSVppsZ0icialcRKRX0czbk6icxcerSR0coP0qLdyVYU4uwdQ2NPDh4b5DcF2mYbOdLl5pcXwNsC4fiacUvTzmzpu2FCJaY/0" //用户logo
	chatrecord.RoleName = "nl_ordinary"                                                                                                                      //用户角色[vip,silver,gold,jewel]
	chatrecord.RoleTitle = "RoleTitle"                                                                                                                       //用户角色名[会员,白银会员,黄金会员,钻石会员]
	chatrecord.Sendtype = "IMG"                                                                                                                              //用户发送消息类型('TXT','IMG','VOICE')
	chatrecord.RoleTitleCss = "#992BAC"                                                                                                                      //头衔颜色
	chatrecord.RoleTitleBack = 1                                                                                                                             //角色聊天背景
	chatrecord.Insider = 1                                                                                                                                   //1内部人员或0外部人员
	chatrecord.IsLogin = 1                                                                                                                                   //状态 [1、登录 0、未登录]
	chatrecord.Status = 0
	chatrecord.Uname = "ooex5xBIIiGr4YVk9-02dTTBZVlI"
	chatrecord.Content = "x9CxJpvVPyD41BDGATQmBE0sN1O_CtFFBtTBx3icGMki5IqnMAd_fy0p" //消息内容
	chatrecord.Datatime = time.Now()                                                //添加时间
	chatrecord.DatatimeStr = time.Now().Format("2006-01-02 15:04:05")
	// 插入成功广播
	rpc.Broadcast("chat", chatrecord, func(result []string) { beego.Debug("result", result) })
}
