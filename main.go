package main

import (
	"github.com/gin-gonic/gin"
	v1handler "pqan.com/go-api/api/v1/handler"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			users.GET("", userHandlerV1.GetUsers)
			users.GET("/:id", userHandlerV1.GetUser)
			users.POST("", userHandlerV1.CreateUser)
			users.PUT("/:id", userHandlerV1.UpdateUser)
			users.DELETE("/:id", userHandlerV1.DeleteUser)
		}

		products := v1.Group("/products")
		{
			productHandlerV1 := v1handler.NewProductHandler()
			products.GET("", productHandlerV1.GetProducts)
			products.GET("/:id", productHandlerV1.GetProduct)
			products.POST("", productHandlerV1.CreateProduct)
			products.PUT("/:id", productHandlerV1.UpdateProduct)
			products.DELETE("/:id", productHandlerV1.DeleteProduct)
		}

	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
