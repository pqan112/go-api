package services

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"pqan.com/go-api/internal/models"
	"pqan.com/go-api/internal/repositories"
	"pqan.com/go-api/internal/utils"
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

func (us *userService) CreateUser(user models.User) (models.User, error) {
	user.Email = utils.NomalizeText(user.Email)
	if _, exists := us.repo.FindByEmail(user.Email); exists {
		return models.User{}, utils.NewError("error.email_already_exist", http.StatusConflict)
	}

	user.UUID = uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, utils.WrapError("error.internal_server_error", http.StatusInternalServerError, err)
	}

	user.Password = string(hashedPassword)

	if err := us.repo.Create(user); err != nil {
		return models.User{}, utils.WrapError("error.failed_to_create_user", http.StatusInternalServerError, err)
	}

	return user, nil
}

func (us *userService) UpdateUser() {
}

func (us *userService) DeleteUser() {
}
