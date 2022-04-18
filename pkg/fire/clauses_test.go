package fire_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
	"testing"
)

func TestCount(t *testing.T) {
	DBDryRun.Select("?", fire.COUNT{}).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT COUNT(*) FROM "user"`)

	DBDryRun.Select("?", fire.COUNT{Column: clause.Column{Name: "age"}}).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT COUNT("age") FROM "user"`)

	DBDryRun.Select("?", fire.COUNT{Column: clause.Column{Name: "age", Alias: "count_age"}}).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT COUNT("age")AS "count_age" FROM "user"`)

	DBDryRun.Clauses(fire.Select("id", "name", fire.COUNT{})).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT "user"."id", "user"."name", COUNT(*) FROM "user"`)
}

func TestSum(t *testing.T) {
	DBDryRun.Select("?", fire.SUM{Column: clause.Column{Name: "age"}}).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT COALESCE(SUM("age"),0) FROM "user"`)

	DBDryRun.Select("?", fire.SUM{Column: clause.Column{Name: "age", Alias: "count_age"}}).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT COALESCE(SUM("age"),0)AS "count_age" FROM "user"`)

	DBDryRun.Clauses(fire.Select("id", "name", fire.SUM{})).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT "user"."id", "user"."name", COALESCE(SUM("),0) FROM "user"`)
}

func TestUnnest(t *testing.T) {
	DBDryRun.Select("?", fire.UNNEST{Column: clause.Column{Name: "age"}}).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT UNNEST("age") FROM "user"`)

	DBDryRun.Select("?", fire.UNNEST{Column: clause.Column{Name: "age", Alias: "count_age"}}).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT UNNEST("age")AS "count_age" FROM "user"`)

	DBDryRun.Clauses(fire.Select("id", "name", fire.UNNEST{})).Find(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT "user"."id", "user"."name", UNNEST(") FROM "user"`)
}

func TestArithmetic(t *testing.T) {
	DBDryRun.Model(User{}).Updates(map[string]interface{}{
		"age": fire.SetArithmetic{Column: clause.Column{Table: clause.CurrentTable, Name: "age"}, Type: fire.ArithmeticMultiply, Value: 33},
	})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `UPDATE "user" SET "age"="user"."age" * 33`)
}
