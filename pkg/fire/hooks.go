/**
* @Author: TheLife
* @Date: 2021/6/23 下午9:37
 */
package fire

import (
	"gorm.io/plugin/soft_delete"
)

// 重新定义 gorm.Model

type TimeFieldsModel struct {
	// 13位毫秒时间戳
	CreatedAt int64 `gorm:"autoCreateTime:milli;type:bigint(13);unsigned;comment:创建时间" json:"created_at"`
	// 13位毫秒时间戳,兼容用户一秒内多次点击
	UpdatedAt int64 `gorm:"autoUpdateTime:milli;type:bigint(13);unsigned;comment:修改时间" json:"updated_at"`
	// 软删除,默认为0,删除时设置当前10位秒数时间戳
	DeletedAt soft_delete.DeletedAt `gorm:"type:bigint(13);unsigned;comment:删除时间" json:"deleted_at"`
}
