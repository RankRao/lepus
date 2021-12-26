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
	"fmt"
	"github.com/gorhill/cronexpr"
	"lepus/src/libary/conf"
	"lepus/src/libary/conv"
	"lepus/src/libary/logger"
	"lepus/src/libary/mysql"
	"lepus/src/libary/utils"
	"os/exec"
	"runtime"
	"time"
)

var opt = conf.Option
var db = mysql.InitConnect()
var log = logger.NewLog(opt["log_dir"]+"/lepus_task.log", conv.StrToInt(conf.Option["debug"]))

var (
	taskMap = make(map[string][]string)
	memMap  = make(map[string][]string)
)

func cronTabRunNextTime(cronFormat string) string {
	expr, err := cronexpr.Parse(cronFormat)
	if err != nil {
		log.Error(fmt.Sprintln("Parse cronTab err,", err))
	}
	nextRunTime := expr.Next(time.Now()).Format("2006-01-02 15:04:05")
	return nextRunTime

}

func periodRunNextTime(periodTime string) string {
	var periodTimeSecond int64
	periodTimeSecond = utils.StrToInt64(periodTime)
	nextRunTime := time.Now().Unix() + periodTimeSecond
	return time.Unix(nextRunTime, 0).Format("2006-01-02 15:04:05")
}

func addTask() {
	sql := fmt.Sprintf("select id,type_id,task_command,schedule_type,crontab_time,period_time,status from task where (enable=1 and schedule_type in ('crontab','period') and status in ('waiting','success','failed')  and next_time<='%s') or (enable=1 and schedule_type='manual' and status='waiting') ", utils.GetCurrentTime())
	rows, err := mysql.QueryAll(db, sql)
	if err != nil {
		log.Error(fmt.Sprintln("Query task info err,", err))
	}
	for _, row := range rows {
		var nextRunTime string
		switch row["schedule_type"] {
		case "crontab":
			nextRunTime = cronTabRunNextTime(row["crontab_time"].(string))
		case "period":
			nextRunTime = periodRunNextTime(row["period_time"].(string))
		default:
			nextRunTime = utils.GetCurrentTime()
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
		if err = mysql.Execute(db, fmt.Sprintf("update task set status='running',next_time='%s' where id='%s' ", nextRunTime, row["id"].(string))); err != nil {
			log.Error(fmt.Sprintln("Update task status to running err,", err))
		}
		if err = mysql.Execute(db, fmt.Sprintf("insert into task_run(task_id,run_start_time,run_status) values('%s','%s','%s') ", row["id"].(string), utils.GetCurrentTime(), "waiting")); err != nil {
			log.Error(fmt.Sprintln("Insert task to task run err,", err))
		}
	}
}

func getTask() {
	taskMap = make(map[string][]string)
	sql := "select task_run.id as run_id,task_id,task_name,task_command,timeout from task_run join task on task_run.task_id=task.id and task_run.run_status='waiting' and task.status='running' where now()>run_start_time"
	rows, err := mysql.QueryAll(db, sql)
	if err != nil {
		log.Error(fmt.Sprintln("Get task err,", err))
		return
	}
	if rows == nil || len(rows) == 0 {
		return
	}
	for _, row := range rows {
		taskMap[row["run_id"].(string)] = []string{row["task_id"].(string), row["task_name"].(string), row["task_command"].(string), row["timeout"].(string)}
	}

}

func startTask() {
	if taskMap == nil || len(taskMap) == 0 {
		return
	}
	for k, v := range taskMap {
		startOneSubTask(k, v)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func startOneSubTask(taskKey string, taskInfo []string) {
	_, in := memMap[taskKey]
	if in && memMap[taskKey][0] == taskInfo[0] {
		return
	}
	if in {
		delete(memMap, taskKey)
	}
	runTaskCmd(taskKey, taskInfo[0], taskInfo[1], taskInfo[2], taskInfo[3])
	memMap[taskKey] = taskInfo
}

func runTaskCmd(runId, taskId, taskName, taskCommand, timeout string) {
	log.Info(fmt.Sprintf("Task %s execute command start, taskId:%s, runId:%s, taskCommand:%s", taskName, taskId, runId, taskCommand))
	if err := mysql.Execute(db, fmt.Sprintf("update task_run set run_status='running' where id='%s' ", runId)); err != nil {
		log.Error(fmt.Sprintf("Task %s set status to running err:%s", taskName, err))
	}
	var (
		status, logContent string
	)

	cmd := exec.Command("/bin/sh", "-c", taskCommand)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		status = "failed"
		logContent = fmt.Sprint(err)
		log.Error(fmt.Sprintf("Task %s execute command start with err:%s", taskName, err))
	} else {
		mysql.Execute(db, fmt.Sprintf("update task_run set run_pid='%d' where id='%s' ", cmd.Process.Pid, runId))
		doneChan := make(chan bool, 1)
		errorChan := make(chan error, 1)
		go func() {
			if err = cmd.Wait(); err != nil {
				errorChan <- err
			} else {
				doneChan <- true
			}
		}()

		timeoutSecond := utils.StrToInt(timeout)
		if timeoutSecond == 0 {
			timeoutSecond = 86400 * 365
		}
		select {
		case <-time.After(time.Duration(timeoutSecond) * time.Second):
			cmd.Process.Kill()
			status = "failed"
			logContent = "exec command timeout and process killed "
			log.Error(fmt.Sprintf("Task %s execute command timeout and process exit.", taskName))
		case <-doneChan:
			status = "success"
			logContent = out.String()
			log.Info(fmt.Sprintf("Task %s execute command finished and success", taskName))
		case err := <-errorChan:
			status = "failed"
			logContent = fmt.Sprint(err)
			log.Error(fmt.Sprintf("Task %s execute command finished with err:%s", taskName, err))
		}
	}
	mysql.Execute(db, fmt.Sprintf("update task_run set run_status='%s',run_end_time='%s' where id='%s' ", status, utils.GetCurrentTime(), runId))
	mysql.Execute(db, fmt.Sprintf("update task set status='%s' where id='%s' ", status, taskId))
	mysql.Execute(db, fmt.Sprintf("insert into task_log(task_id,run_id,content) values('%s','%s','%s')", taskId, runId, logContent))
}

func resetAllDeadTask() {
	mysql.Execute(db, fmt.Sprintf("update task_run set run_status='failed',run_end_time='%s' where run_status='running' ", utils.GetCurrentTime()))
	mysql.Execute(db, fmt.Sprintf("update task set status='failed' where status='running'"))
}

func main() {
	start := time.Now()
	fmt.Printf("Task schedule server start at %s", start)
	log.Info(fmt.Sprintf("Task schedule server start at %s", start))
	runtime.GOMAXPROCS(runtime.NumCPU())
	resetAllDeadTask()
	for true {
		addTask()
		getTask()
		startTask()
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
