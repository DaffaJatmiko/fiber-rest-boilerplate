package services

import (
	"errors"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/models"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/schemas"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/utils"
)

// UserRepositoryInterface defines what AuthService and UserService needs from repository
type UserRepositoryInterface interface {
	// Auth needs

	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint) (*models.User, error)

	// Management needs

	Update(id uint, user *models.User) error
	Delete(id uint) error
	GetAll(params *utils.PaginationParams) ([]*models.User, int64, error)
}

// UserService handles user management logic
type UserService struct {
	userRepo UserRepositoryInterface
}

// NewUserService crate a new UserService instance
func NewUserService(userRepo UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAll(params *utils.PaginationParams) ([]schemas.UserResponse, *response.Pagination, error) {
	// set default value
	params.GetDefaults()

	// get users from repository
	users, total, err := s.userRepo.GetAll(params)
	if err != nil {
		return nil, nil, err
	}

	// convert ke response format
	var userResponses []schemas.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, schemas.UserToResponse(user))
	}

	// Calculate pagination
	pagination := utils.CalculatePagination(params.Page, params.Size, total)

	return userResponses, pagination, nil
}

func (s *UserService) GetByID(id uint) (*schemas.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	response := schemas.UserToResponse(user)
	return &response, nil
}

func (s *UserService) Update(id uint, req *schemas.UpdateUserRequest) (*schemas.UserResponse, error) {
	// Get user by id
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	// update field if provide
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Username != "" {
		user.Name = req.Username
	}
	if req.Role != "" && models.IsValidRole(req.Role) {
		user.Role = req.Role
	}
	if req.Password != "" {
		hashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashPassword
	}

	// save to database
	if err := s.userRepo.Update(id, user); err != nil {
		return nil, errors.New("user update failed")
	}

	response := schemas.UserToResponse(user)
	return &response, nil
}

func (s *UserService) Delete(id uint) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return errors.New("User not found")
	}

	if err := s.userRepo.Delete(id); err != nil {
		return errors.New("user delete failed")
	}

	return nil
}
