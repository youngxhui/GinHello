package main

import (
	"GinHello/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run(":8181")
}
