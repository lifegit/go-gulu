package tpl

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/rand"
)

var tAppApp = tplNode{
	NameFormat: "app/app.go",
	TplContent: `
package app

func init() {
	SetUpConf()
	SetUpBasics()

	SetUpCache()
	SetUpDB()

	SetUpResult()
	SetUpFileUpload()
}

func Close()  {
	_ = Cache.Close()
	_ = DB.Close()
}
`,
}
var tAppBasics = tplNode{
	NameFormat: "app/basics.go",
	TplContent: `
package app

import (
	"github.com/lifegit/go-gulu/v2/pkg/logging"
	"github.com/sirupsen/logrus"
	"time"
)

var Log *logrus.Logger

func SetUpBasics() {
	// timeZone
	_, _ = time.LoadLocation(Global.App.TimeZone)

	// log
	Log = logging.NewLogger(Global.App.Log, 3, &logrus.TextFormatter{}, nil)
}
`,
}
var tAppCache = tplNode{
	NameFormat: "app/cache.go",
	TplContent: `
package app

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/lifegit/go-gulu/v2/pkg/gredis"
)

var Cache *redis.Client

func SetUpCache() {
	redisClient, err := gredis.CreateRedis(fmt.Sprintf("%s:%d", Global.Redis.Addr, Global.Redis.Port), Global.Redis.Password, Global.Redis.DbIdx)
	if err != nil {
		Log.WithError(err).Fatal("could not connect to the redis server")
	}
	Cache = redisClient
}
`,
}
var tAppConf = tplNode{
	NameFormat: "app/conf.go",
	TplContent: `
package app

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"github.com/lifegit/go-gulu/v2/pkg/viperine"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

var Global GlobalConf

type GlobalConf struct {
	App struct {
		Name     string _[BACKQUOTE]_toml:"name"_[BACKQUOTE]_
		Version  int    _[BACKQUOTE]_toml:"version"_[BACKQUOTE]_
		TimeZone string _[BACKQUOTE]_toml:"timeZone"_[BACKQUOTE]_
		Env      string _[BACKQUOTE]_toml:"env"_[BACKQUOTE]_
		IsCron     bool   _[BACKQUOTE]_toml:"cron"_[BACKQUOTE]_
		Log      string _[BACKQUOTE]_toml:"log"_[BACKQUOTE]_
	} _[BACKQUOTE]_toml:"app"_[BACKQUOTE]_
	Server struct {
		Addr         string _[BACKQUOTE]_toml:"addr"_[BACKQUOTE]_
		Port         int    _[BACKQUOTE]_toml:"port"_[BACKQUOTE]_
		RunMode      string _[BACKQUOTE]_toml:"runMode"_[BACKQUOTE]_
		ReadTimeout  int    _[BACKQUOTE]_toml:"readTimeout"_[BACKQUOTE]_
		WriteTimeout int    _[BACKQUOTE]_toml:"writeTimeout"_[BACKQUOTE]_
		APIPrefix    string _[BACKQUOTE]_toml:"apiPrefix"_[BACKQUOTE]_
		StaticPath   string _[BACKQUOTE]_toml:"staticPath"_[BACKQUOTE]_
		HTTPLogDir   string _[BACKQUOTE]_toml:"httpLogDir"_[BACKQUOTE]_
		IsNotRoute     bool   _[BACKQUOTE]_toml:"isNotRoute"_[BACKQUOTE]_
		IsSwagger      bool   _[BACKQUOTE]_toml:"isSwagger"_[BACKQUOTE]_
		IsCors         bool   _[BACKQUOTE]_toml:"isCors"_[BACKQUOTE]_
		IsHTTPS        bool   _[BACKQUOTE]_toml:"isHttps"_[BACKQUOTE]_
	} _[BACKQUOTE]_toml:"server"_[BACKQUOTE]_
	Jwt struct {
		Secret     string _[BACKQUOTE]_toml:"secret"_[BACKQUOTE]_
		Key        string _[BACKQUOTE]_toml:"key"_[BACKQUOTE]_
		ExpireHour int    _[BACKQUOTE]_toml:"expireHour"_[BACKQUOTE]_
	} _[BACKQUOTE]_toml:"jwt"_[BACKQUOTE]_
	Upload struct {
		Type           string _[BACKQUOTE]_toml:"type"_[BACKQUOTE]_
		Domain         string _[BACKQUOTE]_toml:"domain"_[BACKQUOTE]_
		Local          struct {
			BaseDir string _[BACKQUOTE]_toml:"baseDir"_[BACKQUOTE]_
		} _[BACKQUOTE]_toml:"local"_[BACKQUOTE]_
		Oss struct {
			AccessKeyID     string _[BACKQUOTE]_toml:"accessKeyID"_[BACKQUOTE]_
			AccessKeySecret string _[BACKQUOTE]_toml:"accessKeySecret"_[BACKQUOTE]_
			Endpoint        string _[BACKQUOTE]_toml:"endpoint"_[BACKQUOTE]_
			BucketName      string _[BACKQUOTE]_toml:"bucketName"_[BACKQUOTE]_
		} _[BACKQUOTE]_toml:"oss"_[BACKQUOTE]_
	} _[BACKQUOTE]_toml:"upload"_[BACKQUOTE]_
	Db struct {
		Type     string _[BACKQUOTE]_toml:"type"_[BACKQUOTE]_
		Addr     string _[BACKQUOTE]_toml:"addr"_[BACKQUOTE]_
		Port     int    _[BACKQUOTE]_toml:"port"_[BACKQUOTE]_
		Username string _[BACKQUOTE]_toml:"username"_[BACKQUOTE]_
		Password string _[BACKQUOTE]_toml:"password"_[BACKQUOTE]_
		Database string _[BACKQUOTE]_toml:"database"_[BACKQUOTE]_
		Charset  string _[BACKQUOTE]_toml:"charset"_[BACKQUOTE]_
	} _[BACKQUOTE]_toml:"db"_[BACKQUOTE]_
	Redis struct {
		Addr     string _[BACKQUOTE]_toml:"addr"_[BACKQUOTE]_
		Port     int    _[BACKQUOTE]_toml:"port"_[BACKQUOTE]_
		Password string _[BACKQUOTE]_toml:"password"_[BACKQUOTE]_
		DbIdx    int    _[BACKQUOTE]_toml:"dbIdx"_[BACKQUOTE]_
	} _[BACKQUOTE]_toml:"redis"_[BACKQUOTE]_
	SafeLogin struct {
		PrivateKey string _[BACKQUOTE]_toml:"privateKey"_[BACKQUOTE]_
		TryPerson  struct {
			Max  int _[BACKQUOTE]_toml:"max"_[BACKQUOTE]_
			Time int _[BACKQUOTE]_toml:"time"_[BACKQUOTE]_
		} _[BACKQUOTE]_toml:"tryPerson"_[BACKQUOTE]_
		TryIP struct {
			Max  int _[BACKQUOTE]_toml:"max"_[BACKQUOTE]_
			Time int _[BACKQUOTE]_toml:"time"_[BACKQUOTE]_
		} _[BACKQUOTE]_toml:"tryIp"_[BACKQUOTE]_
	} _[BACKQUOTE]_toml:"safeLogin"_[BACKQUOTE]_
}

const DEV = "dev"

func (g *GlobalConf) isDev() bool {
	return g.getEnv() == DEV
}
func (g *GlobalConf) getEnv() (res string) {
	if res = os.Getenv("GO_ENV"); res == "" {
		res = DEV
	}

	return res
}
func SetUpConf() {
	basePath := recursionPath("conf")
	v, err := viperine.LocalConfToViper([]string{
		path.Join(basePath, "base.toml"),
		path.Join(basePath, Global.getEnv(), "conf.toml"),
	}, &Global, func(event fsnotify.Event, viper *viper.Viper) {
		if event.Op != fsnotify.Remove {
			_ = viper.Unmarshal(&Global)
		}
	})

	if err != nil {
		logrus.WithError(err).Fatal(err, v)
	}
}

func recursionPath(dirName string) (dirPath string) {
	var dir string
	for i := 0; i < 10; i++ {
		dirPath = path.Join(dir, dirName)
		dir = path.Join(dir, "../")

		if file.IsDir(dirPath) {
			return
		}
	}

	return
}`}
var tAppDb = tplNode{
	NameFormat: "app/db.go",
	TplContent: `
package app

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	//"gorm.io/driver/sqlite"
	//"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
)

// init mysql
var DB *fire.Fire

func SetUpDB() {
	if db, err := OpenConnection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		DB = fire.NewInstance(db)
		sqlDB, err := db.DB()
		if err == nil {
			err = sqlDB.Ping()
		}
		if err != nil {
			log.Printf("failed to connect database, got error %v", err)
		}

		// ConnPool db conn pool
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}
}



func OpenConnection() (db *gorm.DB, err error) {
	gormConf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// https://gorm.io/docs/gorm_config.html#NamingStrategy
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NameReplacer:  strings.NewReplacer("Tb", ""),
		}, // 表名不加复数s
	}
	switch Global.Db.Type {
	case "mysql":
		dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", Global.Db.Username, Global.Db.Password, Global.Db.Addr, Global.Db.Port, Global.Db.Database, Global.Db.Charset)
		db, err = gorm.Open(mysql.Open(dbDSN), gormConf)
	//case "postgres":
	//	log.Println("testing postgres...")
	//	if dbDSN == "" {
	//		dbDSN = "user=gorm password=gorm dbname=gorm host=localhost port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//	}
	//	db, err = gorm.Open(postgres.New(postgres.Config{
	//		DSN:                  dbDSN,
	//		PreferSimpleProtocol: true,
	//	}), gormConf)
	//case "sqlserver":
	//	// CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';
	//	// CREATE DATABASE gorm;
	//	// USE gorm;
	//	// CREATE USER gorm FROM LOGIN gorm;
	//	// sp_changedbowner 'gorm';
	//	// npm install -g sql-cli
	//	// mssql -u gorm -p LoremIpsum86 -d gorm -o 9930
	//	log.Println("testing sqlserver...")
	//	if dbDSN == "" {
	//		dbDSN = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	//	}
	//	db, err = gorm.Open(sqlserver.Open(dbDSN), gormConf)
	//default:
	//	log.Println("testing sqlite3...")
	//	db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), gormConf)
	}

	return
}
`,
}
var tAppResult = tplNode{
	NameFormat: "app/result.go",
	TplContent: `
/**
* @Author: TheLife
* @Date: 2020-10-30 2:37 上午
 */
package app

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/ginMiddleware/mwCors"
	"github.com/lifegit/go-gulu/v2/pkg/ginMiddleware/mwJwt"
	"github.com/lifegit/go-gulu/v2/pkg/ginMiddleware/mwLogger"
	"net/http"
	"reflect"
	"time"
)

var Result ResultApi

type ResultApi struct {
	Gin *gin.Engine
	Api *gin.RouterGroup

	Jwt mwJwt.MwJwt
}

func SetUpResult() {
	Result.Setup()
}

func (r *ResultApi) Setup() {
	// 设置模式，设置模式要放在调用Default()函数之前
	addr := fmt.Sprintf("%s:%d", Global.Server.Addr, Global.Server.Port)
	// mode
	gin.SetMode(Global.Server.RunMode)
	r.Gin = gin.New()
	// middleware
	// Recovery
	r.Gin.Use(gin.Recovery())
	// Cors
	if Global.Server.IsCors {
		r.Gin.Use(mwCors.NewCorsMiddleware())
	}
	// Logger
	r.Gin.Use(mwLogger.NewLoggerMiddlewareSmoothFail(true, Global.Server.HTTPLogDir))
	// Jwt
	r.Jwt = mwJwt.NewJwtMiddleware(Global.Jwt.Key, Global.App.Name, Global.Jwt.Secret, Global.Jwt.Key, reflect.TypeOf(JwtUser{}), func(e error) (code int, jsonObj interface{}) {
		return http.StatusUnauthorized, gin.H{"msg": "未登录!请先登录!"}
	})

	// staticPath
	if staticPath := Global.Server.StaticPath; staticPath != "" {
		Log.Info(fmt.Sprintf("visit http://%s/ for front-end static html files", addr))
		r.Gin.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	}
	// notRoute
	if Global.Server.IsNotRoute {
		Log.Info(fmt.Sprintf("visit http://%s/404 for RESTful Is NotRoute", addr))
		r.Gin.NoRoute(func(c *gin.Context) {
			c.Status(http.StatusNotFound)
		})
	}
	// swaggerApi
	if Global.Server.IsSwagger && Global.isDev() {
		Log.Info(fmt.Sprintf("visit http://%s/swagger/index.html for RESTful APIs Document", addr))
		r.Gin.Static("swagger", "docs/swagger")
		r.GET("swagger/*any", func(c *gin.Context) {
			if strings.Contains(c.Request.RequestURI, "doc.json") {
				v, _ := file.ReadFile("docs/swagger/v3/openapi.json")
				_, _ = c.Writer.WriteString(v)
				c.Abort()
			}
		}, ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// appInfo
	if Global.isDev() {
		Log.Info(fmt.Sprintf("visit http://%s/app/info for app info only on not-prod mode", addr))
		r.Gin.GET("/app/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, Global)
		})
	}

	// apiPrefix
	r.Api = r.Gin.Group(Global.Server.APIPrefix)
}

func (r *ResultApi) Running() {
	addr := fmt.Sprintf("%s:%d", Global.Server.Addr, Global.Server.Port)
	Log.Infof("http result server at listening %s", addr)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", Global.Server.Port),
		Handler:        r.Gin,
		ReadTimeout:    time.Second * time.Duration(Global.Server.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(Global.Server.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}
	if Global.Server.IsHTTPS {
		// https
		if err := autotls.Run(r.Gin, addr); err != nil {
			Log.WithError(err).Fatal("https result server fail run !")
		}
		//if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
		//	Log.Errorf("https result server is run err: %v!", err)
		//}
	} else {
		// http
		if err := server.ListenAndServe(); err != nil {
			Log.Errorf("http result server is run err: %v!", err)
		}
	}

	//// endless
	//// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//serverNew := endless.NewServer(endPoint, routersInit)
	//serverNew.BeforeBegin = func(add string) {
	//	Log.Info("Actual pid is %d", syscall.Getpid())
	//}
	//err = serverNew.ListenAndServe()
	//if err != nil {
	//	Log.Error("server err: %v", err)
	//}
}


type JwtUser struct {
	Id       uint
	Username string
}
`,
}
var tAppSMS = tplNode{
	NameFormat: "app/sms.go",
	TplContent: `
package app

import "github.com/lifegit/go-gulu/v2/aliyun"

var SmsClient *aliyun.AliSmsClient

func SetUpSMS() {
	s, err := aliyun.NewSMS(Global.Sms.RegionID, Global.Sms.AccessKeyID, Global.Sms.AccessKeySecret)
	if err != nil {
		Log.WithError(err).Fatal("aliyun sms init error")
	}
	SmsClient = s
}
`,
}
var tAppUpload = tplNode{
	NameFormat: "app/upload.go",
	TplContent: `
package app

import (
	"github.com/lifegit/go-gulu/v2/nice/file/upload"
	"log"
	"path"
)

var (
	FileSaveKey   = "file"
	ImageSavePath = "images/"

	AttrAvatar = upload.FileAttribute{
		Key:     FileSaveKey,
		Exts:    upload.AllowImageExts,
		MaxByte: 3 * 1024 * 1024,
		DirPath: path.Join(ImageSavePath, "avatar"),
	}

	AttrImg = upload.FileAttribute{
		Key:     FileSaveKey,
		Exts:    upload.AllowImageExts,
		MaxByte: 5 * 1024 * 1024,
		DirPath: path.Join(ImageSavePath, "img"),
	}
	AttrResources = upload.FileAttribute{
		Key:     FileSaveKey,
		Exts:    upload.AllowAnyExts,
		MaxByte: 10 * 1024 * 1024,
		DirPath: path.Join("resources"),
	}
)

var FileUploads upload.FileUpload

func SetUpFileUpload() {
	switch Global.Upload.Type {
	case "local":
		FileUploads = upload.NewLocal(Global.Upload.Local.BaseDir, Global.Upload.Domain)
	case "oss":
		c, err := upload.NewOss(Global.Upload.Oss.Endpoint, Global.Upload.Oss.AccessKeyID, Global.Upload.Oss.AccessKeySecret, Global.Upload.Oss.BucketName, Global.Upload.Domain)
		if err != nil {
			log.Fatalln("could not check to the oss server")
		}
		FileUploads = c
	default:
		log.Fatal("upload type is nil")
	}
}
`,
}

var tConfigBaseToml = tplNode{
	NameFormat: "conf/base.toml",
	TplContent: `
# toml 转 struct : https://github.com/xuri/toml-to-go

[app]
    name = "{{.AppPkg}}"
    version = 1
    timeZone = "Asia/Shanghai" # 时区
	isCron = true                # 是否启动内置的后台计划任务
    log = "./runtime/logs/app"
`,
}
var tConfigDevToml = tplNode{
	NameFormat: "conf/dev/conf.toml",
	TplContent: fmt.Sprintf(`
[server]
    addr = "{{.AppAddr}}"          # eg: www.mojotv.cn eg:localhost eg:127.0.0.1
    port = {{.AppPort}}
    runMode = "debug"           # eg:debug eg:release eg:test
    readTimeout = 60
    writeTimeout = 60
    apiPrefix = "services"           # api前缀，一般为版本号,设置为后 {api_prefix}/v1/resource
    staticPath = "./static/"   # 静态路径,必须是绝对路径或相对于go build可执行文件
    httpLogDir = "./runtime/logs/http"   # httpLog目录,空则不开启log
    isNotRoute = true             # 如果为true,则所有未找到的路由都将转到static/index.html。
    isSwagger = true              # 开启 swagger api doc
    isCors = true                 # 开启cors跨域限制
    isHttps = false               # 如果 addr 为域名，则是否开启https,证书来自 LetsEncrypt

[jwt]
    secret = "%s"
    key = "Authorization"
    expireHour = 24            # jwt颁布后生效的小时时间

[upload]
    type = "local"               # eg:oss eg:local
    domain = ""
    [upload.local]
        baseDir = "./runtime/upload/"
    [upload.oss]
        accessKeyID = ""
        accessKeySecret = ""
        endpoint = ""
        bucketName = ""

[db]
    type = "{{.DbType}}"
    addr = "{{.DbAddr}}"
    port = {{.DbPort}}
    username = "{{.DbUser}}"
    password = "{{.DbPassword}}"
    database = "{{.DbName}}"
    charset = "{{.DbChar}}"


[redis]
    addr = "127.0.0.1"          # 如果为空将不会初始化 redis
    port = 6379
    password = ""
    dbIdx = 0


[safeLogin]
    privateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICWgIBAAKBgFdAHsGcPdpPwB1Zz09Ywqyzq0FHv2Es49v1UvfhLnaWnUIyAaeI\nb1ok67P/eJr9zkKQi1GyRKSXQZ9N0W+nXUbPftp7u64uZx0Ihng0QnCCDRxSTNJv\nGTp7UHzpX96dv8kPttl8/nNZkJpsA6aZVKl3qhuwiNMTH9Ye88WOQLyDAgMBAAEC\ngYAN7EN8LcyI+9TyWhSE2usl1/3qCuL1RM6PmRRGTf62Gc66c3RkIZdzURTzwj6i\nrQGvCZXR0Zq2kRR9sVNMd+6gYCgQEHrTD1Zu2lWx4paTdSUTY7or/v0M0gicJOkq\ncSZ3sQlGC0naAnK47prjKA+3UO7Hzj/Jld319Q1CnFqGoQJBAKUzP7SzMSMR4FiY\nD1lyJAeTTKwrGwnDJ0ZTb+ejUeyPt2P+Lhp6PIcD87XXxjGfCXkcFZ276XxWXpyV\nI/xDWksCQQCHNNV26qKwVTt+Ut2QJE8hBFkzVigqgyt2p51+J69YsumWKAWZpTqV\nudq4n6re/4R213E3etB7bkNLp3CGKcOpAkBvB63GdjUNPAOLp8+RL1y11rNOd745\nZndsFcH9blAubT01sG0uEH/Dws02p2omiZwlUNHabKR1k9sM5FQGRQJXAkB6WOzx\ndtFRD0+OuB2WWcTg87ZkJgqirZ+e934ksnSRpxSItB6dMk8ZPd0WRCWzNTUA9WOV\n+KS/jL+IrjO8s/5BAkAsJl0e4IdSPnpL4Av9ZBorueFw/PNeoNKSxe+MD30Wqkrf\nJHKpdOwjs80DYOq0ZXevSXXHQkTVnF6F5+Uyuwcd\n-----END RSA PRIVATE KEY-----"
    [safeLogin.tryPerson]   # 在尝试登录时错误的用户名或密码max次，该用户名在time分钟内再次登录将使用验证码进行验证
        max = 3
        time = 30
    [safeLogin.tryIp]   # 在尝试登录时ip频繁max次，该ip在time分钟内再次登录将使用验证码进行验证
        max = 3
        time = 60
`, rand.RandomString(32)),
}

var tHandlerHandlers = tplNode{
	NameFormat: "handlers/handlers.go",
	TplContent: `
package handlers

import (
	"{{.AppPkg}}/app"
	"{{.AppPkg}}/handlers/v1"
)

func init() {
	v1.SetUp()
}

func Run()  {
	app.Result.Running()
}
`,
}

var tHandlerV1 = tplNode{
	NameFormat: "handlers/v1/v1.go",
	TplContent: `
package v1

import (
	"{{.AppPkg}}/app"
)

// aliyun cdn and wechat minApp not supported PATCH, so use PUT
func SetUp() {
	api := app.Result.Api.Group("v1")
	apiAuth := api.Group("", app.Result.Jwt.Middleware)

	api.Group("account").
		// register
		POST("register-check", AccountRegisterCheck).
		POST("register-trial", AccountRegisterTrial).
		POST("register", AccountRegister).
		// login
		POST("login-ck", AccountLoginCheck).
		POST("login-vc", AccountLoginCaptcha).
		POST("login-in", AccountLoginIn).
		POST("login-out", AccountLoginOut).
		// forget
		POST("forget-check", AccountForgetCheck).
		POST("forget-trial", AccountForgetTrial).
		POST("forget", AccountForget)

	apiAuth.Group("account").
		GET("info", tbAccountsInfoOne).
		PUT("info", tbAccountsInfoUpdate).
		PUT("info-pass", tbAccountsInfoPassUpdate).
		PUT("info-avatar", tbAccountsInfoAvatar)

	{{range .Resources}}
	apiAuth.
		GET("{{.SimpleResourceName}}", {{.HandlerName}}All).
		{{if .HasId}}GET("{{.SimpleResourceName}}/:id", {{.HandlerName}}One).{{end}}
		PUT("{{.SimpleResourceName}}", {{.HandlerName}}Update).
		POST("{{.SimpleResourceName}}", {{.HandlerName}}Create){{if .HasId}}.{{end}}
		{{if .HasId}}DELETE("{{.SimpleResourceName}}/:id", {{.HandlerName}}Delete){{end}}
	{{end}}
}
`,
}

var tHandlerV1AccountLogin = tplNode{
	NameFormat: "handlers/v1/account_login.go",
	TplContent: `
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/nice/person"
	"github.com/lifegit/go-gulu/v2/pkg/gredis/cacheCount"
	"github.com/lifegit/go-gulu/v2/pkg/gredis/cacheStruct"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"github.com/mojocn/base64Captcha"
	"{{.AppPkg}}/app"
	"{{.AppPkg}}/models"
	"{{.AppPkg}}/models/column/{{.AuthTable}}"
	"image/color"
	"time"
)


var (
	accountType = "a"
	againCheck   = gin.H{"pc": false}
	againCaptcha = gin.H{"vc": false}
)

type PersonDate struct {
	Salt          string _[BACKQUOTE]_json:"s,omitempty"_[BACKQUOTE]_
	Code          string _[BACKQUOTE]_json:"c,omitempty"_[BACKQUOTE]_
	IsCaptcha     bool   _[BACKQUOTE]_json:"i,omitempty"_[BACKQUOTE]_
	CaptchaAnswer string _[BACKQUOTE]_json:"v,omitempty"_[BACKQUOTE]_
}

// check
func check(c *gin.Context) {
	// 接收参数
	var param struct {
		Username string _[BACKQUOTE]_binding:"required" json:"p"_[BACKQUOTE]_
		Time     int64  _[BACKQUOTE]_binding:"required" json:"t"_[BACKQUOTE]_
	}
	err := c.ShouldBind(&param)
	if out.HandleError(c, err) {
		return
	}

	// 判断ip是否需要验证码  || 判断用户名是否需要验证码
	client, captcha := NewIp(c.ClientIP()), NewCaptcha(param.Username)
	isCaptcha := client.IsBusy() || captcha.IsBusy()

	// 生产salt和code
	var pass person.Password
	err = pass.RandSaltAndCode(app.Global.SafeLogin.PrivateKey)
	if out.HandleError(c, err) {
		return
	}

	// 保存salt和code
	err = NewData(param.Username, pass.Code).SetStruct(PersonDate{
		Salt:      pass.Salt,
		Code:      pass.Code,
		IsCaptcha: isCaptcha,
	})
	if out.HandleError(c, err) {
		return
	}

	out.JsonData(c, gin.H{
		"c": pass.Code,        // c : string   code 随机码
		"s": pass.SaltEncrypt, // s : string   salt 盐
		"i": isCaptcha,        // i : bool     是否需要图片验证码
	})
}

// captcha
func captcha(c *gin.Context) {
	// 接收参数
	var param struct {
		Username string _[BACKQUOTE]_binding:"required" json:"p"_[BACKQUOTE]_
		Code     string _[BACKQUOTE]_binding:"required" json:"c"_[BACKQUOTE]_
	}
	err := c.ShouldBind(&param)
	if out.HandleError(c, err) {
		return
	}

	data := NewData(param.Username, param.Code)
	if !data.IsEmpty() {
		out.JsonErrorData(c, againCheck, "登录失效,请重新点击登录!")
		return
	}

	// captcha
	captchaConf := base64Captcha.DriverString{
		Length:       4,
		Height:           55,
		Width:            240,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine,
		BgColor: &color.RGBA{
			R: 45,
			G: 95,
			B: 70,
			A: 105,
		},
	}
	captchaConf.ConvertFonts()
	_, content, answer := captchaConf.GenerateIdQuestionAnswer()
	drCap, err := captchaConf.DrawCaptcha(content)
	if out.HandleError(c, err) {
		return
	}

	err = data.SetStruct(PersonDate{CaptchaAnswer: answer}, true)
	if out.HandleError(c, err) {
		return
	}

	binaryData := drCap.(*base64Captcha.ItemChar).BinaryEncoding()
	//fmt.Println("问题:",content,"答案:",answer)
	c.Data(200, "image/GIF", binaryData)
}

// login
func login(c *gin.Context) {
	// 接收参数
	var param struct {
		Username string _[BACKQUOTE]_binding:"required"_[BACKQUOTE]_
		Password string _[BACKQUOTE]_binding:"required"_[BACKQUOTE]_
		Code     string _[BACKQUOTE]_binding:"required"_[BACKQUOTE]_
		Vc       string
	}
	err := c.ShouldBind(&param)
	if out.HandleError(c, err) {
		return
	}

	// 根据username获取salt和code
	data := NewData(param.Username, param.Code)
	var res PersonDate
	err = data.GetStruct(&res)
	if out.HandleError(c, err) {
		return
	}

	// 验证salt
	if res.Salt == "" || res.Code != param.Code {
		out.JsonErrorData(c, againCheck, "登录失效,请重新点击登录")
		return
	}

	client, captcha := NewIp(c.ClientIP()), NewCaptcha(param.Username)
	// 验证验证码
	if res.IsCaptcha && param.Vc != res.CaptchaAnswer {
		out.JsonErrorData(c, againCaptcha, "验证码错误!")
		return
	}

	loginError := func() {
		captchaBool, clientBool := captcha.AddCount(), client.AddCount()
		if captchaBool || clientBool {
			err = data.SetStruct(PersonDate{IsCaptcha: true}, true)
			out.JsonErrorData(c, againCaptcha, "用户名或密码错误!")
			return
		}
		out.JsonError(c, "用户名或密码错误!")
	}

	// 取出密码
	var accountInfo models.{{.AuthTable}}
	err = app.DB.CrudOne(models.{{.AuthTable}}{Username: param.Username}, &accountInfo)
	if err != nil {
		loginError()
		return
	}
	if accountInfo.Use != {{.AuthTable}}.UseAllowed {
		out.JsonError(c, "账号已封号!请联系管理!")
		return
	}

	// 验证密码
	personPass := person.Password{
		Salt: res.Salt,
		Code: res.Code,
	}
	if err = personPass.DecryptAndCheck(accountInfo.Password, param.Password); err != nil {
		loginError()
		return
	}

	data.Destroy()
	captcha.Destroy()

	// jwt
	token, err := app.Result.Jwt.GenerateToken(app.JwtUser{
		Id:       accountInfo.ID,
		Username: accountInfo.Username,
	}, app.Global.Jwt.ExpireHour)
	if err != nil {
		out.JsonError(c, "无法登录")
		return
	}

	_ = app.DB.CrudUpdatePrimaryKey(models.{{.AuthTable}}{ID: accountInfo.ID}, models.{{.AuthTable}}{LastedAt: time.Now().Unix()})

	// authority
	//callData := gin.H{
	//	"token": data,
	//	"data": gin.H{
	//		"authority" : gin.H{
	//			"train": "get,set,add,del",
	//		},
	//	},
	//}
	out.JsonData(c, gin.H{"token": token})
}

func loginOut(c *gin.Context) {
	//	$account = models\account\Agent_sessionModel::getInstance();
	//	$account->loginOut();
	out.JsonSuccess(c)
}

func NewIp(ip string) *cacheCount.Counter {
	return cacheCount.NewCounter(
		app.Global.SafeLogin.TryIP.Max,
		cacheCount.NewCount(
			fmt.Sprintf("pi:%s", ip),
			app.Cache,
			time.Minute * time.Duration(app.Global.SafeLogin.TryIP.Time),
		),
	)
}
func NewCaptcha(username string) *cacheCount.Counter {
	return cacheCount.NewCounter(
		app.Global.SafeLogin.TryPerson.Max,
		cacheCount.NewCount(
			fmt.Sprintf("pc:%s:%s", accountType, username),
			app.Cache,
			time.Minute * time.Duration(app.Global.SafeLogin.TryPerson.Time),
		),
	)
}
func NewData(username, code string) *cacheStruct.CacheStruct {
	return cacheStruct.New(fmt.Sprintf("pd:%s:%s:%s", accountType, username, code), app.Cache)
}
`,
}

var tHandlerV1AccountInfo = tplNode{
	NameFormat: "handlers/v1/account_info.go",
	TplContent: `
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lifegit/go-gulu/v2/nice/person"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"{{.AppPkg}}/app"
	"{{.AppPkg}}/models"
	"{{.AppPkg}}/models/column/{{.AuthTable}}"
)


//One
func tbAccountsInfoOne(c *gin.Context) {
	accountJwt := c.MustGet(app.Global.Jwt.Key).(app.JwtUser)
	data := &models.{{.AuthTable}}{}
	err := app.DB.CrudOne(models.{{.AuthTable}}{ID: accountJwt.Id}, data)
	if out.HandleError(c, err) {
		return
	}

	out.JsonData(c, data)
}

// upload avatar
func tbAccountsInfoAvatar(c *gin.Context) {
	accountJwt := c.MustGet(app.Global.Jwt.Key).(app.JwtUser)
	resUpload := app.FileUploads.Upload(c, app.AttrAvatar)
	if out.HandleError(c, resUpload.Error) {
		return
	}

	err := app.DB.CrudUpdatePrimaryKey(models.{{.AuthTable}}{ID: accountJwt.Id}, models.{{.AuthTable}}{Avatar: resUpload.Url})
	if err != nil {
		_ = app.FileUploads.Remove(resUpload.Save)
		out.JsonError(c, err.Error())
		return
	}

	out.JsonData(c, gin.H{ {{.AuthTable}}.Avatar: resUpload.Url})
}

//Update
func tbAccountsInfoUpdate(c *gin.Context) {
	var mdl models.{{.AuthTable}}
	err := c.ShouldBindWith(&mdl, binding.FormPost)
	if out.HandleError(c, err) {
		return
	}

	accountJwt := c.MustGet(app.Global.Jwt.Key).(app.JwtUser)
	err = fire.NewInstance(app.DB.Omit({{.AuthTable}}.Password)).CrudUpdatePrimaryKey(models.{{.AuthTable}}{ID: accountJwt.Id}, mdl)
	if out.HandleError(c, err) {
		return
	}

	out.JsonSuccess(c)
}

//Update Pass
func tbAccountsInfoPassUpdate(c *gin.Context) {
	// 接收参数
	var param struct {
		Old string _[BACKQUOTE]_binding:"required" form:"old"_[BACKQUOTE]_
		New string _[BACKQUOTE]_binding:"required" form:"new"_[BACKQUOTE]_
	}
	err := c.ShouldBind(&param)
	if out.HandleError(c, err) {
		return
	}

	accountJwt := c.MustGet(app.Global.Jwt.Key).(app.JwtUser)
	// 取出当前密码
	var accountInfo models.{{.AuthTable}}
	err = app.DB.CrudOne(models.{{.AuthTable}}{ID: accountJwt.Id}, &accountInfo)
	if out.HandleError(c, err) {
		return
	}

	// 验证旧密码
	var pass person.Password
	if err = pass.Check(accountInfo.Password, param.Old); err != nil {
		out.JsonError(c, "旧密码不正确")
		return
	}
	// make新密码
	newPass, err := pass.MakePassword(param.New)
	if err != nil {
		out.JsonError(c, "您不能使用该新密码")
		return
	}

	// 修改密码
	err = app.DB.CrudUpdatePrimaryKey(models.{{.AuthTable}}{ID: accountInfo.ID}, models.{{.AuthTable}}{Password: newPass})
	if out.HandleError(c, err) {
		return
	}
	out.JsonSuccess(c)
}
`,
}

var tHandlerV1AccountForgetMobile = tplNode{
	NameFormat: "handlers/v1/account_forget.go",
	TplContent: `
// todo
`,
}

var tHandlerV1AccountRegisterMobile = tplNode{
	NameFormat: "handlers/v1/account_forget.go",
	TplContent: `
// todo
`,
}

var tHandlerV1AccountForgetSecurity = tplNode{
	NameFormat: "handlers/v1/account_forget.go",
	TplContent: `
// todo
`,
}

var tHandlerV1AccountRegisterSecurity = tplNode{
	NameFormat: "handlers/v1/account_forget.go",
	TplContent: `
// todo
`,
}


var tStaticIndex = tplNode{
	"static/index.html",
	`
<!DOCTYPE html><html lang="en">
<head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta name="viewport" content="width=device-width">
  <title>Golang+gin+gorm+sql created by ginplay</title>
  <body>
    <h2 style="text-align: center">
	  created by GinPlay
	  <a class="github-button" href="https://github.com/lifegit/go-gulu" data-show-count="true" aria-label="Star lifegit/go-gulu on GitHub">Star</a>
  </h2>
  <script async defer src="https://buttons.github.io/buttons.js"></script>
</body>
</html>
`,
}

var tTaskCore = tplNode{
	NameFormat: "tasks/core.go",
	TplContent: `
package tasks

import "github.com/lifegit/go-gulu/v2/nice/core"

func RunTasks() {
	tasks := core.NewScheduler()
	// tasks.Every(1).Seconds().Do(task)
	tasks.Start()
}


`,
}
var tTaskExample = tplNode{
	NameFormat: "tasks/example.go",
	TplContent: `
package tasks

import "fmt"

//defining schedule task function here
//then add the function in manger.go
func task() {
	fmt.Println("task one is called")
}
`,
}

var tGitIgnore = tplNode{
	NameFormat: ".gitignore",
	TplContent: `
main
*.exe
*.log
conf/*
!conf/dev
!conf/base.toml
.idea/*
.idea
.vscode/*
.vscode
runtime/*
{{.AppPkg}}
{{.AppPkg}}.test
`,
}
var tDockerfile = tplNode{
	NameFormat: "Dockerfile",
	TplContent: `
# =============== build and run ===============
# build:  docker build -t spider-data .
# run:    docker run hello-world


# =============== build stage ===============
FROM golang:1.16.5-buster AS build
# env
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
# dependent
WORKDIR /app
COPY go.* ./
RUN go mod download -x all
# build
COPY . ./
# ldflags:
#  -s: disable symbol table
#  -w: disable DWARF generation
# run "go tool link -help" to get the full list of ldflags
RUN go env && go build -ldflags "-s -w" -o spider-data -v ./main.go



# =============== final stage ===============
FROM alpine:latest AS final
# resources
WORKDIR /app
COPY --from=build /app/spider-data ./
COPY --from=build /app/conf/base.toml ./conf/base.toml
COPY --from=build /app/conf/prod ./conf/prod
EXPOSE 8881
ENTRYPOINT ["env","GO_ENV=prod","/app/spider-data", "-other", "flags"]
`,
}
var tMod = tplNode{
	NameFormat: "go.mod",
	TplContent: `
module {{.AppPkg}}

go 1.16

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-contrib/static v0.0.1
	github.com/gin-gonic/autotls v0.0.3
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/lifegit/go-gulu/v2 v2.0.0-incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.11
)

replace github.com/lifegit/go-gulu/v2 => /Users/yxs/GolandProjects/src/go-gulu

`,
}

//var tMod = tplNode{
//	NameFormat: "go.mod",
//	TplContent: `module {{.AppPkg}}
//
//go 1.16
//`,
//}
var tLICENSE = tplNode{
	NameFormat: "LICENSE",
	TplContent: `
The MIT License (MIT)

Copyright (c) The lifegit Authors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

`,
}
var tMain = tplNode{
	NameFormat: "main.go",
	TplContent: `
package main

import (
	"{{.AppPkg}}/app"
	"{{.AppPkg}}/handlers"
	"{{.AppPkg}}/tasks"
)

// @title GinPlay Example API
// @version 1.0.0
// @description This is a sample Server ginplay
// @securityDefinitions.apikey ApiKeyAuth
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /services
func main() {
	defer app.Close()

	app.Log.Infof("app: %s , version: %d at start ...", app.Global.App.Name, app.Global.App.Version)
	if app.Global.App.IsCron {
		go tasks.RunTasks()
	}
	handlers.Run()
}
`,
}
var tReadMe = tplNode{
	NameFormat: "README.md",
	TplContent: `

# 介绍

本项目基于[ginplay](https://github.com/lifegit/go-gulu/tree/master/pkg/ginplay)脚手架生成的 RESTful Api 项目

## 运行
_[BACKQUOTE]__[BACKQUOTE]__[BACKQUOTE]_
	$ set env GOPROXY=https://goproxy.io
	$ go mod tidy
	$ go run
	$ go build
_[BACKQUOTE]__[BACKQUOTE]__[BACKQUOTE]_

### 项目结构
    project
    |-- app                                       # 应用的一些工具
    |   |-- app                                   # 对工具集合的初始化与销毁等操作
    |   |-- basics                                # 基础初始化。包含 logrus 实现整个应用的 log
    |   |-- cache                                 # redis，使用 go-redis 实现缓存能力
    |   |-- conf                                  # 配置文件，使用 viper、fsnotify 实现 toml 格式的配置文件读取
    |   |-- db                                    # db，使用 gorm 对数据的持久化
    |   |-- result                                # ResultApi, 使用 gin 初始化、路由 等。
	|   |-- sms                                	  # sms, 短信通道
    |   |-- upload                                # upload，支持文件上传到本地或阿里云oss。
    |-- conf                                      # 配置文件
    |-- docs                                      # 文档
    |   |-- openapi                               # openApi文档
    |   |   |-- v2                                # openApi v2 文档
    |   |   |-- v3                                # openApi v3 文档
    |-- handlers                                  # 路由逻辑处理
    |   |-- v1                                    # 接口
    |-- models                                    # 应用数据库模型
    |-- static                                    # 静态文件
    |-- tasks                                     # 定时器
 
## 用法
- [swagger DOC](http://{{.AppAddr}}:{{.AppPort}}/swagger/index.html)_[BACKQUOTE]_http://{{.AppAddr}}:{{.AppPort}}/swagger/index.html_[BACKQUOTE]_
- [static](http://{{.AppAddr}}:{{.AppPort}})_[BACKQUOTE]_http://{{.AppAddr}}:{{.AppPort}}_[BACKQUOTE]_
- [App INFO](http://{{.AppAddr}}:{{.AppPort}}/app/info)_[BACKQUOTE]_http://{{.AppAddr}}:{{.AppPort}}/app/info_[BACKQUOTE]_ //todo
- [API baseURL](http://{{.AppAddr}}:{{.AppPort}}/api/v1)

## 🌟生成swagger文档
更多见: [swagr](https://github.com/lifegit/go-gulu/pkg/swagr)
_[BACKQUOTE]_[BACKQUOTE]_[BACKQUOTE]bash
	# go modules 方式
	$ go env -w GO111MODULE=on
	$ go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
	$ go get -u github.com/swaggo/swag/cmd/swag
	$ swagr init
	
	# go get 方式 (国内可能不好使)
	$ go get https://github.com/lifegit/go-gulu/pkg/swagr
	$ swagr init
_[BACKQUOTE]_[BACKQUOTE]_[BACKQUOTE]
执行上述命令后，会生成docs/swagger文件夹。存在swaggerV2与openapi3文档。




## 感谢
- [gin](https://github.com/gin-gonic/gin)
- [GORM](http://gorm.io/)、[fire](https://github.com/lifegit/go-gulu/tree/master/pkg/fire)
- [core](https://github.com/lifegit/go-gulu/tree/master/nice/core)
- [viper](https://github.com/spf13/viper)、[fsnotify](https://github.com/fsnotify/fsnotify
- [logrus](https://github.com/sirupsen/logrus)
- [go-redis](https://github.com/go-redis/redis)
- [base64Captcha](https://github.com/mojocn/base64Captcha)
- [go-gulu](https://github.com/lifegit/go-gulu)
`,
}

var ParseOneList = []tplNode{
	tAppApp, tAppBasics, tAppCache, tAppConf, tAppDb, tAppResult, tAppSMS, tAppUpload,
	tConfigBaseToml, tConfigDevToml,
	tHandlerHandlers,
	tHandlerV1, tHandlerV1AccountLogin, tHandlerV1AccountInfo,
	tStaticIndex,
	tTaskCore, tTaskExample,
	tGitIgnore, tDockerfile, tMod, tLICENSE, tMain, tReadMe,
}

var ParseTypeMobileList = []tplNode{
	tHandlerV1AccountRegisterMobile, tHandlerV1AccountForgetMobile,
}

var ParseTypeSecurityList = []tplNode{
	tHandlerV1AccountRegisterSecurity, tHandlerV1AccountForgetSecurity,
}


var tModelObj = tplNode{
	NameFormat: "models/%s.go",
	TplContent: `
package models

// todo 新的

const (
	{{.ModelName}}Table = "{{.TableName}}"
	{{range $i,$p := .Properties}}{{range $z := .ColumnCommentTag}}
	{{$.ModelName}}{{$p.ModelProp}}{{.Tag}}	= {{.State}}{{if ne .Remarks ""}} // {{.Remarks}}{{end}}{{end}}{{end}}
	{{range .Properties}}
	{{$.ModelName}}{{.ModelProp}}	= "{{.ColumnName}}"{{end}}	
)

//{{.ModelName}}
type {{.ModelName}} struct {
	{{range .Properties}}
	 {{.ModelProp}}      {{.ModelType}}         _[BACKQUOTE]_gormCreate:"required" {{.ModelTag}}_[BACKQUOTE]_{{end}}
}
`,
}

var tHandlersObj = tplNode{
	NameFormat: "handlers/v1/%s.go",
	TplContent: `
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"{{.AppPkg}}/app"
	"{{.AppPkg}}/models"
)

{{if .HasId}}
// @Tags {{.SimpleResourceName}}
// @Summary find one {{.SimpleResourceName}}
// @accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} out.Result{data=models.{{.ModelName}}} {{.SimpleResourceName}}} ""
// @Router /v1/{{.SimpleResourceName}}/:id [get]
func {{.HandlerName}}One(c *gin.Context) {
	id, err := out.ParseParamID(c)
	if out.HandleError(c, err) {
		return
	}

	data := &models.{{.ModelName}}{}
	err = app.DB.CrudOne(models.{{.ModelName}}{ID: id}, data)
	if out.HandleError(c, err) {
		return
	}
	out.JsonData(c, data)
}
{{end}}
//Create
// @Tags {{.SimpleResourceName}}
// @Summary find one {{.SimpleResourceName}}
// @Accept application/json
// @Produce application/json
// @Param id body {{.models.ModelName}} true "model"
// @Success 200 {object} out.Result{data=models.{{.ModelName}}} {{.SimpleResourceName}}}} ""
// @Router /v1/card/:id [get]
func {{.HandlerName}}Create(c *gin.Context) {
	var mdl models.{{.ModelName}}
	err := c.ShouldBind(&mdl)
	if out.HandleError(c, err) {
		return
	}
	err = app.DB.CrudCreate(mdl)
	if out.HandleError(c, err) {
		return
	}
	out.JsonData(c, mdl)
}
{{if .HasId}}
//Update
func {{.HandlerName}}Update(c *gin.Context) {
	var mdl models.{{.ModelName}}
	err := c.ShouldBindWith(&mdl, binding.FormPost)
	if out.HandleError(c, err) {
		return
	}
	err = app.DB.CrudUpdatePrimaryKey(models.{{.ModelName}}{ID: mdl.ID}, mdl)
	if out.HandleError(c, err) {
		return
	}

	out.JsonSuccess(c)
}
{{end}}
//All
func {{.HandlerName}}All(c *gin.Context) {
	var param fire.PageParam
	err := c.ShouldBind(&param)
	if out.HandleError(c, err) {
		return
	}
	pageResult, err :=  app.DB.CrudAllPage(models.{{.ModelName}}{}, &[]models.{{.ModelName}}{}, param.Page)
	if out.HandleError(c, err) {
		return
	}

	out.JsonPagination(c, pageResult)
}
{{if .HasId}}
//Delete
func {{.HandlerName}}Delete(c *gin.Context) {
	id, err := out.ParseParamID(c)
	if out.HandleError(c, err) {
		return
	}

	err = app.DB.CrudDelete(models.{{.ModelName}}{ID: id})
	if out.HandleError(c, err) {
		return
	}
	out.JsonSuccess(c)
}
{{end}}
`,
}

var ParseObjList = []tplNode{tModelObj, tHandlersObj}

// todo swagger Doc
