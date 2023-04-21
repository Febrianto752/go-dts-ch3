package service

import (
	"errors"
	"testing"

	"github.com/Febrianto752/go-dts-ch3/entity"
	"github.com/Febrianto752/go-dts-ch3/repository"
	"github.com/Febrianto752/go-dts-ch3/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = service.NewProductService(productRepository)

func TestProductServiceGetOneProductFound(t *testing.T) {
	product := entity.Product{
		Id:   1,
		Title: "Gitar",
		Description: "Gitar dengan warna pelangi yang indah",
		UserId: 1,
	}
	productRepository.Mock.On("FindById", uint(1)).Return(product)

	result, err := productService.GetById(uint(1))

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be 1")
	assert.Equal(t, product.Title, result.Title, "result has to be 'Gitar'")
	assert.Equal(t, product, result, "result has to be a product data with id '1'")
}
func TestProductServiceGetOneProductNotFound(t *testing.T) {
	product := entity.Product{}
	errorMessage := errors.New("Product not found")
	productRepository.Mock.On("FindById", uint(2)).Return(nil)

	result, err := productService.GetById(uint(2))

	assert.NotNil(t, err)

	assert.Equal(t, product, result, "result is empty object product")
	assert.Equal(t, err, errorMessage, "error message must be 'Product not found'")
}

func TestProductServiceGetAllProductFound(t *testing.T) {
	products := []entity.Product{
		{
			Id: 1,
			Title: "Laptop A6",
			Description: "Laptop dengan spesifikasi intel core i11",
			UserId: 1,
		},
		{
			Id: 2,
			Title: "Kipas Mantul",
			Description: "Kipas mantap betul",
			UserId: 1,
		},
	}
	
	productRepository.Mock.On("FindAll").Return(products)

	result, err := productService.GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, products, result, "result is products")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	var productService = service.NewProductService(productRepository)
	
	products := []entity.Product{
	}
	errorMessage := errors.New("Products not found")
	productRepository.Mock.On("FindAll").Return(nil)
	result, err := productService.GetAll()
	assert.NotNil(t, err)
	assert.Equal(t, products, result, "result is empty array of object product")
	assert.Equal(t, err, errorMessage, "error message must be 'Products not found'")
}
