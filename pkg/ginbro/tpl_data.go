package ginbro

import "strings"

var tDocYaml = tplNode{
	NameFormat: "doc/swagger.yaml",
	TplContent: `swagger: "2.0"
info:
  description: "A GinBro RESTful APIs"
  version: "1.0.0"
  title: "GinBro RESTful APIs Application"
host: "{{.AppAddr}}"
basePath: "/api/v1"

schemes:
- "http"
paths:
  {{range .Resources}}
  {{if .IsAuthTable}}
  /login:
    post:
      tags:
      - "auth"
      summary: "login by {{.ResourceName}}"
      description: "login by {{.ResourceName}}"
      consumes:
      - "application/json"
      produces:
      - "application/json"
      parameters:
      - in: "body"
        name: "body"
        description: "create {{.ResourceName}}"
        required: true
        schema:
          $ref: "#/definitions/{{.ModelName}}"
      responses:
        200:
          description: "successful operation"
          schema:
            $ref: "#/definitions/{{.ModelName}}Auth"

  {{end}}
  /{{.ResourceName}}:
    get:
      tags:
      - "{{.ResourceName}}"
      summary: "get all {{.ResourceName}} by pagination"
      description: ""
      consumes:
      - "application/json"
      produces:
      - "application/json"
      parameters:
      - name: "where"
        in: "query"
        description: "column:value will use sql LIKE for search eg:id:67 will where id >67 eg2: name:eric => where name LIKE '%eric%'"
        required: false
        type: "array"
        items:
          type: "string"
      - name: "fields"
        in: "query"
        description: "{$tableColumn},{$tableColumn}... "
        required: false
        type: "string"
      - name: "order"
        in: "query"
        description: "eg: id desc, name desc"
        required: false
        type: "string"
      - name: "offset"
        in: "query"
        description: "sql offset eg: 10"
        required: false
        type: "integer"
      - name: "limit"
        in: "query"
        default: "2"
        description: "limit returning object count"
        required: false
        type: "integer"

      responses:
        200:
          description: "successful operation"
          schema:
            type: "array"
            items:
              $ref: "#/definitions/{{.ModelName}}Pagination"
    post:
      tags:
      - "{{.ResourceName}}"
      summary: "create {{.ResourceName}}"
      description: "create {{.ResourceName}}"
      consumes:
      - "application/json"
      produces:
      - "application/json"
      parameters:
      - in: "body"
        name: "body"
        description: "create {{.ResourceName}}"
        required: true
        schema:
          $ref: "#/definitions/{{.ModelName}}"
      responses:
        200:
          description: "successful operation"
          schema:
            $ref: "#/definitions/ApiResponse"

    patch:
      tags:
      - "{{.ResourceName}}"
      summary: "update {{.ResourceName}}"
      description: "update {{.ResourceName}}"
      consumes:
      - "application/json"
      produces:
      - "application/json"
      parameters:
      - in: "body"
        name: "body"
        description: "create {{.ResourceName}}"
        required: true
        schema:
          $ref: "#/definitions/{{.ModelName}}"
      responses:
        200:
          description: "successful operation"
          schema:
            $ref: "#/definitions/ApiResponse"

  /{{.ResourceName}}/{ID}:
    get:
      tags:
      - "{{.ResourceName}}"
      summary: "get a {{.ResourceName}} by ID"
      description: "get a {{.ResourceName}} by ID"
      consumes:
      - "application/json"
      produces:
      - "application/json"
      parameters:
      - name: "ID"
        in: "path"
        description: "ID of {{.ResourceName}} to return"
        required: true
        type: "integer"
        format: "int64"
      responses:
        200:
          description: "successful operation"
          schema:
            $ref: "#/definitions/{{.ModelName}}"
    delete:
      tags:
      - "{{.ResourceName}}"
      summary: "Destroy a {{.ResourceName}} by ID"
      description: "delete a {{.ResourceName}} by ID"
      consumes:
      - "application/json"
      produces:
      - "application/json"
      parameters:
      - name: "ID"
        in: "path"
        description: "ID of {{.ResourceName}} to delete"
        required: true
        type: "integer"
        format: "int64"
      responses:
        200:
          description: "successful operation"
          schema:
            $ref: "#/definitions/ApiResponse"
  {{end}}


definitions:
  {{range $table := .Resources}}
  {{ if $table.IsAuthTable }}
  {{$table.ModelName}}Auth:
    type: "object"
    properties:
      {{range $row := $table.Properties}}
      {{$row.ColumnName}}:
        type: "{{$row.SwaggerType}}"
        description: "{{$row.ColumnComment}}"
        format: "{{$row.SwaggerFormat}}"
        {{end}}
      token:
        type: "string"
        description: "jwt token"
        format: "string"
      expire:
        type: "string"
        description: "jwt token expire time"
        format: "date-time"
      expire_ts:
        type: "integer"
        description: "expire timestamp unix"
        format: "int64"
  {{end}}
  {{$table.ModelName}}:
    type: "object"
    properties:
    {{range $row := $table.Properties}}
      {{$row.ColumnName}}:
        type: "{{$row.SwaggerType}}"
        description: "{{$row.ColumnComment}}"
        format: "{{$row.SwaggerFormat}}"
      {{end}}
  {{$table.ModelName}}Pagination:
    type: "object"
    properties:
      code:
        type: "integer"
        description: "json repose code"
        format: "int32"
      total:
        type: "integer"
        description: "total numbers"
        format: "int32"
      offset:
        type: "integer"
        description: "offset"
        format: "int32"
      limit:
        type: "integer"
        description: "limit"
        format: "int32"
      list:
        type: "array"
        items:
          $ref: "#/definitions/{{$table.ModelName}}"
{{end}}
  ApiResponse:
    type: "object"
    properties:
      code:
        type: "integer"
        format: "int32"
      msg:
        type: "string"
externalDocs:
  description: "Find out more about Swagger"
  url: "http://swagger.io"
`,
}
var tReadme = tplNode{
	NameFormat: "README_ZH.md",
	TplContent: `
# A RESTful APIs

## 推荐Go版本>1.12
- 对于中国用户: set env GOPROXY=https://goproxy.io
- 运行: go tidy
    
## 用法
- [swagger DOC ](http://{{.AppAddr}}/doc)_[BACKQUOTE]_http://{{.AppAddr}}/swagger/_[BACKQUOTE]_
- [static ](http://{{.AppAddr}})_[BACKQUOTE]_http://{{.AppAddr}}_[BACKQUOTE]_
- [GinbroApp INFO ](http://{{.AppAddr}}/GinbroApp/info)_[BACKQUOTE]_http://{{.AppAddr}}/GinbroApp/info_[BACKQUOTE]_
- API baseURL : _[BACKQUOTE]_http://{{.AppAddr}}/api/v1_[BACKQUOTE]_

## 目录
- app: 一些常用工具
- conf：用于存储配置文件
- docs：文档
- handlers：路由逻辑处理
- models：应用数据库模型
- pkg：第三方包
- runtime 应用运行时数据
- static 默认地址
- tasks：任务


## 感谢
- [felix](https://github.com/dejavuzhou/felix)
- [base64Captcha](https://github.com/mojocn/base64Captcha)
- [swagger Specification](https://swagger.io/specification/)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [GORM](http://gorm.io/)
- [viper](https://github.com/spf13/viper)
- [cobra](https://github.com/spf13/cobra#getting-started)
- [go-redis](https://github.com/go-redis/redis)
`,
}

var tModelObj = tplNode{
	NameFormat: "models/m_%s.go",
	TplContent: `
package models

import (
	"errors"
)

//{{.ModelName}}
type {{.ModelName}} struct {
	DbUtils		*dbUtils.DbUtils	_[BACKQUOTE]_gorm:"-" json:"-"_[BACKQUOTE]_
	{{range .Properties}}
	 {{.ModelProp}}      {{.ModelType}}         _[BACKQUOTE]_bindingCreate:"required" {{.ModelTag}}_[BACKQUOTE]_{{end}}
}
`,
}

var tColumnObj = tplNode{
	NameFormat: "models/fields/%s/%s.go",
	TplContent: `
package {{.TableName}}

const TableName = "{{.TableName}}"
{{range $i,$p := .Properties}}{{range $z := .ColumnCommentTag}}
const {{$p.ModelProp}}{{.Tag}}	= {{.State}}{{if ne .Remarks ""}} // {{.Remarks}}{{end}}{{end}}
{{end}}
{{range .Properties}}
const {{.ModelProp}}	= "{{.ColumnName}}"{{end}}
`,
}

var tHandlersObj = tplNode{
	NameFormat: "handlers/h_%s.go",
	TplContent: `
package handlers

import (
	"{{.AppPkg}}/models"
	"{{.AppPkg}}/models/fields/{{.TableName}}"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func init() {
	result.Api.GET("{{.SimpleResourceName}}",{{if .IsAuthTable}}result.Jwt.Middleware,{{end}} {{.HandlerName}}All)
	{{if .HasId}}result.Api.GET("{{.SimpleResourceName}}/:id", {{if .IsAuthTable}}result.Jwt.Middleware,{{end}} {{.HandlerName}}One){{end}}
	result.Api.POST("{{.SimpleResourceName}}", {{if .IsAuthTable}}result.Jwt.Middleware,{{end}} {{.HandlerName}}Create)
	result.Api.POST("{{.SimpleResourceName}}-update", {{if .IsAuthTable}}result.Jwt.Middleware,{{end}} {{.HandlerName}}Update) // aliyun cdn not supported PATCH
	{{if .HasId}}result.Api.DELETE("{{.SimpleResourceName}}/:id", {{if .IsAuthTable}}result.Jwt.Middleware,{{end}} {{.HandlerName}}Delete){{end}}
}

{{if .HasId}}
//One
func {{.HandlerName}}One(c *gin.Context) {
	id, err := out.ParseParamID(c)
	if out.HandleError(c, err) {
		return
	}

	data := &{{.ModelName}}{}
	err := app.DB.CrudOne({{.ModelName}}{Id: id}, data)
	if out.HandleError(c, err) {
		return
	}
	app.JsonData(c, data)
}
{{end}}
//Create
func {{.HandlerName}}Create(c *gin.Context) {
	var mdl models.{{.ModelName}}
	err := c.ShouldBind(&mdl)
	if out.HandleError(c, err) {
		return
	}
//todo 验证参数全不全
	err = app.DB.Create(mdl).Error
	if out.HandleError(c, err) {
		return
	}
	app.JsonData(c, mdl)
}
//Update
func {{.HandlerName}}Update(c *gin.Context) {
	// limit 1 和 update id
	var mdl models.{{.ModelName}}
	err := c.ShouldBindWith(&mdl, binding.FormPost)
	if out.HandleError(c, err) {
		return
	}
	err = app.DB.CrudUpdate({{.ModelName}}{ID: mdl.ID}, mdl).Error
	if out.HandleError(c, err) {
		return
	}

	app.JsonSuccess(c)
}

//All
func {{.HandlerName}}All(c *gin.Context) {
	var param fire.PageParam
	err := c.ShouldBind(&param)
	if out.HandleError(c, err) {
		return
	}
	res := &[]models.{{.ModelName}}{}
	pageResult, err :=  app.DB.CrudAllPage(models.{{.ModelName}}{}, res, param.Page)
	if out.HandleError(c, err) {
		return
	}

	app.JsonPagination(c, list, pageResult)
}
{{if .HasId}}
//Delete
func {{.HandlerName}}Delete(c *gin.Context) {
	id, err := out.ParseParamID(c)
	if out.HandleError(c, err) {
		return
	}

	err = app.DB.CrudDelete(models.{{.ModelName}}{Id: id})
	if out.HandleError(c, err) {
		return
	}
	app.JsonSuccess(c)
}
{{end}}
`,
}

var tStaticIndex = tplNode{
	"static/index.html",
	`
<!DOCTYPE html><html lang="en">
<head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta name="viewport" content="width=device-width">
  <title>Golang+gin+gorm+sql created by felix ginbro</title>
<body>
  <h2 style="text-align: center">put front end files into this folder</h2>
</body>
</html>
`,
}


var tApp = tplNode{
	NameFormat: "app/app.go",
	TplContent: `
package app

import (
	"context"
)

//ServerRun start the server
func ServerRun() {
	Log.Infof("app: %s , version: %d at start ...", Global.App.Name, Global.App.Version)

	Result.Running()
}

//Close server app
func ServerClose() {
	_ = Cache.Close()
	_ = DB.Close()
}
`,
}

var tAppDb = tplNode{
	NameFormat: "app/db.go",
	TplContent: `
package app

import (
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

// 测试的开始位置
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
	dbDSN := os.Getenv("GORM_DSN")
	switch os.Getenv("GORM_DIALECT") {
	case "mysql":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
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
var tHandlerApp = tplNode{
	NameFormat: "app/app.go",
	TplContent: `
package app

import (
	"context"
)

func init() {
	SetUpConf()
	SetUpBasics()

	SetUpCache()
	SetUpDB()

	SetUpResult()
}

//ServerRun start the server
func ServerRun() {
	Log.Infof("app: %s , version: %d at start ...", Global.App.Name, Global.App.Version)

	Result.Running()
}

//Close server app
func ServerClose() {
	_ = Cache.Close()
	_ = DB.Close(context.Background())
}
`,
}

var tHandlerGin = tplNode{
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
	"net/http"
	"path"
	"time"
)

var Result ResultApi

type ResultApi struct {
	Gin *gin.Engine
	Api *gin.RouterGroup
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
	// middlewareRecovery
	r.Gin.Use(gin.Recovery())
	// middlewareCors
	if Global.Server.IsCors {
		r.Gin.Use(mwCors.NewCorsMiddleware())
	}
	// staticPath
	if staticPath := Global.Server.StaticPath; staticPath != "" {
		Log.Info(fmt.Sprintf("visit http://%s/ for front-end static html files", addr))
		r.Gin.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	}
	// notRoute
	if Global.Server.IsNotRoute {
		Log.Info(fmt.Sprintf("visit http://%s/docs for RESTful Is NotRoute", addr))
		r.Gin.NoRoute(func(c *gin.Context) {
			//file := path.Join(sp, "index.html")
			//c.File(file)
			c.Status(http.StatusNotFound)
		})
	}
	// swaggerApi
	if Global.Server.IsSwagger && Global.isDev() {
		Log.Info(fmt.Sprintf("visit http://%s/docs for RESTful APIs Document", addr))
		//add edit your own swagger.docs.yml file in ./swagger/docs.yml
		//generateSwaggerDocJson()
		r.Gin.Static("docs", "./docs")
	}
	// appInfo
	if Global.isDev() {
		Log.Info(fmt.Sprintf("visit http://%s/app/info for app info only on not-prod mode", addr))
		r.Gin.GET("/app/info", func(c *gin.Context) {
			c.JSON(200, Global)
		})
	}

	// apiPrefix
	r.Api = r.Gin.Group(path.Join("services", Global.Server.APIPrefix))
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
func (r *ResultApi) GetJwtUser(c *gin.Context) (user *JwtUser, err error) {
	//return &utils.User{
	//	Id:1,
	//	Username:"12345678",
	//},nil

	res, exists := c.Get(conf.GetString("jwt.key"))
	if !exists || res == nil {
		return nil, errors.New("参数不存在")
	}

	if err := mapstructure.Decode(res, &user); err != nil {
		return nil, errors.New("参数不存在")
	}

	return
}
`,
}

var tMain = tplNode{
	NameFormat: "main.go",
	TplContent: `
package main

import (
	"{{.AppPkg}}/app"
	"{{.AppPkg}}/tasks"
)

func main() {
	if app.Global.App.IsCron {
		go tasks.RunTasks()
	}
	defer app.ServerClose()

	app.ServerRun()
}
`,
}

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

var tHandlerLogin = tplNode{
	NameFormat: "handlers/h_login.go",
	TplContent: `
package handlers

import (
	"{{.AppPkg}}/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.POST("login", login)
}

func login(c *gin.Context) {
	var mdl models.{{.ModelName }}
	err := c.ShouldBind(&mdl)
	if out.HandleError(c, err) {
		return
	}
	ip := c.ClientIP()
	data, err := mdl.Login(ip)
	if out.HandleError(c, err) {
		return
	}
	jsonData(c, data)
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
.idea/*
.idea
.vscode/*
.vscode
runtime/*
{{.AppPkg}}
{{.AppPkg}}.test
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

var tDocOauth2 = tplNode{
	"doc/oauth-redirect.html",
	`
<!doctype html>
<html lang="en-US">
<body onload="run()">
</body>
</html>
<script>
    'use strict';

    function run() {
        var oauth2 = window.opener.swaggerUIRedirectOauth2;
        var sentState = oauth2.state;
        var redirectUrl = oauth2.redirectUrl;
        var isValid, qp, arr;

        if (/code|token|error/.test(window.location.hash)) {
            qp = window.location.hash.substring(1);
        } else {
            qp = location.search.substring(1);
        }

        arr = qp.split("&")
        arr.forEach(function (v, i, _arr) {
            _arr[i] = '"' + v.replace('=', '":"') + '"';
        })
        qp = qp ? JSON.parse('{' + arr.join() + '}',
                function (key, value) {
                    return key === "" ? value : decodeURIComponent(value)
                }
        ) : {}

        isValid = qp.state === sentState

        if ((
                oauth2.auth.schema.get("flow") === "accessCode" ||
                oauth2.auth.schema.get("flow") === "authorizationCode"
        ) && !oauth2.auth.code) {
            if (!isValid) {
                oauth2.errCb({
                    authId: oauth2.auth.name,
                    source: "auth",
                    level: "warning",
                    message: "Authorization may be unsafe, passed state was changed in server Passed state wasn't returned from auth server"
                });
            }

            if (qp.code) {
                delete oauth2.state;
                oauth2.auth.code = qp.code;
                oauth2.callback({auth: oauth2.auth, redirectUrl: redirectUrl});
            } else {
                let oauthErrorMsg
                if (qp.error) {
                    oauthErrorMsg = "[" + qp.error + "]: " +
                            (qp.error_description ? qp.error_description + ". " : "no accessCode received from the server. ") +
                            (qp.error_uri ? "More info: " + qp.error_uri : "");
                }

                oauth2.errCb({
                    authId: oauth2.auth.name,
                    source: "auth",
                    level: "error",
                    message: oauthErrorMsg || "[Authorization failed]: no accessCode received from the server"
                });
            }
        } else {
            oauth2.callback({auth: oauth2.auth, token: qp, isValid: isValid, redirectUrl: redirectUrl});
        }
        window.close();
    }
</script>

`,
}
var tDocIndex = tplNode{
	"doc/index.html",
	`
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="//unpkg.com/swagger-ui-dist@3/swagger-ui.css">
    <link rel="icon" type="image/png" href="//unpkg.com/swagger-ui-dist@3/favicon-32x32.png" sizes="32x32"/>
    <link rel="icon" type="image/png" href="//unpkg.com/swagger-ui-dist@3/favicon-16x16.png" sizes="16x16"/>
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }

        *,
        *:before,
        *:after {
            box-sizing: inherit;
        }

        body {
            margin: 0;
            background: #fafafa;
        }
    </style>
</head>

<body>
<div id="swagger-ui"></div>

<script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
<script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-standalone-preset.js"></script>
<script>
    window.onload = function () {
        // Build a system
        const ui = SwaggerUIBundle({
            url: "swagger.yaml",
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout"
        })
        window.ui = ui
    }
</script>
</body>
</html>

`,
}

var tAppLogging = tplNode{
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



var tConf = tplNode{
	NameFormat: "app/conf.go",
	TplContent: strings.ReplaceAll(`
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
		Name     string [']toml:"name"[']
		Version  int    [']toml:"version"[']
		TimeZone string [']toml:"timeZone"[']
		Env      string [']toml:"env"[']
		IsCron     bool   [']toml:"cron"[']
		Log      string [']toml:"log"[']
	} [']toml:"app"[']
	Server struct {
		Addr         string [']toml:"addr"[']
		Port         int    [']toml:"port"[']
		RunMode      string [']toml:"runMode"[']
		ReadTimeout  int    [']toml:"readTimeout"[']
		WriteTimeout int    [']toml:"writeTimeout"[']
		APIPrefix    string [']toml:"apiPrefix"[']
		StaticPath   string [']toml:"staticPath"[']
		LogDir       string [']toml:"logDir"[']
		IsNotRoute     bool   [']toml:"notRoute"[']
		IsSwagger      bool   [']toml:"swagger"[']
		IsCors         bool   [']toml:"cors"[']
		IsHTTPS        bool   [']toml:"https"[']
		IsHTTPLog      bool   [']toml:"httpLog"[']
	} [']toml:"server"[']
	Jwt struct {
		Secret     string [']toml:"secret"[']
		Key        string [']toml:"key"[']
		ExpireHour int    [']toml:"expireHour"[']
	} [']toml:"jwt"[']
	Upload struct {
		Type           string [']toml:"type"[']
		ImageSavePath  string [']toml:"imageSavePath"[']
		ImageAllowExts string [']toml:"imageAllowExts"[']
		Domain         string [']toml:"domain"[']
		Local          struct {
			BaseDir string [']toml:"baseDir"[']
		} [']toml:"local"[']
		TypeOss struct {
			AccessKeyID     string [']toml:"accessKeyID"[']
			AccessKeySecret string [']toml:"accessKeySecret"[']
			Endpoint        string [']toml:"endpoint"[']
			BucketName      string [']toml:"bucketName"[']
		} [']toml:"typeOss"[']
	} [']toml:"upload"[']
	Db struct {
		Type     string [']toml:"type"[']
		Addr     string [']toml:"addr"[']
		Port     int    [']toml:"port"[']
		Username string [']toml:"username"[']
		Password string [']toml:"password"[']
		Database string [']toml:"database"[']
		Charset  string [']toml:"charset"[']
	} [']toml:"db"[']
	Redis struct {
		Addr     string [']toml:"addr"[']
		Port     int    [']toml:"port"[']
		Password string [']toml:"password"[']
		DbIdx    int    [']toml:"dbIdx"[']
	} [']toml:"redis"[']
	SafeLogin struct {
		PrivateKey string [']toml:"privateKey"[']
		TryPerson  struct {
			Max  int [']toml:"max"[']
			Time int [']toml:"time"[']
		} [']toml:"tryPerson"[']
		TryIP struct {
			Max  int [']toml:"max"[']
			Time int [']toml:"time"[']
		} [']toml:"tryIp"[']
	} [']toml:"safeLogin"[']
}

const DEV = "dev"

func (g *GlobalConf) isDev() bool {
	res := os.Getenv("GO_ENV")
	return res == "" || res == DEV
}
func (g *GlobalConf) getEnv() string {
	res := os.Getenv("GO_ENV")
	if res == "" {
		return DEV
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
}
`,"[']", "`"),
}


var tFileManager = tplNode{
	NameFormat: "app/fileManager.go",
	TplContent: `
package app
// todo

//import (
//	"{{.AppPkg}}/cite/conf"
//	"errors"
//	"github.com/gin-gonic/gin"
//	"github.com/sirupsen/logrus"
//	"github.com/lifegit/go-gulu/upload"
//	"github.com/lifegit/go-gulu/upload/local"
//	"github.com/lifegit/go-gulu/upload/oss"
//	"strings"
//)
//
//const uploadTypeOSS = "oss"
//const uploadTypeLocal = "local"
//
//var uploadType = conf.GetString("upload.type")
//var filesExtsImage = strings.Split(conf.GetString("upload.imageAllowExts"), ",")
//
//var uploadOss *oss.Oss
//var uploadLocal *local.Local
//
//func init() {
//	if uploadType == uploadTypeOSS {
//		c, err := oss.New(conf.GetString("upload.typeOss.endpoint"), conf.GetString("upload.typeOss.accessKeyID"), conf.GetString("upload.typeOss.accessKeySecret"), conf.GetString("upload.typeOss.bucketName"), conf.GetString("upload.domain"))
//		if err != nil {
//			err = errors.New("could not check to the oss server")
//			logrus.WithError(err).Fatal(err)
//			return
//		}
//		uploadOss = c
//	} else if uploadType == uploadTypeLocal {
//		uploadLocal = local.New(conf.GetString("upload.local.baseDir"), conf.GetString("upload.domain"))
//	} else {
//		err := errors.New("upload type is nil")
//		logrus.WithError(err).Fatal(err)
//		return
//	}
//}
//
//func FileUploadAvatar(c *gin.Context, identity string) *upload.File {
//	return fileUpload(c, &upload.FileAttribute{
//		Exts:    filesExtsImage,
//		DirPath: conf.GetString("upload.imageSavePath") + "avatars/" + identity + "/",
//		MaxSize: 3145728,
//	})
//}
//
//func FileUploadRecoveryAvatar(c *gin.Context, recoveryId string) *upload.File {
//	return fileUpload(c, &upload.FileAttribute{
//		Key:     "avatar",
//		Exts:    filesExtsImage,
//		DirPath: conf.GetString("upload.imageSavePath") + "recovery/avatar/" + recoveryId + "/",
//		MaxSize: 1048576,
//	})
//}
//func FileUploadProductAvatar(c *gin.Context, productId string) *upload.File {
//	return fileUpload(c, &upload.FileAttribute{
//		Key:     "avatar",
//		Exts:    filesExtsImage,
//		DirPath: conf.GetString("upload.imageSavePath") + "product/avatar/" + productId + "/",
//		MaxSize: 1048576,
//	})
//}
//
//func FileUploadRecoveryImg(c *gin.Context, recoveryId string) *upload.File {
//	return fileUpload(c, &upload.FileAttribute{
//		Key:     "img",
//		Exts:    filesExtsImage,
//		DirPath: conf.GetString("upload.imageSavePath") + "recovery/img/" + recoveryId + "/",
//		MaxSize: 2097152,
//	})
//}
//
//func FileUploadResources(c *gin.Context) *upload.File {
//	return fileUpload(c, &upload.FileAttribute{
//		Key:     "file",
//		Exts:    filesExtsImage,
//		DirPath: conf.GetString("upload.imageSavePath") + "resources/",
//		MaxSize: 2097152,
//	})
//}
//func FileUploadError(c *gin.Context) *upload.File {
//	return fileUpload(c, &upload.FileAttribute{
//		Key:     "file",
//		Exts:    filesExtsImage,
//		DirPath: conf.GetString("upload.imageSavePath") + "error/",
//		MaxSize: 20971520, //20M
//	})
//}
//
//func FileDel(file string) (bool, error) {
//	if uploadType == uploadTypeOSS {
//		return uploadOss.Remove(file)
//	} else if uploadType == uploadTypeLocal {
//		return uploadLocal.Remove(file)
//	}
//
//	return false, nil
//}
//
//func fileUpload(c *gin.Context, attribute *upload.FileAttribute) *upload.File {
//	if attribute.Key == "" {
//		attribute.Key = "file"
//	}
//	if uploadType == uploadTypeOSS {
//		return uploadOss.Upload(c, attribute)
//	} else if uploadType == uploadTypeLocal {
//		return uploadLocal.Upload(c, attribute)
//	}
//
//	return &upload.File{
//		Error: errors.New("uploadType not check"),
//	}
//}

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

var tConfigBaseToml = tplNode{
	NameFormat: "conf/base.toml",
	TplContent: `
# toml 转 struct : https://github.com/xuri/toml-to-go

[app]
    name = "{{.AppPkg}}"
    version = 1
    timeZone = "Asia/Shanghai" # 时区
    env = "dev"                # 环境仅允许本地 eg:dev eg:prod
	isCron = true                # 是否启动内置的后台计划任务
    log = "./runtime/logs/app"
`,
}
var tConfigDevToml = tplNode{
	NameFormat: "conf/dev/conf.toml",
	TplContent: `
[server]
    addr = "{{.AppAddr}}"          # eg: www.mojotv.cn eg:localhost eg:127.0.0.1
    port = 8000
    runMode = "debug"           # eg:debug eg:release eg:test
    readTimeout = 60
    writeTimeout = 60
    apiPrefix = "v1"           # api前缀，一般为版本号,设置为后 api/{api_prefix}/resource
    staticPath = "./static/"   # 静态路径,必须是绝对路径或相对于go build可执行文件
    logDir    = "./runtime/logs/http"   # logs的目录
    isNotRoute = true             # 如果为true,则所有未找到的路由都将转到static/index.html。
    isSwagger = true              # 开启 swagger api doc
    isCors = true                 # 开启cors跨域限制
    isHttps = false               # 如果 addr 为域名，则是否开启https,证书来自 LetsEncrypt
    isHttpLog = true             # httplog

[jwt]
    secret = "{{.AppSecret}}"
    key = "Authorization"
    expireHour = 24            # jwt颁布后生效的小时时间

[upload]
    type = "local"               # eg:oss eg:local
    imageSavePath = "images/"
    imageAllowExts = ".jpeg,.jpg,.png,.bmp,.gif"
    domain = ""
    [upload.local]
        baseDir = "./runtime/upload/"
    [upload.typeOss]
        accessKeyID = ""
        accessKeySecret = ""
        endpoint = ""
        bucketName = ""

[db]
    type = "{{.DbType}}"
    addr = "{{.DbAddr}}"
    port = 3306
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
`,
}


var tMod = tplNode{
	NameFormat: "go.mod",
	TplContent: `module {{.AppPkg}}

go 1.16
`,
}

var parseOneList = []tplNode{tDocIndex, tDocOauth2, tStaticIndex, tTaskCore, tTaskExample, tMod, tDockerfile, tConf, tDocYaml, tReadme, tGitIgnore, tFileManager, tAppLogging, tHandlerApp, tHandlerGin, tMain, tConfigBaseToml, tConfigDevToml, tLICENSE, tAppDb, tAppCache}

var parseObjList = []tplNode{tModelObj, tColumnObj, tHandlersObj}
