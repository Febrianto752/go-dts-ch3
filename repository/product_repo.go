package repository

import (
	"github.com/Febrianto752/go-dts-ch3/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	FindById(id uint) (entity.Product, error)
	FindAll() ([]entity.Product, error)
	Update(product entity.Product, id uint) (entity.Product, error)
	Delete(id uint) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func (r *productRepository) Create(product entity.Product) (entity.Product, error) {
	err := r.db.Debug().Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.Debug().Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *productRepository) FindById(id uint) (entity.Product, error) {
	var product entity.Product
	err := r.db.Debug().First(&product, "id = ?", id).Error
	if err != nil {
		return product, err
	}

	return product, err
}

func (r *productRepository) Update(product entity.Product, id uint) (entity.Product, error) {

	err := r.db.Debug().Model(&product).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Delete(id uint) (entity.Product, error) {
	var product entity.Product

	err := r.db.Debug().Where("id = ?", id).Delete(&product).Error
	if err != nil {

		return product, err
	}

	return product, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
