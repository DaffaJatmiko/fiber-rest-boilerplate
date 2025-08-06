package routes

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/config"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes handlers and configures all routes
func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// Initialize all handlers here
	h := NewHandlers(cfg)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "message": "Server is running"})
	})

	// API v1 group
	api := app.Group("/api/v1")

	// Setup route groups with handlers
	setupAuthRoutes(api, h, cfg.JWT.Secret)
}

// setupAuthRoutes configures authentication routes
func setupAuthRoutes(api fiber.Router, h *Handlers, jwtSecret string) {
	auth := api.Group("/auth")
	auth.Post("/register", h.Auth.Register)
	auth.Post("/login", h.Auth.Login)

	// Protected auth routes
	protected := api.Group("/", middleware.AuthMiddleware(jwtSecret))
	protected.Get("/profile", h.Auth.GetProfile)
}
