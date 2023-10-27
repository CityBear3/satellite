CREATE TABLE IF NOT EXISTS `archive`
(
    `id`               CHAR(26)    NOT NULL,
    `device_id`        CHAR(26)    NOT NULL,
    `archive_event_id` CHAR(26)    NOT NULL UNIQUE,
    `content_type`     VARCHAR(64) NOT NULL,
    `created_at`       DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `FK__image__device_id`
        FOREIGN KEY (`device_id`)
            REFERENCES `device` (`id`)
            ON DELETE CASCADE
            ON UPDATE RESTRICT
)
    ENGINE = InnoDB