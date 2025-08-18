package handlers

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/schemas"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// AuthServiceInterface defines what auth handler needs from service
type AuthServiceInterface interface {
	Register(req *schemas.RegisterRequest) (*schemas.AuthResponse, error)
	Login(req *schemas.LoginRequest) (*schemas.AuthResponse, error)
	GetProfile(userID uint) (*schemas.UserResponse, error)
}

// AuthHandler handles http request for authentication
type AuthHandler struct {
	authService AuthServiceInterface
}

// NewAuthHandler create new AuthHandler instance
func NewAuthHandler(authService AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles POST /auth/register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req schemas.RegisterRequest

	// Parse JSON Body
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	// Validate request
	if errors := validator.ValidateStruct(req); errors != nil {
		return response.ValidationFailed(c, "Validation failed", errors)
	}

	// call service
	result, err := h.authService.Register(&req)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, "User registered successfully", result)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req schemas.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if errors := validator.ValidateStruct(req); errors != nil {
		return response.ValidationFailed(c, "Validation failed", errors)
	}

	result, err := h.authService.Login(&req)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "User logged in successfully", result)
}

func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return response.BadRequest(c, "User ID not found in context")
	}
	user, err := h.authService.GetProfile(userID)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "Profile retrieved successfully", user)
}
