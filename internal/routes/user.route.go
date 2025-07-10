package routes

import (
	"github.com/gin-gonic/gin"
	"pqan.com/go-api/internal/controllers"
)

type UserRoutes struct {
	controller *controllers.UserController
}

func NewUserRoutes(controller *controllers.UserController) *UserRoutes {
	return &UserRoutes{
		controller: controller,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", ur.controller.GetUsers)
		users.GET("/:uuid", ur.controller.GetUser)
		users.POST("", ur.controller.CreateUser)
		users.PUT("/:uuid", ur.controller.UpdateUser)
		users.DELETE("/:uuid", ur.controller.DeleteUser)
	}
}
