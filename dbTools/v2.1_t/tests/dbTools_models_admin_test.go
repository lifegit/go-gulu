package tests

import (
	"errors"
	"github.com/lifegit/go-gulu/dbTools/v2/dbUtils"
	"github.com/lifegit/go-gulu/dbTools/v2/hooks"
	"github.com/lifegit/go-gulu/pagination/v1"
	"github.com/lifegit/go-gulu/paramValidator"
	"github.com/lifegit/go-gulu/structure"
)

/*

CREATE TABLE `db_clouddream_prod`.`无标题`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `use` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '可用 [1,allowed,可用|2,disabled,禁止]',
  `username` varchar(18) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '账号',
  `password` char(60) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `name` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` char(11) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '手机号 (可改)',
  `time_lasted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后一次登录时间',
  `time_lasted_mes` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后一次同步消息时间',
  `time_created` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `time_updated` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `time_deleted` bigint(13) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `un_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '表_管理员_信息' ROW_FORMAT = Compact;

INSERT INTO `db_clouddream_prod`.`tb_admins_info`(`id`, `use`, `username`, `password`, `avatar`, `name`, `mobile`, `time_lasted`, `time_lasted_mes`, `time_created`, `time_updated`, `time_deleted`) VALUES (1, 1, 'user', 'pass', '', '婉儿', '', 1604910552, 0, 1577808000, 1604910552392, 0);


*/

//TbAdminsInfo
type TbAdminsInfo struct {
	hooks.TimeFieldsModel
	DbUtils *dbUtils.DbUtils `gorm:"-" json:"-" form:"-"`

	Id            uint   `gorm:"column:id" json:"id" comment:"" columnType:"int(10) unsigned" dataType:"int" columnKey:"PRI"`
	Use           uint   `bindingCreate:"required" gorm:"column:use" json:"use" comment:"可用 [1,allowed,可用|2,disabled,禁止]" columnType:"tinyint(1) unsigned" dataType:"tinyint" columnKey:""`
	Username      string `bindingCreate:"required" gorm:"column:username" json:"username" comment:"账号" columnType:"varchar(18)" dataType:"varchar" columnKey:"UNI"`
	Password      string `bindingCreate:"required" bindingCreate:"required" gorm:"column:password" json:"-" comment:"密码" columnType:"char(60)" dataType:"char" columnKey:""`
	Avatar        string `bindingCreate:"required" gorm:"column:avatar" form:"avatar" json:"avatar" comment:"头像" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	Name          string `gorm:"column:name" form:"name" json:"name" comment:"姓名" columnType:"varchar(15)" dataType:"varchar" columnKey:""`
	Mobile        string `gorm:"column:mobile" form:"mobile" json:"mobile" comment:"手机号 (可改)" columnType:"char(11)" dataType:"char" columnKey:""`
	TimeLasted    uint   `gorm:"column:time_lasted" json:"time_lasted" comment:"最后一次登录时间" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	TimeLastedMes uint   `gorm:"column:time_lasted_mes" json:"time_lasted_mes" comment:"最后一次同步消息时间" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
}

//TableName
func (m *TbAdminsInfo) TableName() string {
	return "tb_admins_info"
}

//isExists
func (m *TbAdminsInfo) IsExists() (b bool) {
	one := &TbAdminsInfo{}
	err := m.DbUtils.CrudOne([]string{"1"}, m, one, db)

	return err == nil
}

//One
func (m *TbAdminsInfo) One(fields []string) (one *TbAdminsInfo, err error) {
	one = &TbAdminsInfo{}
	err = m.DbUtils.CrudOne(fields, m, one, db)

	return
}

//All
func (m *TbAdminsInfo) All(fields []string) (list *[]TbAdminsInfo, err error) {
	list = &[]TbAdminsInfo{}
	err = m.DbUtils.CrudAll(fields, m, list, db)

	return
}

//AllPage
func (m *TbAdminsInfo) AllPage(fields []string, list interface{}, pageSize uint) (page pagination.Page, err error) {
	count, err := m.DbUtils.CrudAllPage(fields, m, list, pageSize, db)

	return pagination.Page{Total: count, Size: pageSize}, err
}

//Create
func (m *TbAdminsInfo) Create() (err error) {
	// bindingCreate:"required"
	if err = paramValidator.ValidateCreate.Struct(m); err != nil {
		return
	}

	m.Id = 0

	return dbUtils.InitDb(m.DbUtils, db).Create(m).Error
}

//Update
func (m *TbAdminsInfo) Update(limit1 bool) (err error) {
	if m.Id == 0 && m.DbUtils.WhereIsEmpty() {
		return errors.New("update condition is not exist")
	}

	where := TbAdminsInfo{Id: m.Id}
	m.Id = 0

	return m.DbUtils.CrudUpdate(structure.StructToMap(*m), where, db, limit1)
}

//Delete
func (m *TbAdminsInfo) Delete() error {
	if m.Id == 0 && m.DbUtils.WhereIsEmpty() {
		return errors.New("resource must not be zero value")
	}
	return m.DbUtils.CrudDelete(m, db)
}

//Count
func (m *TbAdminsInfo) Count() (count uint, err error) {
	count, err = m.DbUtils.CrudCount(m, db)

	return
}
