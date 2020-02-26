/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

import (
	"fmt"
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

func (utils DbUtils) GetWhere(tx *gorm.DB) *gorm.DB {
	where := utils.Where
	if where != nil {
		for _, value := range where.Compare {
			tx = tx.Where(fmt.Sprintf("%s %s ?", value.Field, value.Type), value.Text)
		}
		for _, value := range where.In {
			n := ""
			if value.Not {
				n = "not"
			}
			tx = tx.Where(fmt.Sprintf("%s %s in(?)", value.Field, n), value.In)
		}

		for _, value := range where.Range {
			tx = tx.Where(fmt.Sprintf("%s >= ? AND %s <= ?", value.Field, value.Field), value.Start, value.End)
		}

		for _, value := range where.Like {
			tx = tx.Where(value.Field+" LIKE ?", fmt.Sprintf("%%%s%%", value.Text))
		}
	}
	return tx
}

func (utils DbUtils) GetUpdate() (m *map[string]interface{}) {
	updates := utils.Updates
	if updates != nil {
		maps := make(map[string]interface{})
		for _, value := range updates.Arithmetic {
			maps[value.Field] = gorm.Expr(fmt.Sprintf("%s %s ?", value.Field, value.Type), value.Number)
		}
		for _, value := range updates.Set {
			maps[value.Field] = value.Value
		}
		return &maps
	}
	return nil
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
func (utils DbUtils) GetJoin(tx *gorm.DB) *gorm.DB {
	join := utils.Join

	if join != nil {
		for _, item := range join.LeftJoin {
			tx = tx.Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s", item.TableName, item.Left.Table, item.Left.Field, item.Right.Table, item.Right.Field))
		}
	}
	return tx
}
