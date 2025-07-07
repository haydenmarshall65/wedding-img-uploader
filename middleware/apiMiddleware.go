package middleware

import (
	"hayden/wedding-img-uploader/models"
	"hayden/wedding-img-uploader/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func APIMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if request path includes API, make sure user has api token by grabbing Authorization header and checking the bearer token against
		// a token in the database. If yes, return next(). If no, return 401.
		if !strings.Contains(c.Request.RequestURI, "api/v1") {
			c.Next()
			return
		}
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"message": "Unauthorized: No auth headers provided. Send the correct auth headers."})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(401, gin.H{"message": "Unauthorized: Incorrect auth format. Send the correct auth headers."})
			c.Abort()
			return
		}

		token := tokenParts[1]
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized: Invalid token. Create a new token to continue using the app."})
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(float64)

		if !ok {
			c.JSON(401, gin.H{"message": "Unauthorized: Invalid User. Create a new token and try again."})
			c.Abort()
			return
		}

		authUser := models.User{ID: int(userID)}

		if err := models.DB.First(&authUser).Error; err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized: Invalid User. Please sign in or sign up and try again."})
			c.Abort()
			return
		}

		c.Set("authenticated_user", authUser)
		c.Next()
	}
}
