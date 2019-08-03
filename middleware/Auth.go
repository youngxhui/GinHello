package middleware

import (
	"GinHello/config"
	"GinHello/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := model.Result{
			Code:    http.StatusUnauthorized,
			Message: "无法认证，重新登录",
			Data:    nil,
		}
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		_, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result.Message = "token 过期" + err.Error()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
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
