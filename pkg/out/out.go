/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package out

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//type Result struct {
//	Code ErrCode     `json:"code"`
//	Msg  string      `json:"msg"`
//	Data interface{} `json:"data"`
//}
//
//// success
//func JsonSuccess(c *gin.Context) {
//	JsonData(c, gin.H{})
//}
//func JsonData(c *gin.Context, data interface{}) {
//	c.JSON(http.StatusOK, Result{
//		Code: 200,
//		Msg:  "ok",
//		Data: data,
//	})
//}
//
//// fail
//type ErrCode int
//
//func JsonError(c *gin.Context, msg string, code ...ErrCode) {
//	JsonErrorData(c, gin.H{}, msg, code...)
//}
//func JsonErrorData(c *gin.Context, data interface{}, msg string, code ...ErrCode) {
//	o := ErrCode(-1)
//	if code != nil {
//		o = code[0]
//	}
//	c.AbortWithStatusJSON(http.StatusOK, Result{
//		Code: o,
//		Msg:  msg,
//		Data: data,
//	})
//}

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
type ErrCode int

func JsonError(c *gin.Context, msg string) {
	JsonErrorData(c, nil, msg)
}
func JsonErrorData(c *gin.Context, data interface{}, msg string) {
	c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
		"data": data,
		"msg":  msg,
	})
}

func HandleError(c *gin.Context, err error) bool {
	if err != nil {
		JsonError(c, err.Error())
		return true
	}
	return false
}

func HandleErrorData(c *gin.Context, data interface{}, err error) bool {
	if err != nil {
		JsonErrorData(c, data, err.Error())
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
