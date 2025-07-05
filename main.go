package main

import (
	"github.com/gin-gonic/gin"
	v1handler "pqan.com/go-api/api/v1/handler"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		userHandlerV1 := v1handler.NewUserHandler()
		v1.GET("/users", userHandlerV1.GetUsers)
		v1.GET("/users/:id", userHandlerV1.GetUser)
		v1.POST("/users", userHandlerV1.CreateUser)
		v1.PUT("/users/:id", userHandlerV1.UpdateUser)
		v1.DELETE("/users/:id", userHandlerV1.DeleteUser)
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
