/**
* @Author: TheLife
* @Date: 2020-11-8 5:57 下午
 */
package tests

import (
	"fmt"
	"go-gulu/dbTools/v2/dbUtils"
	"go-gulu/dbTools/v2/tool/update"
	"testing"
)



// 使用 func 进行 update 测试
// 推荐使用该方式，因为代码看起来简洁。虽然会多一些切片入栈代码，不在乎这一丢丢忽略不计的性能啦。
func TestUpdateOneByFunc(t *testing.T) {
	initMysqlDb()

	// UPDATE `tb_record_recharge_agent_info` SET `money` = `money` + 1, `time_updated` = 1604842912643  WHERE `tb_record_recharge_agent_info`.`id` = 5
	record := TbRecordRechargeAgentInfo{DbUtils: &dbUtils.DbUtils{}, Id: 5}
	record.DbUtils = record.DbUtils.Update(update.Arithmetic{
		Field: "money",
		Type: update.ArithmeticIncrease,
		Number: 1,
	})
	err := record.Update(false)
	fmt.Println(err)
}

//使用 struct 进行 update 测试
func TestUpdateOneByStruct(t *testing.T) {
	initMysqlDb()

	//UPDATE `tb_record_recharge_agent_info` SET `money` = `money` + 1, `time_updated` = 1604842912643  WHERE `tb_record_recharge_agent_info`.`id` = 5
	agent := TbRecordRechargeAgentInfo{DbUtils: &dbUtils.DbUtils{
		Updates: &dbUtils.DbUpdates{
			Arithmetic: []update.Arithmetic{
				{
					Field: "money",
					Type: update.ArithmeticIncrease,
					Number: 1,
				},
			},
		},
	},
		Id:5,
	}

	err := agent.Update(false)
	fmt.Println(err)
}