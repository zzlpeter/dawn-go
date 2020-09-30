```
修改conf/conf.toml配置文件
执行SQL创建表
CREATE TABLE `permission`.`t_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` varchar(28) NOT NULL,
  `role_desc` varchar(200) NOT NULL,
  `insert_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `site_url` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

CREATE TABLE `permission`.`t_permission`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `permission_id` varchar(32) NOT NULL,
  `permission_desc` varchar(200) NOT NULL,
  `role_id` varchar(32) NOT NULL,
  `insert_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `permission_id_UNIQUE`(`role_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

CREATE TABLE `task`.`task`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键自增',
  `task_key` varchar(63) NOT NULL COMMENT '任务唯一标识',
  `execute_func` varchar(255) NOT NULL COMMENT '任务执行的方法',
  `trigger` varchar(255) CHARACTER NOT NULL COMMENT '调度方式(cron/date/interval)',
  `spec` varchar(255) CHARACTER NOT NULL COMMENT '调度时间',
  `args` varchar(255) CHARACTER NOT NULL DEFAULT '' COMMENT '执行方法的args参数',
  `is_valid` int(11) NOT NULL COMMENT '是否有效',
  `status` varchar(31) NOT NULL COMMENT '执行状态(ready/doing)',
  `extra` varchar(255) NOT NULL DEFAULT '{}' COMMENT '额外信息',
  `create_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务描述',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ux__task__task_key_status`(`task_key`, `status`) USING BTREE COMMENT 'task_key/status唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务表' ROW_FORMAT = Compact;

CREATE TABLE `task`.`task_execute`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键自增',
  `task_id` int(11) NOT NULL COMMENT '任务ID',
  `status` varchar(255) NOT NULL COMMENT '执行状态',
  `extra` text COMMENT '额外信息',
  `create_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `trace_id` varchar(255) DEFAULT '' COMMENT 'trace id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx__execute_task__task_id_create_at`(`task_id`, `create_at`) USING BTREE COMMENT '任务ID、创建时间'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务执行表' ROW_FORMAT = Compact;

安装依赖包
git mod vendor 

执行main函数
go run main.go

访问swagger
http://localhost:8001/swagger/index.html
```