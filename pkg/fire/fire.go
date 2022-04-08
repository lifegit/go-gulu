// Package fire /**/
// 对 gorm.DB 的补充封装，实现更爽快得使用。属于基础层服务代码。
package fire

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/url"
	"reflect"
)

type Fire struct {
	*gorm.DB
}

func NewInstance(db *gorm.DB) *Fire {
	return &Fire{DB: db}
}

func (d *Fire) Close() (err error) {
	dbs, err := d.DB.DB()
	if err != nil {
		return
	}

	return dbs.Close()
}

// === SELECT ===

// WhereIn
// column IN(?)
// column NOT IN(?)
func (d *Fire) ModelWhere(model interface{}) *Fire {
	tx := d.Model(model).Where(model)

	return NewInstance(tx)
}

// PreloadAll
// TODO：Multiple SQL, gorm bonding data, so query conditions other than the main table are not supported
func (d *Fire) PreloadAll() *Fire {
	tx := d.DB.Preload(clause.Associations)

	return NewInstance(tx)
}

// PreloadJoin
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

func (d *Fire) Allow(params url.Values, allow Allow) *Fire {
	tx := NewInstance(d.DB)
	tx = allow.AllowParams(params, tx)
	tx = allow.AllowSort(params, tx)

	return tx
}

type CompareType string

const (
	CompareEqual        CompareType = "="
	CompareGreaterEqual CompareType = ">="
	CompareGreater      CompareType = ">"
	CompareSmallerEqual CompareType = "<="
	CompareSmaller      CompareType = "<"
)

// === where ===

// WhereCompare
// column CompareEqual ? # column = ?
func (d *Fire) WhereCompare(column string, value interface{}, compare ...CompareType) *Fire {
	return NewInstance(d.Scopes(WhereCompare(column, value, compare...)))
}

func WhereCompare(column string, value interface{}, compare ...CompareType) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		c := If(compare != nil, compare, []CompareType{CompareEqual}).([]CompareType)[0]
		return db.Where(fmt.Sprintf("%s %s ?", FormatColumn(column), c), value)
	}
}

// WhereIn
// column IN(?)
// column NOT IN(?)
func (d *Fire) WhereIn(column string, value interface{}, isNot ...bool) *Fire {
	c := If(isNot != nil && isNot[0], "NOT", "")
	tx := d.Where(fmt.Sprintf("%s %s IN (?)", FormatColumn(column), c), value)

	return NewInstance(tx)
}

// WhereLike
// column LIKE %?%
func (d *Fire) WhereLike(column string, value interface{}) *Fire {
	tx := d.Where(fmt.Sprintf("%s LIKE ?", FormatColumn(column)), fmt.Sprintf("%%%s%%", value))

	return NewInstance(tx)
}

// WhereRange
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

// UpdateArithmetic
// field = field ArithmeticType Number # field = field + 1
func UpdateArithmetic(column string, value interface{}, art ArithmeticType) (m map[string]interface{}) {
	m = make(map[string]interface{})
	m[column] = gorm.Expr(fmt.Sprintf("%s %s ?", FormatColumn(column), art), value)

	return
}
