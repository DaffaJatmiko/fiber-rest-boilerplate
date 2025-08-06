package repositories

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/database"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/models"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/utils"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) Update(id uint, user *models.User) error {
	return database.DB.Model(user).Where("id = ?", id).Updates(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}

func (r *UserRepository) GetAll(params *utils.PaginationParams) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := database.DB.Model(&models.User{})

	// Search functionality
	if params.Search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	// Count total
	query.Count(&total)

	// Apply pagination and sorting
	err := query.Order(params.Sort + " " + params.Order).
		Offset(params.GetOffset()).
		Limit(params.Size).
		Find(&users).Error

	return users, total, err
}
