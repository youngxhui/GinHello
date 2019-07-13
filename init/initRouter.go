package init

import (
	"GinHello/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 添加 Get 请求路由
	router.GET("/", retHelloGinAndMethod)
	// 添加 Post 请求路由
	router.POST("/", retHelloGinAndMethod)
	// 添加 Put 请求路由
	router.PUT("/", retHelloGinAndMethod)
	// 添加 Delete 请求路由
	router.DELETE("/", retHelloGinAndMethod)
	// 添加 Patch 请求路由
	router.PATCH("/", retHelloGinAndMethod)
	// 添加 Head 请求路由
	router.HEAD("/", retHelloGinAndMethod)
	// 添加 Options 请求路由
	router.OPTIONS("/", retHelloGinAndMethod)
	// 添加 user
	router.GET("/user/:name", user.Save)
	router.GET("/user", user.SaveByQuery)
	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
