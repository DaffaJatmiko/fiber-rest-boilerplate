package services

import (
	"errors"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/models"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/internal/schemas"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/utils"
)

// BookRepositoryInterface defines what BookService needs from repository
type BookRepositoryInterface interface {
	Create(book *models.Book) error
	GetAll(params *utils.PaginationParams) ([]*models.Book, int64, error)
	GetById(id uint) (*models.Book, error)
	Update(id uint, book *models.Book) error
	Delete(id uint) error
}

// BookService handles book management logic
type BookService struct {
	bookRepo BookRepositoryInterface
}

// NewBookService create a new BookService instance
func NewBookService(bookRepo BookRepositoryInterface) *BookService {
	return &BookService{bookRepo: bookRepo}
}

func (s *BookService) Create(req *schemas.CreateBookRequest, userId uint) (*schemas.BookResponse, error) {

	// Create a book model
	book := &models.Book{
		Title:  req.Title,
		Author: req.Author,
		Desc:   req.Description,
		UserID: userId,
	}

	// save to database
	err := s.bookRepo.Create(book)
	if err != nil {
		return nil, errors.New("could not create book")
	}

	// Reload book with user data
	bookWithUser, err := s.bookRepo.GetById(book.UserID)
	if err != nil {
		return nil, errors.New("could not get book")
	}

	response := schemas.BookToResponse(bookWithUser)
	return &response, nil
}

func (s *BookService) GetAll(params *utils.PaginationParams) ([]schemas.BookResponse, *response.Pagination, error) {
	// set default value
	params.GetDefaults()

	// get book from repository
	books, total, err := s.bookRepo.GetAll(params)
	if err != nil {
		return nil, nil, err
	}

	// convert format response
	bookResponses := make([]schemas.BookResponse, 0)
	for _, book := range books {
		bookResponses = append(bookResponses, schemas.BookToResponse(book))
	}

	// calculate pagination
	pagination := utils.CalculatePagination(params.Page, params.Size, total)
	return bookResponses, pagination, nil
}

func (s *BookService) GetById(id uint) (*schemas.BookResponse, error) {
	// get book by id from repository
	book, err := s.bookRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	response := schemas.BookToResponse(book)
	return &response, nil
}

func (s *BookService) Update(id uint, req *schemas.UpdateBookRequest) (*schemas.BookResponse, error) {
	// Get book by id
	book, err := s.bookRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	// update fields if provide
	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}
	if req.Description != "" {
		book.Desc = req.Description
	}

	// save update to repository
	err = s.bookRepo.Update(id, book)
	if err != nil {
		return nil, err
	}

	response := schemas.BookToResponse(book)
	return &response, nil
}

func (s *BookService) Delete(id uint) error {
	// get book by id
	_, err := s.bookRepo.GetById(id)
	if err != nil {
		return err
	}

	err = s.bookRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
