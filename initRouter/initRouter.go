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
	//router.Use(middleware.Auth())
	router.GET("/", middleware.Auth(), func(context *gin.Context) {
		time := time.Now()
		//time = time.Nanosecond()
		context.JSON(http.StatusOK, time.Unix())
	})
	router.GET("/login", user.CreateJwt)
	router.POST("/register", user.Register)

	return router
}
