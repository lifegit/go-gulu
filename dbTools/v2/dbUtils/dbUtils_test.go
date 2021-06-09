/**
* @Author: TheLife
* @Date: 2021/5/26 下午6:09
 */
package dbUtils_test

import (
	"fmt"
	"go-gulu/dbTools/v2/dbUtils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"testing"
)

func TestDbUtils(t *testing.T) {
	db := initMysqlDb()

	user := TbUser{DbUtils: &dbUtils.DbUtils{DB: db}, Name: "zhangsan "}
	b := user.DbUtils.IsExists(user)

	fmt.Println(b)
}

//TbUser
type TbUser struct {
	DbUtils *dbUtils.DbUtils `gorm:"-" json:"-" form:"-"`
	Name    string           `bindingCreate:"required" gorm:"column:name" json:"username" comment:"账号" columnType:"varchar(18)" dataType:"varchar" columnKey:"UNI"`
	Time    int64
	Age     int64
	Tag     string
}

//TableName
func (m *TbUser) TableName() string {
	return "user"
}

func initMysqlDb() (db *gorm.DB) {
	gormConf := &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 表名不加复数s
	}

	dbUser := "com1yema1mysq1ok"
	dbPassword := "new1pwd1yema1ok1_"
	dbAddr := "127.0.0.1"
	dbName := "db_test"
	dbCharset := "utf8"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local", dbUser, dbPassword, dbAddr, dbName, dbCharset)), gormConf)
	if err != nil {
		log.Fatalln(err, "create database connection failed")
		return
	}

	// ConnPool db conn pool
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	//// hooks
	//var field hooks.TimeFieldsModel
	//db.Callback().Create().Replace("gorm:update_time_stamp", field.HookUpdateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", field.HookUpdateTimeStampForUpdateCallback)

	return
}
