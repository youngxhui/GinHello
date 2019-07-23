package initRouter

import (
	"GinHello/handler"
	"GinHello/middleware"
	"GinHello/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {

	router := gin.New()
	// 添加自定义的 logger 中间件
	router.Use(middleware.Logger(), gin.Recovery())

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.StaticFile("/favicon.ico", "/favicon.ico")
	router.Static("/statics", "./statics/")
	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}
	return router
}
