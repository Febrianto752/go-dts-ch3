package service

import (
	"github.com/Febrianto752/go-dts-ch3/entity"
	"github.com/Febrianto752/go-dts-ch3/repository"
)

type ProductService interface {
	Create(payload entity.ProductRequest) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	GetById(id uint) (entity.Product, error)
	Update(payload entity.ProductRequest, id uint, userId uint) (entity.Product, error)
	Delete(id uint) (entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func (s *productService) Create(payload entity.ProductRequest) (entity.Product, error) {

	product := entity.Product{
		Title:       payload.Title,
		Description: payload.Description,
		UserId:      payload.UserId,
	}

	newProduct, err := s.productRepository.Create(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *productService) GetAll() ([]entity.Product, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *productService) GetById(id uint) (entity.Product, error) {
	product, err := s.productRepository.FindById(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *productService) Update(payload entity.ProductRequest, id uint, userId uint) (entity.Product, error) {
	product, err := s.productRepository.FindById(id)
	if err != nil {
		panic(err)
	}

	product.Title = payload.Title
	product.Description = payload.Description
	product.UserId = userId

	updatedProduct, err := s.productRepository.Update(product, id)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *productService) Delete(id uint) (entity.Product, error) {
	product, err := s.productRepository.Delete(id)

	if err != nil {
		return product, err
	}

	return product, nil
}

func NewProductService(product repository.ProductRepository) ProductService {
	return &productService{productRepository: product}
}
