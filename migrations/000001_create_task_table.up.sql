CREATE TABLE `tasks` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL COMMENT 'タスク名',
  `done` tinyint(1) unsigned NOT NULL COMMENT 'タスクステータス',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
