/*
SQLyog Community v13.1.9 (64 bit)
MySQL - 8.0.40 : Database - umdp
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`umdp` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `umdp`;

/*Table structure for table `channel` */

DROP TABLE IF EXISTS `channel`;

CREATE TABLE `channel` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `channel_name` varchar(50) NOT NULL COMMENT '渠道名称',
  `channel_tag` varchar(100) DEFAULT NULL COMMENT '渠道标识',
  `channel_config` json DEFAULT NULL COMMENT '渠道配置',
  `channel_status` tinyint(1) DEFAULT '1' COMMENT '渠道状态 0 禁用 1 启用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `tag` (`channel_tag`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `channel` */

/*Table structure for table `log` */

DROP TABLE IF EXISTS `log`;

CREATE TABLE `log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增',
  `profession_id` int DEFAULT NULL COMMENT '业务编号',
  `channel_id` int DEFAULT NULL COMMENT '渠道编号',
  `template_id` int DEFAULT NULL COMMENT '模板编号',
  `parameters` json DEFAULT NULL COMMENT '调用参数',
  `receiver` json DEFAULT NULL COMMENT '接收者',
  `status` tinyint(1) DEFAULT '0' COMMENT '是否成功 1 成功 0 失败',
  `request_id` varchar(100) DEFAULT NULL COMMENT '调用编号',
  `err_message` text COMMENT '错误信息',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '调用时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '调用结束时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=91561 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `log` */

/*Table structure for table `phone_log` */

DROP TABLE IF EXISTS `phone_log`;

CREATE TABLE `phone_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `channel_id` int DEFAULT NULL COMMENT '渠道编号',
  `call_resp` json DEFAULT NULL COMMENT '状态返回',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `custom_voice_id` varchar(50) DEFAULT NULL COMMENT '自定义呼叫编号',
  `phone` char(11) DEFAULT NULL COMMENT '手机号码',
  `call_count` int DEFAULT NULL COMMENT '已拨打',
  `call_limit` int DEFAULT NULL COMMENT '拨打限制',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3566 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `phone_log` */

/*Table structure for table `profession` */

DROP TABLE IF EXISTS `profession`;

CREATE TABLE `profession` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `profession_name` varchar(100) DEFAULT NULL COMMENT '业务名称',
  `token` varchar(100) DEFAULT NULL COMMENT '校验密钥',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `profession` */

insert  into `profession`(`id`,`profession_name`,`token`,`created_at`,`updated_at`) values 
(5,'短信渠道测试','sdwdsadwqedfa','2023-12-08 11:41:36','2023-12-11 19:37:51'),
(6,'腾讯短信测试','sasadadasda','2023-12-11 18:44:46','2023-12-11 18:44:46'),
(8,'钉钉测试','338849b8-9fca-44ee-9403-eb2dcdc926c4','2023-12-18 19:28:12','2023-12-18 19:28:12'),
(9,'飞书测试','80aa886d-7e6b-4a10-8e08-bc8f0e5a6fe7','2023-12-20 14:09:48','2023-12-20 14:09:48');

/*Table structure for table `profession_channel` */

DROP TABLE IF EXISTS `profession_channel`;

CREATE TABLE `profession_channel` (
  `profession_id` int NOT NULL COMMENT '业务编号',
  `channel_id` int NOT NULL COMMENT '渠道编号',
  UNIQUE KEY `pc` (`profession_id`,`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `profession_channel` */

insert  into `profession_channel`(`profession_id`,`channel_id`) values 
(5,47),
(6,48),
(8,49),
(9,50);

/*Table structure for table `template` */

DROP TABLE IF EXISTS `template`;

CREATE TABLE `template` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增',
  `template_name` varchar(100) DEFAULT NULL COMMENT '模板名称',
  `profession_id` int DEFAULT NULL COMMENT '业务编号',
  `retry` int DEFAULT '0' COMMENT '重试次数 默认0次不重试',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `template` */

insert  into `template`(`id`,`template_name`,`profession_id`,`retry`,`created_at`,`updated_at`) values 
(25,'钉钉测试',8,0,'2023-12-19 10:22:47','2023-12-20 13:12:09'),
(26,'飞书测试',9,0,'2023-12-20 14:36:04','2023-12-20 14:36:04');

/*Table structure for table `template_channel` */

DROP TABLE IF EXISTS `template_channel`;

CREATE TABLE `template_channel` (
  `template_id` int DEFAULT NULL COMMENT '模板编号',
  `channel_id` int DEFAULT NULL COMMENT '渠道编号',
  `config` json DEFAULT NULL COMMENT '渠道配置'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `template_channel` */

insert  into `template_channel`(`template_id`,`channel_id`,`config`) values 
(25,49,'{\"id\": 49, \"type\": \"钉钉\", \"title\": \"{$title}\", \"content\": \"# 这是支持markdown的文本   \\n## 标题2      \\n* 列表1  \\n![alt 啊](https://img.alicdn.com/tps/TB1XLjqNVXXXXc4XVXXXXXXXXXX-170-64.png)\", \"messageType\": 2}'),
(26,50,'{\"id\": 50, \"type\": \"飞书\", \"content\": \"{$content}\", \"messageType\": 1}');

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增',
  `username` varchar(100) NOT NULL COMMENT '登录用户名',
  `password` char(64) NOT NULL COMMENT '密码',
  `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '显示名称',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `user` */

insert  into `user`(`id`,`username`,`password`,`nickname`,`created_at`,`updated_at`) values 
(1,'admin','21232f297a57a5a743894a0e4a801fc3','admin','2025-04-14 15:03:54','2025-04-14 15:03:54');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
