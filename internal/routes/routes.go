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
	setupUserRoutes(api, h, cfg.JWT.Secret)
	setupBookRoutes(api, h, cfg.JWT.Secret)
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

// setupUserRoutes configures user routes
func setupUserRoutes(api fiber.Router, h *Handlers, jwtSecret string) {
	protected := api.Group("/", middleware.AuthMiddleware(jwtSecret))
	protected.Get("/users", h.User.GetAll)
	protected.Get("/users/:id", h.User.GetByID)
	protected.Put("/users/:id", h.User.Update)
	protected.Delete("/users/:id", h.User.Delete)
}

// setupBookRoutes configuras book routes
func setupBookRoutes(api fiber.Router, h *Handlers, jwtSecret string) {
	protected := api.Group("/", middleware.AuthMiddleware(jwtSecret))
	protected.Get("/books", h.Book.GetAll)
	protected.Get("/books/:id", h.Book.GetById)
	protected.Post("/books", h.Book.Create)
	protected.Put("/books/:id", h.Book.Update)
	protected.Delete("/books/:id", h.Book.Delete)

}
