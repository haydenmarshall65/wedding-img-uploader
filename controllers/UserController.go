package controllers

import (
	"hayden/wedding-img-uploader/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

}

func GetUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
