package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messsage": "List all users",
	})
}

func (u *UserHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messsage": "User detail",
	})
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"messsage": "Create user",
	})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messsage": "Update user",
	})
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"messsage": "Update user",
	})
}
