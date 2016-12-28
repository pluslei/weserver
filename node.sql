/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50520
Source Host           : localhost:3306
Source Database       : haolive

Target Server Type    : MYSQL
Target Server Version : 50520
File Encoding         : 65001

Date: 2016-12-28 11:03:07
*/

SET FOREIGN_KEY_CHECKS=0;

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
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8;

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
INSERT INTO `node` VALUES ('33', '主题设置', 'theme/index', '2', '30', '系统设置/主题设置', '2', '1', '50', 'weserver/sysconfig/theme_index', '1', '');
INSERT INTO `node` VALUES ('34', '数据管理', 'data', '1', '0', '数据管理', '2', '1', '600', 'weserver/data', '1', 'am-icon-database');
INSERT INTO `node` VALUES ('35', '房间管理', 'room/index', '2', '34', '数据管理/房间管理', '2', '1', '50', 'weserver/data/room_index', '1', '');
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
INSERT INTO `node` VALUES ('46', '刷新房间', 'room/refresh', '3', '35', '房间管理/刷新房间', '2', '1', '50', 'weserver/data/room_refresh', '1', '');
INSERT INTO `node` VALUES ('47', '在线用户', 'user/onlineuser', '2', '1', '用户管理/在线用户', '2', '1', '102', 'weserver/user/onlineuser', '1', '');
INSERT INTO `node` VALUES ('48', '首页管理', 'home', '1', '0', '首页管理', '2', '1', '700', 'weserver/home', '1', 'am-icon-home');
INSERT INTO `node` VALUES ('49', '关于我们', 'home/aboutme', '2', '48', '首页管理/关于我们', '2', '1', '50', 'weserver/home/aboutme', '1', '');
INSERT INTO `node` VALUES ('50', '联系我们', 'home/contact', '2', '48', '首页管理/联系我们', '2', '1', '50', 'weserver/home/contact', '1', '');
INSERT INTO `node` VALUES ('51', '讲师简介', 'teacher/index', '2', '34', '数据管理/讲师简介', '2', '1', '50', 'weserver/data/teacher_index', '1', '');
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
INSERT INTO `node` VALUES ('69', '问题解答', 'data/qs_index', '2', '34', '数据管理/问题解答', '2', '1', '50', 'weserver/data/qs_index', '1', '');
INSERT INTO `node` VALUES ('70', '问题增加', 'data/qs_addqs', '3', '69', '数据管理/问题增加', '2', '1', '50', 'weserver/data/qs_addqs', '1', '');
INSERT INTO `node` VALUES ('71', '更新问题', 'data/qs_updateqs', '3', '69', '数据管理/更新问题', '2', '1', '50', 'weserver/data/qs_updateqs', '1', '');
INSERT INTO `node` VALUES ('72', '删除问题', 'data/qs_delqs', '3', '69', '数据管理/删除问题', '2', '1', '50', 'weserver/data/qs_delqs', '1', '');
INSERT INTO `node` VALUES ('73', '发送广播', 'data/qs_broad', '2', '34', '数据管理/qs_broad', '2', '1', '50', 'weserver/data/qs_broad', '1', '');
INSERT INTO `node` VALUES ('74', '前端功能', 'tool', '1', '0', '前端功能', '2', '2', '50', 'tool', '1', '');
INSERT INTO `node` VALUES ('75', '是否上传聊天图片', 'uploadchatimage', '2', '74', '前端功能/上传聊天图片', '2', '2', '50', 'tool/uploadmsgimage', '1', '');
INSERT INTO `node` VALUES ('76', '是否私聊', 'provicechat', '2', '74', '前端功能/是否私聊', '2', '2', '50', 'tool/provicechat', '1', '');
INSERT INTO `node` VALUES ('77', '是否发送广播', 'sendbroadcast', '2', '74', '前端功能/是否发送广播', '2', '2', '50', 'tool/sendbrodcast', '1', '');
INSERT INTO `node` VALUES ('78', '用户审核', 'user/verifyuser', '2', '1', '用户管理/用户审核', '2', '1', '50', 'weserver/user/verifyuser', '0', '');
INSERT INTO `node` VALUES ('79', '机器人', 'data/robot_speak', '2', '34', '数据管理/机器人', '2', '1', '50', 'weserver/data/robot_speak', '1', '');
INSERT INTO `node` VALUES ('80', '消息库管理', 'data/message_index', '2', '34', '数据管理/消息库管理', '2', '1', '50', 'weserver/data/message_index', '1', '');
INSERT INTO `node` VALUES ('81', '删除消息库', 'data/message_delete', '3', '79', '数据管理/删除消息库', '2', '1', '50', 'weserver/data/message_delete', '1', '');
INSERT INTO `node` VALUES ('82', '修改消息库', 'data/message_edit', '3', '79', '数据管理/修改消息库', '2', '1', '50', 'weserver/data/message_edit', '1', '');
INSERT INTO `node` VALUES ('83', '增加消息库', 'data/message_add', '3', '79', '数据管理/增加消息库', '2', '1', '50', 'weserver/data/messgae_add', '1', '');
INSERT INTO `node` VALUES ('84', '修改消息库分类', 'data/messagetype_edit', '3', '79', '数据管理/修改消息库分类', '2', '1', '50', 'weserver/data/messagetype_edit', '1', '');
INSERT INTO `node` VALUES ('85', '增加消息库分类', 'data/messagetype_add', '3', '79', '数据管理/增加消息库分类', '2', '1', '50', 'weserver/data/messagetype_add', '1', '');
INSERT INTO `node` VALUES ('86', '删除消息库分类', 'data/messagetype_delete', '3', '79', '数据管理/删除消息库分类', '2', '1', '50', 'weserver/data/messagetype_delete', '1', '');
INSERT INTO `node` VALUES ('98', '飞屏', 'flyscreen', '2', '36', '前端功能/飞屏', '2', '2', '50', 'flyscreen', '1', '');
