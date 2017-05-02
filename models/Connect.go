package models

import (
	"fmt"
	"os"
	"weserver/src/tools"

	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 链接数据库
func Connect() {
	dns, _ := getConfig()
	beego.Info("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		beego.Error("数据库连接失败")
	} else {
		beego.Info("数据库连接成功")
		// writeSiteConf()
	}

	for _, v := range os.Args {
		if v == "db" {
			inserData()
		}
	}
}

func getConfig() (string, string) {
	db_host := beego.AppConfig.String("host")
	db_port := beego.AppConfig.String("port")
	db_user := beego.AppConfig.String("username")
	db_pass := beego.AppConfig.String("password")
	db_name := beego.AppConfig.String("dbname")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	return dns, db_name
}

func inserData() {
	inserGroup()
	inserRole()
	insertTitle()
	inserUser()
	insertNodes()
	inserSys()
	inserRoominfo()

	// 删除User唯一索引
	o := orm.NewOrm()
	_, err := o.Raw("DROP INDEX role_id ON `user`").Exec()
	if err != nil {
		beego.Error("del role_index error", err)
	}
	_, err = o.Raw("DROP INDEX title_id ON `user`").Exec()
	if err != nil {
		beego.Error("del title_index error", err)
	}

	// 删除Regist唯一索引
	_, err = o.Raw("DROP INDEX role_id ON `regist`").Exec()
	if err != nil {
		beego.Error("del role_index error", err)
	}
	_, err = o.Raw("DROP INDEX title_id ON `regist`").Exec()
	if err != nil {
		beego.Error("del title_index error", err)
	}

}

func inserGroup() {
	group := new(Group)
	group.Id = 1
	group.Name = "admin"
	group.Title = "后台"
	group.Status = 1
	group.Sort = 1

	_, err := AddGroup(group)
	if err != nil {
		beego.Error("init group error")
	}
}

func inserRole() {
	role := [...]Role{
		{Id: 1, Title: "管理员", Name: "manager", Remark: "管理员", Status: 1, Weight: 1, Delay: 0, IsInsider: 1},
		{Id: 2, Title: "客服", Name: "customer", Remark: "客服", Status: 1, Weight: 1, Delay: 0, IsInsider: 1},
		{Id: 3, Title: "助理", Name: "zhuli", Remark: "助理", Status: 1, Weight: 1, Delay: 0, IsInsider: 1},
		{Id: 4, Title: "讲师", Name: "teacher", Remark: "讲师", Status: 1, Weight: 1, Delay: 0, IsInsider: 1},
		{Id: 5, Title: "普通", Name: "customer", Remark: "普通", Status: 1, Weight: 1, Delay: 0, IsInsider: 0},
		{Id: 6, Title: "游客", Name: "customer", Remark: "游客", Status: 1, Weight: 1, Delay: 0, IsInsider: 0},
	}
	for _, v := range role {
		AddRole(&v)
	}
	beego.Info("init role")
}

func insertTitle() {
	title := [...]Title{
		{Id: 1, Name: "管理员", Css: "#CC0000", Background: 1, Weight: 999, Remark: "管理员"},
		{Id: 2, Name: "至尊", Css: "#000000", Background: 0, Weight: 999, Remark: "至尊"},
		{Id: 3, Name: "铂金", Css: "#000000", Background: 0, Weight: 999, Remark: "铂金"},
		{Id: 4, Name: "黄金", Css: "#000000", Background: 0, Weight: 999, Remark: "黄金"},
		{Id: 5, Name: "白银", Css: "#000000", Background: 0, Weight: 999, Remark: "Remark"},
		{Id: 6, Name: "VIP", Css: "#000000", Background: 0, Weight: 999, Remark: "Remark"},
		{Id: 7, Name: "普通", Css: "#000000", Background: 0, Weight: 999, Remark: "Remark"},
		{Id: 8, Name: "分析师-胡老师", Css: "#CC0000", Background: 1, Weight: 999, Remark: "Remark"},
		{Id: 9, Name: "分析师-徐老师", Css: "#CC0000", Background: 1, Weight: 999, Remark: "Remark"},
		{Id: 10, Name: "分析师-王老师", Css: "#CC0000", Background: 1, Weight: 999, Remark: "Remark"},
	}
	for _, v := range title {
		AddTitle(&v)
	}
}

func inserUser() {
	user := new(User)
	user.Id = 1
	user.Username = "admin"
	user.Password = "4bc08e686673c541e4c70815763955b4"
	user.Nickname = "管理员"
	user.Status = 1
	user.Lastlogintime = time.Now()
	user.Createtime = time.Now()
	user.UserIcon = ""
	user.RegStatus = 1
	user.Role = &Role{Id: 1}
	user.Title = &Title{Id: 1}
	_, err := AddUser(user)
	if err != nil {
		beego.Error("add user error", err)
	}
}

func inserRoominfo() {
	roomGroupId := beego.AppConfig.String("roomGroupId")
	roomUrl := beego.AppConfig.String("roomUrl")
	roomAccess := beego.AppConfig.String("roomAccess")
	roomSecretKey := beego.AppConfig.String("roomSecretKey")
	roominfo := [...]RoomInfo{
		{Id: 1, Qos: 0, RoomTitle: "互动社区", RoomTeacher: "胡老师", RoomNum: "1352", GroupId: roomGroupId, Url: roomUrl, Port: 80, Tls: false, Access: roomAccess, SecretKey: roomSecretKey, RoomIcon: "i/allchannel/icon1.png", RoomIntro: "时刻为您提供最新的红木行情分析，为广大客户提供交流的平台", RoomBanner: "i/allchannel/banner1.png", Title: "互动社区", MidPage: 1},
		{Id: 2, Qos: 0, RoomTitle: "会员专区", RoomTeacher: "徐老师", RoomNum: "1982", GroupId: roomGroupId, Url: roomUrl, Port: 80, Tls: false, Access: roomAccess, SecretKey: roomSecretKey, RoomIcon: "i/allchannel/icon1.png", RoomIntro: "专业的分析给您更准确的投资建议", RoomBanner: "i/allchannel/banner1.png", Title: "会员专区", MidPage: 1},
		{Id: 3, Qos: 0, RoomTitle: "贵宾专区", RoomTeacher: "王老师", RoomNum: "650", GroupId: roomGroupId, Url: roomUrl, Port: 80, Tls: false, Access: roomAccess, SecretKey: roomSecretKey, RoomIcon: "i/allchannel/icon3.png", RoomIntro: "尊享专人一对一的贴心服务，给您专属的投资指导", RoomBanner: "i/allchannel/banner1.png", Title: "贵宾专区", MidPage: 1},
		{Id: 4, Qos: 0, RoomTitle: "铂金VIP", RoomTeacher: "雷老师", RoomNum: "960", GroupId: roomGroupId, Url: roomUrl, Port: 80, Tls: false, Access: roomAccess, SecretKey: roomSecretKey, RoomIcon: "i/allchannel/icon4.png", RoomIntro: "网罗最新最全的红木行情，由数位具有十数年经验的高级分析师为您提供专属服务", RoomBanner: "i/allchannel/banner1.png", Title: "铂金VIP", MidPage: 1},
	}

	for _, v := range roominfo {
		v.RoomId = beego.AppConfig.String("mqServerTopic") + "/" + getRoomId()
		AddRoom(&v)
	}
}

func insertNodes() {
	countNode := 60
	count, _ := GetNodeCount()
	if int(count) >= countNode {
		fmt.Println("node haved")
	} else {
		fmt.Println("insert node start")
		nodes := [...]Node{
			{Id: 1, Title: "用户管理", Name: "user", Level: 1, Pid: 0, Remark: "用户管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user", Hide: 1, Ico: "am-icon-user"},
			{Id: 2, Title: "用户列表", Name: "user/index", Level: 2, Pid: 1, Remark: "用户管理/用户列表", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/index", Hide: 1, Ico: ""},
			{Id: 3, Title: "用户设置列表", Name: "user/usersetlist", Level: 2, Pid: 1, Remark: "用户管理/用户设置列表", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/usersetlist", Hide: 1, Ico: ""},
			{Id: 4, Title: "更新用户", Name: "user/update", Level: 3, Pid: 2, Remark: "用户管理/增加用户", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/updateuser", Hide: 1, Ico: ""},
			{Id: 5, Title: "删除用户", Name: "user/deluser", Level: 3, Pid: 2, Remark: "用户管理/删除用户", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/deluser", Hide: 1, Ico: ""},
			{Id: 6, Title: "用户赋予角色", Name: "user/usertorole", Level: 3, Pid: 2, Remark: "用户管理/用户赋予角色", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/usertorole", Hide: 1, Ico: ""},
			{Id: 7, Title: "用户赋予头衔", Name: "user/setusertitle", Level: 3, Pid: 2, Remark: "用户管理/用户赋予头衔", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/setusertitle", Hide: 1, Ico: ""},
			{Id: 8, Title: "房间用户审核", Name: "user/regstatus", Level: 3, Pid: 2, Remark: "用户管理/房间用户审核", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/regstatus", Hide: 1, Ico: ""},
			{Id: 9, Title: "踢出房间", Name: "user/kictuser", Level: 3, Pid: 2, Remark: "用户管理/踢出房间", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/kictuser", Hide: 1, Ico: ""},
			{Id: 10, Title: "批量删除用户", Name: "user/preparedel", Level: 3, Pid: 2, Remark: "用户管理/踢出房间", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/preparedel", Hide: 1, Ico: ""},
			{Id: 11, Title: "增加用户", Name: "user/add", Level: 2, Pid: 1, Remark: "用户管理/增加用户", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/adduser", Hide: 1, Ico: ""},
			{Id: 12, Title: "初始化用户", Name: "user/setusername", Level: 3, Pid: 11, Remark: "用户管理/初始化用户", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/setusername", Hide: 1, Ico: ""},
			{Id: 13, Title: "解除禁言", Name: "user/UnShutUp", Level: 3, Pid: 2, Remark: "用户管理/解除禁言", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/UnShutUp", Hide: 1, Ico: ""},
			{Id: 14, Title: "头衔管理", Name: "title", Level: 1, Pid: 0, Remark: "头衔管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/title", Hide: 1, Ico: "am-icon-header"},
			{Id: 15, Title: "头衔列表", Name: "title/index", Level: 2, Pid: 14, Remark: "头衔列表", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/title/index", Hide: 1, Ico: ""},
			{Id: 16, Title: "新增头衔", Name: "title/addtitle", Level: 2, Pid: 14, Remark: "新增头衔", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/title/addtitle", Hide: 1, Ico: ""},
			{Id: 17, Title: "删除头衔", Name: "title/deltitle", Level: 3, Pid: 15, Remark: "删除头衔", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/title/deltitle", Hide: 1, Ico: ""},
			{Id: 18, Title: "更新头衔", Name: "title/updatetitle", Level: 3, Pid: 15, Remark: "更新头衔", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/title/updatetitle", Hide: 1, Ico: ""},
			{Id: 19, Title: "获取头衔", Name: "title/getalltitle", Level: 3, Pid: 15, Remark: "获取头衔", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/title/getalltitle", Hide: 1, Ico: ""},
			{Id: 20, Title: "头衔图片上传", Name: "title/upload", Level: 3, Pid: 15, Remark: "头衔图片上传", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/title/upload", Hide: 1, Ico: ""},
			{Id: 21, Title: "角色管理", Name: "role", Level: 1, Pid: 0, Remark: "角色管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role", Hide: 1, Ico: "am-icon-paper-plane"},
			{Id: 22, Title: "角色列表", Name: "role/index", Level: 2, Pid: 21, Remark: "角色列表", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role/index", Hide: 1, Ico: ""},
			{Id: 23, Title: "增加角色", Name: "role/addrole", Level: 3, Pid: 22, Remark: "增加角色", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role/addrole", Hide: 1, Ico: ""},
			{Id: 24, Title: "更新角色", Name: "role/updaterole", Level: 3, Pid: 22, Remark: "更新角色", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role/updaterole", Hide: 1, Ico: ""},
			{Id: 25, Title: "删除角色", Name: "role/delrole", Level: 3, Pid: 22, Remark: "删除角色", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role/delrole", Hide: 1, Ico: ""},
			{Id: 26, Title: "角色赋权", Name: "role/addaccess", Level: 3, Pid: 22, Remark: "角色赋权", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role/addaccess", Hide: 1, Ico: ""},
			{Id: 27, Title: "角色赋权获取", Name: "role/accesstonode", Level: 3, Pid: 22, Remark: "角色赋权获取", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role/accesstonode", Hide: 1, Ico: ""},
			{Id: 28, Title: "获取所有角色", Name: "role/getallrole", Level: 3, Pid: 22, Remark: "获取所有角色", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role/getallrole", Hide: 1, Ico: ""},
			{Id: 29, Title: "角色图片上传", Name: "role/upload", Level: 3, Pid: 22, Remark: "角色图片上传", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/role/upload", Hide: 1, Ico: ""},
			{Id: 30, Title: "数据管理", Name: "data", Level: 1, Pid: 0, Remark: "数据管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data", Hide: 1, Ico: "am-icon-database"},
			{Id: 31, Title: "公告管理", Name: "data/qs_broad", Level: 2, Pid: 30, Remark: "公告列表", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/qs_broad", Hide: 1, Ico: ""},
			{Id: 32, Title: "发送公告", Name: "data/sendbroad", Level: 3, Pid: 31, Remark: "发送公告", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/sendbroad", Hide: 1, Ico: ""},
			{Id: 33, Title: "编辑公告", Name: "data/notice_edit", Level: 3, Pid: 31, Remark: "编辑公告", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/notice_edit", Hide: 1, Ico: ""},
			{Id: 34, Title: "删除公告", Name: "data/notice_del", Level: 3, Pid: 31, Remark: "删除公告", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/notice_del", Hide: 1, Ico: ""},
			{Id: 35, Title: "房间管理", Name: "data/room_index", Level: 2, Pid: 30, Remark: "房间管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/room_index", Hide: 1, Ico: ""},
			{Id: 36, Title: "增加房间", Name: "data/room_add", Level: 3, Pid: 35, Remark: "增加房间", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/room_add", Hide: 1, Ico: ""},
			{Id: 37, Title: "修改房间", Name: "data/room_edit", Level: 3, Pid: 35, Remark: "修改房间", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/room_edit", Hide: 1, Ico: ""},
			{Id: 38, Title: "删除房间", Name: "data/room_del", Level: 3, Pid: 35, Remark: "删除房间", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/room_del", Hide: 1, Ico: ""},
			{Id: 39, Title: "策略管理", Name: "data/strategy_index", Level: 2, Pid: 30, Remark: "策略管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/strategy_index", Hide: 1, Ico: ""},
			{Id: 40, Title: "增加策略", Name: "data/strategy_add", Level: 3, Pid: 39, Remark: "策略管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/strategy_add", Hide: 1, Ico: ""},
			{Id: 41, Title: "删除策略", Name: "data/strategy_del", Level: 3, Pid: 39, Remark: "策略管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/strategy_del", Hide: 1, Ico: ""},
			{Id: 42, Title: "讲师管理", Name: "data/teacher_index", Level: 2, Pid: 30, Remark: "讲师管理", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/teacher_index", Hide: 1, Ico: ""},
			{Id: 43, Title: "增加讲师", Name: "data/teacher_add", Level: 3, Pid: 42, Remark: "增加讲师", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/teacher_add", Hide: 1, Ico: ""},
			{Id: 44, Title: "修改讲师", Name: "data/teacher_edit", Level: 3, Pid: 42, Remark: "修改讲师", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/teacher_edit", Hide: 1, Ico: ""},
			{Id: 45, Title: "删除讲师", Name: "data/teacher_del", Level: 3, Pid: 42, Remark: "删除讲师", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/teacher_del", Hide: 1, Ico: ""},
			{Id: 46, Title: "查询房间讲师", Name: "data/teacher_room", Level: 3, Pid: 42, Remark: "查询房间讲师", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/teacher_room", Hide: 1, Ico: ""},
			{Id: 47, Title: "操作建议", Name: "data/suggest_index", Level: 2, Pid: 30, Remark: "操作建议", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/suggest_index", Hide: 1, Ico: ""},
			{Id: 48, Title: "增加建仓", Name: "data/suggest_add", Level: 3, Pid: 47, Remark: "增加建仓", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/suggest_add", Hide: 1, Ico: ""},
			{Id: 49, Title: "编辑建仓", Name: "data/suggest_edit", Level: 3, Pid: 47, Remark: "编辑建仓", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/suggest_edit", Hide: 1, Ico: ""},
			{Id: 50, Title: "删除建仓", Name: "data/suggest_del", Level: 3, Pid: 47, Remark: "删除建仓", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/suggest_del", Hide: 1, Ico: ""},
			{Id: 51, Title: "增加平仓", Name: "data/suggest_addclose", Level: 3, Pid: 47, Remark: "增加平仓", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/suggest_addclose", Hide: 1, Ico: ""},
			{Id: 52, Title: "获取平仓", Name: "data/suggest_getclose", Level: 3, Pid: 47, Remark: "获取平仓", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/suggest_getclose", Hide: 1, Ico: ""},
			{Id: 53, Title: "编辑平仓", Name: "data/suggest_editclose", Level: 3, Pid: 47, Remark: "编辑平仓", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/suggest_editclose", Hide: 1, Ico: ""},
			{Id: 54, Title: "删除平仓", Name: "data/suggest_delclose", Level: 3, Pid: 47, Remark: "删除平仓", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/suggest_delclose", Hide: 1, Ico: ""},
			{Id: 55, Title: "系统设置", Name: "sysconfig", Level: 1, Pid: 0, Remark: "系统设置", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/sysconfig", Hide: 1, Ico: "am-icon-cog"},
			{Id: 56, Title: "全局设置", Name: "sysconfig/index", Level: 2, Pid: 55, Remark: "全局设置", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/sysconfig/index", Hide: 1, Ico: ""},
			{Id: 57, Title: "图片上传", Name: "data/upload", Level: 3, Pid: 56, Remark: "图片上传", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/upload", Hide: 1, Ico: ""},
			{Id: 58, Title: "用户状态修改", Name: "user/userstatus", Level: 3, Pid: 3, Remark: "用户状态修改", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/user/userstatus", Hide: 1, Ico: ""},
			{Id: 59, Title: "聊天记录", Name: "data/chatrecord", Level: 2, Pid: 0, Remark: "聊天记录", Status: 2, Group: &Group{Id: 1}, Sort: 100, Url: "weserver/data/chatrecord", Hide: 1, Ico: ""},
		}
		for _, v := range nodes {
			AddNode(&v)
		}
	}
}

func inserSys() {
	sys := new(SysConfig)
	sys.Id = 1
	sys.Systemname = "互动社区"
	sys.ChatInterval = 0
	sys.Registerrole = 6
	sys.Registertitle = 7
	sys.HistoryMsg = 0
	sys.HistoryCount = 50
	sys.NoticeCount = 10
	sys.StrategyCount = 0
	sys.TeacherCount = 0
	sys.PositionCount = 0
	sys.WelcomeMsg = "欢迎来到互动社区"
	sys.Verify = 0
	sys.LoginSys = 1
	sys.AuditMsg = 0
	sys.VirtualUser = 10
	_, err := AddSysConfig(sys)
	if err != nil {
		beego.Error("init db", err)
	}
}

func getRoomId() string {
	random := tools.RandomAlphanumeric(6)
	if IsRoomInfo(random) {
		getRoomId()
	}
	return random
}
