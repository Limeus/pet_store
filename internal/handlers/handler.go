package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"petStore/internal/model"
	"petStore/internal/service"
)

// PetHandler содержит зависимости для хэндлеров питомцев.
type PetHandler struct {
	service service.PetServiceInterface
}

// NewPetHandler создает новый экземпляр PetHandler.
func NewPetHandler(service service.PetServiceInterface) PetHandlerInterface {
	return &PetHandler{service: service}
}

// GetPets godoc
// @Summary      Получить список питомцев
// @Description  Возвращает список всех питомцев
// @Tags         pets
// @Produce      json
// @Success      200  {array}  model.Pet
// @Router   	 /pets [get]
func (h *PetHandler) GetPets(c *gin.Context) {
	pets := h.service.GetPets()
	c.IndentedJSON(http.StatusOK, pets)
}

// PostPets godoc
// @Summary      Добавить нового питомца
// @Description  Добавляет нового питомца в список
// @Tags         pets
// @Accept       json
// @Produce      json
// @Param        pet  body      model.PostPet  true  "Информация о питомце"
// @Success      201  {object}  model.Pet
// @Failure      400  {object}  model.ErrorResponse  "Ошибка при обработке запроса"
// @Router       /pets [post]
func (h *PetHandler) PostPets(c *gin.Context) {
	var newPet model.PostPet

	if err := c.BindJSON(&newPet); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	err := h.service.AddPet(newPet)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newPet)
}

// GetPetByID godoc
// @Summary      Получить питомца по ID
// @Description  Возвращает информацию о питомце по его идентификатору
// @Tags         pets
// @Produce      json
// @Param        id   path      string     true  "ID питомца"
// @Success      200  {object}  model.Pet
// @Failure      404  {object}  model.ErrorResponse  "Питомец не найден"
// @Router       /pets/{id} [get]
func (h *PetHandler) GetPetByID(c *gin.Context) {
	id := c.Param("id")
	pet, err := h.service.GetPetByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, model.ErrorResponse{Message: "pet not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, pet)
}

// DeletePetById godoc
// @Summary      Удалить питомца по ID
// @Description  Удаляет питомца из списка по его идентификатору
// @Tags         pets
// @Produce      json
// @Param        id   path      string     true  "ID питомца"
// @Success      200  {object}  model.DeleteResponse  "Питомец успешно удалён"
// @Failure      404  {object}  model.ErrorResponse  "Питомец не найден"
// @Router       /pets/{id} [delete]
func (h *PetHandler) DeletePetById(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeletePetByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, model.ErrorResponse{Message: "pet not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, model.DeleteResponse{Message: "pet deleted successfully"})
}

// UpdatePetById godoc
// @Summary      Обновить питомца по ID
// @Description  Обновляет информацию о питомце по его идентификатору
// @Tags         pets
// @Accept       json
// @Produce      json
// @Param        id   path      string     true  "ID питомца"
// @Param        pet  body      model.UpdatePet  true  "Обновленная информация о питомце"
// @Success      200  {object}  model.Pet
// @Failure      400  {object}  model.ErrorResponse  "Ошибка при обработке запроса"
// @Failure      404  {object}  model.ErrorResponse  "Питомец не найден"
// @Router       /pets/{id} [put]
func (h *PetHandler) UpdatePetById(c *gin.Context) {
	id := c.Param("id")
	var updatedPet model.UpdatePet

	if err := c.BindJSON(&updatedPet); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	pet, err := h.service.UpdatePetByID(id, updatedPet)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.IndentedJSON(http.StatusNotFound, model.ErrorResponse{Message: "pet not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, pet)
}
