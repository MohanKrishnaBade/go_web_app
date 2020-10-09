# go_web_app
Dockerized golang web application.

# setup
1. install docker & docker-compose on your machine.
2. install mysql workbench or sequel Pro on your machine.
3. clone repository.
4. make sure ports 80 and 3306 are free to run this application.
5. run docker-compose  build && docker-compose up to spin up your go and mysql containers
6. connect workbench/sequel pro to mysql db
7. create a user table by running this Query 
```
CREATE TABLE `users` (
  `id` varchar(100) NOT NULL,
  `email` varchar(200) NOT NULL,
  `firstName` varchar(200) DEFAULT NULL,
  `lastName` varchar(200) DEFAULT NULL,
  `picture` text NOT NULL,
  `password` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT COLLATE utf8mb4_unicode_ci;
```
8 . you are all set, welcome to our golang world.
#
```
hit this Url to see the home page. -> http://localhost/
```
# Authentication Methods.
1. Cookie-Based authentication
2. OpenId


