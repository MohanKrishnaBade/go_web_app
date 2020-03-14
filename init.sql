
CREATE DATABASE IF NOT EXISTS go_project;
USE go_project;
CREATE TABLE `users` (
  `id` varchar(100) NOT NULL,
  `email` varchar(200) NOT NULL,
  `firstName` varchar(200) DEFAULT NULL,
  `lastName` varchar(200) DEFAULT NULL,
  `picture` text NOT NULL,
  `password` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT COLLATE utf8mb4_unicode_ci;