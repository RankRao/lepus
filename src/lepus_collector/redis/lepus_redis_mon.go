/*
Copyright 2014-2022 The Lepus Team Group, website: https://www.lepus.cc
Licensed under the GNU General Public License, Version 3.0 (the "GPLv3 License");
You may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.gnu.org/licenses/gpl-3.0.html
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
Special note:
Please do not use this source code for any commercial purpose,
or use it for commercial purposes after secondary development, otherwise you may bear legal risks.
*/

package main

import (
	"fmt"
	"lepus/src/libary/conf"
	"lepus/src/libary/conv"
	"lepus/src/libary/http"
	"lepus/src/libary/logger"
	"lepus/src/libary/mysql"
	_ "lepus/src/libary/redis"
	"lepus/src/libary/tool"
	"lepus/src/libary/utils"
	_ "reflect"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_redis_mon.log", conv.StrToInt(conf.Option["debug"]))
var dbClient = mysql.InitConnect()

func collectorRedis(dbType string, dbGroup string, ip string, port string, origPass string, tag string) {
	log.Info(fmt.Sprintf("Start check instance %s:%s at %s", ip, port, time.Now()))
	eventEntity := fmt.Sprintf("%s:%s", ip, port)
	rdb, err := redis.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Error(fmt.Sprintln("Can't dial redis, ", err))
		errInfo := strings.Replace(fmt.Sprint(err), "'", "", -1)
		eventEntity := fmt.Sprintf("%s:%s", ip, port)
		events := make([]map[string]interface{}, 0)
		event := map[string]interface{}{
			"event_time":   tool.GetNowTime(),
			"event_type":   dbType,
			"event_group":  dbGroup,
			"event_entity": eventEntity,
			"event_key":    "connect",
			"event_value":  "0",
			"event_tag":    tag,
			"event_unit":   "",
		}
		events = append(events, event)
		_, err := http.Post(conf.Option["proxy"], events)
		if err != nil {
			log.Error(fmt.Sprintln("Send events to proxy error:", err))
		}
		insertSQL := fmt.Sprintf("insert into dashboard_redis(host,port,tag,connect,error_info) values('%s','%s','%s','%d','%s')", ip, port, tag, 0, errInfo)
		err = mysql.Execute(dbClient, insertSQL)
		if err != nil {
			log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
			return
		}
		return
	}
	defer rdb.Close()

	if origPass != "" {
		if _, err := rdb.Do("AUTH", origPass); err != nil {
			rdb.Close()
			log.Error(fmt.Sprintln("Redis Auth error, ", err))
			return
		}
	}

	info, err := redis.String(rdb.Do("INFO"))
	if err != nil {
		log.Error(fmt.Sprintln("Can't do redis info query, ", err))
		return
	}

	infoMap := make(map[string]string)
	infoArray := strings.Split(info, "\n")
	for _, item := range infoArray {
		if strings.Contains(item, ":") {
			v := strings.Split(item, ":")
			infoMap[v[0]] = v[1]
		}

	}

	connect := 1
	redisVersion := infoMap["redis_version"]
	redisMode := infoMap["redis_mode"]
	os := infoMap["os"]
	archBits := infoMap["arch_bits"]
	gccVersion := infoMap["gcc_version"]
	processId := infoMap["process_id"]
	runId := infoMap["run_id"]
	tcpPort := infoMap["tcp_port"]
	uptimeInSeconds := infoMap["uptime_in_seconds"]
	uptimeInDays := infoMap["uptime_in_days"]
	connectedClients := infoMap["connected_clients"]
	blockedClients := infoMap["blocked_clients"]
	usedMemory := infoMap["used_memory"]
	usedMemoryHuman := infoMap["used_memory_human"]
	usedMemoryRss := infoMap["used_memory_rss"]
	usedMemoryRssHuman := infoMap["used_memory_rss_human"]
	usedMemoryPeak := infoMap["used_memory_peak"]
	usedMemoryPeakHuman := infoMap["used_memory_peak_human"]
	usedMemoryLua := infoMap["used_memory_lua"]
	usedMemoryLuaHuman := infoMap["used_memory_lua_human"]
	memFragmentationRatio := infoMap["mem_fragmentation_ratio"]
	memAllocator := infoMap["mem_allocator"]
	rdbBgsaveInProgress := infoMap["rdb_bgsave_in_progress"]
	rdbLastSaveTime := infoMap["rdb_last_save_time"]
	rdbLastBgsaveStatus := infoMap["rdb_last_bgsave_status"]
	rdbLastBgsaveTimeSec := infoMap["rdb_last_bgsave_time_sec"]
	aofEnabled := infoMap["aof_enabled"]
	aofRewriteInProgress := infoMap["aof_rewrite_in_progress"]
	aofRewriteScheduled := infoMap["aof_rewrite_scheduled"]
	aofLastRewriteTimeSec := infoMap["aof_last_rewrite_time_sec"]
	aofLastBgrewriteStatus := infoMap["aof_last_bgrewrite_status"]
	totalConnectionsReceived := infoMap["total_connections_received"]
	totalCommandsProcessed := infoMap["total_commands_processed"]
	instantaneousOpsPerSec := infoMap["instantaneous_ops_per_sec"]
	rejectedConnections := infoMap["rejected_connections"]
	expiredKeys := infoMap["expired_keys"]
	evictedKeys := infoMap["evicted_keys"]
	keyspaceHits := infoMap["keyspace_hits"]
	keyspaceMisses := infoMap["keyspace_misses"]
	usedCpuSys := infoMap["used_cpu_sys"]
	usedCpuUser := infoMap["used_cpu_user"]
	usedCpuSysChildren := infoMap["used_cpu_sys_children"]
	usedCpuUserChildren := infoMap["used_cpu_user_children"]

	maxClientsConfig, err := redis.Strings(rdb.Do("config", "GET", "maxclients"))
	if err != nil {
		log.Error(fmt.Sprintln("Can't do redis maxclients query, ", err))
		return
	}
	maxClients := maxClientsConfig[1]

	maxMemoryConfig, err := redis.Strings(rdb.Do("config", "GET", "maxmemory"))
	if err != nil {
		log.Error(fmt.Sprintln("Can't do redis maxMemory query, ", err))
		return
	}

	maxMemory := maxMemoryConfig[1]

	events := make([]map[string]interface{}, 0)

	event := map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "connect",
		"event_value":  1,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "uptimeInSeconds",
		"event_value":  utils.StrToFloat(strings.Replace(uptimeInSeconds, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "uptimeInDays",
		"event_value":  utils.StrToFloat(strings.Replace(uptimeInDays, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "connectedClients",
		"event_value":  utils.StrToFloat(strings.Replace(connectedClients, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "blockedClients",
		"event_value":  utils.StrToFloat(strings.Replace(blockedClients, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "usedMemory",
		"event_value":  utils.StrToFloat(strings.Replace(usedMemory, "\r", "", -1)) / 1024 / 1024,
		"event_tag":    tag,
		"event_unit":   "MB",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "usedMemoryRss",
		"event_value":  utils.StrToFloat(strings.Replace(usedMemoryRss, "\r", "", -1)) / 1024 / 1024,
		"event_tag":    tag,
		"event_unit":   "MB",
	}

	events = append(events, event)
	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "usedMemoryPeak",
		"event_value":  utils.StrToFloat(strings.Replace(usedMemoryPeak, "\r", "", -1)) / 1024 / 1024,
		"event_tag":    tag,
		"event_unit":   "MB",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "usedMemoryLua",
		"event_value":  utils.StrToFloat(strings.Replace(usedMemoryLua, "\r", "", -1)) / 1024 / 1024,
		"event_tag":    tag,
		"event_unit":   "MB",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "memFragmentationRatio",
		"event_value":  utils.StrToFloat(strings.Replace(memFragmentationRatio, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "%",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "totalConnectionsReceived",
		"event_value":  utils.StrToFloat(strings.Replace(totalConnectionsReceived, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "totalCommandsProcessed",
		"event_value":  utils.StrToFloat(strings.Replace(totalCommandsProcessed, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "instantaneousOpsPerSec",
		"event_value":  utils.StrToFloat(strings.Replace(instantaneousOpsPerSec, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "rejectedConnections",
		"event_value":  utils.StrToFloat(strings.Replace(rejectedConnections, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "expiredKeys",
		"event_value":  utils.StrToFloat(strings.Replace(expiredKeys, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "evictedKeys",
		"event_value":  utils.StrToFloat(strings.Replace(evictedKeys, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "keyspaceHits",
		"event_value":  utils.StrToFloat(strings.Replace(keyspaceHits, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "keyspaceMisses",
		"event_value":  utils.StrToFloat(strings.Replace(keyspaceMisses, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "usedCpuSys",
		"event_value":  utils.StrToFloat(strings.Replace(usedCpuSys, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "usedCpuUser",
		"event_value":  utils.StrToFloat(strings.Replace(usedCpuUser, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "usedCpuSysChildren",
		"event_value":  utils.StrToFloat(strings.Replace(usedCpuSysChildren, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "usedCpuUserChildren",
		"event_value":  utils.StrToFloat(strings.Replace(usedCpuUserChildren, "\r", "", -1)),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	_, err = http.Post(conf.Option["proxy"], events)
	if err != nil {
		log.Error(fmt.Sprintln("Send events to proxy error:", err))
	}

	insertSQL := fmt.Sprintf("insert into dashboard_redis("+
		"host,port,tag,connect,redis_version,redis_mode,os,arch_bits, gcc_version, process_id, run_id, tcp_port, uptime_in_seconds, uptime_in_days, connected_clients, blocked_clients, maxclients,maxmemory,used_memory, used_memory_human, used_memory_rss, used_memory_rss_human, used_memory_peak, used_memory_peak_human,used_memory_lua, used_memory_lua_human, mem_fragmentation_ratio, mem_allocator, rdb_bgsave_in_progress, rdb_last_save_time, rdb_last_bgsave_status, rdb_last_bgsave_time_sec, aof_enabled, aof_rewrite_in_progress,aof_rewrite_scheduled, aof_last_rewrite_time_sec, aof_last_bgrewrite_status, total_connections_received, total_commands_processed, instantaneous_ops_per_sec, rejected_connections, expired_keys, evicted_keys,keyspace_hits, keyspace_misses, used_cpu_sys, used_cpu_user, used_cpu_sys_children, used_cpu_user_children) "+
		"values('%s','%s','%s','%d','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')",
		ip, port, tag, connect, redisVersion, redisMode, os, archBits, gccVersion, processId, runId, tcpPort, uptimeInSeconds, uptimeInDays, connectedClients, blockedClients, maxClients, maxMemory, usedMemory, usedMemoryHuman, usedMemoryRss, usedMemoryRssHuman, usedMemoryPeak, usedMemoryPeakHuman, usedMemoryLua, usedMemoryLuaHuman, memFragmentationRatio, memAllocator, rdbBgsaveInProgress, rdbLastSaveTime, rdbLastBgsaveStatus, rdbLastBgsaveTimeSec, aofEnabled, aofRewriteInProgress, aofRewriteScheduled, aofLastRewriteTimeSec, aofLastBgrewriteStatus, totalConnectionsReceived, totalCommandsProcessed, instantaneousOpsPerSec, rejectedConnections, expiredKeys, evictedKeys, keyspaceHits, keyspaceMisses, usedCpuSys, usedCpuUser, usedCpuSysChildren, usedCpuUserChildren)

	err = mysql.Execute(dbClient, insertSQL)
	if err != nil {
		log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
		return
	}

	log.Info(fmt.Sprintf("Complete check instance %s:%s at %s", ip, port, time.Now()))

}

func scanRedis() {
	rows, err := mysql.QueryAll(dbClient, "select ip,port,user,pass,module_name,cluster_name,env_name,idc_name from meta_nodes a join meta_clusters b on a.cluster_id=b.id join meta_modules c on b.module_id=c.id join meta_hosts d on a.ip=d.ip_address join meta_envs e on d.env_id=e.id join meta_idcs f on d.idc_id=f.id where a.monitor=1 and c.module_name='Redis'")
	if err != nil {
		log.Error(fmt.Sprintln("Can't query mysql database, ", err))
		return
	}
	for _, row := range rows {
		var origPass string
		if row["pass"] != "" && row["pass"] != nil {
			origPass, err = utils.AesPassDecode(row["pass"].(string), conf.Option["db_pass_key"])
			if err != nil {
				log.Error("Encrypt Password Error.")
				return
			}
		}
		go collectorRedis(row["module_name"].(string), row["env_name"].(string), row["ip"].(string), row["port"].(string), origPass, row["cluster_name"].(string))
	}
}
func main() {
	startTime := time.Now()
	fmt.Println("Server start at ", startTime)
	log.Info(fmt.Sprintln("Server start at ", startTime))

	//for true {
	//	scanRedis()
	//	time.Sleep(time.Duration(conv.StrToInt(conf.Option["interval"])) * time.Second)
	//}
	scanRedis()
	time.Sleep(time.Duration(8) * time.Second)
	defer dbClient.Close()
}
