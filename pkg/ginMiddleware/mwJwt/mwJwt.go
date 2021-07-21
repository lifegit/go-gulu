package mwJwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/jwt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strings"
)

const tokenPrefix = "Bearer "
const bearerLength = len(tokenPrefix)

type MwJwt struct {
	Middleware gin.HandlerFunc

	tokenKey string
	appName  string
	secret   string
}

func NewJwtMiddleware(tokenKey, secret, appName, ginKey string, p reflect.Type, abortFunc func(error) (code int, jsonObj interface{})) MwJwt {
	return MwJwt{
		Middleware: func(c *gin.Context) {
			// header or query
			token := c.GetHeader(tokenKey)
			if token == "" {
				token = c.Query(tokenKey)
			}
			if len(token) < bearerLength {
				c.AbortWithStatusJSON(abortFunc(errors.New(fmt.Sprintf("%s has not Bearer token", tokenKey))))
				return
			}
			token = strings.TrimSpace(token)
			token = strings.TrimPrefix(token, tokenPrefix)
			value, err := jwt.Parse(token, appName, secret) // value is map[string]interface{}
			if err != nil {
				c.AbortWithStatusJSON(abortFunc(err))
				return
			}

			data := reflect.New(p).Elem().Interface()
			if err := mapstructure.Decode(value, &data); err != nil {
				c.AbortWithStatusJSON(abortFunc(err))
				return
			}

			c.Set(ginKey, data)
			c.Next()
		},
		tokenKey: tokenKey,
		appName:  appName,
		secret:   secret,
	}
}

func (j *MwJwt) GenerateToken(data interface{}, expireHour int) (res *jwt.JwtObj, e error) {
	if res, e = jwt.GenerateToken(data, j.appName, j.secret, j.tokenKey, expireHour); e == nil {
		res.Token = tokenPrefix + res.Token
	}
	return
}
