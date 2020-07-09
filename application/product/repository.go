package product

import (
	"fmt"
	"github.com/alexandrebrundias/product-crud/domain"
	"github.com/jinzhu/gorm"
	UUID "github.com/satori/go.uuid"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(Db *gorm.DB) *Repository {
	return &Repository{Db}
}

func (r Repository) Insert(product *domain.Product) (*domain.Product, error) {
	product.ID = UUID.NewV4().String()

	if err := r.Db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r Repository) FindAll() ([]*domain.Product, error) {
	panic("implement me")
}

func (r Repository) FindById(id string) (*domain.Product, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	var product domain.Product
	r.Db.First(&product, "id = ?", id)

	if product.ID == ""{
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil

}

func (r Repository) Update(product *domain.Product) (*domain.Product, error) {
	panic("implement me")
}

func (r Repository) Delete(id string) error {
	panic("implement me")
}

