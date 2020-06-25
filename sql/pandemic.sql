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
  `name` varchar(64) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `FK_disease_tier_id` (`tier`),
  CONSTRAINT `FK_disease_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `disease` DISABLE KEYS */;
/*!40000 ALTER TABLE `disease` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `event` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  `tier` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `FK_event_tier_id` (`tier`),
  CONSTRAINT `FK_event_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `event` DISABLE KEYS */;
/*!40000 ALTER TABLE `event` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `manufacture` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tier` int(11) DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  `name` varchar(64) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `production_speed` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `FK_manufacture_tier_id` (`tier`),
  CONSTRAINT `FK_manufacture_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `manufacture` DISABLE KEYS */;
/*!40000 ALTER TABLE `manufacture` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `medication` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `base_value` int(11) DEFAULT NULL,
  `research_cost` int(11) DEFAULT NULL,
  `maximum_traits` int(11) DEFAULT NULL,
  `tier` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `FK_medication_tier_id` (`tier`),
  CONSTRAINT `FK_medication_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `medication` DISABLE KEYS */;
INSERT INTO `medication` (`id`, `name`, `description`, `base_value`, `research_cost`, `maximum_traits`, `tier`) VALUES
	(1, 'Acetaminophen', 'Acetaminophen is a drug used to treat mild to moderate pain.', 5, 100, 1, 2),
	(2, 'Ibuprofen', 'Ibuprofen reduces inflammation and pain in the body.', 7, 200, 1, 2),
	(3, 'Aspirin', 'Aspirin can be used to treat mild pain, fever or inflammation.', 6, 150, 1, 2),
	(4, 'Adderall', 'Adderall is used to treat attention deficit hyperactivity disorder (ADHD) and narcolepsy.', 15, 1000, 1, 3),
	(5, 'Multi-Vitamin', 'Multi-Vitamin is used to treat vitamin deficiency.', 2, 50, 0, 1),
	(6, 'Vitamin D', 'Vitamin D is used to treat vitamin D deficiency.', 1, 20, 0, 1),
	(7, 'Penicillin G benzathine', 'Penicillin is an antibiotic that is used to treat many types of mild to moderate infections.', 50, 5000, 3, 4),
	(8, 'Mixatrone', 'Mixatrone is a cancer medication that interferes with the growth and spread of cancer cells in the body.', 200, 20000, 5, 5);
/*!40000 ALTER TABLE `medication` ENABLE KEYS */;

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

/*!40000 ALTER TABLE `medication_disease` DISABLE KEYS */;
/*!40000 ALTER TABLE `medication_disease` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `medication_trait` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `tier` int(11) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `FK_medication_trait_tier_id` (`tier`),
  CONSTRAINT `FK_medication_trait_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `medication_trait` DISABLE KEYS */;
INSERT INTO `medication_trait` (`id`, `name`, `description`, `tier`) VALUES
	(1, 'Efficient', '', 2),
	(2, 'Superior', '', 3),
	(3, 'Groundbreaking', '', 5),
	(4, 'Ineffective', '', 1),
	(5, 'Headache inducing', '', 1),
	(6, 'Nausiating', '', 1),
	(7, 'Painful', '', 2),
	(8, 'Multi-Purpouse', '', 1);
/*!40000 ALTER TABLE `medication_trait` ENABLE KEYS */;

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

/*!40000 ALTER TABLE `researcher` DISABLE KEYS */;
/*!40000 ALTER TABLE `researcher` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `researcher_trait` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tier` int(11) DEFAULT NULL,
  `name` varchar(64) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `rarity` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `FK_researcher_trait_tier_id` (`tier`),
  CONSTRAINT `FK_researcher_trait_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `researcher_trait` DISABLE KEYS */;
/*!40000 ALTER TABLE `researcher_trait` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `tier` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `color` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `tier` DISABLE KEYS */;
/*!40000 ALTER TABLE `tier` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `user` (
  `id` varchar(64) NOT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `accesslevel` int(11) DEFAULT 1,
  `tier` int(11) DEFAULT 1,
  `balance` int(20) DEFAULT 0,
  `manufacture` int(11) DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `user_disease` (
  `user` varchar(64) NOT NULL,
  `disease` int(11) NOT NULL,
  PRIMARY KEY (`user`,`disease`),
  KEY `FK_user_disease_disease_id` (`disease`),
  KEY `FK_user_disease_user_id` (`user`),
  CONSTRAINT `FK_user_disease_disease_id` FOREIGN KEY (`disease`) REFERENCES `disease` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_disease_user_id` FOREIGN KEY (`user`) REFERENCES `user` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40000 ALTER TABLE `user_disease` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_disease` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `user_tier` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user` varchar(64) NOT NULL,
    `tier` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    KEY `FK_user_tier_tier_id` (`tier`),
    KEY `FK_user_tier_user_id` (`user`),
    CONSTRAINT `FK_user_tier_tier_id` FOREIGN KEY (`tier`) REFERENCES `tier` (`id`) ON UPDATE NO ACTION,
    CONSTRAINT `FK_user_tier_user_id` FOREIGN KEY (`user`) REFERENCES `user` (`id`) ON UPDATE NO ACTION
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

/*!40000 ALTER TABLE `user_event` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_event` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `user_friend` (
  `user` varchar(64) NOT NULL,
  `friend` varchar(64) NOT NULL,
#   0 - pending
#   1 - accepted
#   2 - declined
  `status` int(11),
  PRIMARY KEY (`user`, `friend`),
  KEY `FK_user_friend_user_id` (`user`),
  KEY `FK_user_friend_friend_id` (`friend`),
  CONSTRAINT `FK_user_friend_id` FOREIGN KEY (`friend`) REFERENCES `user` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_friend_user_id` FOREIGN KEY (`user`) REFERENCES `user` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `user_friend` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_friend` ENABLE KEYS */;

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

/*!40000 ALTER TABLE `user_medication` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_medication` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `user_medication_trait` (
  `user_medication` int(11) NOT NULL,
  `medication_trait` int(11) NOT NULL,
  PRIMARY KEY (`user_medication`,`medication_trait`),
  KEY `FK_user_medication_trait_medication_trait_id` (`medication_trait`),
  KEY `FK_user_medication_trait_user_medication_id` (`user_medication`),
  CONSTRAINT `FK_user_medication_trait_medication_trait_id` FOREIGN KEY (`medication_trait`) REFERENCES `medication_trait` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_medication_trait_user_medication_id` FOREIGN KEY (`user_medication`) REFERENCES `user_medication` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `user_medication_trait` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_medication_trait` ENABLE KEYS */;

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

/*!40000 ALTER TABLE `user_researcher` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_researcher` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `user_researcher_trait` (
  `user_researcher` int(11) NOT NULL,
  `researcher_trait` int(11) NOT NULL,
  PRIMARY KEY (`user_researcher`,`researcher_trait`),
  KEY `FK_user_researcher_trait_researcher_trait_id` (`researcher_trait`),
  KEY `FK_user_researcher_trait_user_researcher_id` (`user_researcher`),
  CONSTRAINT `FK_user_researcher_trait_researcher_trait_id` FOREIGN KEY (`researcher_trait`) REFERENCES `researcher_trait` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_researcher_trait_user_researcher_id` FOREIGN KEY (`user_researcher`) REFERENCES `user_researcher` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40000 ALTER TABLE `user_researcher_trait` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_researcher_trait` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
