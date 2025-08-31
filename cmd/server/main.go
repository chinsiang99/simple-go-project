package main

import (
	// _ "github.com/simple-go-project/docs" // Import swagger docs

	"net/http"

	"github.com/chinsiang99/simple-go-project/internal/api/middlewares"
	"github.com/chinsiang99/simple-go-project/internal/config"

	// "github.com/chinsiang99/simple-go-project/internal/database"
	"github.com/chinsiang99/simple-go-project/internal/utils/logger"
	"github.com/gin-gonic/gin"
)

// @title github.com/simple-go-project API
// @version 1.0
// @description This is a sample API server with authentication
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Load configuration
	cfg := config.New()

	// Initialize logger with file rotation
	logger.Init(cfg.LOG)

	// Initialize database
	// db, err := database.Init(cfg.DB)
	// if err != nil {
	// 	logger.Fatal("Failed to initialize database:", err)
	// }

	// Set Gin mode if it is production
	if cfg.APP.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	// router := routes.SetupRoutes(db, cfg)

	// Swagger endpoint (disable in production)
	// if config.Environment != "production" {
	// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// }

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middlewares.Log())

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// Start server
	logger.Infof("Starting server on port %v", cfg.APP.AppPort)
	// if err := router.Run(":" + cfg.APP.AppPort); err != nil {
	// 	logger.Fatal("Failed to start server:", err)
	// }
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}
