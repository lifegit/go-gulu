/**
* @Author: TheLife
* @Date: 2021/7/3 下午2:52
 */
package fire_test

import (
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// 对于常见的分页示例

// 分页-单表
func TestPageParamAllowCrudAllPage(t *testing.T) {
	// server
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		// gin
		var param fire.PageParam
		err := c.ShouldBind(&param)
		if out.HandleError(c, err) {
			return
		}

		userList := &[]TbUser{}
		pageResult, err := DBDryRun.Allow(param.Param, fire.Allow{
			Where: []string{"age"},
			Like:  []string{"name"},
			Range: []string{"height"},
			In:    []string{"tag"},
			Sorts: []string{"age"},
		}).CrudAllPage(TbUser{CompanyID: 1}, userList, param.Page)
		assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(2), "SELECT count(*) FROM `user` WHERE `age` = 18.000000 AND (`height` >= 160.000000 AND `height` <= 190.000000) AND `tag`  IN ('学生','儿子','青年') AND `name` LIKE '%Wang%' AND `user`.`company_id` = 1")
		assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` = 18.000000 AND (`height` >= 160.000000 AND `height` <= 190.000000) AND `tag`  IN ('学生','儿子','青年') AND `name` LIKE '%Wang%' AND `user`.`company_id` = 1 ORDER BY `age` asc LIMIT 5 OFFSET 10")
		if out.HandleError(c, err) {
			return
		}

		out.JsonPageResult(c, userList, pageResult)
	})
	go router.Run(":9991")

	// client
	client := &http.Client{}
	_, _ = client.Get(`http://127.0.0.1:9991/?current=3&page_size=5&params={"height":[160,190],"name":"Wang","age":18,"tag":["学生","儿子","青年"]}&sort={"age":"ascend"}`)
}

// 分页-外建（join）
func TestPageParamAllowPreloadJoin(t *testing.T) {
	// server
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		// gin
		var param fire.PageParam
		err := c.ShouldBind(&param)
		if out.HandleError(c, err) {
			return
		}

		type User struct {
			TbUser
			Company TbCompany
		}

		userList := &[]User{}
		pageResult, err := DBDryRun.Allow(param.Param, fire.Allow{
			Where: []string{"age"},
			Like:  []string{"user.name", "company.name"},
			Range: []string{"height"},
			In:    []string{"tag"},
			Sorts: []string{"age"},
		}).CrudAllPagePreloadJoin(User{}, userList, param.Page)
		assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(2), "SELECT count(*) FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id` WHERE `age` = 18.000000 AND (`height` >= 160.000000 AND `height` <= 190.000000) AND `tag`  IN ('student','儿子','青年') AND `user`.`name` LIKE '%Wang%' AND `company`.`name` LIKE '%Shanghai%'")
		assert.Equal(t, DBDryRun.Logger.(*Diary).LastSql(), "SELECT `user`.`id`,`user`.`company_id`,`user`.`name`,`user`.`tag`,`user`.`age`,`user`.`height`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`id` AS `Company__id`,`Company`.`address` AS `Company__address`,`Company`.`name` AS `Company__name` FROM `user` LEFT JOIN `company` `Company` ON `user`.`company_id` = `Company`.`id` WHERE `age` = 18.000000 AND (`height` >= 160.000000 AND `height` <= 190.000000) AND `tag`  IN ('student','儿子','青年') AND `user`.`name` LIKE '%Wang%' AND `company`.`name` LIKE '%Shanghai%' ORDER BY `age` asc LIMIT 5 OFFSET 10")
		if out.HandleError(c, err) {
			return
		}

		out.JsonPageResult(c, userList, pageResult)
	})
	go router.Run(":9992")

	// client
	client := &http.Client{}
	_, _ = client.Get(`http://127.0.0.1:9992/?current=3&page_size=5&params={"height":[160,190],"user.name":"Wang","company.name":"Shanghai","age":18,"tag":["student","儿子","青年"]}&sort={"age":"ascend"}`)
}

//分页-外键（多次查询）
func TestPageParamAllowPreloadAll(t *testing.T) {
	// server
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		// gin
		var param fire.PageParam
		err := c.ShouldBind(&param)
		if out.HandleError(c, err) {
			return
		}

		type User struct {
			TbUser
			Company TbCompany
		}

		userList := &[]User{}
		pageResult, err := DB.Allow(param.Param, fire.Allow{
			Where: []string{"age"},
			Like:  []string{"user.name", "company.name"}, // company.name not support
			Range: []string{"height"},
			In:    []string{"tag"},
			Sorts: []string{"age"},
		}).CrudAllPagePreloadAll(User{}, userList, param.Page)
		assert.Equal(t, DB.Logger.(*Diary).LastSql(3), "SELECT count(*) FROM `user` WHERE `age` = 18.000000 AND (`height` >= 160.000000 AND `height` <= 190.000000) AND `tag`  IN ('student','儿子','青年') AND `user`.`name` LIKE '%Wang%'")
		assert.Equal(t, DB.Logger.(*Diary).LastSql(2), "SELECT * FROM `company` WHERE `company`.`id` = 1 AND `company`.`deleted_at` = 0")
		assert.Equal(t, DB.Logger.(*Diary).LastSql(), "SELECT * FROM `user` WHERE `age` = 18.000000 AND (`height` >= 160.000000 AND `height` <= 190.000000) AND `tag`  IN ('student','儿子','青年') AND `user`.`name` LIKE '%Wang%' ORDER BY `age` asc LIMIT 20")
		if out.HandleError(c, err) {
			return
		}

		out.JsonPageResult(c, userList, pageResult)
	})
	go router.Run(":9993")

	// client
	client := &http.Client{}
	_, _ = client.Get(`http://127.0.0.1:9993/?params={"height":[160,190],"user.name":"Wang","age":18,"tag":["student","儿子","青年"]}&sort={"age":"ascend"}`)
}
