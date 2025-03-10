package controllers

import (
	"hayden/wedding-img-uploader/models"
	"hayden/wedding-img-uploader/utils"
	"time"

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

	// create a personal_access_token in the database:
	// id, user_id, token, expiry_date, last_accessed_on, created_at, updated_at
	var personalToken models.PersonalAccessToken
	personalToken.UserID = dbUser.ID
	personalToken.Token = token
	now := time.Now().AddDate(0, 0, 30)
	personalToken.ExpiryDate = now.Format("2006-01-02 3:4:5 pm")

	if err := models.DB.Create(&personalToken).Error; err != nil {
		c.JSON(500, gin.H{"message": "Unable to generate token"})
	}

	// for production, include Domain=<domain>; Secure; SameSite=Strict;
	c.Header("Set-Cookie", "secrettoken="+token+"; HttpOnly; Max-Age:2592000")
	c.Status(204)
}

func Register(c *gin.Context) {
	// todo
}
