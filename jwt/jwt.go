/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	User interface{}
	jwt.StandardClaims
}

type JwtObj struct {
	Key      string  `json:"key"`
	Token    string  `json:"token"`
	ExpireTs int64   `json:"expire"`
	Alive    float64 `json:"alive"`
}

// https://studygolang.com/articles/13062

// GenerateToken generate tokens used for auth
func GenerateToken(m interface{}, appName string, appSecret string, key string, expireHour int) (*JwtObj, error) {
	expireAfterTime := time.Hour * time.Duration(expireHour)
	expireTime := time.Now().Add(expireAfterTime)
	stdClaims := Claims{
		m,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    appName,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := tokenClaims.SignedString([]byte(appSecret))
	if err != nil {
		return nil, errors.New("config is wrong, can not generate jwt")
	}
	// fmt.Println("tokenString" , tokenString)
	data := &JwtObj{Key: key, Token: tokenString, ExpireTs: expireTime.Unix(), Alive: expireAfterTime.Seconds()}
	return data, err
}

// JwtParseUser parsing token
func Parse(tokenString string, appName string, appSecret string) (interface{}, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := Claims{}
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(appSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return nil, errors.New("token is expired")
	}

	if !claims.VerifyIssuer(appName, true) {
		return nil, errors.New("token's issuer is wrong,greetings Hacker")
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return &claims.User, nil
	}

	return nil, err
}
