package routers

import (
	"github.com/chinsiang99/simple-go-project/internal/api/handlers"
	"github.com/chinsiang99/simple-go-project/internal/api/middlewares"
	"github.com/chinsiang99/simple-go-project/internal/config"
	"github.com/gin-gonic/gin"
)

func NewRouterManager(handlers *handlers.HandlerManager, appConfig *config.AppConfig, securityConfig *config.SecurityConfig) *gin.Engine {
	// Set Gin mode if it is production
	if appConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Swagger endpoint (disable in production)
	// if config.Environment != "production" {
	// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// }

	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middlewares.Log())
	router.Use(middlewares.SecurityMiddleware(appConfig, securityConfig))

	RegisterAuthRoutes(router, handlers.AuthHandler)
	RegisterAppRoutes(router)
	// RegisterUserRoutes(router, h.UserHandler)
	// RegisterTicketRoutes(router, h.TicketHandler)

	return router
}
