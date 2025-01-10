package dto

type CreateCategoryRequest struct {
	CategoryName   string `json:"category_name"`
	CategorySlug   string `json:"category_slug"`
	CategoryStatus string `json:"category_status"`
}

type UpdateCategoryRequest struct {
	CategoryName   string `json:"category_name"`
	CategorySlug   string `json:"category_slug"`
	CategoryStatus string `json:"category_status"`
}
