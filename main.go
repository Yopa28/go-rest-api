package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/config"
	"go-rest-api/routes"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Product API is running",
		})
	})

	routes.ProductRoutes(r)
	routes.AuthRoutes(r)

	r.Run(":8080")
}
