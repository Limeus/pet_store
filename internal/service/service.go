package service

import (
	"errors"
	"petStore/internal/model"
	"petStore/internal/repository"
)

func GetPets() []model.Pet {
	return repository.GetAll()
}

func AddPet(newPet model.Pet) {
	repository.Save(newPet.Age, string(newPet.Type), newPet.Price)
}

func GetPetByID(id string) (model.Pet, error) {
	pet, found := repository.FindById(id)
	if !found {
		return model.Pet{}, errors.New("pet not found")
	}
	return pet, nil
}

func DeletePetByID(id string) error {
	_, found := repository.DeleteById(id)
	if !found {
		return errors.New("pet not found")
	}
	return nil
}
