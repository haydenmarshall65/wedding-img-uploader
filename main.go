package main

import (
	"hayden/wedding-img-uploader/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("users/:id", controllers.DeleteUser)

	r.Run(":3000")
}
