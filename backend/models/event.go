package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title            string    `json:"title"`
	ShortDescription string    `json:"shortDescription"`
	LongDescription  string    `json:"longDescription"`
	EventDate        time.Time `json:"dateTime"`
	Organizer        string    `json:"organizer"`
	Location         string    `json:"location"`
	State            string    `json:"state"`
}
