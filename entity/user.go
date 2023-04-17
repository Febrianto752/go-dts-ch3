package entity

import (
	"time"

	"github.com/Febrianto752/go-dts-ch3/helper"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	FullName  string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
	Role      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Email    string `valid:"required~Your email is required,email~Invalid email format" json:"email"`
	Password string `valid:"required,minstringlength(6)~Password has to have a minimum length of 6 characters" json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)
	err = nil
	return
}
