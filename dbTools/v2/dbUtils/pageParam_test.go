/**
* @Author: TheLife
* @Date: 2021/5/27 上午10:31
 */
package dbUtils_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gulu/dbTools/v2/dbUtils"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

var db *gorm.DB

func TestPageParam(t *testing.T) {
	db = initMysqlDb()

	// server
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {

		var param dbUtils.PageParam
		err := c.ShouldBind(&param)
		if err != nil {
			fmt.Println("ShouldBind: ", err)
			return
		}

		user, userList := &TbUser{DbUtils: &dbUtils.DbUtils{DB: db}}, &[]TbUser{}
		count, err := user.DbUtils.CrudAllPage(user, userList, param, dbUtils.Allow{
			Where: []string{"age"},
			Like:  []string{"name"},
			Range: []string{"time"},
			In:    []string{"tags"},
		})
		fmt.Println(count, err)

	})
	go router.Run(":8979")

	// client
	client := &http.Client{}
	_, _ = client.Get(`http://127.0.0.1:8979/?params={"time":[1622173028456,1622259431456],"name":"张明","age":18,"tag":["学生","儿子","青年"]}&sort={"id":"ascend"}`)
}
