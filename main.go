package main

import (
	"app/router"
)

func main() {
	router := router.SetupRouter()
	router.Run()
}
