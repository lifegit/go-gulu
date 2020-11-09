/**
* @Author: TheLife
* @Date: 2020-4-24 1:35 下午
 */
package hoos

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type TimeFieldsModel struct {
	TimeCreated uint `gorm:"column:time_created" json:"time_created" comment:"添加时间" columnType:"bigint(13) unsigned" dataType:"bigint" columnKey:""`
	TimeUpdated uint `gorm:"column:time_updated" json:"time_updated" comment:"更新时间" columnType:"bigint(13) unsigned" dataType:"bigint" columnKey:""`
	TimeDeleted uint `gorm:"column:time_deleted" json:"-" comment:"删除时间" columnType:"bigint(13) unsigned" dataType:"bigint" columnKey:""`
}

// 新建钩子在持久化之前
func (*TimeFieldsModel) HookUpdateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().UnixNano() / 1e6
		if createTimeField, ok := scope.FieldByName("TimeCreated"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("TimeUpdated"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 更新钩子在持久化之前
func (*TimeFieldsModel) HookUpdateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("TimeUpdated", time.Now().UnixNano()/1e6)
	}
}

// 删除钩子在删除之前
func (*TimeFieldsModel) HookDeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("TimeDeleted")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().UnixNano()/1e6),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
