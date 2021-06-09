/**
* @Author: TheLife
* @Date: 2020-4-24 1:35 下午
 */
package hooks

import (
	"gorm.io/gorm"
	"time"
)

type TimeFieldsModel struct {
	TimeUpdated uint `gorm:"column:time_updated" json:"time_updated" comment:"更新时间" columnType:"bigint(13) unsigned" dataType:"bigint" columnKey:""`
	TimeDeleted uint `gorm:"column:time_deleted" json:"-" comment:"删除时间" columnType:"bigint(13) unsigned" dataType:"bigint" columnKey:""`
	TimeCreated uint `gorm:"column:time_created" json:"time_created" comment:"添加时间" columnType:"bigint(13) unsigned" dataType:"bigint" columnKey:""`
}

// 新建钩子在持久化之前
func (*TimeFieldsModel) HookUpdateTimeStampForCreateCallback(db *gorm.DB) {
	//nowTime := time.Now().UnixNano() / 1e6
	//if db.Statement.Error == nil{
	//	if field := db.Statement.Schema.LookUpField("TimeCreated"); field != nil{
	//
	//		//vv,v2 := field.FieldType
	//
	//		// field.Set(db.Statement.ReflectValue, nowTime)
	//
	//
	//		//field
	//		//field.ReflectValueOf
	//		//field.ReflectValueOf
	//		//
	//		//fmt.Println("2", db.Statement.ReflectValue.Index(3))
	//		//
	//		//if _, isZero := field.ValueOf(db.Statement.ReflectValue.Index(3)); isZero {
	//		//	_ = field.Set(db.Statement.ReflectValue, nowTime)
	//		//}
	//	}
	//	//
	//	//if field := db.Statement.Schema.LookUpField("TimeUpdated"); field != nil{
	//	//	if _, isZero := field.ValueOf(db.Statement.ReflectValue); isZero {
	//	//		_ = field.Set(db.Statement.ReflectValue, nowTime)
	//	//	}
	//	//}
	//}
}

// 更新钩子在持久化之前
func (*TimeFieldsModel) HookUpdateTimeStampForUpdateCallback(db *gorm.DB) {
	if _, ok := db.Statement.Get("gorm:update_column"); !ok {
		db.Statement.SetColumn("TimeUpdated", time.Now().UnixNano()/1e6)
	}
}
