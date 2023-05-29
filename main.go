package main

import (
	"log"

	"gin.hello/router"
)

func main() {
	router := router.InitRouter()
	err := router.Run()
	if err != nil {
		log.Fatalln("gin run error",err)
	}
}
