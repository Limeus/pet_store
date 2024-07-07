package handlers

import "github.com/gin-gonic/gin"

type PetHandlerInterface interface {
	GetPets(c *gin.Context)
	PostPets(c *gin.Context)
	GetPetByID(c *gin.Context)
	DeletePetById(c *gin.Context)
	UpdatePetById(c *gin.Context)
}
