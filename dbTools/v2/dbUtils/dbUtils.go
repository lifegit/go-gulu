/**
* @Author: TheLife
* @Date: 2021/5/26 下午5:01
 */
package dbUtils

import (
	"fmt"
	"gorm.io/gorm"
)

type DbUtils struct {
	*gorm.DB
}

func If(isA bool, a, b interface{}) interface{} {
	if isA {
		return a
	}

	return b
}

type CompareType string

const (
	CompareEqual      CompareType = "="
	CompareAboutEqual CompareType = ">="
	CompareAbout      CompareType = ">"
	CompareLessEqual  CompareType = "<="
	CompareLess       CompareType = "<"
)

// === where ===

// column CompareEqual ? # column = ?
func (d *DbUtils) WhereCompare(column string, value interface{}, compare ...CompareType) *DbUtils {
	c := If(compare != nil, compare, []CompareType{CompareEqual})
	d.DB = d.Where(fmt.Sprintf("`%s` %s ?", column, c.([]CompareType)[0]), value)

	return d
}

// column IN(?)
// column NOT IN(?)
func (d *DbUtils) WhereIn(column string, value interface{}, isNot ...bool) *DbUtils {
	c := If(isNot != nil && isNot[0], "NOT", "")

	d.DB = d.Where(fmt.Sprintf("`%s` %s IN (?)", column, c), value)

	return d
}

// column LIKE %?%
func (d *DbUtils) WhereLike(column string, value interface{}) *DbUtils {
	d.DB = d.Where(fmt.Sprintf("`%s` LIKE ?", column), fmt.Sprintf("%%%s%%", value))

	return d
}

// column >= start ANd column <= end
func (d *DbUtils) WhereRange(column string, start interface{}, end interface{}) *DbUtils {
	d.DB = d.Where(fmt.Sprintf("`%s` >= ? AND `%s` <= ?", column, column), start, end)

	return d
}

// === update ===
type ArithmeticType string

const (
	ArithmeticIncrease ArithmeticType = "+"
	ArithmeticReduce   ArithmeticType = "-"
	ArithmeticMultiply ArithmeticType = "*"
	ArithmeticExcept   ArithmeticType = "/"
)

// field = field ArithmeticType Number # field = field + 1
func (d *DbUtils) UpdateArithmetic(column string, value float64, art ArithmeticType) *DbUtils {
	d.DB = d.Update(column, gorm.Expr(fmt.Sprintf("`%s` %s ?", column, art), value))

	return d
}

// === order ===
type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

func (d *DbUtils) OrderByColumn(column string, order OrderType) *DbUtils {
	d.DB = d.Order(fmt.Sprintf("%s %s", column, order))

	return d
}
