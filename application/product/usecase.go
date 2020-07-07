package product

import (
	"github.com/alexandrebrundias/product-crud/core"
)

type Usecase struct {
	ProductRepository core.ProductRepository
}

func NewUsecase(repository core.ProductRepository) *Usecase {
	return &Usecase{repository}
}

func (p Usecase) Create(product *core.Product) (*core.Product, error) {
	_, err := p.ProductRepository.Insert(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p Usecase) FindById(id string) (*core.Product, error) {
	return p.ProductRepository.FindById(id)
}

func (p Usecase) Update(product *core.Product) (*core.Product, error) {
	panic("implement me")
}

func (p Usecase) Delete(id string) error {
	panic("implement me")
}