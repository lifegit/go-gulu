/**
* @Author: TheLife
* @Date: 2020-11-8 5:57 下午
 */
package dbTools

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-gulu/dbTools/hoos"
	"log"
)

var db *gorm.DB

func getMysqlDb(dbAddr, dbName, dbCharset, dbUser, dbPassword string) (db *gorm.DB) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local", dbUser, dbPassword, dbAddr, dbName, dbCharset))
	if err != nil {
		log.Panic("mysql conn fail:",err)
		return
	}
	var field hoos.TimeFieldsModel
	db.Callback().Create().Replace("gorm:update_time_stamp", field.HookUpdateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", field.HookUpdateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", field.HookDeleteCallback)

	return
}

func initMysqlDb()  {
	db = getMysqlDb("localhost:3306", "user", "pass", "text", "")
}