package repositories

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/database"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/models"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/utils"
)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) Create(book *models.Book) error {
	return database.DB.Create(&book).Error
}

func (r *BookRepository) GetAll(params *utils.PaginationParams) ([]*models.Book, int64, error) {
	var books []*models.Book
	var total int64
	query := database.DB.Model(&models.Book{})

	// Search functionality
	if params.Search != "" {
		query = query.Where("title ILIKE ? or author ILIKE ?", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	// Count total
	query.Count(&total)

	// Apply pagination and sorting
	err := query.Preload("User").Order(params.Sort + " " + params.Order).
		Offset(params.GetOffset()).
		Limit(params.Size).
		Find(&books).Error

	return books, total, err
}

func (r *BookRepository) GetById(id uint) (*models.Book, error) {
	var book models.Book
	err := database.DB.Preload("User").Where("id = ?", id).First(&book).Error
	return &book, err
}

func (r *BookRepository) Update(id uint, book *models.Book) error {
	return database.DB.Model(&models.Book{}).Where("id = ?", id).Updates(&book).Error
}

func (r *BookRepository) Delete(id uint) error {
	return database.DB.Where("id = ?", id).Delete(&models.Book{}).Error
}
