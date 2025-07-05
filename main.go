package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messsage": "List all users",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messsage": "List user detail",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"messsage": "Create user",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messsage": "Update user",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"messsage": "Update user",
	})
}

func main() {
	router := gin.Default()

	router.GET("/api/v1/users", GetUsers)
	router.GET("/api/v1/users/:id", GetUser)
	router.POST("/api/v1/users", CreateUser)
	router.PUT("/api/v1/users/:id", UpdateUser)
	router.DELETE("/api/v1/users/:id", DeleteUser)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
