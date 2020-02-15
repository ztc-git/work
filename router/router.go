package router

import (
	db "app/Database"
	"app/handler/user"
	"app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/user/login/:userID/:password/:role", user.Login)
	router.GET("/user/register/:userID/:password/:role", user.Register)
	router.GET("/user/logout",user.Logout)
	router.POST("/user/insert", user.InsertMsg)

	student := router.Group("/student", middleware.Verify())
	{
		student.GET("/search",db.GetStudentMsg)
	}

	teacher := router.Group("/teacher",middleware.Verify())
	{
		teacher.GET("/search/student", db.GetStudentMsg)
		teacher.GET("/search/teacher", db.GetTeacherMsg)
		teacher.PUT("/update/student", db.UpdateStudentMsg)
	}

	admin := router.Group("/admin", middleware.Verify())
	{
		admin.GET("/search/student", db.GetStudentMsg)
		admin.GET("/search/teacher", db.GetTeacherMsg)
		admin.PUT("update/student", db.UpdateStudentMsg)
		admin.PUT("update/teacher", db.UpdateTeacherMsg)
		admin.DELETE("/delete/student",db.DeleteStudentMsg)
		admin.DELETE("/delete/teacher", db.DeleteTeacherMsg)
	}

	return router
}
