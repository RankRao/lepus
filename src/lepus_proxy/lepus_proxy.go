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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"

	"lepus/src/libary/conf"
	"lepus/src/libary/logger"
	"lepus/src/libary/mysql"
	"lepus/src/libary/tool"

	influx "github.com/influxdata/influxdb1-client/v2"
)

/*
init log
*/
var log = logger.InitLog()

/*
init kafka client
*/
//var kafkaClient = kafka.InitClient()

/*
init mysql db
*/
var db = mysql.InitConnect()

type Adapter struct {
	pool sync.Pool
}

func New() *Adapter {
	return &Adapter{
		pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 4096))
			},
		},
	}
}

type EventSlice []struct {
	EventTime   string  `json:"event_time"`   //事件发生时间
	EventType   string  `json:"event_type"`   //事件类型
	EventGroup  string  `json:"event_group"`  //事件分组
	EventEntity string  `json:"event_entity"` //事件实体
	EventKey    string  `json:"event_key"`    //事件指标
	EventValue  float64 `json:"event_value"`  //事件数据
	EventTag    string  `json:"event_tag"`    //事件标签
	EventUnit   string  `json:"event_unit"`   //事件单位
}

func (p *EventSlice) UnmarshalJsonList(data []byte) error {
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	return nil
}

func (api *Adapter) sendEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		buffer := api.pool.Get().(*bytes.Buffer)
		buffer.Reset()
		defer func() {
			if buffer != nil {
				api.pool.Put(buffer)
				buffer = nil
			}
		}()

		_, err := io.Copy(buffer, r.Body)
		if err != nil {
			log.Error(fmt.Sprintln("Io copy error:", err))
			return
		}
		/*
			convert json str to array map
		*/
		data := EventSlice{}
		data.UnmarshalJsonList([]byte(buffer.String()))

		/*
			producer, err := kafka.NewProducer(kafkaClient)
			if err != nil {
				log.Error(fmt.Sprintln("Create new kafka producer err:", err))
			}
			defer func() {
				if producer != nil {
					producer.Close()
				}
			}()
		*/

		producer, err := nsq.NewProducer(conf.Option["nsq_server"], nsq.NewConfig())
		if err != nil {
			panic(err)
		}
		defer func() {
			if producer != nil {
				producer.Stop()
			}
		}()

		for _, val := range data {
			/*
				init event map
			*/
			uuid := tool.GetUUID()
			eventValue, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", val.EventValue), 64)
			m := map[string]interface{}{
				"event_uuid":   uuid,
				"event_time":   val.EventTime,
				"event_type":   val.EventType,
				"event_group":  val.EventGroup,
				"event_entity": val.EventEntity,
				"event_key":    val.EventKey,
				"event_value":  eventValue,
				"event_tag":    val.EventTag,
				"event_unit":   val.EventUnit,
			}
			/*
				convert event map to json string
			*/
			d, _ := json.Marshal(m)
			eventStr := string(d)
			/*
				send event json str to kafka
			*/
			log.Debug(fmt.Sprintf("Event map data:%s", m))
			log.Debug(fmt.Sprintf("Event json data:%s", eventStr))
			//producer.Input() <- &sarama.ProducerMessage{Topic: "lepus_events", Key: nil, Value: sarama.StringEncoder(eventStr)}
			if err := producer.Publish("lepus_events", []byte(eventStr)); err != nil { // 发布消息
				panic(err)
			}

			insertEventSql := fmt.Sprintf("insert into events(event_uuid,event_time,event_type,event_group,event_entity,event_key,event_value,event_tag,event_unit) values('%s','%s','%s','%s','%s','%s','%f','%s','%s')", uuid, val.EventTime, val.EventType, val.EventGroup, val.EventEntity, val.EventKey, eventValue, val.EventTag, val.EventUnit)
			err = mysql.Execute(db, insertEventSql)
			if err != nil {
				log.Error(fmt.Sprintln("Can't insert event data to mysql database, ", err))
				return
			}

			enableInfluxdb := conf.Option["enable_influxdb"]
			if enableInfluxdb == "1" {
				cli, err := influx.NewHTTPClient(influx.HTTPConfig{
					Addr:     fmt.Sprintf("http://%s:%s", conf.Option["influx_host"], conf.Option["influx_port"]),
					Username: conf.Option["influx_user"],
					Password: conf.Option["influx_password"],
				})
				if err != nil {
					log.Error(fmt.Sprintln("Can't connect influxdb database, ", err))
				}
				defer cli.Close()

				bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
					Database:  conf.Option["influx_database"],
					Precision: "s", //精度，默认ns
				})
				if err != nil {
					log.Error(fmt.Sprintln("Can't select influxdb database, ", err))
				}
				tags := map[string]string{"event_uuid": uuid, "event_time": val.EventTime, "event_type": val.EventType, "event_group": val.EventGroup, "event_entity": val.EventEntity, "event_key": val.EventKey, "event_tag": val.EventTag}
				fields := map[string]interface{}{"event_time": val.EventTime, "event_type": val.EventType, "event_group": val.EventGroup, "event_entity": val.EventEntity, "event_tag": val.EventTag, "event_key": val.EventKey, "event_value": eventValue, "event_unit": val.EventUnit}
				pt, err := influx.NewPoint("events", tags, fields, time.Now())
				if err != nil {
					log.Error(fmt.Sprintln("Can't create point on influxdb, ", err))
				}
				bp.AddPoint(pt)
				err = cli.Write(bp)
				if err != nil {
					log.Error(fmt.Sprintln("Can't write event data to influxdb database, ", err))
				}
			}

		}
		api.pool.Put(buffer)
		buffer = nil
		return
	} else {
		fmt.Fprint(w, "Request type must be post.")
	}
}

func (api *Adapter) sendSql(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		buffer := api.pool.Get().(*bytes.Buffer)
		buffer.Reset()
		defer func() {
			if buffer != nil {
				api.pool.Put(buffer)
				buffer = nil
			}
		}()

		_, err := io.Copy(buffer, r.Body)
		if err != nil {
			log.Error(fmt.Sprintln("Io copy error:", err))
			return
		}

		data := buffer.String()
		log.Debug(fmt.Sprintf("receive sql data: %s", data))
		err = mysql.Execute(db, data)
		if err != nil {
			log.Error(fmt.Sprintln("Can't insert sql data to mysql database, ", err))
			return
		}
	}
}

func main() {
	start := time.Now()
	fmt.Printf("Proxy server start on port %s at %s \n", conf.Option["port"], start)
	log.Info(fmt.Sprintf("Proxy server start on port %s at %s", conf.Option["port"], start))

	runtime.GOMAXPROCS(runtime.NumCPU())

	n := New()
	http.HandleFunc("/", n.sendEvent)
	http.HandleFunc("/proxy/event", n.sendEvent)
	http.HandleFunc("/proxy/sql", n.sendSql)
	err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Option["port"]), nil)
	if err != nil {
		log.Error(fmt.Sprintln("Start proxy server err:", err))
	}
}
