USE filestore;

--------------------
---- User table ----
--------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "用户id",
    `username` VARCHAR(50) UNIQUE NOT NULL COMMENT "用户名",
    `password` VARCHAR(100) NOT NULL COMMENT "用户密码",
    `email` VARCHAR(50) NULL DEFAULT '' COMMENT "用户邮箱",
    `phone` VARCHAR(30) NULL DEFAULT '' COMMENT "联系方式",
    `avatar` Text NULL COMMENT "用户头像",
    `profile` Text NULL COMMENT "用户简介",
    `status` SMALLINT NOT NULL COMMENT "用户状态, 0表示启用，1表示禁用，2表示锁定，3表示标记删除",
    `created_at` DATETIME COMMENT '创建时间',
    `updated_at` DATETIME COMMENT '更新时间'
) ENGINE=InnoDB CHARSET=utf8mb4;


--------------------
--- folder table ---
--------------------
DROP TABLE IF EXISTS `folder`;
CREATE TABLE `folder` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "文件夹id",
    `foldername` VARCHAR(50) NOT NULL COMMENT "文件夹名称",
    `userid` INT NOT NULL COMMENT "用户id",
    `created_at` DATETIME COMMENT '创建时间',
    `updated_at` DATETIME COMMENT '更新时间',
    FOREIGN KEY(`userid`) REFERENCES user(`id`),
    UNIQUE INDEX(`foldername`, `userid`)
) ENGINE=InnoDB CHARSET=utf8mb4;


--------------------
---- file table ----
--------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "文件id",
    `filename` VARCHAR(255) NOT NULL UNIQUE COMMENT "文件名",
    `filesha1` VARCHAR(255) NOT NULL COMMENT "文件sha1",
    `filesize` INT NOT NULL COMMENT "文件大小",
    `location` VARCHAR(255) NOT NULL COMMENT "文件存储位置",
    `folderid` INT NOT NULL COMMENT "所属文件夹",
    `userid` INT NOT NULL COMMENT "所属用户",
    `created_at` DATETIME COMMENT '创建时间',
    `updated_at` DATETIME COMMENT '更新时间',
    FOREIGN KEY(`folderid`) REFERENCES folder(`id`),
    FOREIGN KEY(`userid`) REFERENCES user(`id`),
    UNIQUE INDEX(`folderid`, `userid`, `filesha1`)
) ENGINE=InnoDB CHARSET=utf8mb4;