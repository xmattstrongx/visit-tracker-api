package services

import (
	"RestApiProject/models"
	"RestApiProject/repositories"
)

func GetAllUsers() ([]models.User, error) {
	u := repositories.UsersRepo{}
	users, err := u.GetAllUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}
