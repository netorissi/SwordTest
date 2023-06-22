BEGIN;
CREATE TABLE IF NOT EXISTS `sessions` (
    `id` int(11) auto_increment primary key,
    `user_id` int(11) NOT NULL,
    `access_token` varchar(255) NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES users (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;

INSERT INTO `sessions` (`user_id`, `access_token`) VALUES (1, "example_token_manager");
INSERT INTO `sessions` (`user_id`, `access_token`) VALUES (2, "example_token_tech_1");
INSERT INTO `sessions` (`user_id`, `access_token`) VALUES (3, "example_token_tech_2");
COMMIT;