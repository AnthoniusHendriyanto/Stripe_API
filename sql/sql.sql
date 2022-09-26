CREATE DATABASE IF NOT EXISTS `stripe`;

CREATE TABLE IF NOT EXISTS `card` (
    `id` tinyint unsigned NOT NULL AUTO_INCREMENT,
    `cardId` varchar(100) NOT NULL,
    `brand` varchar(100) NOT NULL,
    `customerId` varchar(100) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `cardId_UN` (`cardId`),
    CONSTRAINT `card_customer_FK` FOREIGN KEY (`customerId`) REFERENCES `customer` (`customerId`)
);

CREATE TABLE IF NOT EXISTS `customer` (
    `id` tinyint unsigned NOT NULL AUTO_INCREMENT,
    `customerId` varchar(100) NOT NULL,
    `name` varchar(100) NOT NULL,
    `phone` varchar(100) NOT NULL,
    `status` tinyint NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `customer_id_UN` (`customerId`),
    KEY `status_customer_FK` (`status`),
    CONSTRAINT `customer_status_FK` FOREIGN KEY (`status`) REFERENCES `customer_status` (`status`)
);

CREATE TABLE IF NOT EXISTS `customer_status` (
    `id` tinyint unsigned NOT NULL AUTO_INCREMENT,
    `status` tinyint NOT NULL,
    `description` char(32) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `customer_status_UN` (`status`)
);

INSERT INTO stripe.customer_status (status, description) VALUES (0, "Inactive");
INSERT INTO stripe.customer_status (status, description) VALUES (1, "Active");
