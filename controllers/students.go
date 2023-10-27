package controllers

import (
	"github.com/andreldsr/alura-gin-api/database"
	"github.com/andreldsr/alura-gin-api/models"
	"github.com/gin-gonic/gin"
)

func FindAllStudent(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func FindStudentById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Student " + id + " not found",
		})
		return
	}
	c.JSON(200, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		panic(err.Error())
		return
	}
	if err := models.ValidateStudent(student); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Student not found",
		})
		return
	}
	database.DB.Delete(&student, id)
	c.JSON(200, gin.H{
		"message": "Deleted with success",
	})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Student not found",
		})
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidateStudent(student); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Updates(&student)
	c.JSON(200, gin.H{
		"message": "Student updated succesfully",
	})
}

func FindStudentByDocument(c *gin.Context) {
	var student models.Student
	doc := c.Param("doc")
	database.DB.First(&student, models.Student{Document: doc})
	if student.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Student not found",
		})
		return
	}
	c.JSON(200, student)
}
