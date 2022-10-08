package validate_test

import (
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"github.com/lifegit/go-gulu/v2/pkg/validate"
	"net/http"
	"strings"
	"testing"
)

type User struct {
	Phone string `binding:"required,phone" label:"手机号" json:"phone"`
}

func TestAlias(t *testing.T) {
	// server
	router := gin.New()
	router.POST("/", func(c *gin.Context) {
		// gin
		var param User
		errs, first := validate.Translate(c.ShouldBind(&param))
		if out.HandleError(c, first, errs) {
			return
		}

		out.JsonSuccess(c)
	})
	go router.Run(":9991")

	// client
	client := &http.Client{}
	client.Post(`http://127.0.0.1:9991`, "application/json", strings.NewReader(`{"phone":"2130000000"}`))
}
