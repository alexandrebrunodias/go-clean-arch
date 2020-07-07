package core

type Product struct {
	ID          string  `json:"id" gorm:"type:uuid;primary_key"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
}

func NewProduct(ID, Name, Descript string, Quantity int, Price float32) *Product {
	return &Product{ID, Name, Descript, Quantity, Price}
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
