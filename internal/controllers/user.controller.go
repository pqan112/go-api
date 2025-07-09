package controllers

import (
	"github.com/gin-gonic/gin"
	"pqan.com/go-api/internal/services"
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
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
}
