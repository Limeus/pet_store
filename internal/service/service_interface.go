package service

import "petStore/internal/model"

type PetServiceInterface interface {
	GetPets() []model.Pet
	AddPet(newPet model.PostPet) error
	GetPetByID(id string) (model.Pet, error)
	DeletePetByID(id string) error
	UpdatePetByID(id string, updatedPet model.UpdatePet) (model.Pet, error)
}
