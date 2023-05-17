package main

import (
	"events-app/controller"
	"events-app/database"
	"events-app/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	database.StartConnection()

	args := os.Args

	if len(args) > 1 {
		if args[1] == "migrate" {
			database.Migrate()
		}
	}

	r := gin.Default()

	r.POST("/user/login", controller.Login)

	r.Use(middleware.Auth)

	r.GET("/event", controller.GetEventList)
	r.GET("/event/:id", controller.GetEvent)
	r.POST("/event", controller.CreateEvent)
	r.PUT("/event/:id", controller.UpdateEvent)
	r.DELETE("/event/:id", controller.DeleteEvent)
	r.POST("/event/:id/register", controller.RegisterEvent)
	r.GET("/event/registrations", controller.RegisterEventList)

	r.Run("0.0.0.0:4000")
}
