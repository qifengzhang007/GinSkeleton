
CREATE DATABASE /*!32312 IF NOT EXISTS*/`db_goskeleton` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `db_goskeleton`;

/*Table structure for table `tb_users` */

DROP TABLE IF EXISTS `tb_users`;

CREATE TABLE `tb_users` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(30) DEFAULT '' COMMENT '账号',
  `pass` VARCHAR(128) DEFAULT '' COMMENT '密码',
  `real_name` VARCHAR(30) DEFAULT '' COMMENT '姓名',
  `phone` CHAR(11) DEFAULT '' COMMENT '手机',
  `status` TINYINT(4) DEFAULT 1 COMMENT '状态',
  `remark` VARCHAR(300) DEFAULT '' COMMENT '备注',
  `last_login_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `last_login_ip` CHAR(30) DEFAULT '' COMMENT '最近一次登录ip',
  `login_times` INT(11) DEFAULT 0 COMMENT '累计登录次数',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/* oauth 表，主要控制一个用户可以同时拥有几个有效的token，通俗地说就是允许一个账号同时有几个人登录，超过将会导致最前面的人的token失效，而退出登录*/
DROP TABLE IF EXISTS `tb_oauth_access_tokens`;

CREATE TABLE `tb_oauth_access_tokens` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `fr_user_id` INT(11) DEFAULT 0 COMMENT '外键:tb_users表id',
  `client_id` INT(10) UNSIGNED DEFAULT 1 COMMENT '普通用户的授权，默认为1',
  `token` VARCHAR(600) DEFAULT NULL,
  `action_name` VARCHAR(128) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT 'login|refresh|reset表示token生成动作',
  `scopes` VARCHAR(128) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '[*]' COMMENT '暂时预留,未启用',
  `revoked` TINYINT(1) DEFAULT 0 COMMENT '是否撤销',
  `client_ip` VARCHAR(128) DEFAULT NULL COMMENT 'ipv6最长为128位',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `expires_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `oauth_access_tokens_user_id_index` (`fr_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/* 创建基于casbin控制接口访问的权限表*/
DROP TABLE IF EXISTS `tb_auth_casbin_rule`;
CREATE TABLE `tb_auth_casbin_rule` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`ptype` varchar(100) DEFAULT '',
`v0` varchar(100) DEFAULT '',
`v1` varchar(100) DEFAULT '',
`v2` varchar(100) DEFAULT '*',
`v3` varchar(100) DEFAULT '',
`v4` varchar(100) DEFAULT '',
`v5` varchar(100) DEFAULT '',
PRIMARY KEY (`id`),
UNIQUE KEY `unique_index` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8


