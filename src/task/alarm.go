/*
Copyright 2014-2024 The Lepus Team Group, website: https://www.lepus.cc
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

package task

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/nsqio/go-nsq"

	"dbmcloud/log"
	"dbmcloud/setting"
	"dbmcloud/src/database"
	"dbmcloud/src/libary/conv"
	"dbmcloud/src/libary/html"
	"dbmcloud/src/libary/mail"
	"dbmcloud/src/libary/utils"
	"dbmcloud/src/model"
)

func init() {
	go eventConsumer()
}

func eventConsumer() {
	time.Sleep(time.Second * time.Duration(60))
	start := time.Now()
	fmt.Printf("Alarm server start at %s \n", start)
	log.Logger.Info(fmt.Sprintf("Alarm server start at %s", start))

	runtime.GOMAXPROCS(runtime.NumCPU())

	consumer, err := nsq.NewConsumer("lepus_events", "lepus-channel", nsq.NewConfig()) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	consumer.AddHandler(&ConsumerT{})
	//fmt.Println(setting.Setting.NsqServer)                                    // 添加消息处理
	if err := consumer.ConnectToNSQD(setting.Setting.NsqServer); err != nil { // 建立连接
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
}

// 订阅NSQ消息
type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	//fmt.Println(string(msg.Body))
	alarm(string(msg.Body))
	return nil
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
	eventValue := utils.StrToFloat64(event["event_value"].(string))

	alarmRuleList := getAlarmRule(eventType, eventGroup, eventEntity, eventKey)
	log.Logger.Debug(fmt.Sprintln("get Alarm Rule:", alarmRuleList))
	if len(alarmRuleList) == 0 {
		return
	}
	for _, rule := range alarmRuleList {
		alarmRule := rule["alarm_rule"].(string)
		alarmValue := utils.StrToFloat64(rule["alarm_value"].(string))
		match := matchAlarmRule(alarmRule, alarmValue, eventValue)
		log.Logger.Debug(fmt.Sprintln("Alarm match result:", match))
		sendAlarm(event, rule, match)
		if match {
			break
		} else {
			continue
		}
	}
}

func getAlarmRule(eventType, eventGroup, eventEntity, eventKey string) []map[string]interface{} {
	var db = database.DB
	var dataList []map[string]interface{}
	if eventEntity != "" {
		// sql = fmt.Sprintf("select id,title,alarm_rule,alarm_value,alarm_sleep,alarm_times,level_id,channel_id from alarm_rules "+
		// 	"where enable=1 and event_type='%s' and event_key='%s' and event_entity='%s'  order by level_id asc", eventType, eventKey, eventEntity)
		// res, _ := database.DB.Model(model.AlarmRule).Find()
		db.Model(&model.AlarmRule{}).Where("enable=1").Where("event_type=?", eventType).Where("event_key=?", eventKey).Where("event_entity=?", eventEntity).Order("level_id asc").Find(&dataList)
		if len(dataList) > 0 {
			return dataList
		}
	}
	if eventGroup != "" {
		db.Model(&model.AlarmRule{}).Where("enable=1").Where("event_type=?", eventType).Where("event_key=?", eventKey).Where("event_group=?", eventGroup).Order("level_id asc").Find(&dataList)
		if len(dataList) > 0 {
			return dataList
		}
	}

	db.Model(&model.AlarmRule{}).Where("enable=1").Where("event_type=?", eventType).Where("event_key=?", eventKey).Order("level_id asc").Find(&dataList)
	return dataList
}

func matchAlarmRule(alarmRule string, alarmValue float64, eventValue float64) bool {
	//alarmValueFloat := conv.StrToFloat(alarmValue)
	//eventValueFloat := conv.StrToFloat(eventValue)
	log.Logger.Debug(fmt.Sprintf("matchAlarmRule, alarmRule:%s,alarmValue:%f,eventValue:%f", alarmRule, alarmValue, eventValue))
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

func sendAlarm(event, rule map[string]interface{}, match bool) {
	//fmt.Println(event)
	eventUuid := event["event_uuid"].(string)
	eventTime := event["event_time"].(string)
	eventType := event["event_type"].(string)
	eventGroup := event["event_group"].(string)
	eventKey := event["event_key"].(string)
	eventEntity := event["event_entity"].(string)
	eventValue := utils.StrToFloat64(event["event_value"].(string))
	eventTag := event["event_tag"].(string)
	eventUnit := event["event_unit"].(string)

	ruleId := rule["id"].(int64)
	alarmTitle := rule["title"].(string)
	alarmRule := rule["alarm_rule"].(string)
	alarmValue := rule["alarm_value"].(string)
	alarmSleep := rule["alarm_sleep"].(int)
	alarmTimes := rule["alarm_times"].(int)
	levelId := rule["level_id"].(int)
	channelId := rule["channel_id"].(int)

	keyName := fmt.Sprintf("%s:%s:%s:%s", eventType, eventKey, eventTag, eventEntity)
	alarmCountKeyName := "alarm_count." + keyName
	alarmAtKeyName := "alarm_at." + keyName
	//fmt.Println(alarmAtKeyName)
	//fmt.Println(match)
	if match {
		alarmCount, _ := database.RDS.Get(alarmCountKeyName).Result()
		alarmAt, _ := database.RDS.Get(alarmAtKeyName).Result()
		if alarmCount == "" {
			alarmCount = "0"
			if alarmAt == "" {
				database.RDS.Set(alarmAtKeyName, time.Now().Unix(), time.Hour*time.Duration(72))
				log.Logger.Info(fmt.Sprintf("Set alarm at key %s", alarmAtKeyName))
			}
		}
		alarmCountInt := conv.StrToInt(alarmCount)
		if alarmCountInt < alarmTimes {
			var alarmLevel = ""
			var db = database.DB
			var record model.AlarmLevel
			db.Model(model.AlarmLevel{}).Select("level_name").Where("enable=1").Where("id=?", levelId).Take(&record)
			alarmLevel = record.LevelName

			var (
				sendMail = 0
			)
			//getChannelSql := fmt.Sprintf("select name,mail_enable,sms_enable,wechat_enable,phone_enable,webhook_enable,mail_list,sms_list,wechat_list,phone_list,webhook_url from alarm_channels where enable=1 and id=%d ", channelId)
			//channelList, _ := mysql.QueryAll(db, getChannelSql)

			var alarmChannel model.AlarmChannel
			db.Model(model.AlarmChannel{}).Where("enable=1").Where("id=?", channelId).Take(&alarmChannel)

			database.RDS.Incr(alarmCountKeyName)
			database.RDS.Expire(alarmCountKeyName, time.Second*time.Duration(alarmSleep))
			log.Logger.Info(fmt.Sprintf("Set alarm count key %s", alarmCountKeyName))

			mailEnable := alarmChannel.MailEnable
			mailList := alarmChannel.MailList

			if mailEnable == 1 && mailList != "" {
				log.Logger.Info(fmt.Sprintf("Start to send email to %s", mailList))
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
				/*/
				get event description
				*/
				var eventDescription model.EventDescription
				db.Model(model.EventDescription{}).Where("description is not null").Where("event_type=?", eventType).Where("event_key=?", eventKey).Take(&eventDescription)
				if eventDescription.Description != "" {
					data = append(data, "事件解释", eventDescription.Description)
				}
				eventContent := html.CreateTable(tableTitle, tableHeader, dataList)

				mailHello := fmt.Sprintf("尊敬的用户：<p></p>您好！您收到一条【%s】事件：【%s】，请您及时关注和处理。", alarmLevel, alarmTitle)
				mailContent := "<span style='margin-top:1px;'>" + mailHello + "</span><p></p>" + eventContent + "<p style='style=\"white-space: pre-wrap;\"'></p><div style='margin-top:20px'></div><div style='margin-top:30px; color:#666'><hr color='#ccc' style='border:1px dashed #cccccc;' />本邮件来自Lepus实时事件告警组件，请勿直接回复本邮件。如需获得技术支持，可联系我们：<a href='https://www.lepus.cc' target='_blank'>https://www.lepus.cc</a></div>"
				//fmt.Println(mailContent)
				//return
				mailTitle := fmt.Sprintf("[%s][%s]%s", alarmLevel, eventEntity, alarmTitle)
				var sendErrorInfo string
				if err := mail.Send(mailTo, mailTitle, mailContent); err != nil {
					sendErrorInfo = err.Error()
					sendMail = 2
					log.Logger.Error(fmt.Sprintf("Failed to send email %s,%s: %s", mailTitle, mailList, err))
				} else {
					sendErrorInfo = "OK"
					sendMail = 1
					log.Logger.Info(fmt.Sprintf("Success to send email %s,%s", mailTitle, mailList))
				}
				db.Create(&model.AlarmSendLog{SendType: "mail", Receiver: mailList, Content: mailContent, Status: sendMail, ErrorInfo: sendErrorInfo})
			}

			var createRecord model.AlarmEvent
			createRecord.AlarmTitle = alarmTitle
			createRecord.AlarmLevel = alarmLevel
			createRecord.AlarmRule = alarmRule
			createRecord.AlarmValue = alarmValue
			createRecord.EventUnit = eventUuid
			formatTime, _ := time.Parse("2006-01-02 15:04:05", eventTime)
			createRecord.EventTime = formatTime
			createRecord.EventType = eventType
			createRecord.EventGroup = eventGroup
			createRecord.EventEntity = eventEntity
			createRecord.EventKey = eventKey
			createRecord.EventValue = eventValue
			createRecord.EventUnit = eventUnit
			createRecord.EventTag = eventTag
			createRecord.RuleId = ruleId
			createRecord.LevelId = levelId
			createRecord.ChannelId = channelId
			createRecord.SendMail = sendMail

			result := db.Create(&createRecord)
			if result.Error != nil {
				log.Logger.Error(fmt.Sprintln("Failed insert alarm event data to database, ", err))
				return
			}

		}
	} else {
		//没有匹配到告警，如果有key信息，说明之前有告警，记录恢复信息，并删除key重新计算告警限流次数
		alarmCount, _ := database.RDS.Get(alarmCountKeyName).Result()
		alarmAt, _ := database.RDS.Get(alarmAtKeyName).Result()
		if alarmCount != "" && alarmAt != "" {
			database.RDS.Del(alarmCountKeyName)
			database.RDS.Del(alarmAtKeyName)
		}

		return
	}
}
