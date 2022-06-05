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
	"fmt"
	"lepus/src/libary/conf"
	"lepus/src/libary/conv"
	"lepus/src/libary/http"
	"lepus/src/libary/logger"
	"lepus/src/libary/tool"
	"lepus/src/libary/utils"
	"math"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"

	network "net"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var (
	eventType   = "OS"
	eventGroup  = conf.Option["server_group"]
	eventEntity = conf.Option["server_ip"]
	eventTag    = conf.Option["server_tag"]
)

var log = logger.NewLog(conf.Option["log_dir"]+"/lepus_server_agent.log", conv.StrToInt(conf.Option["debug"]))

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
	sysType := runtime.GOOS
	if sysType != "linux" {
		return ""
	}
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

func arrayCalc(dataArray []int) (maxVal, minVal, sumVal int) {
	// 定义最大值、最小值、求和变量
	max := dataArray[0]
	min := dataArray[0]
	sum := 0
	for i := 0; i < len(dataArray); i++ {
		// 循环判断数组的元素是否小于自定义的最大值，如果是就把值赋值给max，作为当前最小值
		if dataArray[i] > max {
			max = dataArray[i]
		}
		// 循环判断数组的元素是否小于自定义的最小值，如果是就把值赋值给max，作为当前最小值
		if dataArray[i] < min {
			min = dataArray[i]
		}
		sum += dataArray[i]
	}
	return max, min, sum
}

func osInfoCollector() {
	eventTime := tool.GetNowTime()
	hostInfo, _ := host.Info()
	hostname := hostInfo.Hostname
	if eventGroup == "" {
		eventGroup = "Default"
	}
	ip, _ := GetLocalIP()
	if eventEntity == "" {
		eventEntity = ip
	}

	log.Info(fmt.Sprintf("Start collector os data at ip:%s, hostname:%s ", ip, hostname))
	os := hostInfo.OS
	timestamp, _ := host.BootTime()
	//uptime := int(timestamp)
	uptime := time.Now().Unix() - int64(timestamp)
	unixTime := time.Unix(int64(timestamp), 0)
	bootTime := unixTime.Local().Format("2006-01-02 15:04:05")
	platform, _, version, _ := host.PlatformInformation()
	loadAvg, _ := load.Avg()
	load := loadAvg.Load1
	cpuPhysicalNum, _ := cpu.Counts(false)
	cpuLogicalNum, _ := cpu.Counts(true)
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

	itemList := []map[string]interface{}{
		{"event_key": "cpu.load", "event_value": load, "event_tag": "", "event_unit": ""},
		{"event_key": "cpu.percent", "event_value": cpuPercent, "event_tag": "", "event_unit": "%"},
		{"event_key": "memory.total", "event_value": memTotal / 1024 / 1024, "event_tag": "", "event_unit": "MB"},
		{"event_key": "memory.used", "event_value": memUsed / 1024 / 1024, "event_tag": "", "event_unit": "MB"},
		{"event_key": "memory.free", "event_value": memFree / 1024 / 1024, "event_tag": "", "event_unit": "MB"},
		{"event_key": "memory.available", "event_value": memAvailable / 1024 / 1024, "event_tag": "", "event_unit": "MB"},
		{"event_key": "memory.usedPercent", "event_value": memUsedPercent, "event_tag": "", "event_unit": "%"},
		{"event_key": "memory.swapTotal", "event_value": swapTotal / 1024 / 1024, "event_tag": "", "event_unit": "MB"},
		{"event_key": "memory.swapFree", "event_value": swapFree / 1024 / 1024, "event_tag": "", "event_unit": "MB"},
		{"event_key": "memory.swapCached", "event_value": swapCached / 1024 / 1024, "event_tag": "", "event_unit": "MB"},
	}

	parts, err := disk.Partitions(true)
	if err != nil {
		log.Error(fmt.Sprintf("Get Disk Partition failed, err:%v\n", err))
	}

	diskUsedPctArray := make([]int, 0)  //用于计算磁盘最大值的数组
	inodeUsedPctArray := make([]int, 0) //用于计算inodes最大值的数组
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		if strings.Index(part.Mountpoint, "/dev") == -1 && strings.Index(part.Mountpoint, "/proc") == -1 && strings.Index(part.Mountpoint, "/run") == -1 && strings.Index(part.Mountpoint, "/sys") == -1 {
			itemList = append(itemList, map[string]interface{}{"event_key": "disk.used", "event_value": diskInfo.Used / 1024 / 1024 / 1024, "event_tag": part.Mountpoint, "event_unit": "GB"})
			itemList = append(itemList, map[string]interface{}{"event_key": "disk.free", "event_value": diskInfo.Free / 1024 / 1024 / 1024, "event_tag": part.Mountpoint, "event_unit": "GB"})
			itemList = append(itemList, map[string]interface{}{"event_key": "disk.usedPercent", "event_value": diskInfo.UsedPercent, "event_tag": part.Mountpoint, "event_unit": "%"})
			itemList = append(itemList, map[string]interface{}{"event_key": "disk.inodesUsed", "event_value": diskInfo.InodesUsed / 1000, "event_tag": part.Mountpoint, "event_unit": "K"})
			itemList = append(itemList, map[string]interface{}{"event_key": "disk.inodesFree", "event_value": diskInfo.InodesFree / 1000, "event_tag": part.Mountpoint, "event_unit": "K"})
			itemList = append(itemList, map[string]interface{}{"event_key": "disk.inodesUsedPercent", "event_value": diskInfo.InodesUsedPercent, "event_tag": part.Mountpoint, "event_unit": "%"})
			diskUsedPctArray = append(diskUsedPctArray, int(math.Floor(diskInfo.UsedPercent)))         //float抓int放到数组
			inodeUsedPctArray = append(inodeUsedPctArray, int(math.Floor(diskInfo.InodesUsedPercent))) //float抓int放到数组
		}
	}
	maxDiskUsedPct, _, _ := arrayCalc(diskUsedPctArray)
	maxInodeUsedPct, _, _ := arrayCalc(inodeUsedPctArray)

	diskIoStatPrev, _ := disk.IOCounters()
	time.Sleep(time.Duration(2 * time.Second))
	diskIoStat, _ := disk.IOCounters()

	diskIoReadCountArray := make([]int, 0)
	diskIoWriteCountArray := make([]int, 0)
	diskIoReadBytesArray := make([]int, 0)
	diskIoWriteBytesArray := make([]int, 0)
	diskIoReadTimeArray := make([]int, 0)
	diskIoWriteTimeArray := make([]int, 0)
	for i, v := range diskIoStat {
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.readCount", "event_value": (diskIoStat[i].ReadCount - diskIoStatPrev[i].ReadCount) / 2, "event_tag": v.Name, "event_unit": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.writeCount", "event_value": (diskIoStat[i].WriteCount - diskIoStatPrev[i].WriteCount) / 2, "event_tag": v.Name, "event_unit": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.readBytes", "event_value": (diskIoStat[i].ReadBytes - diskIoStatPrev[i].ReadBytes) / 1024 / 2, "event_tag": v.Name, "event_unit": "KB"})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.writeBytes", "event_value": (diskIoStat[i].WriteBytes - diskIoStatPrev[i].WriteBytes) / 1024 / 2, "event_tag": v.Name, "event_unit": "KB"})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.mergedReadCount", "event_value": (diskIoStat[i].MergedReadCount - diskIoStatPrev[i].MergedReadCount) / 2, "event_tag": v.Name, "event_unit": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.mergedWriteCount", "event_value": (diskIoStat[i].MergedWriteCount - diskIoStatPrev[i].MergedWriteCount) / 2, "event_tag": v.Name, "event_unit": ""})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.readTime", "event_value": (diskIoStat[i].ReadTime - diskIoStatPrev[i].ReadTime) / 1000 / 2, "event_tag": v.Name, "event_unit": "ms"})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.writeTime", "event_value": (diskIoStat[i].WriteTime - diskIoStatPrev[i].WriteTime) / 1000 / 2, "event_tag": v.Name, "event_unit": "ms"})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.iopsInProgress", "event_value": v.IopsInProgress, "event_tag": v.Name, "event_unit": "iops"})
		itemList = append(itemList, map[string]interface{}{"event_key": "diskio.ioTime", "event_value": (diskIoStat[i].IoTime - diskIoStatPrev[i].IoTime) / 1000 / 2, "event_tag": v.Name, "event_unit": "ms"})
		diskIoReadCountArray = append(diskIoReadCountArray, int(diskIoStat[i].ReadCount-diskIoStatPrev[i].ReadCount)/2)
		diskIoWriteCountArray = append(diskIoWriteCountArray, int(diskIoStat[i].WriteCount-diskIoStatPrev[i].WriteCount)/2)
		diskIoReadBytesArray = append(diskIoReadBytesArray, int(diskIoStat[i].ReadBytes-diskIoStatPrev[i].ReadBytes)/2)
		diskIoWriteBytesArray = append(diskIoWriteBytesArray, int(diskIoStat[i].WriteBytes-diskIoStatPrev[i].WriteBytes)/2)
		diskIoReadTimeArray = append(diskIoReadTimeArray, int(diskIoStat[i].ReadTime-diskIoStatPrev[i].ReadTime)/2)
		diskIoWriteTimeArray = append(diskIoWriteTimeArray, int(diskIoStat[i].WriteTime-diskIoStatPrev[i].WriteTime)/2)
	}
	_, _, diskIoReadCountSum := arrayCalc(diskIoReadCountArray)
	_, _, diskIoWriteCountSum := arrayCalc(diskIoWriteCountArray)
	_, _, diskIoReadBytesSum := arrayCalc(diskIoReadBytesArray)
	_, _, diskIoWriteBytesSum := arrayCalc(diskIoWriteBytesArray)
	_, _, diskIoReadTimeSum := arrayCalc(diskIoReadTimeArray)
	_, _, diskIoWriteTimeSum := arrayCalc(diskIoWriteTimeArray)

	netInfoPrev, _ := net.IOCounters(true)
	time.Sleep(time.Duration(2 * time.Second))
	netInfo, _ := net.IOCounters(true)

	netBytesSentArray := make([]int, 0)
	netBytesRecvArray := make([]int, 0)
	netPacketsSentArray := make([]int, 0)
	netPacketsRecvArray := make([]int, 0)
	for i, v := range netInfo {
		if strings.Index(v.Name, "lo") == -1 {
			itemList = append(itemList, map[string]interface{}{"event_key": "network.bytesSent", "event_value": (netInfo[i].BytesSent - netInfoPrev[i].BytesSent) / 1024 / 8 / 2, "event_tag": v.Name, "event_unit": "KB"})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.bytesRecv", "event_value": (netInfo[i].BytesRecv - netInfoPrev[i].BytesRecv) / 1024 / 8 / 2, "event_tag": v.Name, "event_unit": "KB"})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.packetsSent", "event_value": (netInfo[i].PacketsSent - netInfoPrev[i].PacketsSent) / 2, "event_tag": v.Name, "event_unit": ""})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.packetsRecv", "event_value": (netInfo[i].PacketsRecv - netInfoPrev[i].PacketsRecv) / 2, "event_tag": v.Name, "event_unit": ""})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.errin", "event_value": (netInfo[i].Errin - netInfoPrev[i].Errin) / 2, "event_tag": v.Name, "event_unit": ""})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.errout", "event_value": (netInfo[i].Errout - netInfoPrev[i].Errout) / 2, "event_tag": v.Name, "event_unit": ""})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.dropin", "event_value": (netInfo[i].Dropin - netInfoPrev[i].Dropin) / 2, "event_tag": v.Name, "event_unit": ""})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.dropout", "event_value": (netInfo[i].Dropout - netInfoPrev[i].Dropout) / 2, "event_tag": v.Name, "event_unit": ""})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.fifoin", "event_value": (netInfo[i].Fifoin - netInfoPrev[i].Fifoin) / 2, "event_tag": v.Name, "event_unit": ""})
			itemList = append(itemList, map[string]interface{}{"event_key": "network.fifoout", "event_value": (netInfo[i].Fifoout - netInfoPrev[i].Fifoout) / 2, "event_tag": v.Name, "event_unit": ""})
			netBytesSentArray = append(netBytesSentArray, int(netInfo[i].BytesSent-netInfoPrev[i].BytesSent)/2)
			netBytesRecvArray = append(netBytesRecvArray, int(netInfo[i].BytesRecv-netInfoPrev[i].BytesRecv)/2)
			netPacketsSentArray = append(netPacketsSentArray, int(netInfo[i].PacketsSent-netInfoPrev[i].PacketsSent)/2)
			netPacketsRecvArray = append(netPacketsRecvArray, int(netInfo[i].PacketsRecv-netInfoPrev[i].PacketsRecv)/2)
		}
	}
	_, _, netBytesSentSum := arrayCalc(netBytesSentArray)
	_, _, netBytesRecvSum := arrayCalc(netBytesRecvArray)
	_, _, netPacketsSentSum := arrayCalc(netPacketsSentArray)
	_, _, netPacketsRecvSum := arrayCalc(netPacketsRecvArray)

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

	/*
		split array map
	*/
	eventsSegment := utils.SplitArrayMap(events, 10)
	for _, events := range eventsSegment {
		sendEvent(events, "event")
	}

	sqlData := fmt.Sprintf("insert into dashboard_server(ip,hostname,os,platform,version,uptime,boot_time,cpu_physical_num,cpu_logical_num,tag,cpu_load,cpu_percent,memory_total,memory_used,memory_free,memory_available,memory_used_percent,swap_total,swap_free,swap_cached,disk_used_percent,inodes_used_percent,diskio_read_count,diskio_write_count,diskio_read_bytes,diskio_write_bytes,diskio_read_time,diskio_write_time,network_bytes_sent,network_bytes_recv,network_packets_sent,network_packets_recv)"+
		"values('%s','%s','%s','%s','%s','%d','%s','%d','%d','%s','%f','%f','%d','%d','%d','%d','%f','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d','%d')", eventEntity, hostname, os, platform, version, uptime, bootTime, cpuPhysicalNum, cpuLogicalNum, eventTag, load, cpuPercent, memTotal/1024/1024, memUsed/1024/1024, memFree/1024/1024, memAvailable/1024/1024, memUsedPercent, swapTotal/1024/1024, swapFree/1024/1024, swapCached/1024/1024, maxDiskUsedPct, maxInodeUsedPct, diskIoReadCountSum, diskIoWriteCountSum, diskIoReadBytesSum, diskIoWriteBytesSum, diskIoReadTimeSum, diskIoWriteTimeSum, netBytesSentSum, netBytesRecvSum, netPacketsSentSum, netPacketsRecvSum)
	sendEvent(sqlData, "sql")

}

func sendEvent(Data interface{}, sendType string) {

	proxy := conf.Option["proxy"]

	var proxyUrl string
	if sendType == "event" {
		proxyUrl = proxy + "/proxy/event"
	}
	if sendType == "sql" {
		proxyUrl = proxy + "/proxy/sql"
	}
	_, err := http.Post(proxyUrl, Data)

	log.Debug(fmt.Sprintf("Start send events to proxy %s ", proxyUrl))
	if err != nil {
		log.Error(fmt.Sprintf("Failed send events to proxy %s error: %s", proxyUrl, err))

	}

}

func main() {
	for true {
		go osInfoCollector()
		time.Sleep(time.Duration(utils.StrToInt(conf.Option["interval"])) * time.Second)
	}
}
