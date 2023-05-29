package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	return router
}
