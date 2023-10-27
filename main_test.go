package main

import (
	"encoding/json"
	"github.com/andreldsr/alura-gin-api/controllers"
	"github.com/andreldsr/alura-gin-api/database"
	"github.com/andreldsr/alura-gin-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var ID int

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func MockStudent() {
	student := models.Student{Name: "Mock Student", Document: "12345678901"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteMockStudent() {
	var aluno models.Student
	database.DB.Delete(&aluno, ID)
}

func TestListAllStudents(t *testing.T) {
	database.Connect()
	MockStudent()
	defer DeleteMockStudent()
	r := SetupRoutes()
	r.GET("/students", controllers.FindAllStudent)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var mockStudents []models.Student
	err := json.Unmarshal(res.Body.Bytes(), &mockStudents)
	if err != nil {
		return
	}
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, 1, len(mockStudents))
}

func TestFindStudentByID(t *testing.T) {
	database.Connect()
	MockStudent()
	defer DeleteMockStudent()
	r := SetupRoutes()
	r.GET("/students/:id", controllers.FindStudentById)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var mockStudent models.Student
	json.Unmarshal(res.Body.Bytes(), &mockStudent)
	assert.Equal(t, "Mock Student", mockStudent.Name, "The names should be equal")
	assert.Equal(t, "12345678901", mockStudent.Document)
	assert.Equal(t, http.StatusOK, res.Code)
}
