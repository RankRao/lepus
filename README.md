# dbmcloud

#### 介绍
通过Golang重构lepus项目

#### 运行后端服务

安装golang
# cd /usr/local/
# wget  https://golang.google.cn/dl/go1.18.10.linux-amd64.tar.gz
# tar -zxvf go1.18.10.linux-amd64.tar.gz
# vim ~/.bash_profile

配置go环境变量
PATH=$PATH:$HOME/bin:/usr/local/go/bin/
# source ~/.bash_profile
# go version
go version go1.18.10 linux/amd64

配置环境变量
# mkdir /home/golang
# go env -w GO111MODULE=on
# go env -w GOPROXY=https://goproxy.cn,direct
# go env -w GOPATH=/home/golang
# go env

安装依赖包
# git clone https://gitee.com/lepus/dbmcloud.git
# cd dbmcloud/
# go mod tidy
# go mod vendor

配置文件
复制setting.example.yml文件为setting.yml,并修改配置信息



运行服务
go run main.go


#### 运行前端服务



#### 前后端编译打包

1.前端web目录执行npm run build编译前端代码，编译后的文件位于dist目录
2.将前端编译好的dist目录复制到后端根目录（和main.go同一级），并重命名为static
3.将static里面的index.html文件复制到后端根目录（和main.go同一级）
4.进入后端bin目录，执行以下命令打包：
bin> go build -ldflags -w -a -o dbmcloud.exe  ..\main.go

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
