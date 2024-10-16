package repositories

import (
	"api/src/models"
	"database/sql"

	"github.com/guregu/null"
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
	statement, erro := citiesRepository.db.Prepare("insert into cities (cityName, cityUf) values (?, ?)")
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
func (citiesRepository CitiesRepository) GetCityByUF(ufSerach string) ([]models.City, error) {
	rows, erro := citiesRepository.db.Query("select * from cities where cityUf = ?", ufSerach)
	if erro != nil {
		return nil, erro
	}

	defer rows.Close()
	var cities []models.City

	for rows.Next() {
		var cityObj models.City

		if erro = rows.Scan(
			&cityObj.ID,
			&cityObj.Name,
			&cityObj.UF,
			&cityObj.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		cities = append(cities, cityObj)
	}
	return cities, nil
}

// Get city by id
func (citiesRepository CitiesRepository) GetCityById(id uint64) (models.City, error) {
	rows, erro := citiesRepository.db.Query("select * from cities where cityId = ?", id)
	if erro != nil {
		return models.City{
			ID:        0,
			Name:      "",
			UF:        "",
			CreatedAt: &null.Time{},
		}, erro
	}

	defer rows.Close()
	var cityObj models.City

	for rows.Next() {
		if erro = rows.Scan(
			&cityObj.ID,
			&cityObj.Name,
			&cityObj.UF,
			&cityObj.CreatedAt,
		); erro != nil {
			return models.City{
				ID:        0,
				Name:      "",
				UF:        "",
				CreatedAt: &null.Time{},
			}, erro
		}
	}
	return cityObj, nil
}

func (citiesRepository CitiesRepository) UpdateCityById(id uint64, cityObj models.City) error {

	statement, erro := citiesRepository.db.Prepare("update  cities set cityName = ?, cityUf = ? where cityId = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(cityObj.Name, cityObj.UF, id); erro != nil {
		return erro
	}
	return nil
}
