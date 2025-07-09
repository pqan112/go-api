package services

import (
	"log"

	"pqan.com/go-api/internal/repositories"
)

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetUsers() {
	log.Print("Get All Users service")
	us.repo.FindAll()
}

func (us *userService) GetUserByUUID() {
}

func (us *userService) CreateUser() {
}

func (us *userService) UpdateUser() {
}

func (us *userService) DeleteUser() {
}
