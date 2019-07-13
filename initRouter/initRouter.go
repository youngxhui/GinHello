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
		index.Any("", retHelloGinAndMethod)
	}
	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", user.Save)
		userRouter.GET("", user.SaveByQuery)
	}
	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
