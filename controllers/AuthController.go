package controllers

import (
	"hayden/wedding-img-uploader/models"
	"hayden/wedding-img-uploader/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loggingInUser models.User
	var dbUser models.User

	c.ShouldBindJSON(&loggingInUser)

	password := loggingInUser.Password

	if err := models.DB.Where("email = ?", loggingInUser.Email).First(&dbUser).Error; err != nil {
		c.JSON(404, gin.H{"message": "User not found."})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Password."})
		return
	}

	token, err := utils.GenerateToken(uint(dbUser.ID))

	if err != nil {
		c.JSON(400, gin.H{"message": "Unable to generate token."})
		return
	}

	c.Header("Set-Cookie", "secrettoken="+token+"; HttpOnly; Max-Age:2592000")
	// c.Header("Set-Cookie", "<cookie-name>=<cookie-value>; HttpOnly; Max-Age:2592000")
	c.Status(204)
}

func Register(c *gin.Context) {
	// todo
}
