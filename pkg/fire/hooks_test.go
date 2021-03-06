/**
* @Author: TheLife
* @Date: 2021/6/23 下午10:46
 */
package fire_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatedAt(t *testing.T) {
	user := &TbCompany{
		Address: "Shanghai",
		Name:    "lu",
	}
	tx := DBDryRun.Create(user)
	// `created_at`=1624612236447  `updated_at`=1624612236447
	sql := fmt.Sprintf("INSERT INTO `company` (`created_at`,`updated_at`,`deleted_at`,`address`,`name`) VALUES (%d,%d,'0','Shanghai','lu')", tx.Statement.Vars[0], tx.Statement.Vars[1])
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), sql)
}

func TestUpdatedAt(t *testing.T) {
	tx := DBDryRun.Where(TbCompany{Address: "Shanghai"}).Updates(TbCompany{Address: "Jinan"})
	// `updated_at`=1624612236447
	sql := fmt.Sprintf("UPDATE `company` SET `updated_at`=%d,`address`='Jinan' WHERE `company`.`address` = 'Shanghai'", tx.Statement.Vars[0])
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), sql)
}

func TestDeletedAt(t *testing.T) {
	// Delete
	tx := DBDryRun.Delete(TbCompany{ID: 1})
	// `deleted_at`=1624612236
	sql := fmt.Sprintf("UPDATE `company` SET `deleted_at`=%d WHERE `company`.`id` = 1 AND `company`.`deleted_at` = 0", tx.Statement.Vars[0])
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), sql)

	// Select
	tx = DBDryRun.Model(TbCompany{}).Where(TbCompany{Address: "Shanghai"}).Find(nil)
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `company` WHERE `company`.`address` = 'Shanghai' AND `company`.`deleted_at` = 0")
}
