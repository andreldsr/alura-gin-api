package main

import (
	"github.com/andreldsr/alura-gin-api/database"
	"github.com/andreldsr/alura-gin-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
