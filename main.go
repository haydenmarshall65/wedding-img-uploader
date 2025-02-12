package main

import (
	"hayden/wedding-img-uploader/controllers"
	"hayden/wedding-img-uploader/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.InitDB()

	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("users/:id", controllers.DeleteUser)

	router.Run(":3000")
}
