package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/handlers"
	"go-rest-api/middlewares"
)

func ProductRoutes(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/products", handlers.GetProducts)
	protected.GET("/products/:id", handlers.GetProductByID)
	protected.PUT("/products/:id", handlers.UpdateProduct)
	protected.DELETE("/products/:id", handlers.DeleteProduct)
}
