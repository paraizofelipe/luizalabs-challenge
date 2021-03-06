// Code generated by MockGen. DO NOT EDIT.
// Source: ./buyer/service/service.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/paraizofelipe/luizalabs-challenge/buyer/domain"
	domain0 "github.com/paraizofelipe/luizalabs-challenge/product/domain"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockService) Add(arg0 domain.Buyer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockServiceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockService)(nil).Add), arg0)
}

// AddFavoriteProduct mocks base method.
func (m *MockService) AddFavoriteProduct(buyerID, productID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavoriteProduct", buyerID, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavoriteProduct indicates an expected call of AddFavoriteProduct.
func (mr *MockServiceMockRecorder) AddFavoriteProduct(buyerID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavoriteProduct", reflect.TypeOf((*MockService)(nil).AddFavoriteProduct), buyerID, productID)
}

// FindAll mocks base method.
func (m *MockService) FindAll() ([]domain.Buyer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]domain.Buyer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockServiceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockService)(nil).FindAll))
}

// FindByEmail mocks base method.
func (m *MockService) FindByEmail(email string) (domain.Buyer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(domain.Buyer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockServiceMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockService)(nil).FindByEmail), email)
}

// FindByID mocks base method.
func (m *MockService) FindByID(id string) (domain.Buyer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(domain.Buyer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockServiceMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockService)(nil).FindByID), id)
}

// FindFavoriteProduct mocks base method.
func (m *MockService) FindFavoriteProduct(buyerID string) ([]domain0.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFavoriteProduct", buyerID)
	ret0, _ := ret[0].([]domain0.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFavoriteProduct indicates an expected call of FindFavoriteProduct.
func (mr *MockServiceMockRecorder) FindFavoriteProduct(buyerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFavoriteProduct", reflect.TypeOf((*MockService)(nil).FindFavoriteProduct), buyerID)
}

// RemoveByID mocks base method.
func (m *MockService) RemoveByID(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveByID indicates an expected call of RemoveByID.
func (mr *MockServiceMockRecorder) RemoveByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByID", reflect.TypeOf((*MockService)(nil).RemoveByID), id)
}

// Update mocks base method.
func (m *MockService) Update(arg0 domain.Buyer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), arg0)
}
