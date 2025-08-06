package middleware

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/jwt"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware validates JWT token
func AuthMiddleware(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Unauthorized(c, "Authorization header not found")
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Validate token
		claims, err := jwt.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			return response.Unauthorized(c, err.Error())
		}

		// Set user info in context for use in handler
		c.Locals("user_id", claims.UserID)
		c.Locals("user_email", claims.Email)
		c.Locals("user_role", claims.Role)

		return c.Next()
	}
}
