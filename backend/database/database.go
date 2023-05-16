package database

import (
	"events-app/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartConnection() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./database/event.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func Migrate() {
	DB.AutoMigrate(&models.Event{})
	DB.AutoMigrate(&models.User{})
	Populate()
}

func Populate() {
	event1 := &models.Event{
		Title:            "Conferencia de tecnología",
		ShortDescription: "Conferencia Tecnologia",
		LongDescription:  "Una conferencia sobre las últimas tendencias en tecnología",
		EventDate:        time.Date(2023, 10, 10, 10, 0, 0, 0, time.UTC),
		Organizer:        "Tech Conferences Inc.",
		Location:         "Centro de convenciones",
		State:            "publicada",
	}
	event2 := &models.Event{
		Title:            "Taller de desarrollo de aplicaciones móviles",
		ShortDescription: "Conferencia Desarrollo",
		LongDescription:  "Un taller práctico sobre cómo desarrollar aplicaciones móviles",
		EventDate:        time.Date(2023, 06, 10, 15, 0, 0, 0, time.UTC),
		Organizer:        "Mobile Development Academy",
		Location:         "Oficinas de MDA",
		State:            "publicada",
	}
	event3 := &models.Event{
		Title:            "Reunión de planificación",
		ShortDescription: "Conferencia Proyectos",
		LongDescription:  "Una reunión para planificar el próximo proyecto",
		EventDate:        time.Date(2023, 5, 10, 15, 0, 0, 0, time.UTC),
		Organizer:        "Equipo de proyecto",
		Location:         "Sala de conferencias",
		State:            "borrador",
	}
	DB.Create(&event1)
	DB.Create(&event2)
	DB.Create(&event3)

	admin := &models.User{
		Username: "admin",
		Password: "admin",
		Rol:      "admin",
		Hash:     nil,
	}

	user := &models.User{
		Username: "Juan",
		Password: "123",
		Rol:      "user",
		Hash:     nil,
	}

	DB.Create(&admin)
	DB.Create(&user)
}
