CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(200) NOT NULL,
  `last_name` varchar(200) NOT NULL,
  `email` varchar(200) DEFAULT NULL,
  `mobile_number` bigint NOT NULL,
  `address` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS `orders` (
  `id` bigint UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL,
  `price` bigint,
  `order_date` DATETIME DEFAULT Now(),
  `is_deleted` tinyint DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;