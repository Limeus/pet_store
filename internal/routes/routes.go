package routes

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"petStore/docs"
	_ "petStore/docs"
	"petStore/internal/handlers"
)

// SetupRouter initializes the Gin router and registers routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/api/v1/pets", handlers.GetPets)
	router.GET("/api/v1/pets/:id", handlers.GetPetByID)
	router.POST("/api/v1/pets", handlers.PostPets)
	router.DELETE("/api/v1/pets/:id", handlers.DeletePetById)
	router.PUT("/api/v1/pets/:id", handlers.UpdatePetById)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	router.Run(":8080")
	return router
}
