/**
* @Author: TheLife
* @Date: 2020-11-8 5:57 下午
 */
package tests

import (
	"fmt"
	"github.com/lifegit/go-gulu/dbTools/v2/dbUtils"
	"github.com/lifegit/go-gulu/dbTools/v2/tool/update"
	"github.com/lifegit/go-gulu/dbTools/v2/tool/where"
	"testing"
	"time"
)

// 使用 func 进行 update 测试
// 推荐使用该方式，因为代码看起来简洁。虽然会多一些切片入栈代码，不在乎这一丢丢忽略不计的性能啦。
func TestUpdateListByFunc(t *testing.T) {
	initMysqlDb()

	// UPDATE `tb_record_recharge_agent_info` SET `code` = 0, `money` = `money` + 1, `time_updated` = 1604842519491  WHERE (`time_created` >= 1604842793)
	record := TbRecordRechargeAgentInfo{DbUtils: &dbUtils.DbUtils{}}
	record.DbUtils = record.DbUtils.Update(update.Set{
		Field: "Code",
		Value: 0,
	}).Update(update.Arithmetic{
		Field:  "money",
		Type:   update.ArithmeticIncrease,
		Number: 1,
	}).Where(where.Compare{
		Field: "time_created",
		Type:  where.CompareAboutEqual,
		Text:  time.Now().Unix(),
	})
	err := record.Update(false)
	fmt.Println(err)
}

//使用 struct 进行 update 测试
func TestUpdateListByStruct(t *testing.T) {
	initMysqlDb()

	// UPDATE `tb_record_recharge_agent_info` SET `code` = 0, `money` = `money` + 1, `time_updated` = 1604842519491  WHERE (`time_created` >= 1604842793)
	agent := TbRecordRechargeAgentInfo{DbUtils: &dbUtils.DbUtils{
		Updates: &dbUtils.DbUpdates{
			Set: []update.Set{
				{
					Field: "Code",
					Value: 0,
				},
			},
			Arithmetic: []update.Arithmetic{
				{
					Field:  "money",
					Type:   update.ArithmeticIncrease,
					Number: 1,
				},
			},
		},
		Wheres: &dbUtils.DbWheres{
			Compare: []where.Compare{
				{
					Field: "time_created",
					Type:  where.CompareAboutEqual,
					Text:  time.Now().Unix(),
				},
			},
		},
	}}

	err := agent.Update(false)
	fmt.Println(err)
}
