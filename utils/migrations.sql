USE testdb;

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_profile` (
  `user_id` bigint(20) NOT NULL,
  `first_name` varchar(32) NOT NULL,
  `last_name` varchar(64) NOT NULL,
  `phone` varchar(64) NOT NULL,
  `address` varchar(64) NOT NULL,
  `city` varchar(64) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_data` (
  `user_id` bigint(20) NOT NULL,
  `school` varchar(32) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `auth` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `api-key` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `user_profile`
ADD CONSTRAINT `fk_user_profile_user`
FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `user_data`
ADD CONSTRAINT `fk_user_data_user`
FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
ON DELETE CASCADE ON UPDATE CASCADE;

INSERT INTO `auth` VALUES (1,'www-dfq92-sqfwf'),(2,'ffff-2918-xcas');
INSERT INTO `user` VALUES (1,'test'),(2,'admin'),(3,'guest');
INSERT INTO `user_data` VALUES (1,'гімназія №179 міста Києва'),(2,'ліцей №227'),(3,'Медична гімназія №33 міста Києва');
INSERT INTO `user_profile` VALUES (1,'Олександр','Шкільний','+38050123455','вул. Сибірська 2','Київ'),(2,'Дмитро','Кавун','+38065133223','вул. Біла 4','Харків'),(3,'Василь','Шпак','+38055221166','вул. Срібляста 5','Житомир');