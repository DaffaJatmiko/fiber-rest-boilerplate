package response

import (
	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorData  `json:"error,omitempty"`
}

type ErrorData struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Details []ValidationError `json:"details,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page       int  `json:"page"`
	Size       int  `json:"size"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

// Success response helper
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created response helper
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// BadRequest response helper
func BadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(BaseResponse{
		Success: false,
		Message: message,
		Error: &ErrorData{
			Code:    fiber.StatusBadRequest,
			Message: message,
		},
	})
}

func ValidationFailed(c *fiber.Ctx, message string, details []ValidationError) error {
	return c.Status(fiber.StatusBadRequest).JSON(BaseResponse{
		Success: false,
		Message: message,
		Error: &ErrorData{
			Code:    fiber.StatusBadRequest,
			Message: message,
			Details: details,
		},
	})
}

// Paginated response helper
func Paginated(c *fiber.Ctx, message string, data interface{}, pagination Pagination) error {
	// safety check for nil slices

	return c.Status(fiber.StatusOK).JSON(PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

// Unauthorized response helper
func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(BaseResponse{
		Success: false,
		Message: message,
		Error: &ErrorData{
			Code:    fiber.StatusUnauthorized,
			Message: message,
		},
	})
}

// InternalError response helper - NEW for global error handler
func InternalError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(BaseResponse{
		Success: false,
		Message: message,
		Error: &ErrorData{
			Code:    fiber.StatusInternalServerError,
			Message: message,
		},
	})
}
