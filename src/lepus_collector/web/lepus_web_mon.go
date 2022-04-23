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
	_ "reflect"
	"strings"
	"time"
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_web_mon.log", conv.StrToInt(conf.Option["debug"]))
var dbClient = mysql.InitConnect()

func collectorUrlStatus(name, url, method, clusterName, moduleName, envName string) {
	log.Info(fmt.Sprintf("Start check web url %s:%s at %s", name, url, time.Now()))
	timeStart := time.Now()
	response, err := http.Get(url)
	if err != nil {
		errInfo := fmt.Sprintf("Did not get any response to web url on %s:%s, %s", name, url, err)
		log.Error(errInfo)
		events := make([]map[string]interface{}, 0)
		detail := make([]map[string]string, 0)
		detail = append(detail, map[string]string{"Error": strings.ReplaceAll(errInfo, "\"", "'")})
		event := map[string]interface{}{
			"event_time":   tool.GetNowTime(),
			"event_type":   moduleName,
			"event_group":  envName,
			"event_entity": name,
			"event_key":    "respStatus",
			"event_value":  0,
			"event_tag":    clusterName,
			"event_unit":   "",
			"event_detail": detail,
		}
		events = append(events, event)
		_, err := http.Post(conf.Option["proxy"], events)
		if err != nil {
			log.Error(fmt.Sprintln("Send events to proxy error:", err))
		}

		insertSQL := fmt.Sprintf("insert into dashboard_web(name,url,tag,resp_status) values('%s','%s','%s','%d')", name, url, clusterName, 0)
		err = mysql.Execute(dbClient, insertSQL)
		if err != nil {
			log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
			return
		}
		return
	}
	defer response.Body.Close()
	//fmt.Println(response)
	timeEnd := time.Now()
	httpTime := timeEnd.Sub(timeStart).Milliseconds()
	httpCode := response.StatusCode
	httpProto := response.Proto

	events := make([]map[string]interface{}, 0)

	event := map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   moduleName,
		"event_group":  envName,
		"event_entity": name,
		"event_key":    "respStatus",
		"event_value":  1,
		"event_tag":    clusterName,
		"event_unit":   "",
		"event_detail": "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   moduleName,
		"event_group":  envName,
		"event_entity": name,
		"event_key":    "HttpCode",
		"event_value":  httpCode,
		"event_tag":    clusterName,
		"event_unit":   "",
		"event_detail": "",
	}
	events = append(events, event)

	event = map[string]interface{}{
		"event_time":   tool.GetNowTime(),
		"event_type":   moduleName,
		"event_group":  envName,
		"event_entity": name,
		"event_key":    "httpTime",
		"event_value":  httpTime,
		"event_tag":    clusterName,
		"event_unit":   "毫秒",
		"event_detail": "",
	}
	events = append(events, event)

	_, err = http.Post(conf.Option["proxy"], events)
	if err != nil {
		log.Error(fmt.Sprintln("Send events to proxy error:", err))
	}

	insertSQL := fmt.Sprintf("insert into dashboard_web(name,url,tag,resp_status,http_code,http_proto,http_time) values('%s','%s','%s','%d','%d','%s','%d')", name, url, clusterName, 1, httpCode, httpProto, httpTime)

	err = mysql.Execute(dbClient, insertSQL)
	if err != nil {
		log.Error(fmt.Sprintln("Can't insert data to mysql database, ", err))
		return
	}

	log.Info(fmt.Sprintf("Complete check web url %s:%s at %s", name, url, time.Now()))

}

func scanWebUrl() {
	rows, err := mysql.QueryAll(dbClient, "select name,url,method,cluster_name,module_name,env_name from meta_webs a join meta_clusters b on a.cluster_id=b.id join meta_modules c on b.module_id=c.id join meta_envs d on a.env_id=d.id where monitor=1 and module_name='Web'")
	if err != nil {
		log.Error(fmt.Sprintln("Can't query mysql database, ", err))
		return
	}
	for _, row := range rows {
		go collectorUrlStatus(row["name"].(string), row["url"].(string), row["method"].(string), row["cluster_name"].(string), row["module_name"].(string), row["env_name"].(string))
	}
}
func main() {
	startTime := time.Now()
	fmt.Println("Server start at ", startTime)
	log.Info(fmt.Sprintln("Server start at ", startTime))
	scanWebUrl()
	time.Sleep(time.Duration(5) * time.Second)
	defer dbClient.Close()
}
