/**
* @Author: TheLife
* @Date: 2020-11-8 6:50 下午
 */
package dbUtils

import (
	"github.com/jinzhu/gorm"
	"go-gulu/dbTools/v1/dbUtils/tool/join"
)

type DbJoin struct {
	LeftJoin []join.LeftJoin
}

func (d *DbUtils) Join(w interface{}) *DbUtils {
	if d.Joins == nil {
		d.Joins = &DbJoin{}
	}

	switch t := w.(type) {
	case join.LeftJoin:
		d.Joins.LeftJoin = append(d.Joins.LeftJoin, t)
	}

	return d
}

func (d *DbUtils) GetJoin(tx *gorm.DB) *gorm.DB {
	if d == nil || d.Joins == nil {
		return tx
	}

	for _, value := range d.Joins.LeftJoin {
		tx = tx.Joins(value.String())
	}

	return tx
}
