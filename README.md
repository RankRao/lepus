# lepus

#### 介绍
Lepus核心代码模块,更多模块请进入 https://gitee.com/lepus-group

#### 编译代码

$ go build -ldflags -w -a ../src/lepus_proxy/lepus_proxy.go

$ go build -ldflags -w -a ../src/lepus_alarm/lepus_alarm.go

$ go build -ldflags -w -a ../src/lepus_task/lepus_task.go

$ go build -ldflags -w -a ../src/lepus_collector/mysql/lepus_mysql_mon.go

$ go build -ldflags -w -a ../src/lepus_collector/redis/lepus_redis_mon.go

#### 启动Lepus

$ ./lepus_proxy --config=../etc/proxy.ini

$ ./lepus_alarm --config=../etc/alarm.ini

$ ./lepus_task --config=../etc/config.ini

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 支持
https://www.lepus.cc

https://gitee.com/lepus-group

