package controller

import (
	"events-app/database"
	"events-app/model"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type EventFilter struct {
	Title     *string    `form:"title"`
	StartDate *time.Time `form:"start"`
	EndDate   *time.Time `form:"end"`
	State     *string    `form:"state"`
}

func GetEventList(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	var filters EventFilter
	c.ShouldBind(&filters)

	var fields = []string{}
	var value = []any{}

	if filters.Title != nil {
		fields = append(fields, "title LIKE ?")
		value = append(value, "%"+*filters.Title+"%")
	}

	if filters.State != nil {
		fields = append(fields, "state = ?")

		if user.Rol != "admin" {
			value = append(value, "publicada")
		} else {
			value = append(value, *filters.State)
		}
	}

	if filters.State == nil && user.Rol != "admin" {
		fields = append(fields, "state = ?")
		value = append(value, "publicada")
	}

	if filters.StartDate != nil && filters.EndDate != nil && filters.StartDate.Before(*filters.EndDate) {
		fields = append(fields, "event_date BETWEEN ? AND ?")
		value = append(value, *filters.StartDate, *filters.EndDate)
	}

	var events []model.Event

	database.DB.Where(strings.Join(fields, " and "), value...).Debug().Find(&events)

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
	log.Println(event)
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&event)

	c.Status(http.StatusCreated)
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

	c.Status(http.StatusOK)
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

	var filterByStatus = c.Query("status")
	var eventsFilter = []model.Event{}
	var now = time.Now()

	for _, event := range events {
		var active = event.EventDate.After(now)

		switch {
		case filterByStatus == "active" && active:
			eventsFilter = append(eventsFilter, event)
		case filterByStatus == "completed" && !active:
			eventsFilter = append(eventsFilter, event)
		case filterByStatus == "":
			eventsFilter = append(eventsFilter, event)
		}
	}

	c.JSON(http.StatusOK, eventsFilter)
}
