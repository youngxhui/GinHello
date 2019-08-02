package middleware

import (
	"GinHello/config"
	"GinHello/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := model.Result{
			Code:    http.StatusUnauthorized,
			Message: "无法认证，重新登录",
			Data:    nil,
		}
		auth := context.Request.Header.Get("Authorization")
		auth = strings.Fields(auth)[1]
		println("index is ", auth)
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		}
		// 校验token
		claims, err := parseToken(auth)
		if err != nil {
			println("解析token失败")
		} else if time.Now().Unix() > claims.ExpiresAt {
			println("token 过期")
		} else {
			println("token 正确")
		}
		context.Next()
	}
}
func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err

}
