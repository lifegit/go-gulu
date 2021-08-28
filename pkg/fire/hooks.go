// Package fire /**
// 重新定义 gorm.Model

package fire

import (
	"gorm.io/plugin/soft_delete"
)

type TimeFieldsEditModel struct {
	// 13位毫秒时间戳
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli;type:bigint(13);unsigned;comment:创建时间" json:"createdAt"`
	// 13位毫秒时间戳,兼容用户一秒内多次点击
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime:milli;type:bigint(13);unsigned;comment:修改时间" json:"updatedAt"`
}

type TimeFieldsModel struct {
	TimeFieldsEditModel
	// 软删除,默认为0,删除时设置当前10位秒数时间戳
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(13);unsigned;comment:删除时间" json:"deletedAt"`
}



