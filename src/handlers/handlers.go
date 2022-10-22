package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers server configuration - Set port and listen server
func Handlers() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "11000"
	}

	cors := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, cors))

}
