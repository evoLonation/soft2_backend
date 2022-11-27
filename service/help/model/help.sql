CREATE TABLE `literature_request` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL COMMENT '求助者id',
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文献标题',
    `author` varchar(255)  NOT NULL DEFAULT '' COMMENT '作者名称',
    `magazine` varchar(255)  NOT NULL DEFAULT '' COMMENT '期刊名称',
    `link` char(5)  NOT NULL COMMENT '链接',
    `request_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '求助时间',
    `request_content` varchar(255)  NOT NULL DEFAULT '' COMMENT '求助描述',
    `wealth` bigint NOT NULL COMMENT '财富值',
    `request_status` bigint NULL DEFAULT 0 COMMENT '求助状态',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;
CREATE TABLE `literature_help` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL COMMENT '应助者id',
    `request_id` bigint NOT NULL COMMENT '求助id',
    `wealth` bigint NOT NULL COMMENT '财富值',
    `help_status` bigint NOT NULL COMMENT '应助状态',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;
CREATE TABLE `user_help` (
    `user_id` bigint NOT NULL COMMENT '用户id',
    `request` bigint NOT NULL COMMENT '求助次数',
    `help` bigint NOT NULL COMMENT '应助次数',
    `wealth` bigint DEFAULT 0 COMMENT '财富值',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;