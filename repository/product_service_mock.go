package repository

import (
	"fmt"

	"github.com/Febrianto752/go-dts-ch3/entity"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id uint) (entity.Product, error) {
	arguments := repository.Mock.Called(id)
	fmt.Println(arguments)

	if arguments.Get(0) == nil {
		return entity.Product{}, nil
	}

	product := arguments.Get(0).(entity.Product)

	return product, nil
}

func (repository *ProductRepositoryMock) FindAll() ([]entity.Product, error) {

	return []entity.Product{}, nil
}
func (repository *ProductRepositoryMock) Create(product entity.Product) (entity.Product, error) {

	return entity.Product{}, nil
}
func (repository *ProductRepositoryMock) Update(product entity.Product, id uint) (entity.Product, error) {

	return entity.Product{}, nil
}
func (repository *ProductRepositoryMock) Delete(id uint) (entity.Product, error) {

	return entity.Product{}, nil
}
