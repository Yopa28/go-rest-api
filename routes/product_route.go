package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/handlers"
	"go-rest-api/middlewares"
)

func ProductRoutes(r *gin.Engine) {

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	// semua user bisa GET
	protected.GET("/products", handlers.GetProducts)
	protected.GET("/products/:id", handlers.GetProductByID)

	// hanya admin yang bisa POST, PUT, DELETE
	admin := protected.Group("/")
	admin.Use(middlewares.AdminMiddleware())

	admin.POST("/products", handlers.CreateProduct)
	admin.PUT("/products/:id", handlers.UpdateProduct)
	admin.DELETE("/products/:id", handlers.DeleteProduct)
}