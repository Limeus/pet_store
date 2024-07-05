package main

import (
	"petStore/internal/routes"
)

func main() {
	router := routes.SetupRouter()
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
