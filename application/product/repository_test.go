package product_test

import (
	"github.com/alexandrebrundias/product-crud/application/product"
	"github.com/alexandrebrundias/product-crud/core"
	"github.com/alexandrebrundias/product-crud/infrastructure/database"
	UUID "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductRepoistory_Insert_Find(t *testing.T) {
	db, err := database.NewDatabaseTest().Connect()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := product.NewRepoistory(db)

	p := &core.Product{
		ID:          UUID.NewV4().String(),
		Name:        "TESTE",
		Description: "DESCRIPTION TESTE",
		Quantity:    50,
		Price:       30.2,
	}

	_, err = repo.Insert(p)
	require.Nil(t, err)

	pFind, err := repo.FindById(p.ID)

	require.Nil(t, err)
	require.Equal(t, p.ID, pFind.ID)
	require.Equal(t, p.Price, pFind.Price)
}
