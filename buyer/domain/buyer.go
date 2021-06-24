package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/paraizofelipe/luizalabs-challenge/product/domain"
)

type Buyer struct {
	ID               uuid.UUID        `json:"id" db:"id"`
	Name             string           `json:"name" db:"name"`
	Email            string           `json:"email" db:"email"`
	FavoriteProducts []domain.Product `json:"favorite_products,omitempty" db:"favorite_products"`
	CreatedAt        time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at" db:"updated_at"`
}

func (d Buyer) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}

	return string(b)
}
