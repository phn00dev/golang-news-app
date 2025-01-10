package models

import "time"

type Category struct {
	ID             int       `json:"id"`
	CategoryName   string    `json:"category_name"`
	CategorySlug   string    `json:"category_slug"`
	CategoryStatus string    `json:"category_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
