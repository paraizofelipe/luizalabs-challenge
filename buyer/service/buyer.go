package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/domain"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/repository"
)

type BuyerService struct {
	repository repository.BuyerRepository
}

func NewService(db *sqlx.DB) Service {
	return &BuyerService{
		repository: repository.NewPostgreRepository(db),
	}
}

func (s BuyerService) FindAll() (listBuyer []domain.Buyer, err error) {
	return s.repository.FindAll()
}

func (s BuyerService) FindByEmail(email string) (buyer domain.Buyer, err error) {
	return s.repository.FindByEmail(email)
}

func (s BuyerService) Add(buyer domain.Buyer) (err error) {
	return s.repository.Add(buyer)
}

func (s BuyerService) Update(buyer domain.Buyer) (err error) {
	return s.repository.Update(buyer)
}

func (s BuyerService) RemoveByEmail(email string) (err error) {
	return s.repository.RemoveByEmail(email)
}
