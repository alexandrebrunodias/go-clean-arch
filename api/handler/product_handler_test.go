package handler_test

import (
	"encoding/json"
	"fmt"
	"github.com/alexandrebrundias/product-crud/api/handler"
	"github.com/alexandrebrundias/product-crud/domain"
	"github.com/alexandrebrundias/product-crud/product/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var productFake *domain.Product
	err := faker.FakeData(&productFake)
	productFake.ID = uuid.NewV4().String()
	assert.NoError(t, err)

	j, err := json.Marshal(productFake)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/products", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/products")

	usecaseMock := mocks.NewMockProductUsecase(ctrl)
	usecaseMock.EXPECT().Create(productFake).Return(productFake, nil)

	h := handler.ProductHandler{ProductUsecase: usecaseMock}
	err = h.Create(ctx)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestProductHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var productFake domain.Product
	err := faker.FakeData(&productFake)
	productFake.ID = uuid.NewV4().String()
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/products/" + productFake.ID, nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/products/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(productFake.ID)

	usecaseMock := mocks.NewMockProductUsecase(ctrl)
	usecaseMock.EXPECT().FindById(productFake.ID).Return(&productFake, nil)

	h := handler.ProductHandler{ProductUsecase: usecaseMock}
	err = h.GetByID(ctx)
	assert.NoError(t, err)

	fmt.Println()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, http.StatusOK, rec.Code)

	var productResponse = domain.Product{}
	err = json.NewDecoder(rec.Body).Decode(&productResponse)
	require.NoError(t, err)

	assert.Equal(t, productFake, productResponse)
}