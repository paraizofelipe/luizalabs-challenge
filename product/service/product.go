package service

import (
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

func (s ProdutService) FindAll() (listProduct []domain.Product, err error) {
	return s.repository.FindAll()
}

func (s ProdutService) FindByID(email string) (product domain.Product, err error) {
	return s.repository.FindByID(email)
}

func (s ProdutService) Add(product domain.Product) (err error) {
	return s.repository.Add(product)
}

func (s ProdutService) Update(product domain.Product) (err error) {
	return s.repository.Update(product)
}

func (s ProdutService) RemoveByID(id string) (err error) {
	return s.repository.RemoveByID(id)
}
