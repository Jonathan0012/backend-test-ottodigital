package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authentication")
	if authHeader != "abc" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}
}
