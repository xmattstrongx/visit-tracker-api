package repositories

import (
	"RestApiProject/models"
	"fmt"
)

type StatesRepository interface {
	GetAllStates() ([]models.State, error)
	GetStateID(string) (int, error)
	GetSpecificState(string) (models.StateDetails, error)
}

type StatesRepo struct {
}

func (s StatesRepo) GetAllStates() ([]models.State, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	nameAbbreviation := "name, abbreviation"

	states := []models.State{}
	if err := db.
		Select(nameAbbreviation).
		From(States).
		QueryStructs(&states); err != nil {
		return nil, err
	}
	return states, nil
}

func (s StatesRepo) GetStateID(abbreviation string) (int, error) {
	db, err := initDB()
	if err != nil {
		return 0, err
	}

	abbreviationEqualsState := fmt.Sprintf("abbreviation = '%s'", abbreviation)

	// select * from cities where stateid = (select id from states where abbreviation = 'AL')
	var stateID int
	if err := db.
		Select(StateID).
		From(States).
		Where(abbreviationEqualsState).
		QueryScalar(&stateID); err != nil {
		return 0, err
	}
	return stateID, nil
}

func (s StatesRepo) GetSpecificState(abbreviation string) (models.StateDetails, error) {
	db, err := initDB()
	if err != nil {
		return models.StateDetails{}, err
	}

	abbreviationEqualsState := fmt.Sprintf("abbreviation = '%s'", abbreviation)

	state := models.StateDetails{}
	if err := db.
		Select(Star).
		From(States).
		Where(abbreviationEqualsState).
		QueryStruct(&state); err != nil {
		return models.StateDetails{}, err
	}
	return state, nil
}
