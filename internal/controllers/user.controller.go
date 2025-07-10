package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pqan.com/go-api/internal/models"
	"pqan.com/go-api/internal/services"
	"pqan.com/go-api/internal/validation"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	uc.service.GetUsers()
}

func (uc *UserController) GetUser(ctx *gin.Context) {
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, validation.HandleValidationErrors(err))
		return
	}

	user, err := uc.service.CreateUser(user)
	// TODO ve nha lam
	ctx.JSON(200, gin.H{"message": "User created"})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
}
