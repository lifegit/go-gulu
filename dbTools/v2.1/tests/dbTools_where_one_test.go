/**
* @Author: TheLife
* @Date: 2020-11-8 5:57 下午
 */
package tests

import (
	"fmt"
	"go-gulu/dbTools/v2/dbUtils"
	"go-gulu/dbTools/v2/tool/order"
	"go-gulu/dbTools/v2/tool/where"
	"testing"
	"time"
)

// 使用 func 进行 where 测试
// 推荐使用该方式，因为代码看起来简洁。虽然会多一些切片入栈代码，不在乎这一丢丢忽略不计的性能啦。
func TestWhereOneByFunc(t *testing.T) {
	initMysqlDb()

	// SELECT * FROM `tb_record_recharge_agent_info`  WHERE (`type` >= 1) AND (`time_updated` >= 0 AND `time_updated` <= 1604839616) AND (`adminid`  IN(1)) AND (`remark` LIKE '%消费%') ORDER BY `time_created` asc LIMIT 1
	record := TbRecordRechargeAgentInfo{DbUtils: &dbUtils.DbUtils{}}
	record.DbUtils = record.DbUtils.Where(where.Compare{
		Field: "type",
		Type:  where.CompareAboutEqual,
		Text:  1,
	}).Where(where.Range{
		Field: "time_updated",
		Start: 0,
		End:   time.Now().Unix(),
	}).Where(where.In{
		Field: "adminid",
		In:    1,
	}).Where(where.Like{
		Field: "remark",
		Text:  "消费",
	}).Order(order.Order{
		Field: "time_created",
		Type:  order.OrderAsc,
	})
	one, err := record.One(nil)
	fmt.Println(one, err)
}

// 使用 struct 进行 where 测试
func TestWhereOneByStruct(t *testing.T) {
	initMysqlDb()

	// SELECT * FROM `tb_record_recharge_agent_info`  WHERE (`type` >= 1) AND (`time_updated` >= 0 AND `time_updated` <= 1604839616) AND (`adminid`  IN(1)) AND (`remark` LIKE '%消费%') ORDER BY `time_created` asc LIMIT 1
	record := TbRecordRechargeAgentInfo{DbUtils: &dbUtils.DbUtils{
		Wheres: &dbUtils.DbWheres{
			Compare: []where.Compare{
				{
					Field: "type",
					Type:  where.CompareAboutEqual,
					Text:  1,
				},
			},
			Range: []where.Range{
				{
					Field: "time_updated",
					Start: 0,
					End:   time.Now().Unix(),
				},
			},
			In: []where.In{
				{
					Field: "adminid",
					In:    1,
				},
			},
			Like: []where.Like{
				{
					Field: "remark",
					Text:  "消费",
				},
			},
		},
		Orders: &dbUtils.DbOrder{
			Field: "time_created",
			Type:  order.OrderAsc,
		},
	}}

	one, err := record.One(nil)
	fmt.Println(one, err)
}
