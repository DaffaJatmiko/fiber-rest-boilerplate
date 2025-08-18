package schemas

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/models"
	"time"
)

type UpdateUserRequest struct {
	Email    string `json:"email" validate:"omitempty,email"`
	Username string `json:"username" validate:"omitempty,min=2"`
	Password string `json:"password" validate:"omitempty,min=8"`
	Role     string `json:"role" validate:"omitempty,oneof=USER ADMIN"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Helper function that convert model to response

func UserToResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
