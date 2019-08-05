package main

import (
	"GinHello/initRouter"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := initRouter.SetupRouter()
	_ = router.Run()
}
