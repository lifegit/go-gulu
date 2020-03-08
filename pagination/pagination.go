/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package pagination

import (
	"fmt"
	"go-gulu/arrayconv"
	"go-gulu/dbUtils"
)

type Searched struct {
	Name   string
	AsName string
	Vague  bool
}

type Pagination struct {
	Page  uint
	Limit uint

	DbUtils *dbUtils.DbUtils
}

func New(page uint) *Pagination {
	var limit uint = 15

	return &Pagination{
		Page:  page,
		Limit: limit,

		DbUtils: &dbUtils.DbUtils{
			Where:  &dbUtils.DbWheres{},
			Join:   &dbUtils.DbJoin{},
			Offset: limit * (page - 1),
		},
	}
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
						p.DbUtils.Where.In = append(p.DbUtils.Where.In, dbUtils.In{
							Field: fmt.Sprintf("%s.`%s`", tableName, key),
							In:    val,
						})
					} else if length == 1 {
						p.DbUtils.Where.Compare = append(p.DbUtils.Where.Compare, dbUtils.Compare{
							Field: fmt.Sprintf("%s.`%s`", tableName, key),
							Type:  dbUtils.CompareEqual,
							Text:  val,
						})
					}
				}
			} else if sea := arrayGet(key, searched); sea != nil {
				if sea.Vague {
					p.DbUtils.Where.Like = append(p.DbUtils.Where.Like, dbUtils.Like{
						Field: fmt.Sprintf("%s.`%s`", tableName, sea.Name),
						Text:  val.(string),
					})
				} else {
					p.DbUtils.Where.Compare = append(p.DbUtils.Where.Compare, dbUtils.Compare{
						Field: fmt.Sprintf("%s.`%s`", tableName, sea.Name),
						Type:  dbUtils.CompareEqual,
						Text:  val,
					})
				}
			}
		}
	}
}

func (p *Pagination) AllowSorted(tableName string, o *dbUtils.Order, allow []string, defaultOrder *dbUtils.Order) {
	if o != nil && arrayconv.StringIn(o.Field, allow) {
		p.DbUtils.Order = &dbUtils.Order{
			Field: fmt.Sprintf("%s.`%s`", tableName, o.Field),
			Type:  o.Type,
		}
		return
	}

	if defaultOrder != nil && p.DbUtils.Order == nil {
		p.DbUtils.Order = &dbUtils.Order{
			Field: fmt.Sprintf("%s.`%s`", tableName, defaultOrder.Field),
			Type:  defaultOrder.Type,
		}
	}
}

func (p *Pagination) AddJoin(joinTableName string, left dbUtils.JoinOn, right dbUtils.JoinOn) {
	p.DbUtils.Join.LeftJoin = append(p.DbUtils.Join.LeftJoin, dbUtils.LeftJoin{
		TableName: joinTableName,
		Left: dbUtils.JoinOn{
			Table: left.Table,
			Field: left.Field,
		},
		Right: dbUtils.JoinOn{
			Table: right.Table,
			Field: right.Field,
		},
	})
}
