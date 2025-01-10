package models

import "time"

type Post struct {
	ID         int       `json:"id"`
	PostTitle  string    `json:"post_title"`
	PostSlug   string    `json:"post_slug"`
	PostDesc   string    `json:"post_desc"`
	PostImage  string    `json:"post_image"`
	PostStatus string    `json:"post_status"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
