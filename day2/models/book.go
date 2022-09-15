package models

import (
	"time"
)

type (
	Books struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Title     string    `json:"title" validate:"required"`
		Author    string    `json:"author" validate:"required"`
		Price     int       `json:"price" validate:"required"`
	}
)
