package app

import (
	"pqan.com/go-api/internal/controllers"
	"pqan.com/go-api/internal/repositories"
	"pqan.com/go-api/internal/routes"
	"pqan.com/go-api/internal/services"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {
	userRepo := repositories.NewUserRepo()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRoutes := routes.NewUserRoutes(userController)

	return &UserModule{
		routes: userRoutes,
	}
}

func (um *UserModule) Routes() routes.Route {
	return um.routes
}
