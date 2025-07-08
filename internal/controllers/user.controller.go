package controllers

import (
	"github.com/gin-gonic/gin"
	"pqan.com/go-api/internal/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{"message": uc.service.GetUsers})
	uc.service.GetUsers()
}
