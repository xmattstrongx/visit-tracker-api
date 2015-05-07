package services

import (
	"RestApiProject/models"
	"RestApiProject/repositories"
	"strings"
)

type StatesService interface {
	GetAllStates() ([]models.State, error)
	GetSpecificState(string) (models.StateDetails, error)
}

func GetAllStates() ([]models.State, error) {
	s := repositories.StatesRepo{}
	states, err := s.GetAllStates()
	if err != nil {
		return states, err
	}
	return states, nil
}

func GetSpecificState(stateAbbreviation string) (models.StateDetails, error) {
	s := repositories.StatesRepo{}
	states, err := s.GetSpecificState(strings.ToUpper(stateAbbreviation))
	if err != nil {
		return states, err
	}

	return states, nil
}

func GetCitiesForSpecificState(stateAbbreviation string, paginationParameters models.PaginationParameters) ([]models.City, error) {
	s := repositories.StatesRepo{}
	stateID, err := s.GetStateID(strings.ToUpper(stateAbbreviation))
	if err != nil {
		return nil, err
	}

	c := repositories.CityRepo{}
	cities, err := c.GetCitiesForSpecificState(stateID, paginationParameters)
	if err != nil {
		return nil, err
	}

	// This seems redundant since in the API we already know what state we are in.
	// I have just added this to match the example in the README
	for i := range cities {
		cities[i].State = stateAbbreviation
	}

	return cities, nil
}

func GetDetailsForSpecificCity(stateAbbreviation, cityName string) (models.CityDetails, error) {
	s := repositories.StatesRepo{}
	stateID, err := s.GetStateID(strings.ToUpper(stateAbbreviation))
	if err != nil {
		return models.CityDetails{}, err
	}

	c := repositories.CityRepo{}
	cities, err := c.GetDetailsForSpecificCity(stateID, cityName)
	if err != nil {
		return cities, err
	}

	return cities, nil
}
