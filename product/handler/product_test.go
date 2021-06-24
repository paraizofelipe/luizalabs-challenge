package handler

import (
	"bytes"
	"errors"
	"log"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paraizofelipe/luizalabs-challenge/product/domain"
	"github.com/paraizofelipe/luizalabs-challenge/product/service"
	"github.com/paraizofelipe/luizalabs-challenge/router"
)

func TestHandlerProductCreate(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		setupMock          func() *service.MockService
		requestBody        string
		expectedStatusCode int
		authorization      bool
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
				"title": "Galaxy S20",
				"brand": "Samsung",
				"image": "http://samsung.com/galaxys20.png",
				"price": 100.00,
				"review_score": 4.5
			}`,
			authorization:      true,
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
			authorization:      true,
			expectedStatusCode: 400,
		},
		{
			description: "error when creating the product",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Add(gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			requestBody: `{
				"title": "Galaxy S20",
				"brand": "Samsung",
				"image": "http://samsung.com/galaxys20.png",
				"price": 100.00,
				"review_score": 4.5
			}`,
			authorization:      true,
			expectedStatusCode: 500,
		},
		{
			description: "error not authorization",
			setupMock: func() *service.MockService {
				return nil
			},
			authorization:      false,
			expectedStatusCode: 403,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock()
				body    = bytes.NewBuffer([]byte(test.requestBody))
				w       = httptest.NewRecorder()
				req     = httptest.NewRequest("POST", "http://luizalabs.com/api/product", body)
			)

			handler := Product{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				Authorization: router.Authorization{
					Write: test.authorization,
				},
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

func TestHandlerProductUpdate(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		setupMock          func() *service.MockService
		requestID          string
		requestBody        string
		expectedStatusCode int
		authorization      bool
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
				"title": "Galaxy S21",
				"price": 150.00,
				"review_score": 4.5
			}`,
			authorization:      true,
			expectedStatusCode: 200,
		},
		{
			description: "Should return error when updating the product",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Update(gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			requestID: "e190a597-e7a3-4672-8a08-da3825e87244",
			requestBody: `{
				"title": "Galaxy S21",
				"price": 150.00,
				"review_score": 4.5
			}`,
			authorization:      true,
			expectedStatusCode: 500,
		},
		{
			description: "Should return error when updating the product",
			setupMock: func() *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().Update(gomock.Any()).Return(errors.New("")).AnyTimes()
				return service
			},
			authorization:      true,
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
				"title": "Galaxy S21",
				"price": 150.00,
				"review_score": 4.5
			}`,
			authorization:      true,
			expectedStatusCode: 400,
		},
		{
			description: "error not authorization",
			setupMock: func() *service.MockService {
				return nil
			},
			authorization:      false,
			expectedStatusCode: 403,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock()
				body    = bytes.NewBuffer([]byte(test.requestBody))
				w       = httptest.NewRecorder()
				req     = httptest.NewRequest("PATCH", "http://luizalabs.com/api/product", body)
			)

			handler := Product{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				Authorization: router.Authorization{
					Write: test.authorization,
				},
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

func TestHandlerProductRemove(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.LstdFlags|log.Lshortfile)

	tests := []struct {
		description        string
		setupMock          func(string) *service.MockService
		requestID          string
		expectedStatusCode int
		authorization      bool
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
			authorization:      true,
			expectedStatusCode: 200,
		},
		{
			description: "error when removing the product",
			requestID:   "000000000",
			setupMock: func(id string) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().RemoveByID(id).Return(errors.New("")).AnyTimes()
				return service
			},
			authorization:      true,
			expectedStatusCode: 400,
		},
		{
			description: "error not authorization",
			setupMock: func(id string) *service.MockService {
				return nil
			},
			authorization:      false,
			expectedStatusCode: 403,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock(test.requestID)
				w       = httptest.NewRecorder()
			)

			handler := Product{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				Authorization: router.Authorization{
					Write: test.authorization,
				},
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

func TestHandlerProductList(t *testing.T) {
	var logger = log.New(os.Stdout, "", log.Lshortfile)

	tests := []struct {
		description        string
		requestPage        int
		setupMock          func(int) *service.MockService
		expectedStatusCode int
		authorization      bool
	}{
		{
			description: "simple test",
			requestPage: 1,
			setupMock: func(page int) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().ListByPage(page-1).Return([]domain.Product{}, nil).AnyTimes()
				return service
			},
			authorization:      true,
			expectedStatusCode: 200,
		},
		{
			description: "error of invalid page",
			requestPage: 0,
			setupMock: func(id int) *service.MockService {
				return &service.MockService{}
			},
			authorization:      true,
			expectedStatusCode: 400,
		},
		{
			description: "error when listing product",
			requestPage: 1,
			setupMock: func(page int) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().ListByPage(page-1).Return([]domain.Product{}, errors.New("")).AnyTimes()
				return service
			},
			authorization:      true,
			expectedStatusCode: 500,
		},
		{
			description: "error not authorization",
			requestPage: 1,
			setupMock: func(page int) *service.MockService {
				return nil
			},
			authorization:      false,
			expectedStatusCode: 403,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock(test.requestPage)
				w       = httptest.NewRecorder()
			)

			handler := Product{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				Authorization: router.Authorization{
					Read: test.authorization,
				},
				ResponseWriter: w,
				QueryString: url.Values{
					"page": []string{strconv.Itoa(test.requestPage)},
				},
			}

			handler.list(&ctx)
			resp := w.Result()

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("cuttent: %d ---> expected: %d", resp.StatusCode, test.expectedStatusCode)
			}
		})
	}
}

func TestHandlerProductDetail(t *testing.T) {
	var logger = log.New(&bytes.Buffer{}, "", log.Lshortfile)

	tests := []struct {
		description        string
		requestID          string
		setupMock          func(string) *service.MockService
		expectedStatusCode int
		authorization      bool
	}{
		{
			description: "simple test",
			requestID:   "e190a597-e7a3-4672-8a08-da3825e87244",
			setupMock: func(id string) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().FindByID(id).Return(domain.Product{}, nil).AnyTimes()
				return service
			},
			authorization:      true,
			expectedStatusCode: 200,
		},

		{
			description: "error of invalid ID",
			requestID:   "000000000",
			setupMock: func(id string) *service.MockService {
				return &service.MockService{}
			},
			authorization:      true,
			expectedStatusCode: 400,
		},

		{
			description: "error when fetching product",
			requestID:   "e190a597-e7a3-4672-8a08-da3825e87244",
			setupMock: func(id string) *service.MockService {
				ctrl := gomock.NewController(t)
				service := service.NewMockService(ctrl)
				service.EXPECT().FindByID(id).Return(domain.Product{}, errors.New("")).AnyTimes()
				return service
			},
			authorization:      true,
			expectedStatusCode: 500,
		},

		{
			description: "error not authorization",
			setupMock: func(id string) *service.MockService {
				return nil
			},
			authorization:      false,
			expectedStatusCode: 403,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var (
				service = test.setupMock(test.requestID)
				w       = httptest.NewRecorder()
			)

			handler := Product{
				Logger:  logger,
				Service: service,
			}

			ctx := router.Context{
				Authorization: router.Authorization{
					Read: test.authorization,
				},
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
