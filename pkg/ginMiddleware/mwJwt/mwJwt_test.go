package mwJwt_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/ginMiddleware/mwJwt"
	"github.com/lifegit/go-gulu/v2/pkg/jwt"
	"github.com/lifegit/go-gulu/v2/pkg/out"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

type Name struct {
	Id       uint
	Username string
}

func Get(url string) (res string) {
	client := &http.Client{}
	response, err := client.Get(url)
	if err == nil {
		bytes, _ := ioutil.ReadAll(response.Body)
		res = string(bytes)
	}

	return
}

func TestName(t *testing.T) {
	g := gin.New()
	tokenKey, ginKey := "Authorization", "token"
	tokenObj, tokenVal := &jwt.JwtObj{}, Name{
		Id:       1,
		Username: "he",
	}
	jwtM := mwJwt.NewJwtMiddleware(tokenKey, "n&jL2q8QrxEsSPxx7$JZiTD3A3vr.bCq", "app", ginKey, reflect.TypeOf(Name{}), func(e error) (code int, jsonObj interface{}) {
		return http.StatusUnauthorized, gin.H{"msg": e.Error()}
	})
	go g.Run(":8009")

	func() {
		g.GET("generate", func(c *gin.Context) {
			tokenObj, _ = jwtM.GenerateToken(tokenVal, 2400)
			out.JsonData(c, gin.H{tokenKey: tokenObj})
		})
		Get(fmt.Sprintf(`http://127.0.0.1:8009/generate`))
	}()

	func() {
		g.Use(jwtM.Middleware).GET("verify", func(c *gin.Context) {
			name := c.MustGet(ginKey).(Name)
			assert.Equal(t, name, tokenVal)
		})
		Get(fmt.Sprintf(`http://127.0.0.1:8009/verify?%s=%s`, tokenKey, url.QueryEscape(tokenObj.Token)))
	}()
}
