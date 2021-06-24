package handler

import (
	"bytes"
	"errors"
	"log"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/domain"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/service"
	"github.com/paraizofelipe/luizalabs-challenge/router"
)

func TestDetail(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.Lshortfile)

	tests := []struct {
		description        string
		requestID          string
		setupMock          func(string) *service.MockService
		expectedStatusCode int
	}{
		{
			description: "simple test",
			requestID:   "e190a597-e7a3-4672-8a08-da3825e87244",
			setupMock: func(id string) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().FindByID(id).Return(domain.Buyer{}, nil).AnyTimes()
				return service
			},
			expectedStatusCode: 200,
		},

		{
			description: "error of invalid ID",
			requestID:   "000000000",
			setupMock: func(id string) *service.MockService {
				return &service.MockService{}
			},
			expectedStatusCode: 500,
		},

		{
			description: "error when fetching buyer",
			requestID:   "e190a597-e7a3-4672-8a08-da3825e87244",
			setupMock: func(id string) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().FindByID(id).Return(domain.Buyer{}, errors.New("")).AnyTimes()
				return service
			},
			expectedStatusCode: 500,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock(test.requestID)
				w       = httptest.NewRecorder()
			)

			handler := Buyer{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				ResponseWriter: w,
				Params: map[string]string{
					"id": test.requestID,
				},
			}

			handler.detail(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		setupMock          func() *service.MockService
		requestBody        string
		expectedStatusCode int
	}{
		{
			description: "simple test",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Add(gomock.Any()).Return(nil).AnyTimes()
				return service
			},
			requestBody: `{
				"name":  "Fulano",
				"email": "fulano@gmail.com"
			}`,
			expectedStatusCode: 201,
		},
		{
			description: "error when parser json of body",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Add(gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			expectedStatusCode: 400,
		},
		{
			description: "error when creating the buyer",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Add(gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			requestBody: `{
				"name":  "Fulano",
				"email": "fulano@gmail.com"
			}`,
			expectedStatusCode: 500,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock()
				body    = bytes.NewBuffer([]byte(test.requestBody))
				w       = httptest.NewRecorder()
				req     = httptest.NewRequest("POST", "http://luizalabs.com/api/buyer", body)
			)

			handler := Buyer{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				ResponseWriter: w,
				Request:        req,
			}

			handler.create(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		setupMock          func(string) *service.MockService
		requestID          string
		expectedStatusCode int
	}{
		{
			description: "simple test",
			requestID:   "e190a597-e7a3-4672-8a08-da3825e87244",
			setupMock: func(id string) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().RemoveByID(id).Return(nil).AnyTimes()
				return service
			},
			expectedStatusCode: 200,
		},
		{
			description: "error when removing the buyer",
			requestID:   "000000000",
			setupMock: func(id string) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().RemoveByID(id).Return(errors.New("")).AnyTimes()
				return service
			},
			expectedStatusCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock(test.requestID)
				w       = httptest.NewRecorder()
			)

			handler := Buyer{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				ResponseWriter: w,
				Params:         map[string]string{"id": test.requestID},
			}

			handler.remove(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		setupMock          func() *service.MockService
		requestID          string
		requestBody        string
		expectedStatusCode int
	}{
		{
			description: "Simple test",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Update(gomock.Any()).Return(nil).AnyTimes()
				return service
			},
			requestID: "e190a597-e7a3-4672-8a08-da3825e87244",
			requestBody: `{
				"name": "fulano",
				"email": "fulano@gmail.com"
			}`,
			expectedStatusCode: 200,
		},
		{
			description: "Should return error when updating the buyer",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Update(gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			requestID: "e190a597-e7a3-4672-8a08-da3825e87244",
			requestBody: `{
				"name": "fulano",
				"email": "fulano@gmail.com"
			}`,
			expectedStatusCode: 500,
		},
		{
			description: "Should return error when updating the buyer",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Update(gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			expectedStatusCode: 400,
		},
		{
			description: "Should return error invalid ID",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Update(gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			requestID: "00000000",
			requestBody: `{
				"name": "fulano",
				"email": "fulano@gmail.com"
			}`,
			expectedStatusCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock()
				body    = bytes.NewBuffer([]byte(test.requestBody))
				w       = httptest.NewRecorder()
				req     = httptest.NewRequest("PATCH", "http://luizalabs.com/api/buyer", body)
			)

			handler := Buyer{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				ResponseWriter: w,
				Params:         map[string]string{"id": test.requestID},
				Request:        req,
			}

			handler.update(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}

func TestAddFavoriteProduct(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		setupMock          func() *service.MockService
		requestBuyerID     string
		requestProductID   string
		expectedStatusCode int
	}{
		{
			description: "Simple test",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().AddFavoriteProduct(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				return service
			},
			requestBuyerID:     "e190a597-e7a3-4672-8a08-da3825e87244",
			requestProductID:   "7c40ba69-5d12-458a-aa99-8ef4bfaf8180",
			expectedStatusCode: 200,
		},
		{
			description: "Should return error when updating the buyer",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().AddFavoriteProduct(gomock.Any(), gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			requestBuyerID:     "e190a597-e7a3-4672-8a08-da3825e87244",
			requestProductID:   "7c40ba69-5d12-458a-aa99-8ef4bfaf8180",
			expectedStatusCode: 500,
		},
		{
			description: "Should return error invalid buyer ID",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().AddFavoriteProduct(gomock.Any(), gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			requestBuyerID:     "",
			requestProductID:   "e190a597-e7a3-4672-8a08-da3825e87244",
			expectedStatusCode: 500,
		},
		{
			description: "Should return error invalid product ID",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().AddFavoriteProduct(gomock.Any(), gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			requestBuyerID:     "e190a597-e7a3-4672-8a08-da3825e87244",
			requestProductID:   "",
			expectedStatusCode: 500,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock()
				w       = httptest.NewRecorder()
			)

			handler := Buyer{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				ResponseWriter: w,
				Params: map[string]string{
					"id":         test.requestBuyerID,
					"id_product": test.requestProductID,
				},
			}

			handler.addFavoriteProduct(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}
