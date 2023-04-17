package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	Title       string `json:"title" form:"title" valid:"required~Title of your product is required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserId      uint
	User        *User
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
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
