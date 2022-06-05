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
	"lepus/src/libary/postgres"
	"lepus/src/libary/tool"
	"lepus/src/libary/utils"
	"strings"
	"time"
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_postgres_mon.log", conv.StrToInt(conf.Option["debug"]))
var dbClient = mysql.InitConnect()

var (
	queryVersionSQL            = "select version() as version limit 1"
	queryStartTimeSQL          = "select pg_postmaster_start_time();"
	queryMaxConnectionsSQL     = "show max_connections"
	queryConnectionsSQL        = "select count(*) from pg_stat_activity"
	queryConnectionsDetailSQL  = "select client_addr,datname,usename,count(*) count from pg_stat_activity group by client_addr,datname,usename order by count desc"
	queryActiveDetailSQL       = "select datname,usename,application_name,client_addr,client_hostname,backend_start,query_start,state,query from pg_stat_activity where state='active' and pid<> pg_backend_pid() order by query_start asc"
	queryPreparedXactDetailSQL = "select * from pg_prepared_xacts"
	queryCheckpointSQL         = "SELECT (100 * checkpoints_req) / (checkpoints_timed + checkpoints_req) AS checkpoints_req_pct,pg_size_pretty(buffers_checkpoint * block_size / (checkpoints_timed + checkpoints_req)) AS avg_checkpoint_write,pg_size_pretty(block_size * (buffers_checkpoint + buffers_clean + buffers_backend)) AS total_written,\n100 * buffers_checkpoint / (buffers_checkpoint + buffers_clean + buffers_backend) AS checkpoint_write_pct, 100 * buffers_backend / (buffers_checkpoint + buffers_clean + buffers_backend) AS backend_write_pct FROM pg_stat_bgwriter,(SELECT cast(current_setting('block_size') AS integer) AS block_size) AS bs;"
	queryStatDatabaseSQL       = "select sum(xact_commit) xact_commit,sum(xact_rollback) xact_rollback,sum(tup_returned) tup_returned,sum(tup_fetched) tup_fetched,sum(tup_inserted) tup_inserted,sum(tup_updated) tup_updated,sum(tup_deleted) tup_deleted,sum(conflicts) conflicts,sum(deadlocks) deadlocks from pg_stat_database"
)

func collectorPostgres(dbType, dbGroup, ip, port, user, pass, tag string) {
	log.Info(fmt.Sprintf("Start check instance %s:%s at %s", ip, port, time.Now()))
	pgdb, err := postgres.NewConnect(ip, port, user, pass, "postgres")
	eventEntity := fmt.Sprintf("%s:%s", ip, port)

	if err != nil {
		log.Error(fmt.Sprintf("Can't connect to postgres database on %s:%s, %s", ip, port, err))
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

		insertSQL := fmt.Sprintf("insert into dashboard_postgresql(host,port,tag,connect,error_info) values('%s','%s','%s','%d','%s')", ip, port, tag, 0, errInfo)
		err = mysql.Execute(dbClient, insertSQL)
		if err != nil {
			log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
			return
		}
		return
	}

	defer pgdb.Close()

	connect := 1
	row := pgdb.QueryRow(queryVersionSQL)
	var version string
	if err := row.Scan(&version); err != nil {
		log.Error(fmt.Sprintf("Can't scan version on %s:%s, %s ", ip, port, err))
		return
	}
	version = strings.Fields(version)[0] + "-" + strings.Fields(version)[1]

	queryStartTime, _ := postgres.QueryAll(pgdb, queryStartTimeSQL)
	startTime := queryStartTime[0]["pg_postmaster_start_time"]
	uptime := time.Now().Unix() - startTime.(time.Time).Unix()
	queryMaxConnections, _ := postgres.QueryAll(pgdb, queryMaxConnectionsSQL)
	maxConnections := queryMaxConnections[0]["max_connections"]

	queryConnections, _ := postgres.QueryAll(pgdb, queryConnectionsSQL)
	connections := queryConnections[0]["count"]

	queryActiveDetail, _ := postgres.QueryAll(pgdb, queryActiveDetailSQL)
	queryPreparedXactDetail, _ := postgres.QueryAll(pgdb, queryPreparedXactDetailSQL)

	queryCheckkpoint, _ := postgres.QueryAll(pgdb, queryCheckpointSQL)
	checkpointsReqPct := queryCheckkpoint[0]["checkpoints_req_pct"]
	avgCheckpointWrite := strings.Replace(queryCheckkpoint[0]["avg_checkpoint_write"].(string), "bytes", "", -1)
	totalWritten := strings.Replace(queryCheckkpoint[0]["total_written"].(string), "kB", "", -1)
	checkpointWritePct := queryCheckkpoint[0]["checkpoint_write_pct"]
	backendWritePct := queryCheckkpoint[0]["backend_write_pct"]

	statDatabasePrev, err := postgres.QueryAll(pgdb, queryStatDatabaseSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query pg_stat_database on %s:%s, %s", ip, port, err))
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	statDatabase, err := postgres.QueryAll(pgdb, queryStatDatabaseSQL)
	if err != nil {
		log.Error(fmt.Sprintf("Can't query pg_stat_database on %s:%s, %s", ip, port, err))
		return
	}
	tupFetched := conv.StrToInt(statDatabase[0]["tup_fetched"].(string)) - conv.StrToInt(statDatabasePrev[0]["tup_fetched"].(string))
	tupReturned := conv.StrToInt(statDatabase[0]["tup_returned"].(string)) - conv.StrToInt(statDatabasePrev[0]["tup_returned"].(string))
	tupInserted := conv.StrToInt(statDatabase[0]["tup_inserted"].(string)) - conv.StrToInt(statDatabasePrev[0]["tup_inserted"].(string))
	tupDeleted := conv.StrToInt(statDatabase[0]["tup_deleted"].(string)) - conv.StrToInt(statDatabasePrev[0]["tup_deleted"].(string))
	tupUpdated := conv.StrToInt(statDatabase[0]["tup_updated"].(string)) - conv.StrToInt(statDatabasePrev[0]["tup_updated"].(string))
	xactCommit := conv.StrToInt(statDatabase[0]["xact_commit"].(string)) - conv.StrToInt(statDatabasePrev[0]["xact_commit"].(string))
	xactRollback := conv.StrToInt(statDatabase[0]["xact_rollback"].(string)) - conv.StrToInt(statDatabasePrev[0]["xact_rollback"].(string))
	conflicts := conv.StrToInt(statDatabase[0]["conflicts"].(string)) - conv.StrToInt(statDatabasePrev[0]["conflicts"].(string))
	deadlocks := conv.StrToInt(statDatabase[0]["deadlocks"].(string)) - conv.StrToInt(statDatabasePrev[0]["deadlocks"].(string))

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
		"event_key":    "connections",
		"event_value":  connections,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "activeSQL",
		"event_value":  len(queryActiveDetail),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "prepared_xacts",
		"event_value":  len(queryPreparedXactDetail),
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "tup_fetched",
		"event_value":  tupFetched,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "tup_returned",
		"event_value":  tupReturned,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "tup_inserted",
		"event_value":  tupInserted,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "tup_deleted",
		"event_value":  tupDeleted,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "tup_updated",
		"event_value":  tupUpdated,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "xact_commit",
		"event_value":  xactCommit,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "xact_rollback",
		"event_value":  xactRollback,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "conflicts",
		"event_value":  conflicts,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "deadlocks",
		"event_value":  deadlocks,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	_, err = http.Post(conf.Option["proxy"], events)
	if err != nil {
		log.Error(fmt.Sprintln("Send events to proxy error:", err))
	}

	insertSQL := fmt.Sprintf("insert into dashboard_postgresql(host,port,tag,connect,start_time,uptime,version,max_connections,connections,active_sql,prepared_xacts,tup_fetched,tup_returned,tup_inserted,tup_deleted,tup_updated,xact_commit,xact_rollback,conflicts,deadlocks,checkpoint_req_pct,checkpoint_avg_write,checkpoint_total_write,checkpoint_write_pct,checkpoint_backend_write_pct)"+
		"values('%s','%s','%s','%d','%s','%d','%s','%s','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%s','%s','%d','%d')", ip, port, tag, connect, startTime, uptime, version, maxConnections, connections, len(queryActiveDetail), len(queryPreparedXactDetail), tupFetched, tupReturned, tupInserted, tupDeleted, tupUpdated, xactCommit, xactRollback, conflicts, deadlocks, checkpointsReqPct, avgCheckpointWrite, totalWritten, checkpointWritePct, backendWritePct)

	err = mysql.Execute(dbClient, insertSQL)
	if err != nil {
		log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
		return
	}

	log.Info(fmt.Sprintf("Complete check instance %s:%s at %s", ip, port, time.Now()))

}

func scanPostgres() {
	rows, err := mysql.QueryAll(dbClient, "select ip,port,user,pass,module_name,cluster_name,env_name,idc_name from meta_nodes a join meta_clusters b on a.cluster_id=b.id join meta_modules c on b.module_id=c.id join meta_hosts d on a.ip=d.ip_address join meta_envs e on d.env_id=e.id join meta_idcs f on d.idc_id=f.id where a.monitor=1 and c.module_name='PostgreSQL'")
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
		go collectorPostgres(row["module_name"].(string), row["env_name"].(string), row["ip"].(string), row["port"].(string), row["user"].(string), origPass, row["cluster_name"].(string))
	}
}

func main() {
	startTime := time.Now()
	fmt.Println("Server start at ", startTime)
	log.Info(fmt.Sprintln("Server start at ", startTime))

	//for true {
	//	scanPostgres()
	//	time.Sleep(time.Duration(conv.StrToInt(conf.Option["interval"])) * time.Second)
	//}
	scanPostgres()
	time.Sleep(time.Duration(8) * time.Second)
	defer dbClient.Close()
}
