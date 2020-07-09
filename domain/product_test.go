package domain_test

import (
	"github.com/alexandrebrundias/product-crud/domain"
	"github.com/bxcodec/faker/v3"
	UUID "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestProduct_Validate(t *testing.T) {
	p := &domain.Product {
		ID:          faker.UUIDDigit(),
		Name:        faker.Name(),
		Description: faker.Name(),
		Quantity:    rand.Int63(),
		Price:       rand.Float32(),
	}

	err := p.Validate()

	assert.Nil(t, err)
}

func TestProduct_ValidateFailId(t *testing.T) {
	var p = &domain.Product{}
	p.ID = UUID.NewV4().String()

	err := p.Validate()
	assert.NotNil(t, err)
}


func TestProduct_ValidateFailName(t *testing.T) {
	var p = &domain.Product{}
	p.ID = UUID.NewV4().String()
	p.Name = ""

	err := p.Validate()
	assert.NotNil(t, err)
}

func TestProduct_ValidateFailPrice(t *testing.T) {
	var p = &domain.Product{}
	p.ID = UUID.NewV4().String()
	p.Price = 0

	err := p.Validate()
	assert.NotNil(t, err)
}