USE garden;
SET NAMES utf8;
SET time_zone = 'Europe/Prague';

CREATE TABLE `plan`
(
  `id`   INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(30) NOT NULL,
  UNIQUE (`name`)
) ENGINE = InnoDB;

CREATE TABLE `plan_value`
(
  `id`      INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `plan_id` INT UNSIGNED NOT NULL,
  `name`    VARCHAR(30)  NOT NULL,
  UNIQUE (`plan_id`, `name`),
  FOREIGN KEY (`plan_id`) REFERENCES `plan` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE TABLE `raspberry`
(
  `id`         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name`       VARCHAR(30)  NOT NULL,
  `plan_id`    INT UNSIGNED NOT NULL,
  `address`    VARCHAR(100) NOT NULL,
  `is_active`  TINYINT(1)   NOT NULL,
  `user`       VARCHAR(50)  NOT NULL,
  `password`   VARCHAR(50)  NOT NULL,
  `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`plan_id`) REFERENCES `plan` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE TABLE `raspberry_plan_value`
(
  `raspberry_id`  INT UNSIGNED NOT NULL,
  `plan_value_id` INT UNSIGNED NOT NULL,
  `value`         DECIMAL(5, 2) NOT NULL,
  `note`          VARCHAR(255) NOT NULL,
  PRIMARY KEY (`plan_value_id`, `raspberry_id`),
  FOREIGN KEY (`raspberry_id`) REFERENCES `raspberry` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`plan_value_id`) REFERENCES `plan_value` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE TABLE `periphery`
(
  `id`            INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name`          VARCHAR(30) NOT NULL,
  `is_measurable` TINYINT(1)  NOT NULL,
  UNIQUE (`name`)
) ENGINE = InnoDB;

CREATE TABLE `pinout`
(
  `id`           INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `raspberry_id` INT UNSIGNED NOT NULL,
  `periphery_id` INT UNSIGNED NOT NULL,
  `pin`          VARCHAR(2)   NOT NULL,
  `created_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`raspberry_id`) REFERENCES `raspberry` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`periphery_id`) REFERENCES `periphery` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE TABLE `job`
(
  `id`   INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(30) NOT NULL,
  UNIQUE (`name`)
) ENGINE = InnoDB;

CREATE TABLE `work`
(
  `id`           INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `raspberry_id` INT UNSIGNED NOT NULL,
  `job_id`       INT UNSIGNED NOT NULL,
  `params`       JSON         NULL,
  `created_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`raspberry_id`) REFERENCES `raspberry` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`job_id`) REFERENCES `job` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE TABLE `measurement`
(
  `id`           INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `work_id`      INT UNSIGNED NOT NULL,
  `periphery_id` INT UNSIGNED NOT NULL,
  `value`        DECIMAL(5, 2) NOT NULL,
  `created_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`work_id`) REFERENCES `work` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`periphery_id`) REFERENCES `periphery` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;
