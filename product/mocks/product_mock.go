// Code generated by MockGen. DO NOT EDIT.
// Source: domain/product.go

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "github.com/alexandrebrundias/product-crud/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProductUsecase is a mock of ProductUsecase interface
type MockProductUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockProductUsecaseMockRecorder
}

// MockProductUsecaseMockRecorder is the mock recorder for MockProductUsecase
type MockProductUsecaseMockRecorder struct {
	mock *MockProductUsecase
}

// NewMockProductUsecase creates a new mock instance
func NewMockProductUsecase(ctrl *gomock.Controller) *MockProductUsecase {
	mock := &MockProductUsecase{ctrl: ctrl}
	mock.recorder = &MockProductUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductUsecase) EXPECT() *MockProductUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockProductUsecase) Create(product *domain.Product) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", product)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockProductUsecaseMockRecorder) Create(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductUsecase)(nil).Create), product)
}

// FindById mocks base method
func (m *MockProductUsecase) FindById(id string) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockProductUsecaseMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockProductUsecase)(nil).FindById), id)
}

// Update mocks base method
func (m *MockProductUsecase) Update(product *domain.Product) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", product)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockProductUsecaseMockRecorder) Update(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductUsecase)(nil).Update), product)
}

// Delete mocks base method
func (m *MockProductUsecase) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockProductUsecaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductUsecase)(nil).Delete), id)
}

// MockProductRepository is a mock of ProductRepository interface
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockProductRepository) Insert(product *domain.Product) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", product)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockProductRepositoryMockRecorder) Insert(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockProductRepository)(nil).Insert), product)
}

// FindAll mocks base method
func (m *MockProductRepository) FindAll() ([]*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockProductRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockProductRepository)(nil).FindAll))
}

// FindById mocks base method
func (m *MockProductRepository) FindById(id string) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockProductRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockProductRepository)(nil).FindById), id)
}

// Update mocks base method
func (m *MockProductRepository) Update(product *domain.Product) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", product)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockProductRepositoryMockRecorder) Update(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepository)(nil).Update), product)
}

// Delete mocks base method
func (m *MockProductRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockProductRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductRepository)(nil).Delete), id)
}