/**
* @Author: TheLife
* @Date: 2020-11-8 6:58 下午
 */
package dbUtils

import (
	"github.com/lifegit/go-gulu/dbTools/v2/tool/order"
	"gorm.io/gorm"
)

type DbOrder order.Order

func (d *DbUtils) Order(w interface{}) *DbUtils {
	switch t := w.(type) {
	case DbOrder:
		d.Orders = &t
	}

	return d
}

func (d *DbUtils) GetOrder(tx *gorm.DB) *gorm.DB {
	if d == nil || d.Orders == nil {
		return tx
	}

	tx = tx.Order(order.Order(*d.Orders).String())

	return tx
}
