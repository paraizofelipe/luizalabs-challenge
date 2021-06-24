package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/domain"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/repository"
	productDomain "github.com/paraizofelipe/luizalabs-challenge/product/domain"
)

func TestFindByID(t *testing.T) {
	var (
		ctrl    = gomock.NewController(t)
		service BuyerService
	)

	tests := []struct {
		description string
		in          string
		setupMock   func(string)
		expected    domain.Buyer
		hasError    bool
	}{
		{
			description: "simple test",
			in:          "1a6dbc7f-7dee-4122-b9b8-abaf5a97eed7",
			setupMock: func(id string) {
				repo := repository.NewMockBuyerRepository(ctrl)
				repo.EXPECT().FindByID(id).Return(domain.Buyer{}, nil).AnyTimes()
				repo.EXPECT().FindFavoriteProduct(id).Return([]productDomain.Product{}, nil).AnyTimes()
				service.repository = BuyerService{
					repository: repo,
				}
			},
			expected: domain.Buyer{},
			hasError: false,
		},
		{
			description: "error when find buyer",
			in:          "1a6dbc7f-7dee-4122-b9b8-abaf5a97eed7",
			setupMock: func(id string) {
				repo := repository.NewMockBuyerRepository(ctrl)
				repo.EXPECT().FindByID(id).Return(domain.Buyer{}, errors.New("")).AnyTimes()
				service.repository = BuyerService{
					repository: repo,
				}
			},
			expected: domain.Buyer{},
			hasError: true,
		},
		{
			description: "error when find favorite products",
			in:          "1a6dbc7f-7dee-4122-b9b8-abaf5a97eed7",
			setupMock: func(id string) {
				repo := repository.NewMockBuyerRepository(ctrl)
				repo.EXPECT().FindByID(id).Return(domain.Buyer{}, nil).AnyTimes()
				repo.EXPECT().FindFavoriteProduct(id).Return([]productDomain.Product{}, errors.New("")).AnyTimes()
				service.repository = BuyerService{
					repository: repo,
				}
			},
			expected: domain.Buyer{},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			test.setupMock(test.in)
			currentBuyer, err := service.FindByID(test.in)
			if test.hasError && err == nil {
				t.Error(err)
			}

			if currentBuyer.String() != test.expected.String() {
				t.Errorf("current: %v ---> expected: %v", currentBuyer, test.expected)
			}
		})
	}
}
