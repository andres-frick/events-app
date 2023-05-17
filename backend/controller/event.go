package controller

import (
	"events-app/database"
	"events-app/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEventList(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	var events []model.Event

	if user.Rol == "admin" {
		database.DB.Find(&events)
	} else {
		database.DB.Where("state = ?", "publicada").Find(&events)
	}

	c.JSON(http.StatusOK, events)
}

func GetEvent(c *gin.Context) {
	var event model.Event

	if err := database.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	if user.Rol != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	var event model.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&event)

	c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	if user.Rol != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	var event model.Event

	if err := database.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&event)

	c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	if user.Rol != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	var event model.Event

	if err := database.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	database.DB.Delete(&event)

	c.Status(http.StatusNoContent)
}

func RegisterEvent(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	var event model.Event

	if err := database.DB.Where("id = ?", c.Param("id")).Where("state = ?", "publicada").First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if event.EventDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event date is in the past"})
		return
	}

	err := database.DB.Model(&user).Association("Events").Append(&event)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already registered"})
		return
	}

	c.Status(http.StatusCreated)
}

func RegisterEventList(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	var events []model.Event

	database.DB.Model(&user).Association("Events").Find(&events)

	c.JSON(http.StatusOK, events)
}
