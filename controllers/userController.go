package controllers

import (
	"DTS-GO-FINAL/helpers"
	"DTS-GO-FINAL/models"
	"DTS-GO-FINAL/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var (
	appJSON = "application/json"
)

type Credentials struct {
	Username string `json:"username" form:"username" valid:"required~Your username is required"`
	Password string `json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password minimum lengths is 6 characters"`
}

type RegisterInput struct {
	Username string `json:"username" form:"username"`
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	Age      int32  `json:"age" form:"age"`
}

func Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.FullName = strings.TrimSpace(User.FullName)
	User.Email = strings.TrimSpace(User.Email)
	User.Username = strings.TrimSpace(User.Username)
	User.Password = strings.TrimSpace(User.Password)

	userExist := repositories.FindUser(User.Username, User.Email)
	if userExist.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "sorry, your username or email already registered",
		})
		return
	}

	err := repositories.CreateUser(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"full_name": User.FullName,
		"username":  User.Username,
		"email":     User.Email,
		"message":   "Thank you for joining us ^^",
	})
}

func Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Credentials := Credentials{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Credentials)
	} else {
		c.ShouldBind(&Credentials)
	}

	user, err := repositories.FindByUsername(Credentials.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": fmt.Sprintf("user %s is not exist", Credentials.Username),
		})
		return
	}
	log.Println(user)

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(Credentials.Password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "wrong password, please recheck your password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
