package main

import (
	"hayden/wedding-img-uploader/controllers"
	"hayden/wedding-img-uploader/models"

	"hayden/wedding-img-uploader/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	models.InitDB()

	router.GET("/api/v1/users", controllers.GetUsers)
	router.GET("/api/v1/users/:id", controllers.GetUser)
	router.POST("/api/v1/users", controllers.CreateUser)
	router.PUT("/api/v1/users/:id", controllers.UpdateUser)
	router.DELETE("/api/v1/users/:id", controllers.DeleteUser)

	router.Run(":3000")
}
