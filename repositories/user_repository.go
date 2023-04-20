package repositories

import (
	"DTS-GO-FINAL/database"
	"DTS-GO-FINAL/models"
)

func FindByUsername(username string) (*models.User, error) {
	db := database.GetDB()
	User := models.User{}
	err := db.Debug().Where("username = ?", username).Take(&User).Error
	return &User, err
}

func FindUser(username, email string) *models.User {
	db := database.GetDB()
	userExist := models.User{}
	_ = db.Debug().Where("username = ?", username).Or("email = ?", email).First(&userExist).Error

	return &userExist
}

func CreateUser(user *models.User) error {
	db := database.GetDB()
	err := db.Debug().Create(&user).Error
	return err
}
