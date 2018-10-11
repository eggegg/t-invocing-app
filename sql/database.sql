create database if not exists t_invocing_app;
USE  t_invocing_app;


CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(255) DEFAULT '' COMMENT '用户名，不可修改',
  `password` varchar(255) DEFAULT '' COMMENT '用户密码',
  `email` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `Index_users_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户表';

INSERT into users values (1, "admin", "$2a$10$sNLIMmdXqM0zR47dvwLj0.eaN.AJvQESMTwIvSLurE/SJjwH58fZO", "admin@admin.com");


CREATE TABLE IF NOT EXISTS `invoices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) DEFAULT '',
  `user_id` int(10) unsigned,
  `paid` tinyint(1),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `transactions` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name` varchar(255) DEFAULT '',
   `price` int(10) unsigned NOT NULL,
   `invoice_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;