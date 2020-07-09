package product

import (
	"fmt"
	"github.com/alexandrebrundias/product-crud/domain"
)

type Usecase struct {
	ProductRepository domain.ProductRepository
}

func NewUsecase(repository domain.ProductRepository) *Usecase {
	return &Usecase{repository}
}

func (u Usecase) Create(product *domain.Product) (*domain.Product, error) {
	_, err := u.ProductRepository.Insert(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u Usecase) FindById(id string) (*domain.Product, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	product, err := u.ProductRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	if product.ID == ""{
		return nil, fmt.Errorf("product not found")
	}

	return product, nil
}

func (u Usecase) Update(product *domain.Product) (*domain.Product, error) {
	panic("implement me")
}

func (u Usecase) Delete(id string) error {
	panic("implement me")
}