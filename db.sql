CREATE TABLE `comments` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime DEFAULT NULL,
                            `updated_at` datetime DEFAULT NULL,
                            `deleted_at` datetime DEFAULT NULL,
                            `content` longtext,
                            `user_id` bigint(20) unsigned DEFAULT NULL,
                            `post_id` bigint(20) unsigned DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `idx_comments_deleted_at` (`deleted_at`),
                            KEY `fk_posts_comments` (`post_id`),
                            KEY `fk_users_comments` (`user_id`),
                            CONSTRAINT `fk_posts_comments` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
                            CONSTRAINT `fk_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE `posts` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime DEFAULT NULL,
                         `updated_at` datetime DEFAULT NULL,
                         `deleted_at` datetime DEFAULT NULL,
                         `title` longtext,
                         `content` longtext,
                         `user_id` bigint(20) unsigned DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `idx_posts_deleted_at` (`deleted_at`),
                         KEY `fk_users_posts` (`user_id`),
                         CONSTRAINT `fk_users_posts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime DEFAULT NULL,
                         `updated_at` datetime DEFAULT NULL,
                         `deleted_at` datetime DEFAULT NULL,
                         `username` varchar(191) DEFAULT NULL,
                         `password` longtext,
                         `email` varchar(191) DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uni_users_email` (`email`),
                         UNIQUE KEY `uni_users_username` (`username`),
                         KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
