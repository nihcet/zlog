package main

import (
	"github.com/nihcet/zlog/src/router"
)

func main() {
	router := router.SetupRouter()
	router.Run()
}
