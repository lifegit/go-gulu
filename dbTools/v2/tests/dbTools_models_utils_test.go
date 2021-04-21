/**
* @Author: TheLife
* @Date: 2020-11-8 5:57 下午
 */
package tests

import (
	"fmt"
	"go-gulu/dbTools/v2/hooks"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func getMysqlDb(dbAddr, dbName, dbCharset, dbUser, dbPassword string) (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local", dbUser, dbPassword, dbAddr, dbName, dbCharset)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic("mysql conn fail:",err)
		return
	}

	// ConnPool db conn pool
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	var field hooks.TimeFieldsModel
	db.Callback().Create().Replace("gorm:update_time_stamp", field.HookUpdateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", field.HookUpdateTimeStampForUpdateCallback)


	return
}

func initMysqlDb()  {
	db = getMysqlDb("localhost:3306", "user", "pass", "text", "")
}