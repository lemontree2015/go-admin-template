/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : localhost
 Source Database       : quick_backend

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : utf-8

 Date: 06/26/2021 16:39:05 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `sys_menu`
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) DEFAULT '0' COMMENT '父级id',
  `name` varchar(36) DEFAULT '' COMMENT '资源名',
  `icon` varchar(128) DEFAULT NULL COMMENT '资源图标',
  `path` varchar(128) DEFAULT NULL COMMENT '后台访问路由',
  `action` varchar(100) DEFAULT '' COMMENT 'api路由',
  `component` varchar(128) DEFAULT NULL COMMENT 'Vue组件',
  `sort` int(10) DEFAULT '0',
  `menu_type` tinyint(4) DEFAULT '1' COMMENT '资源类型 1菜单 2接口',
  `method` char(10) DEFAULT '' COMMENT '接口请求方式：GET POST PUT DELETE',
  `permission` varchar(128) DEFAULT NULL COMMENT '权限标识',
  `status` tinyint(2) DEFAULT '1' COMMENT '状态1正常 0删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `sys_menu`
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` VALUES ('1', '0', '系统管理', 'example', '/system', '', 'Layout', '0', '1', '', null, '1'), ('2', '1', '用户管理', 'user', 'user', '', '/user/index', '0', '1', '', null, '1'), ('3', '1', '角色管理', 'peoples', 'role', '', '/role/index', '0', '1', '', '', '1'), ('4', '1', '菜单管理', 'tree-table', 'menu', '', '/menu/index', '0', '1', '', null, '1'), ('5', '2', '添加用户', 'edit', '', '/user/add', '', '0', '2', 'POST', 'system:user:add', '1'), ('6', '0', '内容管理', 'documentation', '/content', '', 'Layout', '0', '1', 'GET', '', '1'), ('7', '6', '新闻管理', 'documentation', 'news', '', '/news/index', '0', '1', 'GET', '', '1'), ('8', '7', '分类管理', 'list', 'category', '', '/category/index', '0', '1', 'GET', '', '1'), ('9', '7', '文章管理', 'documentation', 'publish', '', '/news/list', '0', '1', 'GET', '', '1'), ('10', '2', '用户详情', 'dict', '', '/user/get/:id', '', '0', '2', 'GET', 'system:user:get', '1'), ('11', '2', '用户列表', 'list', '', '/user/list', '', '0', '2', 'GET', 'system:user:list', '1'), ('12', '2', '编辑用户', 'edit', '', '/user/edit/:id', '', '0', '2', 'PUT', 'system:user:edit', '1'), ('13', '3', '添加角色', 'edit', '', '/role/add', '', '0', '2', 'POST', 'system:role:add', '1'), ('14', '3', '角色详情', 'dict', '', '/role/get/:id', '', '0', '2', 'GET', 'system:role:get', '1'), ('15', '3', '角色列表', 'list', '', '/role/list', '', '0', '2', 'GET', 'system:role:list', '1'), ('16', '3', '角色菜单', 'dict', '', '/role/roleMenus/:roleId', '', '0', '2', 'GET', 'sys:role:roleMenu', '1'), ('17', '3', '所有角色', 'list', '', '/role/all', '', '0', '2', 'GET', 'sys:role:all', '1'), ('18', '3', '编辑角色', 'edit', '', '/role/edit/:id', '', '0', '2', 'POST', 'system:role:edit', '1'), ('19', '4', '菜单列表', 'list', '', '/menu/list', '', '0', '2', 'GET', 'system:menu:list', '1'), ('20', '4', '添加菜单', 'edit', '', '/menu/add', '', '0', '2', 'GET', 'system:menu:add', '1'), ('21', '4', '菜单详情', 'dict', '', '/menu/get/:id', '', '0', '2', 'GET', 'sys:menu:get', '1'), ('22', '4', '编辑菜单', 'edit', '', '/menu/edit/:id', '', '0', '2', 'PUT', 'sys:menu:edit', '1'), ('23', '1', '网站设置', 'system', 'setting', '', '/setting/index', '0', '1', 'GET', '', '1'), ('24', '23', '配制列表', 'list', 'list', '', '/setting/list', '0', '1', 'GET', '', '1'), ('25', '23', '编辑配制', 'edit', 'edit', '', '/setting/edit', '0', '1', 'GET', '', '1'), ('26', '7', '编辑文章', 'edit', 'edit', '', '/news/edit', '0', '1', 'GET', '', '1'), ('27', '6', '轮播广告', 'documentation', 'adMaterials', '', '/adMaterials/index', '0', '1', 'GET', '', '1'), ('28', '27', '添加广告', 'edit', '', '/material/add', '', '0', '2', 'POST', 'system:material:add', '1'), ('29', '27', '广告列表', 'list', '', '/material/list', '', '0', '2', 'GET', 'system:material:list', '1'), ('30', '27', '广告详情', 'dict', '', '/material/get/:id', '', '0', '2', 'GET', 'system:material:get', '1'), ('31', '27', '编辑广告', 'edit', '', '/material/edit/:id', '', '0', '2', 'PUT', 'system:material:edit', '1'), ('32', '8', '分类列表', 'list', '', '/news/category/list', '', '0', '2', 'GET', 'system:category:list', '1'), ('33', '8', '所有分类', 'list', '', '/news/category/all', '', '0', '2', 'GET', 'system:category:all', '1'), ('34', '8', '添加分类', 'edit', '', '/news/category/add', '', '0', '2', 'POST', 'system:category:add', '1'), ('35', '8', '分类详情', 'dict', '', '/news/category/get/:id', '', '0', '2', 'GET', 'system:category:get', '1'), ('36', '8', '编辑分类', 'edit', '', '/news/category/edit/:id', '', '0', '2', 'PUT', 'system:category:edit', '1'), ('37', '9', '文章列表', 'list', '', '/news/publish/list', '', '0', '2', 'GET', 'system:publish:list', '1'), ('38', '26', '添加文章', 'edit', '', '/news/publish/add', '', '0', '2', 'POST', 'system:publish:add', '1'), ('39', '26', '文章详情', 'dict', '', '/news/publish/get/:id', '', '0', '2', 'GET', 'system:publish:get', '1'), ('40', '26', '修改文章', 'edit', '', '/news/publish/edit/:id', '', '0', '2', 'PUT', 'system:publish:edit', '1'), ('41', '26', '上传图片', 'upload', '', '/news/publish/uploadImg', '', '0', '2', 'POST', 'system:publish:uploadImg', '1'), ('42', '23', '添加配制', 'edit', '', '/setting/add', '', '0', '2', 'POST', 'system:setting:add', '1'), ('43', '23', '配制详情', 'dict', '', '/setting/get/:id', '', '0', '2', 'GET', 'system:setting:get', '1');
COMMIT;

-- ----------------------------
--  Table structure for `sys_role`
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '角色标识',
  `title` varchar(50) DEFAULT NULL COMMENT '角色名',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(4) DEFAULT '1' COMMENT '记录状态（1正常 0删除）',
  `created` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `sys_role`
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES ('1', 'admin', '超级管理员', '超级管理员', '1', '2021-06-10 10:34:00', '2021-06-10 10:34:00'), ('2', 'audit', '运营', '运营', '1', '2021-06-11 11:11:25', '2021-06-20 21:09:29');
COMMIT;

-- ----------------------------
--  Table structure for `sys_role_menu`
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `role_id` int(10) DEFAULT NULL,
  `menu_id` int(10) DEFAULT NULL,
  `status` tinyint(2) DEFAULT '1' COMMENT '1正常 0删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `sys_role_menu`
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` VALUES ('1', '2', '1', '0'), ('2', '2', '2', '0'), ('3', '2', '3', '0'), ('4', '2', '4', '0'), ('5', '2', '5', '0'), ('6', '2', '10', '0'), ('7', '2', '11', '0'), ('8', '2', '12', '0'), ('9', '2', '13', '0'), ('10', '2', '14', '0'), ('11', '2', '15', '0'), ('12', '2', '16', '0'), ('13', '2', '17', '0'), ('14', '2', '18', '0'), ('15', '2', '19', '0'), ('16', '2', '20', '0'), ('17', '2', '21', '0'), ('18', '2', '22', '0'), ('19', '2', '23', '0'), ('20', '2', '6', '1'), ('21', '2', '7', '1'), ('22', '2', '8', '1'), ('23', '2', '9', '1'), ('24', '2', '32', '1'), ('25', '2', '33', '1'), ('26', '2', '34', '1'), ('27', '2', '35', '1'), ('28', '2', '36', '0'), ('29', '2', '37', '1'), ('30', '2', '26', '1'), ('31', '2', '38', '1'), ('32', '2', '39', '1'), ('33', '2', '40', '1'), ('34', '2', '41', '1'), ('35', '2', '27', '1'), ('36', '2', '28', '1'), ('37', '2', '29', '1'), ('38', '2', '30', '1'), ('39', '2', '31', '1');
COMMIT;

-- ----------------------------
--  Table structure for `sys_settings`
-- ----------------------------
DROP TABLE IF EXISTS `sys_settings`;
CREATE TABLE `sys_settings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '标识符',
  `title` varchar(50) DEFAULT NULL COMMENT '配制名',
  `content` text,
  `status` tinyint(2) DEFAULT '1' COMMENT '状态1正常0删除',
  `created` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `sys_settings`
-- ----------------------------
BEGIN;
INSERT INTO `sys_settings` VALUES ('1', 'aboutUs', '关于我们', '关于我们。。。。。', '1', '2021-06-11 16:00:26', '2021-06-11 16:04:12'), ('2', 'concactUs', '联系我们', '<p>联系我们</p>', '1', '2021-06-11 23:02:13', '2021-06-11 23:02:13');
COMMIT;

-- ----------------------------
--  Table structure for `sys_user`
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '',
  `real_name` varchar(64) NOT NULL DEFAULT '',
  `password` varchar(40) NOT NULL DEFAULT '',
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `avatar` varchar(200) DEFAULT NULL COMMENT '头像',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1正常 0删除',
  `last_login_time` datetime DEFAULT NULL,
  `last_login_ip` varchar(50) DEFAULT NULL,
  `created` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `sys_user`
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES ('1', 'admin', '梁胖', '7c4a8d09ca3762af61e59520943dc26494f8941b', '415593463@qq.com', '13817056806', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '1', '2021-06-22 11:22:09', '127.0.0.1', '2021-06-10 10:32:28', '2021-06-20 22:32:23'), ('2', 'test', '测试', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', '123456789@qq.com', '1234567', '', '1', '2021-06-20 21:09:52', '127.0.0.1', '2021-06-19 21:26:19', '2021-06-19 21:26:19');
COMMIT;

-- ----------------------------
--  Table structure for `sys_user_role`
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL DEFAULT '0',
  `role_id` int(10) NOT NULL DEFAULT '0',
  `status` tinyint(2) DEFAULT '1' COMMENT '记录状态（1正常 0删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `sys_user_role`
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` VALUES ('1', '1', '1', '1'), ('2', '2', '2', '1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
