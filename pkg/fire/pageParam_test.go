// Package fire /***/
// 对于常见的分页示例

package fire_test

import (
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
	"net/http"
	"testing"
)

// 分页-单表
func TestPageParamAllowCrudAllPage(t *testing.T) {
	// server
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		pageResult, err := DBDryRun.
			Allow(c.Request.URL.Query(), fire.Allow{
				Where: fire.Filtered{"age", "age2"}, // age2 not where
				Like:  fire.Filtered{"name"},
				Range: fire.Filtered{"height"},
				In:    fire.Filtered{"tag"},
				Sorts: fire.Filtered{"age"},
			}).
			CrudAllPage(User{CompanyID: 1}, &[]User{}, c.Request.URL)
		assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(2), `SELECT count(*) FROM "user" WHERE "user"."age" = '18' AND "user"."height" >= '160' AND "user"."height" <= '190' AND "user"."tag" IN ('student','youth') AND "user"."name" LIKE '%Wang%' AND "user"."company_id" = 1`)
		assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" = '18' AND "user"."height" >= '160' AND "user"."height" <= '190' AND "user"."tag" IN ('student','youth') AND "user"."name" LIKE '%Wang%' AND "user"."company_id" = 1 ORDER BY "user"."age" LIMIT 5 OFFSET 10`)
		if out.HandleError(c, err) {
			return
		}

		out.JsonData(c, pageResult)
	})
	go router.Run(":9991")

	// client
	client := &http.Client{}
	_, _ = client.Get(`http://127.0.0.1:9991/?current=3&pageSize=5&height=160&height=190&name=Wang&age=18&tag=student&tag=youth&sort={"age":"ascend"}`)
}

// 分页-外建（join）
func TestPageParamAllowPreloadJoin(t *testing.T) {
	// server
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		// gin
		var param fire.Page
		err := c.ShouldBind(&param)
		if out.HandleError(c, err) {
			return
		}

		type TbUser struct {
			User
			Company Company
		}

		pageResult, err := DBDryRun.
			Allow(c.Request.URL.Query(), fire.Allow{
				Where: fire.Filtered{"age"},
				Like:  fire.Filtered{"name", clause.Column{Table: "company", Name: "name", Alias: "companyName"}},
				Range: fire.Filtered{"height"},
				In:    fire.Filtered{"tag"},
				Sorts: fire.Filtered{"age"},
			}).
			CrudAllPagePreloadJoin(TbUser{}, &[]TbUser{}, param)
		assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(2), `SELECT count(*) FROM "user" LEFT JOIN "company" "Company" ON "user"."company_id" = "Company"."id" WHERE "user"."age" = '18' AND "user"."height" >= '160' AND "user"."height" <= '190' AND "user"."tag" IN ('student','youth') AND "user"."name" LIKE '%Wang%' AND "company"."name" LIKE '%Shanghai%'`)
		assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT "user"."id","user"."company_id","user"."name","user"."tag","user"."age","user"."height","Company"."created_at" AS "Company__created_at","Company"."updated_at" AS "Company__updated_at","Company"."deleted_at" AS "Company__deleted_at","Company"."id" AS "Company__id","Company"."address" AS "Company__address","Company"."name" AS "Company__name" FROM "user" LEFT JOIN "company" "Company" ON "user"."company_id" = "Company"."id" WHERE "user"."age" = '18' AND "user"."height" >= '160' AND "user"."height" <= '190' AND "user"."tag" IN ('student','youth') AND "user"."name" LIKE '%Wang%' AND "company"."name" LIKE '%Shanghai%' ORDER BY "user"."age" LIMIT 5 OFFSET 10`)
		if out.HandleError(c, err) {
			return
		}

		out.JsonData(c, pageResult)
	})
	go router.Run(":9992")

	// client
	client := &http.Client{}
	_, _ = client.Get(`http://127.0.0.1:9992/?current=3&pageSize=5&height=160&height=190&name=Wang&companyName=Shanghai&age=18&tag=student&tag=youth&sort={"age":"ascend"}`)
}

//分页-外键（多次查询）
func TestPageParamAllowPreloadAll(t *testing.T) {
	// server
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		// gin
		var param fire.Page
		err := c.ShouldBind(&param)
		if out.HandleError(c, err) {
			return
		}

		type TbUser struct {
			User
			Company Company
		}

		pageResult, err := DB.
			Allow(c.Request.URL.Query(), fire.Allow{
				Where: fire.Filtered{"age"},
				Like:  fire.Filtered{"name", clause.Column{Table: "company", Name: "name", Alias: "companyName"}}, // company.name not support
				Range: fire.Filtered{"height"},
				In:    fire.Filtered{"tag"},
				Sorts: fire.Filtered{"age"},
			}).
			CrudAllPagePreloadAll(TbUser{}, &[]TbUser{}, param)
		assert.Equal(t, DB.Logger.(*fire.Diary).LastSql(3), `SELECT count(*) FROM "user" WHERE "user"."age" = '18' AND "user"."height" >= '160' AND "user"."height" <= '190' AND "user"."tag" IN ('student','youth') AND "user"."name" LIKE '%Wang%'`)
		assert.Equal(t, DB.Logger.(*fire.Diary).LastSql(2), `SELECT * FROM "company" WHERE "company"."id" = 1 AND "company"."deleted_at" = 0`)
		assert.Equal(t, DB.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" = '18' AND "user"."height" >= '160' AND "user"."height" <= '190' AND "user"."tag" IN ('student','youth') AND "user"."name" LIKE '%Wang%' ORDER BY "user"."age" LIMIT 30`)
		if out.HandleError(c, err) {
			return
		}

		out.JsonData(c, pageResult)
	})
	go router.Run(":9993")

	// client
	client := &http.Client{}
	_, _ = client.Get(`http://127.0.0.1:9993/?pageSize=30&height=160&height=190&name=Wang&age=18&tag=student&tag=youth&sort={"age":"ascend"}`)
}
