package fire_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatColumn(t *testing.T) {
	var res, success string

	// column
	success = "`a1`"
	res = fire.FormatColumn("a1")
	assert.Equal(t, res, success)
	res = fire.FormatColumn("`a1`")
	assert.Equal(t, res, success)

	// table column
	success = "`table`.`a1`"
	res = fire.FormatColumn("table.a1")
	assert.Equal(t, res, success)
	res = fire.FormatColumn("`table`.`a1`")
	assert.Equal(t, res, success)
	res = fire.FormatColumn("`table`.a1")
	assert.Equal(t, res, success)
	res = fire.FormatColumn("table.`a1`")
	assert.Equal(t, res, success)
}

func TestWhereCompare(t *testing.T) {
	DBDryRun.WhereCompare("age", 18, fire.CompareAboutEqual).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` >= 18 LIMIT 1")

	DBDryRun.WhereCompare("age", 18, fire.CompareAbout).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` > 18 LIMIT 1")

	DBDryRun.WhereCompare("age", 18, fire.CompareLessEqual).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` <= 18 LIMIT 1")

	DBDryRun.WhereCompare("age", 18, fire.CompareLess).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` < 18 LIMIT 1")
}

func TestIn(t *testing.T) {
	DBDryRun.WhereIn("age", []int{18, 19, 20}).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age`  IN (18,19,20) LIMIT 1")

	DBDryRun.WhereIn("age", []int{18, 19, 20}, true).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` NOT IN (18,19,20) LIMIT 1")
}

func TestLike(t *testing.T) {
	DBDryRun.WhereLike("name", "Wang").Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `name` LIKE '%Wang%' LIMIT 1")
}

func TestRange(t *testing.T) {
	DBDryRun.WhereRange("age", 10, 20).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` >= 10 AND `age` <= 20 LIMIT 1")
}

func TestUpdateArithmetic(t *testing.T) {
	DBDryRun.Model(User{}).Where(User{ID: 1}).Updates(fire.UpdateArithmetic("age", 2, fire.ArithmeticIncrease))
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "UPDATE `user` SET `age`=`age` + 2 WHERE `user`.`id` = 1")
}

func TestOrderByColumn(t *testing.T) {
	DBDryRun.OrderByColumn("age", fire.OrderAsc).Model(User{}).Find(&[]User{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` ORDER BY `age` asc")
}

func TestPreloadJoin(t *testing.T) {
	type TbUser struct {
		User
		Company Company
	}
	DBDryRun.PreloadJoin(TbUser{}).WhereCompare("user.age", 18).Find(&[]TbUser{})
	assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT `user`.`id`,`user`.`company_id`,`user`.`name`,`user`.`tag`,`user`.`age`,`user`.`height`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`id` AS `Company__id`,`Company`.`address` AS `Company__address`,`Company`.`name` AS `Company__name` FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id` WHERE `user`.`age` = 18")
}

// 分页属于(Belongs To)
// user.company_id -> company.id
// 一个用户属于一个公司
// https://gorm.io/zh_CN/docs/belongs_to.html
func TestAssociationsBelongsTo(t *testing.T) {
	type TbUser struct {
		User
		Company Company
	}
	res := &[]TbUser{}
	DB.PreloadAll().Find(res)

	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `company` WHERE `company`.`id` IN (1,2) AND `company`.`deleted_at` = 0")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user`")
}

// 分页一对一(Has One)
// user.id -> card.user_id
// 一个用户有一张唱片
// https://gorm.io/zh_CN/docs/has_one.html
func TestAssociationsHasOne(t *testing.T) {
	type TbUser struct {
		User
		Card Card
	}
	res := &[]TbUser{}
	DB.PreloadAll().Find(res)

	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `card` WHERE `card`.`user_id` IN (1,2,3,4)")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user`")
}

// 分页一对多
// user.id -> []card.user_id
// 一个用户有多张唱片
// https://gorm.io/zh_CN/docs/has_many.html
func TestAssociationsHasMany(t *testing.T) {
	type TbUser struct {
		User
		Card []Card
	}
	res := &[]TbUser{}
	DB.PreloadAll().Find(res)

	assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `card` WHERE `card`.`user_id` IN (1,2,3,4)")
	assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user`")
}

// 分页多对多,`user_languages` 是连接表
// https://gorm.io/zh_CN/docs/many_to_many.html
func TestAssociationsManyToMany(t *testing.T) {
	type TbLanguage struct {
		Language
		Users []*User `gorm:"many2many:user_languages;"`
	}
	type TbUser struct {
		User
		Languages []*Language `gorm:"many2many:user_languages;"`
	}

	// 正向: 一个人会多种语言 user <- []language
	// user.id -> user_languages.user_id,[]language_id -> language.id
	func() {
		res := &[]TbUser{}
		DB.PreloadAll().Find(res)
		assert.Equal(t, DB.Logger.(*Diary).LastSql(3), "SELECT * FROM `user_languages` WHERE `user_languages`.`user_id` IN (1,2,3,4)")
		assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `language` WHERE `language`.`id` IN (1,2,3)")
		assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user`")
	}()

	// 反向: 一群人会一种语言 language <- []user
	// language.id -> user_languages.language_id,[]user_id -> user.id
	func() {
		res := &[]TbLanguage{}
		DB.PreloadAll().Find(res)
		assert.Equal(t, DB.Logger.(*Diary).LastSql(3), "SELECT * FROM `user_languages` WHERE `user_languages`.`language_id` IN (1,2,3)")
		assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `user` WHERE `user`.`id` IN (1,2)")
		assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `language`")
	}()
}
