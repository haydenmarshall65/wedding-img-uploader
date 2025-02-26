package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func APIMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if request path includes API, make sure user has api token by grabbing Authorization header and checking the bearer token against
		// a token in the database. If yes, return next(). If no, return 401.
		if strings.Contains(c.Request.RequestURI, "api/v1") {
			authHeader := c.Request.Header.Get("Authorization")
			if authHeader == "" {
				c.JSON(401, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		} else {
			c.Next()
		}
	}
}
