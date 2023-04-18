package service

import (
	"testing"

	"github.com/Febrianto752/go-dts-ch3/repository"
	"github.com/Febrianto752/go-dts-ch3/service"

	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = service.NewProductService(productRepository)

func TestProductServiceGetOneProduct(t *testing.T) {
	// product := entity.Product{
	// 	Id:   "2",
	// 	Name: "Kacamata",
	// }
	// productRepository.Mock.On("FindById", "2").Return(product)

	// result, err := productService.GetOneProduct("2")

	// assert.Nil(t, err)

	// assert.NotNil(t, result)

	// assert.Equal(t, product.Id, result.Id, "result has to be 2")
	// assert.Equal(t, product.Name, result.Name, "result has to be 'Kacamata'")
	// assert.Equal(t, &product, result, "result has to be a product data with id '2'")

}
