package main

import (
	"app/handler/user"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.GET("/user/login/:username/:password/", user.Login)
	r.GET("/user/register/:username/:password", user.Register)

	r.Run()
}


