package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"petStore/internal/converter"
	"petStore/internal/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "pets"
)

func getDBConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return sql.Open("postgres", psqlconn)
}

func SavePet(age int, petType string, price float64) {
	db, err := getDBConnection()
	CheckError(err)
	defer db.Close()

	insertStmt := `insert into pets("age", "type", "price") values($1, $2, $3)`
	_, e := db.Exec(insertStmt, age, petType, price)
	CheckError(e)
}

func GetPets() []model.Pet {
	db, err := getDBConnection()
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT "age", "type", "price" FROM pets`)
	CheckError(err)
	defer rows.Close()

	var pets []model.Pet
	pets, err = converter.ToModels(rows)
	if err != nil {
		CheckError(err)
	}
	return pets
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func FindPetById(id string) (model.Pet, bool) {
	db, err := getDBConnection()
	if err != nil {
		return model.Pet{}, false
	}
	defer db.Close()

	var pet model.Pet
	var petType string

	query := "SELECT age, type, price FROM pets WHERE id = $1"
	err = db.QueryRow(query, id).Scan(&pet.Age, &petType, &pet.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Pet{}, false
		}
		return model.Pet{}, false
	}

	switch petType {
	case "cat":
		pet.Type = model.Cat
	case "dog":
		pet.Type = model.Dog
	default:
		return model.Pet{}, false
	}

	return pet, true
}
