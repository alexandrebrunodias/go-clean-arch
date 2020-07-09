package product

import (
	"github.com/alexandrebrundias/product-crud/domain"
)

type Usecase struct {
	ProductRepository domain.ProductRepository
}

func NewUsecase(repository domain.ProductRepository) *Usecase {
	return &Usecase{repository}
}

func (p Usecase) Create(product *domain.Product) (*domain.Product, error) {
	_, err := p.ProductRepository.Insert(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p Usecase) FindById(id string) (*domain.Product, error) {
	return p.ProductRepository.FindById(id)
}

func (p Usecase) Update(product *domain.Product) (*domain.Product, error) {
	panic("implement me")
}

func (p Usecase) Delete(id string) error {
	panic("implement me")
}