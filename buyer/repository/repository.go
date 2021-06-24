package repository

import (
	"github.com/paraizofelipe/luizalabs-challenge/buyer/domain"
	productDomain "github.com/paraizofelipe/luizalabs-challenge/product/domain"
)

type Reader interface {
	FindAll() ([]domain.Buyer, error)
	FindByEmail(email string) (domain.Buyer, error)
	FindByID(id string) (domain.Buyer, error)
	FindFavoriteProduct(buyerID string) ([]productDomain.Product, error)
}

type Writer interface {
	Add(domain.Buyer) error
	AddFavoriteProduct(buyerID string, productID string) error
	Update(domain.Buyer) error
	RemoveByID(id string) error
}

type BuyerRepository interface {
	Reader
	Writer
}
