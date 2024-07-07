package routes

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"petStore/docs"
	_ "petStore/docs"
	"petStore/internal/handlers"
	"petStore/internal/service"
)

// SetupRouter initializes the Gin router and registers routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	petService := service.NewPetService()
	petHandler := handlers.NewPetHandler(petService)

	router.GET("/api/v1/pets", petHandler.GetPets)
	router.GET("/api/v1/pets/:id", petHandler.GetPetByID)
	router.POST("/api/v1/pets", petHandler.PostPets)
	router.DELETE("/api/v1/pets/:id", petHandler.DeletePetById)
	router.PUT("/api/v1/pets/:id", petHandler.UpdatePetById)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	router.Run(":8080")

	return router
}
