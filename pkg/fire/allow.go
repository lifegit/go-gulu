// Package fire /**
// 这是一个接收请求参数, 过滤到sql条件的漏斗工具。
// 实现开箱即用, 快速匹配筛选条件。

package fire

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/url"
	"strings"
)

type Allow struct {
	// where
	Where []string
	Range []string
	In    []string
	Like  []string

	// order
	Sorts []string
}

// AllowSort allow sort
func (a *Allow) AllowSort(sort url.Values, db *Fire) *Fire {
	v := sort["sort"]
	if v == nil || len(v) <= 0 {
		return db
	}
	m := make(map[string]string)
	if err := json.Unmarshal([]byte(v[0]), &m); err != nil || len(m) <= 0 {
		return db
	}
	if err := validator.New().Var(m, "omitempty,max=1,dive,keys,required,endkeys,eq=ascend|eq=descend"); err != nil {
		return db
	}
	for _, condItem := range a.Sorts {
		for column, value := range m {
			if caColumn := Camel2Case(column); caColumn == condItem {
				db = db.OrderByColumn(caColumn, If(strings.HasPrefix(value, string(OrderAsc)), OrderAsc, OrderDesc).(OrderType))
			}
		}
	}

	return db
}

// AllowParams allow params
func (a *Allow) AllowParams(params url.Values, db *Fire) *Fire {
	// used Allow.key loop: fixed SQL order, we can put the condition of low energy consumption in the front
	// not used Params loop: range map is no order, it may result in different SQL generated each time

	toCamel2Case(params)

	// where
	for _, condItem := range a.Where {
		for column, value := range params {
			if column == condItem {
				if len(value) >= 1 {
					db = db.WhereCompare(column, value[0])
				}
			}
		}
	}

	// Range
	for _, condItem := range a.Range {
		for column, value := range params {
			if column == condItem {
				if len(value) >= 2 {
					db = db.WhereRange(column, value[0], value[1])
				}
			}
		}
	}

	// In
	for _, condItem := range a.In {
		for column, value := range params {
			if column == condItem {
				db = db.WhereIn(column, value)
			}
		}
	}

	// condItem: mkt_roads.name
	// column:road		value: {"name":"11"}
	// Like
	for _, condItem := range a.Like {
		for column, value := range params {
			if column == condItem {
				if len(value) >= 1 {
					db = db.WhereLike(column, value[0])
				}
			}
		}
	}

	return db
}
