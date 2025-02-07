package controllers

import (
	"hayden/wedding-img-uploader/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// var user models.User
	user := models.User{ID: 1, FirstName: "hayden", LastName: "marshall", Email: "hmarshall@example.com", Password: "hayden"}
	models.DB.Create(&user)

	// fmt.Println(main.DB)
	c.IndentedJSON(200, user)
}

func GetUsers(c *gin.Context) {

}

func GetUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
