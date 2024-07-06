package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petStore/internal/model"
	"petStore/internal/service"
)

func GetPets(c *gin.Context) {
	pets := service.GetPets()
	c.IndentedJSON(http.StatusOK, pets)
}

func PostPets(c *gin.Context) {
	var newPet model.Pet

	if err := c.BindJSON(&newPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.AddPet(newPet)
	c.IndentedJSON(http.StatusCreated, newPet)
}

func GetPetByID(c *gin.Context) {
	id := c.Param("id")
	pet, err := service.GetPetByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pet not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, pet)
}

func DeletePetById(c *gin.Context) {
	id := c.Param("id")
	err := service.DeletePetByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pet not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "pet deleted successfully"})
}
