package main

import (
	"log"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/initialize"
)

// @title           Scholar AI Backend API
// @version         0.1.0
// @description     REST API for Scholar AI backend services.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @schemes   http https

func main() {
	// Bootstrap all services
	if err := initialize.Bootstrap(); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to bootstrap application", "error", err)
		}
		log.Fatal("Failed to bootstrap application:", err)
	}

	// Create and run the application
	app, err := NewApp()
	if err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to create application", "error", err)
		}
		log.Fatal("Failed to create application:", err)
	}

	// Start the server
	if err := app.Run(); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to start server", "error", err)
		}
		log.Fatal("Failed to start server:", err)
	}
}
