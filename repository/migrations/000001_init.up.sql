BEGIN;
CREATE TABLE IF NOT EXISTS `users` (
    `id` int(11) auto_increment primary key,
    `role` VARCHAR(25) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1;

INSERT INTO `users` (`role`) VALUES ("manager");
INSERT INTO `users` (`role`) VALUES ("tech");
INSERT INTO `users` (`role`) VALUES ("tech");
COMMIT;