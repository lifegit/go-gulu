/**
* @Author: TheLife
* @Date: 2021/6/23 下午10:30
 */
package fire_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsExists(t *testing.T) {
	DBDryRun.IsExists(TbUser{ID: 1, Name: "Mr Wang"})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT 1 FROM `user` WHERE `user`.`id` = 1 AND `user`.`name` = 'Mr Wang' LIMIT 1")
}

func TestCrudOne(t *testing.T) {
	_ = DBDryRun.WhereLike("name", "M").CrudOne(TbUser{ID: 1, Age: 20}, &TbUser{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `name` LIKE '%M%' AND `user`.`id` = 1 AND `user`.`age` = 20 LIMIT 1")
}

func TestCrudAll(t *testing.T) {
	_ = DBDryRun.WhereRange("age", 18, 20).CrudAll(TbUser{}, &TbUser{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` >= 18 AND `age` <= 20")
}

func TestCrudAllPage(t *testing.T) {
	_, _ = DBDryRun.CrudAllPage(TbUser{}, &TbUser{}, fire.Page{
		Current:  3,
		PageSize: 5,
	})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(2), "SELECT count(*) FROM `user`")
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` LIMIT 5 OFFSET 10")
}

func TestCrudOnePreloadJoin(t *testing.T) {
	type User struct {
		TbUser
		Company TbCompany
	}
	_ = DBDryRun.CrudOnePreloadJoin(User{}, &User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT `user`.`id`,`user`.`company_id`,`user`.`name`,`user`.`tag`,`user`.`age`,`user`.`height`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`id` AS `Company__id`,`Company`.`address` AS `Company__address`,`Company`.`name` AS `Company__name` FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id` LIMIT 1")
}

func TestCrudAllPreloadJoin(t *testing.T) {
	type User struct {
		TbUser
		Company TbCompany
	}
	_ = DBDryRun.CrudAllPreloadJoin(User{}, &[]User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT `user`.`id`,`user`.`company_id`,`user`.`name`,`user`.`tag`,`user`.`age`,`user`.`height`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`id` AS `Company__id`,`Company`.`address` AS `Company__address`,`Company`.`name` AS `Company__name` FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id`")
}

func TestCrudAllPagePreloadJoin(t *testing.T) {
	type User struct {
		TbUser
		Company TbCompany
	}
	_, _ = DBDryRun.CrudAllPagePreloadJoin(User{}, &[]User{}, fire.Page{
		Current:  3,
		PageSize: 5,
	})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(2), "SELECT count(*) FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id`")
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT `user`.`id`,`user`.`company_id`,`user`.`name`,`user`.`tag`,`user`.`age`,`user`.`height`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`id` AS `Company__id`,`Company`.`address` AS `Company__address`,`Company`.`name` AS `Company__name` FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id` LIMIT 5 OFFSET 10")
}

func TestCrudOnePreloadAll(t *testing.T) {
	type User struct {
		TbUser
		Company TbCompany
	}
	_ = DB.CrudOnePreloadAll(User{TbUser: TbUser{Age: 18}}, &User{})
	xx := DB
	fmt.Println(xx)
	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `company` WHERE `company`.`id` = 1 AND `company`.`deleted_at` = 0")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `user`.`age` = 18 LIMIT 1")
}

func TestCrudAllPreloadAll(t *testing.T) {
	type User struct {
		TbUser
		Company TbCompany
	}
	_ = DB.CrudAllPreloadAll(User{}, &[]User{})
	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `company` WHERE `company`.`id` IN (1,2) AND `company`.`deleted_at` = 0")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user`")
}

func TestCrudAllPagePreloadAll(t *testing.T) {
	type User struct {
		TbUser
		Company TbCompany
	}
	_, _ = DB.CrudAllPagePreloadAll(User{TbUser: TbUser{Age: 18}}, &[]User{}, fire.Page{
		Current:  1,
		PageSize: 5,
	})
	assert.Equal(t, DB.Logger.(*Diary).LastSql(3), "SELECT count(*) FROM `user` WHERE `user`.`age` = 18")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `company` WHERE `company`.`id` = 1 AND `company`.`deleted_at` = 0")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `user`.`age` = 18 LIMIT 5")
}

func TestCrudCount(t *testing.T) {
	_, _ = DBDryRun.WhereRange("age", 18, 20).CrudCount(TbUser{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT count(*) FROM `user` WHERE `age` >= 18 AND `age` <= 20")
}

func TestCrudSum(t *testing.T) {
	_, _ = DBDryRun.WhereRange("age", 18, 20).CrudSum(TbUser{}, "id")
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT IFNULL(SUM(`id`),0) FROM `user` WHERE `age` >= 18 AND `age` <= 20")
}

func TestCrudUpdate(t *testing.T) {
	// simple
	_ = DBDryRun.CrudUpdate(TbUser{Name: "Mr"}, TbUser{Name: "LI"})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `name`='LI' WHERE `user`.`name` = 'Mr'")

	_ = DBDryRun.CrudUpdate(TbUser{Name: "Mr"}, TbUser{Name: "LI"}, fire.UpdateArithmetic("age", 1, fire.ArithmeticIncrease), fire.M{"tag": ""})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `age`=`age` + 1,`name`='LI',`tag`='' WHERE `user`.`name` = 'Mr'")

	// Select And Omit
	_ = fire.NewInstance(DBDryRun.Select("name", "tag")).CrudUpdate(TbUser{Name: "Mr"}, TbUser{Name: "LI"}, fire.UpdateArithmetic("age", 1, fire.ArithmeticReduce))
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `name`='LI',`tag`='' WHERE `user`.`name` = 'Mr'")

	_ = fire.NewInstance(DBDryRun.Select("*").Omit("age")).CrudUpdate(TbUser{Name: "Mr"}, TbUser{Name: "LI"}, fire.UpdateArithmetic("age", 1, fire.ArithmeticMultiply))
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `company_id`=0,`height`=0,`id`=0,`name`='LI',`tag`='' WHERE `user`.`name` = 'Mr'")

	_ = fire.NewInstance(DBDryRun.Select("name").Omit("age")).CrudUpdate(TbUser{Name: "Mr"}, TbUser{Name: "LI"}, fire.UpdateArithmetic("age", 1, fire.ArithmeticExcept))
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `name`='LI' WHERE `user`.`name` = 'Mr'")
}

func TestCrudDelete(t *testing.T) {
	_ = DBDryRun.WhereRange("age", 18, 20).CrudDelete(TbUser{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "DELETE FROM `user` WHERE `age` >= 18 AND `age` <= 20")
}
