package repository

import (
	"github.com/Febrianto752/go-dts-ch3/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Debug().First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) AddUser(user entity.User) (entity.User, error) {
	err := r.db.Debug().Create(&user).Error
	if err != nil {

		return user, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
