package controllers

import (
	"errors"
	"hayden/wedding-img-uploader/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var newUser models.User
	c.BindJSON(&newUser)

	key := newUser.Password
	token := jwt.New(jwt.SigningMethodHS256)
	hashedPassword, err := token.SignedString([]byte(key))
	if err != nil {
		c.IndentedJSON(400, struct{ message string }{message: "Invalid Password."})
	}
	newUser.Password = hashedPassword

	models.DB.Create(&newUser)

	c.IndentedJSON(200, newUser)
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
		c.IndentedJSON(400, struct{ message string }{message: "Invalid User search."})
		return
	}

	var user models.User
	user.ID = id

	if err := models.DB.First(&user).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(404, struct{ message string }{message: "User not found."})
		return
	}

	c.IndentedJSON(200, user)
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
