package main

import (
	"GinHello/init"
)

func main() {
	router := init.SetupRouter()
	_ = router.Run()
}
