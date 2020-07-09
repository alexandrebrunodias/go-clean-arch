package product_test

import (
	"github.com/alexandrebrundias/product-crud/domain"
	"github.com/alexandrebrundias/product-crud/infrastructure/database"
	"github.com/alexandrebrundias/product-crud/product"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductRepoistory_Insert_Find(t *testing.T) {
	db, err := database.NewDatabaseTest().Connect()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	repo := product.NewRepository(db)

	var productFake *domain.Product
	err = faker.FakeData(&productFake)
	assert.NoError(t, err)

	_, err = repo.Insert(productFake)
	assert.NoError(t, err)

	p, err := repo.FindById(productFake.ID)

	assert.NoError(t, err)
	assert.Equal(t, productFake.ID, p.ID)
	assert.Equal(t, productFake.Price, p.Price)
}
