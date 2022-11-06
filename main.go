package main

import (
	"log"

	dataBase "modules/dataBase"
	handlers "modules/src/handlers"
	utils "modules/src/utils"
)

func main() {
	// Load Env
	utils.LoadEnv()
	// Load DB
	dataBase.ConnectMongoDb()

	if dataBase.CheckConnection() {
		log.Fatal("Without connextion to DB")
	}

	// Load Server
	handlers.Handlers()
}
