DROP TABLE IF EXISTS `Users`;
CREATE TABLE `Users` (
	`id` int(6) unsigned NOT NULL AUTO_INCREMENT,
	`user` varchar(30) NOT NULL,
	`password` varchar(30) NOT NULL,
	`role` varchar(30) NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;

