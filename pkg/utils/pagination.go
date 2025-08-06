package utils

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"math"
)

// PaginationParams untuk handle query parameters dari request
type PaginationParams struct {
	Page   int    `query:"page" validate:"min=1"`
	Size   int    `query:"size" validate:"min=1,max=100"`
	Sort   string `query:"sort"`
	Order  string `query:"order" validate:"oneof=asc desc"`
	Search string `query:"search"`
}

// GetOffset menghitung offset untuk database query
func (p *PaginationParams) GetOffset() int {
	return (p.Page - 1) * p.Size
}

// GetDefaults set default values jika tidak ada di request
func (p *PaginationParams) GetDefaults() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Size == 0 {
		p.Size = 10
	}
	if p.Sort == "" {
		p.Sort = "id"
	}
	if p.Order == "" {
		p.Order = "asc"
	}
}

// CalculatePagination helper untuk convert ke response.Pagination
func CalculatePagination(page, size int, total int64) *response.Pagination {
	totalPages := int(math.Ceil(float64(total) / float64(size)))

	return &response.Pagination{
		Page:       page,
		Size:       size,
		Total:      int(total), // Convert int64 to int
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
}
