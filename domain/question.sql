CREATE DATABASE IF NOT EXISTS archi;
USE archi;

CREATE TABLE question (
  `id` varchar(36) NOT NULL PRIMARY KEY,
  `user_id` VARCHAR(32) NOT NULL,
  `title` varchar(250) NOT NULL ,
  `content` text NOT NULL,
  `create_time` TIMESTAMP NOT NULL DEFAULT current_timestamp,
  `follow_count` int DEFAULT 0,
  KEY `idx_user_id` (`user_id`)
)ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;