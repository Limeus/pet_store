package converter

import (
	"database/sql"
	"petStore/internal/model"
)

func ToModels(rows *sql.Rows) ([]model.Pet, error) {
	var pets []model.Pet
	for rows.Next() {
		var pet model.Pet
		var description sql.NullString
		err := rows.Scan(&pet.Age, &pet.Type, &pet.Price, &pet.DateAdded, &description)
		if err != nil {
			CheckError(err)
		}
		if description.Valid {
			pet.Description = description.String
		} else {
			pet.Description = ""
		}
		pets = append(pets, pet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pets, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
