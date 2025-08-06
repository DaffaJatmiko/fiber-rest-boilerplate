package routes

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/config"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/handlers"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/repositories"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/services"
)

// Handlers holds all application handlers
type Handlers struct {
	Auth *handlers.AuthHandler
	//User *handlers.UserHandler
	//Book *handlers.BookHandler

	// Easy to add more handlers:
	// Order *handlers.OrderHandler
	// Payment *handlers.PaymentHandler
}

// NewHandlers creates and initializes all application handlers with their dependencies
func NewHandlers(cfg *config.Config) *Handlers {
	// Initialize repositories (data layer)
	userRepo := repositories.NewUserRepository()
	//bookRepo := repositories.NewBookRepository()

	// Initialize services (business layer)
	authService := services.NewAuthService(userRepo, cfg.JWT.Secret)
	//userService := services.NewUserService(userRepo)
	//bookService := services.NewBookService(bookRepo, userRepo)

	// Initialize handlers (presentation layer)
	return &Handlers{
		Auth: handlers.NewAuthHandler(authService),
		//User: handlers.NewUserHandler(userService),
		//Book: handlers.NewBookHandler(bookService),
	}
}
