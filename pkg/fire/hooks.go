// Package fire /**
// 重新定义 gorm.Model

package fire

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type TimeFields1Model struct {
	// 13位毫秒时间戳
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli;type:bigint(13);unsigned;comment:创建时间;<-:create" json:"createdAt"`
}

func (t *TimeFields1Model) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now().UnixNano() / 1e6
	return
}

type TimeFields2Model struct {
	TimeFields1Model
	// 13位毫秒时间戳,兼容用户一秒内多次点击修改
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime:milli;type:bigint(13);unsigned;comment:修改时间;<-" json:"updatedAt"`
}

type TimeFields3Model struct {
	TimeFields2Model
	// 软删除,默认为0,删除时设置当前10位秒数时间戳
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(13);unsigned;comment:删除时间;<-:false" json:"deletedAt"`
}
