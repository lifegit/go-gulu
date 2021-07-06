/**
* @Author: TheLife
* @Date: 2021/5/26 下午5:01
 */
package fire

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"strings"
)

// 对 gorm.DB 的补充封装，实现更爽快的使用。属于基础层服务代码。

type Fire struct {
	*gorm.DB
}

func If(isA bool, a, b interface{}) interface{} {
	if isA {
		return a
	}

	return b
}

func NewInstance(db *gorm.DB) *Fire {
	return &Fire{DB: db}
}

// 格式化列
func FormatColumn(column string) (res string) {
	list := strings.Split(column, ".")
	for key, value := range list {
		if len(value) >= 1 {
			// first and last is not `
			if value[:1] != "`" && value[len(value)-1:] != "`" {
				value = fmt.Sprintf("`%s`", value)
			}
			res += value
			// isLast
			if key != len(list)-1 {
				res += "."
			}
		}
	}

	return
}

func (d *Fire) Close() {
	if dbs, err := d.DB.DB(); err == nil {
		_ = dbs.Close()
	}
}

// === SELECT ===
// TODO：Multiple SQL, gorm bonding data, so query conditions other than the main table are not supported
func (d *Fire) PreloadAll() *Fire {
	tx := d.DB.Preload(clause.Associations)

	return NewInstance(tx)
}

// TODO：Single SQL, mysql bonding data, so the conditions of all query tables are supported. use Join you need to pay attention to performance
func (d *Fire) PreloadJoin(model interface{}) *Fire {
	if reflect.TypeOf(model).Kind() != reflect.Struct {
		return nil
	}
	tx := d.DB
	key := reflect.TypeOf(model)
	val := reflect.ValueOf(model)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			if val.Field(i).CanInterface() {
				// anonymous not join
				if !key.Field(i).Anonymous {
					tx = tx.Joins(key.Field(i).Name)
				}
			}
		}
	}

	return NewInstance(tx)
}

// allow
func (d *Fire) Allow(param Param, allow Allow) *Fire {
	tx := NewInstance(d.DB)
	tx = allow.AllowParams(param.Params, tx)
	tx = allow.AllowSort(param.Sort, tx)

	return tx
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
func (d *Fire) WhereCompare(column string, value interface{}, compare ...CompareType) *Fire {
	c := If(compare != nil, compare, []CompareType{CompareEqual})
	tx := d.DB.Where(fmt.Sprintf("%s %s ?", FormatColumn(column), c.([]CompareType)[0]), value)

	return NewInstance(tx)
}

// column IN(?)
// column NOT IN(?)
func (d *Fire) WhereIn(column string, value interface{}, isNot ...bool) *Fire {
	c := If(isNot != nil && isNot[0], "NOT", "")
	tx := d.Where(fmt.Sprintf("%s %s IN (?)", FormatColumn(column), c), value)

	return NewInstance(tx)
}

// column LIKE %?%
func (d *Fire) WhereLike(column string, value interface{}) *Fire {
	tx := d.Where(fmt.Sprintf("%s LIKE ?", FormatColumn(column)), fmt.Sprintf("%%%s%%", value))

	return NewInstance(tx)
}

// column >= start ANd column <= end
func (d *Fire) WhereRange(column string, start interface{}, end interface{}) *Fire {
	formatColumn := FormatColumn(column)
	tx := d.Where(fmt.Sprintf("%s >= ? AND %s <= ?", formatColumn, formatColumn), start, end)

	return NewInstance(tx)
}

// === order ===
type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

func (d *Fire) OrderByColumn(column string, order OrderType, many ...bool) *Fire {
	if many == nil || !many[0] {
		delete(d.Statement.Clauses, "ORDER BY")
	}
	tx := d.Order(fmt.Sprintf("%s %s", FormatColumn(column), order))

	return NewInstance(tx)
}

// === update ===
type ArithmeticType string

const (
	ArithmeticIncrease ArithmeticType = "+"
	ArithmeticReduce   ArithmeticType = "-"
	ArithmeticMultiply ArithmeticType = "*"
	ArithmeticExcept   ArithmeticType = "/"
)

//field = field ArithmeticType Number # field = field + 1
func UpdateArithmetic(column string, value interface{}, art ArithmeticType) (m map[string]interface{}) {
	m = make(map[string]interface{})
	m[column] = gorm.Expr(fmt.Sprintf("%s %s ?", FormatColumn(column), art), value)

	return
}
