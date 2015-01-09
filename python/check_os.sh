#!/bin/bash

ip=$1
host=$2
port=$3
user=$4
password=$5
dbname=$6
community=$7
filter_os_disk=$8
tags=$9
#echo $tags
if [ x"$filter_os_disk" = x ];then
    filter_os_disk="none"
fi

mysql_client="mysql -h${host}  -P${port}  -u${user}  -p${password}"
RETRIES=3
TIMEOUT=10
snmpwalk="/usr/bin/snmpwalk -r $RETRIES -t $TIMEOUT"
snmpget="/usr/bin/snmpget -r $RETRIES -t $TIMEOUT"
snmpdf="/usr/bin/snmpdf -r $RETRIES -t $TIMEOUT"

snmpdata=$($snmpget -v1 -c ${community} ${ip} SNMPv2-MIB::sysName.0 \
    SNMPv2-MIB::sysDescr.0 \
    HOST-RESOURCES-MIB::hrSystemDate.0 \
    HOST-RESOURCES-MIB::hrSystemUptime.0 \
    HOST-RESOURCES-MIB::hrSystemProcesses.0 \
    UCD-SNMP-MIB::laLoad.1 \
    UCD-SNMP-MIB::laLoad.2 \
    UCD-SNMP-MIB::laLoad.3 \
    UCD-SNMP-MIB::ssCpuUser.0 \
    UCD-SNMP-MIB::ssCpuSystem.0 \
    UCD-SNMP-MIB::ssCpuIdle.0 \
    UCD-SNMP-MIB::memTotalSwap.0 \
    UCD-SNMP-MIB::memAvailSwap.0 \
    UCD-SNMP-MIB::memTotalReal.0 \
    UCD-SNMP-MIB::memAvailReal.0 \
    UCD-SNMP-MIB::memTotalFree.0 \
    UCD-SNMP-MIB::memShared.0 \
    UCD-SNMP-MIB::memBuffer.0 \
    UCD-SNMP-MIB::memCached.0 \
    2>/dev/null)

if [ "$snmpdata" != "" ]; then
    hostname=`echo "$snmpdata" | grep SNMPv2-MIB::sysName.0|awk '{print $NF}'`
    kernel=`echo "$snmpdata" | grep SNMPv2-MIB::sysDescr.0|awk '{print $4 " " $6 " " $15}'`
    system_date=`echo "$snmpdata" | grep HOST-RESOURCES-MIB::hrSystemDate.0|cut -d '=' -f2|cut -d ' ' -f3`
    system_uptime=`echo "$snmpdata" | grep HOST-RESOURCES-MIB::hrSystemUptime.0|cut -d ')' -f2`
    process=`echo "$snmpdata" | grep HOST-RESOURCES-MIB::hrSystemProcesses.0|cut -d ' ' -f4`

    load_1=`echo "$snmpdata" | grep UCD-SNMP-MIB::laLoad.1|awk '{print $NF}'`
    load_5=`echo "$snmpdata" | grep UCD-SNMP-MIB::laLoad.2|awk '{print $NF}'`
    load_15=`echo "$snmpdata" | grep UCD-SNMP-MIB::laLoad.3|awk '{print $NF}'`

    cpu_user_time=`echo "$snmpdata" | grep UCD-SNMP-MIB::ssCpuUser.0 |awk '{print $NF}'`
    cpu_system_time=`echo "$snmpdata" | grep UCD-SNMP-MIB::ssCpuSystem.0 |awk '{print $NF}'`
    cpu_idle_time=`echo "$snmpdata" | grep UCD-SNMP-MIB::ssCpuIdle.0 |awk '{print $NF}'`

    swap_total=`echo "$snmpdata" | grep UCD-SNMP-MIB::memTotalSwap.0 |cut -d= -f2 |awk -F ' ' '{print $2}'`
    swap_avail=`echo "$snmpdata" | grep UCD-SNMP-MIB::memAvailSwap.0 |cut -d= -f2 |awk -F ' ' '{print $2}'`
    mem_total=`echo "$snmpdata" | grep UCD-SNMP-MIB::memTotalReal.0 |cut -d= -f2 |awk -F ' ' '{print $2}'`
    mem_used=`echo "$snmpdata" | grep UCD-SNMP-MIB::memAvailReal.0 |cut -d= -f2 |awk -F ' ' '{print $2}'`
    mem_free=`echo "$snmpdata" | grep UCD-SNMP-MIB::memTotalFree.0 |cut -d= -f2 |awk -F ' ' '{print $2}'`
    mem_shared=`echo "$snmpdata" | grep UCD-SNMP-MIB::memShared.0 |cut -d= -f2 |awk -F ' ' '{print $2}'`
    mem_buffered=`echo "$snmpdata" | grep UCD-SNMP-MIB::memBuffer.0 |cut -d= -f2 |awk -F ' ' '{print $2}'`
    mem_cached=`echo "$snmpdata" | grep UCD-SNMP-MIB::memCached.0 |cut -d= -f2 |awk -F ' ' '{print $2}'`

    mem_usage_rate=`$snmpdf -v1 -c ${community}  ${ip}  |grep "Physical"|awk '{print $6}'`
    mem_available=$[$mem_free+$mem_buffered+$mem_cached]

   #disk usage
   IFS=$'\n'
   disk_all=`$snmpdf -v1 -c  ${community} ${ip} |grep -E "/"|grep -vE "/dev|/boot" |grep -vE "$filter_os_disk"`
   for line in $disk_all
   do
      IFS=' '
      mounted=`echo $line|awk -F ' ' '{print $1}' `
      total_size=`echo $line|awk -F ' ' '{print $2}' `
      used_size=`echo $line|awk -F ' ' '{print $3}' `
      avail_size=`echo $line|awk -F ' ' '{print $4}' `
      used_rate=`echo $line|awk -F ' ' '{print $5}' `
      $mysql_client -N -e "insert  into $dbname.os_disk(ip,tags,mounted,total_size,used_size,avail_size,used_rate) values('${ip}','${tags}','${mounted}','${total_size}','${used_size}','${avail_size}','${used_rate}')"
   done

   #disk io
   disk_io_reads_total=0
   disk_io_writes_total=0
   fdisk_io_string=`$snmpwalk -v1 -c ${community} ${ip}  UCD-DISKIO-MIB::diskIODevice |grep -ivE "ram|loop|md"`
   IFS=$'\n'
   for line in $fdisk_io_string
   do
      IFS=' '
      fdisk=`echo $line|awk -F ' ' '{print $4}'`
      fdisk_id=`echo $line|awk -F ' ' '{print $1}'|awk -F '.' '{print $2}'`
      disk_io_reads_1=`$snmpwalk -v1 -c ${community} ${ip} UCD-DISKIO-MIB::diskIOReads.$fdisk_id|awk '{print $NF}'`
      disk_io_writes_1=`$snmpwalk -v1 -c ${community} ${ip} UCD-DISKIO-MIB::diskIOWrites.$fdisk_id|awk '{print $NF}'`
      sleep 1
      disk_io_reads_2=`$snmpwalk -v1 -c ${community} ${ip} UCD-DISKIO-MIB::diskIOReads.$fdisk_id|awk '{print $NF}'`
      disk_io_writes_2=`$snmpwalk -v1 -c ${community} ${ip} UCD-DISKIO-MIB::diskIOWrites.$fdisk_id|awk '{print $NF}'`
      disk_io_reads=$(($disk_io_reads_2-$disk_io_reads_1))
      disk_io_writes=$(($disk_io_writes_2-$disk_io_writes_1))
      #disk_io_reads=`expr $disk_io_reads / 3`
      #disk_io_writes=`expr $disk_io_writes / 3`
      $mysql_client -N -e "insert  into $dbname.os_diskio(ip,tags,fdisk,disk_io_reads,disk_io_writes) values('${ip}','${tags}','${fdisk}','${disk_io_reads}','${disk_io_writes}')"
      let disk_io_reads_total=$disk_io_reads_total+$disk_io_reads
      let disk_io_writes_total=$disk_io_writes_total+$disk_io_writes
   done


   #net
   net_in_bytes_total=0
   net_out_bytes_total=0
   net_descr_string=`$snmpwalk -v1 -c ${community} ${ip} IF-MIB::ifDescr`
   IFS=$'\n'
   for line in $net_descr_string
   do
       IFS=' '
       net_descr=`echo $line|awk -F '=' '{print $2}' |awk -F ': ' '{print $2}'`
       net_descr_id=`echo $line|awk -F ' ' '{print $1}'|awk -F '.' '{print $2}'`
       in_bytes_1=`$snmpwalk -v1 -c ${community} ${ip} IF-MIB::ifInOctets.$net_descr_id|awk '{print $NF}'`
       out_bytes_1=`$snmpwalk -v1 -c ${community} ${ip} IF-MIB::ifOutOctets.$net_descr_id|awk '{print $NF}'`
       sleep 1
       in_bytes_2=`$snmpwalk -v1 -c ${community} ${ip} IF-MIB::ifInOctets.$net_descr_id|awk '{print $NF}'`
       out_bytes_2=`$snmpwalk -v1 -c ${community} ${ip} IF-MIB::ifOutOctets.$net_descr_id|awk '{print $NF}'`
       in_bytes=$(($in_bytes_2-$in_bytes_1))
       out_bytes=$(($out_bytes_2-$out_bytes_1))
       #in_bytes=`expr $in_bytes / 3`
       #out_bytes=`expr $out_bytes / 3`
       $mysql_client -N -e "insert  into $dbname.os_net(ip,tags,if_descr,in_bytes,out_bytes) values('${ip}','${tags}','${net_descr}','${in_bytes}','${out_bytes}')"
       let net_in_bytes_total=$net_in_bytes_total+$in_bytes
       let net_out_bytes_total=$net_out_bytes_total+$out_bytes
   done

   $mysql_client -N -e "insert into  $dbname.os_status_history select *, LEFT(REPLACE(REPLACE(REPLACE(create_time,'-',''),' ',''),':',''),12) from $dbname.os_status where ip='${ip}';"
   $mysql_client -N -e "delete from  $dbname.os_status where ip='${ip}';"
   $mysql_client -N -e "insert  into $dbname.os_status(ip,snmp,tags,hostname,kernel,system_date,system_uptime,process,load_1,load_5,load_15,cpu_user_time,cpu_system_time,cpu_idle_time,swap_total,swap_avail,mem_total,mem_used,mem_free,mem_shared,mem_buffered,mem_cached,mem_usage_rate,mem_available,disk_io_reads_total,disk_io_writes_total,net_in_bytes_total,net_out_bytes_total) values('${ip}',1,'${tags}','${hostname}','${kernel}','${system_date}','${system_uptime}','${process}','${load_1}','${load_5}','${load_15}','${cpu_user_time}','${cpu_system_time}','${cpu_idle_time}','${swap_total}','${swap_avail}','${mem_total}','${mem_used}','${mem_free}','${mem_shared}','${mem_buffered}','${mem_cached}','${mem_usage_rate}','${mem_available}','${disk_io_reads_total}','${disk_io_writes_total}','${net_in_bytes_total}','${net_out_bytes_total}')"


else
   $mysql_client -N -e "insert  into $dbname.os_status(ip,tags,snmp) values('${ip}','${tags}',0)"
fi
