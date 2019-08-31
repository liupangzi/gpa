-- gpa mysql model -s=test/mysql/sql/customer.sql -t=test/mysql/model/customer.go -p=test_mysql_model && gofmt -w test/mysql/model/customer.go

CREATE TABLE `customer`
(
    `id`          SERIAL PRIMARY KEY,
    `card_id`     SMALLINT     NOT NULL DEFAULT 0 COMMENT 'card id',
    `member_id`   MEDIUMINT             DEFAULT 0 COMMENT 'member id',
    `order`       INT          not null default 12 COMMENT 'order by',
    `name`        VARCHAR(255) NOT NULL,
    `email`       VARCHAR(255) NOT NULL,
    `address`     text                  DEFAULT NULL,
    `create_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (`card_id`),
    INDEX (email)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;