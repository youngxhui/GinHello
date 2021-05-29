package main

import (
	"GinHello/router"
)

func main() {
	eng := router.SetupRouter()

	_ = eng.Run(":8181")
}
