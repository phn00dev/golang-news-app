package dto

import (
	"github.com/phn00dev/golang-news-app/internal/models"
	timeformat "github.com/phn00dev/golang-news-app/internal/utils/timeFormat"
)

type CategoryResponse struct {
	ID             int    `json:"id"`
	CategoryName   string `json:"category_name"`
	CategorySlug   string `json:"category_slug"`
	CategoryStatus string `json:"category_status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}

func GetOneCategoryResponse(category *models.Category) *CategoryResponse {
	if category == nil {
		return nil
	}
	return &CategoryResponse{
		ID:             category.ID,
		CategoryName:   category.CategoryName,
		CategorySlug:   category.CategorySlug,
		CategoryStatus: category.CategoryStatus,
		CreatedAt:      timeformat.FormatTime(category.CreatedAt),
		UpdatedAt:      timeformat.FormatTime(category.UpdatedAt),
		DeletedAt:      timeformat.FormatTime(category.UpdatedAt),
	}

}

func GetAllCategoryResponses(categories []models.Category) []CategoryResponse {
	categoryResponses := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		categoryResponses[i] = *GetOneCategoryResponse(&category)
	}
	return categoryResponses
}
