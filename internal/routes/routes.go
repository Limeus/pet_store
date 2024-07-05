package routes

import (
	"github.com/gin-gonic/gin"
	"petStore/internal/handlers"
)

// SetupRouter initializes the Gin router and registers routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/pets", handlers.GetPets)
	router.GET("/pets/:id", handlers.GetPetByID)
	router.POST("/pets", handlers.PostPets)
	return router
}
