/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type JoinOn struct {
	Table string
	Field string
}
type LeftJoin struct {
	TableName string
	Left      JoinOn
	Right     JoinOn
}
type DbJoin struct {
	LeftJoin []LeftJoin
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
