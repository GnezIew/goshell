CREATE TABLE `question_teacher_user` (
                                         `id` int unsigned NOT NULL AUTO_INCREMENT,
                                         `question_teacher_id` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
                                         `teacher_name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '老师名称',
                                         `account` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '账号',
                                         `password` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '密码',
                                         `is_available` tinyint NOT NULL DEFAULT '1' COMMENT '是否启用 0 停用 1启用',
                                         `auth_content_status` tinyint NOT NULL DEFAULT '0' COMMENT '授权内容状态： 0 无授权内容 、 1 使用中 2 全部已停用',
                                         `channel_batch` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '第一次授权渠道批次',
                                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                         PRIMARY KEY (`id`),
                                         UNIQUE KEY `question_teacher_id` (`question_teacher_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='老师表\n';