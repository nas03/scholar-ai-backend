package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/controllers"
	"github.com/nas03/scholar-ai/backend/internal/repositories"
	"github.com/nas03/scholar-ai/backend/internal/services"
)

// SetupUserRoutes configures all user-related routes
func SetupUserRoutes(apiV1 *gin.RouterGroup) {

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(global.Mdb)
	mailRepo := repositories.NewMailRepository(global.Mdb)
	userService := services.NewUserService(userRepo, mailRepo)
	userController := controllers.NewUserController(userService)

	// User routes
	users := apiV1.Group("/users")
	{
		users.POST("/create", userController.CreateUser)
		users.POST("/activate", userController.ActivateUserAccount)
		users.GET("/ping", controllers.Ping) // Keep ping for testing
	}
}
