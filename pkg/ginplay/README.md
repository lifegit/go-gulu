# ginfire(gin and gorm's fire) 详解

## 📌是什么

Go语言的RESTful APIs脚手架工具。


## 📦安装felix

```bash
# 自动编译
go get https://github.com/lifegit/go-gulu/pkg/ginplay

echo "go build && ./ginplay create -h"

# 源码安装
git clone https://github.com/lifegit/go-gulu
cd pkg/ginplay
go mod download

go install
echo "添加 GOBIN 到 PATH环境变量"
```

### ⚠️命令行参数详解

```bash
[root@ericzhou felix]# go build && ./ginplay create -h
generate a RESTful APIs app with gin and gorm for gophers

Usage:
  create [flags]

Examples:
go build -o main && ./main create \
-d ./go-admin \
-k go-admin \
-b admin \
-u root \
-s pass \
-n db_test


Flags:
  -a, --appAddr string      http service bind address (default "127.0.0.1")                 生成app接口监听的地址
  -d, --appDir string       code project output directory                                   golang代码输出的目录
  -k, --appPkg string       go.mod module name                                              生成go app 包名称同时生成go.mod文件
  -o, --appPort int         http service bind port (default 8080)                           生成app接口监听的端口
  -l, --authColumn string   bcrypt password column (default "password")                     使用bcrypt方式加密的用户表密码字段名称
  -b, --authTable string    login user table (default "users")                              认知登陆用户表名称
  -e, --authType int        generate 1:mobile or 2:security on register、forget (default 1) 注册与找回密码的方式(手机号、密保问题)
  -r, --dbAddr string       database connection addr (default "127.0.0.1")                  数据库地址
  -c, --dbChar string       database charset (default "utf8")                               数据库字符集合
  -n, --dbName string       database name                                                   数据库名称
  -s, --dbPassword string   database user password                                          数据库密码
  -p, --dbPort int          database connection addr (default 3306)                         数据库端口
  -t, --dbType string       database type: mysql/postgres/mssql/sqlite (default "mysql")    数据库类型: mysql/postgres/mssql/sqlite
  -u, --dbUser string       database username (default "root")                              数据库用户名
  -h, --help                help for ginplay                                                帮助

```

### ❤️功能简介
- 每一张数据库表生成一个RESTful规范的资源(`GET-pagination/GET-one/POST/PUT/DELETE`)
- 支持`jwt-token`认证和`Bearer Token`路由中间件; `gin autotls` 开启免证书 https。
- 开箱即用的定时任务[core](https://github.com/lifegit/go-gulu/tree/master/nice/core)。
- 开箱即用的 [logrus](https://github.com/sirupsen/logrus)日志工具。
- 开箱即用的 [fire](https://github.com/lifegit/go-gulu/tree/master/pkg/fire)，[gorm](https://gorm.cn/) 增强工具。
- 开箱即用的 [OpenApi 3](https://swagger.io/specification/)  和 [OpenApi 2](https://swagger.io/specification/v2/) 文档。
- 开箱即用的 Dockerfile、Docker Compose
- 开箱即用的[viper](https://github.com/spf13/viper) 、[fsnotify](https://github.com/fsnotify/fsnotify)对toml配置文件读取，监听。
- 开箱即用的upload上传功能，支持上传到本地或aliyunOss


### ✏️工作原理

1. 通过 [cobra](github.com/spf13/cobra) 获取命令行参数。
2. 使用 [gormt](https://github.com/xxjwxc/gormt) 连接数据库并获取数据库表的名称和字段类型等，同时生成 GORM 模型字段。
5. 使用标准库 [`text/template`](https://golang.google.cn/pkg/text/template/) 生成 GORM 模型文件, GIN handler 文件 ...
6. 根据备注生成 OpenApi2和3的 `.json` `.yaml` 文件。
7. 使用 `go fmt ./...` 格式化代码


### 👍支持数据库大多数SQL数据库
- mysql
- SQLite
- postgreSQL
- mssql(TODO:: sqlserver)


## 🙏感谢
- [dejavuzhou/felix：鸣谢 ginbro](https://github.com/dejavuzhou/felix)，[作者 wiki MojoTech](https://tech.mojotv.cn/2019/05/22/golang-felix-ginbro)
- [lifegit/go-gulu 开箱即用工具集](https://github.com/lifegit/go-gulu)