package product_test

import (
	"github.com/alexandrebrundias/product-crud/application/product"
	"github.com/alexandrebrundias/product-crud/domain"
	"github.com/alexandrebrundias/product-crud/infrastructure/database"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestProductRepoistory_Insert_Find(t *testing.T) {
	db, err := database.NewDatabaseTest().Connect()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := product.NewRepository(db)

	productFake := &domain.Product{
		ID:          faker.UUIDDigit(),
		Name:        faker.Name(),
		Description: faker.Name(),
		Quantity:    rand.Int63(),
		Price:       rand.Float32(),
	}

	_, err = repo.Insert(productFake)
	require.Nil(t, err)

	p, err := repo.FindById(productFake.ID)

	require.Nil(t, err)
	require.Equal(t, productFake.ID, p.ID)
	require.Equal(t, productFake.Price, p.Price)
}
