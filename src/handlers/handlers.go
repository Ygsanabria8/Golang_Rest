package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	middleWares "modules/src/middleWares"
	routes "modules/src/routes"
)

// Handlers server configuration - Set port and listen server
func Handlers() {
	router := createRoutes()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "11000"
	}

	cors := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, cors))

}

// createRoutes Create routes for API
func createRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleWares.CheckConnectionDataBase(routes.RegisterUser)).Methods("POST")

	return router
}
