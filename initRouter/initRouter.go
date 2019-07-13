package initRouter

import (
	"GinHello/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/")
	{
		// 添加 Get 请求路由
		index.GET("", retHelloGinAndMethod)
		// 添加 Post 请求路由
		index.POST("", retHelloGinAndMethod)
		// 添加 Put 请求路由
		index.PUT("", retHelloGinAndMethod)
		// 添加 Delete 请求路由
		index.DELETE("", retHelloGinAndMethod)
		// 添加 Patch 请求路由
		index.PATCH("", retHelloGinAndMethod)
		// 添加 Head 请求路由
		index.HEAD("", retHelloGinAndMethod)
		// 添加 Options 请求路由
		index.OPTIONS("", retHelloGinAndMethod)
		//index.Any("", retHelloGinAndMethod)
	}
	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", user.Save)
		userRouter.GET("", user.SaveByQuery)
	}

	//router.GET("/user/:name", user.Save)
	//router.GET("/user", user.SaveByQuery)
	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
