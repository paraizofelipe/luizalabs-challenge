package service

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/product/domain"
	"github.com/paraizofelipe/luizalabs-challenge/product/repository"
)

type ProdutService struct {
	repository repository.ProductRepository
}

func NewService(db *sqlx.DB) Service {
	return &ProdutService{
		repository: repository.NewPostgreRepository(db),
	}
}

func (s ProdutService) ListByPage(page int) (listProduct []domain.Product, err error) {
	return s.repository.ListByPage(page)
}

func (s ProdutService) FindByID(id string) (product domain.Product, err error) {
	return s.repository.FindByID(id)
}

func (s ProdutService) FindByTitleAndBrand(brand string, title string) (product domain.Product, err error) {
	return s.repository.FindByTitleAndBrand(brand, title)
}

func (s ProdutService) Add(product domain.Product) (err error) {
	var p domain.Product

	if p, err = s.FindByTitleAndBrand(product.Brand, product.Title); err != nil {
		return
	}
	if p.Brand != "" && p.Title != "" {
		err = errors.New("Product already registered")
		return
	}

	return s.repository.Add(product)
}

func (s ProdutService) Update(product domain.Product) (err error) {
	return s.repository.Update(product)
}

func (s ProdutService) RemoveByID(id string) (err error) {
	return s.repository.RemoveByID(id)
}
