-- 用户基础信息表 
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `account_num` bigint(20) NOT NULL COMMENT '账号(数字)',
  `name` varchar(64) DEFAULT '' COMMENT '昵称',
  `email` varchar(128) DEFAULT '' COMMENT '邮箱',
  `phone` varchar(20) DEFAULT '' COMMENT '手机号',
  `avatar_url` varchar(255) DEFAULT '' COMMENT '头像URL',
  `desc` varchar(255) DEFAULT '' COMMENT '个人简介',
  `password` varchar(255) NOT NULL COMMENT '加密后的密码',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:冻结',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_account_num` (`account_num`),
  UNIQUE KEY `uk_email` (`email`),
  UNIQUE KEY `uk_phone` (`phone`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户基础信息表';

-- 用户扩展信息表 (合并档案和偏好设置)
CREATE TABLE IF NOT EXISTS `user_profile` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  
  -- 基本信息
  `real_name` varchar(64) DEFAULT '' COMMENT '真实姓名',
  `gender` tinyint(1) DEFAULT '0' COMMENT '性别 0:未知 1:男 2:女',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `location` varchar(128) DEFAULT '' COMMENT '所在地',
  `company` varchar(128) DEFAULT '' COMMENT '公司',
  `position` varchar(64) DEFAULT '' COMMENT '职位',
  `bio` text COMMENT '个人介绍',
  
  -- 联系方式扩展
  `github` varchar(100) DEFAULT '' COMMENT 'GitHub账号',
  `website` varchar(255) DEFAULT '' COMMENT '个人网站',
  
  -- 偏好设置 (JSON存储，灵活扩展)
  `preferences` json DEFAULT NULL COMMENT '用户偏好设置 {
    "theme": "light",
    "language": "zh-CN", 
    "timezone": "Asia/Shanghai",
    "notifications": {
      "email": true,
      "push": true
    },
    "workflow": {
      "autoSave": true,
      "defaultView": "list"
    }
  }',
  
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_id` (`user_id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户扩展信息表';

-- 用户会话表 (合并安全控制和会话管理)
CREATE TABLE IF NOT EXISTS `user_session` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `session_id` varchar(128) NOT NULL COMMENT '会话ID',
  
  -- 登录信息
  `ip_address` varchar(64) NOT NULL COMMENT 'IP地址',
  `user_agent` text COMMENT 'User-Agent',
  `device_info` varchar(255) DEFAULT '' COMMENT '设备信息',
  
  -- 时间管理
  `login_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
  `last_activity` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最后活动时间',
  `expires_at` datetime NOT NULL COMMENT '过期时间',
  `logout_time` datetime DEFAULT NULL COMMENT '登出时间',
  
  -- 状态控制
  `is_active` tinyint(1) DEFAULT '1' COMMENT '是否活跃',
  `logout_reason` varchar(32) DEFAULT '' COMMENT '登出原因',
  
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_session_id` (`session_id`),
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_expires_at` (`expires_at`),
  INDEX `idx_last_activity` (`last_activity`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户会话表';

-- 用户安全设置表 (登录安全控制和MFA)
CREATE TABLE IF NOT EXISTS `user_security` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  
  -- 登录安全控制
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(64) DEFAULT '' COMMENT '最后登录IP',
  `failed_login_attempts` int DEFAULT '0' COMMENT '连续失败登录次数',
  `account_locked_until` datetime DEFAULT NULL COMMENT '账户锁定到何时',
  
  -- MFA设置
  `mfa_enabled` tinyint(1) DEFAULT '0' COMMENT 'MFA是否启用',
  `mfa_type` varchar(32) DEFAULT '' COMMENT 'MFA类型 totp/email/sms',
  `mfa_secret` varchar(255) DEFAULT '' COMMENT 'MFA密钥',
  `backup_codes` json DEFAULT NULL COMMENT '备用验证码',
  
  -- 密码安全
  `password_updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '密码最后更新时间',
  `require_password_change` tinyint(1) DEFAULT '0' COMMENT '需要修改密码',
  
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_id` (`user_id`),
  INDEX `idx_last_login` (`last_login_at`),
  INDEX `idx_locked_until` (`account_locked_until`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户安全设置表';

-- 用户角色表 (简化权限控制)
CREATE TABLE IF NOT EXISTS `user_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `role` varchar(32) NOT NULL DEFAULT 'user' COMMENT '角色 user/admin',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_id` (`user_id`),
  INDEX `idx_role` (`role`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色表';

-- 初始化默认角色数据
INSERT INTO `user_role` (`user_id`, `role`) 
SELECT `id`, 'user' FROM `user` WHERE 1=1 ON DUPLICATE KEY UPDATE `role` = `role`;

