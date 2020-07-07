package product

import (
	"github.com/alexandrebrundias/product-crud/application/delivery/http/response"
	"github.com/alexandrebrundias/product-crud/core"
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct {
	ProductUsecase core.ProductUsecase
}

func NewHandler(e *echo.Echo, usecase core.ProductUsecase) {
	handler := &Handler{usecase}

	e.GET("/products/:id", handler.GetByID)
	e.POST("/products", handler.Create)
	e.POST("/products/", handler.Create)
}

func (h *Handler) Create(c echo.Context) error {
	var product core.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: err.Error()})
	}

	if err := product.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: err.Error()})
	}

	_, err := h.ProductUsecase.Create(&product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

func (h *Handler) GetByID(c echo.Context) error {
	id := c.Param("id")
	product, err := h.ProductUsecase.FindById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ItemNotFound())
	}

	return c.JSON(http.StatusOK, product)
}