package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/controllers"
	"github.com/nas03/scholar-ai/backend/internal/repositories"
	"github.com/nas03/scholar-ai/backend/internal/services"
)

// SetupUserRoutes configures all user-related routes
func SetupAuthRoutes(apiV1 *gin.RouterGroup) {

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(global.Mdb)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	// User routes
	auth := apiV1.Group("/auth")
	{

		auth.POST("/login", authController.Login)
		auth.GET("/refresh", authController.RotateAuthToken)
	}
}
