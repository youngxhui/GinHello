package main

import (
	"gin.hello/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
