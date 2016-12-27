package socket

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	"github.com/googollee/go-socket.io"
	m "weserver/models"
	p "weserver/src/parameter"
	//"haolive/controllers/haoindex"
	"menteslibres.net/gosexy/redis"
	"weserver/controllers"
	. "weserver/src/tools"
	//"github.com/astaxie/beego/context"
	"math/rand"
	"path"
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
	//redis
	p.Client = redis.New()
	p.Rediserr = p.Client.Connect("127.0.0.1", uint(6379))
	if p.Rediserr != nil {
		beego.Info("Connect failed:", p.Rediserr)
	}

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
				islogin := js.Get("IsLogin").MustBool()
				codeid := js.Get("Codeid").MustString()
				userrole.Uname = js.Get("Uname").MustString()            //用户名
				userrole.UserIcon = js.Get("UserIcon").MustString()      //用户Icon
				userrole.RoleName = js.Get("RoleName").MustString()      //用户角色
				userrole.Titlerole = js.Get("Titlerole").MustString()    //用户类型
				userrole.Authorcss = js.Get("Authorcss").MustString()    //用户头衔
				userrole.InSider = js.Get("Insider").MustInt()           //1内部人员或0外部人员
				userrole.Logintime = time.Now().Format("01-02 15:04:05") //登入的时间
				userrole.Datatime = time.Now()                           //登入的时间
				codeid = Transformname(codeid, "", -1)                   //解码公司代码和房间号
				userrole.Roomid = Transformname(codeid, "", 2)           //房间号
				dataem := make(map[string]interface{})
				var (
					userroom map[string]Usertitle //公司房间号对应的用户列表信息
					urole    []Usertitle
				)
				if p.Rediserr == nil {
					roomval := fmt.Sprintf("%d", userrole.InSider) + "_" + userrole.RoleName + "_" + codeid
					so.Join(roomval)
					//用户列表信息
					userroom = make(map[string]Usertitle) //房间对应的用户信息
					jobroom := "coderoom_" + codeid
					roomdata, _ := p.Client.Get(jobroom)
					if len(roomdata) > 0 {
						userroom, _ = Jsontoroommap(roomdata)
					}
					codename := Transformname(codeid, EncodeB64(userrole.Uname), 0) //公司代码用户名互转
					//如果有用户重复登入就私聊退出该用户
					if _, ok := job.socketidso[codename]; ok == true {
						if islogin == true {
							data := make(map[string]interface{})
							data["user"] = userroom[codename]
							data["leave"] = "leave"
							job.socketidso[codename].Emit("all disconnection", data)
						}

						if _, ok := job.socketiduser[job.socketidso[codename].Id()]; ok == true {
							delete(job.socketiduser, job.socketidso[codename].Id())
						}
						if _, ok := userroom[codename]; ok == true {
							delete(userroom, codename)
						}
						delete(job.socketidso, codename)
					}
					//新用户信息
					job.socketiduser[so.Id()] = codename
					job.socketidso[codename] = so
					userroom[codename] = userrole
					for rolval, userId := range userroom {
						if _, ok := job.socketidso[rolval]; ok == true {
							if len(userId.Uname) > 0 {
								urole = append(urole, userId)
							} else {
								delete(userroom, rolval)
							}
						} else {
							delete(userroom, rolval)
						}
					}
					jsonRes, _ := json.Marshal(userroom)
					jsonstr := string(jsonRes)
					p.Client.Set(jobroom, jsonstr)
					dataem["utype"] = "emit"
					var casttomsg string
					body, err01 := json.Marshal(urole)
					if err01 == nil {
						casttomsg = string(body)
					}
					casttomsg = EncodeB64(casttomsg)
					//dataem["msg"] = urole
					dataem["msg"] = casttomsg
					so.Emit("all connection", dataem)
					dataem["utype"] = "broad"
					if _, ok := job.socketiotoroom[codeid]; ok == true {
						roleroom := job.socketiotoroom[codeid].Roomval
						rolelen := len(roleroom[0])
						for i := 0; i < 2; i++ {
							for j := 0; j < rolelen; j++ {
								so.BroadcastTo(roleroom[i][j], "all connection", dataem)
							}
						}
					}
				}
			}
		})

		so.On("all message", func(msg string) {
			if len(msg) > 0 {
				//var ObjBroad  bool    //是否广播
				var sojson Socketjson
				msg = DecodeB64(msg)
				key := []byte(msg)
				js, err := simplejson.NewJson(key)
				if err != nil {
					beego.Error(err)
				}
				sojson.Author = js.Get("Author").MustString()
				codeid := js.Get("Codeid").MustString()
				codeid = Transformname(codeid, "", -1)          //解码公司代码和房间号
				coderoom := Transformname(codeid, "", 2)        //房间号
				roomid, _ := strconv.ParseInt(coderoom, 10, 64) //房间号
				sojson.Coderoom = int(roomid)                   //公司代码房间号
				sojson.Codeid = codeid                          //解码公司代码和房间号
				//禁言控制
				chatstate := true
				var gag m.GagControl
				gag, err = m.SelectGagControl(sojson.Author)
				if err == nil && gag.Gatstate == 0 {
					unixtime := -1
					switch gag.Gagmode {
					case 5:
						unixtime = int(time.Now().Unix() - 300)
					default:
					}
					chatstate = false
					if int64(unixtime) >= gag.Gagtime {
						chatstate = true
						gag.Gatstate = 1
						_, err = m.UpdateGagControl(gag)
						if err != nil {
							beego.Debug(err)
						}
					}
				}
				if chatstate {
					sojson.AuthorInSider = 0 //1内部人员或0外部人员
					sojson.AuthorRole = js.Get("AuthorRole").MustString()
					sojson.Authortype = js.Get("Authortype").MustString()
					sojson.AuditStatus = js.Get("AuditStatus").MustInt()
					sojson.Sendtype = js.Get("Sendtype").MustString() //用户发送消息类型，txt, img
					sojson.Sendtype = EncodeB64(sojson.Sendtype)
					sojson.AuthorRole = EncodeB64(sojson.AuthorRole)
					sojson.Authortype = EncodeB64(sojson.Authortype)
					sojson.Authorcss = EncodeB64(js.Get("UserIcon").MustString())
					sojson.Author = EncodeB64(sojson.Author)
					sojson.Content = js.Get("Content").MustString()
					sojson.Content = EncodeB64(sojson.Content)
					sojson.Chat = EncodeB64(js.Get("Chat").MustString())
					sojson.Time = time.Now().Format("15:04")
					sojson.Newtime = time.Now().Format("2006-01-02 15:04:05") //发言时间
					sojson.Datatime = time.Now()
					//数据传给后台
					uname := sojson.Author
					codename := Transformname(codeid, uname, 0) //公司代码用户名互转
					//从map中获取数据
					if _, ok := job.socketiduserip[codename]; ok == true {
						sojson.Ipaddress = job.socketiduserip[codename].Ip  //获取ip地址
						sojson.Procities = job.socketiduserip[codename].Pro //省市
					}
					data := make(map[string]interface{})
					data["msg"] = sojson
					sendmessage := true
					if sojson.AuthorDelay > 0 {
						if _, ok := job.socketidusertime[codename]; ok == true {
							authorunixtime := int(time.Now().Unix()) - job.socketidusertime[codename]
							gagtime := sojson.AuthorDelay
							if authorunixtime < gagtime {
								sendmessage = false
								if _, ok := job.socketidso[codename]; ok == true {
									//发言间隔时间提示
									sendshowmsg := fmt.Sprintf("请间隔%d秒在发言!", gagtime)
									job.socketidso[codename].Emit("all showmsg", sendshowmsg)
								}
							} else {
								job.socketidusertime[codename] = int(time.Now().Unix())
							}
						} else {
							job.socketidusertime[codename] = int(time.Now().Unix())
						}
					}
					if sojson.AuthorInSider == 0 && sojson.AuditStatus != 100 {
						sysconfig, _ := m.GetAllSysConfig()              //系统设置
						sojson.AuthorDelay = int(sysconfig.ChatInterval) //用户禁言时间
						sojson.AuditStatus = sysconfig.AuditStatus       //1，不需要审核，2，需要审核
					}
					//var sendmsgtext string
					if sendmessage {
						switch DecodeB64(sojson.Chat) {
						case "allchat":
							//保存数据
							if sojson.AuditStatus != 100 {
								sojson.IsEmitBroad = true
								/*
									body, err01 := json.Marshal(sojson)
									if err01 == nil {
										sendmsgtext = string(body)
									}
									sendmsgtext = EncodeB64(sendmsgtext)
									data["msg"] = sendmsgtext
								*/
								data["msg"] = sojson
								so.Emit("all message", data)
							}
							if _, ok := job.socketiotoroom[codeid]; ok == true {
								roleroom := job.socketiotoroom[codeid].Roomval
								rolelen := len(roleroom[0])
								k := 0
								/*
									if sojson.AuthorInSider == 0 && sojson.AuditStatus == 2 {
										k = 1
									}
								*/
								sojson.IsEmitBroad = false
								/*
									body, err01 := json.Marshal(sojson)
									if err01 == nil {
										sendmsgtext = string(body)
									}
									sendmsgtext = EncodeB64(sendmsgtext)
								*/
								data["msg"] = sojson
								for i := k; i < 2; i++ {
									for j := 0; j < rolelen; j++ {
										so.BroadcastTo(roleroom[i][j], "all message", data)
									}
								}
								Savechatdata(sojson, 0)
							}
						default:
						}
					} else {
						sendshowmsg := "你没有此权限,请联系管理员!"
						so.Emit("all showmsg", sendshowmsg)
					}
				} else {
					sendshowmsg := "你没有此权限,请联系管理员!"
					so.Emit("all showmsg", sendshowmsg)
				}
			}
		})

		//模拟消息
		so.On("all analogmessage", func(msg string) {
			if len(msg) > 0 {
				var sojson Socketjson
				msg = DecodeB64(msg)
				key := []byte(msg)
				js, err := simplejson.NewJson(key)
				if err != nil {
					beego.Error(err)
				}
				objauthor := js.Get("Objauthor").MustString()
				codeid := js.Get("Codeid").MustString()
				codeid = Transformname(codeid, "", -1)          //解码公司代码和房间号
				coderoom := Transformname(codeid, "", 2)        //房间号
				roomid, _ := strconv.ParseInt(coderoom, 10, 64) //房间号
				sojson.Coderoom = int(roomid)                   //公司代码房间号
				sojson.Codeid = codeid                          //解码公司代码和房间号
				author := new(m.User)
				author.Username = objauthor
				_, err = m.LoadRelatedUser(author, "Username")
				if err == nil {
					sojson.Author = js.Get("Author").MustString()
					rolename := js.Get("AuthorRole").MustString() //角色名称
					utype := js.Get("AuthorTitle").MustString()   //用户类型
					ucss := js.Get("AuthorCss").MustString()      //用户头衔
					sojson.AuthorRole = EncodeB64(rolename)       //用户角色
					sojson.Authortype = EncodeB64(utype)          //用户类别
					sojson.Authorcss = EncodeB64(ucss)            //用户头衔
					sojson.Author = EncodeB64(sojson.Author)
					sojson.Content = EncodeB64(js.Get("Content").MustString())
					sojson.Chat = js.Get("Chat").MustString()
					switch sojson.Chat {
					case "sayhim":
						sojson.Username = js.Get("Username").MustString()
						sojson.Username = EncodeB64(sojson.Username)
						sojson.UserRole = js.Get("UserRole").MustString()
						sojson.UserRole = EncodeB64(sojson.UserRole)
						sojson.Usertype = js.Get("Usertype").MustString()
						sojson.Usertype = EncodeB64(sojson.Usertype)
						sojson.Usercss = js.Get("Usercss").MustString()
						sojson.Usercss = EncodeB64(sojson.Usercss)
					default:
					}
					sojson.Chat = EncodeB64(sojson.Chat)
					sojson.Time = time.Now().Format("15:04")
					sojson.Newtime = time.Now().Format("2006-01-02 15:04:05") //发言时间
					sojson.Datatime = time.Now()
					//数据传给后台
					data := make(map[string]interface{})
					/*
						var sendmsgtext string
						body, err01 := json.Marshal(sojson)
						if err01 == nil {
							sendmsgtext = string(body)
						}
						sendmsgtext = EncodeB64(sendmsgtext)
					*/
					data["msg"] = sojson
					switch DecodeB64(sojson.Chat) {
					case "allchat":
						if _, ok := job.socketiotoroom[codeid]; ok == true {
							so.Emit("all analogsuccess", "")
							so.Emit("all analogmessage", data)
							roleroom := job.socketiotoroom[codeid].Roomval
							rolelen := len(roleroom[0])
							k := 0
							for i := k; i < 2; i++ {
								for j := 0; j < rolelen; j++ {
									so.BroadcastTo(roleroom[i][j], "all analogmessage", data)
								}
							}
							Savechatdata(sojson, 0)
						}
					case "sayhim":
						if _, ok := job.socketiotoroom[codeid]; ok == true {
							so.Emit("all analogsuccess", "")
							so.Emit("all analogmessage", data)
							roleroom := job.socketiotoroom[codeid].Roomval
							rolelen := len(roleroom[0])
							k := 0
							for i := k; i < 2; i++ {
								for j := 0; j < rolelen; j++ {
									so.BroadcastTo(roleroom[i][j], "all analogmessage", data)
								}
							}
							Savechatdata(sojson, 0)
						}
					default:
					}
				} else {
					sendshowmsg := "你没有此权限,请联系管理员!"
					so.Emit("all showmsg", sendshowmsg)
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
				}
				sojson.Author = js.Get("Author").MustString()
				sojson.Content = js.Get("Content").MustString()
				sojson.Ipaddress = js.Get("WebIp").MustString()  //获取ip地址
				sojson.Procities = js.Get("WebPro").MustString() //省市

				codeid := js.Get("Codeid").MustString()
				codeid = Transformname(codeid, "", -1)          //解码公司代码和房间号
				coderoom := Transformname(codeid, "", 2)        //房间号
				roomid, _ := strconv.ParseInt(coderoom, 10, 64) //房间号
				if roomid > 0 {
					sojson.Coderoom = int(roomid) //房间号
					sojson.Codeid = codeid        //解码公司代码和房间号
				}
				author := new(m.User)
				author.Username = sojson.Author
				userUsername, err := m.LoadRelatedUser(author, "Username")
				if err == nil {
					body, err := json.Marshal(userUsername)
					if err == nil {
						js, _ := simplejson.NewJson(body)
						Rolemap := js.Get("Role").MustMap()
						rolename := fmt.Sprintf("%s", Rolemap["Name"])      //角色名称
						rolename = strings.Replace(rolename, " ", "", -1)   //去空格
						utype := fmt.Sprintf("%s", Rolemap["Title"])        //用户类型
						uinsider := fmt.Sprintf("%s", Rolemap["IsInsider"]) //1内部人员或0外部人员
						insider, _ := strconv.ParseInt(uinsider, 10, 64)
						//改动的
						Titlemap := js.Get("Title").MustMap()
						ucss := "/upload/usertitle/" + fmt.Sprintf("%s", Titlemap["Css"])
						sojson.Authorcss = EncodeB64(ucss)
						sojson.AuthorRole = EncodeB64(rolename) //用户角色
						sojson.Authortype = EncodeB64(utype)
						sojson.Chat = EncodeB64("allchat")
						sojson.Author = EncodeB64(sojson.Author)
						sojson.Content = EncodeB64(sojson.Content)
						sojson.Time = time.Now().Format("15:04")
						sojson.Datatime = time.Now()
						sojson.AuthorInSider = int(insider) //1内部人员或0外部人员

						//是否能广播
						uroleid := fmt.Sprintf("%s", Rolemap["Id"]) //发言间隔时间
						roleid, _ := strconv.ParseInt(uroleid, 10, 64)
						nodelist, _ := m.GetUserByRoleId(roleid)
						listlength := len(nodelist)
						for i := 0; i < listlength; i++ {
							if nodelist[i]["Name"] == "sendbroadcast" {
								sojson.IsBroadCast = true
								break
							}
						}
						//var sendmsgtext string
						if sojson.IsBroadCast {
							if "all" == coderoom {
								roominfo, num, _ := m.GetAllRoomDate()
								if num > 0 {
									length := int(num)
									for i := 0; i < length; i++ {
										sojson.Coderoom = int(roominfo[i].RommNumber)
										codeid = roominfo[i].CompanyCode + "_" + fmt.Sprintf("%d", roominfo[i].RommNumber)
										//数据传给后台
										data := make(map[string]interface{})
										/*
											body, err01 := json.Marshal(sojson)
											if err01 == nil {
												sendmsgtext = string(body)
											}
											sendmsgtext = EncodeB64(sendmsgtext)
										*/
										data["msg"] = sojson
										//so.Emit("all broadmsg", data)
										if _, ok := job.socketiotoroom[codeid]; ok == true {
											roleroom := job.socketiotoroom[codeid].Roomval
											rolelen := len(roleroom[0])
											k := 0
											if sojson.AuthorInSider == 0 {
												k = 1
											}
											for i := k; i < 2; i++ {
												for j := 0; j < rolelen; j++ {
													so.BroadcastTo(roleroom[i][j], "all broadmsg", data)
												}
											}
											//保存数据
											Savechatdata(sojson, 0)
											Savechatdata(sojson, 2)
										}
									}
								}
							} else {
								//数据传给后台
								data := make(map[string]interface{})
								/*
									body, err01 := json.Marshal(sojson)
									if err01 == nil {
										sendmsgtext = string(body)
									}
									sendmsgtext = EncodeB64(sendmsgtext)
								*/
								data["msg"] = sojson
								//so.Emit("all broadmsg", data)
								if _, ok := job.socketiotoroom[codeid]; ok == true {
									roleroom := job.socketiotoroom[codeid].Roomval
									rolelen := len(roleroom[0])
									k := 0
									if sojson.AuthorInSider == 0 {
										k = 1
									}
									for i := k; i < 2; i++ {
										for j := 0; j < rolelen; j++ {
											so.BroadcastTo(roleroom[i][j], "all broadmsg", data)
										}
									}
									//保存数据
									Savechatdata(sojson, 0)
									Savechatdata(sojson, 2)
								}
							}
							sendshowmsg := "广播发送成功!"
							so.Emit("all success", sendshowmsg)
						} else {
							//广播消息提示
							sendshowmsg := "你被禁止广播,请联系管理员!"
							so.Emit("all showmsg", sendshowmsg)
						}
					}
				}
			}
		})

		so.On("disconnection", func() {
			if _, ok := job.socketiduser[so.Id()]; ok == true {
				codename := job.socketiduser[so.Id()] //用户名
				if len(codename) > 0 {
					data := make(map[string]interface{})
					if p.Rediserr == nil {
						//用户信息
						codeid := Transformname("", codename, 1) //公司代码用户名互转
						jobroom := "coderoom_" + codeid
						roomdata, _ := p.Client.Get(jobroom)
						if len(roomdata) > 0 {
							userroom, _ := Jsontoroommap(roomdata)
							if _, ok = userroom[codename]; ok == true {
								var sendmsgtext string
								body, err01 := json.Marshal(userroom[codename])
								if err01 == nil {
									sendmsgtext = string(body)
								}
								sendmsgtext = EncodeB64(sendmsgtext)
								data["user"] = sendmsgtext
								if _, ok := job.socketiotoroom[codeid]; ok == true {
									roleroom := job.socketiotoroom[codeid].Roomval
									rolelen := len(roleroom[0])
									for i := 0; i < 2; i++ {
										for j := 0; j < rolelen; j++ {
											so.BroadcastTo(roleroom[i][j], "all disconnection", data)
										}
									}
								}

								if userroom[codename].RoleName != EncodeB64("guest") {
									user := new(m.User)
									user.Username = DecodeB64(userroom[codename].Uname)
									userList, err := m.ReadFieldUser(user, "Username")
									if err == nil {
										onlineunix := time.Now().Unix() - userroom[codename].Datatime.Unix()
										userList.OnlineTime += onlineunix
										userList.UpdateUserFields("OnlineTime")
									}
								}

								delete(userroom, codename)
								jsonRes, _ := json.Marshal(userroom)
								jsonstr := string(jsonRes)
								p.Client.Set(jobroom, jsonstr)

								beego.Info("on disconnect", so.Id(), userroom[codename].Uname)
							}
						}

						//删除map中的内容
						if _, ok = job.socketiduser[so.Id()]; ok == true {
							delete(job.socketiduser, so.Id())
						}
						if _, ok = job.socketidso[codename]; ok == true {
							delete(job.socketidso, codename)
						}
						if _, ok := job.socketidusertime[codename]; ok == true {
							delete(job.socketidusertime, codename)
						}

						if _, ok := job.socketiduserip[codename]; ok == true {
							delete(job.socketiduserip, codename)
						}
						//超过500就删除
						/*
							length := len(job.socketiduserip)
							if length > 1000 {
								for nameip, _ := range job.socketiduserip {
									delete(job.socketiduserip, nameip)
								}
							}
						*/
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
					codeid := js.Get("Codeid").MustString()       //公司房间标识符
					codeid = Transformname(codeid, "", -1)        //解码公司代码和房间号
					codeuser := Transformname(codeid, objname, 0) //公司代码用户名互转
					if _, ok := job.socketidso[codeuser]; ok == true {
						user := new(m.User)
						user.Username = uname
						_, err := m.LoadRelatedUser(user, "Username")
						if err == nil {
							job.socketidso[codeuser].Emit("all kickout", "")
						}
					}
				}
			}
		})

		so.On("all marquuefly", func(msg string) {
			if len(msg) > 0 {
				msg = DecodeB64(msg)
				key := []byte(msg)
				js, err := simplejson.NewJson(key)
				if err != nil {
					beego.Error(err)
				} else {
					content := js.Get("Content").MustString()
					uname := js.Get("Uname").MustString()
					codeid := js.Get("Codeid").MustString()                //公司房间标识符
					codeid = Transformname(codeid, "", -1)                 //解码公司代码和房间号
					codeuser := Transformname(codeid, EncodeB64(uname), 0) //公司代码用户名互转
					if _, ok := job.socketidso[codeuser]; ok == true {
						user := new(m.User)
						user.Username = uname
						_, err := m.LoadRelatedUser(user, "Username")
						if err == nil {
							data := make(map[string]interface{})
							randnano := rand.NewSource(time.Now().UnixNano())
							r := rand.New(randnano)
							index := r.Intn(3)
							data["index"] = index
							data["content"] = EncodeB64(content)
							so.Emit("all marquuefly", data)
							roleroom := job.socketiotoroom[codeid].Roomval
							rolelen := len(roleroom[0])
							for i := 0; i < 2; i++ {
								for j := 0; j < rolelen; j++ {
									so.BroadcastTo(roleroom[i][j], "all marquuefly", data)
								}
							}
						}
					}
				}
			}
		})

		so.On("all robotspeak", func(msg string) {
			dataem := make(map[string]interface{})
			var shuijunval string
			body, err := json.Marshal(Resultuser)
			if err == nil {
				shuijunval = string(body)
			}
			shuijunval = EncodeB64(shuijunval)
			dataem["shuijun"] = shuijunval
			so.Emit("all robotspeak", dataem)
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
		method := this.GetString("method")              //用户方法
		myname := this.GetString("myname")              //用户名
		username := this.GetString("username")          //用户名
		userdata := this.GetString("mydata")            //数据
		codeid := this.GetString("ucodeid")             //公司房间标识符
		codeid = Transformname(codeid, "", -1)          //解码公司代码和房间号
		coderoom := Transformname(codeid, "", 2)        //房间号
		roomid, _ := strconv.ParseInt(coderoom, 10, 64) //房间号
		data := make(map[string]interface{})
		if p.Rediserr == nil {
			switch method {
			case "chatdata":
				{
					//开线程获取省市
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
						socketchat  [2][]Socketjson
						chatlen     [2]int
					)
					switch sysconfig.HistoryMsg { //是否显示历史消息 0显示  1 不显示
					case 0:
						{
							//用户聊天历史记录信息
							jobmark := "chat_new_" + codeid
							jobnewdata, _ := p.Client.Get(jobmark)
							if len(jobnewdata) > 0 {
								socketchat[0], _ = Jsontosocket(jobnewdata)
							}
							chatlen[0] = len(socketchat[0])
							jobmark = "chat_old_" + codeid
							jobolddata, _ := p.Client.Get(jobmark)
							if len(jobolddata) > 0 {
								socketchat[1], _ = Jsontosocket(jobolddata)
							}
							//p.Client.Del(jobmark)
							chatlen[1] = len(socketchat[1])
							length := chatlen[0] + chatlen[1]
							if length > 0 {
								if chatlen[0] == 0 {
									socketchat[0] = socketchat[1]
								} else if chatlen[1] != 0 {
									var newchat []Socketjson
									if chatlen[0] > recordcount {
										jobmark := "chat_new_" + codeid
										p.Client.Del(jobmark)
										jobmark = "chat_old_" + codeid
										p.Client.Set(jobmark, jobnewdata)
										//写数据库
										var chat []m.ChatRecord
										chat = make([]m.ChatRecord, chatlen[0])
										for i := 0; i < chatlen[0]; i++ {
											chat[i].Uname = DecodeB64(socketchat[0][i].Author)
											chat[i].Sendtype = 0
											chat[i].Data = DecodeB64(socketchat[0][i].Content)
											chat[i].Coderoom = int(roomid) //房间号
											chat[i].Ipaddress = socketchat[0][i].Ipaddress
											chat[i].Procities = socketchat[0][i].Procities
											chat[i].Datatime = socketchat[0][i].Datatime
										}
										err = m.AddChatdata(chat, int(chatlen[0]))
										if err != nil {
											beego.Debug(err)
										}
									} else {
										for i := chatlen[0]; i < chatlen[1]; i++ {
											newchat = append(newchat, socketchat[1][i])
										}
										for i := 0; i < chatlen[0]; i++ {
											newchat = append(newchat, socketchat[0][i])
										}
										socketchat[0] = newchat
									}
								} else if chatlen[1] == 0 {
									var newchat []Socketjson
									if chatlen[0] > recordcount {
										jobmark := "chat_new_" + codeid
										p.Client.Del(jobmark)
										//写数据库
										var chat []m.ChatRecord
										chat = make([]m.ChatRecord, chatlen[0])
										for i := 0; i < chatlen[0]; i++ {
											chat[i].Uname = DecodeB64(socketchat[0][i].Author)
											chat[i].Sendtype = 0
											chat[i].Data = DecodeB64(socketchat[0][i].Content)
											chat[i].Coderoom = int(roomid) //房间号
											chat[i].Ipaddress = socketchat[0][i].Ipaddress
											chat[i].Procities = socketchat[0][i].Procities
											chat[i].Datatime = socketchat[0][i].Datatime
											if i-recordcount >= 0 {
												newchat = append(newchat, socketchat[0][i])
											}
										}
										err = m.AddChatdata(chat, int(chatlen[0]))
										if err != nil {
											beego.Debug(err)
										}
										jsonRes, _ := json.Marshal(newchat)
										jobnewdata = string(jsonRes)
										jobmark = "chat_old_" + codeid
										p.Client.Set(jobmark, jobnewdata)
										socketchat[0] = newchat
									}
								}
								historychat = socketchat[0]
							}
						}
					default:
					}
					data["historydata"] = historychat //聊天的历史信息
					//从数据库中获取公告中的最后一条内容
					broaddata, _ := m.GetBroadcastData(int(roomid))
					data["notice"] = broaddata //公告
				}
			case "read":
			case "gagcontrol": //禁言操作
				codeuser := Transformname(codeid, EncodeB64(username), 0) //公司代码用户名互转
				if _, ok := job.socketidso[codeuser]; ok == true {
					if this.GetSession("indexUserInfo") != nil {
						user := new(m.User)
						user.Username = myname
						_, err := m.LoadRelatedUser(user, "Username")
						if err == nil {
							gattime, _ := strconv.ParseInt(userdata, 10, 64) //禁言的时间
							var gag m.GagControl
							gag, err = m.SelectGagControl(username)
							switch gattime {
							case 0:
								if err == nil {
									gag.Uname = myname
									gag.Gagtime = 0
									gag.Gagmode = 0
									gag.Gatstate = 1
									gag.Coderoom = int(roomid) //房间号
									gag.Ipaddress = job.socketiduserip[codeuser].Ip
									gag.Procities = job.socketiduserip[codeuser].Pro
									gag.Datatime = time.Now()
									_, err = m.UpdateGagControl(gag)
									if err != nil {
										beego.Debug(err)
									}
								}
							case 5:
								if err == nil {
									gag.Uname = myname
									gag.Gagtime = time.Now().Unix() + 300
									gag.Gagmode = 300
									gag.Gatcount += 1
									gag.Gatstate = 0
									gag.Coderoom = int(roomid) //房间号
									gag.Ipaddress = job.socketiduserip[codeuser].Ip
									gag.Procities = job.socketiduserip[codeuser].Pro
									gag.Datatime = time.Now()
									_, err = m.UpdateGagControl(gag)
									if err != nil {
										beego.Debug(err)
									}
								} else {
									//添加信息
									gag.Uname = myname
									gag.Objname = username
									gag.Gagtime = time.Now().Unix() + 300
									gag.Gagmode = 300
									gag.Gatcount = 1
									gag.Gatstate = 0
									gag.Coderoom = int(roomid) //房间号
									gag.Ipaddress = job.socketiduserip[codeuser].Ip
									gag.Procities = job.socketiduserip[codeuser].Pro
									gag.Datatime = time.Now()
									_, err = m.AddGagControl(&gag)
									if err != nil {
										beego.Debug(err)
									}
								}
							default:
							}
						}
					}
				}
			case "blacklist": //加入黑名单操作
				codeuser := Transformname(codeid, EncodeB64(username), 0) //公司代码用户名互转
				if _, ok := job.socketidso[codeuser]; ok == true {
					if this.GetSession("indexUserInfo") != nil {
						user := new(m.User)
						user.Username = myname
						_, err := m.LoadRelatedUser(user, "Username")
						if err == nil {
							var black m.BlackList
							black, err = m.SelectBlackList(username)
							if err == nil {
								black.Uname = myname
								switch black.Status {
								case 0:
									black.Status = 1
								case 1:
									black.Status = 0
								default:
								}
								black.Coderoom = int(roomid) //房间号
								black.Ipaddress = job.socketiduserip[codeuser].Ip
								black.Procities = job.socketiduserip[codeuser].Pro
								black.Datatime = time.Now()
								_, err = m.UpdateBlackList(black)
								if err != nil {
									beego.Debug(err)
								}
							} else {
								//添加信息
								black.Uname = myname
								black.Objname = username
								black.Coderoom = int(roomid) //房间号
								black.Ipaddress = job.socketiduserip[codeuser].Ip
								black.Procities = job.socketiduserip[codeuser].Pro
								black.Status = 0
								black.Datatime = time.Now()
								_, err = m.AddBlackList(&black)
								if err != nil {
									beego.Debug(err)
								}
							}
						}
					}
				}
			case "kickout":
				codeuser := Transformname(codeid, EncodeB64(username), 0) //公司代码用户名互转
				if _, ok := job.socketidso[codeuser]; ok == true {
					if this.GetSession("indexUserInfo") != nil {
						user := new(m.User)
						user.Username = myname
						_, err := m.LoadRelatedUser(user, "Username")
						if err == nil {
							var kickout m.KickOut
							kickout, err = m.SelectKickOut(username)
							if err == nil {
								kickout.Uname = myname
								kickout.Kicktime = time.Now().Unix() + 3600
								kickout.Ipaddress = job.socketiduserip[codeuser].Ip
								kickout.Procities = job.socketiduserip[codeuser].Pro
								kickout.Status = 0
								kickout.Coderoom = int(roomid) //房间号
								kickout.Datatime = time.Now()
								_, err = m.UpdateKickControl(kickout)
								if err != nil {
									beego.Debug(err)
								}
							} else {
								//添加信息
								kickout.Uname = myname
								kickout.Objname = username
								kickout.Kicktime = time.Now().Unix() + 3600
								kickout.Ipaddress = job.socketiduserip[codeuser].Ip
								kickout.Procities = job.socketiduserip[codeuser].Pro
								kickout.Status = 0
								kickout.Coderoom = int(roomid) //房间号
								kickout.Datatime = time.Now()
								_, err = m.AddKickOut(&kickout)
								if err != nil {
									beego.Debug(err)
								}
							}
						}
					}
				}
			case "checkrole":
				codeuser := Transformname(codeid, EncodeB64(username), 0) //公司代码用户名互转
				if _, ok := job.socketidso[codeuser]; ok == true {
					var checkresult = "no"
					if this.GetSession("indexUserInfo") != nil {
						user := new(m.User)
						user.Username = myname
						userUsername, err := m.LoadRelatedUser(user, "Username")
						if err == nil {
							body, err := json.Marshal(userUsername)
							if err == nil {
								js, _ := simplejson.NewJson(body)
								Rolemap := js.Get("Role").MustMap()
								uweight := fmt.Sprintf("%s", Rolemap["Weight"])
								myweight, _ := strconv.ParseInt(uweight, 10, 64) //用户权重

								user.Username = username
								userUsername, err = m.LoadRelatedUser(user, "Username")
								if err == nil {
									body, err = json.Marshal(userUsername)
									if err == nil {
										js, _ = simplejson.NewJson(body)
										Rolemap = js.Get("Role").MustMap()
										uobjweight := fmt.Sprintf("%s", Rolemap["Weight"])
										objweight, _ := strconv.ParseInt(uobjweight, 10, 64) //用户权重
										if myweight > objweight {
											checkresult = "yes"
										} else {
											checkresult = "no"
										}
									}
								} else {
									checkresult = "yes"
								}
							}
						}
					}
					data["msg"] = checkresult
				}
			case "leave":
				if this.GetSession("indexUserInfo") != nil {
					this.DelSession("indexUserInfo")
				}
			default:
			}
		}
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

func (this *SocketController) ChatUpload() {
	if "POST" == this.Ctx.Request.Method {
		if this.GetSession("indexUserInfo") == nil {
			this.Ctx.Redirect(302, "/")
			return
		}
		file, fh, err := this.GetFile("uploadfilelib")
		if err == nil {
			file.Close() //关闭文件
			uploadfilename := fh.Filename
			index := strings.Index(uploadfilename, ".")
			if -1 != index {
				randnum := time.Now().UnixNano()
				uploadfilename = fmt.Sprintf("%d", randnum) + uploadfilename[index:]
				fileloadurl := path.Join("..", "upload", "img", uploadfilename)
				if len(fh.Filename) > 0 {
					err = this.SaveToFile("uploadfilelib", fileloadurl)
					if err != nil {
						//上传文件消息提示
						this.Rsp(false, "图片上传失败", "")
					} else {
						this.Rsp(true, "图片上传成功", fileloadurl)
					}
				} else {
					this.Rsp(false, "图片上传失败", "")
				}
			} else {
				this.Rsp(false, "图片上传失败", "")
			}
		} else {
			this.Rsp(false, "图片上传失败", "")
		}
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

func (this *SocketController) ChatKickOut() {
	this.TplName = "haoindex/kickout.html"
}

func (this *SocketController) ChatModifyIcon() {
	if this.GetSession("indexUserInfo") != nil && this.IsAjax() {
		uname := EncodeB64(this.GetString("uname"))
		uicon := EncodeB64(this.GetString("uicon"))
		codeid := this.GetString("ucodeid")         //公司房间标识符
		codeid = Transformname(codeid, "", -1)      //解码公司代码和房间号
		codename := Transformname(codeid, uname, 0) //公司代码用户名互转
		if p.Rediserr == nil {
			var (
				userroom  map[string]Usertitle
				usertitle Usertitle
			)
			//用户列表信息
			userroom = make(map[string]Usertitle) //房间对应的用户信息
			jobroom := "coderoom_" + codeid
			roomdata, _ := p.Client.Get(jobroom)
			if len(roomdata) > 0 {
				userroom, _ = Jsontoroommap(roomdata)
				usertitle = userroom[codename]
				usertitle.UserIcon = uicon
				userroom[codename] = usertitle
				jsonRes, _ := json.Marshal(userroom)
				jsonstr := string(jsonRes)
				p.Client.Set(jobroom, jsonstr)
			}
		}
		data := make(map[string]interface{})
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Ctx.Redirect(302, "/")
	}
}

//保存数据
func Savechatdata(sojson Socketjson, sel int64) {
	if p.Rediserr == nil {
		switch sel {
		case 0:
			{
				var (
					socketchat []Socketjson
					jsonstr    string
				)
				jobmark := "chat_new_" + sojson.Codeid
				jobdata, _ := p.Client.Get(jobmark)
				if len(jobdata) > 0 {
					socketchat, _ = Jsontosocket(jobdata)
				}
				socketchat = append(socketchat, sojson)
				jsonRes, _ := json.Marshal(socketchat)
				jsonstr = string(jsonRes)
				p.Client.Set(jobmark, jsonstr)
				length := len(socketchat)
				if length == recordcount {
					p.Client.Del(jobmark)
					jobmark = "chat_old_" + sojson.Codeid
					p.Client.Set(jobmark, jsonstr)
					//写数据库
					var chat []m.ChatRecord
					chat = make([]m.ChatRecord, recordcount)
					for i := 0; i < recordcount; i++ {
						chat[i].Uname = DecodeB64(socketchat[i].Author)
						chat[i].Sendtype = 0
						chat[i].Data = DecodeB64(socketchat[i].Content)
						chat[i].Coderoom = socketchat[i].Coderoom
						chat[i].Ipaddress = socketchat[i].Ipaddress
						chat[i].Procities = socketchat[i].Procities
						chat[i].Datatime = socketchat[i].Datatime
						id, errchat := m.AddChat(&chat[i])
						if errchat == nil {
							socketchat[i].Id = id
						} else {
							beego.Debug(err)
						}
					}
				}
			}
		case 1:
			{
				//写数据库
				var chat m.ChatRecord
				chat.Uname = DecodeB64(sojson.Author)
				chat.Objname = DecodeB64(sojson.Username)
				chat.Sendtype = 1
				chat.Data = DecodeB64(sojson.Content)
				chat.Coderoom = sojson.Coderoom
				chat.Ipaddress = sojson.Ipaddress
				chat.Procities = sojson.Procities
				chat.Datatime = time.Now()
				_, err = m.AddChat(&chat)
				if err != nil {
					beego.Debug(err)
				}
			}
		case 2:
			{
				//写数据库
				var broad m.Broadcast
				broad.Uname = DecodeB64(sojson.Author)
				broad.Data = DecodeB64(sojson.Content)
				broad.Coderoom = sojson.Coderoom
				broad.Ipaddress = sojson.Ipaddress
				broad.Procities = sojson.Procities
				broad.Datatime = time.Now()
				_, err = m.AddBroadcast(&broad)
				if err != nil {
					beego.Debug(err)
				}
			}
		default:
		}
	}
}

//数据解析
func AnalyticContentdata(content string) string {
	return content
}
