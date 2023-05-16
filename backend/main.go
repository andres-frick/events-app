package main

import (
	"events-app/database"
	"os"
)

func main() {
	database.StartConnection()

	args := os.Args

	if len(args) > 1 {
		if args[1] == "migrate" {
			database.Migrate()
		}
	}
}
