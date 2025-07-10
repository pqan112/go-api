package services

import "pqan.com/go-api/internal/models"

type UserService interface {
	GetUsers()
	GetUserByUUID()
	CreateUser(user models.User) (models.User, error)
	UpdateUser()
}
