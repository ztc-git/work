package db

import (
	"app/config"
	"app/initDB"
	"app/model"
	"github.com/gin-gonic/gin"
)


func GetStudentMsg(c *gin.Context) {
	StudentID := c.Query("StudentID")

	student := new(model.Student)
	initDB.Db.Where("StudentID = ?",StudentID).Find(&student)
	if student.Name == "" {
		c.JSON(200, gin.H{"msg":config.ErrorFind})
	}else {
		c.JSON(200, gin.H{"msg":student})
	}
}


func GetTeacherMsg(c *gin.Context) {
	TeacherID := c.Query("TeacherID")

	teacher := new(model.Teacher)
	initDB.Db.Where("TeacherID = ?",TeacherID).Find(&teacher)
	if teacher.Name == "" {
		c.JSON(200, gin.H{"msg":config.ErrorFind})
	}else {
		c.JSON(200, gin.H{"msg":teacher})
	}
}


func UpdateStudentMsg(c *gin.Context) {
	StudentID := c.PostForm("StudentID")
	ClassIDNow := c.PostForm("ClassIDNow")
	MajorNow := c.PostForm("MajorNow")
	CollegeNow := c.PostForm("CollegeNow")

	student := new(model.Student)
	initDB.Db.Where("StudentID = ?", StudentID).Find(&student)

	student.ClassID = ClassIDNow
	student.College = CollegeNow
	student.Major = MajorNow
	initDB.Db.Save(&student)
}


func UpdateTeacherMsg(c *gin.Context) {
	TeacherID := c.PostForm("TeacherID")
	ClassIDNow := c.PostForm("ClassIDNow")
	MajorNow := c.PostForm("MajorNow")

	teacher := new(model.Teacher)
	initDB.Db.Where("TeacherID = ?", TeacherID).Find(&teacher)

	teacher.Major = MajorNow
	teacher.ClassID = ClassIDNow
	initDB.Db.Save(&teacher)
}


func DeleteStudentMsg(c * gin.Context) {
	StudentId := c.PostForm("StudentID")
	initDB.Db.Where("StudentID LIKE ?", StudentId).Delete(model.Student{})
}


func DeleteTeacherMsg(c *gin.Context) {
	TeacherID := c.PostForm("TeacherID")
	initDB.Db.Where("TeacherID LIKE ?", TeacherID).Delete(model.Teacher{})
}

