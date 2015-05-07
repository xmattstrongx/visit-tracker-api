package repositories

import "RestApiProject/models"

type UsersRepository interface {
	GetAllUsers() ([]models.User, error)
}

type UsersRepo struct {
}

func (u UsersRepo) GetAllUsers() ([]models.User, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	IDFirstNameLastName := "id, firstname, lastname"

	users := []models.User{}
	if err := db.
		Select(IDFirstNameLastName).
		From(Users).
		QueryStructs(&users); err != nil {
		return nil, err
	}
	return users, nil
}
