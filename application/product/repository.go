package product

import (
	"fmt"
	"github.com/alexandrebrundias/product-crud/core"
	"github.com/jinzhu/gorm"
	UUID "github.com/satori/go.uuid"
)

type Repoistory struct {
	Db *gorm.DB
}

func NewRepoistory(Db *gorm.DB) *Repoistory {
	return &Repoistory{Db}
}

func (r Repoistory) Insert(product *core.Product) (*core.Product, error) {
	product.ID = UUID.NewV4().String()

	if err := r.Db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r Repoistory) FindAll() ([]*core.Product, error) {
	panic("implement me")
}

func (r Repoistory) FindById(id string) (*core.Product, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	var product core.Product
	r.Db.First(&product, "id = ?", id)

	if product.ID == ""{
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil

}

func (r Repoistory) Update(product *core.Product) (*core.Product, error) {
	panic("implement me")
}

func (r Repoistory) Delete(id string) error {
	panic("implement me")
}

