// Package fire /**/
// 对 gorm.DB 的补充封装，实现更爽快得使用。属于基础层服务代码。
package fire

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/url"
	"reflect"
	"strings"
	"unicode"
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

func If(isA bool, a, b interface{}) interface{} {
	if isA {
		return a
	}

	return b
}

const ColumnAll = "*"

type FormatColumnType string

const (
	FormatColumnBackQuote      FormatColumnType = "`"
	FormatColumnQuotationMarks FormatColumnType = `"`
)

var formatColumn = string(FormatColumnBackQuote)

func SetFormatColumnType(v FormatColumnType) {
	formatColumn = string(v)
}

func FormatColumn(column ...string) (res string) {
	var list []string
	for _, item := range column {
		list = append(list, strings.Split(item, ".")...)
	}

	for key, value := range list {
		if len(value) >= 1 {
			// first and last is not `
			if value == ColumnAll {
				value = ColumnAll
			} else if value[:1] != formatColumn && value[len(value)-1:] != formatColumn {
				value = fmt.Sprintf("%s%s%s", formatColumn, value, formatColumn)
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

type Column struct {
	Table  string
	Column string
}

func (c *Column) String() string {
	return FormatColumn(c.Table, c.Column)
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
func toCamel2Case(m url.Values) {
	for key, value := range m {
		if !strings.Contains(key, "_") {
			delete(m, key)
			m[Camel2Case(key)] = value
		}
	}
}
func Camel2Case(name string) string {
	buffer := strings.Builder{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteString("_")
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
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
	c := If(compare != nil, compare, []CompareType{CompareEqual})
	tx := d.DB.Where(fmt.Sprintf("%s %s ?", FormatColumn(column), c.([]CompareType)[0]), value)

	return NewInstance(tx)
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
