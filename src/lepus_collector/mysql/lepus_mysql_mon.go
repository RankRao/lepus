/*
Copyright 2014-2021 The Lepus Team Group, website: https://www.lepus.cc
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
	"lepus/src/libary/mysql"
	"lepus/src/libary/tool"
	"lepus/src/libary/utils"
	"time"
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_mysql_mon.log", conv.StrToInt(conf.Option["debug"]))
var dbClient = mysql.InitConnect()

var (
	queryVersionSQL   = "select version() as version limit 1"
	queryStatusSQL    = "show global status"
	queryVariablesSQL = "show global variables"
)

func collectorMySQL(dbType, dbGroup, ip, port, user, pass, tag string) {
	log.Info(fmt.Sprintf("Start check instance %s:%s at %s", ip, port, time.Now()))
	mydb, err := mysql.Connect(ip, port, user, pass, "information_schema")
	eventEntity := fmt.Sprintf("%s:%s", ip, port)

	if err != nil {
		log.Error(fmt.Sprintf("Can't connect to mysql database on %s:%s, %s", ip, port, err))
		events := make([]map[string]interface{}, 0)
		detail := make([]map[string]string, 0)
		detail = append(detail, map[string]string{"Error": fmt.Sprint(err)})
		//detail = append(detail, map[string]string{"Error": fmt.Sprint(err), "Into":"BBB"})
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
		//fmt.Println(events)
		_, err := http.Post(conf.Option["proxy"], events)
		if err != nil {
			log.Error(fmt.Sprintln("Send events to proxy error:", err))
		}

		insertSQL := fmt.Sprintf("insert into dashboard_mysql(ip,port,tag,connect) values('%s','%s','%s','%d')", ip, port, tag, 0)
		err = mysql.Execute(dbClient, insertSQL)
		if err != nil {
			log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
			return
		}
		return
	}

	row := mydb.QueryRow(queryVersionSQL)
	var version string
	if err := row.Scan(&version); err != nil {
		log.Error(fmt.Sprintf("Can't scan mysql version on %s:%d, %s", ip, port, err))
		return
	}

	rows, err := mydb.Query(queryStatusSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query mysql status on %s:%d, %s", ip, port, err))
		return
	}

	defer rows.Close()
	var key, value string
	globalStatusPrev := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			log.Error(fmt.Sprintf("Can't scan mysql status on %s:%d, %s", ip, port, err))
			return
		}
		globalStatusPrev[key] = value
	}

	time.Sleep(time.Duration(1) * time.Second)

	rows, err = mydb.Query(queryStatusSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query mysql status on %s:%d, %s", ip, port, err))
		return
	}
	defer rows.Close()
	globalStatus := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			log.Error(fmt.Sprintf("Can't scan mysql status on %s:%s, %s", ip, port, err))
			return
		}
		globalStatus[key] = value
	}

	//fmt.Println(globalStatus)
	rows, err = mydb.Query(queryVariablesSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query mysql variables on %s:%s, %s", ip, port, err))
		return
	}
	defer rows.Close()
	globalVariables := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			log.Error(fmt.Sprintf("Can't scan mysql variables on %s:%s, %s", ip, port, err))
			return
		}
		globalVariables[key] = value
	}

	connect := 1
	//variables
	timezone := globalVariables["time_zone"]
	readonly := globalVariables["read_only"]
	hostname := globalVariables["hostname"]
	keyBufferSize := globalVariables["key_buffer_size"]
	sortBufferSize := globalVariables["sort_buffer_size"]
	joinBufferSize := globalVariables["join_buffer_size"]
	maxConnections := globalVariables["max_connections"]
	openFilesLimit := globalVariables["open_files_limit"]
	tableOpenCache := globalVariables["table_open_cache"]

	uptime := conv.StrToInt(globalStatus["Uptime"])
	openFiles := conv.StrToInt(globalStatus["open_files"])
	openTables := conv.StrToInt(globalStatus["Open_tables"])
	threadsConnected := conv.StrToInt(globalStatus["Threads_connected"])
	threadsRunning := conv.StrToInt(globalStatus["Threads_running"])
	threadsCreated := conv.StrToInt(globalStatus["Threads_created"])
	threadsCached := conv.StrToInt(globalStatus["Threads_cached"])
	connections := conv.StrToInt(globalStatus["Connections"])
	abortedClients := conv.StrToInt(globalStatus["Aborted_clients"])
	abortedConnects := conv.StrToInt(globalStatus["Aborted_connects"])

	bytesReceived := conv.StrToInt(globalStatus["Bytes_received"]) - conv.StrToInt(globalStatusPrev["Bytes_received"])
	bytesSent := conv.StrToInt(globalStatus["Bytes_sent"]) - conv.StrToInt(globalStatusPrev["Bytes_sent"])
	comSelect := conv.StrToInt(globalStatus["Com_select"]) - conv.StrToInt(globalStatusPrev["Com_select"])
	comInsert := conv.StrToInt(globalStatus["Com_insert"]) - conv.StrToInt(globalStatusPrev["Com_insert"])
	comUpdate := conv.StrToInt(globalStatus["Com_update"]) - conv.StrToInt(globalStatusPrev["Com_update"])
	comDelete := conv.StrToInt(globalStatus["Com_delete"]) - conv.StrToInt(globalStatusPrev["Com_delete"])
	comCommit := conv.StrToInt(globalStatus["Com_commit"]) - conv.StrToInt(globalStatusPrev["Com_commit"])
	comRollback := conv.StrToInt(globalStatus["Com_rollback"]) - conv.StrToInt(globalStatusPrev["Com_rollback"])
	questions := conv.StrToInt(globalStatus["Questions"]) - conv.StrToInt(globalStatusPrev["Questions"])
	queries := conv.StrToInt(globalStatus["Queries"]) - conv.StrToInt(globalStatusPrev["Queries"])
	slowQueries := conv.StrToInt(globalStatus["Slow_queries"])

	//innodb status
	innodbPagesCreated := conv.StrToInt(globalStatus["Innodb_pages_created"])
	innodbPagesRead := conv.StrToInt(globalStatus["Innodb_pages_read"])
	innodbPagesWritten := conv.StrToInt(globalStatus["Innodb_pages_written"])
	innodbRowLockCurrentWaits := conv.StrToInt(globalStatus["Innodb_row_lock_current_waits"])
	innodbBufferPoolReadRequests := conv.StrToInt(globalStatus["Innodb_buffer_pool_read_requests"]) - conv.StrToInt(globalStatusPrev["Innodb_buffer_pool_read_requests"])
	innodbBufferPoolWriteRequests := conv.StrToInt(globalStatus["Innodb_buffer_pool_write_requests"]) - conv.StrToInt(globalStatusPrev["Innodb_buffer_pool_write_requests"])
	innodbRowsDeleted := conv.StrToInt(globalStatus["Innodb_rows_deleted"]) - conv.StrToInt(globalStatusPrev["Innodb_rows_deleted"])
	innodbRowsInserted := conv.StrToInt(globalStatus["Innodb_rows_inserted"]) - conv.StrToInt(globalStatusPrev["Innodb_rows_inserted"])
	innodbRowsRead := conv.StrToInt(globalStatus["Innodb_rows_read"]) - conv.StrToInt(globalStatusPrev["Innodb_rows_read"])
	innodbRowsUpdated := conv.StrToInt(globalStatus["Innodb_rows_updated"]) - conv.StrToInt(globalStatusPrev["Innodb_rows_updated"])

	events := make([]map[string]interface{}, 0)

	event := map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "threadsConnected",
		"event_value":  threadsConnected,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "threadsRunning",
		"event_value":  threadsRunning,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)
	threadsWait, _ := mysql.QueryAll(mydb, "select * from information_schema.processlist where db is not null and db != 'information_schema' and  command !='Sleep' and time >0 order by time desc limit 50;")
	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "threadWait",
		"event_value":  len(threadsWait),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	longQuery, _ := mysql.QueryAll(mydb, "select host,db,user,command,info from information_schema.processlist where db is not null and db != 'information_schema' and  command !='Sleep' and time > 10 order by time desc limit 50;")
	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "longQuery",
		"event_value":  len(longQuery),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	ActiveTrx, _ := mysql.QueryAll(mydb, "select trx_id,trx_mysql_thread_id,trx_started,trx_isolation_level,trx_state,trx_rows_locked,trx_lock_structs,trx_tables_locked,trx_unique_checks,trx_is_read_only,trx_query from information_schema.INNODB_TRX limit 50;")
	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "ActiveTrx",
		"event_value":  len(ActiveTrx),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	LongTrx, _ := mysql.QueryAll(mydb, "select trx_id,trx_mysql_thread_id,trx_started,trx_isolation_level,trx_state,trx_rows_locked,trx_lock_structs,trx_tables_locked,trx_unique_checks,trx_is_read_only,trx_query from information_schema.INNODB_TRX where timestampdiff(second, trx_started,now())>10 limit 50;")
	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "LongTrx",
		"event_value":  len(LongTrx),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "slowQueries",
		"event_value":  slowQueries,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "queries",
		"event_value":  queries,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	mydb.Close()

	_, err = http.Post(conf.Option["proxy"], events)
	if err != nil {
		log.Error(fmt.Sprintln("Send events to proxy error:", err))
	}

	insertSQL := fmt.Sprintf("insert into dashboard_mysql(ip,port,tag,connect,hostname,version,timezone,uptime,readonly,max_connections,open_files_limit,open_files,"+
		"table_open_cache,open_tables,long_query,active_trx,threads_connected,threads_running,threads_wait,threads_created,threads_cached,connections,aborted_clients,aborted_connects,bytes_received,bytes_sent,"+
		"com_select,com_insert,com_update,com_delete,com_commit,com_rollback,questions,queries,slow_queries,key_buffer_size,sort_buffer_size,join_buffer_size,innodb_pages_created,innodb_pages_read,innodb_pages_written,"+
		"innodb_row_lock_current_waits,innodb_buffer_pool_read_requests,innodb_buffer_pool_write_requests,innodb_rows_read,innodb_rows_inserted,innodb_rows_updated,innodb_rows_deleted )"+
		"values('%s','%s','%s','%d','%s','%s','%s','%d','%s','%s','%s','%d','%s','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%s'"+
		",'%s','%s','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d')", ip, port, tag, connect, hostname, version, timezone, uptime, readonly, maxConnections, openFilesLimit, openFiles,
		tableOpenCache, openTables, len(longQuery), len(ActiveTrx), threadsConnected, threadsRunning, len(threadsWait), threadsCreated, threadsCached, connections, abortedClients, abortedConnects, bytesReceived, bytesSent,
		comSelect, comInsert, comUpdate, comDelete, comCommit, comRollback, questions, queries, slowQueries, keyBufferSize, sortBufferSize, joinBufferSize, innodbPagesCreated, innodbPagesRead, innodbPagesWritten,
		innodbRowLockCurrentWaits, innodbBufferPoolReadRequests, innodbBufferPoolWriteRequests, innodbRowsRead, innodbRowsInserted, innodbRowsUpdated, innodbRowsDeleted)

	err = mysql.Execute(dbClient, insertSQL)
	if err != nil {
		log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
		return
	}

	log.Info(fmt.Sprintf("Complete check instance %s:%s at %s", ip, port, time.Now()))

}

func scanMysql() {
	rows, err := mysql.QueryAll(dbClient, "select ip,port,user,pass,module_name,cluster_name,env_name,idc_name from meta_nodes a join meta_clusters b on a.cluster_id=b.id join meta_modules c on b.module_id=c.id join meta_hosts d on a.ip=d.ip_address join meta_envs e on d.env_id=e.id join meta_idcs f on d.idc_id=f.id where a.monitor=1 and c.module_name='MySQL'")
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
		go collectorMySQL(row["module_name"].(string), row["env_name"].(string), row["ip"].(string), row["port"].(string), row["user"].(string), origPass, row["cluster_name"].(string))
	}
}

func main() {
	startTime := time.Now()
	fmt.Println("Server start at ", startTime)
	log.Info(fmt.Sprintln("Server start at ", startTime))

	//for true {
	//	scanMysql()
	//	time.Sleep(time.Duration(conv.StrToInt(conf.Option["interval"])) * time.Second)
	//}
	scanMysql()
	time.Sleep(time.Duration(5) * time.Second)
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
