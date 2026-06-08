package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/handlers"
	"go-rest-api/middlewares"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/profile", middlewares.AuthMiddleware(), handlers.Profile)
}
