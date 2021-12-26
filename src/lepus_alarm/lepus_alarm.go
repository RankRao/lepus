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
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"lepus/src/libary/conf"
	"lepus/src/libary/conv"
	"lepus/src/libary/html"
	_ "lepus/src/libary/kafka"
	"lepus/src/libary/logger"
	"lepus/src/libary/mail"
	"lepus/src/libary/mysql"
	"lepus/src/libary/redis"
	"lepus/src/libary/utils"
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

/*
init redis cache
*/
var rds = redis.InitClient()

/*
send alarm function
*/
func sendAlarm(event, rule map[string]interface{}, match bool) {
	eventTime := event["event_time"].(string)
	eventType := event["event_type"].(string)
	eventGroup := event["event_group"].(string)
	eventKey := event["event_key"].(string)
	eventEntity := event["event_entity"].(string)
	eventValue := event["event_value"].(float64)
	eventTag := event["event_tag"].(string)
	eventUnit := event["event_unit"].(string)

	ruleId := rule["id"].(string)
	alarmTitle := rule["title"].(string)
	alarmLevel := rule["alarm_level"].(string)
	alarmSleep := conv.StrToInt(rule["alarm_sleep"].(string))
	alarmTimes := conv.StrToInt(rule["alarm_times"].(string))
	channelId := conv.StrToInt(rule["channel_id"].(string))

	keyName := fmt.Sprintf("%s:%s:%s:%s", eventType, eventKey, eventTag, eventEntity)
	alarmCountKeyName := "alarm_count." + keyName
	alarmAtKeyName := "alarm_at." + keyName

	if match == true {
		alarmCount, _ := rds.Get(alarmCountKeyName).Result()
		alarmAt, _ := rds.Get(alarmAtKeyName).Result()
		if alarmCount == "" {
			alarmCount = "0"
			if alarmAt == "" {
				rds.Set(alarmAtKeyName, time.Now().Unix(), time.Hour*time.Duration(72))
			}
		}
		alarmCountInt := conv.StrToInt(alarmCount)
		if alarmCountInt < alarmTimes {
			var (
				sendMail  = 0
				sendPhone = 0
			)
			sql := fmt.Sprintf("select name,mail_list,phone_list from alarm_channels where enable=1 and id=%d ", channelId)
			channelList, _ := mysql.QueryAll(db, sql)
			if len(channelList) > 0 {
				rds.Incr(alarmCountKeyName)
				rds.Expire(alarmCountKeyName, time.Second*time.Duration(alarmSleep))
				channel := channelList[0]
				mailList := channel["mail_list"].(string)
				phoneList := channel["phone_list"].(string)
				if mailList != "" {
					mailTo := strings.Split(mailList, ";")
					tableHeader := []string{"名称", "内容"}
					dataList := make([][]string, 0)
					data := make([]string, 0)
					data = append(data, "事件时间", eventTime)
					dataList = append(dataList, data)
					data = make([]string, 0)
					data = append(data, "事件类型", eventType)
					dataList = append(dataList, data)
					data = make([]string, 0)
					data = append(data, "事件组别", eventGroup)
					dataList = append(dataList, data)
					data = make([]string, 0)
					data = append(data, "事件实体", eventEntity)
					dataList = append(dataList, data)
					data = make([]string, 0)
					data = append(data, "事件指标", eventKey)
					dataList = append(dataList, data)
					data = make([]string, 0)
					data = append(data, "事件标签", eventTag)
					dataList = append(dataList, data)
					data = make([]string, 0)
					data = append(data, "事件数值", utils.FloatToStr(eventValue)+eventUnit)
					dataList = append(dataList, data)
					tableTitle := alarmTitle
					eventContent := html.CreateTable(alarmTitle, tableTitle, tableHeader, dataList)

					mailContent := eventContent
					//fmt.Println(mailContent)
					//return
					mailTitle := fmt.Sprintf("[%s][%s]%s", alarmLevel, eventEntity, alarmTitle)
					log.Info(fmt.Sprintln("Start to send mail :", mailTitle))
					if err := mail.Send(mailTo, mailTitle, mailContent); err != nil {
						sendMail = -1
						log.Error(fmt.Sprintln("send alarm mail error:", err))
					} else {
						sendMail = 1
					}
				}
				if phoneList != "" {
					log.Info(fmt.Sprintln("Start to fake send phone :", phoneList))
				}
			}
			log.Info(fmt.Sprintln("Insert alarm event data to mysql database"))
			insertAlarmSql := fmt.Sprintf("insert into alarm_events(alarm_title,alarm_level,event_time,event_type,event_group,event_entity,event_key,event_value,event_tag,rule_id,send_mail,send_phone) values(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%f\",\"%s\",\"%s\",\"%d\",\"%d\")", alarmTitle, alarmLevel, eventTime, eventType, eventGroup, eventEntity, eventKey, eventValue, eventTag, ruleId, sendMail, sendPhone)
			err := mysql.Execute(db, insertAlarmSql)
			if err != nil {
				log.Error(fmt.Sprintln("Can't insert alarm event data to mysql database, ", err))
				return
			}

		}
	} else {
		rds.Del(alarmCountKeyName)
		rds.Del(alarmAtKeyName)
		return
	}
}

func getAlarmRule(eventType, eventGroup, eventEntity, eventKey string) []map[string]interface{} {
	var sql string
	if eventEntity != "" {
		sql = fmt.Sprintf("select id,title,alarm_rule,alarm_value,alarm_level,alarm_sleep,alarm_times,channel_id from alarm_rules "+
			"where enable=1 and event_type='%s' and event_key='%s' and event_entity='%s'  order by alarm_level asc", eventType, eventKey, eventEntity)
		res, _ := mysql.QueryAll(db, sql)
		if len(res) > 0 {
			return res
		}
	}
	if eventGroup != "" {
		sql = fmt.Sprintf("select id,title,alarm_rule,alarm_value,alarm_level,alarm_sleep,alarm_times,channel_id from alarm_rules "+
			"where enable=1 and  event_type='%s' and event_key='%s' and event_group='%s'  order by alarm_level asc", eventType, eventKey, eventGroup)
		res, _ := mysql.QueryAll(db, sql)
		if len(res) > 0 {
			return res
		}
	}

	sql = fmt.Sprintf("select id,title,alarm_rule,alarm_value,alarm_level,alarm_sleep,alarm_times,channel_id from alarm_rules "+
		"where enable=1 and  event_type='%s' and event_key='%s'  order by alarm_level asc", eventType, eventKey)
	res, err := mysql.QueryAll(db, sql)
	if err != nil {
		log.Error(fmt.Sprintln("query alarm rule err:", err))
	}
	return res
}

func matchAlarmRule(alarmRule string, alarmValue float64, eventValue float64) bool {
	//alarmValueFloat := conv.StrToFloat(alarmValue)
	//eventValueFloat := conv.StrToFloat(eventValue)
	log.Debug(fmt.Sprintf("matchAlarmRule, alarmRule:%s,alarmValue:%f,eventValue:%f", alarmRule, alarmValue, eventValue))
	if alarmRule == "=" && (alarmValue == eventValue) {
		return true
	}
	if alarmRule == "!=" && (alarmValue != eventValue) {
		return true
	}
	if alarmRule == ">" && (eventValue > alarmValue) {
		return true
	}
	if alarmRule == ">=" && (eventValue >= alarmValue) {
		return true
	}
	if alarmRule == "<" && (eventValue < alarmValue) {
		return true
	}
	if alarmRule == "<=" && (eventValue <= alarmValue) {
		return true
	}
	return false
}

func alarm(value string) {
	/*
		convert event json str to  map
	*/
	var event map[string]interface{}
	err := json.Unmarshal([]byte(value), &event)
	if err != nil {
		log.Error(fmt.Sprintln("unmarshal json event value err:", err))
		return
	}
	/*
		.(string) convert interface{} to string
	*/
	eventType := event["event_type"].(string)
	eventGroup := event["event_group"].(string)
	eventKey := event["event_key"].(string)
	//eventTag := event["event_tag"].(string)
	eventEntity := event["event_entity"].(string)
	eventValue := event["event_value"].(float64)

	alarmRuleList := getAlarmRule(eventType, eventGroup, eventEntity, eventKey)
	log.Debug(fmt.Sprintln("get Alarm Rule:", alarmRuleList))
	if len(alarmRuleList) == 0 {
		return
	}
	for _, rule := range alarmRuleList {
		alarmRule := rule["alarm_rule"].(string)
		alarmValue := utils.StrToFloat64(rule["alarm_value"].(string))
		match := matchAlarmRule(alarmRule, alarmValue, eventValue)
		log.Debug(fmt.Sprintln("Alarm match result:", match))
		sendAlarm(event, rule, match)
		if match {
			break
		} else {
			continue
		}
	}
}

//nsq订阅消息
type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	//fmt.Println(string(msg.Body))
	alarm(string(msg.Body))
	return nil
}

func main() {
	start := time.Now()
	fmt.Printf("Alarm server start at %s \n", start)
	log.Info(fmt.Sprintf("Alarm server start at %s", start))

	runtime.GOMAXPROCS(runtime.NumCPU())

	/*
		consumer, err := kafka.NewConsumer(kafkaClient)
		if err != nil {
			log.Error(fmt.Sprintln("Create new kafka consumer err:", err))
		}
		defer consumer.Close()

		partitions, err := consumer.Partitions("lepus_events")
		if err != nil {
			log.Error(fmt.Sprintln("Get partitions info err:", err))
		}
		for _, partitionId := range partitions {
			partitionConsumer, err := consumer.ConsumePartition("lepus_events", partitionId, sarama.OffsetNewest)
			if err != nil {
				log.Error(fmt.Sprintln("Consume partition topic err:", err))
			}
			go func(pc *sarama.PartitionConsumer) {
				defer (*pc).Close()
				//block
				for message := range (*pc).Messages() {
					value := string(message.Value)
					log.Debug(fmt.Sprintln("Kafka message value:", value))
					alarm(value)
				}
			}(&partitionConsumer)
		}
	*/

	consumer, err := nsq.NewConsumer("lepus_events", "lepus-channel", nsq.NewConfig()) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	consumer.AddHandler(&ConsumerT{})                                         // 添加消息处理
	if err := consumer.ConnectToNSQD(conf.Option["nsq_server"]); err != nil { // 建立连接
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
}
