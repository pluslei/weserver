package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"weserver/controllers/haoadmin"
	"weserver/controllers/haoindex"
	"weserver/controllers/haophone"
	s "weserver/src/rpcserver"
	"weserver/src/socket"
)

func init() {
	// 注册路由
	Router()
	// 注册RPC
	Rpc()
	// 开启调试模式
	orm.Debug = true
	beego.SetStaticPath("/upload", "../upload")
	beego.SetStaticPath("/css", "./views/dist/css")
	beego.SetStaticPath("/i", "./views/dist/i")
	beego.SetStaticPath("/js", "./views/dist/js")
	beego.SetStaticPath("/fonts", "./views/dist/fonts")
	/*
		beego.SetViewsPath("../weclient/dist")
		beego.SetStaticPath("/css", "../weclient/dist/css")
		beego.SetStaticPath("/i", "../weclient/dist/i")
		beego.SetStaticPath("/js", "../weclient/dist/js")
		beego.SetStaticPath("/fonts", "../weclient/dist/fonts")
	*/
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

	// rbac
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

	//在线用户
	beego.Router("/weserver/user/onlineuser", &haoadmin.UserController{}, "*:Onlineuser")
	beego.Router("/weserver/user/verifyuser", &haoadmin.UserController{}, "*:Rerifyuser")

	// 节点管理
	beego.Router("/weserver/node/addnode", &haoadmin.NodeController{}, "*:AddNode")
	beego.Router("/weserver/node/updatenode", &haoadmin.NodeController{}, "*:UpdateNode")
	beego.Router("/weserver/node/getnodetree", &haoadmin.NodeController{}, "*:GetNodeTree")
	beego.Router("/weserver/node/delnode", &haoadmin.NodeController{}, "*:DelNode")
	beego.Router("/weserver/node/index", &haoadmin.NodeController{}, "*:Index")
	beego.Router("/weserver/node/getNode", &haoadmin.NodeController{}, "*:GetNode")

	// 分组管理
	beego.Router("/weserver/group/addgroup", &haoadmin.GroupController{}, "*:AddGroup")
	beego.Router("/weserver/group/updategroup", &haoadmin.GroupController{}, "*:UpdateGroup")
	beego.Router("/weserver/group/delgroup", &haoadmin.GroupController{}, "*:DelGroup")
	beego.Router("/weserver/group/index", &haoadmin.GroupController{}, "*:Index")

	beego.Router("/weserver/title/addtitle", &haoadmin.TitleController{}, "*:AddTitle")
	beego.Router("/weserver/title/updatetitle", &haoadmin.TitleController{}, "*:UpdateTitle")
	beego.Router("/weserver/title/deltitle", &haoadmin.TitleController{}, "*:DelTitle")
	beego.Router("/weserver/title/index", &haoadmin.TitleController{}, "*:Index")
	beego.Router("/weserver/title/getalltitle", &haoadmin.TitleController{}, "*:GetAllTitle")
	beego.Router("/weserver/title/upload", &haoadmin.TitleController{}, "*:UploadTitle")

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
	// 主题设置
	beego.Router("/weserver/sysconfig/theme_index", &haoadmin.ThemeController{}, "*:Index")
	beego.Router("/weserver/sysconfig/theme_addtheme", &haoadmin.ThemeController{}, "*:AddTheme")
	beego.Router("/weserver/sysconfig/theme_updatetheme", &haoadmin.ThemeController{}, "*:UpdateTheme")
	beego.Router("/weserver/sysconfig/theme_deltheme", &haoadmin.ThemeController{}, "*:DelTheme")

	// 数据管理
	beego.Router("/weserver/data/room_index", &haoadmin.RoomController{}, "*:Index")
	beego.Router("/weserver/data/room_updateroom", &haoadmin.RoomController{}, "*:UpdateRoom")
	beego.Router("/weserver/data/room_refresh", &haoadmin.RoomController{}, "*:RefreshRoom")

	beego.Router("/weserver/data/chatrecord", &haoadmin.ChatRecordController{}, "*:ChatRecordList")

	// 讲师介绍
	beego.Router("/weserver/data/teacher_index", &haoadmin.TeacherController{}, "*:Index")
	beego.Router("/weserver/data/teacher_addteacher", &haoadmin.TeacherController{}, "*:AddTeacher")
	beego.Router("/weserver/data/teacher_updateteacher", &haoadmin.TeacherController{}, "*:UpdateTeacher")
	beego.Router("/weserver/data/teacher_delteacher", &haoadmin.TeacherController{}, "*:DelTeacher")
	beego.Router("/weserver/data/teacher_upload", &haoadmin.TeacherController{}, "*:UploadTeacher")
	// 问题解答
	beego.Router("/weserver/data/qs_index", &haoadmin.QsController{}, "*:Index")
	beego.Router("/weserver/data/qs_addqs", &haoadmin.QsController{}, "*:AddQs")
	beego.Router("/weserver/data/qs_updateqs", &haoadmin.QsController{}, "*:UpdateQs")
	beego.Router("/weserver/data/qs_delqs", &haoadmin.QsController{}, "*:DelQs")
	// 广播
	beego.Router("/weserver/data/qs_broad", &haoadmin.QsController{}, "*:DataBroadQs")
	beego.Router("/weserver/data/sendbroad", &haoadmin.QsController{}, "*:SendBroad")
	// 消息管理
	beego.Router("/weserver/data/message_index", &haoadmin.MessageController{}, "*:MessageList")
	beego.Router("/weserver/data/message_add", &haoadmin.MessageController{}, "*:AddMessage")
	beego.Router("/weserver/data/message_byid", &haoadmin.MessageController{}, "*:GetMessageById")
	beego.Router("/weserver/data/message_edit", &haoadmin.MessageController{}, "*:EditMessage")
	beego.Router("/weserver/data/message_delete", &haoadmin.MessageController{}, "*:DeleteMessage")
	// 消息分类管理
	beego.Router("/weserver/data/messagetype_index", &haoadmin.MessageTypeController{}, "*:MessageTypeList")
	beego.Router("/weserver/data/messagetype_add", &haoadmin.MessageTypeController{}, "*:AddMessageType")
	beego.Router("/weserver/data/messagetype_delete", &haoadmin.MessageTypeController{}, "*:DeleteMessageType")
	beego.Router("/weserver/data/messagetype_byid", &haoadmin.MessageTypeController{}, "*:GetMessageTypeById")
	beego.Router("/weserver/data/messagetype_edit", &haoadmin.MessageTypeController{}, "*:EditMessageType")
	//机器人发言
	beego.Router("/weserver/data/robot_speak", &haoadmin.QsController{}, "*:DataRobotSpeak")

	beego.Router("/weserver/data/face_index", &haoadmin.FaceController{}, "*:List")
	beego.Router("/weserver/data/face_getfaces", &haoadmin.FaceController{}, "*:GetFaces")
	beego.Router("/weserver/data/face_delete", &haoadmin.FaceController{}, "*:DeleteFace")
	beego.Router("/weserver/data/face_getface", &haoadmin.FaceController{}, "*:GetFace")
	beego.Router("/weserver/data/face_upload", &haoadmin.FaceController{}, "*:Upload")
	beego.Router("/weserver/data/face_submit", &haoadmin.FaceController{}, "*:AddOrEdit")
	beego.Router("/weserver/data/face_getMaxValue", &haoadmin.FaceController{}, "*:GetMaxGroupValue")

	// 首页管理
	beego.Router("/weserver/home/aboutme", &haoadmin.HomeController{}, "*:AboutUs")
	beego.Router("/weserver/home/contact", &haoadmin.HomeController{}, "*:ContactUs")
	// 课程管理
	beego.Router("/weserver/home/course_index", &haoadmin.CourseController{}, "*:Index")
	beego.Router("/weserver/home/course_addcourse", &haoadmin.CourseController{}, "*:AddCourse")
	beego.Router("/weserver/home/course_updatecourse", &haoadmin.CourseController{}, "*:UpdateCourse")
	beego.Router("/weserver/home/course_delcourse", &haoadmin.CourseController{}, "*:DelCourse")
	beego.Router("/weserver/home/course_coursejson", &haoadmin.CourseController{}, "*:GetCourseJson")
	// 客服管理
	beego.Router("/weserver/home/custservice_index", &haoadmin.CustserviceController{}, "*:Index")
	beego.Router("/weserver/home/custservice_addcust", &haoadmin.CustserviceController{}, "*:AddCust")
	beego.Router("/weserver/home/custservice_updatecust", &haoadmin.CustserviceController{}, "*:UpdateCust")
	beego.Router("/weserver/home/custservice_delcust", &haoadmin.CustserviceController{}, "*:DelCust")
	// 首页讲师管理
	beego.Router("/weserver/home/teachbanner_index", &haoadmin.TeachBannerController{}, "*:Index")
	beego.Router("/weserver/home/teachbanner_addbanner", &haoadmin.TeachBannerController{}, "*:AddBanner")
	beego.Router("/weserver/home/teachbanner_updatebanner", &haoadmin.TeachBannerController{}, "*:UpdateBanner")
	beego.Router("/weserver/home/teachbanner_delbanner", &haoadmin.TeachBannerController{}, "*:DelBanner")
	beego.Router("/weserver/home/teachbanner_upload", &haoadmin.TeachBannerController{}, "*:UploadBanner")
	//首页客服电话管理
	beego.Router("/weserver/home/telbanner_index", &haoadmin.TelBannerController{}, "*:TelBannerIndex")
	beego.Router("/weserver/data/telbanner_upload", &haoadmin.TelBannerController{}, "*:UploadTelBanner")
	// 水军
	beego.Router("/weserver/data/robot_speak", &haoadmin.QsController{}, "*:DataRobotSpeak")

	//首页财经新闻
	// beego.Router("/finance_index", &haoadmin.FinanceNewsController{}, "*:GetFinanceNews")
	beego.Router("/online/public/index", &haoadmin.MainController{}, "*:OnlineIndex")

	// 测试
	beego.Router("/test", &haoadmin.TestController{}, "*:Test")
	beego.Router("/test/postapi", &haoadmin.TestController{}, "*:PostApi")

	// index
	beego.Router("/", &haoindex.IndexController{})
	beego.Router("/?:id([0-9]+)", &haoindex.IndexController{}, "*:Index")
	beego.Router("/face", &haoindex.FaceController{}, "*:Add")
	beego.Router("/index", &haoindex.IndexController{}, "*:Index")
	beego.Router("/voice", &haoindex.IndexController{}, "*:Voice")
	beego.Router("/mediaurl", &haoindex.IndexController{}, "*:GetMediaURL")
	beego.Router("/setnickname", &haoindex.IndexController{}, "*:SetNickname")

	//socket
	// 获取系统信息
	beego.Router("/chat/user/list", &socket.SocketController{}, "*:ChatUserList")
	// 获取在线人数
	beego.Router("/chat/user/online/msg", &socket.SocketController{}, "*:ChatOnlineUserMsg")
	// 以下暂时没用
	beego.Router("/chat/modify/icon", &socket.SocketController{}, "*:ChatModifyIcon")
	beego.Router("/chat/upload", &socket.SocketController{}, "*:ChatUpload")
	beego.Router("/chat/kickout", &socket.SocketController{}, "*:ChatKickOut")

	// 添加网名
	beego.Router("/weserver/data/netname_add", &haoadmin.NetNameController{}, "*:AddNetName")

	//手机的访问路由
	beego.Router("/phone/index", &haophone.LoginController{}, "*:PhoneIndex")
	beego.Router("/phone/login", &haophone.LoginController{}, "*:PhoneLogin")
	beego.Router("/phone/logout", &haophone.LoginController{}, "*:PhoneLogout")
	beego.Router("/phone/sendcode", &haophone.LoginController{}, "*:SendCode")
	beego.Router("/phone/register", &haophone.LoginController{}, "*:PhoneRegister")
	beego.Router("/phone/getgroups", &haophone.LoginController{}, "*:GetPhoneGroups")
	beego.Router("/phone/getfaces", &haophone.LoginController{}, "*:GetPhoneFaces")
}
