package repositories

import (
	"log"

	"pqan.com/go-api/internal/models"
)

type userRepo struct {
	users []models.User
}

func NewUserRepo() UserRepo {
	return &userRepo{
		users: make([]models.User, 0),
	}
}

func (ur *userRepo) FindAll() {
	log.Print("Get All Users repo")
}

func (ur *userRepo) FindByUUID() {
}

func (ur *userRepo) Create() {
}

func (ur *userRepo) Update() {
}

func (ur *userRepo) Delete() {
}
