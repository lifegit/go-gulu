package ginbro

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
- cite: 一些常用工具
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
	"github.com/lifegit/go-gulu/dbTools/dbUtils"
	"github.com/lifegit/go-gulu/pagination"
	"github.com/lifegit/go-gulu/paramValidator"
	"github.com/lifegit/go-gulu/structure"
)

//{{.ModelName}}
type {{.ModelName}} struct {
	DbUtils		*dbUtils.DbUtils	_[BACKQUOTE]_gorm:"-" json:"-"_[BACKQUOTE]_
	{{range .Properties}}
	 {{.ModelProp}}      {{.ModelType}}         _[BACKQUOTE]_bindingCreate:"required" {{.ModelTag}}_[BACKQUOTE]_{{end}}
}
//TableName
func (m *{{.ModelName}}) TableName() string {
	return "{{.TableName}}"
}
//isExists
func (m *{{.ModelName}}) IsExists() (b bool) {
	one := &{{.ModelName}}{}
	err := m.DbUtils.CrudOne([]string{"1"}, m, one, db)

	return err == nil
}
//One
func (m *{{.ModelName}}) One(fields []string) (one *{{.ModelName}}, err error) {
	one = &{{.ModelName}}{}
	err = m.DbUtils.CrudOne(fields, m, one, db)

	return
}
//All
func (m *{{.ModelName}}) All(fields []string) (list *[]{{.ModelName}}, err error) {
	list = &[]{{.ModelName}}{}
	err = m.DbUtils.CrudAll(fields, m, list, db)

	return
}
//AllPage
func (m *{{.ModelName}}) AllPage(fields []string, list interface{}, pageSize uint) (page pagination.Page, err error) {
	count, err := m.DbUtils.CrudAllPage(fields, m, list, pageSize, db)

	return pagination.Page{Total: count, Size: pageSize}, err
}
//Create
func (m *{{.ModelName}}) Create() (err error) {
	// bindingCreate:"required"
	if err = paramValidator.ValidateCreate.Struct(m); err != nil {
		return
	}

	m.Id = 0

	return dbUtils.InitDb(m.DbUtils, db).Create(m).Error
}
//Update
func (m *{{.ModelName}}) Update(limit1 bool) (err error) {
	if m.Id == 0 && m.DbUtils.WhereIsEmpty() {
		return errors.New("update condition is not exist")
	}

	where := {{.ModelName}}{Id: m.Id}
	m.Id = 0

	return m.DbUtils.CrudUpdate(structure.StructToMap(*m), where, db, limit1)
}
//Delete
func (m *{{.ModelName}}) Delete() error {
	if m.Id == 0 && m.DbUtils.WhereIsEmpty() {
		return errors.New("resource must not be zero value")
	}
	return m.DbUtils.CrudDelete(m, db)
}
//Count
func (m *{{.ModelName}}) Count() (count uint, err error) {
	count, err = m.DbUtils.CrudCount(m, db)

	return
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
	"github.com/lifegit/go-gulu/app"
	"github.com/lifegit/go-gulu/dbTools/dbUtils"
	"github.com/lifegit/go-gulu/pagination"
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
	id, err := app.ParseParamID(c)
	if app.HandleError(c, err) {
		return
	}
	mdl := models.{{.ModelName}}{Id: id}
	data, err := mdl.One(nil)
	if app.HandleError(c, err) {
		return
	}
	app.JsonData(c, data)
}
{{end}}
//Create
func {{.HandlerName}}Create(c *gin.Context) {
	var mdl models.{{.ModelName}}
	err := c.ShouldBind(&mdl)
	if app.HandleError(c, err) {
		return
	}
	err = mdl.Create()
	if app.HandleError(c, err) {
		return
	}
	app.JsonData(c, mdl)
}
//Update
func {{.HandlerName}}Update(c *gin.Context) {
	var mdl models.{{.ModelName}}
	err := c.ShouldBindWith(&mdl, binding.FormPost)
	if app.HandleError(c, err) {
		return
	}

	err = mdl.Update(true)
	if app.HandleError(c, err) {
		return
	}

	app.JsonSuccess(c)
}

//All
func {{.HandlerName}}All(c *gin.Context) {
	var param pagination.Param
	err := c.ShouldBind(&param)
	if app.HandleError(c, err) {
		return
	}

	query := pagination.New(param.Page)
	query.AllowFiltered({{.TableName}}.TableName, param.Filtered,
		[]string{ },
		[]pagination.Searched{ },
	)
	query.AllowSorted({{.TableName}}.TableName, param.Sorted,
		[]string{ {{.TableName}}.TimeCreated },
		&dbUtils.Order{
			Field: {{.TableName}}.TimeCreated,
			Type:  dbUtils.OrderDesc,
		},
	)

	var list = &[]models.{{.ModelName}}{}
	var mdl models.{{.ModelName}}
	page, err := mdl.AllPage(nil, list, query)
	if app.HandleError(c, err) {
		return
	}

	app.JsonPagination(c, list, page)
}
{{if .HasId}}
//Delete
func {{.HandlerName}}Delete(c *gin.Context) {
	id, err := app.ParseParamID(c)
	if app.HandleError(c, err) {
		return
	}
	mdl := models.{{.ModelName}}{Id: id}
	err = mdl.Delete()
	if app.HandleError(c, err) {
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
var tModelJwt = tplNode{
	NameFormat: "model/m_jwt.go",
	TplContent: `
package model

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"{{.AppPkg}}/config"
	"time"
)

func jwtGenerateToken(m *{{.ModelName}}) (*jwtObj, error) {
	m.{{.PasswordPropertyName}} = ""
	expireAfterTime := time.Hour * time.Duration(config.GetInt("GinbroApp.jwt_expire_hour"))
	iss := config.GetString("GinbroApp.name")
	appSecret := config.GetString("GinbroApp.secret")
	expireTime := time.Now().Add(expireAfterTime)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", m.Id),
		Issuer:    iss,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(appSecret))
	if err != nil {
		logrus.WithError(err).Fatal("config is wrong, can not generate jwt")
	}
	data := &jwtObj{     {{.ModelName}}: *m, Token: tokenString, Expire: expireTime, ExpireTs: expireTime.Unix()}
	return data, err
}

type jwtObj struct {
	{{.ModelName}}
	Token    string    _[BACKQUOTE]_json:"token"_[BACKQUOTE]_
	Expire   time.Time _[BACKQUOTE]_json:"expire"_[BACKQUOTE]_
	ExpireTs int64     _[BACKQUOTE]_json:"expire_ts"_[BACKQUOTE]_
}
//JwtParseUser
func JwtParseUser(tokenString string) (*{{.ModelName}}, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := config.GetString("GinbroApp.secret")
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return nil, errors.New("token is expired")
	}
	appName := config.GetString("GinbroApp.name")
	if !claims.VerifyIssuer(appName, true) {
		return nil, errors.New("token's issuer is wrong,greetings Hacker")
	}
	key := fmt.Sprintf("login:%s", claims.Id)
	jwtObj, err := mem.GetJwtObj(key)
	if err != nil {
		return nil, err
	}
	return &jwtObj.{{.ModelName}}, err
}
//GetJwtObj
func (s *memoryStore) GetJwtObj(id string) (value *jwtObj, err error) {
	vv, err := s.Get(id, false)
	if err != nil {
		return nil, err
	}
	value, ok := vv.(*jwtObj)
	if ok {
		return value, nil
	}
	return nil, errors.New("mem:has value of this id, but is not type of *jwtObj")
}

`,
}

var tModelBase = tplNode{
	NameFormat: "models/app.go",
	TplContent: `
package models

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var RedisDB *redis.Client

func init() {
	db = initMysqlDb()
	RedisDB = initCache()
}

//Close clear db collection
func Close() {
	if db != nil {
		_ = db.Close()
	}
	if RedisDB != nil {
		_ = RedisDB.Close()
	}
}

`,
}

var tModelDb = tplNode{
	NameFormat: "models/app_db.go",
	TplContent: `
package models

import (
	"{{.AppPkg}}/cite/appLogging"
	"{{.AppPkg}}/cite/conf"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lifegit/go-gulu/dbTools/dbUtils"
	"strings"
)

func initMysqlDb() (db *gorm.DB) {
	db, err := createDatabase()
	if err != nil {
		appLogging.Log.WithError(err).Fatal("create database connection failed")
		return
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return setting.DatabaseSetting.TablePrefix + defaultTableName
	//}
	//db.SingularTable(true)

	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//enable Gorm mysql log
	if flag := conf.GetBool("enable.sqlLog"); flag {
		db.LogMode(flag)
		//f, err := os.OpenFile("mysql_gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		//if err != nil {
		//	logrus.WithError(err).Fatalln("could not create mysql gorm log file")
		//}
		//logger :=  New(f,"", Ldate)
		//db.SetLogger(logger)
	}
	//db.AutoMigrate()
	var field hooks.TimeFieldsModel
	db.Callback().Create().Replace("gorm:update_time_stamp", field.HookUpdateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", field.HookUpdateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", field.HookDeleteCallback)

	return
}

func createDatabase() (*gorm.DB, error) {
	dbType := conf.GetString("db.type")
	dbAddr := fmt.Sprintf("%s:%d", conf.GetString("db.addr"), conf.GetInt("db.port"))
	dbName := conf.GetString("db.database")
	dbUser := conf.GetString("db.username")
	dbPassword := conf.GetString("db.password")
	dbCharset := conf.GetString("db.charset")
	conn := ""
	switch dbType {
	case "mysql":
		conn = fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local", dbUser, dbPassword, dbAddr, dbName, dbCharset)
	case "sqlite":
		conn = dbAddr
	case "mssql":
		return nil, errors.New("TODO:suport sqlServer")
	case "postgres":
		hostPort := strings.Split(dbAddr, ":")
		if len(hostPort) == 2 {
			return nil, errors.New("db.addr must be like this host:ip")
		}
		conn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", hostPort[0], hostPort[1], dbUser, dbName, dbPassword)
	default:
		return nil, fmt.Errorf("database type %s is not supported by felix db", dbType)
	}
	return gorm.Open(dbType, conn)
}

`,
}

var tModelCache = tplNode{
	NameFormat: "models/app_cache.go",
	TplContent: `
package models

import (
	"{{.AppPkg}}/cite/appLogging"
	"{{.AppPkg}}/cite/conf"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/lifegit/go-gulu/gredis"
)

func initCache() *redis.Client {
	redisClient, err := gredis.CreateRedis(fmt.Sprintf("%s:%d", conf.GetString("redis.addr"), conf.GetInt("redis.port")), conf.GetString("redis.password"), conf.GetInt("redis.dbIdx"))
	if err != nil {
		appLogging.Log.WithError(err).Fatal("could not connect to the redis server")
	}
	return redisClient
}
`,
}

var tMain = tplNode{
	NameFormat: "main.go",
	TplContent: `
package main

import (
	"{{.AppPkg}}/cite/conf"
	"{{.AppPkg}}/handlers"
	"{{.AppPkg}}/tasks"
)

func main() {
	if conf.GetBool("enable.cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()

	handlers.ServerRun()
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
	if app.HandleError(c, err) {
		return
	}
	ip := c.ClientIP()
	data, err := mdl.Login(ip)
	if app.HandleError(c, err) {
		return
	}
	jsonData(c, data)
}
`,
}

var tHandlerApp = tplNode{
	NameFormat: "handlers/app.go",
	TplContent: `
package handlers

import (
	"{{.AppPkg}}/cite/appLogging"
	"{{.AppPkg}}/cite/conf"
	"{{.AppPkg}}/models"
	"github.com/lifegit/go-gulu/paramValidator"
	"time"
)

func init() {
	// timeZone
	_, _ = time.LoadLocation(conf.GetString("app.timeZone"))

	// resultApi
	result.Setup()

	// paramValidator
	paramValidator.Setup()
}

//ServerRun start the server
func ServerRun() {
	appLogging.Log.Infof("app: %s , version: %d at start ...", conf.GetString("app.name"), conf.GetInt("app.version"))

	result.Running()
}

//Close server app
func Close() {
	models.Close()
}
`,
}

var tHandlerGin = tplNode{
	NameFormat: "handlers/app_gin.go",
	TplContent: `
package handlers

import (
	"{{.AppPkg}}/cite/appLogging"
	"{{.AppPkg}}/cite/conf"
	"errors"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/lifegit/go-gulu/ginMiddleware/mwCors"
	"github.com/lifegit/go-gulu/ginMiddleware/mwJwt"
	"github.com/lifegit/go-gulu/ginMiddleware/mwLogger"
	"net/http"
	"path"
	"time"
)

var result ResultApi

type ResultApi struct {
	Gin *gin.Engine
	Api *gin.RouterGroup

	Jwt  mwJwt.MwJwt
	Addr string
}

func (r *ResultApi) Setup() {
	// 设置模式，设置模式要放在调用Default()函数之前
	r.Addr = fmt.Sprintf("%s:%d", conf.GetString("server.addr"), conf.GetInt("server.port"))

	// mode
	gin.SetMode(conf.GetString("server.runMode"))
	r.Gin = gin.New()
	// middlewareRecovery
	r.Gin.Use(gin.Recovery())
	// middlewareLogger
	middlewareLogger, err := mwLogger.NewLoggerMiddleware(true, conf.GetBool("enable.httpLog"), conf.GetString("server.logDir"))
	if err != nil {
		appLogging.Log.WithError(err).Fatal("gin middleware logger is io error")
	}
	r.Gin.Use(middlewareLogger)
	// middlewareCors
	if conf.GetBool("enable.cors") {
		r.Gin.Use(mwCors.NewCorsMiddleware())
	}
	// middlewareJwt
	r.Jwt = mwJwt.NewJwtMiddleware(conf.GetString("jwt.key"), conf.GetString("app.name"), conf.GetString("jwt.secret"), conf.GetString("jwt.key"))

	// staticPath
	if staticPath := conf.GetString("server.staticPath"); staticPath != "" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/ for front-end static html files", r.Addr))
		r.Gin.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	}
	// notRoute
	if conf.GetBool("enable.notRoute") {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/docs for RESTful Is NotRoute", r.Addr))
		r.Gin.NoRoute(func(c *gin.Context) {
			//file := path.Join(sp, "index.html")
			//c.File(file)
			c.Status(http.StatusNotFound)
		})
	}
	// swaggerApi
	if conf.GetBool("enable.swagger") && conf.GetString("app.env") != "prod" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/docs for RESTful APIs Document", r.Addr))
		//add edit your own swagger.docs.yml file in ./swagger/docs.yml
		//generateSwaggerDocJson()
		r.Gin.Static("docs", "./docs")
	}
	// appInfo
	if conf.GetString("app.env") != "prod" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/app/info for app info only on not-prod mode", r.Addr))
		r.Gin.GET("/app/info", func(c *gin.Context) {
			m := make(map[string]map[string]string)
			for _, val := range []string{"app", "server", "enable", "upload"} {
				m[val] = conf.GetStringMapString(val)
			}
			c.JSON(200, m)
		})
	}
	// apiPrefix
	r.Api = r.Gin.Group(path.Join("services", conf.GetString("server.apiPrefix")))
}

func (r *ResultApi) Running() {
	appLogging.Log.Infof("http result server at listening %s", r.Addr)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.GetInt("server.port")),
		Handler:        r.Gin,
		ReadTimeout:    time.Second * time.Duration(conf.GetInt("server.readTimeout")),
		WriteTimeout:   time.Second * time.Duration(conf.GetInt("server.writeTimeout")),
		MaxHeaderBytes: 1 << 20,
	}
	if conf.GetBool("enable.https") {
		// https
		if err := autotls.Run(r.Gin, r.Addr); err != nil {
			appLogging.Log.WithError(err).Fatal("https result server fail run !")
		}
		//if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
		//	appLogging.Log.Errorf("https result server is run err: %v!", err)
		//}
	} else {
		// http
		if err := server.ListenAndServe(); err != nil {
			appLogging.Log.Errorf("http result server is run err: %v!", err)
		}
	}

	//// endless
	//// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//serverNew := endless.NewServer(endPoint, routersInit)
	//serverNew.BeforeBegin = func(add string) {
	//	logging.Logger.Info("Actual pid is %d", syscall.Getpid())
	//}
	//err = serverNew.ListenAndServe()
	//if err != nil {
	//	appLogging.Logger.Error("server err: %v", err)
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

func RunTasks() {
	//// Do jobs with params
	//core.Every(1).Second().Do(taskWithParams, 1, "hello")
	//
	//// Do jobs without params
	//core.Every(1).Second().Do(task)
	//core.Every(2).Seconds().Do(task)
	//core.Every(1).Minute().Do(task)
	//core.Every(2).Minutes().Do(task)
	//core.Every(1).Hour().Do(task)
	//core.Every(2).Hours().Do(task)
	//core.Every(1).Day().Do(task)
	//core.Every(2).Days().Do(task)
	//
	//// Do jobs on specific weekday
	//core.Every(1).Monday().Do(task)
	//core.Every(1).Thursday().Do(task)
	//
	//// function At() take a string like 'hour:min'
	//core.Every(1).Day().At("10:30").Do(task)
	//core.Every(1).Monday().At("18:30").Do(task)
	//
	//// remove, clear and next_run
	//_, time := NextRun()
	//fmt.Println(time)
	//
	//// Remove(task)
	//// Clear()
	//
	//// function Start start all the pending jobs
	//<-core.Start()
	//
	//// also , you can create a your new scheduler,
	//// to run two scheduler concurrently
	//s := NewScheduler()
	//s.Every(3).Seconds().Do(task)
	//<-s.Start()
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
func taskWithParams(a int, b string) {
	fmt.Println(a, b)
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
	NameFormat: "cite/appLogging/appLogging.go",
	TplContent: `
package appLogging

import (
	"{{.AppPkg}}/cite/conf"
	"github.com/sirupsen/logrus"
	"github.com/lifegit/go-gulu/logging"
)

var Log *logrus.Logger

func init() {
	Log = logging.NewLogger(conf.GetString("log.baseDir"), 3, &logrus.JSONFormatter{}, nil)
}

`,
}
var tConf = tplNode{
	NameFormat: "cite/conf/conf.go",
	TplContent: `
package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

const path = "./conf"
const confType = "toml"
const confMain = "base"
const confNormal = "conf"

func init() {
	watch := handWatchFileChange
	// mainConf init
	v, err := getLocalConfToViper(path, confMain, confType, &watch)
	if err != nil {
		logrus.WithError(err).Fatal(err)
		return
	}
	// setting to default
	AllSettingsToDefault(v)

	// check ENV
	//env := os.Getenv("GO_ENV")
	v, err = getLocalConfToViper(fmt.Sprintf("%s/%s/", path, viper.GetString("app.env")), confNormal, confType, &watch)
	if err != nil {
		logrus.WithError(err).Fatal(err)
		return
	}
	// setting to default
	AllSettingsToDefault(v)

	//fmt.Println("local","AllSettings",viper.AllSettings())
	//return nil
}
func AllSettingsToDefault(setting *viper.Viper) {
	configs := setting.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
}

// handWatchFileChange
func handWatchFileChange(event fsnotify.Event) {
	if event.Op == fsnotify.Create || event.Op == fsnotify.Write || event.Op == fsnotify.Chmod {
		path := event.Name[:strings.LastIndex(event.Name, "/")]
		pathname := event.Name[strings.LastIndex(event.Name, "/")+1 : strings.LastIndex(event.Name, ".")]
		confType := event.Name[strings.LastIndex(event.Name, ".")+1:]
		v, err := getLocalConfToViper(path, pathname, confType, nil)
		if err == nil {
			// setting to default
			AllSettingsToDefault(v)
			_ = fmt.Sprintf("application configuration'initialization watch success in %s", event.Name)
			return
		}
	}

	_ = fmt.Sprintf("application configuration'initialization watch fail in %s, file op is %s", event.Name, event.Op)
}

// getFileConfToViper
func getLocalConfToViper(path, pathname, confType string, WatchChange *func(fsnotify.Event)) (*viper.Viper, error) {
	// viper init
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigType(confType)
	v.SetConfigName(pathname)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if WatchChange != nil {
		v.WatchConfig()
		v.OnConfigChange(*WatchChange)
	}
	return v, nil
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}
`,
}

var tFileManager = tplNode{
	NameFormat: "cite/fileManager/fileManager.go",
	TplContent: `
package fileManager

import (
	"{{.AppPkg}}/cite/conf"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/lifegit/go-gulu/upload"
	"github.com/lifegit/go-gulu/upload/local"
	"github.com/lifegit/go-gulu/upload/oss"
	"strings"
)

const uploadTypeOSS = "oss"
const uploadTypeLocal = "local"

var uploadType = conf.GetString("upload.type")
var filesExtsImage = strings.Split(conf.GetString("upload.imageAllowExts"), ",")

var uploadOss *oss.Oss
var uploadLocal *local.Local

func init() {
	if uploadType == uploadTypeOSS {
		c, err := oss.New(conf.GetString("upload.typeOss.endpoint"), conf.GetString("upload.typeOss.accessKeyID"), conf.GetString("upload.typeOss.accessKeySecret"), conf.GetString("upload.typeOss.bucketName"), conf.GetString("upload.domain"))
		if err != nil {
			err = errors.New("could not check to the oss server")
			logrus.WithError(err).Fatal(err)
			return
		}
		uploadOss = c
	} else if uploadType == uploadTypeLocal {
		uploadLocal = local.New(conf.GetString("upload.local.baseDir"), conf.GetString("upload.domain"))
	} else {
		err := errors.New("upload type is nil")
		logrus.WithError(err).Fatal(err)
		return
	}
}

func FileUploadAvatar(c *gin.Context, identity string) *upload.File {
	return fileUpload(c, &upload.FileAttribute{
		Exts:    filesExtsImage,
		DirPath: conf.GetString("upload.imageSavePath") + "avatars/" + identity + "/",
		MaxSize: 3145728,
	})
}

func FileUploadRecoveryAvatar(c *gin.Context, recoveryId string) *upload.File {
	return fileUpload(c, &upload.FileAttribute{
		Key:     "avatar",
		Exts:    filesExtsImage,
		DirPath: conf.GetString("upload.imageSavePath") + "recovery/avatar/" + recoveryId + "/",
		MaxSize: 1048576,
	})
}
func FileUploadProductAvatar(c *gin.Context, productId string) *upload.File {
	return fileUpload(c, &upload.FileAttribute{
		Key:     "avatar",
		Exts:    filesExtsImage,
		DirPath: conf.GetString("upload.imageSavePath") + "product/avatar/" + productId + "/",
		MaxSize: 1048576,
	})
}

func FileUploadRecoveryImg(c *gin.Context, recoveryId string) *upload.File {
	return fileUpload(c, &upload.FileAttribute{
		Key:     "img",
		Exts:    filesExtsImage,
		DirPath: conf.GetString("upload.imageSavePath") + "recovery/img/" + recoveryId + "/",
		MaxSize: 2097152,
	})
}

func FileUploadResources(c *gin.Context) *upload.File {
	return fileUpload(c, &upload.FileAttribute{
		Key:     "file",
		Exts:    filesExtsImage,
		DirPath: conf.GetString("upload.imageSavePath") + "resources/",
		MaxSize: 2097152,
	})
}
func FileUploadError(c *gin.Context) *upload.File {
	return fileUpload(c, &upload.FileAttribute{
		Key:     "file",
		Exts:    filesExtsImage,
		DirPath: conf.GetString("upload.imageSavePath") + "error/",
		MaxSize: 20971520, //20M
	})
}

func FileDel(file string) (bool, error) {
	if uploadType == uploadTypeOSS {
		return uploadOss.Remove(file)
	} else if uploadType == uploadTypeLocal {
		return uploadLocal.Remove(file)
	}

	return false, nil
}

func fileUpload(c *gin.Context, attribute *upload.FileAttribute) *upload.File {
	if attribute.Key == "" {
		attribute.Key = "file"
	}
	if uploadType == uploadTypeOSS {
		return uploadOss.Upload(c, attribute)
	} else if uploadType == uploadTypeLocal {
		return uploadLocal.Upload(c, attribute)
	}

	return &upload.File{
		Error: errors.New("uploadType not check"),
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
    env = "dev"                # 环境仅允许本地 eg:dev eg:prod
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

[log]
    baseDir = "./runtime/logs/logging"


[enable]
    notRoute = true             # 如果为true,则所有未找到的路由都将转到static/index.html。
    swagger = true              # 开启 swagger api doc
    cors = true                 # 开启cors跨域限制
    sqlLog = true               # 在控制台显示gorm的logs
    https = false               # 如果 addr 为域名，则是否开启https,证书来自 LetsEncrypt
    cron = true                # 是否启动内置的后台计划任务
    httpLog = true             # httplog

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

//"tpl/config.toml": "config.toml",

var tMod = tplNode{
	NameFormat: "go.mod",
	TplContent: `module {{.AppPkg}}

go 1.12

replace go-gulu => /Users/yxs/GolandProjects/src/go-gulu
`,
}

var parseOneList = []tplNode{tDocIndex, tDocOauth2, tStaticIndex, tTaskCore, tTaskExample, tMod, tConf, tDocYaml, tReadme, tGitIgnore, tFileManager, tAppLogging, tHandlerApp, tHandlerGin, tMain, tConfigBaseToml, tConfigDevToml, tLICENSE, tModelBase, tModelDb, tModelCache}

var parseObjList = []tplNode{tModelObj, tColumnObj, tHandlersObj}
