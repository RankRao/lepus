### Lepus介绍

* Lepus数据库监控系统是简洁、直观、强大的开源数据库监控产品，支持MySQL/Oracle/MongoDB/Redis等数据库一站式性能监控，让数据库监控更加简单和专业。远程云中监控、实时邮件告警、丰富的指标和图表、MySQL慢查询分析和AWR性能报告。

* Lepus致力于打造开源的企业级智能化监控和运维管理平台。


### 安装部署

##### 1.编译代码

$ go build -ldflags -w -a ../src/lepus_proxy/lepus_proxy.go

$ go build -ldflags -w -a ../src/lepus_alarm/lepus_alarm.go

$ go build -ldflags -w -a ../src/lepus_task/lepus_task.go

$ go build -ldflags -w -a ../src/lepus_collector/mysql/lepus_mysql_mon.go

$ go build -ldflags -w -a ../src/lepus_collector/redis/lepus_redis_mon.go

##### 2.启动Lepus

$ ./lepus_proxy --config=../etc/proxy.ini

$ ./lepus_alarm --config=../etc/alarm.ini

$ ./lepus_task --config=../etc/config.ini

##### 3.其他模块

Lepus5.0版本之后将会以Group的方式运行，此仓库为Lepus核心代码模块, 更多仓库模块请进入 [https://gitee.com/lepus-group](https://gitee.com/lepus-group)


##### 4.使用3.X版本
- 使用Lepus3.X版本，请使用lepus仓库代码 v3.8版本标签下载：[https://gitee.com/lepus-group/lepus/tree/v3.8](https://gitee.com/lepus-group/lepus/tree/v3.8)
部署文档参考官方网站手册。
- 使用Lepus5.X及以上版本，需编译部署lepus、lepus-console两个仓库，如果你不需要修改源码，推荐使用我们编译好的二进制包安装（lepus-bin仓库）快速安装部署，部署文档：[点击查看部署文档](https://discuss.lepus.cc/d/5-lepus-50/7)。


### 使用案例
据以往统计和反馈，这些公司曾经或者目前在使用Lepus（历史统计不代表目前继续在使用，如有出入可以联系我们删除）：
- 飞牛网
- OPPO
- 招商银行信用卡
- 乐视
- PHPOK
- 平安好房
- 同程旅游
- 时光网

### 软件快照


Lepus3.x版本：
![Image description](https://discuss.lepus.cc/assets/files/2022-03-20/1647748068-110473-lepus3.jpg)

lepus5.X版本：
![Image description](https://discuss.lepus.cc/assets/files/2022-03-20/1647748134-751835-20220319194313.png)

### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

### 如何加入
参与项目开发请优先提交requests，如果想加入团队一起开发请发送申请邮件至ruyi@139.com （requests将作为能否加入我们的重要参考依据）

### 捐助作者
LEPUS项目由作者和贡献者业余时间开发和维护，作者负责软件的设计、前后端开发、系统测试、数据库和服务运维。除了耗费大量的业余时间开发和维护，每年还需要支付很多服务器费用部署数据库。如果您觉得Lepus对您的工作有帮助，欢迎捐助我们，捐助费用将用于服务器的续费和软件维护,您的捐助是我们持续的动力，[点击进行捐助](https://www.lepus.cc/crowdfunding/)

### 联系和支持
网站：
[https://www.lepus.cc](https://www.lepus.cc)

社区：
[https://discuss.lepus.cc](https://discuss.lepus.cc)

公众号：
公众号定时推送Lepus更新通知和技术分享，以及数据库、大数据干货技术，欢迎关注作者公众号。
![输入图片说明](https://discuss.lepus.cc/assets/files/2022-02-06/1644160856-809885-qrcode-for-gh-0ae0d3832970-258.jpg)