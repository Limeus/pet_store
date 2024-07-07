package service

import (
	"errors"
	"petStore/internal/model"
	"petStore/internal/repository"
)

// PetService реализует интерфейс PetServiceInterface.
type PetService struct{}

// NewPetService создает новый экземпляр PetService.
func NewPetService() PetServiceInterface {
	return &PetService{}
}

func (s *PetService) GetPets() []model.Pet {
	return repository.GetAll()
}

func (s *PetService) AddPet(newPet model.PostPet) error {
	if string(newPet.Type) != "cat" && string(newPet.Type) != "dog" {
		return errors.New("pet must be cat or dog")
	}
	if newPet.Age < 0 {
		return errors.New("pet age must be >= 0")
	}
	if newPet.Price < 0 {
		return errors.New("pet price must be >= 0")
	}
	repository.Save(newPet.Age, string(newPet.Type), newPet.Price, newPet.Description)
	return nil
}

func (s *PetService) GetPetByID(id string) (model.Pet, error) {
	pet, found := repository.FindById(id)
	if !found {
		return model.Pet{}, errors.New("pet not found")
	}
	return pet, nil
}

func (s *PetService) DeletePetByID(id string) error {
	_, found := repository.DeleteById(id)
	if !found {
		return errors.New("pet not found")
	}
	return nil
}

func (s *PetService) UpdatePetByID(id string, updatedPet model.UpdatePet) (model.Pet, error) {
	pet, found := repository.FindById(id)
	if !found {
		return model.Pet{}, errors.New("pet not found")
	}

	pet.Price = updatedPet.Price
	pet.Description = updatedPet.Description

	err := repository.UpdatePet(id, pet)
	if err != nil {
		return model.Pet{}, err
	}

	return pet, nil
}
