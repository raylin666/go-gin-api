CREATE TABLE `srv_keyword_tests`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `keyword`     varchar(36) NOT NULL COMMENT '关键词',
    `res_content` varchar(60) NOT NULL COMMENT '返回内容',
    `created_at`  datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY uni_keyword (`keyword`) USING BTREE
) ENGINE=InnoDB COMMENT='关键词测试';