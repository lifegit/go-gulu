/**
* @Author: TheLife
* @Date: 2020-11-8 6:58 下午
 */
package dbUtils

import (
	"github.com/jinzhu/gorm"
	order "github.com/lifegit/go-gulu/dbTools/v1/dbUtils/tool/order"
)

type DbOrder order.Order

func (d *DbUtils) Order(w interface{}) *DbUtils {
	switch t := w.(type) {
	case order.Order:
		o := DbOrder(t)
		d.Orders = &o
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
