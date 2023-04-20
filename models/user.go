package models

import (
	"DTS-GO-FINAL/helpers"
	"errors"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;unique;type:varchar(100)" json:"username" form:"username" valid:"required~Your username is required"`
	FullName string `gorm:"not null;type:varchar(100)" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string `gorm:"not null;unique;type:varchar(100)" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required,minstringlength(6)~Password minimum lengths is 6 characters"`
	Age      int32  `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age <= 8 {
		return errors.New("you must be over 8 years old")
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
