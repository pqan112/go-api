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

func (ur *userRepo) Create(user models.User) error {
	ur.users = append(ur.users, user)
	return nil
}

func (ur *userRepo) Update() {
}

func (ur *userRepo) Delete() {
}

func (ur *userRepo) FindByEmail(email string) (models.User, bool) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, true
		}
	}

	return models.User{}, false
}
