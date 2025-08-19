package schemas

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/models"
	"time"
)

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Author      string `json:"author" validate:"required,min=1,max=200"`
	Description string `json:"desc" validate:"omitempty,max=1000"`
}

type UpdateBookRequest struct {
	Title       string `json:"title" validate:"omitempty,min=1,max=200"`
	Author      string `json:"author" validate:"omitempty,min=1,max=200"`
	Description string `json:"desc" validate:"omitempty,max=1000"`
}

type BookResponse struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Author    string       `json:"author"`
	Desc      string       `json:"desc"`
	UserID    uint         `json:"user_id"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

// Helper function that convert model to response
func BookToResponse(book *models.Book) BookResponse {

	response := BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Desc:      book.Desc,
		UserID:    book.UserID,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	// Cek nil dulu
	if book.User != nil {
		response.User = UserToResponse(book.User)
	}

	return response
}
