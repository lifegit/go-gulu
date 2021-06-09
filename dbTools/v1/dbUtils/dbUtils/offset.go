/**
* @Author: TheLife
* @Date: 2020-11-8 7:04 下午
 */
package dbUtils

import (
	"github.com/jinzhu/gorm"
	"github.com/lifegit/go-gulu/dbTools/v1/dbUtils/tool/offset"
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
		tx = tx.Offset(*d.Offset)
	}
	return tx
}
