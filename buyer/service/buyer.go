package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/domain"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/repository"
	productDomain "github.com/paraizofelipe/luizalabs-challenge/product/domain"
)

type BuyerService struct {
	repository repository.BuyerRepository
}

func NewService(db *sqlx.DB) Service {
	return &BuyerService{
		repository: repository.NewPostgreRepository(db),
	}
}
func (s BuyerService) FindFavoriteProduct(buyerID string) (listProduct []productDomain.Product, err error) {
	return s.repository.FindFavoriteProduct(buyerID)
}

func (s BuyerService) FindAll() (listBuyer []domain.Buyer, err error) {
	return s.repository.FindAll()
}

func (s BuyerService) FindByEmail(email string) (buyer domain.Buyer, err error) {
	return s.repository.FindByEmail(email)
}

func (s BuyerService) FindByID(id string) (buyer domain.Buyer, err error) {
	if buyer, err = s.repository.FindByID(id); err != nil {
		return
	}
	if buyer.FavoriteProducts, err = s.FindFavoriteProduct(id); err != nil {
		return
	}

	return
}

func (s BuyerService) Add(buyer domain.Buyer) (err error) {
	return s.repository.Add(buyer)
}

func (s BuyerService) AddFavoriteProduct(buyerID string, productID string) (err error) {
	return s.repository.AddFavoriteProduct(buyerID, productID)
}

func (s BuyerService) Update(buyer domain.Buyer) (err error) {
	return s.repository.Update(buyer)
}

func (s BuyerService) RemoveByID(email string) (err error) {
	return s.repository.RemoveByID(email)
}
