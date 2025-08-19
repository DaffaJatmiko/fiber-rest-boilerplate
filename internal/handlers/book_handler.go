package handlers

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/schemas"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/utils"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// BookServiceInterface defines what book handler need from service
type BookServiceInterface interface {
	Create(req *schemas.CreateBookRequest, userId uint) (*schemas.BookResponse, error)
	GetById(id uint) (*schemas.BookResponse, error)
	GetAll(params *utils.PaginationParams) ([]schemas.BookResponse, *response.Pagination, error)
	Update(id uint, req *schemas.UpdateBookRequest) (*schemas.BookResponse, error)
	Delete(id uint) error
}

// BookHandler handles http request for book management
type BookHandler struct {
	bookService BookServiceInterface
}

// NewBookHandler create new BookHandler instance
func NewBookHandler(bookService BookServiceInterface) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) Create(c *fiber.Ctx) error {
	// Get user id from context
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return response.BadRequest(c, "User ID is not in context.")
	}

	var req schemas.CreateBookRequest

	// Parse JSON Body
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	// validate request
	if errors := validator.ValidateStruct(req); errors != nil {
		return response.ValidationFailed(c, "Validation failed", errors)
	}

	// call service
	book, err := h.bookService.Create(&req, userID)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "Success create book", book)

}

func (h *BookHandler) GetById(c *fiber.Ctx) error {
	// parse int
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	// convert int ke uint
	id := uint(idInt)

	// get book from service
	book, err := h.bookService.GetById(id)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "Success get book", book)
}

func (h *BookHandler) GetAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	params := &utils.PaginationParams{
		Page:   page,
		Size:   size,
		Sort:   c.Query("sort", ""),
		Order:  c.Query("order", ""),
		Search: c.Query("search", ""),
	}

	books, pagination, err := h.bookService.GetAll(params)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Paginated(c, "Books retrieved successfully", books, *pagination)
}

func (h *BookHandler) Update(c *fiber.Ctx) error {
	// get int userId
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	var req schemas.UpdateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if errors := validator.ValidateStruct(req); errors != nil {
		return response.ValidationFailed(c, "Validation failed", errors)
	}

	id := uint(idInt)

	// get book by id
	book, err := h.bookService.Update(id, &req)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "Success update book", book)
}

func (h *BookHandler) Delete(c *fiber.Ctx) error {
	// get int userId
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	id := uint(idInt)

	err = h.bookService.Delete(id)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "Success delete book", id)
}
