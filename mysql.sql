-- 创建数据库
CREATE DATABASE IF NOT EXISTS `go1_admin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `go1_admin`;

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
                                       `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    `username` varchar(64) NOT NULL COMMENT '用户名',
    `password` varchar(128) NOT NULL COMMENT '密码',
    `nickname` varchar(128) DEFAULT NULL COMMENT '昵称',
    `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
    `role` varchar(64) DEFAULT 'user' COMMENT '角色',
    `status` tinyint(1) DEFAULT '1' COMMENT '状态：0-禁用 1-启用',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`),
    KEY `idx_email` (`email`),
    KEY `idx_deleted_at` (`deleted_at`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 操作日志表
CREATE TABLE IF NOT EXISTS `operation_logs` (
                                                `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户ID',
    `ip` varchar(64) DEFAULT NULL COMMENT 'IP地址',
    `method` varchar(20) DEFAULT NULL COMMENT '请求方法',
    `path` varchar(255) DEFAULT NULL COMMENT '请求路径',
    `status` int(11) DEFAULT NULL COMMENT '状态码',
    `latency` bigint(20) DEFAULT NULL COMMENT '延迟时间(ms)',
    `user_agent` varchar(255) DEFAULT NULL COMMENT '用户代理',
    `request` text COMMENT '请求内容',
    `response` text COMMENT '响应内容',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_created_at` (`created_at`),
    KEY `idx_deleted_at` (`deleted_at`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志表';

-- 文件上传表
CREATE TABLE IF NOT EXISTS `uploads` (
                                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '上传用户ID',
    `file_name` varchar(255) NOT NULL COMMENT '文件名',
    `file_path` varchar(255) NOT NULL COMMENT '文件路径',
    `file_size` bigint(20) DEFAULT NULL COMMENT '文件大小',
    `file_type` varchar(64) DEFAULT NULL COMMENT '文件类型',
    `status` tinyint(1) DEFAULT '1' COMMENT '状态：0-禁用 1-启用',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_created_at` (`created_at`),
    KEY `idx_deleted_at` (`deleted_at`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件上传表';

-- 插入管理员用户（密码：admin123）
INSERT INTO `users` (`created_at`, `updated_at`, `username`, `password`, `nickname`, `email`, `role`, `status`)
VALUES
    (NOW(), NOW(), 'admin', '$2a$10$YL0xEkHwLyH5kKyT1LS8aOHU8WkRe/su9RCl7O9zyJXv7QE8.zszy', '系统管理员', 'admin@example.com', 'admin', 1);

-- 插入测试用户（密码：test123）
INSERT INTO `users` (`created_at`, `updated_at`, `username`, `password`, `nickname`, `email`, `role`, `status`)
VALUES
    (NOW(), NOW(), 'test', '$2a$10$YL0xEkHwLyH5kKyT1LS8aOHU8WkRe/su9RCl7O9zyJXv7QE8.zszy', '测试用户', 'test@example.com', 'user', 1);

-- 创建触发器，在插入用户时自动生成昵称（如果为空）
DELIMITER //
CREATE TRIGGER `users_before_insert` BEFORE INSERT ON `users`
    FOR EACH ROW
BEGIN
    IF NEW.nickname IS NULL THEN
        SET NEW.nickname = CONCAT('用户_', SUBSTRING(MD5(RAND()), 1, 8));
END IF;
END;//
DELIMITER ;