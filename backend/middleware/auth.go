package middleware

import (
	"events-app/database"
	"events-app/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	parts := strings.Split(bearer, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		c.Abort()
		return
	}

	hash := parts[1]

	var user model.User
	tx := database.DB.Where("hash = ?", hash).First(&user)

	if tx.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		c.Abort()
		return
	}

	c.Next()
}
