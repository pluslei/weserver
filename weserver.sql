/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50712
Source Host           : localhost:3306
Source Database       : weserver

Target Server Type    : MYSQL
Target Server Version : 50712
File Encoding         : 65001

Date: 2017-04-08 16:10:56
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for black_list
-- ----------------------------
DROP TABLE IF EXISTS `black_list`;
CREATE TABLE `black_list` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `coderoom` int(11) NOT NULL DEFAULT '0',
  `uname` varchar(128) NOT NULL DEFAULT '',
  `objname` varchar(128) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `ipaddress` varchar(128) NOT NULL DEFAULT '',
  `procities` varchar(128) NOT NULL DEFAULT '',
  `datatime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of black_list
-- ----------------------------

-- ----------------------------
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS `chat_record`;
CREATE TABLE `chat_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room` varchar(255) NOT NULL DEFAULT '',
  `uname` varchar(255) NOT NULL DEFAULT '',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `user_icon` varchar(255) NOT NULL DEFAULT '',
  `role_name` varchar(255) NOT NULL DEFAULT '',
  `role_title` varchar(255) NOT NULL DEFAULT '',
  `sendtype` varchar(255) NOT NULL DEFAULT '',
  `role_title_css` varchar(255) NOT NULL DEFAULT '',
  `role_title_back` int(11) NOT NULL DEFAULT '0',
  `insider` int(11) NOT NULL DEFAULT '1',
  `is_login` int(11) NOT NULL DEFAULT '0',
  `content` longtext NOT NULL,
  `datatime` datetime NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `uuid` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chat_record
-- ----------------------------
INSERT INTO `chat_record` VALUES ('1', 'Topic_ChatMessage/sub1', 'ooex5xK49v54rbsWVA2KUWT0TNb8', 'Sir. Lei', 'http://wx.qlogo.cn/mmopen/KetjXWSVppsZ0icialcRKRXwLaCpCcoa664FHnSrnL5w6u9x1qyb8FfD35MiavwjibBQKiaQtdzyX9nyMibIicAx5htkLyIGEKJTAhB/0', 'guest', '普通', 'TXT', '1472629522.jpg', '0', '1', '1', 'hello&nbsp;<span style=\"font-size: 1.6rem;\">world</span>', '2017-03-08 12:00:22', '1', '90161783-93c1-48dc-bad7-598ae02871a3');
INSERT INTO `chat_record` VALUES ('2', 'Topic_ChatMessage/sub2', 'ooex5xFQ6NnAs7ZmJ1DyuZlrbRm8', '旭', 'http://wx.qlogo.cn/mmopen/KetjXWSVppsZ0icialcRKRX2DCfHcppb9GohWY00UQiaLibcxdm5auDwz8zHQXhEkgrmHqUuhM0eYic3sQ27ws1ohepiaDOkzjdXX2/0', 'guest', '普通', 'TXT', '1472629522.jpg', '0', '1', '1', '哈哈', '2017-04-01 13:01:35', '1', 'ae259730-da72-40ea-b546-65eb9d447477');

-- ----------------------------
-- Table structure for collect
-- ----------------------------
DROP TABLE IF EXISTS `collect`;
CREATE TABLE `collect` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uname` varchar(255) NOT NULL DEFAULT '',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `room_icon` varchar(255) NOT NULL DEFAULT '',
  `room_title` varchar(255) NOT NULL DEFAULT '',
  `room_teacher` varchar(255) NOT NULL DEFAULT '',
  `is_collect` tinyint(1) NOT NULL DEFAULT '0',
  `is_attention` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of collect
-- ----------------------------

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '',
  `title` varchar(128) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '1',
  `sort` int(11) NOT NULL DEFAULT '50',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of group
-- ----------------------------
INSERT INTO `group` VALUES ('1', 'admin', '后台', '1', '1');
INSERT INTO `group` VALUES ('2', 'index', '前台', '1', '50');

-- ----------------------------
-- Table structure for kick_out
-- ----------------------------
DROP TABLE IF EXISTS `kick_out`;
CREATE TABLE `kick_out` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `coderoom` varchar(255) NOT NULL DEFAULT '',
  `operuid` varchar(255) NOT NULL DEFAULT '',
  `opername` varchar(128) NOT NULL DEFAULT '',
  `objuid` varchar(255) NOT NULL DEFAULT '',
  `objname` varchar(128) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `opertime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kick_out
-- ----------------------------

-- ----------------------------
-- Table structure for node
-- ----------------------------
DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `level` int(11) NOT NULL DEFAULT '1',
  `pid` bigint(20) NOT NULL DEFAULT '0',
  `remark` varchar(200) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '2',
  `group_id` bigint(20) NOT NULL,
  `sort` int(10) DEFAULT '50',
  `url` varchar(100) NOT NULL DEFAULT '',
  `hide` int(11) NOT NULL DEFAULT '1',
  `ico` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of node
-- ----------------------------
INSERT INTO `node` VALUES ('1', '用户管理', 'user', '1', '0', '用户管理', '2', '1', '100', 'weserver/user', '1', 'am-icon-user');
INSERT INTO `node` VALUES ('2', '用户列表', 'user/index', '2', '1', '用户管理/用户列表', '2', '1', '101', 'weserver/user/index', '1', '');
INSERT INTO `node` VALUES ('3', '增加用户', 'user/add', '2', '1', '用户管理/增加用户', '2', '1', '103', 'weserver/user/adduser', '1', '');
INSERT INTO `node` VALUES ('4', '更新用户', 'user/update', '3', '2', '用户管理/更新用户', '2', '1', '50', 'weserver/user/updateuser', '1', '');
INSERT INTO `node` VALUES ('5', '删除用户', 'user/del', '3', '2', '用户管理/删除用户', '2', '1', '50', 'weserver/user/deluser', '1', '');
INSERT INTO `node` VALUES ('6', '角色管理', 'role', '1', '0', '角色管理', '2', '1', '200', 'weserver/role', '1', 'am-icon-paper-plane');
INSERT INTO `node` VALUES ('7', '角色列表', 'role/index', '2', '6', '角色管理/角色列表', '2', '1', '50', 'weserver/role/index', '1', '');
INSERT INTO `node` VALUES ('8', '增加角色', 'role/add', '2', '6', '角色管理/增加角色', '2', '1', '50', 'weserver/role/addrole', '1', '');
INSERT INTO `node` VALUES ('9', '编辑角色', 'role/updaterole', '3', '7', '角色管理/编辑角色', '2', '1', '50', 'weserver/role/updaterole', '1', '');
INSERT INTO `node` VALUES ('10', '删除角色', 'role/del', '3', '7', '角色管理/删除角色', '2', '1', '50', 'weserver/role/delrole', '1', '');
INSERT INTO `node` VALUES ('11', '角色赋予权限', 'role/addaccess', '3', '6', '角色管理/角色赋予权限', '2', '1', '50', 'weserver/role/addaccess', '1', '');
INSERT INTO `node` VALUES ('12', '获取角色节点', 'role/accesstonode', '3', '6', '角色管理/获取角色节点', '2', '1', '50', 'weserver/role/accesstonode', '1', '');
INSERT INTO `node` VALUES ('13', '获取角色', 'role/getallrole', '3', '6', '角色管理/获取角色', '2', '1', '50', 'weserver/role/getallrole', '1', '');
INSERT INTO `node` VALUES ('14', '头衔管理', 'title', '1', '0', '头衔管理', '2', '1', '300', 'weserver/title', '1', 'am-icon-header');
INSERT INTO `node` VALUES ('15', '头衔列表', 'title/index', '2', '14', '头衔管理/头衔列表', '2', '1', '50', 'weserver/title/index', '1', '');
INSERT INTO `node` VALUES ('16', '新增头衔', 'title/addtitle', '2', '14', '头衔管理/新增头衔', '2', '1', '50', 'weserver/title/addtitle', '1', '');
INSERT INTO `node` VALUES ('17', '更新头衔', 'title/updatetitle', '3', '15', '头衔管理/更新头衔', '2', '1', '50', 'weserver/title/updatetitle', '1', '');
INSERT INTO `node` VALUES ('18', '删除头衔', 'title/deltitle', '3', '15', '头衔管理/删除头衔', '2', '1', '50', 'weserver/title/deltitle', '1', '');
INSERT INTO `node` VALUES ('19', '节点管理', 'node', '1', '0', '节点管理', '2', '1', '400', 'weserver/node', '0', '');
INSERT INTO `node` VALUES ('20', '节点列表', 'node/index', '2', '19', '节点管理/节点列表', '2', '1', '50', 'weserver/node/index', '1', '');
INSERT INTO `node` VALUES ('21', '增加节点', 'node/addnode', '2', '19', '节点管理/增加节点', '2', '1', '50', 'weserver/node/addnode', '1', '');
INSERT INTO `node` VALUES ('22', '更新节点', 'node/updatenode', '3', '19', '节点管理/更新节点', '2', '1', '50', 'weserver/node/updatenode', '1', '');
INSERT INTO `node` VALUES ('23', '删除节点', 'node/delnode', '3', '19', '节点管理/删除节点', '2', '1', '50', 'weserver/node/delnode', '1', '');
INSERT INTO `node` VALUES ('24', '获取节点树', 'node/getnodetree', '3', '19', '节点管理/获取节点树', '2', '1', '50', 'weserver/node/getnodetree', '1', '');
INSERT INTO `node` VALUES ('25', '组别管理', 'group', '1', '0', '组别管理', '2', '1', '500', 'weserver/group', '0', '');
INSERT INTO `node` VALUES ('26', '组别列表', 'group/index', '2', '25', '组别管理/组别列表', '2', '1', '50', 'weserver/group/index', '1', '');
INSERT INTO `node` VALUES ('27', '新增组别', 'group/addgroup', '3', '25', '组别管理/新增组别', '2', '1', '50', 'weserver/group/addgroup', '1', '');
INSERT INTO `node` VALUES ('28', '更新组别', 'group/updategroup', '3', '25', '组别管理/更新组别', '2', '1', '50', 'weserver/group/updategroup', '1', '');
INSERT INTO `node` VALUES ('29', '删除组别', 'group/delgroup', '3', '25', '组别管理/删除组别', '2', '1', '50', 'weserver/group/delgroup', '1', '');
INSERT INTO `node` VALUES ('30', '系统设置', 'sysconfig', '1', '0', '系统设置', '2', '1', '800', 'weserver/sysconfig', '1', 'am-icon-cog');
INSERT INTO `node` VALUES ('31', '全局设置', 'sysconfig/index', '2', '30', '系统设置/全局设置', '2', '1', '50', 'weserver/sysconfig/index', '1', '');
INSERT INTO `node` VALUES ('32', '全局设置更新', 'sysonfig/updateconfig', '3', '31', '系统设置/全局设置更新', '2', '1', '50', 'weserver/sysconfig/updateconfig', '1', '');
INSERT INTO `node` VALUES ('33', '主题设置', 'theme/index', '2', '30', '系统设置/主题设置', '2', '1', '50', 'weserver/sysconfig/theme_index', '0', '');
INSERT INTO `node` VALUES ('34', '数据管理', 'data', '1', '0', '数据管理', '2', '1', '600', 'weserver/data', '1', 'am-icon-database');
INSERT INTO `node` VALUES ('35', '房间管理', 'room/index', '2', '34', '数据管理/房间管理', '2', '1', '50', 'weserver/data/room_index', '0', '');
INSERT INTO `node` VALUES ('36', '前端菜单', 'index', '1', '0', '前端管理', '2', '2', '50', 'index', '1', '');
INSERT INTO `node` VALUES ('37', '对TA说', 'saidto', '2', '36', '前端菜单/对TA说', '2', '2', '50', 'saidto', '1', '');
INSERT INTO `node` VALUES ('38', '对TA私聊', 'saidtosecret', '2', '36', '前端菜单/对TA私聊', '2', '2', '50', 'saidtosecert', '1', '');
INSERT INTO `node` VALUES ('39', '踢出1小时', 'out1hour', '2', '36', '前端菜单/踢出1小时', '2', '2', '50', 'out1hour', '1', '');
INSERT INTO `node` VALUES ('40', '加入黑名单', 'addblacklist', '2', '36', '前端菜单/加入黑名单', '2', '2', '50', 'addblacklist', '1', '');
INSERT INTO `node` VALUES ('41', '禁言五分钟', 'disablemsg', '2', '36', '前端菜单/禁言五分钟', '2', '2', '50', 'disablemsg', '1', '');
INSERT INTO `node` VALUES ('42', '恢复发言', 'enablemsg', '2', '36', '前端菜单/恢复发言', '2', '2', '50', 'enablemsg', '1', '');
INSERT INTO `node` VALUES ('43', '增加主题', 'theme/addtheme', '2', '33', '系统设置/增加主题', '2', '1', '50', 'weserver/sysconfig/theme_addtheme', '0', '');
INSERT INTO `node` VALUES ('44', '更新主题', 'theme/updatetheme', '3', '33', '系统设置/更新主题', '2', '1', '50', 'weserver/sysconfig/theme_updatetheme', '0', '');
INSERT INTO `node` VALUES ('45', '更新房间', 'room/updateroom', '3', '35', '房间管理/更新房间', '2', '1', '50', 'weserver/sysconfig/room_updateroom', '1', '');
INSERT INTO `node` VALUES ('46', '刷新房间', 'room/refresh', '3', '35', '房间管理/刷新房间', '2', '1', '50', 'weserver/data/room_refresh', '0', '');
INSERT INTO `node` VALUES ('47', '在线用户', 'user/onlineuser', '2', '1', '用户管理/在线用户', '2', '1', '102', 'weserver/user/onlineuser', '1', '');
INSERT INTO `node` VALUES ('48', '首页管理', 'home', '1', '0', '首页管理', '2', '1', '700', 'weserver/home', '0', 'am-icon-home');
INSERT INTO `node` VALUES ('49', '关于我们', 'home/aboutme', '2', '48', '首页管理/关于我们', '2', '1', '50', 'weserver/home/aboutme', '1', '');
INSERT INTO `node` VALUES ('50', '联系我们', 'home/contact', '2', '48', '首页管理/联系我们', '2', '1', '50', 'weserver/home/contact', '1', '');
INSERT INTO `node` VALUES ('51', '讲师简介', 'teacher/index', '2', '34', '数据管理/讲师简介', '2', '1', '50', 'weserver/data/teacher_index', '0', '');
INSERT INTO `node` VALUES ('52', '增加讲师', 'teacher/addteacher', '3', '51', '数据管理/增加讲师', '2', '1', '50', 'weserver/data/teacher_addteacher', '1', '');
INSERT INTO `node` VALUES ('53', '更新讲师', 'teacher/updateteacher', '3', '51', '数据管理/更新讲师', '2', '1', '50', 'weserver/data/teacher_updateteacher', '1', '');
INSERT INTO `node` VALUES ('54', '删除讲师', 'teacher_delteacher', '3', '51', '数据管理/删除讲师', '2', '1', '50', 'weserver/data/teacher_delteacher', '1', '');
INSERT INTO `node` VALUES ('55', '课程管理', 'home/course_index', '2', '48', '首页管理/课程管理', '2', '1', '50', 'weserver/home/course_index', '1', '');
INSERT INTO `node` VALUES ('56', '增加课程', 'home/course_addcourse', '3', '55', '首页管理/增加课程', '2', '1', '50', 'weserver/home/course_addcourse', '1', '');
INSERT INTO `node` VALUES ('57', '更新课程', 'home/course_updatecourse', '3', '55', '首页管理/更新课程', '2', '1', '50', 'weserver/home/course_updatecourse', '1', '');
INSERT INTO `node` VALUES ('58', '删除课程', 'home/course_delcourse', '3', '55', '首页管理/删除课程', '2', '1', '50', 'weserver/home/course_delcourse', '1', '');
INSERT INTO `node` VALUES ('59', '获取课程信息', 'home/course_coursejson', '3', '55', '首页管理/获取课程json', '2', '1', '50', 'weserver/home/course_coursejson', '1', '');
INSERT INTO `node` VALUES ('60', '客服管理', 'home/custservice_index', '2', '48', '首页管理/客服管理', '2', '1', '50', 'weserver/home/custservice_index', '1', '');
INSERT INTO `node` VALUES ('61', '增加客服', 'home/custservice_addcust', '3', '60', '首页管理/增加客服', '2', '1', '50', 'weserver/home/custservice_addcust', '1', '');
INSERT INTO `node` VALUES ('62', '更新客服', 'home/custservice_updatecust', '3', '60', '首页管理/更新客服', '2', '1', '50', 'weserver/home/custservice_updatecust', '1', '');
INSERT INTO `node` VALUES ('63', '删除客服', 'home/custservice_delcust', '3', '60', '首页管理/删除客服', '2', '1', '50', 'weserver/home/custservice_delcust', '1', '');
INSERT INTO `node` VALUES ('64', '首页幻灯片', 'home/teachbanner_index', '2', '48', '首页管理/首页幻灯片', '2', '1', '50', 'weserver/home/teachbanner_index', '1', '');
INSERT INTO `node` VALUES ('65', '增加幻灯片', 'home/teachbanner_addbanner', '3', '64', '首页管理/增加幻灯片', '2', '1', '50', 'weserver/home/teachbanner_addbanner', '1', '');
INSERT INTO `node` VALUES ('66', '更新幻灯片', 'home/teachbanner_updatebanner', '3', '64', '首页管理/更新幻灯片', '2', '1', '50', 'weserver/home/teachbanner_updatebanner', '1', '');
INSERT INTO `node` VALUES ('67', '删除幻灯片', 'home/teachbanner_delbanner', '3', '64', '首页管理/删除幻灯片', '2', '1', '50', 'weserver/home/teachbanner_delbanner', '1', '');
INSERT INTO `node` VALUES ('68', '幻灯片上传', 'home/teachbanner_upload', '3', '64', '首页管理/上传幻灯片', '2', '1', '50', 'weserver/home/teachbanner_upload', '1', '');
INSERT INTO `node` VALUES ('69', '问题解答', 'data/qs_index', '2', '34', '数据管理/问题解答', '2', '1', '50', 'weserver/data/qs_index', '0', '');
INSERT INTO `node` VALUES ('70', '问题增加', 'data/qs_addqs', '3', '69', '数据管理/问题增加', '2', '1', '50', 'weserver/data/qs_addqs', '1', '');
INSERT INTO `node` VALUES ('71', '更新问题', 'data/qs_updateqs', '3', '69', '数据管理/更新问题', '2', '1', '50', 'weserver/data/qs_updateqs', '1', '');
INSERT INTO `node` VALUES ('72', '删除问题', 'data/qs_delqs', '3', '69', '数据管理/删除问题', '2', '1', '50', 'weserver/data/qs_delqs', '1', '');
INSERT INTO `node` VALUES ('73', '发布公告', 'data/qs_broad', '2', '34', '数据管理/qs_broad', '2', '1', '50', 'weserver/data/qs_broad', '1', '');
INSERT INTO `node` VALUES ('74', '前端功能', 'tool', '1', '0', '前端功能', '2', '2', '50', 'tool', '1', '');
INSERT INTO `node` VALUES ('75', '是否上传聊天图片', 'uploadchatimage', '2', '74', '前端功能/上传聊天图片', '2', '2', '50', 'tool/uploadmsgimage', '1', '');
INSERT INTO `node` VALUES ('76', '是否私聊', 'provicechat', '2', '74', '前端功能/是否私聊', '2', '2', '50', 'tool/provicechat', '1', '');
INSERT INTO `node` VALUES ('77', '是否发送广播', 'sendbroadcast', '2', '74', '前端功能/是否发送广播', '2', '2', '50', 'tool/sendbrodcast', '1', '');
INSERT INTO `node` VALUES ('78', '用户审核', 'user/verifyuser', '2', '1', '用户管理/用户审核', '2', '1', '50', 'weserver/user/verifyuser', '0', '');
INSERT INTO `node` VALUES ('79', '机器人', 'data/robot_speak', '2', '34', '数据管理/机器人', '2', '1', '50', 'weserver/data/robot_speak', '0', '');
INSERT INTO `node` VALUES ('80', '消息库管理', 'data/message_index', '2', '34', '数据管理/消息库管理', '2', '1', '50', 'weserver/data/message_index', '0', '');
INSERT INTO `node` VALUES ('81', '删除消息库', 'data/message_delete', '3', '79', '数据管理/删除消息库', '2', '1', '50', 'weserver/data/message_delete', '1', '');
INSERT INTO `node` VALUES ('82', '修改消息库', 'data/message_edit', '3', '79', '数据管理/修改消息库', '2', '1', '50', 'weserver/data/message_edit', '1', '');
INSERT INTO `node` VALUES ('83', '增加消息库', 'data/message_add', '3', '79', '数据管理/增加消息库', '2', '1', '50', 'weserver/data/messgae_add', '1', '');
INSERT INTO `node` VALUES ('84', '修改消息库分类', 'data/messagetype_edit', '3', '79', '数据管理/修改消息库分类', '2', '1', '50', 'weserver/data/messagetype_edit', '1', '');
INSERT INTO `node` VALUES ('85', '增加消息库分类', 'data/messagetype_add', '3', '79', '数据管理/增加消息库分类', '2', '1', '50', 'weserver/data/messagetype_add', '1', '');
INSERT INTO `node` VALUES ('86', '删除消息库分类', 'data/messagetype_delete', '3', '79', '数据管理/删除消息库分类', '2', '1', '50', 'weserver/data/messagetype_delete', '1', '');
INSERT INTO `node` VALUES ('98', '飞屏', 'flyscreen', '2', '36', '前端功能/飞屏', '2', '2', '50', 'flyscreen', '1', '');
INSERT INTO `node` VALUES ('99', '聊天记录', 'data/chatrecord', '2', '34', '数据管理/聊天记录', '2', '1', '50', 'weserver/data/chatrecord', '1', '');

-- ----------------------------
-- Table structure for node_roles
-- ----------------------------
DROP TABLE IF EXISTS `node_roles`;
CREATE TABLE `node_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `node_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of node_roles
-- ----------------------------
INSERT INTO `node_roles` VALUES ('3', '32', '4');
INSERT INTO `node_roles` VALUES ('4', '33', '4');
INSERT INTO `node_roles` VALUES ('5', '34', '4');
INSERT INTO `node_roles` VALUES ('6', '35', '4');
INSERT INTO `node_roles` VALUES ('17', '74', '5');
INSERT INTO `node_roles` VALUES ('18', '75', '5');
INSERT INTO `node_roles` VALUES ('19', '36', '4');
INSERT INTO `node_roles` VALUES ('20', '37', '4');
INSERT INTO `node_roles` VALUES ('21', '38', '4');
INSERT INTO `node_roles` VALUES ('22', '39', '4');
INSERT INTO `node_roles` VALUES ('23', '40', '4');
INSERT INTO `node_roles` VALUES ('24', '41', '4');
INSERT INTO `node_roles` VALUES ('25', '42', '4');
INSERT INTO `node_roles` VALUES ('26', '74', '4');
INSERT INTO `node_roles` VALUES ('27', '75', '4');
INSERT INTO `node_roles` VALUES ('28', '76', '4');
INSERT INTO `node_roles` VALUES ('29', '77', '4');
INSERT INTO `node_roles` VALUES ('30', '36', '3');
INSERT INTO `node_roles` VALUES ('31', '37', '3');
INSERT INTO `node_roles` VALUES ('32', '38', '3');
INSERT INTO `node_roles` VALUES ('33', '39', '3');
INSERT INTO `node_roles` VALUES ('34', '40', '3');
INSERT INTO `node_roles` VALUES ('35', '41', '3');
INSERT INTO `node_roles` VALUES ('36', '42', '3');
INSERT INTO `node_roles` VALUES ('37', '74', '3');
INSERT INTO `node_roles` VALUES ('38', '75', '3');
INSERT INTO `node_roles` VALUES ('39', '76', '3');
INSERT INTO `node_roles` VALUES ('40', '77', '3');
INSERT INTO `node_roles` VALUES ('41', '36', '2');
INSERT INTO `node_roles` VALUES ('42', '37', '2');
INSERT INTO `node_roles` VALUES ('43', '38', '2');
INSERT INTO `node_roles` VALUES ('44', '39', '2');
INSERT INTO `node_roles` VALUES ('45', '40', '2');
INSERT INTO `node_roles` VALUES ('46', '41', '2');
INSERT INTO `node_roles` VALUES ('47', '42', '2');
INSERT INTO `node_roles` VALUES ('48', '74', '2');
INSERT INTO `node_roles` VALUES ('49', '75', '2');
INSERT INTO `node_roles` VALUES ('50', '76', '2');
INSERT INTO `node_roles` VALUES ('51', '77', '2');
INSERT INTO `node_roles` VALUES ('63', '1', '1');
INSERT INTO `node_roles` VALUES ('64', '2', '1');
INSERT INTO `node_roles` VALUES ('65', '4', '1');
INSERT INTO `node_roles` VALUES ('66', '5', '1');
INSERT INTO `node_roles` VALUES ('67', '3', '1');
INSERT INTO `node_roles` VALUES ('68', '47', '1');
INSERT INTO `node_roles` VALUES ('69', '6', '1');
INSERT INTO `node_roles` VALUES ('70', '7', '1');
INSERT INTO `node_roles` VALUES ('71', '9', '1');
INSERT INTO `node_roles` VALUES ('72', '10', '1');
INSERT INTO `node_roles` VALUES ('73', '8', '1');
INSERT INTO `node_roles` VALUES ('74', '14', '1');
INSERT INTO `node_roles` VALUES ('75', '15', '1');
INSERT INTO `node_roles` VALUES ('76', '17', '1');
INSERT INTO `node_roles` VALUES ('77', '18', '1');
INSERT INTO `node_roles` VALUES ('78', '16', '1');
INSERT INTO `node_roles` VALUES ('79', '30', '1');
INSERT INTO `node_roles` VALUES ('80', '31', '1');
INSERT INTO `node_roles` VALUES ('81', '32', '1');
INSERT INTO `node_roles` VALUES ('82', '33', '1');
INSERT INTO `node_roles` VALUES ('83', '34', '1');
INSERT INTO `node_roles` VALUES ('84', '35', '1');
INSERT INTO `node_roles` VALUES ('85', '45', '1');
INSERT INTO `node_roles` VALUES ('86', '46', '1');
INSERT INTO `node_roles` VALUES ('87', '51', '1');
INSERT INTO `node_roles` VALUES ('88', '52', '1');
INSERT INTO `node_roles` VALUES ('89', '53', '1');
INSERT INTO `node_roles` VALUES ('90', '54', '1');
INSERT INTO `node_roles` VALUES ('91', '69', '1');
INSERT INTO `node_roles` VALUES ('92', '70', '1');
INSERT INTO `node_roles` VALUES ('93', '71', '1');
INSERT INTO `node_roles` VALUES ('94', '72', '1');
INSERT INTO `node_roles` VALUES ('95', '73', '1');
INSERT INTO `node_roles` VALUES ('96', '48', '1');
INSERT INTO `node_roles` VALUES ('97', '49', '1');
INSERT INTO `node_roles` VALUES ('98', '50', '1');
INSERT INTO `node_roles` VALUES ('99', '55', '1');
INSERT INTO `node_roles` VALUES ('100', '56', '1');
INSERT INTO `node_roles` VALUES ('101', '57', '1');
INSERT INTO `node_roles` VALUES ('102', '58', '1');
INSERT INTO `node_roles` VALUES ('103', '59', '1');
INSERT INTO `node_roles` VALUES ('104', '60', '1');
INSERT INTO `node_roles` VALUES ('105', '61', '1');
INSERT INTO `node_roles` VALUES ('106', '62', '1');
INSERT INTO `node_roles` VALUES ('107', '63', '1');
INSERT INTO `node_roles` VALUES ('108', '64', '1');
INSERT INTO `node_roles` VALUES ('109', '65', '1');
INSERT INTO `node_roles` VALUES ('110', '66', '1');
INSERT INTO `node_roles` VALUES ('111', '67', '1');
INSERT INTO `node_roles` VALUES ('112', '68', '1');
INSERT INTO `node_roles` VALUES ('113', '36', '1');
INSERT INTO `node_roles` VALUES ('114', '37', '1');
INSERT INTO `node_roles` VALUES ('115', '38', '1');
INSERT INTO `node_roles` VALUES ('116', '39', '1');
INSERT INTO `node_roles` VALUES ('117', '40', '1');
INSERT INTO `node_roles` VALUES ('118', '41', '1');
INSERT INTO `node_roles` VALUES ('119', '42', '1');
INSERT INTO `node_roles` VALUES ('120', '98', '1');
INSERT INTO `node_roles` VALUES ('121', '74', '1');
INSERT INTO `node_roles` VALUES ('122', '75', '1');
INSERT INTO `node_roles` VALUES ('123', '76', '1');
INSERT INTO `node_roles` VALUES ('124', '77', '1');

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room` varchar(255) NOT NULL DEFAULT '',
  `uname` varchar(128) NOT NULL DEFAULT '',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `data` longtext NOT NULL,
  `time` varchar(255) NOT NULL DEFAULT '',
  `datatime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of notice
-- ----------------------------
INSERT INTO `notice` VALUES ('1', 'Topic_ChatMessage/sub1', 'ooex5xK49v54rbsWVA2KUWT0TNb8', 'Sir. Lei', '大家下午好', '2017-03-29 17:19:50', '2017-03-29 17:19:58');
INSERT INTO `notice` VALUES ('2', 'Topic_ChatMessage/sub1', 'ooex5xK49v54rbsWVA2KUWT0TNb8', 'Sir. Lei', '欢迎光临', '2017-04-8 12:19:50', '2017-04-08 12:19:58');
INSERT INTO `notice` VALUES ('3', 'Topic_ChatMessage/sub2', 'ooex5xK49v54rbsWVA2KUWT0TNb8', 'Sir. Lei', '皓月科技欢迎你', '2017-04-9 8:19:50', '2017-04-09 12:19:58');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL DEFAULT '',
  `name` varchar(128) NOT NULL DEFAULT '',
  `remark` varchar(255) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `weight` int(11) NOT NULL DEFAULT '1',
  `delay` int(11) NOT NULL DEFAULT '0',
  `is_insider` int(11) NOT NULL DEFAULT '0',
  `ico` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('1', '管理员', 'manager', '管理员', '1', '1', '0', '1', '');
INSERT INTO `role` VALUES ('3', '客服', 'customer', '客服', '1', '1', '0', '1', '');
INSERT INTO `role` VALUES ('4', '讲师', 'teacher', '讲师', '1', '1', '0', '1', '');
INSERT INTO `role` VALUES ('6', 'VIP', 'nl_vip', 'VIP', '1', '1', '0', '1', '');
INSERT INTO `role` VALUES ('7', '铂金', 'nl_platinum', '铂金', '1', '1', '0', '1', '');
INSERT INTO `role` VALUES ('8', '黄金', 'nl_gold', '黄金', '1', '1', '0', '1', '');
INSERT INTO `role` VALUES ('9', '白银', 'nl_silver', '白银', '1', '1', '0', '1', '');
INSERT INTO `role` VALUES ('10', '普通', 'nl_ordinary', '普通', '1', '1', '0', '1', '');
INSERT INTO `role` VALUES ('11', '游客', 'guest', '游客', '1', '1', '30', '1', '');

-- ----------------------------
-- Table structure for roominfo
-- ----------------------------
DROP TABLE IF EXISTS `roominfo`;
CREATE TABLE `roominfo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room_id` varchar(255) NOT NULL DEFAULT '',
  `qos` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `room_title` varchar(255) NOT NULL DEFAULT '',
  `room_teacher` varchar(255) NOT NULL DEFAULT '',
  `room_num` varchar(255) NOT NULL DEFAULT '',
  `group_id` varchar(255) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `port` int(11) NOT NULL DEFAULT '0',
  `tls` tinyint(1) NOT NULL DEFAULT '0',
  `access` varchar(255) NOT NULL DEFAULT '',
  `secret_key` varchar(255) NOT NULL DEFAULT '',
  `room_icon` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of roominfo
-- ----------------------------
INSERT INTO `roominfo` VALUES ('1', 'Topic_ChatMessage/sub1', '0', '红木投资', '王老师', '299', 'GID_SUB_ChatMSG', 'mqf-dfbt0o9yg6.mqtt.aliyuncs.com', '80', '0', 'LTAIpFkq7b3IOvXu', 'QnewzgFKXHPBTOEV5BDAOSSGe1iuRA', 'http://wx.qlogo.cn/mmopen/KetjXWSVppsZ0icialcRKRXwLaCpCcoa664FHnSrnL5w6u9x1qyb8FfD35MiavwjibBQKiaQtdzyX9nyMibIicAx5htkLyIGEKJTAhB/0');
INSERT INTO `roominfo` VALUES ('2', 'Topic_ChatMessage/sub2', '0', '期货股指', '马老师', '388', 'GID_SUB_ChatMSG', 'mqf-dfbt0o9yg6.mqtt.aliyuncs.com', '80', '0', 'LTAIpFkq7b3IOvXu', 'QnewzgFKXHPBTOEV5BDAOSSGe1iuRA', 'http://wx.qlogo.cn/mmhead/ver_1/YMMNFRRsMIkFXX2dk0ibQ3ibteRGJDtvdeSryybatJuf8u3S5TIXDibUzuWYrFfyLRK1NrribBTylobSU0SGc6g0zLapBmFjlyJYTlicwtuhlLns/0');
INSERT INTO `roominfo` VALUES ('3', 'Topic_ChatMessage/sub3', '0', '贵金属交易', '雷老师', '768', 'GID_SUB_ChatMSG', 'mqf-dfbt0o9yg6.mqtt.aliyuncs.com', '80', '0', 'LTAIpFkq7b3IOvXu', 'QnewzgFKXHPBTOEV5BDAOSSGe1iuRA', 'http://wx.qlogo.cn/mmhead/ver_1/vHpDPyxQ1HqM0ic4ArXkoLVc45qcyEFDvXaNAw4gs1WZMF6gjZTq9pEqplRdOd4YwULibHOiao7BU2LGCibnYwiaPZDxLCyXnNfdic1MYyXDM60H0/0');
INSERT INTO `roominfo` VALUES ('4', 'Topic_ChatMessage/sub4', '0', '现货交易', '胡老师', '549', 'GID_SUB_ChatMSG', 'mqf-dfbt0o9yg6.mqtt.aliyuncs.com', '80', '0', 'LTAIpFkq7b3IOvXu', 'QnewzgFKXHPBTOEV5BDAOSSGe1iuRA', 'http://wx.qlogo.cn/mmhead/ver_1/sC0pvDJWtfcgaJfHJ3bKV17mmev7sowQFz2LibVLcDFT3n0SaQ7UksIAyVEvxYNsicNibj4FJ9MHMfSdVHLBVSW4iahrXJdOEfqhw2KloXxwvIM/0');
INSERT INTO `roominfo` VALUES ('5', 'Topic_ChatMessage/sub5', '0', '品种合约', '张老师', '325', 'GID_SUB_ChatMSG', 'mqf-dfbt0o9yg6.mqtt.aliyuncs.com', '80', '0', 'LTAIpFkq7b3IOvXu', 'QnewzgFKXHPBTOEV5BDAOSSGe1iuRA', 'http://wx.qlogo.cn/mmhead/ver_1/J9erBP5VgeauymSAtsoyqZssm8Dc7azh6mVoLKFyRanlTErwClcJT9d7eEfH0ThLF3w3HW9IajeJp8YtBZoVgQ/0');

-- ----------------------------
-- Table structure for strategy
-- ----------------------------
DROP TABLE IF EXISTS `strategy`;
CREATE TABLE `strategy` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room` varchar(255) NOT NULL DEFAULT '',
  `icon` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(128) NOT NULL DEFAULT '',
  `titel` varchar(255) NOT NULL DEFAULT '',
  `data` longtext NOT NULL,
  `is_top` tinyint(1) NOT NULL DEFAULT '0',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0',
  `thumb_num` bigint(20) NOT NULL DEFAULT '0',
  `datatime` datetime NOT NULL,
  `time` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of strategy
-- ----------------------------
INSERT INTO `strategy` VALUES ('1', 'Topic_ChatMessage/sub1', 'http://wx.qlogo.cn/mmopen/ZkOjia6CWnQ5zUp6NsyFylzsmoCbMfiay0P00BzTCAlDtoDaSrvtRsicLQlaib42XbzhzQ12peQAeXRpBNGFEicVEXb02SA74eDq1/0', '小雪', '金融分析师', '今天国家宣布建设国家级新区,雄安新区, 建议大家持有廊坊发展股票', '1', '0', '544', '2017-04-02 17:19:58', '2017-04-02 17:19:58');
INSERT INTO `strategy` VALUES ('2', 'Topic_ChatMessage/sub1', 'http://wx.qlogo.cn/mmopen/ZkOjia6CWnQ5zUp6NsyFylzsmoCbMfiay0P00BzTCAlDtoDaSrvtRsicLQlaib42XbzhzQ12peQAeXRpBNGFEicVEXb02SA74eDq1/0', '小雪', '金融分析师', '今天受利空影响,建议大家谨慎买入,高抛低吸', '0', '0', '133', '2017-04-01 12:19:58', '2017-04-01 12:19:58');
INSERT INTO `strategy` VALUES ('3', 'Topic_ChatMessage/sub2', 'http://wx.qlogo.cn/mmopen/ZkOjia6CWnQ7ynZJdVH1rfVmMsKOW0mU3ndkoibNZcE0Ao2GSpws8dj3GK6nicvLJ5AL4tbXr3Z4UH9pgBbxm70kL8RGohHOrgia/0', '马老师', '期货高级咨询师', '国家推进供给侧改革, 建议大家买入焦煤,焦炭合约', '1', '0', '666', '2017-01-01 09:19:58', '2017-01-01 09:19:58');

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `systemname` varchar(255) NOT NULL DEFAULT '',
  `chat_interval` bigint(20) NOT NULL DEFAULT '0',
  `registerrole` bigint(20) NOT NULL DEFAULT '0',
  `registertitle` bigint(20) NOT NULL DEFAULT '0',
  `history_msg` bigint(20) NOT NULL DEFAULT '0',
  `history_count` bigint(20) NOT NULL DEFAULT '0',
  `notice_count` bigint(20) NOT NULL DEFAULT '0',
  `strategy_count` bigint(20) NOT NULL DEFAULT '0',
  `welcome_msg` varchar(255) NOT NULL DEFAULT '',
  `verify` bigint(20) NOT NULL DEFAULT '0',
  `login_sys` bigint(20) NOT NULL DEFAULT '0',
  `audit_msg` bigint(20) NOT NULL DEFAULT '0',
  `virtual_user` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES ('1', '海丝会员交流间', '0', '11', '8', '0', '99', '7', '8', '欢迎来到直播间听课学习', '0', '0', '1', '591');

-- ----------------------------
-- Table structure for title
-- ----------------------------
DROP TABLE IF EXISTS `title`;
CREATE TABLE `title` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '',
  `css` varchar(128) NOT NULL DEFAULT '',
  `background` int(11) NOT NULL DEFAULT '0',
  `weight` int(11) NOT NULL DEFAULT '1',
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of title
-- ----------------------------
INSERT INTO `title` VALUES ('2', '至尊', '#992BAC', '1', '2000', '至尊');
INSERT INTO `title` VALUES ('4', '铂金', '1472629292.jpg', '0', '1', '铂金');
INSERT INTO `title` VALUES ('5', '黄金', '1472629336.jpg', '0', '1', '黄金');
INSERT INTO `title` VALUES ('6', '白银', '1472629375.jpg', '0', '1', '白银');
INSERT INTO `title` VALUES ('7', 'VIP', '1472629400.jpg', '0', '1', 'VIP');
INSERT INTO `title` VALUES ('8', '普通', '1472629522.jpg', '0', '1', '普通');
INSERT INTO `title` VALUES ('19', '分析师-胡老师', '#FF0000', '1', '9999', '分析师');
INSERT INTO `title` VALUES ('20', '分析师-徐老师', '#FF0000', '1', '9999', '分析师');
INSERT INTO `title` VALUES ('21', '分析师-王老师', '#CC0000', '1', '9999', '分析师');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room` varchar(255) NOT NULL DEFAULT '',
  `username` varchar(32) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(32) NOT NULL DEFAULT '',
  `phone` bigint(20) NOT NULL DEFAULT '0',
  `qq` bigint(20) NOT NULL DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `lastlogintime` datetime DEFAULT NULL,
  `createtime` datetime NOT NULL,
  `user_icon` varchar(255) DEFAULT NULL,
  `reg_status` int(11) NOT NULL DEFAULT '1',
  `online_time` bigint(20) NOT NULL DEFAULT '0',
  `openid` varchar(255) NOT NULL DEFAULT '',
  `sex` int(11) NOT NULL DEFAULT '0',
  `province` varchar(255) NOT NULL DEFAULT '',
  `city` varchar(255) NOT NULL DEFAULT '',
  `country` varchar(255) NOT NULL DEFAULT '',
  `headimgurl` varchar(255) NOT NULL DEFAULT '',
  `unionid` varchar(255) NOT NULL DEFAULT '',
  `role_id` bigint(20) NOT NULL,
  `title_id` bigint(20) NOT NULL,
  `is_shutup` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `phone` (`phone`),
  UNIQUE KEY `role_id` (`role_id`),
  UNIQUE KEY `title_id` (`title_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'Topic_ChatMessage/sub1', 'ooex5xK49v54rbsWVA2KUWT0TNb8', '', 'Sir. Lei', '', '0', '0', '', '2', '2017-04-08 16:06:29', '2017-04-07 18:24:57', 'http://wx.qlogo.cn/mmopen/KetjXWSVppsZ0icialcRKRXwLaCpCcoa664FHnSrnL5w6u9x1qyb8FfD35MiavwjibBQKiaQtdzyX9nyMibIicAx5htkLyIGEKJTAhB/0', '2', '0', 'ooex5xK49v54rbsWVA2KUWT0TNb8', '1', '湖北', '武汉', '中国', 'http://wx.qlogo.cn/mmopen/KetjXWSVppsZ0icialcRKRXwLaCpCcoa664FHnSrnL5w6u9x1qyb8FfD35MiavwjibBQKiaQtdzyX9nyMibIicAx5htkLyIGEKJTAhB/0', '', '11', '8', '0');
INSERT INTO `user` VALUES ('2', '', 'admin', '4bc08e686673c541e4c70815763955b4', '管理员', 'admin@ihaoyue.com', '1111', '0', null, '1', '0000-00-00 00:00:00', '2016-01-13 11:41:49', '1471264568.jpg', '2', '43117', '', '0', '', '', '', '', '', '1', '1', '0');

-- ----------------------------
-- Table structure for virtual_user
-- ----------------------------
DROP TABLE IF EXISTS `virtual_user`;
CREATE TABLE `virtual_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `user_icon` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of virtual_user
-- ----------------------------
INSERT INTO `virtual_user` VALUES ('1', 'wxid_sij2khjyocse22', '奋斗', 'http://wx.qlogo.cn/mmhead/ver_1/Sibo1NMq5SmMlVvPtQ98jW9bdiaFJAGasK7anwicke0ooEia5pG5FJsgSp7goR0thOoVkGnrxuJTY5v70Et6DUwl5xICo7J1srymYRbXKshMiajw/0');
INSERT INTO `virtual_user` VALUES ('2', 'wxid_zukvh2hekf8i22', '润心', 'http://wx.qlogo.cn/mmhead/ver_1/WWHicbvIvKAR0ajic7TyTUVcqHOLQP6ianiceYWFUa6cibaeNdRAS5uuLXTX0NC2sMO0gFUFNvvl0lu1xg5bZYdJuicD0Ub2Yyye1fayXOAj6weCg/0');
INSERT INTO `virtual_user` VALUES ('3', 'v1_2d950f094d89cadb224038f38aaf3d120f625ab0fa8e37ec49f37dd15bfb50ddfbaede7bb18e39990c047b2eb3885057@stranger', '渊缘流长', 'http://wx.qlogo.cn/mmhead/ver_1/fiabBIQYYSqHf9O7NGGxX7TrjIdUjn4YdpmWY7Z3eRRPnZehIQahYxv5cJ0yl5Q2BTH1dBiaMDIAbexoJ5mh5qJXuvmWapmKcgVKhFjGOD5x4/0');
INSERT INTO `virtual_user` VALUES ('4', 'wxid_3568225682411', '左边右边', 'http://wx.qlogo.cn/mmhead/ver_1/VoiafKnX8wSicwXjuav711QM4qoibPTXJ5oVjjQcOyvAnCgibopZ7EQKDIN8KVQnGOdic5npFmnvamqj6xt1ia5aCxOCwR7WCeDHTmZTxpgW4MuCg/0');
INSERT INTO `virtual_user` VALUES ('5', 'yangmeijun7999', 'ymj', 'http://wx.qlogo.cn/mmhead/ver_1/A6RU6mxwNWIv9JXmSzkqUdXpo0lQdnsauZZgKajdbwE2JWdRXkajuKjerBFuxH0WKPKJAvdHiblibIbfdte2zhicg/0');
INSERT INTO `virtual_user` VALUES ('6', 'wxid_7q39clu4gjj422', '醉相思', 'http://wx.qlogo.cn/mmhead/ver_1/3icicWzoatCHUBa8SM8Idibib8mmdcE9w9z1z6cqT0Yia6Xic9g1hNx5EG9RyN5qTia3w1ibnfYddXX0BpITXV0kjl6fbTwjwb4zfw9cibIYuEibfwJibk/0');
INSERT INTO `virtual_user` VALUES ('7', 'wxid_ft9r7u2s2v2o22', '朱', 'http://wx.qlogo.cn/mmhead/ver_1/Y5aBKTVbhgv1wgSf0p64QobPYoLYOfPfFjN9S1FrRsWo4iceMYFjI5Xicqvuu6jicCsSfJ7Bbfo6ibu1uDmSpvn8exiabkrAWsR0Ov643k0MF3nc/0');
INSERT INTO `virtual_user` VALUES ('8', 'wxid_2b8knj9ofcbm22', '兵哥哥', 'http://wx.qlogo.cn/mmhead/ver_1/JZQFGMQ1pTuMjKtUUpaCprDdCHyEo7wY6Lbxs2woTJ1koFjRvR1f2oo9EP5W4WQaiaTKI0VUhxiceRTfptXcVTOgfXO4Hpx7DwJd6icaibwWiaS8/0');
INSERT INTO `virtual_user` VALUES ('9', 'wxid_u3adqqh1c87f12', '幸福', 'http://wx.qlogo.cn/mmhead/ver_1/YMMNFRRsMIkFXX2dk0ibQ3ibteRGJDtvdeSryybatJuf8u3S5TIXDibUzuWYrFfyLRK1NrribBTylobSU0SGc6g0zLapBmFjlyJYTlicwtuhlLns/0');
INSERT INTO `virtual_user` VALUES ('10', 'gh_3f705be685cc', '搞笑语录', 'http://wx.qlogo.cn/mmhead/ver_1/wKnic0ickKxTjmbfTKF4wL3Gw24C09dJicVkZKyWE3iawMkX1mLVPp0rGy72nbhbzia7YdbsV6x7r0nwL9GCjz85G1BrdjdZiaHKjAjmNbrdRdHE4/0');
INSERT INTO `virtual_user` VALUES ('11', 'wxid_32fwgj2o5gdy11', '零诫', 'http://wx.qlogo.cn/mmhead/ver_1/vHpDPyxQ1HqM0ic4ArXkoLVc45qcyEFDvXaNAw4gs1WZMF6gjZTq9pEqplRdOd4YwULibHOiao7BU2LGCibnYwiaPZDxLCyXnNfdic1MYyXDM60H0/0');
INSERT INTO `virtual_user` VALUES ('12', 'wxid_6354743547222', 'Furure', 'http://wx.qlogo.cn/mmhead/ver_1/ClIoaQSydiaHOnTsSh4Vod0C5JfQQahFLiaXQOROtMykSIP22kxQ5NRvM8TnaMzx3gmNuicvjOBdvnOuIqT7w6cpu2QVcVbalrfNaoWsNOxunI/0');
INSERT INTO `virtual_user` VALUES ('13', 'wxid_f6ehx434htwv22', '刁蛮明', 'http://wx.qlogo.cn/mmhead/ver_1/sC0pvDJWtfcgaJfHJ3bKV17mmev7sowQFz2LibVLcDFT3n0SaQ7UksIAyVEvxYNsicNibj4FJ9MHMfSdVHLBVSW4iahrXJdOEfqhw2KloXxwvIM/0');
INSERT INTO `virtual_user` VALUES ('14', 'wxid_ycuux9hk9za621', '黄坤', 'http://wx.qlogo.cn/mmhead/ver_1/G7XD6xiaUibDE0JwMyrH0RxS2K4GLI12wyUhMhJ6KRzazjAsXQJrftZPa1vjctfc8fpxMgVelx0OYMub88nLaTGPDD3OJOOibFQpWhnBVyyuF4/0');
INSERT INTO `virtual_user` VALUES ('16', 'hl346220515', '’HaN、', 'http://wx.qlogo.cn/mmhead/ver_1/cV6NXLOOwf23Vxm2YVdewwI6444pndoZKs18byyGsZaGeczdP2LFBzyzcJeWwZTE7w6NgBNE5P1hIkfVVWuB8nLgPYoklLuCVftbwNvibA7c/0');
INSERT INTO `virtual_user` VALUES ('17', 'wxid_hltx5pe9rs7y21', '一指通·iptv@黄全瑞', 'http://wx.qlogo.cn/mmhead/ver_1/TVyuhlZypZnuNicrtBIo3AibQHCg8A1NuTutP2De5ic96dgtCHeFOFXdDWEHPoUoS5IVx29MSibczk2gJVUiceQ36fybJYvOUDFuGXJyPLtj8eR8/0');
INSERT INTO `virtual_user` VALUES ('18', 'xuanzi_lover', '何博', 'http://wx.qlogo.cn/mmhead/ver_1/J9erBP5VgeauymSAtsoyqZssm8Dc7azh6mVoLKFyRanlTErwClcJT9d7eEfH0ThLF3w3HW9IajeJp8YtBZoVgQ/0');
INSERT INTO `virtual_user` VALUES ('19', 'zhangxiang1200', '宾克斯的酒。', 'http://wx.qlogo.cn/mmhead/ver_1/RiackNO5iaLjGzDRBv6yvJJ4YKXCrXD3nHTbICTrwcTV1VjuG7OOUibvVicHXibgIgEKIMjHx7TIKH7ERzYdMFe1dTZNlDlCGH9CNWSXwOc2IGwU/0');
INSERT INTO `virtual_user` VALUES ('20', 'xiaolinor', 'lidy', 'http://wx.qlogo.cn/mmhead/ver_1/Cab3afbbhicCh7kicaGWvqpTLuszNvcf7Bu17jH1nAyCE5MoItPV8VFZc8eYFZbJj7l9WFM4RM8y2ho5jGVPuhVA/0');
INSERT INTO `virtual_user` VALUES ('21', 'wxid_ys7c4ih3kedy12', '陈佰康', 'http://wx.qlogo.cn/mmhead/ver_1/d8722YYeLowTjkTuSkJzibkHe8S3SFyGpFY0nXkcdniaql7feMnznrZGaNVRSbhicWQjwwyYJsPwicPB4vbTVh0ibD8aCmVfyQRnGTdDy0cH4SibY/0');
INSERT INTO `virtual_user` VALUES ('22', 'wxid_4263852638422', 'Alin', 'http://wx.qlogo.cn/mmhead/ver_1/n7AZTnqowpLkLxlr6mLdmxlTvkr6BjibeRB7GA8P3oBs2ezsbVj5DhzZUo2LGtg6FMum2xgrL6hiaqaQSGP6XxxFMjzliaZuPraPLn7gEf8QUk/0');
INSERT INTO `virtual_user` VALUES ('23', 'wxid_3b3kphl305yn22', 'jvaemape', 'http://wx.qlogo.cn/mmhead/ver_1/FcBbPVW1r94NemDGRDzXCibXcCW0mI0jreuGgw7GZaicqGPLEJn3bbsfRuxbicYyoO0sBt8YnFx5rcy2q01pNtQBw9uian4bdznZicEtyvfxzGyQ/0');
INSERT INTO `virtual_user` VALUES ('24', 'wxid_e32xoxbyypoq21', '孟宪涛', 'http://wx.qlogo.cn/mmhead/ver_1/LZwGISJj9Cdyaia7tIbJoCIGgXMBOPPzV6TLTKAOe5ASh3kEFswGp7pzPeYS8mGYxrKF7qaEHVGh2mEzCrMZ6Mk62ic6p6sPqNGVicxS2bzFib4/0');
INSERT INTO `virtual_user` VALUES ('25', 'manzi_11', '曼', 'http://wx.qlogo.cn/mmhead/ver_1/Utlog87lTws9oCGE6aaJDISkiaBeq7ibuTykIFxC6qmfpgpxAbCrztkzjDSp5o5zU8puQcR2HTZiaBXRvzqQ0qShrX9j4yu9iaCgFHMRmLqsJTY/0');
INSERT INTO `virtual_user` VALUES ('26', 'wxid_kxsprllrxzt021', '李煜nl', 'http://wx.qlogo.cn/mmhead/ver_1/vynlfvxYF8vaU8lTL3vy86XhQLEzUnMiaGXv0vspeI7ibRzwce9LjowLu7sFpE7r1jgBVsY9vzynaUjhxCZqnUw0zmoggUibN5zAYibnLicsrqOM/0');
INSERT INTO `virtual_user` VALUES ('27', 'wxid_s7v05mui1lok21', '姚と杰', 'http://wx.qlogo.cn/mmhead/ver_1/rIdaPjeRZdzlgzEcE8BgDwP6PQM6v311qeiclSuY6OJE2UneEgYdO1UKTeics9LCibdSUOBmUbPW564g5sFA8rbfM38XIO1LiaibibytPmWbMicrs4/0');
INSERT INTO `virtual_user` VALUES ('28', 'Rong1208c', '张荣', 'http://wx.qlogo.cn/mmhead/ver_1/hXtGIMLEd76MmqWsIFS39pgs0OZGn44wvZNebnXmtVdibiaut6TuxvzkpoWZzziaASSiawZw6ibhPqpP0YCyU3wOoj1gerqLPyCz3TOSjcZuJtRA/0');
INSERT INTO `virtual_user` VALUES ('29', 'gh_27bc217d5422', '快点扫我', 'http://wx.qlogo.cn/mmhead/ver_1/YA69e6DJxVhib4v0sbjabKT8icIXbp85o3ibXG61wRLIGiaoZ5MDUswGL2rPmTcv7nr4ObZmS3HzmD1URoTKNHvOUKwc4AUcB0dz1OL4PXSJibqM/0');
INSERT INTO `virtual_user` VALUES ('30', 'syj331343', '宋义杰', 'http://wx.qlogo.cn/mmhead/ver_1/ReqhbbD2mB5vibluRTHAIocRXblJQpN6sQ9XbvazlUIx0fTibHVhrstwb2S0H1BEPB7nktuaMJujoyvibjZhqmnqg/0');
INSERT INTO `virtual_user` VALUES ('31', 'wxid_p4lfw6bozcg321', '猫.喵～～', 'http://wx.qlogo.cn/mmhead/ver_1/z37F4IwYU0qacFwf5hLySPqxgfJSlWGOL0skzZHoWs4QZnVkUd7gFoVGicwfWL5XykT9ZnOP7TuBRrLNiayBSTuEfekWOzGZOl9OaibGmpNNaA/0');
INSERT INTO `virtual_user` VALUES ('32', 'gh_5787e1d847f6', '骆驼', 'http://wx.qlogo.cn/mmhead/ver_1/UGeWEdUShwtCS9oP7eSibx4RD7EqQ4BLtTYKWahq5I3193oS69lLyU4XMuWsibQicIhJt7pnZs2KfiasznvVQMrm2MLsIQg3Ed1q0ueuYlnHKng/0');
INSERT INTO `virtual_user` VALUES ('33', 'wxid_pw3c22hlyvfz22', '木木唏', 'http://wx.qlogo.cn/mmhead/ver_1/JLv5LRZIqdhNFgEz5AtOumNypcibkucNpHRYhjUyokdXWNHEBMTUlEDk1Nr8voQK4hcEBricOiazf1zibwGgxAdfEWqTzM8sibNpzwqkBbXGw4zg/0');
INSERT INTO `virtual_user` VALUES ('34', 'q13197175551', '谦谦谦牵手', 'http://wx.qlogo.cn/mmhead/ver_1/ITFE5UrKbBFmTQfF6WTS2KaUAibXCOClZvhqIXPBliaTfny8BkfMz8Dun266XE0YibyH7sPeybLX5H6URRyHHQNVeVnOGwKd3eZ5FH3eJicHdico/0');
INSERT INTO `virtual_user` VALUES ('35', 'wxid_9136x1kh7h8i21', '王馨', 'http://wx.qlogo.cn/mmhead/ver_1/5OJOf3rrEXtogXd5HNdKLgdtPh2riayFVzumDicItWSoZswvZ362wLCrh7ubHg00Qaia2vHGvib2rI8f9ehL8HsrE11nw102xGzmoUJgHcv4eQE/0');
INSERT INTO `virtual_user` VALUES ('36', 'wxid_4047330486612', '123.木头人', 'http://wx.qlogo.cn/mmhead/ver_1/FOV4xGyRuK7eG8mkuZafSicRLnAE1eWEfxdfxZRmbDvEsMlg4TXMduTgzC2ehy08Zy46oZA1z6s09nfNFETLr7g/0');
INSERT INTO `virtual_user` VALUES ('37', 'zl198801151664', '张玲', 'http://wx.qlogo.cn/mmhead/ver_1/xvI7N0z3Cw5tLmfhenkjIauBGD2O0d2RyneTa52CbHO2E2KEIFZib0lyoUXTVhj9NVWgvm1dmhtXxWakSPqOsPjPagDYJ1rOfXTK7tiaoMgQ8/0');
INSERT INTO `virtual_user` VALUES ('38', 'xiaojian1207', '肖建', 'http://wx.qlogo.cn/mmhead/ver_1/L236NhNNb8b6EEJ6DHl9qAtCm2gmptLDvBO2qH3AxibpqFlwFQRtp61QLiapicyQToppl1QoQU1kDqh2GjUf8US1oHspDvgHoZ43TA5wpMZ1tc/0');
INSERT INTO `virtual_user` VALUES ('39', 'wxid_7879208801412', '         ', 'http://wx.qlogo.cn/mmhead/ver_1/g4HOn8Cnx3fzibVXUnWyaOXXqlRYS9k7MUxpksAibE5ap0jrnjss970GzA51ITOiaYStFe78oYDvn1ueQtJTou6rdUmN3Mo6LGj8tzVkj9kdWw/0');
INSERT INTO `virtual_user` VALUES ('40', 'wxid_5ycptelqr9an22', '有位佳人！', 'http://wx.qlogo.cn/mmhead/ver_1/7qicibyG017FBiac3uwGRHLvrnk882Ko0R43T7icSx9WbJ1SEiaBibFAzYYHzoWSEzU9PiaYdSPlqmoHnicSWwRomOZ1kfibvoyeJAYT3ndJ82OsaWwE/0');
INSERT INTO `virtual_user` VALUES ('41', 'wxid_hzr0hz4bqdbm21', '王潮', 'http://wx.qlogo.cn/mmhead/ver_1/aicu6VN8N3zNquOf1cD4BRFoC8ZibXMLWT72u8N5YAdib3kX8X628voj47zSUDmkxNNu9RRuN6yzm8uT9MgdkxwSO1GibOxp6rNxMsIe7icVFibdU/0');
INSERT INTO `virtual_user` VALUES ('42', 'wf-297939339', 'ゴホウ', 'http://wx.qlogo.cn/mmhead/ver_1/Pb9ksyefYibOtIIopOlXVmgc6xdpFbw0Y09y2LZ9niaBHMFg5qtqAhGB0SC8L4gpCZEuWHZ6QSWpxCfibAg89tibxqytDNhvkdm9edib0gqlJgyk/0');
INSERT INTO `virtual_user` VALUES ('43', 'luyaping2653', '小曦', 'http://wx.qlogo.cn/mmhead/ver_1/8oXiaGaVn319G9hbibgzeqHGdjpEoABM5n4AthjW5MAgibmNZoMIHh8PkFofjJes6lJkf57v3B3VAVAoxC4XUMOWw/0');
INSERT INTO `virtual_user` VALUES ('44', 'wxid_kkm818kwql1m21', '总有刁民想害朕！', 'http://wx.qlogo.cn/mmhead/ver_1/z834J475x1lgW8hKBnChlia9ulAtVBN76tNKqNjwb6iaicZB87d18D2lGrFO7BrmA2sWn4ibP1bg71txUcpxC3P02w/0');
INSERT INTO `virtual_user` VALUES ('45', 'Mzero_', 'Shei.....', 'http://wx.qlogo.cn/mmhead/ver_1/1aP6wlib3YYdeqMm4LcE1ngkbia2HxSG3G1eDHW5o6BVfLqAtXiasrwSpPMibMaIFsMiaTLBVtEGRb0CibwrpKicZ2cGVlwYFJmzxqC635S6uK7ol4/0');
INSERT INTO `virtual_user` VALUES ('46', 'wxid_sinajevrgca022', '清风醉蝶', 'http://wx.qlogo.cn/mmhead/ver_1/wScNJATbFyN6qGmeRHmRLz1eQRqickgOjxOy1csmOdvkvAUYhVXPZ9hKZwYvUhVpy0vHKMBkoHSwibh55X4TyBulLMSTV8XibxpZAe54CjqUCs/0');
INSERT INTO `virtual_user` VALUES ('47', 'wxid_zaplbh6ug63y21', '壮志凌云', 'http://wx.qlogo.cn/mmhead/ver_1/EL1nYpwD3Xys9KYWhByCpnXibOlianQ1zOCQ9Qhub2N3ygS9t4clKsGbLTmW68MnDia78Hiaa5LRD71iaSmjcsFxTx39SLQ1jELLcWicnwt5HkneQ/0');
INSERT INTO `virtual_user` VALUES ('48', 'wxid_8yk5txygsv6m11', '阳锦            ҉҉҉҉҉҉҉҉', 'http://wx.qlogo.cn/mmhead/ver_1/DohSt9kX27OemoUDWAgJSwY4GiaWhTMcaXJQL5bH4VJiapsmXVfDfD7XpKAb2abiaFqQRfCImS3kWgfqkTGR4GmZrW8w4Z5aEvLPWvSmETula4/0');
INSERT INTO `virtual_user` VALUES ('49', 'shen__jing', 'vip_静', 'http://wx.qlogo.cn/mmhead/ver_1/3RyWpRP44TffS2BzfYQh7lB5Td4WHAXM3DYs51cTKCtGzQePOAvIbpulx0kaHuqcgGVEqdeicvsuE4bxq6Tf1Y8ibc8aujwBw0Ot2qDdmv2ek/0');
INSERT INTO `virtual_user` VALUES ('50', 'gh_3dca91af5abf', '梦幻西游', 'http://wx.qlogo.cn/mmhead/ver_1/S1g8vylIWLPYyCaRxxYPnQ5YUGLV5ujzUSr52or2R8ibMiceFicDYicBfxBZLuoN63r5zpe3eMIbvXjrUdTK6BibhO8EMBx375VuNOxId4u2w2tc/0');
INSERT INTO `virtual_user` VALUES ('51', 'wxid_jvebqbmr9zm522', '曾梦洁', 'http://wx.qlogo.cn/mmhead/ver_1/IWK3et50PILymgFnjdb4QUUxSft0NXezSJhWud4UM4fQzSVxicuc9dnUY6hOtCVB8mOoUiapRHzPhGCXzMI3LxoZX9TYy3akjLcZxKYbxXakg/0');
INSERT INTO `virtual_user` VALUES ('52', 'wxid_m2dkhi1io4lp22', '纯', 'http://wx.qlogo.cn/mmhead/ver_1/Hg8IE93vMkDUa4Gbayu01ngjo62R8iaSz44CWKlwPDBR5rKlxtVUJ5AL0tUKOfISAOhzb6TG06M5oyGR0Qa7KMJpNg5RibEdBseBtzmZYXop0/0');
INSERT INTO `virtual_user` VALUES ('53', 'suyong66173', '本帝不亡，尔等终究是臣', 'http://wx.qlogo.cn/mmhead/ver_1/X0L850z3GQKOL5WWPEsniccbVYjKuh8UzIyxhZwf9EOOKUMBLNicxic3Hhk7pXCFSPXL1S5sUwUttOQKcpdEL6ic1O98Otvg4vF5VMXiat7O5JVE/0');
INSERT INTO `virtual_user` VALUES ('54', 'shuangshuang306868', 'carol', 'http://wx.qlogo.cn/mmhead/ver_1/NNjdibW68LlT0yypcjsoRqYNTIkia2ia43xBaqklFVAmicrlHTR3weB163FE4ojiaMoYIlHibzhJ0J3ziawnn6Ktwv9TQ/0');
INSERT INTO `virtual_user` VALUES ('55', 'ni_233', '大米西西', 'http://wx.qlogo.cn/mmhead/ver_1/FQ3cKdf5UfXKAKTjQlqL1KtTjsebZlTjfKQldX9teVAbc7iaNKib97yrRBv77wJMbPDMh2GznAYbb8ibZhRWicHyzJ80brEr0AYQ8ib6I2zmM6uY/0');
INSERT INTO `virtual_user` VALUES ('56', 'wxid_75u4imy8hisy22', '夜晚的星星', 'http://wx.qlogo.cn/mmhead/ver_1/r8qxl8ZlEabYwXmCJ9icnKWicc001VgsseEaXKibFVic9y7EKIeFYot7pW9yzEXDIjwia8ludoY9aibIsaPsV3adD8W9KyHclqdoj8iaJWTV0Bp59Q/0');
INSERT INTO `virtual_user` VALUES ('57', 'wxid_wp2fqk3s3hq611', '坦桑石批发', 'http://wx.qlogo.cn/mmhead/ver_1/K2mtdIURGHYxrHk4LkRVL562CMbSyGYtOUBwqMl68DRE8vuYspibHLcicTZpQ2AFAH0fNMFPLe7wzQleV93ddHyg/0');
INSERT INTO `virtual_user` VALUES ('58', 'wxid_yv8i3n9y3im822', 'LYT', 'http://wx.qlogo.cn/mmhead/ver_1/dGDQM0fIsKZum2X3JJsUjCxpDicd6ldW1LSCwrYKGvm8Fbv1JLWBfYzIQymU4bDlfUComOU4YhI4C7LmcSxibuhWDWAcHfJEXtQKeRjIDrQ9k/0');
INSERT INTO `virtual_user` VALUES ('59', 'zhangyawen_forever', '词穷所以沉默丶', 'http://wx.qlogo.cn/mmhead/ver_1/2e2WFAiaNa4P4X1dfDwribwLjeOgMqsqJwcQOPWB7Nmry4aR9z1CHzMk7GO5HduhoFvNvCZRAor9osTQfoFclkRRlW1ZKibPGb4fkWMASTK3r4/0');
INSERT INTO `virtual_user` VALUES ('60', 'wkw563684134', '王', 'http://wx.qlogo.cn/mmhead/ver_1/VuX0TDBJBRWkD2OPj612nJFdLBlvL3iajzLBEIVQ0jSAyicATbWiaxU17mp6ibflr4zaSboxlRtx3Jp8mN1EAOnU0uicwl2kd42QPqJgJQxcMV2o/0');
INSERT INTO `virtual_user` VALUES ('61', 'you_with_', '周游', 'http://wx.qlogo.cn/mmhead/ver_1/85NN3DkBaktPW203DHiaBp7YwIiczIVl5GExa3JtRloWt4cdpqfOVfXZCK8CibLpsS0uyA1AzTO0iczmhVNQhiacSeBrrC24P00Q3icnkSAsum1AI/0');
INSERT INTO `virtual_user` VALUES ('62', 'wxid_8fu4gki5fax112', 'moooi', 'http://wx.qlogo.cn/mmhead/ver_1/YEFickuhZY5bATCPc9ByRyjXtZA3bRibGiabfGromLAxVNOJECZqr1IKNI74C4SQBJDT4L9x33v9bDeehdTflsxyaygv6vKKzrVqAFiaOs7RHzM/0');
INSERT INTO `virtual_user` VALUES ('63', 'wxid_9jw4o0ypf4px22', '罐罐罐罐罐', 'http://wx.qlogo.cn/mmhead/ver_1/tG6tKr14kibxcicXRbpicNeOhT3uicky7H90alLxcRYDu86mwOU1jaTLC0kYHwHBs7A6M3KYIhHCwnLCntzr2Sa4wE2qErMianZPVClRHbk3xhQo/0');
INSERT INTO `virtual_user` VALUES ('64', 'wxid_wl2sgwft3rft21', '武汉英斯利  徐洪亮', 'http://wx.qlogo.cn/mmhead/ver_1/YOE54m0ymmRGrzlA3JPycrjcGeY61gSxSPzvxJ0kGwLEp4Z47e8yU040HQ44qgMr2ClaN07vWeUO2Y4mFDmQS2IFMRAHSRRfAIOBEnc42EY/0');
INSERT INTO `virtual_user` VALUES ('65', 'wxid_fllio97oyx7b22', 'Sandm° 旧梦颜', 'http://wx.qlogo.cn/mmhead/ver_1/xe0iacCHqylibIXlayNySryYwoYVYQeN62eNu3cdiaa71Ie3O4U0uBbsKf5d60ZOp9TYN16yhGjibiaXC0EgqwXgufzOsqicm1Kof8nhpEHqXjicHQ/0');
INSERT INTO `virtual_user` VALUES ('66', 'panweizhou', '', 'http://wx.qlogo.cn/mmhead/ver_1/qVCZnJmC0lkDKgibyMVBpSZ5S28AzWx1Pssd5fXlfmYUkrFxdHVPRFy905HmT2kYzgLIsHGbgbcHicU1ByhiaTmxCJydlQ7pTle1LsibrciaqMYE/0');
INSERT INTO `virtual_user` VALUES ('67', 'showtimexyy', '亚雨、Come。', 'http://wx.qlogo.cn/mmhead/ver_1/ic82y8mP4S3bUs5MGbgHRWFJB8JWnnJibTjrXxggnNuiagkFibwWrjSdg4JctaD6jMYnxZYjyFOMPeOelZoP65WnuStR9AGWOnJtUVzo8HkeHs0/0');
INSERT INTO `virtual_user` VALUES ('68', 'wxid_0346063459212', ' MIDI&慧慧', 'http://wx.qlogo.cn/mmhead/ver_1/tIRs58XY4OblS9kicnXsFnxNIPH2S29HfypC0ibaw2qgOUJ9UBz90yibdCvIBxfFBT8rP0rwedkeGSJHaTnzU2icIQmGqmxcVSG2IwzJN01yHmI/0');
INSERT INTO `virtual_user` VALUES ('69', 'wxid_ngr8v5yqs0er22', 'Ai 苗苗宠物', 'http://wx.qlogo.cn/mmhead/ver_1/Gl42Vn9n8NTRSk6beIGkFesvnibRibqNASiauYIushLopSGHFERNkxXzVe1x8Zegl5PWh4c49qgKFIzialOF5FNI0k8HboYdtz2zV68p8SxJoxI/0');
INSERT INTO `virtual_user` VALUES ('70', 'wxid_vnhaa0cqbhci22', '襄河味儿', 'http://wx.qlogo.cn/mmhead/ver_1/xPO6mVE6Ct7aOsS5Cr2xvdeCWdNLZVHtaP3x4FyokCmt43nme6L376DLAuiaYUAJ6K1rsQicsyKto8yAS3dlrLFNNkVBU5ibngqHic218ib2Yhgo/0');
INSERT INTO `virtual_user` VALUES ('71', 'bryankun', 'loic', 'http://wx.qlogo.cn/mmhead/ver_1/tLXFgkuv3YM7GSced3sCibEvhwNKYYJlhcDtwecSWQs3NcvL2icDLIn6ArGGFwmazXnAxibKM8ykcYzcqYPDVibdibA/0');
INSERT INTO `virtual_user` VALUES ('72', 'wxid_f7w2sgrh7rmi22', '小雷', 'http://wx.qlogo.cn/mmhead/ver_1/fvUtL3FMby0T6FzP8J2picyicFXbs6sGuUxXGbibZ548rPI9PT448pfPMtrtIicvsFicPXnbaLHrIib05VtouGvvEoZloW8YVx3Y82R4iamkurutZY/0');
INSERT INTO `virtual_user` VALUES ('73', 'cbw799327195', '鱼儿', 'http://wx.qlogo.cn/mmhead/ver_1/zibRKxG4xymG0CpCNM5SJiaAMaxdiafko6adxhrcKmzGO5g4dwTFmmMOAfvTIWUCQfQ0crleB8icAob2BIFlYduINQ/0');
INSERT INTO `virtual_user` VALUES ('74', 'wxid_zf3betm76sli21', 'ly琦', 'http://wx.qlogo.cn/mmhead/ver_1/uvJKPzCypOEefuDM7qdxXUSUiaYz7KKQrL9AiarVzkicwhGHWSf0iaibvTGs2ib1K4RgqalyYAJ2crGJhekV61gqZ5oTPv4biaPQHNopBHfictgAS8s/0');
INSERT INTO `virtual_user` VALUES ('75', 'wxid_ce3oncu0jc0a11', '黄乔', 'http://wx.qlogo.cn/mmhead/ver_1/c1GGogaKwmkDT35PJ7icXTKXPwL0evfIP2EMGAuRgBbibX38FETXndLXGTvJaoPjKS7cnaqVv1aajZibfUf5sq5OLdIYXjus7yicsAEqH9J4LrY/0');
INSERT INTO `virtual_user` VALUES ('77', 'wxid_uq03rs7rzpox22', '超越', 'http://wx.qlogo.cn/mmhead/ver_1/Td2uwFu3NWMqvREKq5QKkCpQKGE5KjDibPYX2aHdMjMym87MRTRYjdwJfQ0qTMEDSWNIgRkNgWppWdP2AKeicAcfF0SSByz0iazX0kswj25Evw/0');
INSERT INTO `virtual_user` VALUES ('78', 'sh1qi-c', '飘飘儿', 'http://wx.qlogo.cn/mmhead/ver_1/gIgwTfQicicfGWzkYsSRjGyGvkKwXJJtoCV3ek6QC5zhLkxIs0yYZxxXoiaicTKicCoGta4jQPSSr0UTuWz5rEp2qCoM4iaPNQ55GP1rHMbpABp28/0');
INSERT INTO `virtual_user` VALUES ('79', 'wangbo2128', 'AAAAA', 'http://wx.qlogo.cn/mmhead/ver_1/sNiaMHic6ibuVI2RKrdcZcpGF5Ojq5gvZ1VQdrMfj8ibQkTnicQZCtlHkiaeeIEw72icIxYWY4J1jYMNlFUicBUeNib1qia1KrzNe2wpbbfgUO1EjNvyw/0');
INSERT INTO `virtual_user` VALUES ('81', 'wxid_24ki3cdxiy6p21', '呵呵', 'http://wx.qlogo.cn/mmhead/ver_1/Nj7cxLz78c1PahR94QGTbJYqZSrDh46PcNftfNTp6R7eY3e0lxl0Mhsxqf2DpLaNHG65GuFezibBctulECmq4sZqGg1w6sWAFwgVn30WzreA/0');
INSERT INTO `virtual_user` VALUES ('82', 'wxid_28h2cco1mjf621', '圣雪纷飞', 'http://wx.qlogo.cn/mmhead/ver_1/ZhJ10pkMdSOickjx8DSAIc0sZe25AexomPXSkGrJM2cQiclAXHCCB1n6Y6lriaEE01yGQNOPhicMgUKAFU8vdicRZvQ/0');
INSERT INTO `virtual_user` VALUES ('83', 'zhouwei5708', '天蓝', 'http://wx.qlogo.cn/mmhead/ver_1/Bia1vavPgR1D9yGLpO3YWQ9z8939RMjHUTbFsiaekibsLib7Xab4m1mrB18FOg69eDicrrrXVAC0OeiciaaNiaSQyGrlfLVdxjudL00KquCHrbOK91c/0');
INSERT INTO `virtual_user` VALUES ('84', 'fyc394531669', '付应超', 'http://wx.qlogo.cn/mmhead/ver_1/Gb971ZhzbA9L6DCQibzKYRDd1s5Mmcd5blZyH5vw8RZ2QjNsbhSIS76w4xLjqPGicIjxCibRxibB2bibsEW4unOGY2Q/0');
INSERT INTO `virtual_user` VALUES ('85', 'wxid_izgrpsey08bm12', '  Cao 小曹哥', 'http://wx.qlogo.cn/mmhead/ver_1/3WDtrNiap17Ofd94bGgojkdClS7JNLmIEF6heqjlfoYZMkt6iaOTicdNzjRvia2GU2Kc9rs5AnuFHoagWHEiaFIiaMiaiaplx8lo7UaMP4lbibF6YyQo/0');
INSERT INTO `virtual_user` VALUES ('86', 'wyyanlp1314', '说的跟真的一样', 'http://wx.qlogo.cn/mmhead/ver_1/VUjVzdZSUgkFzSU08rfavQibybbVx2R3IDGJCgZnDEpEnqwZvc81Ljv4PgRTIKIfWIhIicEnZUHMcVVrjrcGNS64XYhxe4S8209poK3h9ke80/0');
INSERT INTO `virtual_user` VALUES ('87', 'wxid_92wm4zi7qxms12', '彭玲', 'http://wx.qlogo.cn/mmhead/ver_1/lLyFBblcYfN35bN31iaAEpEqKGHGmXA28hItS9GMaLEicNnyY9MSJJLRTcQak3icXHcvQjpxPYHVxRr6UktPHl3m6yevVOHVNJGKjvQcWeWgtM/0');
INSERT INTO `virtual_user` VALUES ('88', 'wxid_17qbvgm6x1v922', '王爱红', 'http://wx.qlogo.cn/mmhead/ver_1/fJKzGTOTvOILib3aWyqy7HfZ2jiapAv8jbmHV29DbmH2HIytxLia4Woia3nSWz0FEv4kSWxwMcqkF6nK0S9zyc7TNhicKPicU8VLQJf2KZrtzGUVY/0');
INSERT INTO `virtual_user` VALUES ('89', 'v1_91c6d45d0afff7b70ff59a31e83aff7570c5410cd857057eeecc1e3ff190360f86e34938632a66bb94f9d7579dc32663@stranger', '鼠胆龙辉', 'http://wx.qlogo.cn/mmhead/ver_1/ZWugkbODukibC3LPq2NHrDsgZMn7fIq0C6fQBicJdrPkl3J5PjF0mPLJjFNicYaWEKibKmkBnnBuvj4qhlic8ItuPIZslnkAdI5icdAzOMSa19rpg/0');
INSERT INTO `virtual_user` VALUES ('90', 'w358123395', 'tang•ca夏天', 'http://wx.qlogo.cn/mmhead/ver_1/gic82av7FyS1VfGzQ0rTdQBl4ShB2PaDuDYCJQoeYaPGibufB4k30AxAbLvOIictsynnicSMiav0R1YVFrVcDOLZxbyDqMRKvbibk64tD9u1d0TIw/0');
INSERT INTO `virtual_user` VALUES ('91', 'yhs408397872', '万宝路', 'http://wx.qlogo.cn/mmhead/ver_1/0jIcZvkZgxwqq4u6wkREtsHwoE8MrfuexNoxMtbVKBAJRYMMOicvW9vCsspSuFffGtWxqjy0vJVLficPAAiaWmKXq5plwakgGqGj9UmribDxwBs/0');
INSERT INTO `virtual_user` VALUES ('92', 'v1_2856b94383bda0fe1d4d198959a01e9517a80af987f8f84bddf17131ccb7f63699edbfaf5c49d0fe5c9969b2523f55ac@stranger', '猴哥董志忠', 'http://wx.qlogo.cn/mmhead/ver_1/63qUtngOicAzW2wg5icGqmdnHFAvKaTh0y9DViara9zmInKL8GCYicPp3y02fhthkEav97F8L83vpKVxskCwkGpgWLVHQdVeia40EJPiaJy3F5dFU/0');
INSERT INTO `virtual_user` VALUES ('93', 'v1_64ef36d36150eb1ec976c644d370008bce8c1cf601f3ffd7f8dccbecfbe074a8ed925ee30691f4b0a4cd9d8ab468e845@stranger', '自力', 'http://wx.qlogo.cn/mmhead/ver_1/B26sAqCIXapEURulzSwLDHT2sOHxJic8Un4q9ezOCp0bCPicUwvbX7uJKQooSVZFSF3kg6S6dPFSTTvWxnrLI7T2ArSulzQt0lwib3EJIw1bK4/0');
INSERT INTO `virtual_user` VALUES ('94', 'v1_83f2f3d0ff964f9276eb8e5db2df7697d319555591b027fe95f67916180f65c778729b92a2ab2e2fc54dcc5a4dd16790@stranger', '酷哥:', 'http://wx.qlogo.cn/mmhead/ver_1/3366nKS7u7PggvKuUkicvYLQ76VuzFu3F3ICIlsP2lzp7Tjbmn3xRF8BUAGo76k0uPlHbm3mO9qUQtcicAhqFj2gBHpU6BV9tndKicdxicgNVKU/0');
INSERT INTO `virtual_user` VALUES ('95', 'v1_21d47c5174d732e7bcb24420474c7ecd563b7f1d78fd571407dcfad36aaf2be444f125a74d1d417c0b0070b4baf41a64@stranger', '青青河边草', 'http://wx.qlogo.cn/mmhead/ver_1/HtR9EZ7jB1NVoIofucRVVqj9pmC3B37zhiax7bZQgicvM3dCKnhMUrlzUA646s3vFibicZDSTKXvHbdqhUHCkOgttbHBByOWib7S3NK37eDI2UXo/0');
INSERT INTO `virtual_user` VALUES ('96', 'v1_55b39e3ac3552d903fcb7c1ca841939ffb41571aa4abcb21172c406901b235a912ae0fb35dd4447ed82960c93b7f1dd8@stranger', '老李', 'http://wx.qlogo.cn/mmhead/ver_1/5FiaqLc3KJVG6ZmVlQ9N7PfuXd4bCFiaPSY2GMq8lKQYiaSZCcGyZTVXGVKwgdpMldSiacYbc0TlprLyicNlrcS0CoYAme2CT8eV8fd0N2Uhtbs0/0');
INSERT INTO `virtual_user` VALUES ('97', 'v1_633d9199bdaa7d4249b64371ce044106f55eab7dba876cda25b8749bf4cd944586ee10761a941b7326d297b267e49e16@stranger', '发言权', 'http://wx.qlogo.cn/mmhead/ver_1/GMp8MewTQKicyqj1SklWKFuN4D3iclOjiaTj5uwfBlcBm2SnKwKwaWQoAHR1Zu7vl4OsASLl8bYS4NtHWxFVibibsCGYWNfWCTZ0hzTlZM3veZ00/0');
INSERT INTO `virtual_user` VALUES ('98', 'v1_09ba6d1d5649b9efc4ae9e7ceabc44dd15488fe6e56413e1946d226aa65c8d22bbc94291195c5609f4d8c571eb23aa24@stranger', '李海鑫', 'http://wx.qlogo.cn/mmhead/ver_1/mia5oeGyHUriaM6m6asibhbCVaTOYptCghNdayv7IEG80KcNtic9QJVhDaY6NbK3js21vuhA0anAZtdLvDrUsGATfib8VAWAdQpQV3xM13u6yvnM/0');
INSERT INTO `virtual_user` VALUES ('99', 'v1_1e05f891eb12cadfb9cbc493ae4ae8089ac8bd6978903e6ca17bf58d1064100ab67f9fa6d75e634bd30d02acf5f13668@stranger', '雨露', 'http://wx.qlogo.cn/mmhead/ver_1/alFkBTVWCP5MYVFZ0lJWD3YksgVicIgHZrOTY9l9musWkBhZRMFZY1X0GgIBA53AibCdBOxvjTRa5mo62ay5zMx38tz1MPMwATqou16gYg3hU/0');
INSERT INTO `virtual_user` VALUES ('100', 'v1_159d8a7a4e9c68d07c430f866a6bb332b72a2ee4d77a476eb687bd47d2f82a881d8a63c10d28eee7ec88a32cef64cc14@stranger', '尊香', 'http://wx.qlogo.cn/mmhead/ver_1/HBEwOU6KGMz9icrkKlGRh7BFW9117QmaY8X95C7frAKRj6dPJgYqEicvaueExushutznMxFO587O6fvUeOBwN5SUtIru8zNorwnvsgtSN4EA0/0');
INSERT INTO `virtual_user` VALUES ('101', 'v1_49d4e3782728e37c2412650d03f5f00e646d6f219f180245bfbd2c44a30f532609c6f65cde588c06d9358289aeae21be@stranger', '安静和快了', 'http://wx.qlogo.cn/mmhead/ver_1/UZMSQ6paUFBkPzCIQ81hWria944cos5X0IicWlvXlXEQ4gdhOG73nvYYMiagutIiaOd0UApLev2eyHl7TSexrXQlmXLlz41uPiaFZJ6iaIU4PZJtA/0');
INSERT INTO `virtual_user` VALUES ('102', 'v1_e6d2303253fcd7d8f379888fb38da191505f3470f2bcc4c37180142971a9f39336222c142419befa78921639d12e8929@stranger', '私亨茶园', 'http://wx.qlogo.cn/mmhead/ver_1/iaLXibntibicNAicHLHHe1yrxlUv1TwBSJRZ9gOR0UcK1kCjWdPxBpicicZxwZicysJyxTalwTe2D3QPuo7Q857BjPaTCfEAnFYw4bTYCibx1pSJPjgs/0');
INSERT INTO `virtual_user` VALUES ('103', 'v1_aa8d2d160e7c3bce6951868d4a163d39caa9b116d030373a4db7bd5a8f17ef303d2462f0235e7aeba444fa71717ef075@stranger', '薇', 'http://wx.qlogo.cn/mmhead/ver_1/gibEDZfXb46rxbONN35iaRHNoKAR3udqFick0m7ibgQSxAPwWjg2Eqkcyc3FryYAHJGfTgwUhiaNe6JW6bHUAThYLvH1Kc552Mo0OZeHQCueKlWo/0');
INSERT INTO `virtual_user` VALUES ('104', 'v1_eb659995fbb9b051f7d31a34a16cf3acb9bddb829bf8ff1b1aedb0ec5d1d0899@stranger', '国恒工贸', 'http://wx.qlogo.cn/mmhead/ver_1/kDicQGrehuawTpE3kLKic3ULEsROb9C88fCick6VQgSeT7YA5QVZia3nCewTAckhy3bnaFuibgOoHkxKDztEWUewEM9DdjMuErbPcIBgrqb0l2co/0');
INSERT INTO `virtual_user` VALUES ('105', 'tian152051', '田亚军', 'http://wx.qlogo.cn/mmhead/ver_1/40UBHicLNfeyk8gibia80TZqhlaOSqQibo3eZcyIzR69OLkGCZHAcZ2ybhicSQaayabqiaicXPsE9Qicmuy4OSE0IsjPQxib1aCgBJPVibqH9F4DGibzaw/0');
INSERT INTO `virtual_user` VALUES ('106', 'liujianhua221063', '桃花源', 'http://wx.qlogo.cn/mmhead/ver_1/ibgRRvJQpBm36M7YokLF6ibVs1fnEVJ3f00kanb11b5ad17hv6EpsOF5kiaeDk6euSKJe5XZugXwol6Mfohrf8KFUWibgU8OT9cKCrMDYQ046T8/0');
INSERT INTO `virtual_user` VALUES ('107', 'gh_f8d22bca9822', '星禅佛', 'http://wx.qlogo.cn/mmhead/ver_1/goO0a4KtZ1WY7vxzT0Quknic11MA5cpLf0QvZJ4kc7V49COjnHfKmiaMz3EWpMsM6QsMiaU3CO29ogR7icSzibNH84WobluMCKmhydfsC3JsFcZU/0');
INSERT INTO `virtual_user` VALUES ('108', 'gh_fedeb4485f01', '婴幼儿教育', 'http://wx.qlogo.cn/mmhead/ver_1/oOgZ4X21iaOCPOswrR63F6gEqD06hpKLMzK594oG3yHwoYpvoxTrZNSvDTKiavfvSMiaOeX8DngOMayZFp3W05ryKmO3VSw5XzYLxgJ5miaWhaY/0');
INSERT INTO `virtual_user` VALUES ('109', 'wxid_79kyvkenyw8b22', '顺其自然', 'http://wx.qlogo.cn/mmhead/ver_1/ntS0kLibmsRFPsez7dJJXRvjwreD7HPPMDsFdTAoGLBT5W5NWFcEuOqe9bIwUJ7tdFcibTJ2kZ9EHibK33fKtZqqgWQtPjAK6xpEPAMzKCJ1eM/0');
