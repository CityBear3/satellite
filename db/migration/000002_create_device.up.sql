CREATE TABLE IF NOT EXISTS `device`
(
    `id`          CHAR(26)     NOT NULL,
    `name`        VARCHAR(255) NOT NULL,
    `secret`      VARCHAR(255) NOT NULL,
    `client_id`   CHAR(26)     NOT NULL,
    `is_deleted`  BOOL                  DEFAULT FALSE,
    `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    CONSTRAINT `FK__device__client_id`
        FOREIGN KEY (`client_id`)
            REFERENCES `client` (`id`)
            ON DELETE CASCADE
            ON UPDATE RESTRICT
)
    ENGINE = InnoDB