# DROP TABLE `verifycode`;
CREATE TABLE `verifycode`
(
    `email`       varchar(255) NOT NULL COMMENT '认证邮箱',
    `code`        varchar(255) NOT NULL DEFAULT '' COMMENT '验证码',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;