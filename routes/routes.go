package routes

import (
	"github.com/andreldsr/alura-gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/students", controllers.FindAllStudent)
	r.GET("/students/:id", controllers.FindStudentById)
	r.GET("/students/document/:doc", controllers.FindStudentByDocument)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/index", controllers.LoadIndex)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
