/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

import (
	"github.com/jinzhu/gorm"
)

type DbUtils struct {
	Db      *gorm.DB `json:"-"`
	Where   *DbWheres
	Updates *DbUpdates
	Order   *Order
	Join    *DbJoin
	Offset  uint
}
