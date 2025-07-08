package repositories

import "pqan.com/go-api/internal/models"

type UserRepo struct {
	users []models.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make([]models.User, 0),
	}
}
