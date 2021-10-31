--------------------
---- User table ----
--------------------
USE filestore;
DROP TABLE `user` IF EXISTS;
CREATE TABLE `user` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "用户id",
    `user_name` VARCHAR(50) NOT NULL COMMENT "用户名",
    `email` varchar(50) NULL DEFAULT '' COMMENT "用户邮箱",
    `phone` varchar(30) NULL DEFAULT '' COMMENT "联系方式",
    `avatar` Text NULL COMMENT "用户头像",
    `profile` Text NULL COMMENT "用户简介",
    `status` SMALLINT NOT NULL COMMENT "用户状态, 0表示启用，1表示禁用，2表示锁定，3表示标记删除",
    `created_at` DATETIME COMMENT '创建时间',
    `updated_at` DATETIME COMMENT '更新时间'
) ENGINE=InnoDB CHARSET=utf8mb4;

