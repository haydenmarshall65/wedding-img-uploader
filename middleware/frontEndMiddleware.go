package middleware

import (
	"hayden/wedding-img-uploader/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// todo make middleware to check if user is logged in.
func FrontEndMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains("auth/", c.Request.RequestURI) || strings.Contains("api/v1", c.Request.RequestURI) {
			c.Next()
			return
		}

		secrettoken := c.Request.Header.Get("Cookie")

		tokenParts := strings.Split(secrettoken, "secrettoken=")
		token := tokenParts[1]

		// validate the token and user_id against the database
		// if token is valid, set the user in the context
		// if token is invalid, redirect to login
		//
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		var userAuthToken models.PersonalAccessToken

		err := models.DB.Preload("User").Where("token", token).First(&userAuthToken)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		userAuthToken.LastAccessedOn = time.Now().Format("2006-01-02 3:4:5 pm")
		models.DB.Save(&userAuthToken)

		c.Set("authenticated_user", userAuthToken.User)
		c.Next()
	}
}
