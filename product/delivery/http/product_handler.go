package http

import (
	"errors"
	"github.com/alexandrebrundias/product-crud/core"
	"github.com/labstack/echo"
	"net/http"
)

var (
	invalidId = errors.New("ID is not valid")
	notFound = errors.New("Product not found")
)

type ProductHandler struct {
	productUsecase core.ProductUsecase
}

func NewProductHandler(e *echo.Echo, usecase core.ProductUsecase){
	handler := &ProductHandler{usecase}

	e.GET("/products/:id", handler.getByID)
	e.POST("/products", handler.create)
}

func (p *ProductHandler) create(c echo.Context) error {
	var product core.Product
	c.Bind(&product)
	_, err := p.productUsecase.Create(&product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Sever Error")
	}

	return c.JSON(http.StatusOK, product)
}

func (p *ProductHandler) getByID(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusPreconditionRequired, invalidId)
	}

	product, err := p.productUsecase.FindById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, notFound)
	}

	return c.JSON(http.StatusOK, product)
}
