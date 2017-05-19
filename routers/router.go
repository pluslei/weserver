package routers

import (
	"weserver/controllers/haoadmin"
	"weserver/controllers/haoindex"
	"weserver/controllers/mqtt"
	"weserver/models"
	s "weserver/src/rpcserver"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	models.AccessRegister()
}

// 路由必须三个/以上
// 路由中无需验证的请见如 app.conf -> not_auth_package 以逗号隔开
func Rpc() {
	s.Server.Publish("chat", 0, 0)
	s.Server.Event = s.Event{}

	beego.Handler("/rpc", s.Server)
}

func Router() {
	//common
	beego.Router("/weserver/common/getCompanyInfo", &haoadmin.CommonController{}, "*:GetCompanyInfo")
	beego.Router("/weserver/common/getRoomInfo", &haoadmin.CommonController{}, "*:GetRoomInfoByCompanyId")

	// public
	beego.Router("/haoindex", &haoadmin.MainController{}, "*:Index")
	beego.Router("/public/login", &haoadmin.MainController{}, "*:Login")
	beego.Router("/public/logout", &haoadmin.MainController{}, "*:Logout")
	beego.Router("/public/updateadmin", &haoadmin.MainController{}, "*:UpdateAdminIndex")
	beego.Router("/public/updatepwd", &haoadmin.MainController{}, "*:UpdateAdminPwd")

	beego.Router("/weserver/public/index", &haoadmin.MainController{}, "*:Index")
	beego.Router("/weserver/public/changepwd", &haoadmin.MainController{}, "*:Changepwd")
	// beego.Router("/weserver/public/choosecompany", &haoadmin.UserController{}, "*:ChooseCompany")
	beego.Router("/weserver/user/adduser", &haoadmin.UserController{}, "*:AddUser") 
	beego.Router("/weserver/user/updateuser", &haoadmin.UserController{}, "*:UpdateUser")
	beego.Router("/weserver/user/deluser", &haoadmin.UserController{}, "*:DelUser")
	beego.Router("/weserver/user/delregisteruser", &haoadmin.UserController{}, "*:DelRegistUser")
	beego.Router("/weserver/user/index", &haoadmin.UserController{}, "*:Index") 
	beego.Router("/weserver/user/", &haoadmin.UserController{}, "*:Index") 
	beego.Router("/weserver/user/usertorole", &haoadmin.UserController{}, "*:UserToRole")
	beego.Router("/weserver/user/setusertitle", &haoadmin.UserController{}, "*:SetUserTitle")
	beego.Router("/weserver/user/usertotitle", &haoadmin.UserController{}, "*:UserToTitle")
	beego.Router("/weserver/user/regstatus", &haoadmin.UserController{}, "*:UpdateRegStatus")
	beego.Router("/weserver/user/userstatus", &haoadmin.UserController{}, "*:UpdateStatus")
	beego.Router("/weserver/user/kictuser", &haoadmin.UserController{}, "*:KictUser")
	beego.Router("/weserver/user/preparedel", &haoadmin.UserController{}, "*:PrepareDelUser")
	beego.Router("/weserver/user/preparedelregistuser", &haoadmin.UserController{}, "*:PrepareDelRegistUser")

	beego.Router("/weserver/user/usersetlist", &haoadmin.UserController{}, "*:UserList")
	beego.Router("/weserver/user/setusername", &haoadmin.UserController{}, "*:SetUsername")

	//解除禁言
	beego.Router("/weserver/user/UnShutUp", &haoadmin.UserController{}, "*:SetUnShutUp")
	//beego.Router("/weserver/user/onlineuser", &haoadmin.UserController{}, "*:Onlineuser")

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
	beego.Router("/weserver/data/checkrecord", &haoadmin.ChatRecordController{}, "*:CheckRecord")
	beego.Router("/weserver/data/delrecord", &haoadmin.ChatRecordController{}, "*:DelRecord")

	// 公告消息
	beego.Router("/weserver/data/qs_broad", &haoadmin.QsController{}, "*:SendNoticeList")
	beego.Router("/weserver/data/sendbroad", &haoadmin.QsController{}, "*:SendBroad")
	beego.Router("/weserver/data/notice_edit", &haoadmin.QsController{}, "*:Edit")
	beego.Router("/weserver/data/notice_del", &haoadmin.QsController{}, "*:Del")

	// 房间管理
	beego.Router("/weserver/data/room_index", &haoadmin.RoomController{}, "*:Index")
	beego.Router("/weserver/data/room_add", &haoadmin.RoomController{}, "*:Add")
	beego.Router("/weserver/data/room_edit", &haoadmin.RoomController{}, "*:Edit")
	beego.Router("/weserver/data/room_del", &haoadmin.RoomController{}, "*:Del")
	beego.Router("/weserver/data/upload", &haoadmin.RoomController{}, "*:Upload")

	//公司信息管理
	beego.Router("/weserver/data/company", &haoadmin.CompanyController{}, "*:Index")            //公司信息展示
	beego.Router("/weserver/data/company_add", &haoadmin.CompanyController{}, "*:AddCompany")   //添加公司
	beego.Router("/weserver/data/company_del", &haoadmin.CompanyController{}, "*:DelCompany")   //删除公司
	beego.Router("/weserver/data/company_edit", &haoadmin.CompanyController{}, "*:EditCompany") //编辑公司

	//纸条提问
	beego.Router("weserver/data/question", &haoadmin.QuestionController{}, "*:QuestionList") //纸条提问列表
	beego.Router("weserver/data/question_reply", &haoadmin.QuestionController{}, "*:QuestionReply") //回复纸条提问
	beego.Router("weserver/data/question_del", &haoadmin.QuestionController{}, "*:QuestionDel") //删除纸条提问
	// 策略管理
	beego.Router("/weserver/data/strategy_index", &haoadmin.StrategyController{}, "*:Index")
	beego.Router("/weserver/data/strategy_edit", &haoadmin.StrategyController{}, "*:Edit")
	beego.Router("/weserver/data/strategy_add", &haoadmin.StrategyController{}, "*:Add")
	beego.Router("/weserver/data/strategy_del", &haoadmin.StrategyController{}, "*:Del")

	// 讲师管理
	beego.Router("/weserver/data/teacher_index", &haoadmin.TeacherController{}, "*:Index")
	beego.Router("/weserver/data/teacher_add", &haoadmin.TeacherController{}, "*:Add")
	beego.Router("/weserver/data/teacher_edit", &haoadmin.TeacherController{}, "*:Edit")
	beego.Router("/weserver/data/teacher_del", &haoadmin.TeacherController{}, "*:Del")
	beego.Router("/weserver/data/teacher_room", &haoadmin.TeacherController{}, "*:GetTeacher")

	// 操作建议
	beego.Router("/weserver/data/suggest_index", &haoadmin.SuggestController{}, "*:Index")
	beego.Router("/weserver/data/suggest_add", &haoadmin.SuggestController{}, "*:Add")
	beego.Router("/weserver/data/suggest_edit", &haoadmin.SuggestController{}, "*:Edit")
	beego.Router("/weserver/data/suggest_del", &haoadmin.SuggestController{}, "*:Del")
	beego.Router("/weserver/data/suggest_addclose", &haoadmin.SuggestController{}, "*:AddClose")
	beego.Router("/weserver/data/suggest_getclose", &haoadmin.SuggestController{}, "*:GetClose")
	beego.Router("/weserver/data/suggest_editclose", &haoadmin.SuggestController{}, "*:EditClose")
	beego.Router("/weserver/data/suggest_delclose", &haoadmin.SuggestController{}, "*:DelClose")

	// 测试
	// beego.Router("/test", &haoadmin.TestController{}, "*:Test")
	// beego.Router("/test/postapi", &haoadmin.TestController{}, "*:PostApi")

	//*********************************************************************************************
	// 前端
	beego.Router("/", &haoindex.IndexController{}, "*:Redirectr")
	beego.Router("/login", &haoindex.IndexController{}, "*:WeChatLogin")
	beego.Router("/pclogin", &haoindex.IndexController{}, "*:PCLogin")
	beego.Router("/loginhandle", &haoindex.IndexController{}, "*:LoginHandle")
	beego.Router("/wechat", &haoindex.IndexController{}, "*:GetWeChatInfo")

	// beego.Router("/?:id([0-9]+)", &haoindex.IndexController{}, "*:Index")
	beego.Router("/index", &haoindex.IndexController{}, "*:Index")
	beego.Router("/voice", &haoindex.IndexController{}, "*:Voice")
	beego.Router("/mediaurl", &haoindex.IndexController{}, "*:GetMediaURL")
	beego.Router("/setnickname", &haoindex.IndexController{}, "*:SetNickname")
	beego.Router("/wx", &haoindex.IndexController{}, "*:WxServerImg")

	//聊天管理
	beego.Router("/chat/user/roominfo", &mqtt.MqttController{}, "*:GetCompanyInfo")
	beego.Router("/chat/user/message", &mqtt.MqttController{}, "*:GetMessageToSend")
	beego.Router("/chat/user/historylist", &mqtt.MqttController{}, "*:GetChatHistoryList")
	beego.Router("/chat/user/online/info", &mqtt.MqttController{}, "*:GetOnlineUseInfo")
	beego.Router("/chat/user/online/count", &mqtt.MqttController{}, "*:GetOnlineUseCount")

	//question
	beego.Router("/chat/user/question/teacher", &mqtt.QuestionController{}, "*:GetQuestionTeacher")
	beego.Router("/chat/user/question/message", &mqtt.QuestionController{}, "*:GetQuestionToSend")
	beego.Router("/chat/user/question/historylist", &mqtt.QuestionController{}, "*:GetQuestionHistoryList")

	//公告
	beego.Router("/chat/user/notice", &mqtt.NoticeController{}, "*:GetPublishNotice")
	beego.Router("/chat/user/deleteNotice", &mqtt.NoticeController{}, "*:DeleteNotice")
	beego.Router("/chat/user/noticelist", &mqtt.NoticeController{}, "*:GetRoomNoticeList")

	// //策略
	beego.Router("/chat/user/strategy", &mqtt.StrategyController{}, "*:GetStrategyInfo")
	beego.Router("/chat/user/editstrategy", &mqtt.StrategyController{}, "*:GetEditStrategyInfo")
	beego.Router("/chat/user/operatestrategy", &mqtt.StrategyController{}, "*:OperateStrategy")
	beego.Router("/chat/user/strategyList", &mqtt.StrategyController{}, "*:GetStrategyList")
	beego.Router("/chat/user/strategyMap", &mqtt.StrategyController{}, "*:GetUnameMapInfo")
	beego.Router("/chat/user/upload", &mqtt.StrategyController{}, "*:Upload")

	//登录
	beego.Router("/chat/user/login", &mqtt.ManagerController{}, "*:GetUserLogin")
	//申请审核
	beego.Router("/chat/user/apply", &mqtt.ManagerController{}, "*:GetUserApply")
	// 当前在线人信息
	beego.Router("/chat/user/online", &mqtt.ManagerController{}, "*:GetUserOnline")
	//踢人
	beego.Router("/chat/user/KickOut", &mqtt.ManagerController{}, "*:GetKickOutInfo")
	//禁言
	beego.Router("/chat/user/ShutUp", &mqtt.ManagerController{}, "*:GetShutUpInfo")
	beego.Router("/chat/user/UnShutUp", &mqtt.ManagerController{}, "*:GetUnShutUpInfo")

	//专家团队
	beego.Router("/chat/user/teacherList", &mqtt.TeacherController{}, "*:GetTeacherList")
	beego.Router("/chat/user/allList", &mqtt.TeacherController{}, "*:GetAllTeahcerList")
	beego.Router("/chat/user/AddTeacher", &mqtt.TeacherController{}, "*:AddTeacher")
	beego.Router("/chat/user/operateTeacher", &mqtt.TeacherController{}, "*:OperateTeacher")

	//setting
	beego.Router("/chat/user/set/icon", &mqtt.SetController{}, "*:SetIcon")
	beego.Router("/chat/user/set/Nickname", &mqtt.SetController{}, "*:SetNickname")

	//仓位
	beego.Router("/chat/user/positionInfo", &mqtt.PositionController{}, "*:OperatePosition")
	beego.Router("/chat/user/positionList", &mqtt.PositionController{}, "*:GetPositionList")
	beego.Router("/chat/user/closeposition", &mqtt.PositionController{}, "*:GetClosePositionRecord")
	beego.Router("/chat/user/positionNear", &mqtt.PositionController{}, "*:GetPositionNearRecord")
	beego.Router("/chat/user/positionAllList", &mqtt.PositionController{}, "*:GetAllPositionList")

	//平仓操作
	beego.Router("/chat/user/operateclose", &mqtt.PositionController{}, "*:OperateClosePosition")

	//收藏
	beego.Router("/chat/user/Collect", &mqtt.CollectController{}, "*:GetCollectInfo")
	//free
	beego.Router("/wechatSocket", &haoindex.IndexController{}, "*:WechatFree")

}
