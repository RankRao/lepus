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

package model

import (
	"time"
)

type Users struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"size:30;uniqueIndex" json:"username"`
	Password    string    `gorm:"size:200" json:"password"`
	ChineseName string    `gorm:"size:50" json:"chineseName"`
	Admin       bool      `gorm:"default:false" json:"admin"`
	Remark      string    `gorm:"size:200" json:"remark"`
	CreatedAt   time.Time `json:"createAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Token struct {
	TokenKey  string    `gorm:"primaryKey;size:180"`
	Value     []byte    `gorm:"type:bytes;size:1000"`
	CreatedAt time.Time `json:"createAt"`
	Expired   time.Time `json:"expired"`
}

type DatasourceType struct {
	Id          int       `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"size:30;uniqueIndex" json:"name"`
	Description string    `gorm:"size:100" json:"description"`
	Sort        int8      `gorm:"default:1" json:"sort"`
	Enable      int8      `gorm:"default:1" json:"enable"`
	CreatedAt   time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt   time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (DatasourceType) TableName() string {
	return "datasource_type"
}

type Datasource struct {
	Id            int       `gorm:"primarykey" json:"id"`
	Name          string    `gorm:"size:50;uniqueIndex" json:"name"`
	GroupName     string    `gorm:"size:50" json:"group_name"`
	Idc           string    `gorm:"size:30" json:"idc"`
	Env           string    `gorm:"size:30" json:"env"`
	Type          string    `gorm:"size:30" json:"type"`
	Host          string    `gorm:"size:100;index:uniq_host_port_dbid,unique" json:"host"`
	Port          string    `gorm:"size:30;index:uniq_host_port_dbid,unique" json:"port"`
	User          string    `gorm:"size:30" json:"user"`
	Pass          string    `gorm:"size:100" json:"pass"`
	Dbid          string    `gorm:"size:50;index:uniq_host_port_dbid,unique" json:"dbid"`
	Role          int32     `gorm:"default:1" json:"role"`
	Enable        int32     `gorm:"default:1" json:"enable"`
	Status        int32     `gorm:"default:1" json:"status"`
	StatusText    string    `gorm:"size:500" json:"status_text"`
	DbmetaEnable  int32     `gorm:"default:0" json:"dbmeta_enable"`
	ExecuteEnable int32     `gorm:"default:0" json:"execute_enable"`
	MonitorEnable int32     `gorm:"default:0" json:"monitor_enable"`
	AlarmEnable   int32     `gorm:"default:0" json:"alarm_enable"`
	CreatedAt     time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt     time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (Datasource) TableName() string {
	return "datasource"
}

type Favorite struct {
	ID             int       `gorm:"primarykey" json:"id"`
	Username       string    `gorm:"size:50;index:idx_user_datasource" json:"username"`
	DatasourceType string    `gorm:"size:50;index:idx_user_datasource" json:"datasource_type"`
	Datasource     string    `gorm:"size:100;index:idx_user_datasource" json:"datasource"`
	DatabaseName   string    `gorm:"size:50" json:"database_name"`
	Content        string    `gorm:"size:1000" json:"content"`
	CreatedAt      time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt      time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (Favorite) TableName() string {
	return "favorite"
}

type MetaDatabase struct {
	Id             int64     `gorm:"primarykey" json:"id"`
	DatasourceType string    `gorm:"size:50;index" json:"datasource_type"`
	Host           string    `gorm:"size:100;index:idx_host_port" json:"host"`
	Port           string    `gorm:"size:10;index:idx_host_port" json:"port"`
	DatabaseName   string    `gorm:"size:50;index" json:"database_name"`
	SchemaName     string    `gorm:"size:50" json:"schema_name"`
	Characters     string    `gorm:"size:50" json:"characters"`
	IsDeleted      int       `gorm:"default:0" json:"is_deleted"`
	CreatedAt      time.Time `gorm:"column:gmt_created;index" json:"gmt_created"`
	UpdatedAt      time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (MetaDatabase) TableName() string {
	return "meta_database"
}

type MetaTable struct {
	Id             int       `gorm:"primarykey" json:"id"`
	DatasourceType string    `gorm:"size:50;index" json:"datasource_type"`
	Host           string    `gorm:"size:100;index:idx_host_port" json:"host"`
	Port           string    `gorm:"size:10;index:idx_host_port" json:"port"`
	DatabaseName   string    `gorm:"size:50;index" json:"database_name"`
	TableType      string    `gorm:"size:50" json:"table_type"`
	TableNameX     string    `gorm:"column:table_name;size:50;index" json:"table_name"`
	TableComment   string    `gorm:"size:50" json:"table_comment"`
	Characters     string    `gorm:"size:50" json:"characters"`
	IsDeleted      int8      `gorm:"default:0" json:"is_deleted"`
	CreatedAt      time.Time `gorm:"column:gmt_created;index" json:"gmt_created"`
	UpdatedAt      time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (MetaTable) TableName() string {
	return "meta_table"
}

type MetaColumn struct {
	Id               int       `gorm:"primarykey" json:"id"`
	DatasourceType   string    `gorm:"size:50;index" json:"datasource_type"`
	Host             string    `gorm:"size:100;index:idx_host_port" json:"host"`
	Port             string    `gorm:"size:10;index:idx_host_port" json:"port"`
	DatabaseName     string    `gorm:"size:50;index" json:"database_name"`
	TableNameX       string    `gorm:"column:table_name;size:50;index" json:"table_name"`
	ColumnName       string    `gorm:"size:50;index" json:"column_name"`
	ColumnComment    string    `gorm:"size:50" json:"column_comment"`
	DataType         string    `gorm:"size:50" json:"data_type"`
	IsNullable       string    `gorm:"size:50" json:"is_nullable"`
	DefaultValue     string    `gorm:"size:50" json:"default_value"`
	Ordinal_Position int       `gorm:"default:0" json:"ordinal_position"`
	Characters       string    `gorm:"size:100" json:"characters"`
	IsDeleted        int8      `gorm:"default:0" json:"is_deleted"`
	CreatedAt        time.Time `gorm:"column:gmt_created;index" json:"gmt_created"`
	UpdatedAt        time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (MetaColumn) TableName() string {
	return "meta_column"
}

type Idc struct {
	Id          int64     `gorm:"primarykey" json:"id"`
	IdcKey      string    `gorm:"size:30;index" json:"idc_key"`
	IdcName     string    `gorm:"size:30" json:"idc_name"`
	City        string    `gorm:"size:30" json:"city"`
	Description string    `gorm:"size:300" json:"description"`
	CreatedAt   time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt   time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (Idc) TableName() string {
	return "idc"
}

type Env struct {
	Id          int64     `gorm:"primarykey" json:"id"`
	EnvKey      string    `gorm:"size:30;index" json:"env_key"`
	EnvName     string    `gorm:"size:30" json:"env_name"`
	Description string    `gorm:"size:300" json:"description"`
	CreatedAt   time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt   time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (Env) TableName() string {
	return "env"
}

type TaskOption struct {
	TaskKey         string    `gorm:"size:50;primarykey" json:"task_key"`
	TaskName        string    `gorm:"size:50," json:"task_name"`
	TaskDescription string    `gorm:"size:500," json:"task_description"`
	Crontab         string    `gorm:"size:100," json:"crontab"`
	Enable          int8      `gorm:"default:1" json:"enable"`
	CreatedAt       time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt       time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (TaskOption) TableName() string {
	return "task_option"
}

type TaskHeartbeat struct {
	HeartbeatKey     string    `gorm:"size:50;primarykey" json:"heartbeat_key"`
	HeartbeatTime    time.Time `gorm:"column:heartbeat_time" json:"heartbeat_time"`
	HeartbeatEndTime time.Time `gorm:"column:heartbeat_end_time" json:"heartbeat_end_time"`
	CreatedAt        time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt        time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
	//TaskOption       TaskOption `gorm:"references:task"`
}

func (TaskHeartbeat) TableName() string {
	return "task_heartbeat"
}

type Event struct {
	EventUuid   string    `gorm:"primary_key" json:"event_uuid"`
	EventTime   time.Time `gorm:"column:event_time" json:"event_time"`
	EventType   string    `gorm:"" json:"event_type"`
	EventGroup  string    `gorm:"" json:"event_group"`
	EventEntity string    `gorm:"" json:"event_entity"`
	EventKey    string    `gorm:"" json:"event_key"`
	EventValue  float32   `gorm:"type:decimal(20,2)" json:"event_value"`
	EventTag    string    `gorm:"" json:"event_tag"`
	EventUnit   string    `gorm:"" json:"event_unit"`
	EventDetail string    `gorm:"" json:"event_detail"`
	//CreatedAt   time.Time `gorm:"column:gmt_created;autoCreateTime" json:"gmt_created"`
}

type EventsDescription struct {
	ID          int64  `gorm:"primarykey" json:"id"`
	EventType   string `gorm:"column:event_type" json:"eventType"`
	EventKey    string `gorm:"column:event_key" json:"eventKey"`
	Description string `gorm:"size:300" json:"description"`
}

type EventDescription struct {
	ID          int64  `gorm:"primarykey" json:"id"`
	EventType   string `gorm:"column:event_type" json:"eventType"`
	EventKey    string `gorm:"column:event_key" json:"eventKey"`
	Description string `gorm:"size:300" json:"description"`
}

func (EventDescription) TableName() string {
	return "event_description"
}

type EventGlobal struct {
	Id             int64     `gorm:"primarykey" json:"id"`
	DatasourceType string    `gorm:"size:50;index" json:"datasource_type"`
	DatasourceName string    `gorm:"size:50;index" json:"datasource_name"`
	Host           string    `gorm:"size:100;index:idx_host_port" json:"host"`
	Port           string    `gorm:"size:50;index:idx_host_port" json:"port"`
	Version        string    `gorm:"size:50" json:"version"`
	Role           string    `gorm:"size:50" json:"role"`
	Uptime         int64     `gorm:"default:-1" json:"uptime"`
	Connect        int       `gorm:"default:-1" json:"connect"`
	Session        int       `gorm:"default:-1" json:"session"`
	Active         int       `gorm:"default:-1" json:"active"`
	Wait           int       `gorm:"default:-1" json:"wait"`
	Qps            int       `gorm:"default:-1" json:"qps"`
	Tps            int       `gorm:"default:-1" json:"tps"`
	Repl           int       `gorm:"default:-1" json:"repl"`
	Delay          int       `gorm:"default:-1" json:"delay"`
	Remark         string    `gorm:"size:1000" json:"remark"`
	CreatedAt      time.Time `gorm:"column:gmt_created;index" json:"gmt_created"`
	UpdatedAt      time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (EventGlobal) TableName() string {
	return "event_global"
}

type AlarmChannel struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:30" json:"name"`
	Description string    `gorm:"size:1000" json:"description"`
	Enable      int       `gorm:"default:0" json:"enable"`
	MailEnable  int       `gorm:"default:0" json:"mail_enable"`
	MailList    string    `gorm:"size:500" json:"mail_list"`
	CreatedAt   time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt   time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (AlarmChannel) TableName() string {
	return "alarm_channel"
}

type AlarmLevel struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	LevelName   string    `gorm:"size:30" json:"level_name"`
	Description string    `gorm:"size:1000" json:"description"`
	Enable      int       `gorm:"default:0" json:"enable"`
	CreatedAt   time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt   time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (AlarmLevel) TableName() string {
	return "alarm_level"
}

type AlarmRule struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:30" json:"title"`
	EventType   string    `gorm:"size:50" json:"event_type"`
	EventGroup  string    `gorm:"size:50" json:"event_group"`
	EventKey    string    `gorm:"size:50" json:"event_key"`
	EventEntity string    `gorm:"size:50" json:"event_entity"`
	AlarmRule   string    `gorm:"size:50" json:"alarm_rule"`
	AlarmValue  string    `gorm:"size:50" json:"alarm_value"`
	AlarmSleep  int       `gorm:"default:3600" json:"alarm_sleep"`
	AlarmTimes  int       `gorm:"default:3" json:"alarm_times"`
	ChannelId   int       `gorm:"default:1" json:"channel_id"`
	LevelId     int       `gorm:"default:1" json:"level_id"`
	Enable      int       `gorm:"default:1" json:"enable"`
	CreatedAt   time.Time `gorm:"column:gmt_created" json:"gmt_created"`
	UpdatedAt   time.Time `gorm:"column:gmt_updated" json:"gmt_updated"`
}

func (AlarmRule) TableName() string {
	return "alarm_rule"
}

type AlarmEvent struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	AlarmTitle  string    `gorm:"size:50" json:"alarm_title"`
	AlarmLevel  string    `gorm:"size:50" json:"alarm_level"`
	AlarmRule   string    `gorm:"size:50" json:"alarm_rule"`
	AlarmValue  string    `gorm:"size:50" json:"alarm_value"`
	EventTime   time.Time `gorm:"column:event_time" json:"event_time"`
	EventUuid   string    `gorm:"size:200" json:"event_uuid"`
	EventType   string    `gorm:"size:50" json:"event_type"`
	EventGroup  string    `gorm:"size:50" json:"event_group"`
	EventKey    string    `gorm:"size:50" json:"event_key"`
	EventValue  float64   `gorm:"type:decimal(20,2)" json:"event_value"`
	EventUnit   string    `gorm:"size:50" json:"event_unit"`
	EventEntity string    `gorm:"size:50" json:"event_entity"`
	EventTag    string    `gorm:"size:50" json:"event_tag"`
	RuleId      int64     `gorm:"default:0" json:"rule_id"`
	LevelId     int       `gorm:"default:0" json:"level_id"`
	ChannelId   int       `gorm:"default:0" json:"channel_id"`
	SendMail    int       `gorm:"default:0" json:"send_mail"`
	Status      int       `gorm:"default:0" json:"status"`
	CreatedAt   time.Time `gorm:"column:gmt_created" json:"gmt_created"`
}

func (AlarmEvent) TableName() string {
	return "alarm_event"
}

type AlarmSuggest struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	EventType string    `gorm:"size:50" json:"event_type"`
	EventKey  string    `gorm:"size:50" json:"event_key"`
	Content   string    `gorm:"size:3000" json:"content"`
	CreatedAt time.Time `gorm:"column:gmt_created" json:"gmt_created"`
}

func (AlarmSuggest) TableName() string {
	return "alarm_suggest"
}

type AlarmSendLog struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	SendType  string    `gorm:"size:50" json:"send_type"`
	Receiver  string    `gorm:"size:300" json:"receiver"`
	Content   string    `gorm:"size:5000" json:"content"`
	Status    int       `gorm:"default:0" json:"status"`
	ErrorInfo string    `gorm:"size:500" json:"error_info"`
	CreatedAt time.Time `gorm:"column:gmt_created" json:"gmt_created"`
}

func (AlarmSendLog) TableName() string {
	return "alarm_send_log"
}

type AlarmTrack struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	AlarmId   int64     `gorm:"alarm_id"`
	UserId    int64     `gorm:"user_id"`
	Content   string    `gorm:"size:1000"`
	CreatedAt time.Time `gorm:"column:gmt_created" json:"gmt_created"`
}

func (AlarmTrack) TableName() string {
	return "alarm_track"
}
