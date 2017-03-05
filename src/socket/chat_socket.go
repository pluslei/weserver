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

var w *WriteData

type WriteData struct {
	jsondata chan *Socketjson
}

type SocketController struct {
	controllers.PublicController
}

type Userjobs struct {
	//socket.io
	socketiduser   map[string]string          //so->codeidname
	socketidso     map[string]socketio.Socket //codeidname->so
	socketiotoroom map[string]RoleRoom        //nametype->room
	userroom       map[string]Usertitle       //公司房间号对应的用户列表信息
}

var (
	WechatServer *socketio.Server
	job          Userjobs
	err          error
)

var (
	recordcount int = 10 //历史消息显示数量
)

func init() {
	w = &WriteData{
		jsondata: make(chan *Socketjson, 20480),
	}
	w.runWriteDb()
}

func Chatprogram() {
	if len(job.socketiduser) == 0 {
		//socket.io
		job.socketiduser = make(map[string]string)        //so->codeidname
		job.socketidso = make(map[string]socketio.Socket) //codeidname->so
		job.socketiotoroom = make(map[string]RoleRoom)    //nametype->leiwai->room
		job.userroom = make(map[string]Usertitle)         //公司房间号对应的用户列表信息
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
				codeid = Transformname(codeid, "", -1)                //解码公司代码和房间号
				userrole.Uname = js.Get("Uname").MustString()         //微信唯一标识
				userrole.Nickname = js.Get("Nickname").MustString()   //微信名
				userrole.UserIcon = js.Get("UserIcon").MustString()   //微信头像
				userrole.RoleName = js.Get("RoleName").MustString()   //头衔名称
				userrole.RoleTitle = js.Get("RoleTitle").MustString() //微信昵称
				userrole.InSider = js.Get("UserIcon").MustInt()       //人员类别内部人员或外部人员
				userrole.IsLogin = js.Get("IsLogin").MustBool()       //用户是否登录
				if userrole.IsLogin {
					roomval := fmt.Sprintf("%d", userrole.InSider) + "_" + userrole.RoleName + "_" + codeid
					so.Join(roomval)
					codename := Transformname(codeid, EncodeB64(userrole.Uname), 0) //公司代码用户名互转
					job.socketiduser[so.Id()] = codename
					job.socketidso[codename] = so
					job.userroom[codename] = userrole
					totalonline := 0
					for key, _ := range job.userroom {
						if len(key) > 0 {
							totalonline++
						}
					}
					so.Emit("all totalonline", fmt.Sprintf("%d", totalonline))
					if _, ok := job.socketiotoroom[codeid]; ok == true {
						roleroom := job.socketiotoroom[codeid].Roomval
						rolelen := len(roleroom[0])
						for i := 0; i < 2; i++ {
							for j := 0; j < rolelen; j++ {
								so.BroadcastTo(roleroom[i][j], "all totalonline", fmt.Sprintf("%d", totalonline))
							}
						}
					}
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
				sojson.Id = js.Get("Id").MustInt64()                                   //消息Id
				if sojson.Id > 0 {
					chat, _ := m.GetChatIdData(sojson.Id)
					if chat.Status == 1 {
						return
					}
					sojson.Uname = chat.Uname               //用户名
					sojson.Nickname = chat.Nickname         //用户昵称
					sojson.UserIcon = chat.UserIcon         //用户logo
					sojson.RoleName = chat.RoleName         //用户角色[vip,silver,gold,jewel]
					sojson.RoleTitle = chat.RoleTitle       //用户角色名[会员,白银会员,黄金会员,钻石会员]
					sojson.Sendtype = chat.Sendtype         //用户发送消息类型('TXT','IMG','VOICE')
					sojson.RoleTitleCss = chat.RoleTitleCss //头衔颜色
					if chat.RoleTitleBack == 1 {
						sojson.RoleTitleBack = true //角色聊天背景
					} else {
						sojson.RoleTitleBack = false //角色聊天背景
					}
					if chat.IsLogin == 1 {
						sojson.IsLogin = true //状态 [1、登录 0、未登录]
					} else {
						sojson.IsLogin = false //状态 [1、登录 0、未登录]
					}
					sojson.Insider = chat.Insider //1内部人员或0外部人员
					sojson.Content = chat.Content //消息内容
					sojson.Uuid = chat.Uuid       //uuid
					sojson.IsFilter = true        //消息是否过滤[true: 过滤, false: 不过滤]
					sojson.Status = 1
					sojson.Datatime = chat.Datatime //添加时间
				} else {
					sojson.Uuid = js.Get("Uuid").MustString()                 //uuid
					sojson.Uname = js.Get("Uname").MustString()               //用户名
					sojson.Nickname = js.Get("Nickname").MustString()         //用户昵称
					sojson.UserIcon = js.Get("UserIcon").MustString()         //用户logo
					sojson.RoleName = js.Get("RoleName").MustString()         //用户角色[vip,silver,gold,jewel]
					sojson.RoleTitle = js.Get("RoleTitle").MustString()       //用户角色名[会员,白银会员,黄金会员,钻石会员]
					sojson.Sendtype = js.Get("Sendtype").MustString()         //用户发送消息类型('TXT','IMG','VOICE')
					sojson.RoleTitleCss = js.Get("RoleTitleCss").MustString() //头衔颜色
					sojson.RoleTitleBack = js.Get("RoleTitleBack").MustBool() //角色聊天背景
					sojson.Insider = js.Get("Insider").MustInt()              //1内部人员或0外部人员
					sojson.IsLogin = js.Get("IsLogin").MustBool()             //状态 [1、登录 0、未登录]
					sojson.Content = js.Get("Content").MustString()           //消息内容
					sojson.IsFilter = js.Get("IsFilter").MustBool()           //消息是否过滤[true: 过滤, false: 不过滤]
					sojson.Status = js.Get("Status").MustInt()                //审核状态(0：未审核，1：审核)
					sojson.Datatime = time.Now()                              //添加时间
				}
				//消息入库
				SaveChatMsgdata(sojson)
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
					data := make(map[string]interface{})
					sojson.MsgType = "emit"
					data["msg"] = sojson
					so.Emit("all message", data)
					if !sojson.IsFilter || sojson.Status == 1 {
						if _, ok := job.socketiotoroom[codeid]; ok == true {
							roleroom := job.socketiotoroom[codeid].Roomval
							rolelen := len(roleroom[0])
							for i := 0; i < 2; i++ {
								for j := 0; j < rolelen; j++ {
									sojson.MsgType = "broadcastto"
									so.BroadcastTo(roleroom[i][j], "all message", data)
								}
							}
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

		so.On("disconnection", func() {
			if _, ok := job.socketiduser[so.Id()]; ok == true {
				codename := job.socketiduser[so.Id()]    //用户名
				codeid := Transformname("", codename, 1) //公司代码用户名互转
				totalonline := 0
				for key, _ := range job.userroom {
					if len(key) > 0 && key == codename {
						delete(job.userroom, codename)
					} else {
						totalonline++
					}
				}
				so.Emit("all totalonline", fmt.Sprintf("%d", totalonline))
				if _, ok := job.socketiotoroom[codeid]; ok == true {
					roleroom := job.socketiotoroom[codeid].Roomval
					rolelen := len(roleroom[0])
					for i := 0; i < 2; i++ {
						for j := 0; j < rolelen; j++ {
							so.BroadcastTo(roleroom[i][j], "all totalonline", fmt.Sprintf("%d", totalonline))
						}
					}
				}
			}
		})

		so.On("all deletemsg", func(msg string) {
			if len(msg) > 0 {
				msg = DecodeB64(msg)
				key := []byte(msg)
				js, err := simplejson.NewJson(key)
				if err != nil {
					return
				}
				uuid := js.Get("Uuid").MustString() //uuid
				codeid := js.Get("Codeid").MustString()
				codeid = Transformname(codeid, "", -1) //解码公司代码和房间号
				if len(uuid) > 0 {
					id, err := m.DelChatById(uuid)
					if id > 0 && err == nil {
						if _, ok := job.socketiotoroom[codeid]; ok == true {
							roleroom := job.socketiotoroom[codeid].Roomval
							rolelen := len(roleroom[0])
							for i := 0; i < 2; i++ {
								for j := 0; j < rolelen; j++ {
									so.BroadcastTo(roleroom[i][j], "all deletemsg", uuid)
								}
							}
						}
					}
				}

			}
		})

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
		sysconfig, _ := m.GetAllSysConfig()   //系统设置
		recordcount := sysconfig.HistoryCount //显示历史记录条数
		var historychat []m.ChatRecord
		switch sysconfig.HistoryMsg { //是否显示历史消息 0显示  1 不显示
		case 0:
			historychat, _, _ = m.GetChatMsgData(recordcount)
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

func (this *SocketController) ChatOnlineUserMsg() {
	if this.IsAjax() {
		var usermsg []OnlineUserMsg
		for key, item := range job.userroom {
			if len(key) > 0 {
				var msg OnlineUserMsg
				msg.Nickname = item.Nickname
				msg.UserIcon = item.UserIcon
				usermsg = append(usermsg, msg)
			}
		}
		data := make(map[string]interface{})
		data["onlineuser"] = usermsg //聊天的历史信息
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

//广播入库
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

//时时消息入库
func SaveChatMsgdata(sojson Socketjson) {
	jsondata := &sojson
	select {
	case w.jsondata <- jsondata:
		break
	default:
		beego.Error("write db error!!!")
		break
	}
}

func (w *WriteData) runWriteDb() {
	go func() {
		for {
			sojson, ok := <-w.jsondata
			if ok {
				if sojson.Status == 0 {
					addData(sojson)
				} else {
					UpdateData(sojson)
				}
			}
		}
	}()
}

func addData(sojson *Socketjson) {
	beego.Debug("im here", sojson, sojson.RoleTitleBack)
	if sojson.IsLogin && sojson.Insider == 1 {
		//写数据库
		var chatrecord m.ChatRecord
		chatrecord.Uuid = sojson.Uuid                 //uuid
		chatrecord.Code = sojson.Code                 //公司代码
		chatrecord.Room = sojson.Room                 //房间号
		chatrecord.Uname = sojson.Uname               //用户名
		chatrecord.Nickname = sojson.Nickname         //用户昵称
		chatrecord.UserIcon = sojson.UserIcon         //用户logo
		chatrecord.RoleName = sojson.RoleName         //用户角色[vip,silver,gold,jewel]
		chatrecord.RoleTitle = sojson.RoleTitle       //用户角色名[会员,白银会员,黄金会员,钻石会员]
		chatrecord.Sendtype = sojson.Sendtype         //用户发送消息类型('TXT','IMG','VOICE')
		chatrecord.RoleTitleCss = sojson.RoleTitleCss //头衔颜色
		if sojson.RoleTitleBack {
			chatrecord.RoleTitleBack = 1 //角色聊天背景
		} else {
			chatrecord.RoleTitleBack = 0 //角色聊天背景
		}
		chatrecord.Insider = sojson.Insider   //1内部人员或0外部人员
		chatrecord.IsLogin = 1                //状态 [1、登录 0、未登录]
		chatrecord.Content = sojson.Content   //消息内容
		chatrecord.Datatime = sojson.Datatime //添加时间
		if !sojson.IsFilter {
			chatrecord.Status = 1 //审核状态(0：未审核，1：审核)
		} else {
			chatrecord.Status = sojson.Status //审核状态(0：未审核，1：审核)
		}
		_, err = m.AddChat(&chatrecord)
		if err != nil {
			beego.Debug(err)
		}
	}
}

func UpdateData(sojson *Socketjson) {
	beego.Debug("im here", sojson, sojson.RoleTitleBack)
	if sojson.IsLogin && sojson.Insider == 1 {
		//更新数据库
		_, err = m.UpdateChatStatus(sojson.Id)
		if err != nil {
			beego.Debug(err)
		}
	}
}
