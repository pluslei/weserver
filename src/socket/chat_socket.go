package socket

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	"github.com/googollee/go-socket.io"
	m "weserver/models"
	//"haolive/controllers/haoindex"
	"weserver/controllers"
	. "weserver/src/tools"
	//"github.com/astaxie/beego/context"
	//"path"
	//"sort"
	"strconv"
	"strings"
	"time"
)

type SocketController struct {
	controllers.PublicController
}

type Userjobs struct {
	//socket.io
	socketiduser     map[string]string          //so->codeidname
	socketidso       map[string]socketio.Socket //codeidname->so
	socketiduserip   map[string]OnlineIpPro     //codeidname-ip
	socketidusertime map[string]int             //codeidname-time
	socketiotoroom   map[string]RoleRoom        //nametype->room
}

var (
	WechatServer *socketio.Server
	job          Userjobs
	err          error
)

var (
	recordcount int = 10 //历史消息显示数量
)

func Chatprogram() {
	if len(job.socketiduser) == 0 {
		//socket.io
		job.socketiduser = make(map[string]string)        //so->codeidname
		job.socketidso = make(map[string]socketio.Socket) //codeidname->so
		job.socketiduserip = make(map[string]OnlineIpPro) //codeidname-ip
		job.socketidusertime = make(map[string]int)       //codeidname-time
		job.socketiotoroom = make(map[string]RoleRoom)    //nametype->leiwai->room
	}

	WechatServer, err = socketio.NewServer(nil)
	if err != nil {
		beego.Error(err)
	}
	WechatServer.On("connection", func(so socketio.Socket) {
		so.On("all connection", func(msg string) {
			if len(msg) > 0 {
				msg = DecodeB64(msg)
				key := []byte(msg)
				js, err := simplejson.NewJson(key)
				if err != nil {
					beego.Error(err)
				}
				var userrole Usertitle
				codeid := js.Get("Codeid").MustString()
				codeid = Transformname(codeid, "", -1)              //解码公司代码和房间号
				userrole.Uname = js.Get("Uname").MustString()       //用户名
				userrole.RoleName = js.Get("RoleName").MustString() //用户角色
				userrole.InSider = js.Get("InSider").MustInt()      //1内部人员或0外部人员
				userrole.IsLogin = js.Get("IsLogin").MustBool()     //用户是否登录
				if userrole.IsLogin {
					roomval := fmt.Sprintf("%d", userrole.InSider) + "_" + userrole.RoleName + "_" + codeid
					so.Join(roomval)
					codename := Transformname(codeid, EncodeB64(userrole.Uname), 0) //公司代码用户名互转
					job.socketiduser[so.Id()] = codename
					job.socketidso[codename] = so
				}
			}
		})

		so.On("all message", func(msg string) {
			if len(msg) > 0 {
				var sojson Socketjson
				msg = DecodeB64(msg)
				key := []byte(msg)
				js, err := simplejson.NewJson(key)
				if err != nil {
					beego.Error(err)
				}
				codeid := js.Get("Codeid").MustString()
				codeid = Transformname(codeid, "", -1)                                 //解码公司代码和房间号
				code, _ := strconv.ParseInt(beego.AppConfig.String("company"), 10, 64) //公司代码
				sojson.Code = int(code)                                                //公司代码
				room, _ := strconv.ParseInt(beego.AppConfig.String("room"), 10, 64)    //房间号
				sojson.Room = int(room)                                                //房间号
				sojson.Uname = js.Get("Uname").MustString()                            //用户名
				sojson.Nickname = js.Get("Nickname").MustString()                      //用户昵称
				sojson.UserIcon = js.Get("UserIcon").MustString()                      //用户logo
				sojson.RoleName = js.Get("RoleName").MustString()                      //用户角色[vip,silver,gold,jewel]
				sojson.RoleTitle = js.Get("RoleTitle").MustString()                    //用户角色名[会员,白银会员,黄金会员,钻石会员]
				sojson.Sendtype = js.Get("Sendtype").MustString()                      //用户发送消息类型('TXT','IMG','VOICE')
				sojson.RoleTitleCss = js.Get("RoleTitleCss").MustString()              //头衔颜色
				sojson.RoleTitleBack = js.Get("RoleTitleBack").MustBool()              //角色聊天背景
				sojson.Insider = js.Get("Insider").MustInt()                           //1内部人员或0外部人员
				sojson.IsLogin = js.Get("IsLogin").MustBool()                          //状态 [1、登录 0、未登录]
				sojson.Content = js.Get("Sendtype").MustString()                       //消息内容
				sojson.Datatime = time.Now()                                           //添加时间
				beego.Debug(sojson, "==================================")
				//用户信息进行加密
				sojson.Uname = EncodeB64(sojson.Uname)
				sojson.Nickname = EncodeB64(sojson.Nickname)
				sojson.UserIcon = EncodeB64(sojson.UserIcon)
				sojson.RoleName = EncodeB64(sojson.RoleName)
				sojson.RoleTitle = EncodeB64(sojson.RoleTitle)
				sojson.Sendtype = EncodeB64(sojson.Sendtype)
				sojson.RoleTitleCss = EncodeB64(sojson.RoleTitleCss)
				sojson.Content = EncodeB64(sojson.Content)
				if sojson.IsLogin && sojson.Insider == 1 {
					codename := Transformname(codeid, sojson.Uname, 0) //公司代码用户名互转
					if _, ok := job.socketidso[codename]; ok == true {
						data := make(map[string]interface{})
						data["msg"] = sojson
						job.socketidso[codename].Emit("all message", data)
						if _, ok := job.socketiotoroom[codeid]; ok == true {
							roleroom := job.socketiotoroom[codeid].Roomval
							rolelen := len(roleroom[0])
							for i := 0; i < 2; i++ {
								for j := 0; j < rolelen; j++ {
									job.socketidso[codename].BroadcastTo(roleroom[i][j], "all message", data)
								}
							}
							//保存数据
							// Savechatdata(sojson, 0)
						}
					}
				}
			}
		})

		so.On("all roombroadmsg", func(msg string) {
			if len(msg) > 0 {
				var sojson Socketjson
				msg = DecodeB64(msg)
				key := []byte(msg)
				js, err := simplejson.NewJson(key)
				if err != nil {
					beego.Error(err)
					sendshowmsg := "消息发送失败,请联系管理员!"
					so.Emit("all showmsg", sendshowmsg)
					return
				}
				sojson.Uname = js.Get("Uname").MustString()
				if len(sojson.Uname) <= 0 {
					sendshowmsg := "消息发送失败,请联系管理员!"
					so.Emit("all showmsg", sendshowmsg)
					return
				}
				sojson.Content = js.Get("Content").MustString()
				codeid := js.Get("Codeid").MustString()
				codeid = Transformname(codeid, "", -1)                                 //解码公司代码和房间号
				code, _ := strconv.ParseInt(beego.AppConfig.String("company"), 10, 64) //公司代码
				sojson.Code = int(code)                                                //公司代码
				room, _ := strconv.ParseInt(beego.AppConfig.String("room"), 10, 64)    //房间号
				sojson.Room = int(room)                                                //房间号
				sojson.Uname = EncodeB64(sojson.Uname)
				sojson.Content = EncodeB64(sojson.Content)
				//房间号获取
				roleval, _ := m.GetAllUserRole()
				rolelen := len(roleval)
				prevalue := codeid
				data := make(map[string]interface{})
				data["msg"] = sojson
				for i := 0; i < rolelen; i++ {
					for j := 0; j < 2; j++ {
						roleval[i].Name = strings.Replace(roleval[i].Name, " ", "", -1) //去空格
						roomval := fmt.Sprintf("%d", j) + "_" + roleval[i].Name + "_" + prevalue
						so.BroadcastTo(roomval, "all broadmsg", data)
					}
				}
				SaveBroadCastdata(sojson)
				sendshowmsg := "广播发送成功!"
				so.Emit("all success", sendshowmsg)
			}
		})

		so.On("disconnection", func() {})

		so.On("all kickout", func(msg string) {
			if len(msg) > 0 {
				msg = DecodeB64(msg)
				key := []byte(msg)
				js, err := simplejson.NewJson(key)
				if err != nil {
					beego.Error(err)
				} else {
					uname := js.Get("Uname").MustString()
					objname := EncodeB64(js.Get("Objname").MustString())
					beego.Debug("uname", uname, js.Get("Objname").MustString())
					codeid := js.Get("Codeid").MustString()       //公司房间标识符
					codeid = Transformname(codeid, "", -1)        //解码公司代码和房间号
					codeuser := Transformname(codeid, objname, 0) //公司代码用户名互转
					if _, ok := job.socketidso[codeuser]; ok == true {
						// user := new(m.User)
						// user.Username = uname
						// _, err := m.LoadRelatedUser(user, "Username")
						// if err == nil {
						// beego.Debug("====", user)
						job.socketidso[codeuser].Emit("all kickout", "")
						if _, ok := job.socketiduser[job.socketidso[codeuser].Id()]; ok == true {
							delete(job.socketiduser, job.socketidso[codeuser].Id())
						}
						delete(job.socketidso, codeuser)
					} else {
						sendshowmsg := "操作失败!"
						so.Emit("all showmsg", sendshowmsg)
					}
				}
			}
		})
	})
	WechatServer.On("error", func(so socketio.Socket, err error) {
		beego.Error(err)
	})
	beego.BeeApp.Handlers.Handler("/wechatSocket/", WechatServer)
}

//获取客户的真是IP地址
func (this *SocketController) GetClientip() string {
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

func (this *SocketController) ChatUserList() {
	if this.IsAjax() {
		codeid := this.GetString("codeid")              //公司房间标识符
		codeid = Transformname(codeid, "", -1)          //解码公司代码和房间号
		coderoom := Transformname(codeid, "", 2)        //房间号
		roomid, _ := strconv.ParseInt(coderoom, 10, 64) //房间号
		go func() {
			//房间号获取
			roleval, _ := m.GetAllUserRole()
			rolelen := len(roleval)
			var roleroom RoleRoom
			prevalue := codeid
			for i := 0; i < rolelen; i++ {
				for j := 0; j < 2; j++ {
					roleval[i].Name = strings.Replace(roleval[i].Name, " ", "", -1) //去空格
					roomval := fmt.Sprintf("%d", j) + "_" + roleval[i].Name + "_" + prevalue
					roleroom.Roomval[j] = append(roleroom.Roomval[j], roomval)
				}
			}
			job.socketiotoroom[prevalue] = roleroom
		}()
		sysconfig, _ := m.GetAllSysConfig()       //系统设置
		recordcount = int(sysconfig.HistoryCount) //显示历史记录条数
		var (
			historychat []Socketjson
		)
		switch sysconfig.HistoryMsg { //是否显示历史消息 0显示  1 不显示
		case 0:
		default:
		}
		data := make(map[string]interface{})
		data["historydata"] = historychat //聊天的历史信息
		//从数据库中获取公告中的最后一条内容
		broaddata, _ := m.GetBroadcastData(int(roomid))
		data["notice"] = broaddata //公告
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

func (this *SocketController) ChatUpload() {
	this.Ctx.WriteString("")
}

func (this *SocketController) ChatKickOut() {
	this.Ctx.WriteString("")
}

func (this *SocketController) ChatModifyIcon() {
	if this.GetSession("indexUserInfo") != nil && this.IsAjax() {
		data := make(map[string]interface{})
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

func SaveBroadCastdata(sojson Socketjson) {
	//写数据库
	var broad m.Broadcast
	broad.Code = sojson.Code
	broad.Room = sojson.Room
	broad.Uname = DecodeB64(sojson.Uname)
	broad.Data = DecodeB64(sojson.Content)
	broad.Datatime = time.Now()
	_, err = m.AddBroadcast(&broad)
	if err != nil {
		beego.Debug(err)
	}
}
