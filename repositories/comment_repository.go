package repositories

import (
	"DTS-GO-FINAL/database"
	"DTS-GO-FINAL/models"
)

func FindAllComment(uid uint) ([]models.Comment, error) {
	db := database.GetDB()
	Comments := []models.Comment{}

	err := db.Debug().Preload("Photo").Preload("User").Find(&Comments, "user_id = ?", uid).Error
	return Comments, err
}

func FindCommentUser(id uint) (*models.Comment, error) {
	db := database.GetDB()
	Comment := models.Comment{}

	err := db.Debug().Select("user_id").First(&Comment, id).Error
	return &Comment, err
}

func FindByIdComment(id uint) (*models.Comment, error) {
	db := database.GetDB()
	Comment := models.Comment{}

	err := db.Debug().Preload("Photo").Preload("User").First(&Comment, id).Error
	return &Comment, err
}

func CreateComment(Comment *models.Comment) error {
	db := database.GetDB()

	err := db.Debug().Create(&Comment).Error
	return err
}

func UpdateComment(Commentinput *models.Comment) error {
	db := database.GetDB()

	err := db.Debug().Model(&Commentinput).Where("id = ?", Commentinput.ID).Updates(&Commentinput).Error

	return err
}

func DeleteComment(id uint) error {
	db := database.GetDB()
	Comment := models.Comment{}

	err := db.Debug().Where("id = ?", id).Delete(&Comment).Error

	return err
}
