package app

import (
	"github.com/gin-gonic/gin"
	"pqan.com/go-api/internal/config"
	"pqan.com/go-api/internal/routes"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	router *gin.Engine
	config *config.Config
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		router: r,
		config: cfg,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))

	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}
