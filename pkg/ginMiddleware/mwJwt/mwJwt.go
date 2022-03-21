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

	TokenKey string
	AppName  string
	Secret   string
}

func NewJwtMiddleware(tokenKey, secret, appName, ginKey string, p reflect.Type, abortFunc func(*gin.Context, error)) MwJwt {
	return MwJwt{
		Middleware: func(c *gin.Context) {
			// header or query
			token := c.GetHeader(tokenKey)
			if token == "" {
				token = c.Query(tokenKey)
			}

			var err error
			defer func() {
				if err != nil {
					abortFunc(c, err)
					c.Abort()
				} else {
					c.Next()
				}
			}()

			if len(token) < bearerLength {
				err = errors.New(fmt.Sprintf("%s has not Bearer token", tokenKey))
				return
			}
			token = strings.TrimSpace(token)
			token = strings.TrimPrefix(token, tokenPrefix)
			value, err := jwt.Parse(token, appName, secret) // value is map[string]interface{}
			if err != nil {
				return
			}

			data := reflect.New(p).Elem().Interface()
			err = mapstructure.Decode(value, &data)
			if err != nil {
				return
			}
			c.Set(ginKey, data)
		},
		TokenKey: tokenKey,
		AppName:  appName,
		Secret:   secret,
	}
}

func (j *MwJwt) GenerateToken(data interface{}, expireHour int) (res *jwt.JwtObj, e error) {
	if res, e = jwt.GenerateToken(data, j.AppName, j.Secret, j.TokenKey, expireHour); e == nil {
		res.Token = tokenPrefix + res.Token
	}
	return
}
