package routes

import (
	"github.com/gin-gonic/gin"
	"pqan.com/go-api/internal/middlewares"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	r.Use(middlewares.AuthMiddleware())
	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}
}
