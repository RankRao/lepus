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
	"lepus/src/libary/tool"
	"lepus/src/libary/utils"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_mongodb_mon.log", conv.StrToInt(conf.Option["debug"]))
var dbClient = mysql.InitConnect()

func collectorMongoDB(dbType, dbGroup, ip, port, user, pass, tag string) {
	log.Info(fmt.Sprintf("Start check instance %s:%s at %s", ip, port, time.Now()))
	var mongoUrl string
	if user != "" && pass != "" {
		mongoUrl = fmt.Sprintf("%s:%s@%s:%s", user, pass, ip, port)
	} else {
		mongoUrl = fmt.Sprintf("%s:%s", ip, port)
	}
	mongodb, err := mgo.Dial(mongoUrl)
	eventEntity := fmt.Sprintf("%s:%s", ip, port)

	if err != nil {
		log.Error(fmt.Sprintf("Can't connect to mongodb database on %s:%s, %s", ip, port, err))
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

		insertSQL := fmt.Sprintf("insert into dashboard_mongodb(host,port,tag,connect,error_info) values('%s','%s','%s','%d','%s')", ip, port, tag, 0, errInfo)
		err = mysql.Execute(dbClient, insertSQL)
		if err != nil {
			log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
			return
		}
		return
	}

	defer mongodb.Close()

	result := bson.M{}
	if mongodb.DB("admin").Run("serverStatus", &result); err != nil {
		log.Error(fmt.Sprintf("Can't run server status on %s:%s, %s, ", ip, port, err))
		return
	}

	time.Sleep(time.Duration(1) * time.Second)

	resultNew := bson.M{}
	if mongodb.DB("admin").Run("serverStatus", &resultNew); err != nil {
		log.Error(fmt.Sprintf("Can't run server status on %s:%s, %s, ", ip, port, err))
		return
	}

	connect := 1
	ok := result["ok"]
	version := result["version"]
	uptime := result["uptime"]
	//fmt.Print(uptime.(float64))
	connections := result["connections"].(bson.M)
	connectionsCurrent := connections["current"]
	connectionsAvailable := connections["available"]
	mem := result["mem"].(bson.M)
	memBits := mem["bits"]
	memResident := mem["resident"]
	memVirtual := mem["virtual"]
	memSupported := mem["supported"]
	memMapped := mem["mapped"]
	memMappedWithJournal := mem["mappedWithJournal"]
	var memSupportedInt = 0
	if memSupported.(bool) {
		memSupportedInt = 1
	}
	network := result["network"].(bson.M)
	opcounters := result["opcounters"].(bson.M)
	networkNew := resultNew["network"].(bson.M)
	opcountersNew := resultNew["opcounters"].(bson.M)

	networkBytesIn := networkNew["bytesIn"].(int64) - network["bytesIn"].(int64)
	networkBytesOut := networkNew["bytesOut"].(int64) - network["bytesOut"].(int64)
	networkNumRequests := networkNew["numRequests"].(int64) - network["numRequests"].(int64)

	var (
		opcountersInsert  interface{}
		opcountersQuery   interface{}
		opcountersUpdate  interface{}
		opcountersDelete  interface{}
		opcountersCommand interface{}
		opcountersTotal   interface{}
	)
	opcountersType := fmt.Sprintf("%T", opcounters["insert"])
	if opcountersType == "int64" {
		opcountersInsert = opcountersNew["insert"].(int64) - opcounters["insert"].(int64)
		opcountersQuery = opcountersNew["query"].(int64) - opcounters["query"].(int64)
		opcountersUpdate = opcountersNew["update"].(int64) - opcounters["update"].(int64)
		opcountersDelete = opcountersNew["delete"].(int64) - opcounters["delete"].(int64)
		opcountersCommand = opcountersNew["command"].(int64) - opcounters["command"].(int64)
		opcountersTotal = (opcountersNew["insert"].(int64) - opcounters["insert"].(int64)) + (opcountersNew["query"].(int64) - opcounters["query"].(int64)) + (opcountersNew["update"].(int64) - opcounters["update"].(int64)) + (opcountersNew["delete"].(int64) - opcounters["delete"].(int64)) + (opcountersNew["command"].(int64) - opcounters["command"].(int64))
	} else {
		opcountersInsert = opcountersNew["insert"].(int) - opcounters["insert"].(int)
		opcountersQuery = opcountersNew["query"].(int) - opcounters["query"].(int)
		opcountersUpdate = opcountersNew["update"].(int) - opcounters["update"].(int)
		opcountersDelete = opcountersNew["delete"].(int) - opcounters["delete"].(int)
		opcountersCommand = opcountersNew["command"].(int) - opcounters["command"].(int)
		opcountersTotal = (opcountersNew["insert"].(int) - opcounters["insert"].(int)) + (opcountersNew["query"].(int) - opcounters["query"].(int)) + (opcountersNew["update"].(int) - opcounters["update"].(int)) + (opcountersNew["delete"].(int) - opcounters["delete"].(int)) + (opcountersNew["command"].(int) - opcounters["command"].(int))
	}

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
		"event_key":    "connectionsCurrent",
		"event_value":  connectionsCurrent,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "connectionsAvailable",
		"event_value":  connectionsAvailable,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "memBits",
		"event_value":  memBits,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "memResident",
		"event_value":  memResident,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "memVirtual",
		"event_value":  memVirtual,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "memMapped",
		"event_value":  memMapped,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "memMappedWithJournal",
		"event_value":  memMappedWithJournal,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "networkBytesIn",
		"event_value":  networkBytesIn,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "networkBytesIn",
		"event_value":  networkBytesIn,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "networkNumRequests",
		"event_value":  networkNumRequests,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "opcountersInsert",
		"event_value":  opcountersInsert,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "opcountersQuery",
		"event_value":  opcountersQuery,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "opcountersUpdate",
		"event_value":  opcountersUpdate,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "opcountersDelete",
		"event_value":  opcountersDelete,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "opcountersCommand",
		"event_value":  opcountersCommand,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   dbType,
		"event_group":  dbGroup,
		"event_entity": eventEntity,
		"event_key":    "opcounters",
		"event_value":  opcountersTotal,
		"event_tag":    tag,
		"event_unit":   "",
	}
	events = append(events, event)
	_, err = http.Post(conf.Option["proxy"], events)
	if err != nil {
		log.Error(fmt.Sprintln("Send events to proxy error:", err))
	}

	if memMapped == nil {
		memMapped = 0
	}
	if memMappedWithJournal == nil {
		memMappedWithJournal = 0
	}

	insertSQL := fmt.Sprintf("insert into dashboard_mongodb(host,port,tag,connect,ok,uptime,version,connections_current,connections_available,"+
		"mem_bits,mem_resident,mem_virtual,mem_supported,mem_mapped,mem_mappedWithJournal,network_bytesIn,network_bytesOut,network_numRequests,"+
		"opcounters_insert,opcounters_query,opcounters_update,opcounters_delete,opcounters_command,opcounters )"+
		"values('%s','%s','%s','%d','%f','%f','%s','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d')", ip, port, tag, connect, ok, uptime, version, connectionsCurrent, connectionsAvailable, memBits, memResident, memVirtual, memSupportedInt, memMapped, memMappedWithJournal,
		networkBytesIn, networkBytesOut, networkNumRequests, opcountersInsert, opcountersQuery, opcountersUpdate, opcountersDelete, opcountersCommand, opcountersTotal)

	err = mysql.Execute(dbClient, insertSQL)
	if err != nil {
		log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
		return
	}

	log.Info(fmt.Sprintf("Complete check instance %s:%s at %s", ip, port, time.Now()))

}

func scanMongoDB() {
	rows, err := mysql.QueryAll(dbClient, "select ip,port,user,pass,module_name,cluster_name,env_name,idc_name from meta_nodes a join meta_clusters b on a.cluster_id=b.id join meta_modules c on b.module_id=c.id join meta_hosts d on a.ip=d.ip_address join meta_envs e on d.env_id=e.id join meta_idcs f on d.idc_id=f.id where a.monitor=1 and c.module_name='MongoDB'")
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
		go collectorMongoDB(row["module_name"].(string), row["env_name"].(string), row["ip"].(string), row["port"].(string), row["user"].(string), origPass, row["cluster_name"].(string))
	}
}

func main() {
	startTime := time.Now()
	fmt.Println("Server start at ", startTime)
	log.Info(fmt.Sprintln("Server start at ", startTime))

	//for true {
	//	scanMongoDB()
	//	time.Sleep(time.Duration(conv.StrToInt(conf.Option["interval"])) * time.Second)
	//}
	scanMongoDB()
	time.Sleep(time.Duration(8) * time.Second)
	defer dbClient.Close()
}
