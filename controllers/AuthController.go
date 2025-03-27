package controllers

import (
	"hayden/wedding-img-uploader/models"
	"hayden/wedding-img-uploader/utils"
	"log"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterErrors struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (errs RegisterErrors) HasErrors() bool {
	if errs.Email != "" || errs.FirstName != "" || errs.LastName != "" || errs.Password != "" {
		return false
	} else {
		return true
	}
}

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
	var registerringUser models.User
	c.ShouldBindJSON(&registerringUser)

	var registerErrors RegisterErrors

	if registerringUser.FirstName == "" {
		registerErrors.FirstName = "First name is required."
	}
	if registerringUser.LastName == "" {
		registerErrors.LastName = "Last name is required."
	}
	if registerringUser.Email == "" {
		registerErrors.Email = "Email is required."
	} else {
		// validate user email matches expectations
		passesValidation, err := regexp.Match(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, []byte(registerringUser.Email))

		if err != nil {
			log.Print(err)
			c.JSON(500, gin.H{"message": "There was an issue creating the account."})
			return
		}

		if !passesValidation {
			registerErrors.Email = "Please give a valid email."
		}
	}
	if registerringUser.Password == "" {
		registerErrors.Password = "Password is required."
	}

	if registerErrors.HasErrors() {
		c.JSON(400, registerErrors)
		return
	}

	password := registerringUser.Password

	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Password."})
		return
	}

	registerringUser.Password = string(bcryptHash)

	if err := models.DB.Create(&registerringUser).Error; err != nil {
		c.JSON(500, gin.H{"message": "Unable to create new account at this time."})
		return
	}

	c.JSON(200, gin.H{"message": "New account " + registerringUser.Email + " created!"})
}
