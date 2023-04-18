package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" valid:"required~Title of your product is required"`
	Description string `json:"description" valid:"required~Description of your product is required"`
	UserId      uint
	User        *User
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type ProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
