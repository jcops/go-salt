/*
Navicat MySQL Data Transfer

Source Server         : 云
Source Server Version : 50717
Source Host           : 118.25.39.84:3306
Source Database       : go-salt

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2019-05-06 14:04:28
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for go_menu
-- ----------------------------
DROP TABLE IF EXISTS `go_menu`;
CREATE TABLE "go_menu" (
  "id" int(11) NOT NULL AUTO_INCREMENT,
  "created_on" int(11) DEFAULT NULL,
  "modified_on" int(11) DEFAULT NULL,
  "deleted_on" int(11) DEFAULT NULL,
  "name" varchar(255) DEFAULT NULL,
  "path" varchar(255) DEFAULT NULL,
  "method" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("id")
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_menu
-- ----------------------------
INSERT INTO `go_menu` VALUES ('1', null, null, '0', '查询所有用户', '/api/v1/users/*', 'GET');
INSERT INTO `go_menu` VALUES ('2', null, null, '0', '注册用户', '/api/v1/register', 'POST');
INSERT INTO `go_menu` VALUES ('3', null, null, '0', '查询单个用户', '/api/v1/users/:id', 'GET');
INSERT INTO `go_menu` VALUES ('4', null, null, '0', '删除用户', '/api/v1/users/:id', 'DELETE');
INSERT INTO `go_menu` VALUES ('5', null, null, '0', '修改用户信息', '/api/v1/users', 'PUT');
INSERT INTO `go_menu` VALUES ('6', null, null, '0', '查询所有角色', '/api/v1/role/*', 'GET');
INSERT INTO `go_menu` VALUES ('7', null, null, '0', '添加角色', '/api/v1/role', 'POST');
INSERT INTO `go_menu` VALUES ('8', null, null, '0', '删除角色', '/api/v1/role/:id', 'DELETE');
INSERT INTO `go_menu` VALUES ('9', null, null, '0', '修改角色', '/api/v1/role', 'PUT');
INSERT INTO `go_menu` VALUES ('10', null, null, '0', '查询所有节点', '/api/v1/keylist', 'GET');
INSERT INTO `go_menu` VALUES ('11', null, null, '0', '接受节点认证请求', '/api/v1/keyaccept/:id', 'GET');
INSERT INTO `go_menu` VALUES ('12', null, null, '0', '删除认证节点', '/api/v1/keydelete/:id', 'GET');
INSERT INTO `go_menu` VALUES ('13', null, null, '0', '部署软件', '/api/v1/keydeploy', 'GET');
INSERT INTO `go_menu` VALUES ('14', null, null, '0', '测试节点连接性', '/api/v1/keyping', 'GET');
INSERT INTO `go_menu` VALUES ('15', null, null, '0', '执行命令', '/api/v1/keyrun', 'POST');
INSERT INTO `go_menu` VALUES ('16', null, null, '0', '拷贝文件到节点', '/api/v1/keycopy', 'GET');
INSERT INTO `go_menu` VALUES ('17', null, null, '0', '查询所有菜单', '/api/v1/menu', 'GET');
INSERT INTO `go_menu` VALUES ('18', null, null, '0', '创建单个菜单', '/api/v1/menu', 'POST');
INSERT INTO `go_menu` VALUES ('19', null, null, '0', '更新单个菜单', '/api/v1/menu/:id', 'PUT');
INSERT INTO `go_menu` VALUES ('20', null, null, '0', '删除单个菜单', '/api/v1/menu/:id', 'DELETE');
INSERT INTO `go_menu` VALUES ('21', null, '1555666367', '0', '修改用户密码', '/api/v1/users', 'POST');
INSERT INTO `go_menu` VALUES ('23', null, null, '0', '查看用户详情', '/api/v1/getinfo', 'GET');

-- ----------------------------
-- Table structure for go_role
-- ----------------------------
DROP TABLE IF EXISTS `go_role`;
CREATE TABLE "go_role" (
  "id" int(11) NOT NULL AUTO_INCREMENT,
  "created_on" int(11) DEFAULT NULL,
  "modified_on" int(11) DEFAULT NULL,
  "deleted_on" int(11) DEFAULT NULL,
  "name" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("id")
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_role
-- ----------------------------
INSERT INTO `go_role` VALUES ('29', '1555584565', '1555584565', '1555584653', 'sa1');
INSERT INTO `go_role` VALUES ('30', '1555584631', '1555584631', '0', 'sa2');
INSERT INTO `go_role` VALUES ('31', '1555587422', '1555587422', '0', 'dev');
INSERT INTO `go_role` VALUES ('32', '1555587496', '1555587496', '0', 'dev1');
INSERT INTO `go_role` VALUES ('33', '1555587609', '1555587609', '0', 'dev2');
INSERT INTO `go_role` VALUES ('34', '1555587673', '1555587673', '0', 'dev3');
INSERT INTO `go_role` VALUES ('35', '1555587843', '1555587843', '0', 'dev4');
INSERT INTO `go_role` VALUES ('36', '1555665158', '1555665158', '1556270131', 'test111');
INSERT INTO `go_role` VALUES ('37', '1555665272', '1555665272', '1555665747', 'test11122');

-- ----------------------------
-- Table structure for go_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `go_role_menu`;
CREATE TABLE "go_role_menu" (
  "role_id" int(11) NOT NULL,
  "menu_id" int(11) NOT NULL,
  PRIMARY KEY ("role_id","menu_id")
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_role_menu
-- ----------------------------
INSERT INTO `go_role_menu` VALUES ('30', '1');
INSERT INTO `go_role_menu` VALUES ('30', '2');
INSERT INTO `go_role_menu` VALUES ('30', '3');
INSERT INTO `go_role_menu` VALUES ('30', '4');
INSERT INTO `go_role_menu` VALUES ('30', '9');
INSERT INTO `go_role_menu` VALUES ('30', '11');
INSERT INTO `go_role_menu` VALUES ('30', '23');
INSERT INTO `go_role_menu` VALUES ('32', '1');
INSERT INTO `go_role_menu` VALUES ('32', '2');
INSERT INTO `go_role_menu` VALUES ('32', '3');
INSERT INTO `go_role_menu` VALUES ('32', '4');
INSERT INTO `go_role_menu` VALUES ('32', '5');
INSERT INTO `go_role_menu` VALUES ('32', '6');
INSERT INTO `go_role_menu` VALUES ('33', '1');
INSERT INTO `go_role_menu` VALUES ('33', '2');
INSERT INTO `go_role_menu` VALUES ('33', '3');
INSERT INTO `go_role_menu` VALUES ('33', '4');
INSERT INTO `go_role_menu` VALUES ('33', '5');
INSERT INTO `go_role_menu` VALUES ('33', '6');
INSERT INTO `go_role_menu` VALUES ('34', '1');
INSERT INTO `go_role_menu` VALUES ('34', '2');
INSERT INTO `go_role_menu` VALUES ('34', '3');
INSERT INTO `go_role_menu` VALUES ('34', '4');
INSERT INTO `go_role_menu` VALUES ('34', '5');
INSERT INTO `go_role_menu` VALUES ('34', '6');
INSERT INTO `go_role_menu` VALUES ('34', '7');
INSERT INTO `go_role_menu` VALUES ('35', '1');
INSERT INTO `go_role_menu` VALUES ('35', '2');
INSERT INTO `go_role_menu` VALUES ('35', '3');
INSERT INTO `go_role_menu` VALUES ('35', '4');
INSERT INTO `go_role_menu` VALUES ('35', '5');
INSERT INTO `go_role_menu` VALUES ('35', '6');
INSERT INTO `go_role_menu` VALUES ('35', '7');
INSERT INTO `go_role_menu` VALUES ('35', '8');
INSERT INTO `go_role_menu` VALUES ('35', '9');
INSERT INTO `go_role_menu` VALUES ('35', '10');
INSERT INTO `go_role_menu` VALUES ('35', '11');
INSERT INTO `go_role_menu` VALUES ('35', '12');
INSERT INTO `go_role_menu` VALUES ('35', '13');
INSERT INTO `go_role_menu` VALUES ('35', '14');
INSERT INTO `go_role_menu` VALUES ('35', '15');
INSERT INTO `go_role_menu` VALUES ('35', '16');
INSERT INTO `go_role_menu` VALUES ('35', '17');
INSERT INTO `go_role_menu` VALUES ('35', '18');
INSERT INTO `go_role_menu` VALUES ('35', '19');
INSERT INTO `go_role_menu` VALUES ('35', '20');
INSERT INTO `go_role_menu` VALUES ('35', '21');
INSERT INTO `go_role_menu` VALUES ('35', '23');

-- ----------------------------
-- Table structure for go_user
-- ----------------------------
DROP TABLE IF EXISTS `go_user`;
CREATE TABLE "go_user" (
  "id" int(11) NOT NULL AUTO_INCREMENT,
  "created_on" int(11) DEFAULT NULL,
  "modified_on" int(11) DEFAULT NULL,
  "deleted_on" int(11) DEFAULT NULL,
  "username" varchar(255) DEFAULT NULL,
  "password" varchar(255) DEFAULT NULL,
  PRIMARY KEY ("id")
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_user
-- ----------------------------
INSERT INTO `go_user` VALUES ('31', '1555583521', '1555583521', '0', 'eric', 'e10adc3949ba59abbe56e057f20f883e');
INSERT INTO `go_user` VALUES ('32', '1555583540', '1555583540', '0', 'admin', 'e10adc3949ba59abbe56e057f20f883e');
INSERT INTO `go_user` VALUES ('33', '1555585720', '1555586199', '0', 'ada', 'd41d8cd98f00b204e9800998ecf8427e');
INSERT INTO `go_user` VALUES ('35', '1555586467', '1555660790', '0', 'ada1', 'd41d8cd98f00b204e9800998ecf8427e');
INSERT INTO `go_user` VALUES ('36', '1555586548', '1555586548', '0', 'ada2', 'e10adc3949ba59abbe56e057f20f883e');
INSERT INTO `go_user` VALUES ('37', '1555659720', '1555664165', '0', 'ada11', '4297f44b13955235245b2497399d7a93');
INSERT INTO `go_user` VALUES ('38', null, null, '1556176951', 'dad1', '4297f44b13955235245b2497399d7a93');
INSERT INTO `go_user` VALUES ('39', '1556174841', '1556174841', '1556176943', 'testa', 'f5bb0c8de146c67b44babbf4e6584cc0');
INSERT INTO `go_user` VALUES ('40', '1556174865', '1556174865', '1556176323', 'test1111111', '0e9a4f4ad6d22d135835e1050bae60f7');
INSERT INTO `go_user` VALUES ('41', '1556247883', '1556247883', '0', 'tttttttt', '5abbb6bc10fa486549b7697c9cffcf61');
INSERT INTO `go_user` VALUES ('42', '1556263163', '1556263163', '0', '1111111', 'bbb8aae57c104cda40c93843ad5e6db8');
INSERT INTO `go_user` VALUES ('43', '1556263169', '1556263169', '0', '2222222', '4d18e2c96bb0f39c6d6dc690542b0bdc');
INSERT INTO `go_user` VALUES ('44', '1556263174', '1556263174', '0', '3333333', '4aee3e28df37ea1af64bd636eca59dcb');
INSERT INTO `go_user` VALUES ('45', '1556263183', '1556263183', '0', '44444', 'b857eed5c9405c1f2b98048aae506792');
INSERT INTO `go_user` VALUES ('46', '1556263188', '1556263188', '0', '5555555', '9079b6ee1d5ca04ab00e44e877a222ee');
INSERT INTO `go_user` VALUES ('47', '1556263201', '1556263201', '0', '1211111', 'ba9485f02fc98cdbd2edadb0aa8f6390');
INSERT INTO `go_user` VALUES ('48', '1556263210', '1556263210', '0', '131313131', '1cf34d60ff9390a97bd9a4d487d05df6');
INSERT INTO `go_user` VALUES ('49', '1556264139', '1556264139', '0', 'sssssssssss', '3c1d4baa14ad36f3e5bfb6598caa3995');
INSERT INTO `go_user` VALUES ('50', '1556264148', '1556264148', '0', 'fffffffff', '5da0aebaeea0108d794303443b2490f7');

-- ----------------------------
-- Table structure for go_user_role
-- ----------------------------
DROP TABLE IF EXISTS `go_user_role`;
CREATE TABLE "go_user_role" (
  "user_id" int(11) NOT NULL,
  "role_id" int(11) NOT NULL,
  PRIMARY KEY ("user_id","role_id")
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of go_user_role
-- ----------------------------
INSERT INTO `go_user_role` VALUES ('31', '30');
INSERT INTO `go_user_role` VALUES ('31', '31');
INSERT INTO `go_user_role` VALUES ('31', '32');
INSERT INTO `go_user_role` VALUES ('31', '33');
INSERT INTO `go_user_role` VALUES ('31', '34');
INSERT INTO `go_user_role` VALUES ('31', '35');
INSERT INTO `go_user_role` VALUES ('31', '36');
INSERT INTO `go_user_role` VALUES ('33', '35');
INSERT INTO `go_user_role` VALUES ('33', '36');
INSERT INTO `go_user_role` VALUES ('35', '30');
INSERT INTO `go_user_role` VALUES ('36', '30');
INSERT INTO `go_user_role` VALUES ('37', '30');
INSERT INTO `go_user_role` VALUES ('37', '31');
INSERT INTO `go_user_role` VALUES ('37', '35');
INSERT INTO `go_user_role` VALUES ('37', '36');
INSERT INTO `go_user_role` VALUES ('41', '30');
INSERT INTO `go_user_role` VALUES ('41', '31');
INSERT INTO `go_user_role` VALUES ('41', '35');
INSERT INTO `go_user_role` VALUES ('41', '36');
