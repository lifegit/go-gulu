package validate_test

import (
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"github.com/lifegit/go-gulu/v2/pkg/validate"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

type Staff struct {
	Computer Computer `json:"computer"`
	ID       int64    `binding:"required" label:"id" json:"id"`
	Name     string   `binding:"required" label:"姓名" json:"name"`
	JobNum   string   `binding:"required,len=8,startswith=TD" label:"工号" json:"job"`
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
		var param Staff
		errs, first := validate.Translate(c.ShouldBind(&param))
		if out.HandleError(c, first, errs) {
			return
		}

		out.JsonSuccess(c)
	})
	go router.Run(":9993")

	// client
	client := &http.Client{}
	res, _ := client.Post(`http://127.0.0.1:9993`, "application/json", strings.NewReader(`{"id":123,"job":"1TD000002","computer":{"system":{"os":"win","version":11}}}`))
	robots, _ := io.ReadAll(res.Body)

	assert.Equal(t, string(robots), `{"data":{"computer.price":"价格为必填字段","computer.system.os":"操作系统必须是[mac linux]中的一个","computer.system.version":"版本必须小于或等于10","job":"工号长度必须是8个字符","name":"姓名为必填字段"},"msg":"操作系统必须是[mac linux]中的一个"}`)
}
