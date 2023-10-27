package database

import (
	"github.com/andreldsr/alura-gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=localhost user=root password=root dbname=students port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err.Error())
	}
	err = DB.AutoMigrate(&models.Student{})
	if err != nil {
		panic(err.Error())
	}
}
