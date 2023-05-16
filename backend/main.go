package main

import (
	"events-app/database"
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

	r.POST("/user/login", ToImplement)

	//TODO: Middleware to check if user is logged in

	r.GET("/event", ToImplement)
	r.GET("/event/:id", ToImplement)
	r.GET("/event/:id/register", ToImplement)
	r.POST("/event", ToImplement)
	r.PUT("/event/:id", ToImplement)

	r.Run()
}

func ToImplement(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "To Implement",
	})
}
