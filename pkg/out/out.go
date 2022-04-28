package out

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

// success

func JsonSuccess(c *gin.Context) {
	JsonData(c, gin.H{
		"msg": "ok",
	})
}
func JsonData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// fail

func JsonError(c *gin.Context, msg string, data ...interface{}) {
	var d interface{}
	if data != nil {
		d = data[0]
	} else {
		d = make(map[string]interface{})
	}
	c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
		"data": d,
		"msg":  msg,
	})
}

func HandleError(c *gin.Context, err error, data ...interface{}) bool {
	if err != nil {
		JsonError(c, err.Error(), data...)
		return true
	}
	return false
}

func ParseParamID(c *gin.Context) (uint, error) {
	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return uint(parseId), nil
}

func ParseParamIDString(c *gin.Context, key ...string) string {
	if key == nil {
		return c.Param("id")
	}

	return c.Param(key[0])
}

type ParseParam struct {
	Key      string
	validate string
}

func ParseParamIDStringWithValidator(c *gin.Context, v ParseParam) (err error, res string) {
	res = ParseParamIDString(c, v.Key)
	if v.validate != "" {
		err = binding.Validator.Engine().(*validator.Validate).Var(res, v.validate)
	}

	return
}
