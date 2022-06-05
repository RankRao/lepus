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
	"bytes"
	"encoding/gob"
	"fmt"
	"lepus/src/libary/conf"
	"lepus/src/libary/conv"
	"lepus/src/libary/http"
	"lepus/src/libary/logger"
	"lepus/src/libary/mssql"
	"lepus/src/libary/mysql"
	"lepus/src/libary/tool"
	"lepus/src/libary/utils"
	"strings"
	"time"
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_sqlserver_mon.log", conv.StrToInt(conf.Option["debug"]))
var dbClient = mysql.InitConnect()

var (
	//queryVersionSQL   = "SELECT @@VERSION AS [SQL Server and OS Version Info];"
	queryVariablesSQL            = "select @@VERSION as version,@@MAX_CONNECTIONS as max_connections,@@LOCK_TIMEOUT as lock_timeout,@@TRANCOUNT as trancount,@@CONNECTIONS as connections,@@PACK_RECEIVED as pack_received,@@PACK_SENT as pack_sent,@@PACKET_ERRORS as packet_errors,@@ROWCOUNT as row_count,@@CPU_BUSY as cpu_busy,@@IO_BUSY as io_busy,@@CURSOR_ROWS as cursor_rows,@@TOTAL_WRITE as total_write,@@TOTAL_READ as total_read,@@TOTAL_ERRORS as total_errors"
	queryUptimeSQL               = "SELECT crdate startup_time,GETDATE() AS time_now,DATEDIFF(mi,crdate,GETDATE())*60 AS uptime FROM master..sysdatabases WHERE name = 'tempdb';"
	queryOsSysInfoSQL            = "SELECT cpu_count AS [Logical CPU Count],cpu_count/hyperthread_ratio AS [Physical CPU Count],physical_memory_kb/1024 AS [Physical Memory (MB)], sqlserver_start_time FROM master.sys.dm_os_sys_info WITH (NOLOCK) OPTION (RECOMPILE);"
	queryProcessSQL              = "SELECT COUNT(*) as count FROM [Master].[dbo].[SYSPROCESSES] WHERE [DBID] IN ( SELECT  [dbid] FROM [Master].[dbo].[SYSDATABASES]);"
	queryProcessRunningDetailSQL = "SELECT * FROM [Master].[dbo].[SYSPROCESSES] WHERE [DBID] IN ( SELECT  [dbid] FROM [Master].[dbo].[SYSDATABASES])  AND  status !='SLEEPING' AND status !='BACKGROUND'; "
	queryProcessWaitDetailSQL    = "SELECT * FROM [Master].[dbo].[SYSPROCESSES] WHERE [DBID] IN ( SELECT  [dbid] FROM [Master].[dbo].[SYSDATABASES])  AND  status ='SUSPENDED' AND waittime >1;"
)

func collectorMssql(dbType, dbGroup, ip, port, user, pass, tag string) {
	log.Info(fmt.Sprintf("Start check instance %s:%s at %s", ip, port, time.Now()))
	msdb, err := mssql.NewConnect(ip, port, user, pass)
	eventEntity := fmt.Sprintf("%s:%s", ip, port)
	if err != nil {
		log.Error(fmt.Sprintf("Can't connect to mssql database on %s:%s, %s", ip, port, err))
		errInfo := strings.Replace(fmt.Sprint(err), "'", "", -1)
		events := make([]map[string]interface{}, 0)
		event := map[string]interface{}{
			"event_time":   tool.GetNowTime(),
			"event_type":   dbType,
			"event_group":  dbGroup,
			"event_entity": eventEntity,
			"event_key":    "connect",
			"event_value":  0,
			"event_tag":    tag,
			"event_unit":   "",
		}
		events = append(events, event)
		_, err := http.Post(conf.Option["proxy"], events)
		if err != nil {
			log.Error(fmt.Sprintln("Send events to proxy error:", err))
		}

		insertSQL := fmt.Sprintf("insert into dashboard_sqlserver(host,port,tag,connect,error_info) values('%s','%s','%s','%d','%s')", ip, port, tag, 0, errInfo)
		err = mysql.Execute(dbClient, insertSQL)
		if err != nil {
			log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
			return
		}
		return
	}
	defer msdb.Close()

	connect := 1
	queryUptime, _ := mssql.QueryAll(msdb, queryUptimeSQL)
	startupTime := queryUptime[0]["startup_time"].(time.Time)
	uptime := queryUptime[0]["uptime"].(int64)

	queryProcess, _ := mssql.QueryAll(msdb, queryProcessSQL)
	queryProcessRunningDetail, _ := mssql.QueryAll(msdb, queryProcessRunningDetailSQL)
	queryProcessWaitDetail, _ := mssql.QueryAll(msdb, queryProcessWaitDetailSQL)
	process := queryProcess[0]["count"].(int64)
	processRunning := len(queryProcessRunningDetail)
	processWait := len(queryProcessWaitDetail)

	queryVariablesPrev, _ := mssql.QueryAll(msdb, queryVariablesSQL)
	time.Sleep(time.Duration(1) * time.Second)
	queryVariables, _ := mssql.QueryAll(msdb, queryVariablesSQL)

	version := strings.Replace(strings.Split(queryVariables[0]["version"].(string), "-")[0], "Microsoft SQL Server ", "", -1)
	lockTimeOut := queryVariables[0]["lock_timeout"].(int64)
	tranCount := queryVariables[0]["trancount"].(int64)
	maxConnections := queryVariables[0]["max_connections"].(int64)
	packReceived := queryVariables[0]["pack_received"].(int64)
	packSent := queryVariables[0]["pack_sent"].(int64)
	packetErrors := queryVariables[0]["packet_errors"].(int64)
	rowCount := queryVariables[0]["row_count"].(int64)
	cpuBusy := queryVariables[0]["cpu_busy"].(int64)
	ioBusy := queryVariables[0]["io_busy"].(int64)
	cursorRows := queryVariables[0]["cursor_rows"].(int64)
	currentWrite := queryVariables[0]["total_write"].(int64) - queryVariablesPrev[0]["total_write"].(int64)
	currentRead := queryVariables[0]["total_read"].(int64) - queryVariablesPrev[0]["total_read"].(int64)
	currentError := queryVariables[0]["total_errors"].(int64) - queryVariablesPrev[0]["total_errors"].(int64)
	totalErrors := queryVariables[0]["total_errors"].(int64)

	events := make([]map[string]interface{}, 0)

	event := map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "connect",
		"event_value":  connect,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "process",
		"event_value":  process,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "processRunning",
		"event_value":  processRunning,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "lockTimeOut",
		"event_value":  lockTimeOut,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "tranCount",
		"event_value":  tranCount,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "packReceived",
		"event_value":  packReceived,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "packSent",
		"event_value":  packSent,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "packetErrors",
		"event_value":  packetErrors,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "rowCount",
		"event_value":  rowCount,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "cpuBusy",
		"event_value":  cpuBusy,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "ioBusy",
		"event_value":  ioBusy,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "cursorRows",
		"event_value":  cursorRows,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "currentWrite",
		"event_value":  currentWrite,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "currentRead",
		"event_value":  currentRead,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "currentError",
		"event_value":  currentError,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "totalErrors",
		"event_value":  totalErrors,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	_, err = http.Post(conf.Option["proxy"], events)
	if err != nil {
		log.Error(fmt.Sprintln("Send events to proxy error:", err))
	}

	insertSQL := fmt.Sprintf("insert into dashboard_sqlserver(host,port,tag,connect,version,startup_time,uptime,lock_timeout,trancount,max_connections,processes,"+
		"processes_running,processes_waits,pack_received,pack_sent,packet_errors,row_count,cpu_busy,io_busy,cursor_rows,current_write,current_read,current_error,total_errors )"+
		"values('%s','%s','%s','%d','%s','%s','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d')", ip, port, tag, connect, version, startupTime, uptime, lockTimeOut, tranCount, maxConnections, process,
		processRunning, processWait, packReceived, packSent, packetErrors, rowCount, cpuBusy, ioBusy, cursorRows, currentWrite, currentRead, currentError, totalErrors)

	err = mysql.Execute(dbClient, insertSQL)
	if err != nil {
		log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
		return
	}

	log.Info(fmt.Sprintf("Complete check instance %s:%s at %s", ip, port, time.Now()))

}

func scanMssql() {
	rows, err := mysql.QueryAll(dbClient, "select ip,port,user,pass,module_name,cluster_name,env_name,idc_name from meta_nodes a join meta_clusters b on a.cluster_id=b.id join meta_modules c on b.module_id=c.id join meta_hosts d on a.ip=d.ip_address join meta_envs e on d.env_id=e.id join meta_idcs f on d.idc_id=f.id where a.monitor=1 and c.module_name ='SQLServer' ")
	if err != nil {
		log.Error(fmt.Sprintln("Can't query mysql database, ", err))
		return
	}
	for _, row := range rows {
		origPass, err := utils.AesPassDecode(row["pass"].(string), conf.Option["db_pass_key"])
		if err != nil {
			log.Error("Encrypt Password Error.")
			return
		}
		go collectorMssql(row["module_name"].(string), row["env_name"].(string), row["ip"].(string), row["port"].(string), row["user"].(string), origPass, row["cluster_name"].(string))
	}
}

func main() {
	startTime := time.Now()
	fmt.Println("Server start at ", startTime)
	log.Info(fmt.Sprintln("Server start at ", startTime))

	scanMssql()
	time.Sleep(time.Duration(8) * time.Second)
	defer dbClient.Close()
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
