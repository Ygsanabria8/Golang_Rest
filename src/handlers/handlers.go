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

	router.HandleFunc("/users/register", middleWares.CheckConnectionDataBase(routes.RegisterUser)).Methods("POST")
	router.HandleFunc("/users/login", middleWares.CheckConnectionDataBase(routes.Login)).Methods("POST")
	router.HandleFunc("/users/profile", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(routes.Profile))).Methods("GET")
	router.HandleFunc("/users/update", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(routes.UpdateUser))).Methods("PUT")

	return router
}
