package main

import (
	"GinHello/router"
)

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
