package initRouter

import (
	"GinHello/handler/user"
	"GinHello/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", middleware.Auth(), func(context *gin.Context) {
		context.JSON(http.StatusOK, time.Now().Unix())
	})
	router.GET("/login", user.CreateJwt)
	router.POST("/register", user.Register)

	return router
}
