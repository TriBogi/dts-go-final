package controllers

import (
	"DTS-GO-FINAL/helpers"
	"DTS-GO-FINAL/models"
	"DTS-GO-FINAL/repositories"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetAllPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	photo, err := repositories.FindAllPhoto(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error getting photo data",
			"err":     err.Error(),
		})
		return
	}

	for _, photo := range photo {
		photo.User.Password = ""
	}
	c.JSON(http.StatusOK, photo)
}

func GetOnePhoto(c *gin.Context) {
	photoID, _ := strconv.Atoi(c.Param("id"))
	photo, err := repositories.FindByIdPhoto(uint(photoID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "photo not found",
				"err":     "not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error getting photo",
			"err":     err.Error(),
		})
		return
	}

	photo.User.Password = ""
	c.JSON(http.StatusOK, &photo)
}

func CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.Title = strings.TrimSpace(Photo.Title)
	Photo.PhotoURL = strings.TrimSpace(Photo.PhotoURL)
	Photo.Caption = strings.TrimSpace(Photo.Caption)

	err := repositories.CreatePhoto(&Photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &Photo)
}

func UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	PhotoInput := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&PhotoInput)
	} else {
		c.ShouldBind(&PhotoInput)
	}

	PhotoInput.UserID = userID
	PhotoInput.ID = uint(photoID)
	PhotoInput.Title = strings.TrimSpace(PhotoInput.Title)
	PhotoInput.PhotoURL = strings.TrimSpace(PhotoInput.PhotoURL)
	PhotoInput.Caption = strings.TrimSpace(PhotoInput.Caption)

	_, err := url.ParseRequestURI(PhotoInput.PhotoURL)
	if err != nil && PhotoInput.PhotoURL != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "invalid url",
		})
		return
	}

	updatedPhoto, err := repositories.UpdatePhoto(&PhotoInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &updatedPhoto)
}

func DeletePhoto(c *gin.Context) {
	photoID, _ := strconv.Atoi(c.Param("id"))

	err := repositories.DeletePhoto(uint(photoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Can't delete photo",
		})
		return
	}

	c.JSON(http.StatusOK, "Photo successfully deleted")
}
