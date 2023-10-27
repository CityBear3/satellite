CREATE TABLE IF NOT EXISTS `archive_event`
(
    `id`           CHAR(26) NOT NULL,
    `device_id`    CHAR(26) NOT NULL,
    `client_id`    CHAR(26) NOT NULL,
    `requested_at` DATETIME NOT NULL,
    `created_at`   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB