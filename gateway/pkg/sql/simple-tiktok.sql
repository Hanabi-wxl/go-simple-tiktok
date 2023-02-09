/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 8.0.23 : Database - simple_tiktok
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
USE `simple_tiktok`;

/*Table structure for table `comment` */

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `user_id` int NOT NULL,
                           `video_id` int NOT NULL,
                           `comment_text` text NOT NULL,
                           `comment_time` timestamp NULL DEFAULT NULL,
                           `is_deleted` tinyint DEFAULT '0',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `favorite` */

DROP TABLE IF EXISTS `favorite`;

CREATE TABLE `favorite` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `user_id` int NOT NULL,
                            `video_id` int NOT NULL,
                            `is_deleted` tinyint DEFAULT '0',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

/*Table structure for table `follow` */

DROP TABLE IF EXISTS `follow`;

CREATE TABLE `follow` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `follower_id` int NOT NULL,
                          `follow_id` int NOT NULL,
                          `is_deleted` tinyint DEFAULT '0',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `from_user_id` int NOT NULL,
                           `to_user_id` int NOT NULL,
                           `content` text NOT NULL,
                           `send_time` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
                        `user_id` int NOT NULL AUTO_INCREMENT,
                        `password` varchar(128) DEFAULT NULL,
                        `name` varchar(32) DEFAULT NULL,
                        PRIMARY KEY (`user_id`),
                        UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

/*Table structure for table `video` */

DROP TABLE IF EXISTS `video`;

CREATE TABLE `video` (
                         `video_id` int NOT NULL AUTO_INCREMENT,
                         `title` varchar(32) DEFAULT NULL,
                         `author` int DEFAULT NULL,
                         `play_url` text,
                         `cover_url` text,
                         `upload_time` datetime DEFAULT NULL,
                         PRIMARY KEY (`video_id`),
                         UNIQUE KEY `video_id` (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
