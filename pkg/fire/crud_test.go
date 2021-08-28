package fire_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrudCreate(t *testing.T) {
	// fail
	err := DBDryRun.CrudCreate(User{
		CompanyID: 1, Tag: "student", Age: 18, Height: 185,
	})
	assert.Error(t, err)

	err = DBDryRun.CrudCreate([]User{
		{CompanyID: 1, Name: "san", Tag: "student", Height: 185},
		{CompanyID: 2, Tag: "student", Age: 19},
	})
	assert.Error(t, err)

	// success
	_ = DBDryRun.CrudCreate(User{
		CompanyID: 1, Name: "san", Tag: "student", Age: 18, Height: 185,
	})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "INSERT INTO `user` (`company_id`,`name`,`tag`,`age`,`height`) VALUES (1,'san','student',18,185)")

	_ = DBDryRun.CrudCreate([]User{
		{CompanyID: 1, Name: "san", Tag: "student", Height: 185},
		{CompanyID: 2, Name: "sid", Tag: "student", Age: 19},
	}, 2)
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "INSERT INTO `user` (`company_id`,`name`,`tag`,`age`,`height`) VALUES (1,'san','student',0,185),(2,'sid','student',19,0)")
}

func TestIsExists(t *testing.T) {
	DBDryRun.IsExists(User{ID: 1, Name: "Mr Wang"})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT 1 FROM `user` WHERE `user`.`id` = 1 AND `user`.`name` = 'Mr Wang' LIMIT 1")
}

func TestCrudOne(t *testing.T) {
	_ = DBDryRun.WhereLike("name", "M").CrudOne(User{ID: 1, Age: 20}, &User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `name` LIKE '%M%' AND `user`.`id` = 1 AND `user`.`age` = 20 LIMIT 1")
}

func TestCrudAll(t *testing.T) {
	_ = DBDryRun.WhereRange("age", 18, 20).CrudAll(User{}, &[]User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` >= 18 AND `age` <= 20")
}

func TestCrudAllPage(t *testing.T) {
	_, _ = DBDryRun.CrudAllPage(User{}, &[]User{}, fire.Page{
		Current:  3,
		PageSize: 5,
	})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(2), "SELECT count(*) FROM `user`")
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` LIMIT 5 OFFSET 10")
}

func TestCrudOnePreloadJoin(t *testing.T) {
	type TbUser struct {
		User
		Company Company
	}
	_ = DBDryRun.CrudOnePreloadJoin(TbUser{}, &TbUser{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT `user`.`id`,`user`.`company_id`,`user`.`name`,`user`.`tag`,`user`.`age`,`user`.`height`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`id` AS `Company__id`,`Company`.`address` AS `Company__address`,`Company`.`name` AS `Company__name` FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id` LIMIT 1")
}

func TestCrudAllPreloadJoin(t *testing.T) {
	type TbUser struct {
		User
		Company Company
	}
	_ = DBDryRun.CrudAllPreloadJoin(TbUser{}, &[]TbUser{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT `user`.`id`,`user`.`company_id`,`user`.`name`,`user`.`tag`,`user`.`age`,`user`.`height`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`id` AS `Company__id`,`Company`.`address` AS `Company__address`,`Company`.`name` AS `Company__name` FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id`")
}

func TestCrudAllPagePreloadJoin(t *testing.T) {
	type TbUser struct {
		User
		Company Company
	}
	_, _ = DBDryRun.CrudAllPagePreloadJoin(TbUser{}, &[]TbUser{}, fire.Page{
		Current:  3,
		PageSize: 5,
	})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(2), "SELECT count(*) FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id`")
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT `user`.`id`,`user`.`company_id`,`user`.`name`,`user`.`tag`,`user`.`age`,`user`.`height`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`id` AS `Company__id`,`Company`.`address` AS `Company__address`,`Company`.`name` AS `Company__name` FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id` LIMIT 5 OFFSET 10")
}

func TestCrudOnePreloadAll(t *testing.T) {
	type TbUser struct {
		User
		Company Company
	}
	_ = DB.CrudOnePreloadAll(TbUser{User: User{Age: 18}}, &TbUser{})
	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `company` WHERE `company`.`id` = 1 AND `company`.`deleted_at` = 0")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `user`.`age` = 18 LIMIT 1")
}

func TestCrudAllPreloadAll(t *testing.T) {
	type TbUser struct {
		User
		Company Company
	}
	_ = DB.CrudAllPreloadAll(TbUser{}, &[]TbUser{})
	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `company` WHERE `company`.`id` IN (1,2) AND `company`.`deleted_at` = 0")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user`")
}

func TestCrudAllPagePreloadAll(t *testing.T) {
	type TbUser struct {
		User
		Company Company
	}
	_, _ = DB.CrudAllPagePreloadAll(TbUser{User: User{Age: 18}}, &[]TbUser{}, fire.Page{
		Current:  1,
		PageSize: 5,
	})
	assert.Equal(t, DB.Logger.(*Diary).LastSql(3), "SELECT count(*) FROM `user` WHERE `user`.`age` = 18")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `company` WHERE `company`.`id` = 1 AND `company`.`deleted_at` = 0")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `user`.`age` = 18 LIMIT 5")
}

func TestCrudCount(t *testing.T) {
	_, _ = DBDryRun.WhereRange("age", 18, 20).CrudCount(User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT count(*) FROM `user` WHERE `age` >= 18 AND `age` <= 20")
}

func TestCrudSum(t *testing.T) {
	_, _ = DBDryRun.WhereRange("age", 18, 20).CrudSum(User{}, "id")
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT IFNULL(SUM(`id`),0) FROM `user` WHERE `age` >= 18 AND `age` <= 20")
}

func TestCrudUpdate(t *testing.T) {
	// simple
	_ = DBDryRun.CrudUpdate(User{Name: "Mr"}, User{Name: "LI"})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `name`='LI' WHERE `user`.`name` = 'Mr'")

	_ = DBDryRun.CrudUpdate(User{Name: "Mr"}, User{Name: "LI"}, fire.UpdateArithmetic("age", 1, fire.ArithmeticIncrease), fire.M{"tag": ""})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `age`=`age` + 1,`name`='LI',`tag`='' WHERE `user`.`name` = 'Mr'")

	// Select And Omit
	_ = fire.NewInstance(DBDryRun.Select("name", "tag")).CrudUpdate(User{Name: "Mr"}, User{Name: "LI"}, fire.UpdateArithmetic("age", 1, fire.ArithmeticReduce))
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `name`='LI',`tag`='' WHERE `user`.`name` = 'Mr'")

	_ = fire.NewInstance(DBDryRun.Select("*").Omit("age")).CrudUpdate(User{Name: "Mr"}, User{Name: "LI"}, fire.UpdateArithmetic("age", 1, fire.ArithmeticMultiply))
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `company_id`=0,`height`=0,`id`=0,`name`='LI',`tag`='' WHERE `user`.`name` = 'Mr'")

	_ = fire.NewInstance(DBDryRun.Select("name").Omit("age")).CrudUpdate(User{Name: "Mr"}, User{Name: "LI"}, fire.UpdateArithmetic("age", 1, fire.ArithmeticExcept))
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `name`='LI' WHERE `user`.`name` = 'Mr'")
}

func TestCrudUpdatePrimaryKey(t *testing.T) {
	// fail
	err := DBDryRun.CrudUpdatePrimaryKey(User{Name: "Mr"}, User{Name: "LI"})
	assert.Error(t, err) // primary key ID is zero

	// success
	_ = DBDryRun.CrudUpdatePrimaryKey(User{ID: 1}, User{Name: "LI"})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `name`='LI' WHERE `user`.`id` = 1 AND `id` = 1 LIMIT 1")
}

func TestCrudDelete(t *testing.T) {
	_ = DBDryRun.WhereRange("age", 18, 20).CrudDelete(User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "DELETE FROM `user` WHERE `age` >= 18 AND `age` <= 20")
}
