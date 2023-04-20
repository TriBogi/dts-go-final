package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null;type:varchar(100)" json:"name" form:"name" valid:"required~Your social media name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Your social media url is required,url~Please input valid url"`
	UserID         uint
	User           *User
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (sm *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(sm)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
