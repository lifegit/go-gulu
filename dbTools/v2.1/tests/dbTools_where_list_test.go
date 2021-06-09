/**
* @Author: TheLife
* @Date: 2020-11-8 5:57 下午
 */
package tests
//
//import (
//	"fmt"
//	"github.com/lifegit/go-gulu/dbTools/v2/dbUtils"
//	"github.com/lifegit/go-gulu/dbTools/v2/tool/join"
//	"github.com/lifegit/go-gulu/dbTools/v2/tool/where"
//	"github.com/lifegit/go-gulu/pagination/v1"
//	"testing"
//	"time"
//)
//
//// 使用 func 进行 where 测试 多表连接
//// 推荐使用该方式，因为代码看起来简洁。虽然会多一些切片入栈代码，不在乎这一丢丢忽略不计的性能啦。
//func TestWhereListByFuncOnMultipleTable(t *testing.T) {
//	initMysqlDb()
//
//	//  SELECT `tb_admins_info`.username, `tb_record_recharge_agent_info`.code, `tb_record_recharge_agent_info`.money FROM `tb_record_recharge_agent_info` LEFT JOIN tb_admins_info ON `tb_record_recharge_agent_info`.id = `tb_admins_info`.id  LIMIT 15
//	query := pagination.New(1, pagination.DefaultLimit)
//	record := TbRecordRechargeAgentInfo{DbUtils: query.DbUtils}
//
//	// 潜入其他表的字段
//	var list = &[]struct {
//		TbRecordRechargeAgentInfo
//		Username string `json:"username"`
//	}{}
//
//	// 当多表时，可能字段名称会重复，所以需要显示的给出需要的字段
//	var admin TbAdminsInfo
//	fields := dbUtils.FieldTab(
//		dbUtils.T{TableName: admin.TableName(), Fields: []dbUtils.Field{
//			{Name: "username", AsName: "adminname"},
//		}},
//		dbUtils.T{TableName: record.TableName(), Fields: []dbUtils.Field{
//			{Name: "code"},
//			{Name: "money"},
//		}},
//	)
//	record.DbUtils = record.DbUtils.Join(join.LeftJoin{
//		Left: join.JoinOn{
//			Table: record.TableName(),
//			Field: "id",
//		},
//		Right: join.JoinOn{
//			Table: admin.TableName(),
//			Field: "id",
//		},
//	})
//
//	data, err := record.AllPage(fields, list, query.Limit)
//	fmt.Println(data, err)
//}
//
//// 使用 func 进行 where 测试 单表
//// 推荐使用该方式，因为代码看起来简洁。虽然会多一些切片入栈代码，不在乎这一丢丢忽略不计的性能啦。
//func TestWhereListByFuncOnOneTable(t *testing.T) {
//	initMysqlDb()
//
//	// SELECT * FROM `tb_record_recharge_agent_info`  WHERE (`time_created` >= 0 AND `time_created` <= 1604949852000) LIMIT 15
//	query := pagination.New(1, pagination.DefaultLimit)
//	record := TbRecordRechargeAgentInfo{DbUtils: query.DbUtils}
//
//	var list = &[]TbRecordRechargeAgentInfo{}
//	record.DbUtils = record.DbUtils.Where(where.Range{
//		Field: "time_created",
//		Start: 0,
//		End:   time.Now().Unix() * 1000,
//	})
//
//	data, err := record.AllPage(nil, list, query.Limit)
//	fmt.Println(data, err)
//}
