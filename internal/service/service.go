package service

import (
	"errors"
	"petStore/internal/model"
	"petStore/internal/repository"
)

func GetPets() []model.Pet {
	return repository.GetAll()
}

func AddPet(newPet model.PostPet) error {
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

func UpdatePetByID(id string, updatedPet model.UpdatePet) (model.Pet, error) {
	// Найдите текущего питомца по ID, чтобы убедиться, что он существует
	pet, found := repository.FindById(id)
	if !found {
		return model.Pet{}, errors.New("pet not found")
	}

	// Обновите данные питомца
	pet.Price = updatedPet.Price
	pet.Description = updatedPet.Description

	// Сохраните изменения в репозитории
	repository.UpdatePet(id, pet)

	return pet, nil
}
