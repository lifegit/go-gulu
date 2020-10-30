package mwJwt

import (
	"github.com/gin-gonic/gin"
	"go-gulu/jwt"
	"net/http"
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

func NewJwtMiddleware(tokenKey, appName, secret string, ginKey string) MwJwt {
	return MwJwt{
		Middleware: func(c *gin.Context) {
			//c.Next()
			//return
			//source header or query
			hToken := c.GetHeader(tokenKey)
			if hToken == "" {
				hToken = c.Query(tokenKey)
			}
			if len(hToken) < bearerLength {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": tokenKey + " has not Bearer token"})
				return
			}
			token := strings.TrimSpace(hToken)
			user, err := jwt.Parse(token, appName, secret)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
				return
			}

			c.Set(ginKey, user)

			c.Next()
			// after request
		},
		tokenKey: tokenKey,
		appName:  appName,
		secret:   secret,
	}
}

func (j *MwJwt) GenerateToken(data interface{}, expireHour int) (*jwt.JwtObj, error) {
	return jwt.GenerateToken(data, j.appName, j.secret, j.tokenKey, expireHour)
}
