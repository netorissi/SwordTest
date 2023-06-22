CREATE TABLE IF NOT EXISTS `tasks` (
    `id` int(11) auto_increment primary key,
    `user_id` int(11) NOT NULL,
    `summary` longtext NOT NULL,
    `completed_at` datetime NULL,
    FOREIGN KEY (`user_id`) REFERENCES users (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;