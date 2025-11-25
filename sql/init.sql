-- Create users table
CREATE TABLE IF NOT EXISTS `users`
(
    `user_id`        CHAR(36) PRIMARY KEY,
    `username`       VARCHAR(255) UNIQUE NOT NULL,
    `email`          VARCHAR(255) UNIQUE NOT NULL,
    `password`       TEXT                NOT NULL,
    `phone_number`   VARCHAR(10)         DEFAULT NULL,
    `account_status` TINYINT             NOT NULL DEFAULT 0 COMMENT 'User account status (0=inactive, 1=active)',
    `updated_at`     DATETIME            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at`     DATETIME            NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create semesters table
CREATE TABLE IF NOT EXISTS `semesters`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `name`       VARCHAR(255) NOT NULL,
    `start_date` DATETIME     NOT NULL,
    `end_date`   DATETIME     NOT NULL,
    `updated_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create courses table
CREATE TABLE IF NOT EXISTS `courses`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT,
    `course_id`   VARCHAR(255)  NOT NULL COMMENT 'Course identifier',
    `course_name` VARCHAR(255)  NOT NULL,
    `user_id`     CHAR(36)      NOT NULL,
    `description` TEXT          DEFAULT NULL,
    `lecturers`   TEXT          NOT NULL COMMENT 'Comma-separated lecturer names',
    `credits`     INT           NOT NULL,
    `gpa`         FLOAT         NOT NULL DEFAULT 0,
    `semester_id` INT           NOT NULL,
    `updated_at`  DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create tags table
CREATE TABLE IF NOT EXISTS `tags`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,
    `name`       VARCHAR(255)  NOT NULL,
    `color`      VARCHAR(7)    DEFAULT '#808080' COMMENT 'Hex color code for UI display',
    `updated_at` DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create course_tags junction table
CREATE TABLE IF NOT EXISTS `course_tags`
(
    `course_id` INT NOT NULL,
    `tag_id`    INT NOT NULL,
    PRIMARY KEY (`course_id`, `tag_id`)
);

-- Create reminders table
CREATE TABLE IF NOT EXISTS `reminders`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT,
    `title`       TEXT     NOT NULL,
    `description` TEXT     NOT NULL,
    `due_date`    DATE     NOT NULL,
    `due_time`    TIME     NOT NULL,
    `user_id`     CHAR(36) NOT NULL,
    `course_id`   INT      DEFAULT NULL,
    `type`        INT      NOT NULL DEFAULT 0 COMMENT 'Reminder type (0=course, 1=assignment)',
    `status`      INT      NOT NULL DEFAULT 0 COMMENT 'Reminder status (0=pending, 1=completed, 2=overdue)',
    `updated_at`  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create mail table
CREATE TABLE `mail`
(
    `id`         INT PRIMARY KEY,
    `subject`    TEXT     NOT NULL,
    `header`     TEXT,
    `body`       TEXT     NOT NULL,
    `footer`     TEXT,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- Create indexes
CREATE INDEX `users_index_email` ON `users` (`email`);
CREATE INDEX `users_index_username` ON `users` (`username`);
CREATE INDEX `semesters_index_start_date` ON `semesters` (`start_date`);
CREATE INDEX `courses_index_user_id` ON `courses` (`user_id`);
CREATE INDEX `courses_index_semester_id` ON `courses` (`semester_id`);
CREATE INDEX `courses_index_course_id` ON `courses` (`course_id`);
CREATE INDEX `tags_index_name` ON `tags` (`name`);
CREATE INDEX `course_tags_index_tag_id` ON `course_tags` (`tag_id`);
CREATE INDEX `reminders_index_user_id` ON `reminders` (`user_id`);

-- Add foreign keys
ALTER TABLE `courses`
    ADD CONSTRAINT `fk_courses_user_id`
    FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `courses`
    ADD CONSTRAINT `fk_courses_semester_id`
    FOREIGN KEY (`semester_id`) REFERENCES `semesters` (`id`) ON DELETE RESTRICT;

ALTER TABLE `course_tags`
    ADD CONSTRAINT `fk_course_tags_course_id`
    FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE;

ALTER TABLE `course_tags`
    ADD CONSTRAINT `fk_course_tags_tag_id`
    FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE;

ALTER TABLE `reminders`
    ADD CONSTRAINT `fk_reminders_user_id`
    FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE;

ALTER TABLE `reminders`
    ADD CONSTRAINT `fk_reminders_course_id`
    FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE SET NULL;
