package main

import (
	"hayden/wedding-img-uploader/controllers"
	"hayden/wedding-img-uploader/middleware"
	"hayden/wedding-img-uploader/models"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.APIMiddleware())
	router.Use(middleware.FrontEndMiddleware())

	models.InitDB()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/api/v1/users", controllers.GetUsers)
	router.GET("/api/v1/users/:id", controllers.GetUser)
	router.POST("/api/v1/users", controllers.CreateUser)
	router.PUT("/api/v1/users/:id", controllers.UpdateUser)
	router.DELETE("/api/v1/users/:id", controllers.DeleteUser)

	router.POST("/auth/login", controllers.Login)

	router.Run(":3000")
}
