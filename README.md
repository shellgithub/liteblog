## GoLang 开发 liteblog 练习与学习笔记

## 相关文档地址

- https://beego.me/docs

- https://beego.me/docs/intro/

### bee 工具简介

bee 工具是一个为了协助快速开发 beego 项目而创建的项目，通过 bee 您可以很容易的进行 beego 项目的创建、热编译、开发、测试、和部署。

#### bee 工具的安装

您可以通过如下的方式安装 bee 工具：

```
go get -u github.com/beego/bee/v2

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
echo 'export GOPATH=/data/go/' >> ~/.zshrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
```
安装完之后，bee 可执行文件默认存放在 `$GOPATH/bin `里面，所以您需要把` $GOPATH/bin` 添加到您的环境变量中，才可以进行下一步。

如果你本机设置了 GOBIN，那么上面的bee命令就会安装到 GOBIN 目录下，所以我们需要在环境变量中添加相关的配置信息，如何添加可以查看这篇文档: bee 环境变量配置

#### bee 工具命令详解

我们在命令行输入 bee，可以看到如下的信息：

```
go get -u github.com/beego/bee/v2
cd $GOPATH/src  

go get github.com/beego/bee
go get github.com/astaxie/beego
go get github.com/shiena/ansicolor

bee new liteblog

cd $GOPATH/src/liteblog
go run main.go
```

#### 下载 博客模版（使用QQ登录可以下载）

https://www.sucainiu.com/5921.html

#### 表单生成工具地址

https://www.layui.com/demo/form.html

#### 使用代理下载依赖包

```
go env -w GOPROXY="https://goproxy.cn,direct"
go env -w GOPROXY="https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct"
```

#### 下载 gorm

- https://gorm.io/docs/

```shell script
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

### 下载 orm

```
go get github.com/beego/beego/v2/client/orm
go get github.com/go-sql-driver/mysql@v1.5.0
```

### 添加依赖模块

```
初始化依赖模块
go mod init

添加依赖模块
go mod tidy

```

### 下载wangEditor

- 2021-03-26

- 文件地址：https://www.kancloud.cn/wangfupeng/wangeditor/65722
- Https://github.com/wangfupeng1988/wangEditor
- https://github.com/wangeditor-team/wangEditor/tags?after=v0.0.7

```
# 视频中的版本
https://github.com/wangeditor-team/wangEditor/archive/refs/tags/v3.1.1.tar.gz

# 当前最新版本 2021-03-26
https://github.com/wangeditor-team/wangEditor/archive/refs/tags/v4.6.12.tar.gz
```

### UUID

- 2021-03-26

```
go get github.com/satori/go.uuid
```

### net

- 2021-03-27--2021-03-29

```
go get -u -v golang.org/x/net
go get github.com/PuerkitoBio/goquery
```
### mapstructure

- 2021-03-29
```
go get github.com/goinggo/mapstructure
```













