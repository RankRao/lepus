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
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"lepus/src/libary/conf"
	"lepus/src/libary/conv"
	"lepus/src/libary/http"
	"lepus/src/libary/logger"
	"lepus/src/libary/tool"
	"lepus/src/libary/utils"
	"os/exec"

	network "net"
	"time"
)

var (
	eventType  = "OS"
	eventGroup = conf.Option["event_group"]
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_os_mon.log", conv.StrToInt(conf.Option["debug"]))

func GetLocalIP() (ip string, err error) {
	addrs, err := network.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*network.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return "", nil
}

func doSysCommand(command string) string {
	var (
		dfOut string
	)
	cmd := exec.Command("/bin/sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		fmt.Printf("execute command %s start with err:%s", "df -h", err)
		log.Error(fmt.Sprintf("execute command %s with err:%s", "df -h", err))
	} else {
		//if err = cmd.Wait();err != nil{
		//	fmt.Printf("execute command %s wait with err:%s", "df -h",err)
		//}
		cmd.Wait()
		dfOut = out.String()
	}
	return dfOut
}
func osInfoCollector() {
	eventTime := tool.GetNowTime()
	ip, _ := GetLocalIP()
	hostInfo, _ := host.Info()
	hostname := hostInfo.Hostname
	eventEntity := hostname
	log.Info(fmt.Sprintf("Start collector os data at ip:%s, hostname:%s ", ip, hostname))
	//os := hostInfo.OS
	//timestamp,_ := host.BootTime()
	//uptime := time.Unix(int64(timestamp),0)
	//bootTime := uptime.Local().Format("2006-01-02 15:04:05")
	//platform,_,version,_ := host.PlatformInformation()
	loadAvg, _ := load.Avg()
	load := loadAvg.Load1
	//cpuPhysicalNum := cpu.Counts(false)
	//cpuLogicalNum := cpu.Counts(true)
	cpuPercentAll, _ := cpu.Percent(1*time.Second, false)
	cpuPercent := cpuPercentAll[0]

	memInfo, _ := mem.VirtualMemory()
	memTotal := memInfo.Total
	memAvailable := memInfo.Available
	memUsed := memInfo.Used
	memFree := memInfo.Free
	memUsedPercent := memInfo.UsedPercent
	swapTotal := memInfo.SwapTotal
	swapFree := memInfo.SwapFree
	swapCached := memInfo.SwapCached

	cpuOut := doSysCommand("ps aux|head -1;ps aux|grep -v PID|sort -rn -k +3|head")
	cpuEventDetail := []map[string]interface{}{{"Cpu Usage Top": cpuOut}}
	memOut := doSysCommand("ps aux|head -1;ps aux|grep -v PID|sort -rn -k +4|head")
	memEventDetail := []map[string]interface{}{{"Memory Usage Top": memOut}}
	itemList := []map[string]interface{}{
		{"event_key": "cpu.load", "event_value": load, "event_tag": "", "event_unit": "", "event_detail": cpuEventDetail},
		{"event_key": "cpu.percent", "event_value": cpuPercent, "event_tag": "", "event_unit": "%", "event_detail": cpuEventDetail},
		{"event_key": "memory.total", "event_value": memTotal / 1024 / 1024, "event_tag": "", "event_unit": "MB", "event_detail": ""},
		{"event_key": "memory.used", "event_value": memUsed / 1024 / 1024, "event_tag": "", "event_unit": "MB", "event_detail": memEventDetail},
		{"event_key": "memory.free", "event_value": memFree / 1024 / 1024, "event_tag": "", "event_unit": "MB", "event_detail": ""},
		{"event_key": "memory.available", "event_value": memAvailable / 1024 / 1024, "event_tag": "", "event_unit": "MB", "event_detail": ""},
		{"event_key": "memory.usedPercent", "event_value": memUsedPercent, "event_tag": "", "event_unit": "%", "event_detail": memEventDetail},
		{"event_key": "memory.swapTotal", "event_value": swapTotal / 1024 / 1024, "event_tag": "", "event_unit": "MB", "event_detail": ""},
		{"event_key": "memory.swapFree", "event_value": swapFree / 1024 / 1024, "event_tag": "", "event_unit": "MB", "event_detail": ""},
		{"event_key": "memory.swapCached", "event_value": swapCached / 1024 / 1024, "event_tag": "", "event_unit": "MB", "event_detail": ""},
	}

	parts, err := disk.Partitions(true)
	if err != nil {
		log.Error(fmt.Sprintf("Get Disk Partition failed, err:%v\n", err))
	}
	diskOut := doSysCommand("sudo df -h")
	diskEventDetail := []map[string]interface{}{{"Disk Usage Detail": diskOut}}
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		itemList = append(itemList, map[string]interface{}{"event_key": "disk.used", "event_value": diskInfo.Used / 1024 / 1024, "event_tag": part.Mountpoint, "event_unit": "MB", "event_detail": diskEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "disk.free", "event_value": diskInfo.Free / 1024 / 1024, "event_tag": part.Mountpoint, "event_unit": "MB", "event_detail": diskEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "disk.usedPercent", "event_value": diskInfo.UsedPercent, "event_tag": part.Mountpoint, "event_unit": "%", "event_detail": diskEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "disk.inodesUsed", "event_value": diskInfo.InodesUsed / 1000, "event_tag": part.Mountpoint, "event_unit": "K", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "disk.inodesFree", "event_value": diskInfo.InodesFree / 1000, "event_tag": part.Mountpoint, "event_unit": "K", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "disk.inodesUsedPercent", "event_value": diskInfo.InodesUsedPercent, "event_tag": part.Mountpoint, "event_unit": "%", "event_detail": ""})
	}

	diskIoStat, _ := disk.IOCounters()
	ioOut := doSysCommand("sudo iotop -o -d 0.5 -botq -n 2")
	ioEventDetail := []map[string]interface{}{{"Disk IO Detail": ioOut}}
	for _, v := range diskIoStat {
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.readCount", "event_value": v.ReadCount, "event_tag": v.Name, "event_unit": "", "event_detail": ioEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.writeCount", "event_value": v.WriteCount, "event_tag": v.Name, "event_unit": "", "event_detail": ioEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.readBytes", "event_value": v.ReadBytes, "event_tag": v.Name, "event_unit": "Byte", "event_detail": ioEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.writeBytes", "event_value": v.WriteBytes, "event_tag": v.Name, "event_unit": "Byte", "event_detail": ioEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.mergedReadCount", "event_value": v.MergedReadCount, "event_tag": v.Name, "event_unit": "", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.mergedWriteCount", "event_value": v.MergedWriteCount, "event_tag": v.Name, "event_unit": "", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.readTime", "event_value": v.ReadTime, "event_tag": v.Name, "event_unit": "Byte", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.writeTime", "event_value": v.WriteTime, "event_tag": v.Name, "event_unit": "Byte", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.iopsInProgress", "event_value": v.IopsInProgress, "event_tag": v.Name, "event_unit": "Byte", "event_detail": ""})
	}

	netInfo, _ := net.IOCounters(true)
	netOut := doSysCommand("sudo iftop -NntP -s 2")
	netEventDetail := []map[string]interface{}{{"Network Traffic Detail": netOut}}
	for _, v := range netInfo {
		itemList = append(itemList, map[string]interface{}{"event_key": "network.bytesSent", "event_value": v.BytesSent / 1024 / 8, "event_tag": v.Name, "event_unit": "KB", "event_detail": netEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.bytesRecv", "event_value": v.BytesRecv / 1024 / 8, "event_tag": v.Name, "event_unit": "KB", "event_detail": netEventDetail})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.packetsSent", "event_value": v.PacketsSent / 1000, "event_tag": v.Name, "event_unit": "K", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.packetsRecv", "event_value": v.PacketsRecv / 1000, "event_tag": v.Name, "event_unit": "K", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.errin", "event_value": v.Errin, "event_tag": v.Name, "event_unit": "", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.errout", "event_value": v.Errout, "event_tag": v.Name, "event_unit": "", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.dropin", "event_value": v.Dropin, "event_tag": v.Name, "event_unit": "", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.dropout", "event_value": v.Dropout, "event_tag": v.Name, "event_unit": "", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.fifoin", "event_value": v.Fifoin, "event_tag": v.Name, "event_unit": "", "event_detail": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "network.fifoout", "event_value": v.Fifoout, "event_tag": v.Name, "event_unit": "", "event_detail": ""})
	}

	events := make([]map[string]interface{}, 0)
	for _, item := range itemList {
		event := map[string]interface{}{
			"event_time":   eventTime,
			"event_type":   eventType,
			"event_group":  eventGroup,
			"event_entity": eventEntity,
			"event_key":    item["event_key"],
			"event_value":  item["event_value"],
			"event_tag":    item["event_tag"],
			"event_unit":   item["event_unit"],
			"event_detail": item["event_detail"],
		}
		events = append(events, event)
	}
	//fmt.Println(events)
	_, err = http.Post(conf.Option["proxy"], events)
	if err != nil {
		log.Error(fmt.Sprintln("Send events to proxy error:", err))
	}

}

func main() {
	for true {
		go osInfoCollector()
		time.Sleep(time.Duration(utils.StrToInt(conf.Option["interval"])) * time.Second)
	}
}
