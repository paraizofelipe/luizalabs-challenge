package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Image       string    `json:"image" db:"image"`
	Price       float64   `json:"price" db:"price"`
	Brand       string    `json:"brand" db:"brand"`
	ReviewScore int       `json:"review_score" db:"review_score"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (d Product) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(b)
}
