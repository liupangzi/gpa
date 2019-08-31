-- gpa mysql model -s=test/mysql/sql/ticket.sql -t=test/mysql/model/ticket.go -p=test_mysql_model && gofmt -w test/mysql/model/ticket.go

CREATE TABLE `ticket`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `customer_id` bigint unsigned NOT NULL COMMENT 'customer id',
    `from`        int unsigned    NOT NULL,
    `to`          mediumint                default 10,
    `duration`    double          not null default 12.3,
    `meal`        bool                     default true,
    `climate`     char                     default 'l',
    price         float                    default 1.0 COMMENT '1 dollar',
    operation     varbinary(12),
    `tea`         enum ('Black','Green', 'Oolong', 'Pu-erh', 'White'),
    `token`       mediumtext comment 'A long long text',
    secret_key    tinyblob                 DEFAULT NULL COMMENT 'Whats your secret',
    `create_time` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (`from`),
    INDEX (customer_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;