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
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `comment_id_uindex` (`id`),
                           KEY `comment_video_id_index` (`user_id`),
                           KEY `comment_video_id_index_2` (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

/*Table structure for table `favorite` */

DROP TABLE IF EXISTS `favorite`;

CREATE TABLE `favorite` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `user_id` int NOT NULL,
                            `video_id` int NOT NULL,
                            `is_deleted` tinyint DEFAULT '0',
                            PRIMARY KEY (`id`),
                            KEY `favorite_user_id_index` (`user_id`),
                            KEY `favorite_video_id_index` (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8;

/*Table structure for table `follow` */

DROP TABLE IF EXISTS `follow`;

CREATE TABLE `follow` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `follower_id` int NOT NULL,
                          `follow_id` int NOT NULL,
                          `is_deleted` tinyint DEFAULT '0',
                          PRIMARY KEY (`id`),
                          KEY `follow_id_index` (`follow_id`),
                          KEY `follower_id_index` (`follower_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `from_user_id` int NOT NULL,
                           `to_user_id` int NOT NULL,
                           `content` text NOT NULL,
                           `receiver_read` tinyint DEFAULT '0',
                           `sender_read` tinyint DEFAULT '0',
                           `send_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`),
                           KEY `message_from_user_id_index` (`from_user_id`),
                           KEY `message_to_user_id_index` (`to_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=140 DEFAULT CHARSET=utf8;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
                        `user_id` int NOT NULL AUTO_INCREMENT,
                        `name` varchar(32) DEFAULT NULL,
                        `password` varchar(128) DEFAULT NULL,
                        `avatar` varchar(128) DEFAULT NULL,
                        `signature` varchar(128) DEFAULT NULL,
                        PRIMARY KEY (`user_id`),
                        UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

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
                         UNIQUE KEY `video_id` (`video_id`),
                         KEY `video_author_index` (`author`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
