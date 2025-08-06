package main

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/config"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/database"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/models"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/routes"
	pkgLogger "github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/logger"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Setup database
	setupDatabase(cfg)

	// Setup Fiber app
	app := setupFiberApp()

	// Setup routes (handles all dependencies internally)
	routes.SetupRoutes(app, cfg)

	// Start server
	startServer(app, cfg.App.Port)
}

func setupDatabase(cfg *config.Config) {
	pkgLogger.Init()
	pkgLogger.Info("Starting Go REST API Boilerplate")

	database.ConnectDB(cfg)
	pkgLogger.Info("Database connected")

	if err := database.DB.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		log.Fatal("Database migration failed:", err)
	}
	pkgLogger.Info("Database migration completed")
}

func setupFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Go REST API Boilerplate v1.0.0",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			pkgLogger.Error("Global error: " + err.Error())

			if e, ok := err.(*fiber.Error); ok {
				if e.Code == fiber.StatusBadRequest {
					return response.BadRequest(c, e.Message)
				}
				if e.Code == fiber.StatusUnauthorized {
					return response.Unauthorized(c, e.Message)
				}
			}

			return response.InternalError(c, "Internal Server Error")
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	app.Use(logger.New())

	return app
}

func startServer(app *fiber.App, port string) {
	pkgLogger.Info("Server starting on port " + port)
	log.Fatal(app.Listen(":" + port))
}
