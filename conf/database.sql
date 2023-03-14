-- 创建数据库
CREATE DATABASE IF NOT EXISTS microBlog;

-- 删除用户信息表
-- DROP TABLE IF EXISTS `user_info`;

-- 创建用户信息表
CREATE TABLE IF NOT EXISTS `user_info` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '用户id',
  `name` VARCHAR(30) COMMENT '用户姓名',
  `sex` ENUM('男','女','其他') COMMENT '性别',
  `birth_day` DATETIME COMMENT '出生日期',
  `phone_number` VARCHAR(30) COMMENT '手机号码',
  `email` VARCHAR(30) COMMENT '1表示由系统创建',
  `address` VARCHAR(100) COMMENT '地址',
  `create_by` INT COMMENT '1表示由系统创建',
  `create_at` DATETIME NOT NULL DEFAULT (NOW()) COMMENT '创建时间',
  `update_at` DATETIME COMMENT '更新账户信息时间',
  `recent_login` DATETIME COMMENT '最近登录时间',
  `header_image` VARCHAR(100) COMMENT '头像图片地址',
  `profile` VARCHAR(100) COMMENT '用户个人简介'
);
-- 插入用户信息
-- INSERT INTO  `user_info`( name, createBy, update_at) VALUES ( "bobo", 1, NOW());

-- 创建账号信息表
CREATE TABLE IF NOT EXISTS `account_info` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '账号id',
  `user_id` INT COMMENT '用户id',
  `account_name` VARCHAR(30) COMMENT '账号名称',
  `password` VARCHAR(30) COMMENT '账号密码',
  `create_at` DATETIME NOT NULL DEFAULT (NOW()) COMMENT '账号创建时间',
  `update_at` DATETIME COMMENT '更新账户信息时间',
  `recent_login` DATETIME COMMENT '最近登录时间',
  `status` INT DEFAULT (0) COMMENT '用户账号状态'
);