/*
SQLyog Ultimate
MySQL - 10.3.8-MariaDB : Database - db_stocks
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`db_stocks` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `db_stocks`;

/*Table structure for table `tb_users` */

DROP TABLE IF EXISTS `tb_users`;

CREATE TABLE `tb_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(30) DEFAULT '' COMMENT '账号',
  `pass` varchar(30) DEFAULT '' COMMENT '密码',
  `real_name` varchar(30) DEFAULT '' COMMENT '姓名',
  `phone` char(11) DEFAULT '' COMMENT '手机',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `token` varchar(300) DEFAULT '',
  `remark` varchar(300) DEFAULT '' COMMENT '备注',
  `last_login_time` datetime DEFAULT current_timestamp(),
  `last_login_ip` char(30) DEFAULT NULL COMMENT '最近一次登录ip',
  `login_times` int(11) DEFAULT 1 COMMENT '累计登录次数',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

/*Data for the table `tb_users` */

insert  into `tb_users`(`id`,`username`,`pass`,`real_name`,`phone`,`status`,`token`,`remark`,`last_login_time`,`last_login_ip`,`login_times`,`created_at`,`updated_at`) values
(1,'admin','admin9527','管理员','16601770915',1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjEsIm5hbWUiOiIiLCJwaG9uZSI6IjE2NjAxNzcwOTE1IiwiZXhwIjoxNTg3MzE3NjE0LCJuYmYiOjE1ODczMTQwMDR9._P5BnoYB1fDPHYc98L_GSCVyOMrTBkEVEAN2GUgtFK4','','2020-04-19 23:43:50',NULL,1,'2020-04-19 02:41:54','2020-04-19 02:41:54'),
(2,'zhangsanfeng','wewewewewewe','张三丰','16601770915',1,'','','2020-04-19 23:43:50',NULL,1,'2020-04-19 02:42:19','2020-04-19 02:42:19'),
(3,'testadmin','hello20154','测试1','16601770915',1,'','','2020-04-19 23:43:50',NULL,1,'2020-04-19 02:48:11','2020-04-19 02:48:11'),
(4,'zhangcuishan','yuikjll54','张翠山','15618726171',1,'','','2020-04-19 23:43:50',NULL,1,'2020-04-19 12:32:10','2020-04-19 12:32:10'),
(5,'zhangwuji','secret20212','张无忌','16601770915',1,'','测试数据','2020-04-19 23:43:50',NULL,1,'2020-04-19 22:47:43','2020-04-19 22:47:43'),
(8,'zhangcuishan','yuikjll54','张翠山','15618726171',1,'0','','2020-04-19 23:43:50',NULL,1,'2020-04-19 23:28:35','2020-04-19 23:28:35'),
(9,'lisi2021','secret20212','李四','16601770915',1,'','测试数据','2020-04-20 00:36:06',NULL,1,'2020-04-20 00:36:06','2020-04-20 00:36:06'),
(10,'lisi2025','secret20212','李四','16601770915',1,'','测试数据','2020-04-20 00:38:46',NULL,1,'2020-04-20 00:38:46','2020-04-20 00:38:46');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
