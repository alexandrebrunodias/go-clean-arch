package core

import validator "gopkg.in/go-playground/validator.v9"

type Product struct {
	ID          string  `json:"id" gorm:"type:uuid;primary_key"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
}

func NewProduct(ID, Name, Descript string, Quantity int, Price float32) *Product {
	return &Product{ID, Name, Descript, Quantity, Price}
}

func (p Product) Validate() error{
	validate := validator.New()
	return validate.Struct(p)
}

type ProductUsecase interface {
	Create(product *Product) (*Product, error)
	FindById(id string) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id string) error
}

type ProductRepository interface {
	Insert(product *Product) (*Product, error)
	FindAll() ([]*Product, error)
	FindById(id string) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id string) error
}
