package middleware

import (
	response "dpj-admin-api/support"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"time"
)

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Abort()
			response.WithContext(c).Error(http.StatusForbidden, "Authorization 未携带token")
			return
		}
		parseToken, err := ParseToken(authHeader)
		if err != nil {
			c.Abort()
			response.WithContext(c).Error(http.StatusForbidden, "token 验证失败")
			return
		}

		c.Set("UserId", strconv.Itoa(parseToken.Id))
		c.Next()
	}
}

type JwtStruct struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}

var Secret = []byte("secret")

func ParseToken(tokenString string) (*JwtStruct, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtStruct{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtStruct); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

const TokenExpireDuration = time.Hour * 24

func GenToken(username string, userId int) (string, error) {
	c := JwtStruct{
		Username: username,
		Id:       userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "v-v-v-v",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}
