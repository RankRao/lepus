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
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/nsqio/go-nsq"

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
	eventUuid := event["event_uuid"].(string)
	eventTime := event["event_time"].(string)
	eventType := event["event_type"].(string)
	eventGroup := event["event_group"].(string)
	eventKey := event["event_key"].(string)
	eventEntity := event["event_entity"].(string)
	eventValue := utils.FormatFloat64(event["event_value"].(float64))
	eventTag := event["event_tag"].(string)
	eventUnit := event["event_unit"].(string)

	ruleId := rule["id"].(string)
	alarmTitle := rule["title"].(string)
	alarmRule := rule["alarm_rule"].(string)
	alarmValue := rule["alarm_value"].(string)
	alarmSleep := conv.StrToInt(rule["alarm_sleep"].(string))
	alarmTimes := conv.StrToInt(rule["alarm_times"].(string))
	levelId := conv.StrToInt(rule["level_id"].(string))
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
				log.Info(fmt.Sprintf("Set alarm at key %s", alarmAtKeyName))
			}
		}
		alarmCountInt := conv.StrToInt(alarmCount)
		if alarmCountInt < alarmTimes {
			var alarmLevel = ""
			getLevelSql := fmt.Sprintf("select level_name from alarm_levels where enable=1 and id=%d ", levelId)
			levelList, _ := mysql.QueryAll(db, getLevelSql)
			if len(levelList) > 0 {
				if levelList[0]["level_name"] != nil {
					alarmLevel = levelList[0]["level_name"].(string)
				} else {
					alarmLevel = ""
				}
			}
			var (
				sendMail    = 0
				sendWebhook = 0
			)
			getChannelSql := fmt.Sprintf("select name,mail_enable,webhook_enable,mail_list,webhook_url from alarm_channels where enable=1 and id=%d ", channelId)
			channelList, _ := mysql.QueryAll(db, getChannelSql)
			if len(channelList) > 0 {
				rds.Incr(alarmCountKeyName)
				rds.Expire(alarmCountKeyName, time.Second*time.Duration(alarmSleep))
				log.Info(fmt.Sprintf("Set alarm count key %s", alarmCountKeyName))
				channel := channelList[0]
				mailEnable := utils.StrToInt(channel["mail_enable"].(string))
				webhookEnable := utils.StrToInt(channel["webhook_enable"].(string))
				mailList := channel["mail_list"].(string)
				webhookUrl := channel["webhook_url"].(string)
				if mailEnable == 1 && mailList != "" {
					log.Info(fmt.Sprintf("Start to send email to %s", mailList))
					mailTo := strings.Split(mailList, ";")
					tableTitle := "事件概览"
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
					data = make([]string, 0)
					data = append(data, "触发规则", fmt.Sprintf("%s%s%s", eventKey, alarmRule, alarmValue))
					dataList = append(dataList, data)

					eventContent := html.CreateTable(tableTitle, tableHeader, dataList)

					mailHello := fmt.Sprintf("尊敬的用户：<p></p>您好！您收到一条【%s】事件：【%s】，请您及时关注和处理。", alarmLevel, alarmTitle)
					mailContent := "<span style='margin-top:1px;'>" + mailHello + "</span><p></p>" + eventContent + "</div><div style='margin-top:30px; color:#666'><hr color='#ccc' style='border:1px dashed #cccccc;' />本邮件来自Lepus实时事件告警组件，请勿直接回复本邮件。如需获得技术支持，可联系我们：<a href='https://www.lepus.cc' target='_blank'>https://www.lepus.cc</a></div>"
					//fmt.Println(mailContent)
					//return
					mailTitle := fmt.Sprintf("[%s][%s]%s", alarmLevel, eventEntity, alarmTitle)
					if err := mail.Send(mailTo, mailTitle, mailContent); err != nil {
						sendMail = 2
						log.Error(fmt.Sprintf("Failed to send email %s,%s: %s", mailTitle, mailList, err))
					} else {
						sendMail = 1
						log.Info(fmt.Sprintf("Success to send email %s,%s", mailTitle, mailList))
					}
				}

				if webhookEnable == 1 && webhookUrl != "" {
					log.Info(fmt.Sprintf("Start to call webhook to %s", webhookUrl))
					eventData := make(map[string]interface{})
					//webhookUrl 中包含 "oapi.dingtalk.com" 则构造钉钉自定义机器人的post数据
					if strings.Contains(webhookUrl, "oapi.dingtalk.com") {
						//init dingtalk config
						alarm_url := conf.Option["alarm_url"]
						eventData = map[string]interface{}{
							"msgtype": "markdown",
							"markdown": map[string]string{
								"title": alarmTitle,
								"text": "#### **" + alarmTitle + "** \n " +
									"> ##### 事件时间: " + eventTime + " \n " +
									"> ##### 事件类型: " + eventType + " \n " +
									"> ##### 事件组别: " + eventGroup + " \n " +
									"> ##### 事件实体: " + eventEntity + " \n " +
									"> ##### 事件标签: " + eventTag + " \n " +
									"> ##### 事件指标: " + eventKey + " \n " +
									"> ##### 事件数值: " + utils.FloatToStr(eventValue) + eventUnit + " \n " +
									"> ##### 触发规则: " + eventKey + alarmRule + alarmValue + " \n " +
									"###### [查看详情](" + alarm_url + ") \n"},
						}
					} else {
						//post数据
						eventData = map[string]interface{}{
							"alarm_title":  alarmTitle,
							"alarm_rule":   alarmRule,
							"alarm_value":  alarmValue,
							"event_time":   eventTime,
							"event_type":   eventType,
							"event_group":  eventGroup,
							"event_entity": eventEntity,
							"event_key":    eventKey,
							"event_value":  eventValue,
							"event_tag":    eventTag,
						}
					}
					log.Debug(fmt.Sprintf("webhook post content %s", eventData))
					client := &http.Client{Timeout: 3 * time.Second}
					jsonStr, _ := json.Marshal(eventData)
					resp, err := client.Post(webhookUrl, "application/json", bytes.NewBuffer(jsonStr))
					if err != nil {
						sendWebhook = 2
						log.Error(fmt.Sprintf("Failed to call webhook %s: %s", webhookUrl, err))
					} else {
						sendWebhook = 1
						log.Info(fmt.Sprintf("Success to call webhook %s", webhookUrl))
						resp.Body.Close()
					}
				}
			}

			insertAlarmSql := fmt.Sprintf("insert into alarm_events(alarm_title,alarm_level,alarm_rule,alarm_value,event_uuid,event_time,event_type,event_group,event_entity,event_key,event_value,event_unit,event_tag,rule_id,level_id,channel_id,send_mail,send_webhook) values('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%f','%s','%s','%s','%d','%d','%d','%d')", alarmTitle, alarmLevel, alarmRule, alarmValue, eventUuid, eventTime, eventType, eventGroup, eventEntity, eventKey, eventValue, eventUnit, eventTag, ruleId, levelId, channelId, sendMail, sendWebhook)
			err := mysql.Execute(db, insertAlarmSql)
			if err != nil {
				log.Error(fmt.Sprintln("Failed insert alarm event data to database, ", err))
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
		sql = fmt.Sprintf("select id,title,alarm_rule,alarm_value,alarm_sleep,alarm_times,level_id,channel_id from alarm_rules "+
			"where enable=1 and event_type='%s' and event_key='%s' and event_entity='%s'  order by level_id asc", eventType, eventKey, eventEntity)
		res, _ := mysql.QueryAll(db, sql)
		if len(res) > 0 {
			return res
		}
	}
	if eventGroup != "" {
		sql = fmt.Sprintf("select id,title,alarm_rule,alarm_value,alarm_sleep,alarm_times,level_id,channel_id from alarm_rules "+
			"where enable=1 and  event_type='%s' and event_key='%s' and event_group='%s'  order by level_id asc", eventType, eventKey, eventGroup)
		res, _ := mysql.QueryAll(db, sql)
		if len(res) > 0 {
			return res
		}
	}

	sql = fmt.Sprintf("select id,title,alarm_rule,alarm_value,alarm_sleep,alarm_times,level_id,channel_id from alarm_rules "+
		"where enable=1 and  event_type='%s' and event_key='%s'  order by level_id asc", eventType, eventKey)
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
		kafka消息消费代码，目前已经使用nsq代替，代码暂时保留
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
