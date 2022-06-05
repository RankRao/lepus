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
	"lepus/src/libary/oracle"
	"lepus/src/libary/tool"
	"lepus/src/libary/utils"
	"strings"
	"time"
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_oracle_mon.log", conv.StrToInt(conf.Option["debug"]))
var dbClient = mysql.InitConnect()

var (
	queryInstanceSQL      = "select * from v$instance"
	queryDatabaseSQL      = "select * from v$database"
	queryParameterSQL     = "select name,type,value from v$parameter"
	queryProcessesSQL     = "select name,type,value from v$parameter where name='processes' "
	querySysStatSQL       = "select name,value from v$sysstat "
	querySessionTotalSQL  = "Select a.SID,a.SERIAL#,a.STATUS,a.USERNAME,a.MACHINE,a.MODULE,a.EVENT,b.SQL_ID,b.SQL_TEXT from v$session a, v$sqlarea b where a.sql_hash_value = b.HASH_VALUE and  a.username not in('SYS','SYSTEM') and a.username is not null"
	querySessionActiveSQL = "Select a.SID,a.SERIAL#,a.STATUS,a.USERNAME,a.MACHINE,a.MODULE,a.EVENT,b.SQL_ID,b.SQL_TEXT from v$session a, v$sqlarea b where a.sql_hash_value = b.HASH_VALUE and  a.username not in('SYS','SYSTEM') and a.username is not null and a.status='ACTIVE'"
)

func collectorOracle(dbType, dbGroup, host, port, user, pass, sid, tag string) {

	eventEntity := fmt.Sprintf("%s:%s/%s", host, port, sid)
	log.Info(fmt.Sprintf("Start check instance %s at %s", eventEntity, time.Now()))
	oraCon, err := oracle.NewConnect(host, port, user, pass, sid)
	if err != nil {
		log.Error(fmt.Sprintf("Can't connect to oracle database on %s, %s", eventEntity, err))
		errInfo := strings.Replace(fmt.Sprint(err), "'", "", -1)
		events := make([]map[string]interface{}, 0)

		event := map[string]interface{}{
			"event_time":   tool.GetNowTime(),
			"event_type":   dbType,
			"event_group":  dbGroup,
			"event_entity": eventEntity,
			"event_key":    "connect",
			"event_value":  0.00,
			"event_tag":    tag,
			"event_unit":   "",
		}
		events = append(events, event)
		_, err := http.Post(conf.Option["proxy"], events)
		if err != nil {
			log.Error(fmt.Sprintln("Send events to proxy error:", err))
		}

		insertSQL := fmt.Sprintf("insert into dashboard_oracle(host,port,sid,tag,connect,error_info) values('%s','%s','%s','%s','%d','%s')", host, port, sid, tag, 0, errInfo)
		err = mysql.Execute(dbClient, insertSQL)
		if err != nil {
			log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
			return
		}
		return
	}

	defer oraCon.Close()

	instance, err := oracle.QueryAll(oraCon, queryInstanceSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query oracle instance on %s, %s", eventEntity, err))
		return
	}
	database, err := oracle.QueryAll(oraCon, queryDatabaseSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query oracle database on %s, %s", eventEntity, err))
		return
	}
	processInfo, err := oracle.QueryAll(oraCon, queryProcessesSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query oracle processes on %s, %s", eventEntity, err))
		return
	}

	instanceName := instance[0]["INSTANCE_NAME"]
	instanceRole := instance[0]["INSTANCE_ROLE"]
	instanceStatus := instance[0]["STATUS"]
	startupTime := instance[0]["STARTUP_TIME"]
	version := instance[0]["VERSION"]
	databaseStatus := instance[0]["DATABASE_STATUS"]
	hostname := instance[0]["HOST_NAME"]
	archiver := instance[0]["ARCHIVER"]
	databaseRole := database[0]["DATABASE_ROLE"]
	openMode := database[0]["OPEN_MODE"]
	protectedMode := database[0]["PROTECTION_MODE"]
	processes := processInfo[0]["VALUE"]
	uptime := time.Now().Unix() - startupTime.(time.Time).Unix()

	rows, err := oraCon.Query(querySysStatSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query oracle sysstat on%s, %s", eventEntity, err))
		return
	}
	defer rows.Close()
	var key, value string
	sysStatsPrev := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			log.Error(fmt.Sprintf("Can't scan oracle sysstat on%s, %s", eventEntity, err))
			return
		}
		sysStatsPrev[key] = value
	}

	time.Sleep(time.Duration(1) * time.Second)

	rows, err = oraCon.Query(querySysStatSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query oracle sysstat on%s, %s", eventEntity, err))
		return
	}
	defer rows.Close()
	sysStats := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			log.Error(fmt.Sprintf("Can't scan oracle sysstat on%s, %s", eventEntity, err))
			return
		}
		sysStats[key] = value
	}

	connect := 1
	sessionLogicalReadsPersecond := conv.StrToInt(sysStats["session_logical_reads"]) - conv.StrToInt(sysStatsPrev["session_logical_reads"])
	physicalReadsPersecond := conv.StrToInt(sysStats["physical read"]) - conv.StrToInt(sysStatsPrev["physical read"])
	physicalWritePersecond := conv.StrToInt(sysStats["physical write"]) - conv.StrToInt(sysStatsPrev["physical write"])
	physicalWriteIoRequestsPersecond := conv.StrToInt(sysStats["physical write total IO requests"]) - conv.StrToInt(sysStatsPrev["physical write total IO requests"])
	physicalReadIoRequestsPersecond := conv.StrToInt(sysStats["physical read total IO requests"]) - conv.StrToInt(sysStatsPrev["physical read total IO requests"])
	osCpuWaitTime := conv.StrToInt(sysStats["OS CPU Qt wait time"]) - conv.StrToInt(sysStatsPrev["OS CPU Qt wait time"])
	logonsCumulative := conv.StrToInt(sysStats["logons cumulative"]) - conv.StrToInt(sysStatsPrev["logons cumulative"])
	logonsCurrent := conv.StrToInt(sysStats["logons current"])
	userCommitsPersecond := conv.StrToInt(sysStats["user commits"]) - conv.StrToInt(sysStatsPrev["user commits"])
	userRollbacksPersecond := conv.StrToInt(sysStats["user rollbacks"]) - conv.StrToInt(sysStatsPrev["user rollbacks"])
	userCallsPersecond := conv.StrToInt(sysStats["user calls"]) - conv.StrToInt(sysStatsPrev["user calls"])

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

	sessionTotalDetail, _ := oracle.QueryAll(oraCon, querySessionTotalSQL)
	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "sessionTotal",
		"event_value":  len(sessionTotalDetail),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	sessionActiveDetail, _ := oracle.QueryAll(oraCon, querySessionActiveSQL)
	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "sessionActive",
		"event_value":  len(sessionActiveDetail),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "sessionLogicalReadsPersecond",
		"event_value":  sessionLogicalReadsPersecond,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "physicalReadsPersecond",
		"event_value":  physicalReadsPersecond,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "physicalWritePersecond",
		"event_value":  physicalWritePersecond,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "physicalWriteIoRequestsPersecond",
		"event_value":  physicalWriteIoRequestsPersecond,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "physicalReadIoRequestsPersecond",
		"event_value":  physicalReadIoRequestsPersecond,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "osCpuWaitTime",
		"event_value":  osCpuWaitTime,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "logonsCumulative",
		"event_value":  logonsCumulative,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "logonsCurrent",
		"event_value":  logonsCurrent,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "userCommitsPersecond",
		"event_value":  userCommitsPersecond,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "userRollbacksPersecond",
		"event_value":  userRollbacksPersecond,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "userCallsPersecond",
		"event_value":  userCallsPersecond,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	_, err = http.Post(conf.Option["proxy"], events)
	if err != nil {
		log.Error(fmt.Sprintln("Send events to proxy error:", err))
	}

	insertSQL := fmt.Sprintf("insert into dashboard_oracle(host,port,tag,sid,connect,instance_name,instance_role,instance_status,database_role,open_mode,protection_mode,host_name,database_status,startup_time"+
		",uptime,version,archiver,session_total,session_active,processes,session_logical_read_persecond,physical_read_persecond,physical_write_persecond,physical_read_io_request_persecond"+
		",physical_write_io_request_persecond,os_cpu_wait_time,logons_cumulative,logons_current,user_commits_persecond,user_rollbacks_persecond"+
		",user_calls_persecond)"+
		"values('%s','%s','%s','%s','%d','%s','%s','%s','%s','%s','%s','%s','%s','%s','%d','%s','%s','%d','%d','%s','%d','%d','%d','%d',"+
		"'%d','%d','%d','%d','%d','%d','%d')", host, port, tag, sid, connect, instanceName, instanceRole, instanceStatus, databaseRole, openMode, protectedMode, hostname, databaseStatus, startupTime, uptime,
		version, archiver, len(sessionTotalDetail), len(sessionActiveDetail), processes, sessionLogicalReadsPersecond, physicalReadsPersecond, physicalWritePersecond, physicalReadIoRequestsPersecond,
		physicalWriteIoRequestsPersecond, osCpuWaitTime, logonsCumulative, logonsCurrent, userCommitsPersecond, userRollbacksPersecond, userCallsPersecond)

	err = mysql.Execute(dbClient, insertSQL)
	if err != nil {
		log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
		return
	}

	log.Info(fmt.Sprintf("Complete check instance %s at %s", eventEntity, time.Now()))

}

func scanOracle() {
	rows, err := mysql.QueryAll(dbClient, "select ip,port,user,pass,dbid,module_name,cluster_name,env_name,idc_name from meta_nodes a join meta_clusters b on a.cluster_id=b.id join meta_modules c on b.module_id=c.id join meta_hosts d on a.ip=d.ip_address join meta_envs e on d.env_id=e.id join meta_idcs f on d.idc_id=f.id where a.monitor=1 and c.module_name='Oracle'")
	if err != nil {
		log.Error(fmt.Sprintln("Can't query mysql database, ", err))
		return
	}
	if len(rows) == 0 {
		log.Warning(fmt.Sprintln("Not oracle node found, please add oracle node to meta data"))
		return
	}
	for _, row := range rows {
		origPass, err := utils.AesPassDecode(row["pass"].(string), conf.Option["db_pass_key"])
		if err != nil {
			log.Error("Encrypt Password Error.")
			return
		}
		go collectorOracle(row["module_name"].(string), row["env_name"].(string), row["ip"].(string), row["port"].(string), row["user"].(string), origPass, row["dbid"].(string), row["cluster_name"].(string))
	}
}

func main() {
	startTime := time.Now()
	fmt.Println("Server start at ", startTime)
	log.Info(fmt.Sprintln("Server start at ", startTime))

	scanOracle()
	time.Sleep(time.Duration(8) * time.Second)
	defer dbClient.Close()
}
