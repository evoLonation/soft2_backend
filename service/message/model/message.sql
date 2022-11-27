DROP TABLE `message`;
CREATE TABLE `message`
(
    `msgId`			bigint unsigned		NOT NULL AUTO_INCREMENT,
    `receiverId`	bigint unsigned		NOT NULL DEFAULT '0' COMMENT '接收者Id',
    `content`		varchar(255)        NOT NULL DEFAULT '' COMMENT '消息内容',
    `messageType`	tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '消息类型',
    `read`			boolean 			NOT NULL DEFAULT false COMMENT '消息状态',
    `msgTime`		timestamp			NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `result`		tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '处理结果',
    `uId`		    bigint unsigned		NULL COMMENT '用户id',
    `gId`		    bigint unsigned		NULL COMMENT '误认领申诉Id',
    `pId`		    bigint unsigned		NULL COMMENT '文献id',
    `rId`		    bigint unsigned		NULL COMMENT '文献互助Id',
    PRIMARY KEY (`msgId`),
    KEY `idx_receiverId` (`receiverId`),
    KEY `idx_uId` (`uId`),
    KEY `idx_gId` (`gId`),
    KEY `idx_pId` (`pId`),
    KEY `idx_rId` (`rId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;