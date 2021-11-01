package main

import (
	"thoughts/app"
	"thoughts/app/database"
	"thoughts/app/helpers"
)

// the main function will start the app
func main() {
	helpers.PrepareEnv()
	if database.Connect() != nil {
		println("Connected to the database!!!!")
	}

	server := app.App()
	server.Start()
}
