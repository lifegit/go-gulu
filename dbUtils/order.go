/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const OrderAsc = "asc"
const OrderDesc = "desc"

type Order struct {
	Field string `form:"name" binding:"required"`
	Type  string `binding:"required,eq=asc|eq=desc"`
}

func (utils DbUtils) GetOrder(tx *gorm.DB) *gorm.DB {
	order := utils.Order
	if order != nil {
		tx = tx.Order(fmt.Sprintf("%s %s", order.Field, order.Type))
	}
	return tx
}

func (utils DbUtils) GetOffset(tx *gorm.DB) *gorm.DB {
	offset := utils.Offset
	if offset > 0 {
		tx = tx.Offset(offset)
	}
	return tx
}
