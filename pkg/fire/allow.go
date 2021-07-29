/**
* @Author: TheLife
* @Date: 2021/5/26 下午11:21
 */
package fire

import (
	"reflect"
)

// 这是一个接收请求参数, 过滤到sql条件的漏斗工具。
// 实现开箱即用, 快速匹配筛选条件。

type Allow struct {
	// where
	Where []string
	Range []string
	In    []string
	Like  []string

	// order
	Sorts []string
}

type Sort map[string]interface{}

// allowSort
func (a *Allow) AllowSort(sort Sort, db *Fire) *Fire {
	toCamel2Case(sort)

	for _, condItem := range a.Sorts {
		for column, value := range sort {
			if column == condItem {
				db = db.OrderByColumn(column, If(value == "ascend", OrderAsc, OrderDesc).(OrderType))
			}
		}
	}

	return db
}

type Params map[string]interface{}

// allowParams
func (a *Allow) AllowParams(params Params, db *Fire) *Fire {
	// used Allow.key loop: fixed SQL order, we can put the condition of low energy consumption in the front
	// not used Params loop: range map is no order, it may result in different SQL generated each time

	toCamel2Case(params)

	// where
	for _, condItem := range a.Where {
		for column, value := range params {
			if column == condItem {
				switch data := value.(type) {
				case []interface{}:
					if len(data) >= 1 {
						db = db.WhereCompare(column, data[0])
					}
				case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
					db = db.WhereCompare(column, value)
				}
			}
		}
	}

	// Range
	for _, condItem := range a.Range {
		for column, value := range params {
			if column == condItem {
				switch data := value.(type) {
				case []interface{}, []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float32, []float64:
					dValue := reflect.ValueOf(data)
					if dValue.Len() >= 2 {
						db = db.WhereRange(column, dValue.Index(0).Interface(), dValue.Index(1).Interface())
					}
				}
			}
		}
	}

	// In
	for _, condItem := range a.In {
		for column, value := range params {
			if column == condItem {
				data := reflect.ValueOf(value)
				if data.Kind() == reflect.Slice {
					db = db.WhereIn(column, value)
				}
			}
		}
	}

	// Like
	for _, condItem := range a.Like {
		for column, value := range params {
			if column == condItem {
				switch value.(type) {
				case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
					db = db.WhereLike(column, value)
				}
			}
		}
	}

	return db
}
