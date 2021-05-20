package main

import (
	"GinHello/router"
)

func main() {
	router := router.SetupRouter()
	_ = router.Run()
}
