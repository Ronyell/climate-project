package repositories

import (
	"api/src/models"
	"database/sql"
)

// Represent a city repository
type CitiesRepository struct {
	db *sql.DB
}

// Get to return repository
func GetCitiesRepository(db *sql.DB) *CitiesRepository {
	return &CitiesRepository{db}
}

// Create a city
func (citiesRepository CitiesRepository) Create(cityObj models.City) (uint64, error) {
	statement, erro := citiesRepository.db.Prepare("insert into cities (name, uf) values (?, ?)")
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	result, erro := statement.Exec(cityObj.Name, cityObj.UF)
	if erro != nil {
		return 0, nil
	}

	lastID, erro := result.LastInsertId()

	if erro != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}

// Get all cities
// func (citiesRepository CitiesRepository) getAll() (uint64, error) {
// statement, erro := citiesRepository.db.Prepare("select * from cities")
// if erro != nil {
// 	return 0, nil
// }
// defer statement.Close()

// result, erro := statement.Exec()
// if erro != nil {
// 	return 0, nil
// }

// if erro != nil {
// 	return 0, nil
// }

// return uint64(lastID), nil
// }
