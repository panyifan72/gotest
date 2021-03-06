/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50626
Source Host           : localhost:3306
Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50626
File Encoding         : 65001

Date: 2016-06-23 10:55:16
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='相册';

-- ----------------------------
-- Records of album
-- ----------------------------
INSERT INTO `album` VALUES ('1', '-1', '/static/uploadfile/2016-6/21/6b5cba0bab0d5aa9ef54c8d3cabf2b53.jpg', '', '', '1466504209', '1', '1');
INSERT INTO `album` VALUES ('2', '1-1', '/static/uploadfile/2016-6/21/90e4302d7d4797150a9c5b2d3986a15d.jpg', '', '', '1466504209', '1', '1');
INSERT INTO `album` VALUES ('3', '1-2', '/static/uploadfile/2016-6/21/e407136ef71a28ed881e9106cde2bf29.jpg', '', '', '1466504209', '1', '1');
INSERT INTO `album` VALUES ('4', '1-3', '/static/uploadfile/2016-6/21/84223adaf14394d924caeb961d8eb299.jpg', '', '', '1466504209', '1', '1');
INSERT INTO `album` VALUES ('5', '2-1', '/static/uploadfile/2016-6/21/04a973332b13a4b357d61598c3b4c210.jpg', '', '', '1466504209', '1', '1');
INSERT INTO `album` VALUES ('6', '4-1', '/static/uploadfile/2016-6/21/9aef204d6117aed7d87c6809d769c486.jpg', '', '', '1466504209', '1', '1');
INSERT INTO `album` VALUES ('7', '4-2', '/static/uploadfile/2016-6/21/6bdaf4cb70077bab4793fd06c0fdf443.jpg', '', '', '1466504209', '1', '1');
INSERT INTO `album` VALUES ('8', '5-1', '/static/uploadfile/2016-6/21/d67ec9310dc49e3e5c7d25a59a8c94dc.jpg', '', '', '1466504209', '1', '1');
INSERT INTO `album` VALUES ('9', '6-1', '/static/uploadfile/2016-6/21/e98d4d899bb4fe7324df841ebf297285.jpg', '', '', '1466504210', '1', '1');
INSERT INTO `album` VALUES ('10', '7-1', '/static/uploadfile/2016-6/21/c4b6f3515a0503ecad9dcaa40fa8d0a5.jpg', '', '', '1466504210', '1', '1');
INSERT INTO `album` VALUES ('11', '7-2', '/static/uploadfile/2016-6/21/e9f00f11520a826d2467b9285d6e8699.jpg', '', '', '1466504210', '1', '1');
INSERT INTO `album` VALUES ('12', '8-1', '/static/uploadfile/2016-6/21/63000ee47639c1fde48a5566b8578acf.jpg', '', '', '1466504210', '1', '1');

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
  PRIMARY KEY (`id`),
  KEY `INDEX_TCVS` (`title`,`created`,`viewnum`,`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8 COMMENT='文章';

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('1', '所谓情商高，就是这样说话hello', 'http://www.nahehuo.com', '努力工作，勤奋，情绪管理', '毕业后，很多学习不如你的人，都慢慢赶超成为了存款6、7位数的winner？你却还是刚毕业时的状态，月光族一个，个人资产连5位数的存款都没有？', '<img src=\"/static/uploadfile/2016-6/7/0ef780a2f3bc341c2f84134bf64f9b1f.jpeg\" alt=\"\" />\n<p>\n	长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。\n</p>\n<p>\n	物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。\n</p>\n他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br />\n之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br />\n<p>\n	你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。\n</p>\n<p>\n	<img src=\"/static/uploadfile/2016-6/7/7ce38ffb5cf3872cdb6813679298c9da.jpg\" alt=\"\" style=\"white-space:normal;\" />\n</p>', 'lock', '0', '0', '1');
INSERT INTO `article` VALUES ('26', '社保缴满15年，真的可以不用再交了吗?', 'http://www.milu365.com', '努力，勤奋，情绪', '毕业七年，在职场中吃过的亏不胜枚举。从最初的职场菜鸟一路成长到如今，仍然在学习和摸索。这条摸着石头过河的职场之路，曲折蜿蜒，一不小心就会走入', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'lock', '0', '0', '1');
INSERT INTO `article` VALUES ('27', '只想做个职场“老实人”，可取吗?', 'http://www.nahehuo.com', '工作，勤奋，管理', '这段时间，朋友圈里几乎被“情商高”和“情商低”的文章占领了，尤其是《欢乐颂》热播以来，大家对情商一词更是关注有加，邱莹莹的低情商，樊胜美的高', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'lock', '1465373118', '1', '1');
INSERT INTO `article` VALUES ('28', 'HR干货：怎样让离职面谈不白谈', 'http://www.nahehuo.com', '努力，工作勤奋，情绪管理', '2016年5月27日下午，在CEO黄总的带领下，今翌集团开展了第一期的工作经验分享会——产品总监修炼之道', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'lock', '1465373118', '1', '1');
INSERT INTO `article` VALUES ('29', '职场新人，怎么正确应对领导的负面评价?', 'http://www.nahehuo.com', '努力工作，工作勤奋，情绪管理', '据悉，日前上海市网信办正深入开展“招聘网站严重违规失信”专项整治行动，在调查核实、固证备查的基础上，先后约谈应届生招聘网、若邻网、51job', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'test', '1465373118', '1', '1');
INSERT INTO `article` VALUES ('31', '一份来自资深HR的就业建议，请仔细阅读', 'http://www.nahehuo.com', '努力，工作，工作勤奋，情绪管理', '每一个求职者，最大的困难就是如何回答面试人员的问题了。其实HR提问也都是有套路的，对不同的工作岗位，总有那么几个问题是必须会提的', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'test', '1465373118', '1', '1');
INSERT INTO `article` VALUES ('32', '特别推荐：对不起，你那不叫努力，叫重复劳动！', 'http://www.nahehuo.com', '努力工作，勤奋，情绪', '希望登上职业巅峰的年轻人请注意，最好不要频繁跳槽。一项研究表明，越来越多的大公司更倾向于从公司内部挑选CEO。拥有财务和工程类背景的内部管理', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'test', '1464838083', '1', '1');
INSERT INTO `article` VALUES ('33', '人人过六一，职场人卖萌减压大法——人家还只是个宝宝嘛', 'http://www.nahehuo.com', '工作，勤奋，情绪', '每个单位人力资源部门，都有一份内部岗位职责说明书。这份岗位职责说明书对一个组织内部各个部门、各个岗位做了详细说明，包括工作说明、也就是说哪个', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', '', '1464838576', '1', '1');
INSERT INTO `article` VALUES ('34', '以下迹象一定要警惕，说明你的老板讨厌你', 'http://www.nahehuo.com', '努力工作，工作勤奋，情绪管理', '上司有两层含义，一个是指上司这个人，一个是指上司这个位子。我们即使不敬重上司本人，也要敬重‘上司’这个位子。', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'fdsa', '1464838811', '1', '1');
INSERT INTO `article` VALUES ('35', '俞敏洪：盲目自信比考虑周到强一百倍', 'http://www.nahehuo.com', '努力工作，工作勤奋，情绪管理', '配偶是谁？就是晚上睡在你旁边的人，是你的老婆或者老公。世界上最大的风不是台风，而是枕边风。配偶太重要了，是成功最大的关键', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'tset', '1464839526', '1', '1');
INSERT INTO `article` VALUES ('36', '在绝对实力面前，情商没有任何用', 'http://www.nahehuo.com', '努力工作，工作勤奋，情绪管理', '相信有很多人在工作中都有“电话恐惧症”，具体表现为：一和客户或同事用电话沟通，就支支吾吾词不达意，甚至出现张口结舌的局面。这会导致对方根本不', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'lock', '1464846139', '1', '1');
INSERT INTO `article` VALUES ('37', '老板不说你永远不知道：领导最看重下属身上什么品质?', 'http://www.nahehuo.com', '努力工作，工作勤奋，情绪管理', '作为管理者，最理想的员工就是那些能够领会自己计划或事项的要旨，能够清楚自己的标准和要求，能够按照自己所设想与安排的去拔掉一颗又一颗钉子', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'lock', '1464849844', '1', '1');
INSERT INTO `article` VALUES ('38', '熬夜加班是可耻的!!', 'http://www.nahehuo.com', '努力工作，工作勤奋，情绪管理', '当你的想法与上司不一样的时候，你应该怎么办?这方面许多学者和专家都给出了很好的建议，这里哪合伙小编根据自己的经验总结提出以下5点建议', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。物竞天择适者生存，你不争不抢不去努力，结果只能是在原地打转，于是乎只能高高仰望别人的光芒。我们都听过这么一段话，这个世界上最可怕的不是有人比你优秀，而是比你优秀的人依然还在努力，那么这样的你为什么还不去奋斗。从来不怕大器晚成，怕的是一生平庸。<br/>他孩子打我孩子1下，傲慢的开口说不就一巴掌赔你一万块，我有足够资格拿出五万甚至十万让孩子打回去。<br/>之前在网上流传起来的这句话，乍听之下会觉得好笑，可是在这句话之下隐藏着的是什么本质呢？试想一下，如果你有钱，那么当别人这么侮辱你侮辱你的孩子时，你完全有资格有资本可以拿出更多的钱来还回去，这并不是教育小孩子暴力和虚荣，而是不愿意让自己的孩子或者自己承受如此大的委屈和伤害。不然假设一下你没钱的情境呢，当别人这么说之后，你顶多会发怒的吼回去“有钱了不起啊，”然后在背地里孩子看不见的地方偷偷哭泣，当然你也可以选择让孩子打回去，可是那时候的你要承担接下来的什么索赔，你又没钱又没势拿什么和人家斗。<br/>你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', 'test', '1464851588', '1', '1');

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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES ('1', '26', 'lock', 'http://www.nahehuo.com', 'lock hello', '1465277058', '1');
INSERT INTO `comment` VALUES ('2', '26', 'xiaoxin', 'http://www.milu365.com', '迷路我的家，在哪里~\n让我们一起去飞~', '1465277231', '1');
INSERT INTO `comment` VALUES ('10', '26', 'Lee', 'http://www.milu365.com', '迷路我的家，在哪里~\n让我们一起去飞~', '1465277231', '1');
INSERT INTO `comment` VALUES ('11', '27', '李先人', 'http://www.nahehuo.com', '你努力创造更好的生活之后，你孩子出生的环境也是良好的，他可以接受好的教育，见识更多的精彩，可以不会因为没钱而忍受饥饿和寒冷，早早的见识了社会的黑暗，他能有更多的机会做他想做的事情，而你也有能力满足他一切合理要求。事实证明，家庭优越的孩子比家庭贫苦的孩子要更加自信，成长的也更健康。当然，如果你教育有问题或者太过溺爱造成他嚣张跋扈那就另说了。', '1465373118', '1');
INSERT INTO `comment` VALUES ('12', '26', 'Lei', 'http://www.nahehuo.com', '让青春一起赚钱吧', '1465375588', '1');
INSERT INTO `comment` VALUES ('13', '26', 'lock-up', 'http://www.nahehuo.com', '迷路我的家，在哪里~ 让我们一起去飞~', '1465376002', '1');
INSERT INTO `comment` VALUES ('14', '26', '李先人', 'http://www.milu365.com', '迷路我的家，在哪里~ 让我们一起去飞~', '1465376048', '1');
INSERT INTO `comment` VALUES ('15', '1', 'Tom', 'http://www.nahehuo.com', '长江后浪推前浪，一浪更比一浪强。竞争社会，拼的就是谁更有能力，谁更能在社会中立得一席之地。', '1465700872', '1');
INSERT INTO `comment` VALUES ('16', '1', '李先人', 'http://www.nahehuo.com', '长江后浪推前浪', '1466415229', '1');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `user_profile_id` int(10) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '注册时间',
  `changed` int(10) DEFAULT NULL COMMENT '编辑时间',
  `status` int(4) NOT NULL DEFAULT '1' COMMENT '状态: 0屏蔽，1正常',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '13888888888', 'e10adc3949ba59abbe56e057f20f883e', null, null, null, '1');

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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户详情';

-- ----------------------------
-- Records of user_profile
-- ----------------------------
INSERT INTO `user_profile` VALUES ('1', '李白', '1', '1985-10-1', 'katongbala@163.com', '13888888888', '中国 上海', '看没有看过的电影', '加油，给自己');
