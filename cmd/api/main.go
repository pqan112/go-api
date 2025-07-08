package main

import (
	"github.com/gin-gonic/gin"
	"pqan.com/go-api/internal/config"
	"pqan.com/go-api/internal/controllers"
	"pqan.com/go-api/internal/repositories"
	"pqan.com/go-api/internal/routes"
	"pqan.com/go-api/internal/services"
)

func main() {
	config := config.NewConfig()
	userRepo := repositories.NewUserRepo()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRoutes := routes.NewUserRoutes(userController)
	r := gin.Default()
	routes.RegisterRoutes(r, userRoutes)
	if err := r.Run(config.ServerAddress); err != nil {
		panic(err)
	}
}
