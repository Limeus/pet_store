package service

import (
	"errors"
	"petStore/internal/model"
	"petStore/internal/repository"
)

func GetPets() []model.Pet {
	return repository.GetPets()
}

func AddPet(newPet model.Pet) {
	repository.SavePet(newPet.Age, string(newPet.Type), newPet.Price)
}

func GetPetByID(id string) (model.Pet, error) {
	pet, found := repository.FindPetById(id)
	if !found {
		return model.Pet{}, errors.New("pet not found")
	}
	return pet, nil
}
