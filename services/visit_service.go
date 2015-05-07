package services

import (
	"RestApiProject/models"
	"RestApiProject/repositories"
)

func GetAllOfAUsersVisits(userID string, paginationParameters models.PaginationParameters) ([]models.Visit, error) {
	v := repositories.VisitsRepo{}
	visits, err := v.GetAllOfAUsersVisits(userID, paginationParameters)
	if err != nil {
		return visits, err
	}
	return visits, nil
}

func GetAllOfAUsersStateVisits(userID string) ([]models.State, error) {
	v := repositories.VisitsRepo{}
	visits, err := v.GetAllOfAUsersStateVisits(userID)
	if err != nil {
		return visits, err
	}
	return visits, nil
}

func PostVisit(userID string, visit models.Visit) (models.VisitDetails, error) {
	if err := validVisitPost(visit); err != nil {
		return models.VisitDetails{}, err
	}
	c := repositories.CityRepo{}
	if !c.CityIsInState(visit) {
		return models.VisitDetails{}, newCityNotExistInStateError(visit.City, visit.Abbreviation)
	}

	v := repositories.VisitsRepo{}
	visits, err := v.PostVisit(userID, visit)
	if err != nil {
		return visits, err
	}
	return visits, nil
}

func DeleteVisit(userID, visitID string) error {
	v := repositories.VisitsRepo{}
	err := v.DeleteVisit(userID, visitID)
	if err != nil {
		return err
	}
	return nil

}

func validVisitPost(visit models.Visit) error {
	if visit.City == "" {
		return newVisitValidationError("city")
	} else if visit.Abbreviation == "" {
		return newVisitValidationError("state")
	}
	return nil
}
