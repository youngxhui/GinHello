package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		println("已经授权")
		context.Next()
	}
}
