/**
* @Author: TheLife
* @Date: 2020-11-8 5:57 下午
 */
package dbTools

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gulu/dbTools/dbUtils"
	"go-gulu/dbTools/hoos"
	"go-gulu/pagination"
	"go-gulu/paramValidator"
	"go-gulu/structure"
)

/*
CREATE TABLE `db_clouddream_prod`.`无标题`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `adminid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作的管理员id，对应 tb_admins_info 的id',
  `agentid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值的代理id，对应 tb_agents_formal_info 的id',
  `type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 [1,manualAdd,手动加款|2,manualReduce,手动减款|3,cardAdd,充值卡加款|4,cardReduce,充值卡减款|5,manualAddGive,手动加款赠款|6,cardAddGive,充值卡加款赠款|7,manualGive,手动赠款]',
  `code` varchar(26) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '消费的卡密',
  `money` decimal(8, 2) NOT NULL DEFAULT 0.00 COMMENT '充值金额',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `time_created` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `time_updated` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `time_deleted` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 211 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '表_记录_充值_用户_信息' ROW_FORMAT = Compact;

INSERT INTO `db_clouddream_prod`.`tb_record_recharge_agent_info`(`id`, `adminid`, `agentid`, `type`, `code`, `money`, `remark`, `time_created`, `time_updated`, `time_deleted`) VALUES (1, 2, 1, 1, '', 1000.00, '', 1586149681000, 1586149681528, 0);
INSERT INTO `db_clouddream_prod`.`tb_record_recharge_agent_info`(`id`, `adminid`, `agentid`, `type`, `code`, `money`, `remark`, `time_created`, `time_updated`, `time_deleted`) VALUES (2, 2, 2, 1, '', 60.00, '', 1586336440000, 1586336440940, 0);
INSERT INTO `db_clouddream_prod`.`tb_record_recharge_agent_info`(`id`, `adminid`, `agentid`, `type`, `code`, `money`, `remark`, `time_created`, `time_updated`, `time_deleted`) VALUES (3, 2, 4, 1, '', 100.00, '', 1586338451000, 1586338451450, 0);
INSERT INTO `db_clouddream_prod`.`tb_record_recharge_agent_info`(`id`, `adminid`, `agentid`, `type`, `code`, `money`, `remark`, `time_created`, `time_updated`, `time_deleted`) VALUES (4, 2, 2, 1, '', 500.00, '', 1586340063000, 1586340063282, 0);
INSERT INTO `db_clouddream_prod`.`tb_record_recharge_agent_info`(`id`, `adminid`, `agentid`, `type`, `code`, `money`, `remark`, `time_created`, `time_updated`, `time_deleted`) VALUES (5, 2, 3, 1, '', 100.00, '', 1586358937000, 1586358937312, 0);

*/

//TbRecordRechargeInfo
type TbRecordRechargeAgentInfo struct {
	hoos.TimeFieldsModel
	DbUtils *dbUtils.DbUtils `gorm:"-" json:"-" form:"-"`

	Id      uint    `gorm:"column:id" form:"id" json:"id" comment:"" columnType:"int(10) unsigned" dataType:"int" columnKey:"PRI"`
	Adminid uint    `bindingCreate:"required" gorm:"column:adminid" form:"adminid" json:"adminid" comment:"操作的管理员id，对应 tb_admins_info 的id" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	Agentid uint    `bindingCreate:"required" gorm:"column:agentid" form:"agentid" json:"agentid" comment:"充值的代理id，对应 tb_agents_formal_info 的id" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	Type    uint    `bindingCreate:"required" gorm:"column:type" form:"type" json:"type" comment:"类型 [1,manualAdd,手动加款|2,manualReduce,手动减款|3,cardAdd,充值卡加款|4,cardReduce,充值卡减款|5,manualAddGive,手动加款赠款|6,cardAddGive,充值卡加款赠款|7,manualGive,手动赠款]" columnType:"tinyint(1) unsigned" dataType:"tinyint" columnKey:""`
	Code    string  `gorm:"column:code" form:"code" json:"code" comment:"消费的卡密" columnType:"varchar(26)" dataType:"varchar" columnKey:""`
	Money   float32 `bindingCreate:"required" gorm:"column:money" form:"money" json:"money" comment:"充值金额" columnType:"decimal(8,2)" dataType:"decimal" columnKey:""`
	Remark  string  `gorm:"column:remark" form:"remark" json:"remark" comment:"备注" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
}
//TableName
func (m *TbRecordRechargeAgentInfo) TableName() string {
	return "tb_record_recharge_agent_info"
}

//isExists
func (m *TbRecordRechargeAgentInfo) IsExists() (b bool) {
	one := &TbRecordRechargeAgentInfo{}
	err := m.DbUtils.CrudOne([]string{"1"}, m, one, db)

	return err == nil
}

//One
func (m *TbRecordRechargeAgentInfo) One(fields []string) (one *TbRecordRechargeAgentInfo, err error) {
	one = &TbRecordRechargeAgentInfo{}
	err = m.DbUtils.CrudOne(fields, m, one, db)

	return
}

//All
func (m *TbRecordRechargeAgentInfo) All(fields []string) (list *[]TbRecordRechargeAgentInfo, err error) {
	list = &[]TbRecordRechargeAgentInfo{}
	err = m.DbUtils.CrudAll(fields, m, list, db)

	return
}

//AllPage
func (m *TbRecordRechargeAgentInfo) AllPage(fields []string, list interface{}, pageSize uint) (page pagination.Page, err error) {
	count, err := m.DbUtils.CrudAllPage(fields, m, list, pageSize, db)

	return pagination.Page{Total: count, Size: pageSize}, err
}

//Create
func (m *TbRecordRechargeAgentInfo) Create() (err error) {
	// bindingCreate:"required"
	if err = paramValidator.ValidateCreate.Struct(m); err != nil {
		return
	}

	m.Id = 0

	return dbUtils.InitDb(m.DbUtils, db).Create(m).Error
}

//Update
func (m *TbRecordRechargeAgentInfo) Update(limit1 bool) (err error) {
	if m.Id == 0 && m.DbUtils.WhereIsEmpty() {
		return errors.New("update condition is not exist")
	}

	where := TbRecordRechargeAgentInfo{Id: m.Id}
	m.Id = 0

	return m.DbUtils.CrudUpdate(structure.StructToMap(*m), where, db, limit1)
}

//Delete
func (m *TbRecordRechargeAgentInfo) Delete() error {
	if m.Id == 0 && m.DbUtils.WhereIsEmpty() {
		return errors.New("resource must not be zero value")
	}
	return m.DbUtils.CrudDelete(m, db)
}

//Count
func (m *TbRecordRechargeAgentInfo) Count() (count uint, err error) {
	count, err = m.DbUtils.CrudCount(m, db)

	return
}