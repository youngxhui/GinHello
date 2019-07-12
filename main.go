package main

import (
	initRouter "GinHello/init-router"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
