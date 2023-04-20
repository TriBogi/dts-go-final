package repositories

import (
	"DTS-GO-FINAL/database"
	"DTS-GO-FINAL/models"
)

func FindAllSocialMedia(uid uint) ([]models.SocialMedia, error) {
	db := database.GetDB()
	SocialMedias := []models.SocialMedia{}

	err := db.Debug().Preload("User").Find(&SocialMedias, "user_id = ?", uid).Error
	return SocialMedias, err
}

func FindSocialMediaUser(id uint) (*models.SocialMedia, error) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	err := db.Debug().Select("user_id").First(&SocialMedia, id).Error
	return &SocialMedia, err
}

func FindByIdSocialMedia(id uint) (*models.SocialMedia, error) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	err := db.Debug().Preload("User").First(&SocialMedia, id).Error
	return &SocialMedia, err
}

func CreateSocialMedia(SocialMedia *models.SocialMedia) error {
	db := database.GetDB()

	err := db.Debug().Create(&SocialMedia).Error
	return err
}

func UpdateSocialMedia(SocialMediainput *models.SocialMedia) (*models.SocialMedia, error) {
	db := database.GetDB()

	SocialMedia, _ := FindByIdSocialMedia(SocialMediainput.ID)
	err := db.Model(&SocialMedia).Updates(&SocialMediainput).Error

	return SocialMedia, err
}

func DeleteSocialMedia(id uint) error {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	err := db.Debug().Where("id = ?", id).Delete(&SocialMedia).Error

	return err
}
