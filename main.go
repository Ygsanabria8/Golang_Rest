package main

import (
	"log"

	dataBase "modules/src/dataBase"
	handlers "modules/src/handlers"
)

func main() {
	if dataBase.CheckConnection() {
		log.Fatal("Without connextion to DB")
	}

	handlers.Handlers()
}
