package repository

import "github.com/paraizofelipe/luizalabs-challenge/buyer/domain"

type Reader interface {
	FindAll() ([]domain.Buyer, error)
	FindByEmail(email string) (domain.Buyer, error)
	FindByID(id string) (domain.Buyer, error)
}

type Writer interface {
	Add(domain.Buyer) error
	Update(domain.Buyer) error
	RemoveByID(id string) error
	AddFavoriteProduct(id string, productID string) error
}

type BuyerRepository interface {
	Reader
	Writer
}
