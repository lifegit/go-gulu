// Package fire /**
// 这是一个接收请求参数, 过滤到sql条件的漏斗工具。
// 实现开箱即用, 快速匹配筛选条件。

package fire

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/url"
	"strings"
)

type Filtered []interface{}

type Allow struct {
	// where
	Where Filtered
	Range Filtered
	In    Filtered
	Like  Filtered

	// order
	Sorts Filtered
}

// AllowSort allow sort
func (a *Allow) AllowSort(params url.Values, db *gorm.DB) *gorm.DB {
	param := make(url.Values)
	for _, item := range params["sort"] {
		m := make(map[string]string)
		if err := json.Unmarshal([]byte(item), &m); err != nil || len(m) <= 0 {
			continue
		}
		if err := validator.New().Var(m, "omitempty,max=1,dive,keys,required,endkeys,eq=ascend|eq=descend"); err != nil {
			continue
		}
		for k, v := range m {
			param.Set(k, v)
		}
	}

	a.Filtered(param, a.Sorts, func(column clause.Column, value []string) {
		exp := If(strings.HasPrefix(value[0], string(OrderAsc)), OrderAsc, OrderDesc).(OrderType)
		db = db.Order(Order(column, exp))
	})

	return db
}

// AllowParams allow params
func (a *Allow) AllowParams(params url.Values, db *gorm.DB) *gorm.DB {
	// used Allow.key loop: fixed SQL order, we can put the condition of low energy consumption in the front
	// not used Params loop: range map is no order, it may result in different SQL generated each time

	//where
	a.Filtered(params, a.Where, func(column clause.Column, value []string) {
		for _, v := range value {
			// json
			if strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}") {
				var m map[string]interface{}
				if err := json.Unmarshal([]byte(v), &m); err == nil {
					for k, ve := range m {
						// column.key = ?
						db = db.Clauses(WhereJsonEq(column, ve, k))
					}
				}
				continue
			}
			// column = ?
			db = db.Clauses(WhereCompare(column, v, CompareEq))
		}
	})

	// Range
	a.Filtered(params, a.Range, func(column clause.Column, value []string) {
		if len(value) >= 2 {
			//  column >= ? AND column <= ?
			db = db.Clauses(WhereRange(column, value[0], value[1]))
		}
	})

	// In
	a.Filtered(params, a.In, func(column clause.Column, value []string) {
		// column IN (?)
		db = db.Clauses(WhereIn(column, value))
	})

	// Like
	a.Filtered(params, a.Like, func(column clause.Column, value []string) {
		for _, v := range value {
			// column LIKE %?%
			db = db.Clauses(WhereLike(column, v))
		}
	})

	return db
}

func (a *Allow) ParseColumns(columns []interface{}) (res []clause.Column) {
	for _, column := range columns {
		var col clause.Column
		switch v := column.(type) {
		case clause.Column:
			col = v
		case string:
			col.Name = v
		}

		if col.Table == "" {
			col.Table = clause.CurrentTable
		}
		res = append(res, col)
	}
	return
}

func (a *Allow) Filtered(params url.Values, f Filtered, callBack func(column clause.Column, value []string)) {
	for _, arrowColumn := range a.ParseColumns(f) {
		for param, value := range params {
			if (param == arrowColumn.Name && arrowColumn.Alias == "") || (param == arrowColumn.Alias && arrowColumn.Alias != "") {
				arrowColumn.Alias = ""
				callBack(arrowColumn, value)
			}
		}
	}
}
