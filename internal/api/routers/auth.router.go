package routers

import (
	"github.com/chinsiang99/simple-go-project/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		// auth.POST("/register", handler.Register)
	}
}
