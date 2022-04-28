package translate_test

import (
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"github.com/lifegit/go-gulu/v2/pkg/translate"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

type User struct {
	Computer Computer `json:"computer"`
	ID       int64    `binding:"required" label:"id" json:"id"`
	Name     string   `binding:"required" label:"姓名" json:"name"`
	Phone    string   `binding:"required,numeric,len=11,startswith=1" label:"手机号" json:"phone"`
}

type Computer struct {
	System System `json:"system"`
	Price  int64  `binding:"required" label:"价格" json:"price"`
}

type System struct {
	Os      string `binding:"required,oneof=mac linux" label:"操作系统" json:"os"`
	Version int    `binding:"required,max=10" label:"版本" json:"version"`
}

func TestGin(t *testing.T) {
	// server
	router := gin.New()
	router.POST("/", func(c *gin.Context) {
		// gin
		var param User
		errs, first := translate.Translate(c.ShouldBind(&param))
		if out.HandleError(c, first, errs) {
			return
		}

		out.JsonSuccess(c)
	})
	go router.Run(":9993")

	// client
	client := &http.Client{}
	res, _ := client.Post(`http://127.0.0.1:9993`, "application/json", strings.NewReader(`{"id":123,"phone":"2130000000","computer":{"system":{"os":"win","version":11}}}`))
	robots, _ := io.ReadAll(res.Body)

	assert.Equal(t, string(robots), `{"data":{"computer.price":"价格为必填字段","computer.system.os":"操作系统必须是[mac linux]中的一个","computer.system.version":"版本必须小于或等于10","name":"姓名为必填字段","phone":"手机号长度必须是11个字符"},"msg":"操作系统必须是[mac linux]中的一个"}`)
}
