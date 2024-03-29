# DROP TABLE `apply`;
CREATE TABLE `apply`
(
    `applyId`   bigint unsigned     NOT NULL AUTO_INCREMENT,
    `userId`    bigint unsigned     NOT NULL DEFAULT '0' COMMENT '用户id',
    `scholarId` varchar(255)     NOT NULL DEFAULT '0' COMMENT '学者id',
    `status`    tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '处理状态',
    `applyType` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '认证类型',
    `email`     varchar(255)        NULL COMMENT '认证邮箱',
    `url`       varchar(255)        NULL COMMENT '认证证件',
    `applyTime` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`applyId`),
    KEY `idx_userId` (`userId`),
    KEY `idx_scholarId` (`scholarId`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;