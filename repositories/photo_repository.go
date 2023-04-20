package repositories

import (
	"DTS-GO-FINAL/database"
	"DTS-GO-FINAL/models"
)

func FindAllPhoto(uid uint) ([]models.Photo, error) {
	db := database.GetDB()
	Photos := []models.Photo{}

	err := db.Debug().Preload("User").Preload("Comments").Find(&Photos, "user_id = ?", uid).Error
	return Photos, err
}

func FindPhotoUser(id uint) (*models.Photo, error) {
	db := database.GetDB()
	Photo := models.Photo{}

	err := db.Debug().Select("user_id").First(&Photo, id).Error
	return &Photo, err
}

func FindByIdPhoto(id uint) (*models.Photo, error) {
	db := database.GetDB()
	Photo := models.Photo{}

	err := db.Debug().Preload("User").Preload("Comments").First(&Photo, id).Error
	return &Photo, err
}

func CreatePhoto(photo *models.Photo) error {
	db := database.GetDB()

	err := db.Debug().Create(&photo).Error
	return err
}

func UpdatePhoto(photoinput *models.Photo) (*models.Photo, error) {
	db := database.GetDB()
	Photo := models.Photo{}

	err := db.First(&Photo).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&Photo).Updates(&photoinput).Error

	return &Photo, err
}

func DeletePhoto(id uint) error {
	db := database.GetDB()
	Photo := models.Photo{}

	err := db.Debug().Where("id = ?", id).Delete(&Photo).Error

	return err
}
