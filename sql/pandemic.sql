/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

CREATE DATABASE IF NOT EXISTS `pandemic` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `pandemic`;

CREATE TABLE IF NOT EXISTS `disease` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tier` int(11) DEFAULT NULL,
  `name` varchar(64) NOT NULL UNIQUE,
  `description` varchar(255) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_disease_tier_id` (`tier`),
  CONSTRAINT `FK_disease_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `event` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL UNIQUE,
  `description` varchar(255) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  `tier` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_event_tier_id` (`tier`),
  CONSTRAINT `FK_event_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `manufacture` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tier` int(11) DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  `name` varchar(64) NOT NULL UNIQUE,
  `description` varchar(255) DEFAULT NULL,
  `production_speed` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_manufacture_tier_id` (`tier`),
  CONSTRAINT `FK_manufacture_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `medication` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL UNIQUE,
  `description` varchar(255) DEFAULT NULL,
  `research_cost` int(11) DEFAULT NULL,
  `maximum_traits` int(11) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  `tier` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_medication_tier_id` (`tier`),
  CONSTRAINT `FK_medication_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `medication_disease` (
  `medication` int(11) NOT NULL,
  `disease` int(11) NOT NULL,
  `effectiveness` int(11) NOT NULL,
  PRIMARY KEY (`medication`,`disease`),
  KEY `FK_medication_disease_disease_id` (`disease`),
  KEY `FK_medication_disease_medication_id` (`medication`) USING BTREE,
  CONSTRAINT `FK_medication_disease_disease_id` FOREIGN KEY (`disease`) REFERENCES `disease` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_medication_disease_medication_id` FOREIGN KEY (`medication`) REFERENCES `medication` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `medication_trait` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tier` int(11) DEFAULT NULL,
  `name` varchar(64) NOT NULL UNIQUE,
  `description` varchar(255) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_medication_trait_tier_id` (`tier`),
  CONSTRAINT `FK_medication_trait_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `researcher` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tier` int(11) DEFAULT NULL,
  `research_speed` int(11) DEFAULT NULL,
  `salary` int(11) DEFAULT NULL,
  `maximum_traits` int(11) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_researcher_tier_id` (`tier`),
  CONSTRAINT `FK_researcher_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `researcher_trait` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tier` int(11) DEFAULT NULL,
  `name` varchar(64) NOT NULL UNIQUE,
  `description` varchar(255) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_researcher_trait_tier_id` (`tier`),
  CONSTRAINT `FK_researcher_trait_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `tier` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL UNIQUE,
  `color` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user` (
  `id` varchar(64) NOT NULL UNIQUE,
  `username` varchar(64) NOT NULL UNIQUE,
  `password` varchar(64) NOT NULL,
  `accesslevel` int(11) DEFAULT 1,
  `balance` int(20) DEFAULT 0,
  `manufacture` int(11) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_disease` (
  `user` varchar(64) NOT NULL,
  `disease` int(11) NOT NULL,
  PRIMARY KEY (`user`,`disease`),
  KEY `FK_user_disease_disease_id` (`disease`),
  KEY `FK_user_disease_user_id` (`user`),
  CONSTRAINT `FK_user_disease_disease_id` FOREIGN KEY (`disease`) REFERENCES `disease` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_disease_user_id` FOREIGN KEY (`user`) REFERENCES `user` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_event` (
  `user` varchar(64) NOT NULL,
  `event` int(11) NOT NULL,
  PRIMARY KEY (`user`,`event`),
  KEY `FK_user_event_event_id` (`event`),
  KEY `FK_user_event_user_id` (`user`),
  CONSTRAINT `FK_user_event_event_id` FOREIGN KEY (`event`) REFERENCES `event` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_event_user_id` FOREIGN KEY (`user`) REFERENCES `user` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_medication` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` varchar(64) NOT NULL,
  `medication` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_user_medication_user_id` (`user`),
  KEY `FK_user_medication_medication_id` (`medication`),
  CONSTRAINT `FK_user_medication_medication_id` FOREIGN KEY (`medication`) REFERENCES `medication` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_medication_user_id` FOREIGN KEY (`user`) REFERENCES `user` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_medication_trait` (
  `user_medication` int(11) NOT NULL,
  `medication_trait` int(11) NOT NULL,
  PRIMARY KEY (`user_medication`,`medication_trait`),
  KEY `FK_user_medication_trait_medication_trait_id` (`medication_trait`),
  KEY `FK_user_medication_trait_user_medication_id` (`user_medication`),
  CONSTRAINT `FK_user_medication_trait_medication_trait_id` FOREIGN KEY (`medication_trait`) REFERENCES `medication_trait` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_medication_trait_user_medication_id` FOREIGN KEY (`user_medication`) REFERENCES `user_medication` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_researcher` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` varchar(64) DEFAULT NULL,
  `researcher` int(11) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_user_researcher_user_id` (`user`),
  KEY `FK_user_researcher_researcher_id` (`researcher`),
  CONSTRAINT `FK_user_researcher_researcher_id` FOREIGN KEY (`researcher`) REFERENCES `researcher` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_researcher_user_id` FOREIGN KEY (`user`) REFERENCES `user` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_researcher_trait` (
  `user_researcher` int(11) NOT NULL,
  `researcher_trait` int(11) NOT NULL,
  PRIMARY KEY (`user_researcher`,`researcher_trait`),
  KEY `FK_user_researcher_trait_researcher_trait_id` (`researcher_trait`),
  KEY `FK_user_researcher_trait_user_researcher_id` (`user_researcher`),
  CONSTRAINT `FK_user_researcher_trait_researcher_trait_id` FOREIGN KEY (`researcher_trait`) REFERENCES `researcher_trait` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_researcher_trait_user_researcher_id` FOREIGN KEY (`user_researcher`) REFERENCES `user_researcher` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
