package repositories

import "pqan.com/go-api/internal/models"

type UserRepo interface {
	FindAll()
	FindByUUID()
	Create(user models.User) error
	Update()
	Delete()
	FindByEmail(string) (models.User, bool)
}
