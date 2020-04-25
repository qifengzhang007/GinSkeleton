CREATE DATABASE /*!32312 IF NOT EXISTS*/`db_stocks` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `db_stocks`;

/*Table structure for table `tb_users` */

DROP TABLE IF EXISTS `tb_users`;

CREATE TABLE `tb_users` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(30) DEFAULT '' COMMENT '账号',
  `pass` VARCHAR(128) DEFAULT '' COMMENT '密码',
  `real_name` VARCHAR(30) DEFAULT '' COMMENT '姓名',
  `phone` CHAR(11) DEFAULT '' COMMENT '手机',
  `status` TINYINT(4) DEFAULT 1 COMMENT '状态',
  `token` VARCHAR(300) DEFAULT '',
  `remark` VARCHAR(300) DEFAULT '' COMMENT '备注',
  `last_login_time` DATETIME DEFAULT CURRENT_TIMESTAMP(),
  `last_login_ip` CHAR(30) DEFAULT NULL COMMENT '最近一次登录ip',
  `login_times` INT(11) DEFAULT 1 COMMENT '累计登录次数',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP(),
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP(),
  PRIMARY KEY (`id`)
) ENGINE=MYISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

/*Data for the table `tb_users` */

INSERT  INTO `tb_users`(`id`,`username`,`pass`,`real_name`,`phone`,`status`,`token`,`remark`,`last_login_time`,`last_login_ip`,`login_times`,`created_at`,`updated_at`) VALUES
(1,'admin','87d9bb400c0634691f0e3baaf1e2fd0d','','',1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjExLCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE1ODc4NDAxMDYsIm5iZiI6MTU4NzgzMDQwMn0._mZcHdzzmsYYXPxuoVyXzw7U_9Rku7fCmkoWJ9EEdaQ','','2020-04-25 23:51:28','127.0.0.1',1,'2020-04-25 23:51:28','2020-04-25 23:51:28'),
(2,'hello','188bda0c10088d7c2e6d7c00592679e7','','',1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjEzLCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE1ODc4Mzc4OTYsIm5iZiI6MTU4NzgzNDI4Nn0.qayu_u7mEYjTpHPxhgFJtSdGGFHI9rxkwR_RZx_T51E','','2020-04-26 00:59:25','127.0.0.1',1,'2020-04-26 00:59:25','2020-04-26 00:59:25');


