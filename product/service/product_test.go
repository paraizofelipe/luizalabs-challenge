package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paraizofelipe/luizalabs-challenge/product/domain"
	"github.com/paraizofelipe/luizalabs-challenge/product/repository"
)

func TestProductServiceAdd(t *testing.T) {
	var (
		ctrl    = gomock.NewController(t)
		service ProdutService
	)

	tests := []struct {
		description string
		setupMock   func(domain.Product)
		in          domain.Product
		hasError    bool
	}{
		{
			description: "simple test",
			in: domain.Product{
				Title: "title teste",
				Brand: "brand teste",
			},
			setupMock: func(p domain.Product) {
				repo := repository.NewMockProductRepository(ctrl)
				repo.EXPECT().Add(p).Return(nil).AnyTimes()
				service.repository = ProdutService{
					repository: repo,
				}
			},
		},
		{
			description: "error when find favorite products",
			in:          domain.Product{},
			setupMock: func(p domain.Product) {
				repo := repository.NewMockProductRepository(ctrl)
				repo.EXPECT().Add(p).Return(errors.New("")).AnyTimes()
				service.repository = ProdutService{
					repository: repo,
				}
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			test.setupMock(test.in)
			err := service.Add(test.in)

			if test.hasError && err == nil {
				t.Error(err)
			}
		})
	}
}
