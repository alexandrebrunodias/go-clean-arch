package product_test

import (
	"github.com/alexandrebrundias/product-crud/domain"
	"github.com/alexandrebrundias/product-crud/product"
	"github.com/alexandrebrundias/product-crud/product/mock"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var productFake *domain.Product
	err := faker.FakeData(&productFake)
	assert.NoError(t, err)

	repository := mock.NewMockProductRepository(ctrl)
	repository.EXPECT().Insert(productFake).Return(productFake, nil)

	usecase := product.NewUsecase(repository)
	p, err := usecase.Create(productFake)

	assert.Nil(t, err)
	assert.Equal(t, productFake.ID, p.ID)
	assert.Equal(t, productFake.Price, p.Price)
}

func TestProductUsecase_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var productFake *domain.Product
	err := faker.FakeData(&productFake)
	assert.NoError(t, err)

	repository := mock.NewMockProductRepository(ctrl)
	repository.EXPECT().FindById(productFake.ID).Return(productFake, nil)

	usecase := product.NewUsecase(repository)
	p, err := usecase.FindById(productFake.ID)

	assert.Nil(t, err)
	assert.Equal(t, productFake.ID, p.ID)
	assert.Equal(t, productFake.Price, p.Price)
}
