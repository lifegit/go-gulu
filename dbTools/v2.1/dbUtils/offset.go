/**
* @Author: TheLife
* @Date: 2020-11-8 7:04 下午
 */
package dbUtils

import (
	"go-gulu/dbTools/v2/tool/offset"
	"gorm.io/gorm"
)

type DbOffset offset.Offset

func (d *DbUtils) OffSet(w interface{}) *DbUtils {
	switch t := w.(type) {
	case offset.Offset:
		o := DbOffset(t)
		d.Offset = &o
	}

	return d
}

func (d *DbUtils) GetOffSet(tx *gorm.DB) *gorm.DB {
	if d == nil || d.Offset == nil {
		return tx
	}

	if *d.Offset > 0 {
		tx = tx.Offset(int(offset.Offset(*d.Offset)))
	}
	return tx
}
