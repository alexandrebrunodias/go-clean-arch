package usecase

import (
	"github.com/alexandrebrundias/product-crud/core"
)

type ProductUsecase struct {
	productRepository core.ProductRepository
}

func NewProductUsecase(repository core.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repository}
}

func (p ProductUsecase) Create(product *core.Product) (*core.Product, error) {
	_, err := p.productRepository.Insert(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p ProductUsecase) FindById(id string) (*core.Product, error) {
	return p.productRepository.FindById(id)
}

func (p ProductUsecase) Update(product *core.Product) (*core.Product, error) {
	panic("implement me")
}

func (p ProductUsecase) Delete(id string) error {
	panic("implement me")
}