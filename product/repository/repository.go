package repository

import "github.com/paraizofelipe/luizalabs-challenge/product/domain"

type Reader interface {
	ListByPage(page int) ([]domain.Product, error)
	FindByID(id string) (domain.Product, error)
	FindByTitleAndBrand(brand string, title string) (domain.Product, error)
}

type Writer interface {
	Add(domain.Product) error
	Update(domain.Product) error
	RemoveByID(id string) error
}

type ProductRepository interface {
	Reader
	Writer
}
