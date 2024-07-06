package converter

import (
	"database/sql"
	"fmt"
	"petStore/internal/model"
)

func ToModels(rows *sql.Rows) ([]model.Pet, error) {
	defer rows.Close()

	var pets []model.Pet
	for rows.Next() {
		var pet model.Pet
		var petType string

		err := rows.Scan(&pet.Age, &petType, &pet.Price)
		if err != nil {
			return nil, err
		}

		switch petType {
		case "cat":
			pet.Type = model.Cat
		case "dog":
			pet.Type = model.Dog
		default:
			fmt.Printf("unknown pet type: %s\n", petType)
			continue
		}

		pets = append(pets, pet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pets, nil
}

//func ToModel(row *sql.Row) model.Pet {
//	var age int
//	var petType string
//	var price float64
//	var petTypeMod model.Type
//
//	err := row.Scan(&age, &petType, &price)
//	CheckError(err)
//
//	if petType == "cat" {
//		petTypeMod = model.Cat
//	} else {
//		petTypeMod = model.Dog
//	}
//
//	return model.Pet{
//		Age:   age,
//		Type:  petTypeMod,
//		Price: price,
//	}
//}
//
//func CheckError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
