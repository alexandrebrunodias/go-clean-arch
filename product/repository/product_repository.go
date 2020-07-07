package repository

import (
	"fmt"
	"github.com/alexandrebrundias/product-crud/core"
	"github.com/jinzhu/gorm"
	UUID "github.com/satori/go.uuid"
)

type ProductRepoistory struct {
	db *gorm.DB
}

func NewRepoistory(db *gorm.DB) *ProductRepoistory {
	return &ProductRepoistory{db}
}

func (p ProductRepoistory) Insert(product *core.Product) (*core.Product, error) {
	product.ID = UUID.NewV4().String()

	if err := p.db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p ProductRepoistory) FindAll() ([]*core.Product, error) {
	panic("implement me")
}

func (p ProductRepoistory) FindById(id string) (*core.Product, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	var product core.Product
	p.db.First(&product, "id = ?", id)

	if product.ID == ""{
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil

}

func (p ProductRepoistory) Update(product *core.Product) (*core.Product, error) {
	panic("implement me")
}

func (p ProductRepoistory) Delete(id string) error {
	panic("implement me")
}

