package controllers

import (
	"github.com/andreldsr/alura-gin-api/database"
	"github.com/andreldsr/alura-gin-api/models"
	"github.com/gin-gonic/gin"
)

func LoadIndex(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(200, "index.html", gin.H{
		"students": students,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(404, "404.html", nil)
}
