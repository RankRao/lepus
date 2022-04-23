
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Records of alarm_channels
-- ----------------------------
INSERT INTO `alarm_channels` VALUES (1, '默认渠道', 0, '', 1, '', '默认事件发送渠道', 1, '2021-02-27 20:28:33', '2022-04-22 11:20:56');
INSERT INTO `alarm_channels` VALUES (2, '紧急渠道', 1, '', 1, '', '紧急事件报警渠道', 1, '2021-03-20 16:08:21', '2022-04-22 11:20:59');

-- ----------------------------
-- Records of alarm_levels
-- ----------------------------
INSERT INTO `alarm_levels` VALUES (1, '通知', '用于一般常规通知的级别', 1, '2021-02-27 20:28:33', '2022-03-28 21:11:45');
INSERT INTO `alarm_levels` VALUES (2, '提醒', '用于一般提醒通知的级别', 1, '2021-03-20 16:08:21', '2022-03-19 13:22:29');
INSERT INTO `alarm_levels` VALUES (3, '警告', '用于警告信息通知的级别', 1, '2021-04-14 22:54:41', '2022-03-19 13:23:03');
INSERT INTO `alarm_levels` VALUES (4, '紧急', '用于紧急信息通知的级别', 1, '2021-04-14 22:55:17', '2022-03-19 13:23:15');
INSERT INTO `alarm_levels` VALUES (5, '严重', '用于严重故障通知的级别', 1, '2021-04-14 22:55:59', '2022-03-19 13:23:38');

-- ----------------------------
-- Records of alarm_rules
-- ----------------------------
INSERT INTO `alarm_rules` VALUES (1, 'MySQL QPS过高', 'MySQL', '', 'queries', '', '>', '0', 21600, 2, 1, 3, 1, '2021-02-16 10:54:55', '2022-04-22 10:39:41');
INSERT INTO `alarm_rules` VALUES (2, 'MySQL数据库无法连接', 'MySQL', '', 'connect', '', '!=', '1', 10800, 1, 1, 2, 2, '2021-02-27 20:00:05', '2022-04-21 16:26:13');
INSERT INTO `alarm_rules` VALUES (4, 'MySQL活动事务量过高', 'MySQL', '', 'ActiveTrx', '', '>', '10', 21600, 2, 1, 2, 1, '2021-03-26 23:27:58', '2022-04-20 11:14:22');
INSERT INTO `alarm_rules` VALUES (5, 'MySQL活动会话数过高', 'MySQL', '', 'threadsRunning', '', '>', '20', 10800, 2, 1, 3, 1, '2021-03-29 22:33:34', '2022-04-20 11:14:27');
INSERT INTO `alarm_rules` VALUES (6, 'MySQL连接数过高', 'MySQL', '', 'threadsConnected', '', '>', '1000', 10800, 2, 1, 3, 1, '2021-04-02 23:39:52', '2022-04-20 11:14:33');
INSERT INTO `alarm_rules` VALUES (11, '网站响应超时', 'Web', '', 'resp_status', '', '!=', '1', 10800, 1, 1, 5, 1, '2022-01-26 23:10:06', '2022-04-21 16:26:35');
INSERT INTO `alarm_rules` VALUES (12, '网站状态码异常', 'Web', '', 'http_code', '', '!=', '200', 10800, 1, 1, 5, 2, '2022-01-26 23:11:27', '2022-04-21 16:26:17');
INSERT INTO `alarm_rules` VALUES (13, '网站响应时间异常', 'Web', '', 'http_time', '', '>', '8000', 10800, 1, 1, 2, 1, '2022-01-26 23:13:23', '2022-04-20 11:17:54');
INSERT INTO `alarm_rules` VALUES (20, 'GreatSQL QPS过高', 'GreatSQL', NULL, 'queries', NULL, '>', '1000', 21600, 2, 1, 3, 1, '2022-04-20 15:29:26', '2022-04-20 15:31:46');
INSERT INTO `alarm_rules` VALUES (21, 'GreatSQL数据库无法连接', 'GreatSQL', NULL, 'connect', NULL, '!=', '1', 10800, 1, 1, 2, 2, '2022-04-20 15:29:26', '2022-04-21 16:26:22');
INSERT INTO `alarm_rules` VALUES (22, 'GreatSQL活动事务量过高', 'GreatSQL', NULL, 'ActiveTrx', NULL, '>', '10', 21600, 2, 1, 2, 1, '2022-04-20 15:29:26', '2022-04-20 15:31:50');
INSERT INTO `alarm_rules` VALUES (23, 'GreatSQL活动会话数过高', 'GreatSQL', NULL, 'threadsRunning', NULL, '>', '20', 10800, 2, 1, 3, 1, '2022-04-20 15:29:26', '2022-04-20 15:31:53');
INSERT INTO `alarm_rules` VALUES (24, 'GreatSQL连接数过高', 'GreatSQL', NULL, 'threadsConnected', NULL, '>', '1000', 10800, 2, 1, 3, 1, '2022-04-20 15:29:26', '2022-04-20 15:31:55');
INSERT INTO `alarm_rules` VALUES (35, 'MongoDB数据库无法连接', 'MongoDB', '', 'connect', '', '!=', '1', 21600, 1, 1, 5, 2, '2022-04-20 15:36:06', '2022-04-21 16:26:41');
INSERT INTO `alarm_rules` VALUES (36, 'MongoDB当前连接数过高', 'MongoDB', '', 'connectionsCurrent', '', '>', '500', 21600, 2, 1, 3, 1, '2022-04-20 15:37:35', '2022-04-20 15:37:35');
INSERT INTO `alarm_rules` VALUES (37, 'MongoDB可用连接不足', 'MongoDB', '', 'connectionsAvailable', '', '<=', '100', 10800, 2, 1, 4, 1, '2022-04-20 15:38:32', '2022-04-21 16:26:50');
INSERT INTO `alarm_rules` VALUES (38, 'MongoDB网络请求数过高', 'MongoDB', '', 'networkNumRequests', '', '>', '1000', 10800, 2, 1, 3, 1, '2022-04-20 15:40:34', '2022-04-20 15:40:34');
INSERT INTO `alarm_rules` VALUES (39, 'MongoDB写入流量异常', 'MongoDB', '', 'networkBytesIn', '', '>', '10000000', 21600, 2, 1, 3, 1, '2022-04-20 15:42:17', '2022-04-20 15:42:17');
INSERT INTO `alarm_rules` VALUES (40, 'MongoDB操作量过高', 'MongoDB', '', 'opcounters', '', '>', '1000', 3600, 2, 1, 3, 1, '2022-04-20 16:06:39', '2022-04-20 16:06:39');
INSERT INTO `alarm_rules` VALUES (41, 'Oracle数据库连接失败', 'Oracle', '', 'connect', '', '!=', '1', 10800, 1, 1, 5, 2, '2022-04-20 16:09:56', '2022-04-21 16:26:55');
INSERT INTO `alarm_rules` VALUES (42, 'Oracle会话数过高', 'Oracle', '', 'sessionTotal', '', '>', '1000', 10800, 2, 1, 3, 1, '2022-04-20 16:15:55', '2022-04-20 16:15:55');
INSERT INTO `alarm_rules` VALUES (43, 'Oracle活动会话数过高', 'Oracle', '', 'sessionActive', '', '>', '30', 10800, 2, 1, 3, 1, '2022-04-20 16:18:49', '2022-04-20 16:18:49');
INSERT INTO `alarm_rules` VALUES (45, 'PostgreSQL数据库无法连接', 'PostgreSQL', '', 'connect', '', '!=', '1', 21600, 1, 1, 5, 2, '2022-04-20 16:26:21', '2022-04-21 16:27:00');
INSERT INTO `alarm_rules` VALUES (46, 'PostgreSQL连接数过高', 'PostgreSQL', '', 'connections', '', '>', '1000', 10800, 2, 1, 3, 1, '2022-04-20 16:27:30', '2022-04-20 16:27:30');
INSERT INTO `alarm_rules` VALUES (47, 'PostgreSQL活动SQL数过高', 'PostgreSQL', '', 'activeSQL', '', '>', '30', 10800, 2, 1, 3, 1, '2022-04-20 16:29:45', '2022-04-20 16:29:45');
INSERT INTO `alarm_rules` VALUES (50, 'PostgreSQL检测到死锁', 'PostgreSQL', '', 'deadlocks', '', '>', '0', 10800, 2, 1, 3, 1, '2022-04-20 16:42:25', '2022-04-20 16:42:45');
INSERT INTO `alarm_rules` VALUES (53, 'Redis数据库无法连接', 'Redis', '', 'connect', '', '!=', '1', 10800, 2, 1, 5, 1, '2022-04-20 21:00:02', '2022-04-20 21:04:14');
INSERT INTO `alarm_rules` VALUES (55, 'Redis每秒请求数过高', 'Redis', '', 'instantaneousOpsPerSec', '', '>', '2000', 10800, 2, 1, 3, 1, '2022-04-20 21:02:25', '2022-04-20 21:03:15');
INSERT INTO `alarm_rules` VALUES (56, 'Redis连接数过高', 'Redis', '', 'connectedClients', '', '>', '1000', 10800, 2, 1, 3, 1, '2022-04-20 21:04:58', '2022-04-20 21:04:58');
INSERT INTO `alarm_rules` VALUES (57, 'Redis连接阻塞数过高', 'Redis', '', 'blockedClients', '', '>', '10', 10800, 2, 1, 3, 1, '2022-04-20 21:06:56', '2022-04-20 21:06:56');

-- ----------------------------
-- Records of meta_modules
-- ----------------------------
INSERT INTO `meta_modules` VALUES (1, 'MySQL', '', '2021-04-18 22:35:09', '2021-04-19 21:47:40');
INSERT INTO `meta_modules` VALUES (2, 'Oracle', '', '2021-04-19 21:47:49', '2021-04-19 21:47:49');
INSERT INTO `meta_modules` VALUES (3, 'PostgreSQL', '', '2021-04-19 21:48:04', '2021-04-19 21:48:04');
INSERT INTO `meta_modules` VALUES (4, 'Redis', '', '2021-04-19 21:48:25', '2021-04-19 21:48:25');
INSERT INTO `meta_modules` VALUES (5, 'MongoDB', '', '2021-04-19 21:48:38', '2021-04-19 21:48:38');
INSERT INTO `meta_modules` VALUES (6, 'SQLServer', '', '2021-04-19 21:48:56', '2021-04-19 21:48:56');
INSERT INTO `meta_modules` VALUES (7, 'ClickHouse', '', '2021-04-19 21:49:15', '2021-04-19 21:49:15');
INSERT INTO `meta_modules` VALUES (8, 'Web', '', '2022-01-25 21:56:12', '2022-01-25 21:56:12');
INSERT INTO `meta_modules` VALUES (9, 'GreatSQL', '', '2022-03-18 17:32:41', '2022-04-21 18:17:10');

-- ----------------------------
-- Records of task
-- ----------------------------
INSERT INTO `task` VALUES (2, 1, 'Redis Health Check', './lepus_redis_mon', '', 'period', '', 30, '2022-04-21 17:39:33', 0, 'waiting', 1, NULL, 0, '2021-03-14 21:30:22', '2022-04-21 17:40:37');
INSERT INTO `task` VALUES (3, 1, 'Web Health Check', './lepus_web_mon', '', 'period', '', 180, '2022-04-21 17:39:49', 0, 'waiting', 1, NULL, 0, '2022-01-26 22:56:54', '2022-04-21 17:39:41');
INSERT INTO `task` VALUES (4, 1, 'MongoDB Health Check', './lepus_mongo_mon', '', 'period', '', 60, '2022-04-21 17:40:07', 0, 'waiting', 1, NULL, 0, '2022-01-29 15:33:00', '2022-04-21 17:40:39');
INSERT INTO `task` VALUES (5, 1, 'Oracle Health Check', './lepus_oracle_mon', '', 'period', '', 120, '2022-04-21 17:40:50', 0, 'waiting', 1, NULL, 0, '2022-03-18 19:02:41', '2022-04-21 17:39:43');
INSERT INTO `task` VALUES (6, 1, 'GreatSQL Health Check', './lepus_greatsql_mon', '', 'period', '', 120, '2022-04-21 17:40:51', 0, 'waiting', 1, NULL, 0, '2022-04-12 19:39:05', '2022-04-21 17:39:44');
INSERT INTO `task` VALUES (7, 1, 'PostgreSQL Health Check', './lepus_postgres_mon', '', 'period', '', 120, '2022-04-21 17:40:53', 0, 'waiting', 1, NULL, 0, '2022-04-18 12:36:43', '2022-04-21 17:39:45');
INSERT INTO `task` VALUES (8, 1, 'MySQL Health Check', './lepus_mysql_mon', '', 'period', '', 30, '2022-04-21 17:40:12', 0, 'waiting', 1, NULL, 0, '2021-03-13 21:01:28', '2022-04-21 17:40:40');

-- ----------------------------
-- Records of task_type
-- ----------------------------
INSERT INTO `task_type` VALUES (1, 0, '监控采集', NULL, '2021-03-13 20:53:30', '2021-04-15 10:42:49');

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2021-02-26 22:54:13.353', '2021-04-18 09:53:30.822', NULL, 'guest', 'guest', '7b01f76a3c1ef2114eff7844e1e5d256', 0, 'guest');
INSERT INTO `users` VALUES (2, '2021-02-26 22:58:07.035', '2022-04-21 17:36:43.150', NULL, 'admin', 'Administrator', 'a8a0d32f1abefd3fa996321d5e72c6d6', 1, '');

SET FOREIGN_KEY_CHECKS = 1;
