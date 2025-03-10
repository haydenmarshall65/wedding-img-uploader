package controllers

import (
	"errors"
	"hayden/wedding-img-uploader/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var newUser models.User
	c.BindJSON(&newUser)

	password := newUser.Password
	if password == "" {
		c.JSON(400, gin.H{"message": "Password is required."})
		return
	}

	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Password."})
		return
	}

	newUser.Password = string(bcryptHash)

	models.DB.Create(&newUser)

	// saving this for future when working with authentication
	// c.Header("Set-Cookie", "<cookie-name>=<cookie-value>; HttpOnly; Max-Age:2592000")

	c.JSON(200, newUser)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.IndentedJSON(200, users)
}

func GetUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"message": "Unable to find user."})
		return
	}

	var user models.User
	user.ID = id

	if err := models.DB.First(&user).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"message": "User not found."})
		return
	} else {
		c.JSON(200, user)
		return
	}
}

func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"message": "Unable to find user."})
		return
	}

	var user models.User
	user.ID = id

	if err := models.DB.First(&user).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"message": "User not found."})
		return
	}

	pass := []byte(user.Password)

	c.BindJSON(&user)

	if err := bcrypt.CompareHashAndPassword(pass, []byte(user.Password)); err != nil {
		c.JSON(400, gin.H{"message": "Invalid Password."})
		return
	}

	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Password."})
		return
	}

	user.Password = string(bcryptHash)

	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(400, gin.H{"message": "Unable to update user."})
		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"message": "Unable to find user."})
		return
	}

	var user models.User
	user.ID = id

	if err := models.DB.First(&user).Error; err != nil {
		c.JSON(404, gin.H{"message": "Unable to find user."})
		return
	}

	if err := models.DB.Delete(&user).Error; err != nil {
		c.JSON(400, gin.H{"message": "Unable to delete user."})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted."})
}
