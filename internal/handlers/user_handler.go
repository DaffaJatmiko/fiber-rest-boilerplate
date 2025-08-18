package handlers

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/schemas"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/utils"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// UserServiceInterface defines what user handler needs from service
type UserServiceInterface interface {
	GetAll(params *utils.PaginationParams) ([]schemas.UserResponse, *response.Pagination, error)
	GetByID(id uint) (*schemas.UserResponse, error)
	Update(id uint, req *schemas.UpdateUserRequest) (*schemas.UserResponse, error)
	Delete(id uint) error
}

// UserHandler handles http request for user management
type UserHandler struct {
	userService UserServiceInterface
}

// NewUserHandler create new UserHandler instance
func NewUserHandler(userService UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil || size < 1 {
		size = 10
	}

	params := &utils.PaginationParams{
		Page:   page,
		Size:   size,
		Sort:   c.Query("sort", ""),
		Order:  c.Query("order", ""),
		Search: c.Query("search", ""),
	}

	users, pagination, err := h.userService.GetAll(params)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Paginated(c, "Users rerieved successfully", users, *pagination)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	// Parse int
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil || idInt <= 0 {
		return response.BadRequest(c, "Invalid ID")
	}

	// Convert int ke uint
	id := uint(idInt)

	user, err := h.userService.GetByID(id)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "User retrieved successfully", user)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	// parse id int
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil || idInt <= 0 {
		return response.BadRequest(c, "Invalid ID")
	}

	var req schemas.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if errors := validator.ValidateStruct(req); errors != nil {
		return response.ValidationFailed(c, "Validation failed", errors)
	}

	id := uint(idInt)

	user, err := h.userService.Update(id, &req)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "User updated successfully", user)

}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil || idInt <= 0 {
		return response.BadRequest(c, "Invalid ID")
	}

	id := uint(idInt)

	error := h.userService.Delete(id)
	if error != nil {
		return response.BadRequest(c, error.Error())
	}

	return response.Success(c, "User deleted successfully", nil)
}
