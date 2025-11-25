-- Create "mail" table
CREATE TABLE `mail` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `subject` text NOT NULL,
  `header` text NULL,
  `body` text NOT NULL,
  `footer` text NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
