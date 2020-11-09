/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package pagination

import (
	"fmt"
	"go-gulu/arrayconv"
	"go-gulu/dbTools/dbUtils"
	"go-gulu/dbTools/tool/order"
	"go-gulu/dbTools/tool/where"
)

const DefaultLimit = 15

type Pagination struct {
	Page  uint
	Limit uint

	DbUtils *dbUtils.DbUtils
}

func New(page uint, limit uint) *Pagination {
	offset := dbUtils.DbOffset(limit * (page -1))
	return &Pagination{
		Page:  page,
		Limit: limit,

		DbUtils: &dbUtils.DbUtils{
			Wheres:  &dbUtils.DbWheres{},
			Joins:   &dbUtils.DbJoin{},
			Offset:  &offset,
		},
	}
}

type Searched struct {
	Name   string
	AsName string
	Vague  bool
}
////假设接收来自客户端的数据
//ma:= make(map[string]interface{})
//ma["projectid"] = []int{1}
//ma["type"] = []int{0,1}
//ma["tate"] = []string{"11"} //非法数据
//ma["name"] = "王"
func (p *Pagination) AllowFiltered(tableName string, data *map[string]interface{}, filtered []string, searched []Searched) {
	if data != nil {
		for key, val := range *data {
			if arrayconv.StringIn(key, filtered) {
				if arr, ok := val.([]interface{}); ok {
					length := len(arr)
					if length > 1 {
						p.DbUtils.Where(where.In{
							Not:   false,
							Field: fmt.Sprintf("`%s`.`%s`", tableName, key),
							In:    val,
						})
					} else if length == 1 {
						p.DbUtils.Where(where.Compare{
							Field: fmt.Sprintf("%s.`%s`", tableName, key),
							Type:  where.CompareEqual,
							Text:  val,
						})
					}
				}
			} else if sea := arrayGet(key, searched); sea != nil {
				if sea.Vague {
					p.DbUtils.Where(where.Like{
						Field: fmt.Sprintf("%s.`%s`", tableName, sea.Name),
						Text:  val.(string),
					})
				} else {
					p.DbUtils.Where(where.Compare{
						Field: fmt.Sprintf("%s.`%s`", tableName, sea.Name),
						Type:  where.CompareEqual,
						Text:  val,
					})
				}
			}
		}
	}
}

func (p *Pagination) AllowSorted(tableName string, o *order.Order, allow []string, defaultOrder *order.Order) {
	if o != nil && arrayconv.StringIn(o.Field, allow) {
		p.DbUtils.Order(dbUtils.DbOrder{
			Field: fmt.Sprintf("%s.`%s`", tableName, o.Field),
			Type:  o.Type,
		})
		return
	}

	if defaultOrder != nil && p.DbUtils.Orders == nil {
		p.DbUtils.Order(dbUtils.DbOrder{
			Field: fmt.Sprintf("%s.`%s`", tableName, defaultOrder.Field),
			Type:  defaultOrder.Type,
		})
	}
}