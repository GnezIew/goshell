CREATE TABLE `auth_policy` (
                               `id` int unsigned NOT NULL AUTO_INCREMENT,
                               `policy_id` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
                               `app_id` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
                               `policy_name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '策略名称',
                               `description` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '策略描述',
                               `action` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作: *全部操作、click点击操作、view查看操作',
                               `effect` varchar(10) COLLATE utf8mb4_bin NOT NULL DEFAULT 'deny' COMMENT '效果: 只有两种情况 allow允许、deny 拒绝',
                               `conditions` json NOT NULL COMMENT '条件：由三部分组成：condkey条件键、operate 运算符、condvalue 条件值',
                               `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `policy_id` (`policy_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='权限策略表';