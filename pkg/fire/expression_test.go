package fire_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
	"testing"
)

func TestSelect(t *testing.T) {
	DBDryRun.Clauses(fire.Select("id", "name", fire.COUNT{})).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT "user"."id", "user"."name", COUNT(*) FROM "user"`)

	inner := DBDryRun.Select("age").Limit(1)
	DBDryRun.Clauses(fire.Select(
		"id", "name",
		clause.Expr{
			SQL:  "(?) AS age",
			Vars: []interface{}{inner},
		},
	)).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT "user"."id", "user"."name", (SELECT age FROM " LIMIT 1) AS age FROM "user"`)
}

func TestWhereCompare(t *testing.T) {
	DBDryRun.Clause(fire.WhereCompare("age", 18, fire.CompareGte)).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" >= 18 LIMIT 1`)

	DBDryRun.Clause(fire.WhereCompare("age", 18, fire.CompareGt)).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" > 18 LIMIT 1`)

	DBDryRun.Clause(fire.WhereCompare("age", 18, fire.CompareLte)).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" <= 18 LIMIT 1`)

	DBDryRun.Clause(fire.WhereCompare("age", 18, fire.CompareLt)).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" < 18 LIMIT 1`)
}

func TestWhereJson(t *testing.T) {
	DBDryRun.Clause(fire.WhereJsonEq("name", "kr", "real_name")).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE json_extract_path_text("user"."name"::json,'real_name') = 'kr' LIMIT 1`)

	DBDryRun.Clause(fire.WhereJsonHas("name", "real_name")).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."name"::jsonb ? 'real_name' LIMIT 1`)
}

func TestWhereIn(t *testing.T) {
	DBDryRun.Clause(fire.WhereIn("age", []int{18, 19, 20})).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" IN (18,19,20) LIMIT 1`)

	DBDryRun.Clause(fire.WhereIn("age", []int{18, 19, 20}, true)).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" NOT IN (18,19,20) LIMIT 1`)
}

func TestWhereLike(t *testing.T) {
	DBDryRun.Clause(fire.WhereLike("name", "Wang")).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."name" LIKE '%Wang%' LIMIT 1`)
}

func TestWhereRange(t *testing.T) {
	DBDryRun.Clause(fire.WhereRange("age", 10, 20)).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" >= 10 AND "user"."age" <= 20 LIMIT 1`)
}

func TestOrderByColumn(t *testing.T) {
	DBDryRun.OrderByColumn("age1", fire.OrderDesc).OrderByColumn("age", fire.OrderAsc).Model(User{}).Find(&[]User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" ORDER BY "user"."age"`)
}

func TestUpdateArithmetic(t *testing.T) {
	DBDryRun.Model(User{}).Where(User{ID: 1}).Updates(fire.UpdateArithmetic("age", 2, fire.ArithmeticIncrease))
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `UPDATE "user" SET "age"="age" + 2 WHERE "user"."id" = 1`)
}
