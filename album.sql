/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : go

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2018-07-02 00:00:51
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for album
-- ----------------------------
DROP TABLE IF EXISTS `album`;
CREATE TABLE `album` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `picture` varchar(255) DEFAULT '' COMMENT 'Picture',
  `keywords` varchar(2550) DEFAULT '' COMMENT '关键词',
  `summary` varchar(255) DEFAULT '',
  `created` int(10) DEFAULT '0' COMMENT '发布时间',
  `viewnum` int(10) DEFAULT '0' COMMENT '阅读次数',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态: 0草稿，1已发布',
  PRIMARY KEY (`id`),
  KEY `INDEX_TCVS` (`title`,`created`,`viewnum`,`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='相册';

-- ----------------------------
-- Records of album
-- ----------------------------

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `uri` varchar(255) DEFAULT '' COMMENT 'URL',
  `keywords` varchar(2550) DEFAULT '' COMMENT '关键词',
  `summary` varchar(255) DEFAULT '',
  `content` longtext NOT NULL COMMENT '正文',
  `author` varchar(20) DEFAULT '' COMMENT '作者',
  `created` int(10) DEFAULT '0' COMMENT '发布时间',
  `viewnum` int(10) DEFAULT '0' COMMENT '阅读次数',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态: 0草稿，1已发布',
  `ctm` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `INDEX_TCVS` (`title`,`created`,`viewnum`,`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='文章';

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('1', '测试标题的', '', '测试文字名称', '', '这个是内容里面的内容什么贵什么鬼什么贵', '老李', '0', '0', '1', '2018-06-06 16:40:19', '5');
INSERT INTO `article` VALUES ('2', '妇女节', '', '富基恩', '', '阿斯顿发为阿斯顿发的撒范德萨发的', '老发', '0', '0', '1', '2018-06-06 16:40:13', '4');
INSERT INTO `article` VALUES ('3', '五一节', '', '爱妃', '', '瓦尔特热热同仁堂', '老张', '0', '0', '1', '2018-06-06 16:40:14', '5');
INSERT INTO `article` VALUES ('4', '五一节', '', '水电费', '', '为二位二位二位', '', '0', '0', '1', '2018-06-06 16:40:17', '4');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `article_id` int(10) DEFAULT NULL,
  `nickname` varchar(15) DEFAULT NULL,
  `uri` varchar(255) DEFAULT NULL,
  `content` text,
  `created` int(10) DEFAULT '0',
  `status` tinyint(1) DEFAULT '1' COMMENT '0屏蔽，1正常',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='blog评论';

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for f
-- ----------------------------
DROP TABLE IF EXISTS `f`;
CREATE TABLE `f` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dkd` varchar(11) DEFAULT NULL,
  `qwqqq` varchar(22) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of f
-- ----------------------------
INSERT INTO `f` VALUES ('1', 'asdf', 'asdfasdf');
INSERT INTO `f` VALUES ('2', 'werwer', 'werwer');

-- ----------------------------
-- Table structure for t
-- ----------------------------
DROP TABLE IF EXISTS `t`;
CREATE TABLE `t` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of t
-- ----------------------------
INSERT INTO `t` VALUES ('1', 'asdfasdf');
INSERT INTO `t` VALUES ('2', 'erter');
INSERT INTO `t` VALUES ('3', 'asdf');
INSERT INTO `t` VALUES ('4', 'werwer');

-- ----------------------------
-- Table structure for test_api
-- ----------------------------
DROP TABLE IF EXISTS `test_api`;
CREATE TABLE `test_api` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_name` varchar(50) NOT NULL COMMENT 'api名称',
  `api_param` text CHARACTER SET latin1 NOT NULL COMMENT 'api请求参数json格式',
  `test_rule_id` char(250) CHARACTER SET latin1 NOT NULL COMMENT '测试规则编号',
  `api_url` varchar(50) CHARACTER SET latin1 NOT NULL COMMENT 'api链接地址',
  `api_method` tinyint(4) DEFAULT NULL COMMENT '1 get 2 post',
  `ctm` int(11) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='api接口设置';

-- ----------------------------
-- Records of test_api
-- ----------------------------
INSERT INTO `test_api` VALUES ('1', '测试', 'a b c', '28,27', 'http://www.baidu.com', '0', '1529854769');

-- ----------------------------
-- Table structure for test_rule
-- ----------------------------
DROP TABLE IF EXISTS `test_rule`;
CREATE TABLE `test_rule` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rule_name` varchar(50) DEFAULT NULL COMMENT '规则名称',
  `rule` varchar(255) DEFAULT NULL COMMENT '规则',
  `status` tinyint(4) NOT NULL COMMENT '0启用1未启用',
  `ctm` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=latin1 COMMENT='规则表';

-- ----------------------------
-- Records of test_rule
-- ----------------------------
INSERT INTO `test_rule` VALUES ('27', 'asdasd', 'dfsd', '0', '1529854769');
INSERT INTO `test_rule` VALUES ('28', 'qqqqq', 'wwwww', '0', '1529854769');
INSERT INTO `test_rule` VALUES ('29', 'qwqwqw', 'sdfsdfsdf', '0', '1529854769');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `created` int(10) DEFAULT NULL COMMENT '注册时间',
  `changed` int(10) DEFAULT NULL COMMENT '编辑时间',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 0屏蔽，1正常',
  `user_profile_id` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8 COMMENT='用户';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('4', 'ceshi2', '123456', '0', '0', '0', null);
INSERT INTO `user` VALUES ('5', 'ceshi', '123456', '0', '0', '0', null);
INSERT INTO `user` VALUES ('10', 'dsfsdf', 'sadasdas', '0', '0', '0', null);
INSERT INTO `user` VALUES ('29', 'wqrwerw1er', 'rerer', '0', '0', '0', null);
INSERT INTO `user` VALUES ('30', 'wqerwqer', 'wqerwqer', '0', '0', '0', null);

-- ----------------------------
-- Table structure for user_detail
-- ----------------------------
DROP TABLE IF EXISTS `user_detail`;
CREATE TABLE `user_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) DEFAULT NULL,
  `pic` varchar(150) CHARACTER SET latin1 DEFAULT NULL,
  `ctm` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `dec` varchar(50) CHARACTER SET latin1 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_detail
-- ----------------------------
INSERT INTO `user_detail` VALUES ('1', 'wqerwqer1', 'http://www.bai111', '2018-06-08 17:19:55', '?????????');
INSERT INTO `user_detail` VALUES ('5', 'wqerwqer', 'http:www.baidu.com', '2018-06-08 09:45:30', 'sadfwesadfasdf');
INSERT INTO `user_detail` VALUES ('6', 'ceshi2', '', '2018-06-20 15:25:33', 'asdasd');

-- ----------------------------
-- Table structure for user_profile
-- ----------------------------
DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `user_profile` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `realname` varchar(15) DEFAULT NULL,
  `sex` tinyint(1) DEFAULT '1' COMMENT '1boy,0girl',
  `birth` varchar(20) NOT NULL DEFAULT '' COMMENT '生日',
  `email` varchar(20) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
  `hobby` varchar(255) NOT NULL DEFAULT '' COMMENT '爱好',
  `intro` text COMMENT '介绍',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户详情';

-- ----------------------------
-- Records of user_profile
-- ----------------------------
