-- ----------------------------
-- 1. 网关访问日志表 (Access Logs)
-- 用于统计流量、分析错误、审计操作
-- ----------------------------
CREATE TABLE `gateway_access_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `trace_id` varchar(64) DEFAULT '' COMMENT '链路追踪ID',
  `user_id` bigint unsigned DEFAULT '0' COMMENT '用户ID(从Token解析)',
  `ip` varchar(64) DEFAULT '' COMMENT '访问IP',
  `method` varchar(10) DEFAULT '' COMMENT '请求方法',
  `path` varchar(255) DEFAULT '' COMMENT '请求路径',
  `query` text COMMENT '查询参数',
  `body` text COMMENT '请求体(可选截取)',
  `status_code` int DEFAULT '200' COMMENT 'HTTP状态码',
  `latency` bigint DEFAULT '0' COMMENT '耗时(毫秒)',
  `error_msg` varchar(255) DEFAULT '' COMMENT '错误信息',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '请求时间',
  PRIMARY KEY (`id`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_trace_id` (`trace_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网关访问日志表';

-- ----------------------------
-- 2. 黑名单表 (Blacklist)
-- 用于网关层快速拦截恶意 IP 或用户，无需穿透到业务服务
-- ----------------------------
CREATE TABLE `gateway_blacklist` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型 1:IP 2:UserID',
  `value` varchar(64) NOT NULL COMMENT '黑名单值(IP或ID)',
  `reason` varchar(255) DEFAULT '' COMMENT '封禁原因',
  `expire_at` datetime DEFAULT NULL COMMENT '过期时间(NULL为永久)',
  `enable` tinyint NOT NULL DEFAULT '1' COMMENT '是否生效',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_type_val` (`type`, `value`),
  KEY `idx_expire_at` (`expire_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网关黑名单表';
