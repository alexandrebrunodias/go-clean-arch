package product

import (
	"github.com/alexandrebrundias/product-crud/api/common"
	"github.com/alexandrebrundias/product-crud/domain"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
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
		log.Error(err)
		return nil, common.ErrInternal
	}

	return product, nil
}

func (r Repository) FindAll() ([]*domain.Product, error) {
	panic("implement me")
}

func (r Repository) FindById(id string) (*domain.Product, error) {
	var product domain.Product
	err := r.Db.First(&product, "id = ?", id).Error

	if err != nil {
		log.Error(err)
		return nil, common.ErrInternal
	}

	return &product, nil
}

func (r Repository) Update(product *domain.Product) (*domain.Product, error) {
	panic("implement me")
}

func (r Repository) Delete(id string) error {
	panic("implement me")
}

