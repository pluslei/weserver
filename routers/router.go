package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"weserver/controllers/haoadmin"
	"weserver/controllers/haoindex"
	"weserver/controllers/mqtt"
	s "weserver/src/rpcserver"
)

func init() {
	// 注册路由
	Router()
	// 注册RPC
	Rpc()
	// 开启调试模式
	orm.Debug = false
	beego.SetStaticPath("/upload", "../upload")
	beego.SetStaticPath("/css", "./views/dist/css")
	beego.SetStaticPath("/i", "./views/dist/i")
	beego.SetStaticPath("/js", "./views/dist/js")
	beego.SetStaticPath("/fonts", "./views/dist/fonts")

	// beego.SetViewsPath("../weclient/dist")
	// beego.SetStaticPath("/css", "../weclient/dist/css")
	// beego.SetStaticPath("/i", "../weclient/dist/i")
	// beego.SetStaticPath("/js", "../weclient/dist/js")
	// beego.SetStaticPath("/fonts", "../weclient/dist/fonts")

}

// 路由必须三个/以上
// 路由中无需验证的请见如 app.conf -> not_auth_package 以逗号隔开
func Rpc() {
	s.Server.Publish("chat", 0, 0)
	s.Server.Event = s.Event{}

	beego.Handler("/rpc", s.Server)
}

func Router() {
	// public
	beego.Router("/haoindex", &haoadmin.MainController{}, "*:Index")
	beego.Router("/public/login", &haoadmin.MainController{}, "*:Login")
	beego.Router("/public/logout", &haoadmin.MainController{}, "*:Logout")
	beego.Router("/public/updateadmin", &haoadmin.MainController{}, "*:UpdateAdminIndex")
	beego.Router("/public/updatepwd", &haoadmin.MainController{}, "*:UpdateAdminPwd")

	beego.Router("/weserver/public/index", &haoadmin.MainController{}, "*:Index")
	beego.Router("/weserver/public/changepwd", &haoadmin.MainController{}, "*:Changepwd")

	beego.Router("/weserver/user/adduser", &haoadmin.UserController{}, "*:AddUser")
	beego.Router("/weserver/user/updateuser", &haoadmin.UserController{}, "*:UpdateUser")
	beego.Router("/weserver/user/deluser", &haoadmin.UserController{}, "*:DelUser")
	beego.Router("/weserver/user/index", &haoadmin.UserController{}, "*:Index")
	beego.Router("/weserver/user/usertorole", &haoadmin.UserController{}, "*:UserToRole")
	beego.Router("/weserver/user/setusertitle", &haoadmin.UserController{}, "*:SetUserTitle")
	beego.Router("/weserver/user/usertotitle", &haoadmin.UserController{}, "*:UserToTitle")
	beego.Router("/weserver/user/regstatus", &haoadmin.UserController{}, "*:UpdateRegStatus")
	beego.Router("/weserver/user/userstatus", &haoadmin.UserController{}, "*:UpdateUserStatus")
	beego.Router("/weserver/user/kictuser", &haoadmin.UserController{}, "*:KictUser")
	beego.Router("/weserver/user/preparedel", &haoadmin.UserController{}, "*:PrepareDelUser")
	beego.Router("/weserver/user/onlineuser", &haoadmin.UserController{}, "*:Onlineuser")

	// 节点管理
	// beego.Router("/weserver/node/addnode", &haoadmin.NodeController{}, "*:AddNode")
	// beego.Router("/weserver/node/updatenode", &haoadmin.NodeController{}, "*:UpdateNode")
	// beego.Router("/weserver/node/getnodetree", &haoadmin.NodeController{}, "*:GetNodeTree")
	// beego.Router("/weserver/node/delnode", &haoadmin.NodeController{}, "*:DelNode")
	// beego.Router("/weserver/node/index", &haoadmin.NodeController{}, "*:Index")
	// beego.Router("/weserver/node/getNode", &haoadmin.NodeController{}, "*:GetNode")
	// 节点分组分组管理
	// beego.Router("/weserver/group/addgroup", &haoadmin.GroupController{}, "*:AddGroup")
	// beego.Router("/weserver/group/updategroup", &haoadmin.GroupController{}, "*:UpdateGroup")
	// beego.Router("/weserver/group/delgroup", &haoadmin.GroupController{}, "*:DelGroup")
	// beego.Router("/weserver/group/index", &haoadmin.GroupController{}, "*:Index")

	// 头衔
	beego.Router("/weserver/title/addtitle", &haoadmin.TitleController{}, "*:AddTitle")
	beego.Router("/weserver/title/updatetitle", &haoadmin.TitleController{}, "*:UpdateTitle")
	beego.Router("/weserver/title/deltitle", &haoadmin.TitleController{}, "*:DelTitle")
	beego.Router("/weserver/title/index", &haoadmin.TitleController{}, "*:Index")
	beego.Router("/weserver/title/getalltitle", &haoadmin.TitleController{}, "*:GetAllTitle")
	beego.Router("/weserver/title/upload", &haoadmin.TitleController{}, "*:UploadTitle")

	// 角色
	beego.Router("/weserver/role/delrole", &haoadmin.RoleController{}, "*:DelRole")
	beego.Router("/weserver/role/index", &haoadmin.RoleController{}, "*:Index")
	beego.Router("/weserver/role/addrole", &haoadmin.RoleController{}, "*:AddRole")
	beego.Router("/weserver/role/updaterole", &haoadmin.RoleController{}, "*:UpdataRole")
	beego.Router("/weserver/role/addaccess", &haoadmin.RoleController{}, "*:AddAccess")
	beego.Router("/weserver/role/accesstonode", &haoadmin.RoleController{}, "*:AccessToNode")
	beego.Router("/weserver/role/getallrole", &haoadmin.RoleController{}, "*:GetAllRole")
	beego.Router("/weserver/role/upload", &haoadmin.RoleController{}, "*:Upload")

	// 全局设置
	beego.Router("/weserver/sysconfig/index", &haoadmin.SysConfigController{}, "*:Index")

	// 历史消息
	beego.Router("/weserver/data/chatrecord", &haoadmin.ChatRecordController{}, "*:ChatRecordList")

	// 广播
	beego.Router("/weserver/data/qs_broad", &haoadmin.QsController{}, "*:BroadList")
	beego.Router("/weserver/data/sendbroad", &haoadmin.QsController{}, "*:SendBroad")

	// 测试
	beego.Router("/test", &haoadmin.TestController{}, "*:Test")
	beego.Router("/test/postapi", &haoadmin.TestController{}, "*:PostApi")

	// index
	beego.Router("/", &haoindex.IndexController{})
	beego.Router("/?:id([0-9]+)", &haoindex.IndexController{}, "*:Index")
	beego.Router("/index", &haoindex.IndexController{}, "*:Index")
	beego.Router("/voice", &haoindex.IndexController{}, "*:Voice")
	beego.Router("/mediaurl", &haoindex.IndexController{}, "*:GetMediaURL")
	beego.Router("/setnickname", &haoindex.IndexController{}, "*:SetNickname")

	beego.Router("/chat/user/message", &mqtt.MqttController{}, "*:GetMessageToSend")
	beego.Router("/chat/user/historylist", &mqtt.MqttController{}, "*:GetChatHistoryList")
	beego.Router("/chat/user/online/passid", &mqtt.MqttController{}, "*:GetPassId")
	//获取在线人数信息
	beego.Router("/chat/user/online/info", &mqtt.MqttController{}, "*:GetOnlineUseInfo")
	// 获取在线人数
	beego.Router("/chat/user/online/count", &mqtt.MqttController{}, "*:GetOnlineUseCount")
	// 以下暂时没用
	beego.Router("/chat/modify/icon", &mqtt.MqttController{}, "*:ChatModifyIcon")
	beego.Router("/chat/upload", &mqtt.MqttController{}, "*:ChatUpload")
	beego.Router("/chat/kickout", &mqtt.MqttController{}, "*:ChatKickOut")

}
