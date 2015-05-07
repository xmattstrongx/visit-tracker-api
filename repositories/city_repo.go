package repositories

import (
	"RestApiProject/models"
	"fmt"
)

type CitiesRepository interface {
	GetCitiesForSpecificState(int) ([]models.City, error)
	GetDetailsForSpecificCity(int, string) (models.CityDetails, error)
	CityIsInState(visit models.Visit) bool
}

type CityRepo struct {
}

func (c CityRepo) GetCitiesForSpecificState(stateID int, paginationParameters models.PaginationParameters) ([]models.City, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	nameCounty := "Name, County"
	stateIDEqualsAbbreviation := fmt.Sprintf("stateid = '%d'", stateID)

	query := db.Select(nameCounty).
		From(Cities).
		Where(stateIDEqualsAbbreviation).
		OrderBy(Name)

	ApplyPaging(query, paginationParameters)

	cities := []models.City{}
	if err := query.QueryStructs(&cities); err != nil {
		return nil, err
	}

	return cities, nil
}

func (c CityRepo) GetDetailsForSpecificCity(stateID int, cityName string) (models.CityDetails, error) {
	db, err := initDB()
	if err != nil {
		return models.CityDetails{}, err
	}

	stateIDEqualsAbbreviationAndnameEqualsCityName := fmt.Sprintf("stateid = '%d' AND name = '%s'", stateID, cityName)

	city := models.CityDetails{}
	if err := db.
		Select(Star).
		From(Cities).
		Where(stateIDEqualsAbbreviationAndnameEqualsCityName).
		QueryStruct(&city); err != nil {
		return models.CityDetails{}, err
	}

	return city, nil
}

func (c CityRepo) CityIsInState(visit models.Visit) bool {
	db, err := initDB()
	if err != nil {
		return false
	}

	nameEqualsVisitCityAndStateEqualsVisitState := fmt.Sprintf("name = '%s' and stateid = (select id from states where abbreviation = '%s')", visit.City, visit.Abbreviation)

	city := models.CityDetails{}
	if err := db.
		Select(Star).
		From(Cities).
		Where(nameEqualsVisitCityAndStateEqualsVisitState).
		QueryStruct(&city); err != nil {
		return false
	}
	return true
}
