package router

import (
	"app/handler/user"
	"app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/user/login/:userID/:password/:role", user.Login)
	router.GET("/user/register/:userID/:password/:role", user.Register)

	student := router.Group("/student", middleware.Verify())
	{
		student.GET()
	}

	return router
}
