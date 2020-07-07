package product_test

import (
	"github.com/alexandrebrundias/product-crud/application/product"
	"github.com/alexandrebrundias/product-crud/core"
	"github.com/alexandrebrundias/product-crud/infrastructure/database"
	UUID "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var	productMock = &core.Product{
		ID:          UUID.NewV4().String(),
		Name:        "TESTE",
		Description: "DESCRIPTION TESTE",
		Quantity:    50,
		Price:       30.2,
}

func getRepository() *product.Repoistory {
	db, err := database.NewDatabaseTest().Connect()
	if err != nil {
		log.Fatal(err)
	}

	return product.NewRepoistory(db)
}

func TestProductUsecase_Create(t *testing.T) {
	repo := getRepository()
	defer repo.Db.Close()

	usecase := product.NewUsecase(repo)
	_, err := usecase.Create(productMock)
	require.Nil(t, err)

	pFind, err := repo.FindById(productMock.ID)

	require.Nil(t, err)
	require.Equal(t, productMock.ID, pFind.ID)
	require.Equal(t, productMock.Price, pFind.Price)
}

func TestProductUsecase_FindById(t *testing.T) {
	repo := getRepository()
	defer repo.Db.Close()

	usecase := product.NewUsecase(repo)

	repo.Insert(productMock)
	pFind, err := usecase.FindById(productMock.ID)

	require.Nil(t, err)
	require.Equal(t, productMock.ID, pFind.ID)
	require.Equal(t, productMock.Price, pFind.Price)
}
