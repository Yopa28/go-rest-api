// @title Product API
// @version 1.0
// @description REST API using Golang, Gin, MySQL, JWT Authentication
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	"go-rest-api/config"
	"go-rest-api/routes"
	_ "go-rest-api/docs"
	
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Product API is running",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.ProductRoutes(r)
	routes.AuthRoutes(r)

	r.Run(":8080")
}

