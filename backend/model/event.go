package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title            string    `json:"title"`
	ShortDescription string    `json:"short_description"`
	LongDescription  string    `json:"long_description"`
	EventDate        time.Time `json:"event_date"`
	Organizer        string    `json:"organizer"`
	Location         string    `json:"location"`
	State            string    `json:"state"`
}
