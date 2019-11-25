CREATE DATABASE IF NOT EXISTS GOREST;

use GOREST;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE IF NOT EXISTS `users`
(
    `id`       int(11)                         NOT NULL AUTO_INCREMENT,
    `name`     varchar(255)                    NOT NULL,
    `surname`  varchar(255) CHARACTER SET utf8 NOT NULL,
    `password` varchar(255)                    NOT NULL,
    `email`    varchar(255)                    NOT NULL,
    `isActive` tinyint(4)                      NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  AUTO_INCREMENT = 1;
