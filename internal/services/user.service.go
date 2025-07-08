package services

import (
	"log"

	"pqan.com/go-api/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepo
}

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetUsers() {
	log.Print("Get All Users")
}
