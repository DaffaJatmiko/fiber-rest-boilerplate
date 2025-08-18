package services

import (
	"errors"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/models"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/schemas"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/jwt"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/utils"
)

// AuthService handles authentication business logic
type AuthService struct {
	userRepo  UserRepositoryInterface
	jwtSecret string
}

// NewAuthService create new AuthService instance
func NewAuthService(userRepo UserRepositoryInterface, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Register handles user registration
func (s *AuthService) Register(req *schemas.RegisterRequest) (*schemas.AuthResponse, error) {
	// Check if the user already exists
	_, err := s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("could not hash password")
	}

	// Create a user model
	user := &models.User{
		Email:    req.Email,
		Name:     req.Username,
		Password: hashedPassword,
		Role:     "USER",
	}

	// save to a database
	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("could not create user")
	}

	// Generate jwt token
	token, err := jwt.GenerateToken(user.ID, user.Email, user.Password, s.jwtSecret)
	if err != nil {
		return nil, errors.New("could not generate token")
	}

	// return response
	return &schemas.AuthResponse{
		Token: token,
		User:  schemas.UserToResponse(user),
	}, nil
}

// Login handles user login
func (s *AuthService) Login(req *schemas.LoginRequest) (*schemas.AuthResponse, error) {
	// Find user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("could not find user")
	}

	// check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	// generate jwt token
	token, err := jwt.GenerateToken(user.ID, user.Email, user.Password, s.jwtSecret)
	if err != nil {
		return nil, errors.New("could not generate token")
	}

	return &schemas.AuthResponse{
		Token: token,
		User:  schemas.UserToResponse(user),
	}, nil
}

// GetProfile handles get user profile
func (s *AuthService) GetProfile(userID uint) (*schemas.UserResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("could not find user")
	}

	response := schemas.UserToResponse(user)
	return &response, nil
}
