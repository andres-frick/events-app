package controller

import (
	"events-app/database"
	"events-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	loginReq := LoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format for login request"})
		return
	}
	var user model.User
	tx := database.DB.Where("username = ?", loginReq.Username).First(&user)

	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	if !user.CheckPassword(loginReq.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}
	hash := uuid.New().String()
	user.Hash = &hash

	database.DB.Model(&user).Update("hash", hash)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "hash": hash, "rol": user.Rol})

}
