package repositories

import (
	"RestApiProject/models"
	"fmt"
	"log"
)

type VisitsRepository interface {
	GetAllOfAUsersVisits(string) ([]models.Visit, error)
}

type VisitsRepo struct {
}

func (v VisitsRepo) GetAllOfAUsersVisits(userID string, paginationParameters models.PaginationParameters) ([]models.Visit, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	nameAbbreviation := "city, abbreviation"
	userIDEqualsUserID := fmt.Sprintf("userid = '%s'", userID)

	query := db.
		Select(nameAbbreviation).
		From(Visits).
		Where(userIDEqualsUserID)

	ApplyPaging(query, paginationParameters)

	visits := []models.Visit{}
	if err := query.QueryStructs(&visits); err != nil {
		return nil, err
	}
	return visits, nil
}

func (v VisitsRepo) GetAllOfAUsersStateVisits(userID string) ([]models.State, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	userIDEqualsUserID := fmt.Sprintf("userid = '%s'", userID)

	visits := []models.State{}
	if err := db.
		Select("v.abbreviation, s.name").
		Distinct().
		From(`visits v INNER JOIN states s on (s.abbreviation = v.abbreviation)`).
		Where(userIDEqualsUserID).
		QueryStructs(&visits); err != nil {
		return nil, err
	}

	return visits, nil
}

func (v VisitsRepo) PostVisit(userID string, visit models.Visit) (models.VisitDetails, error) {
	db, err := initDB()
	if err != nil {
		return models.VisitDetails{}, err
	}

	visits := models.VisitDetails{}

	if err := db.
		InsertInto(Visits).
		Columns("userid", "city", "abbreviation").
		Values(userID, visit.City, visit.Abbreviation).
		Returning("*").
		QueryStruct(&visits); err != nil {
		return models.VisitDetails{}, err
	}

	return visits, nil
}

func (v VisitsRepo) DeleteVisit(userID, visitID string) error {
	db, err := initDB()
	if err != nil {
		return err
	}

	idEqualsVisitID := fmt.Sprintf("id = '%s'", visitID)
	userIDEqualsUserID := fmt.Sprintf("userid = '%s'", userID)

	if _, err = db.
		DeleteFrom(Visits).
		Where(idEqualsVisitID).
		Where(userIDEqualsUserID).
		Exec(); err != nil {
		log.Printf("err: %v", err)
		return err
	}

	return nil
}
