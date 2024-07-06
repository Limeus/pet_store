package main

import (
	"petStore/internal/routes"
)

// @title Pet Store API
// @version 1.0
// @description pet microservice
// @host localhost:8080
// @BasePath /

func main() {
	router := routes.SetupRouter()
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
