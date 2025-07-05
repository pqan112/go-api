package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		userId := c.Param("id")
		age := c.Query("age")
		c.JSON(200, gin.H{
			"userId": userId,
			"age":    age,
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
