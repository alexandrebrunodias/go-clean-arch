package handler

import (
	"github.com/alexandrebrundias/product-crud/api/common"
	"github.com/alexandrebrundias/product-crud/domain"
	"github.com/labstack/echo"
	"net/http"
)

type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

func NewProductHandler(e *echo.Echo, usecase domain.ProductUsecase) {
	handler := &ProductHandler{usecase}

	e.GET("/products/:id", handler.GetByID)
	e.POST("/products", handler.Create)
	e.POST("/products/", handler.Create)
}

func (p *ProductHandler) Create(c echo.Context) error {
	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	if err := product.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	_, err := p.ProductUsecase.Create(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

func (p *ProductHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	product, err := p.ProductUsecase.FindById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}