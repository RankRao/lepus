

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for alarm_channels
-- ----------------------------
DROP TABLE IF EXISTS `alarm_channels`;
CREATE TABLE `alarm_channels`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `webhook_enable` int(11) NOT NULL DEFAULT 0,
  `webhook_url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `mail_enable` int(10) NOT NULL DEFAULT 0,
  `mail_list` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `description` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `enable` int(10) NOT NULL DEFAULT 1,
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for alarm_events
-- ----------------------------
DROP TABLE IF EXISTS `alarm_events`;
CREATE TABLE `alarm_events`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `alarm_title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `alarm_level` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `alarm_rule` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `alarm_value` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_uuid` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  `event_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_group` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_entity` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_value` decimal(20, 2) DEFAULT NULL,
  `event_unit` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_tag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `rule_id` int(10) DEFAULT NULL,
  `level_id` int(10) NOT NULL DEFAULT 0,
  `channel_id` int(10) NOT NULL DEFAULT 0,
  `send_mail` int(10) NOT NULL DEFAULT 0 COMMENT '0-未发送，1-成功，2-失败',
  `send_webhook` int(10) NOT NULL DEFAULT 0,
  `status` int(10) NOT NULL DEFAULT 0 COMMENT '0-未处理，1-处理中，2-已完成',
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_event_time`(`event_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3290 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for alarm_levels
-- ----------------------------
DROP TABLE IF EXISTS `alarm_levels`;
CREATE TABLE `alarm_levels`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `level_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `enable` smallint(2) NOT NULL DEFAULT 1,
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for alarm_rules
-- ----------------------------
DROP TABLE IF EXISTS `alarm_rules`;
CREATE TABLE `alarm_rules`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event_group` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_key` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event_entity` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `alarm_rule` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `alarm_value` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `alarm_sleep` int(10) NOT NULL DEFAULT 3600,
  `alarm_times` int(10) NOT NULL DEFAULT 3,
  `enable` int(10) NOT NULL DEFAULT 1,
  `level_id` int(10) NOT NULL DEFAULT 0,
  `channel_id` int(10) NOT NULL DEFAULT 0,
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_db
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_db`;
CREATE TABLE `dashboard_db`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `host` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `port` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `db_type` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `db_type_sort` tinyint(2) NOT NULL DEFAULT 0,
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `role` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `version` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `connect` tinyint(2) NOT NULL DEFAULT -1,
  `connect_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `sessions` tinyint(2) NOT NULL DEFAULT -1,
  `sessions_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `actives` tinyint(2) NOT NULL DEFAULT -1,
  `actives_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `waits` tinyint(2) NOT NULL DEFAULT -1,
  `waits_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `repl` tinyint(2) NOT NULL DEFAULT -1,
  `repl_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `repl_delay` tinyint(2) NOT NULL DEFAULT -1,
  `repl_delay_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `tablespace` tinyint(2) NOT NULL DEFAULT -1,
  `tablespace_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `snmp` tinyint(2) NOT NULL DEFAULT -1,
  `snmp_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `process` tinyint(2) NOT NULL DEFAULT -1,
  `process_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `load_1` tinyint(2) NOT NULL DEFAULT -1,
  `load_1_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `cpu` tinyint(2) NOT NULL DEFAULT -1,
  `cpu_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `network` tinyint(2) NOT NULL DEFAULT -1,
  `network_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `memory` tinyint(2) NOT NULL DEFAULT -1,
  `memory_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `disk` tinyint(2) NOT NULL DEFAULT -1,
  `disk_tips` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'no_data',
  `uptime_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_greatsql
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_greatsql`;
CREATE TABLE `dashboard_greatsql`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `host` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `port` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `connect` tinyint(2) NOT NULL DEFAULT 0,
  `hostname` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `uptime` int(10) NOT NULL DEFAULT -1,
  `version` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `timezone` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `readonly` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `max_connections` int(10) NOT NULL DEFAULT -1,
  `open_files_limit` int(10) NOT NULL DEFAULT -1,
  `open_files` int(10) NOT NULL DEFAULT -1,
  `table_open_cache` int(10) NOT NULL DEFAULT -1,
  `open_tables` int(10) NOT NULL DEFAULT -1,
  `threads_connected` int(10) NOT NULL DEFAULT -1,
  `threads_running` int(10) NOT NULL DEFAULT -1,
  `threads_created` int(10) NOT NULL DEFAULT -1,
  `threads_cached` int(10) NOT NULL DEFAULT -1,
  `connections` int(10) NOT NULL DEFAULT -1,
  `aborted_clients` int(10) NOT NULL DEFAULT -1,
  `aborted_connects` int(10) NOT NULL DEFAULT -1,
  `bytes_received` int(10) NOT NULL DEFAULT -1,
  `bytes_sent` int(10) NOT NULL DEFAULT -1,
  `com_select` int(10) NOT NULL DEFAULT -1,
  `com_insert` int(10) NOT NULL DEFAULT -1,
  `com_update` int(10) NOT NULL DEFAULT -1,
  `com_delete` int(10) NOT NULL DEFAULT -1,
  `com_commit` int(10) NOT NULL DEFAULT -1,
  `com_rollback` int(10) NOT NULL DEFAULT -1,
  `questions` int(10) NOT NULL DEFAULT -1,
  `queries` int(10) NOT NULL DEFAULT -1,
  `slow_queries` int(10) NOT NULL DEFAULT -1,
  `key_buffer_size` bigint(20) NOT NULL DEFAULT -1,
  `sort_buffer_size` bigint(20) NOT NULL DEFAULT -1,
  `join_buffer_size` bigint(20) NOT NULL DEFAULT -1,
  `innodb_pages_created` bigint(20) NOT NULL DEFAULT -1,
  `innodb_pages_read` bigint(20) NOT NULL DEFAULT -1,
  `innodb_pages_written` bigint(20) NOT NULL DEFAULT -1,
  `innodb_row_lock_current_waits` bigint(20) NOT NULL DEFAULT -1,
  `innodb_buffer_pool_read_requests` bigint(20) NOT NULL DEFAULT -1,
  `innodb_buffer_pool_write_requests` bigint(20) NOT NULL DEFAULT -1,
  `innodb_rows_read` bigint(20) NOT NULL DEFAULT -1,
  `innodb_rows_inserted` bigint(20) NOT NULL DEFAULT -1,
  `innodb_rows_updated` bigint(20) NOT NULL DEFAULT -1,
  `innodb_rows_deleted` bigint(20) NOT NULL DEFAULT -1,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_ip_port`(`host`, `port`) USING BTREE,
  INDEX `idx_gmt_create`(`gmt_create`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3468 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_mongodb
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_mongodb`;
CREATE TABLE `dashboard_mongodb`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `host` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `port` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `connect` smallint(6) NOT NULL DEFAULT 0,
  `ok` tinyint(2) NOT NULL DEFAULT -1,
  `uptime` int(11) NOT NULL DEFAULT -1,
  `version` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `connections_current` int(10) NOT NULL DEFAULT -1,
  `connections_available` int(10) NOT NULL DEFAULT -1,
  `mem_bits` smallint(4) NOT NULL DEFAULT -1,
  `mem_resident` int(10) NOT NULL DEFAULT -1,
  `mem_virtual` int(10) NOT NULL DEFAULT -1,
  `mem_supported` tinyint(1) NOT NULL DEFAULT -1,
  `mem_mapped` int(10) NOT NULL DEFAULT -1,
  `mem_mappedWithJournal` int(10) NOT NULL DEFAULT -1,
  `network_bytesIn` int(10) NOT NULL DEFAULT -1,
  `network_bytesOut` int(10) NOT NULL DEFAULT -1,
  `network_numRequests` int(10) NOT NULL DEFAULT -1,
  `opcounters_insert` int(10) NOT NULL DEFAULT -1,
  `opcounters_query` int(10) NOT NULL DEFAULT -1,
  `opcounters_update` int(10) NOT NULL DEFAULT -1,
  `opcounters_delete` int(10) NOT NULL DEFAULT -1,
  `opcounters_command` int(10) NOT NULL DEFAULT -1,
  `opcounters` int(10) NOT NULL DEFAULT -1,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`, `mem_mappedWithJournal`) USING BTREE,
  INDEX `idx_gmt_create`(`gmt_create`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 258087 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_mysql
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_mysql`;
CREATE TABLE `dashboard_mysql`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `host` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `port` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `connect` tinyint(2) NOT NULL DEFAULT 0,
  `hostname` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `uptime` int(10) NOT NULL DEFAULT -1,
  `version` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `timezone` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `readonly` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `max_connections` int(10) NOT NULL DEFAULT -1,
  `open_files_limit` int(10) NOT NULL DEFAULT -1,
  `open_files` int(10) NOT NULL DEFAULT -1,
  `table_open_cache` int(10) NOT NULL DEFAULT -1,
  `open_tables` int(10) NOT NULL DEFAULT -1,
  `threads_connected` int(10) NOT NULL DEFAULT -1,
  `threads_running` int(10) NOT NULL DEFAULT -1,
  `threads_created` int(10) NOT NULL DEFAULT -1,
  `threads_cached` int(10) NOT NULL DEFAULT -1,
  `connections` int(10) NOT NULL DEFAULT -1,
  `aborted_clients` int(10) NOT NULL DEFAULT -1,
  `aborted_connects` int(10) NOT NULL DEFAULT -1,
  `bytes_received` int(10) NOT NULL DEFAULT -1,
  `bytes_sent` int(10) NOT NULL DEFAULT -1,
  `com_select` int(10) NOT NULL DEFAULT -1,
  `com_insert` int(10) NOT NULL DEFAULT -1,
  `com_update` int(10) NOT NULL DEFAULT -1,
  `com_delete` int(10) NOT NULL DEFAULT -1,
  `com_commit` int(10) NOT NULL DEFAULT -1,
  `com_rollback` int(10) NOT NULL DEFAULT -1,
  `questions` int(10) NOT NULL DEFAULT -1,
  `queries` int(10) NOT NULL DEFAULT -1,
  `slow_queries` int(10) NOT NULL DEFAULT -1,
  `key_buffer_size` bigint(20) NOT NULL DEFAULT -1,
  `sort_buffer_size` bigint(20) NOT NULL DEFAULT -1,
  `join_buffer_size` bigint(20) NOT NULL DEFAULT -1,
  `innodb_pages_created` bigint(20) NOT NULL DEFAULT -1,
  `innodb_pages_read` bigint(20) NOT NULL DEFAULT -1,
  `innodb_pages_written` bigint(20) NOT NULL DEFAULT -1,
  `innodb_row_lock_current_waits` bigint(20) NOT NULL DEFAULT -1,
  `innodb_buffer_pool_read_requests` bigint(20) NOT NULL DEFAULT -1,
  `innodb_buffer_pool_write_requests` bigint(20) NOT NULL DEFAULT -1,
  `innodb_rows_read` bigint(20) NOT NULL DEFAULT -1,
  `innodb_rows_inserted` bigint(20) NOT NULL DEFAULT -1,
  `innodb_rows_updated` bigint(20) NOT NULL DEFAULT -1,
  `innodb_rows_deleted` bigint(20) NOT NULL DEFAULT -1,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_ip_port`(`host`, `port`) USING BTREE,
  INDEX `idx_gmt_create`(`gmt_create`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 861522 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_oracle
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_oracle`;
CREATE TABLE `dashboard_oracle`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `host` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `port` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `sid` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `connect` tinyint(2) NOT NULL DEFAULT 0,
  `instance_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `instance_role` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `instance_status` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `database_role` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `open_mode` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `protection_mode` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `host_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `database_status` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `startup_time` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `uptime` bigint(20) NOT NULL DEFAULT -1,
  `version` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `archiver` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `session_total` int(10) NOT NULL DEFAULT -1,
  `session_active` smallint(4) NOT NULL DEFAULT -1,
  `processes` int(10) NOT NULL DEFAULT -1,
  `session_logical_read_persecond` int(10) NOT NULL DEFAULT -1,
  `physical_read_persecond` int(10) NOT NULL DEFAULT -1,
  `physical_write_persecond` int(10) NOT NULL DEFAULT -1,
  `physical_read_io_request_persecond` int(10) NOT NULL DEFAULT -1,
  `physical_write_io_request_persecond` int(10) NOT NULL DEFAULT -1,
  `os_cpu_wait_time` int(10) NOT NULL DEFAULT -1,
  `logons_persecond` int(10) NOT NULL DEFAULT -1,
  `logons_current` int(10) NOT NULL DEFAULT -1,
  `user_commits_persecond` int(10) NOT NULL DEFAULT -1,
  `user_rollbacks_persecond` int(10) NOT NULL DEFAULT -1,
  `user_calls_persecond` int(10) NOT NULL DEFAULT -1,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14355 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_postgresql
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_postgresql`;
CREATE TABLE `dashboard_postgresql`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `host` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `port` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `connect` tinyint(2) NOT NULL DEFAULT 0,
  `start_time` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `uptime` int(10) NOT NULL DEFAULT -1,
  `version` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `max_connections` int(10) NOT NULL DEFAULT -1,
  `connections` int(10) NOT NULL DEFAULT -1,
  `active_sql` int(10) NOT NULL DEFAULT -1,
  `prepared_xacts` int(10) NOT NULL DEFAULT -1,
  `tup_fetched` int(10) NOT NULL DEFAULT -1,
  `tup_returned` int(10) NOT NULL DEFAULT -1,
  `tup_inserted` int(10) NOT NULL DEFAULT -1,
  `tup_deleted` int(10) NOT NULL DEFAULT -1,
  `tup_updated` int(10) NOT NULL DEFAULT -1,
  `xact_commit` int(10) NOT NULL DEFAULT -1,
  `xact_rollback` int(10) NOT NULL DEFAULT -1,
  `conflicts` int(10) NOT NULL DEFAULT -1,
  `deadlocks` int(10) NOT NULL DEFAULT -1,
  `checkpoint_req_pct` int(10) NOT NULL DEFAULT -1,
  `checkpoint_avg_write` int(10) NOT NULL DEFAULT -1,
  `checkpoint_total_write` int(10) NOT NULL DEFAULT -1,
  `checkpoint_write_pct` int(10) NOT NULL DEFAULT -1,
  `checkpoint_backend_write_pct` int(10) NOT NULL DEFAULT -1,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_ip_port`(`host`, `port`) USING BTREE,
  INDEX `idx_gmt_create`(`gmt_create`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5874 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_redis
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_redis`;
CREATE TABLE `dashboard_redis`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `host` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `port` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `connect` smallint(4) NOT NULL DEFAULT 0,
  `redis_version` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `redis_mode` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `os` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `arch_bits` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `gcc_version` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `process_id` int(10) NOT NULL DEFAULT -1,
  `run_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `tcp_port` int(11) NOT NULL DEFAULT -1,
  `uptime_in_seconds` int(11) NOT NULL DEFAULT -1,
  `uptime_in_days` int(11) NOT NULL DEFAULT -1,
  `maxclients` int(11) NOT NULL DEFAULT -1,
  `maxmemory` int(11) NOT NULL DEFAULT -1,
  `connected_clients` smallint(4) NOT NULL DEFAULT -1,
  `blocked_clients` smallint(4) NOT NULL DEFAULT -1,
  `used_memory` bigint(20) NOT NULL DEFAULT -1,
  `used_memory_human` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `used_memory_rss` bigint(20) NOT NULL DEFAULT -1,
  `used_memory_rss_human` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `used_memory_peak` bigint(20) NOT NULL DEFAULT -1,
  `used_memory_peak_human` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `used_memory_lua` bigint(20) NOT NULL DEFAULT -1,
  `used_memory_lua_human` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `mem_fragmentation_ratio` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `mem_allocator` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `rdb_bgsave_in_progress` smallint(4) NOT NULL DEFAULT -1,
  `rdb_last_save_time` bigint(18) NOT NULL DEFAULT -1,
  `rdb_last_bgsave_status` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `rdb_last_bgsave_time_sec` smallint(4) NOT NULL DEFAULT -1,
  `aof_enabled` smallint(4) NOT NULL DEFAULT -1,
  `aof_rewrite_in_progress` smallint(4) NOT NULL DEFAULT -1,
  `aof_rewrite_scheduled` smallint(4) NOT NULL DEFAULT -1,
  `aof_last_rewrite_time_sec` smallint(4) NOT NULL DEFAULT -1,
  `aof_last_bgrewrite_status` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `total_connections_received` bigint(18) NOT NULL DEFAULT -1,
  `total_commands_processed` bigint(18) NOT NULL DEFAULT -1,
  `instantaneous_ops_per_sec` smallint(4) NOT NULL DEFAULT -1,
  `rejected_connections` smallint(4) NOT NULL DEFAULT -1,
  `expired_keys` int(10) NOT NULL DEFAULT -1,
  `evicted_keys` int(10) NOT NULL DEFAULT -1,
  `keyspace_hits` int(10) NOT NULL DEFAULT -1,
  `keyspace_misses` int(10) NOT NULL DEFAULT -1,
  `used_cpu_sys` decimal(10, 2) NOT NULL DEFAULT -1.00,
  `used_cpu_user` decimal(10, 2) NOT NULL DEFAULT -1.00,
  `used_cpu_sys_children` int(10) NOT NULL DEFAULT -1,
  `used_cpu_user_children` int(10) NOT NULL DEFAULT -1,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_create_time`(`gmt_create`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16663 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_sqlserver
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_sqlserver`;
CREATE TABLE `dashboard_sqlserver`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `server_id` smallint(4) NOT NULL DEFAULT 0,
  `host` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `port` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tags` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `connect` smallint(4) NOT NULL DEFAULT 0,
  `role` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `uptime` int(11) NOT NULL DEFAULT -1,
  `version` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `lock_timeout` int(11) NOT NULL DEFAULT -1,
  `trancount` int(11) NOT NULL DEFAULT -1,
  `max_connections` int(11) NOT NULL DEFAULT -1,
  `processes` int(11) NOT NULL DEFAULT -1,
  `processes_running` int(11) NOT NULL DEFAULT -1,
  `processes_waits` int(11) NOT NULL DEFAULT -1,
  `connections_persecond` int(11) NOT NULL DEFAULT -1,
  `pack_received_persecond` int(11) NOT NULL DEFAULT -1,
  `pack_sent_persecond` int(11) NOT NULL DEFAULT -1,
  `packet_errors_persecond` int(11) NOT NULL DEFAULT -1,
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dashboard_web
-- ----------------------------
DROP TABLE IF EXISTS `dashboard_web`;
CREATE TABLE `dashboard_web`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `resp_status` smallint(4) NOT NULL DEFAULT 0,
  `http_code` smallint(4) NOT NULL DEFAULT -1,
  `http_proto` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '-1',
  `http_time` int(11) NOT NULL DEFAULT -1,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 270143 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for events
-- ----------------------------
DROP TABLE IF EXISTS `events`;
CREATE TABLE `events`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `event_uuid` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  `event_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event_group` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_entity` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `event_value` decimal(20, 2) NOT NULL,
  `event_tag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `event_unit` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gmt_created`(`gmt_created`) USING BTREE,
  INDEX `idx_event_time`(`event_time`) USING BTREE,
  INDEX `idx_key_entity_time`(`event_key`, `event_entity`, `event_time`) USING BTREE,
  INDEX `idx_entity_time`(`event_entity`, `event_time`) USING BTREE,
  INDEX `idx_type_group_entity_key_time`(`event_type`, `event_group`, `event_entity`, `event_key`, `event_time`) USING BTREE,
  INDEX `idx_event_uuid`(`event_uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27339278 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for meta_clusters
-- ----------------------------
DROP TABLE IF EXISTS `meta_clusters`;
CREATE TABLE `meta_clusters`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `cluster_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `module_id` int(10) NOT NULL DEFAULT 0,
  `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for meta_envs
-- ----------------------------
DROP TABLE IF EXISTS `meta_envs`;
CREATE TABLE `meta_envs`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `env_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for meta_hosts
-- ----------------------------
DROP TABLE IF EXISTS `meta_hosts`;
CREATE TABLE `meta_hosts`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `idc_id` int(10) NOT NULL DEFAULT 0,
  `env_id` int(10) NOT NULL,
  `ip_address` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `hostname` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `online` int(10) NOT NULL DEFAULT 1,
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for meta_idcs
-- ----------------------------
DROP TABLE IF EXISTS `meta_idcs`;
CREATE TABLE `meta_idcs`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `idc_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for meta_modules
-- ----------------------------
DROP TABLE IF EXISTS `meta_modules`;
CREATE TABLE `meta_modules`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `module_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for meta_nodes
-- ----------------------------
DROP TABLE IF EXISTS `meta_nodes`;
CREATE TABLE `meta_nodes`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `cluster_id` int(10) NOT NULL DEFAULT 0,
  `ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `domain` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `port` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `user` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `pass` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `dbid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `role` int(10) NOT NULL DEFAULT 1,
  `monitor` tinyint(2) NOT NULL DEFAULT 1,
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for meta_webs
-- ----------------------------
DROP TABLE IF EXISTS `meta_webs`;
CREATE TABLE `meta_webs`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `cluster_id` int(10) NOT NULL DEFAULT 0,
  `env_id` int(10) NOT NULL DEFAULT 0,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `method` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'GET',
  `monitor` tinyint(2) NOT NULL DEFAULT 1,
  `gmt_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_updated` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type_id` int(10) NOT NULL DEFAULT 0,
  `task_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `task_command` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `task_attribute` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `schedule_type` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `crontab_time` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `period_time` int(10) DEFAULT NULL,
  `next_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `timeout` int(11) NOT NULL DEFAULT 0,
  `status` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `enable` int(10) NOT NULL DEFAULT 1,
  `remark` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `semd_alarm` int(10) NOT NULL DEFAULT 0,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_modify` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for task_log
-- ----------------------------
DROP TABLE IF EXISTS `task_log`;
CREATE TABLE `task_log`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `task_id` int(10) NOT NULL DEFAULT 0,
  `run_id` int(10) NOT NULL DEFAULT 0,
  `content` varchar(5000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_modify` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2135119 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for task_run
-- ----------------------------
DROP TABLE IF EXISTS `task_run`;
CREATE TABLE `task_run`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `task_id` int(10) NOT NULL DEFAULT 0,
  `run_pid` int(10) NOT NULL DEFAULT 0,
  `run_start_time` datetime(0) DEFAULT NULL,
  `run_end_time` datetime(0) DEFAULT NULL,
  `run_status` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_modify` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gmt_create`(`gmt_create`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3338108 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for task_type
-- ----------------------------
DROP TABLE IF EXISTS `task_type`;
CREATE TABLE `task_type`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) NOT NULL DEFAULT 0,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0',
  `description` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `gmt_create` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gmt_modify` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tokens
-- ----------------------------
DROP TABLE IF EXISTS `tokens`;
CREATE TABLE `tokens`  (
  `token_key` varchar(180) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `value` varbinary(1000) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `expired` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`token_key`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `chinese_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `admin` tinyint(1) DEFAULT 0,
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
