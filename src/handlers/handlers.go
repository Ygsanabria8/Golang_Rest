package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	middleWares "modules/src/middleWares"
	tweet "modules/src/routes/tweet"
	user "modules/src/routes/user"
	utils "modules/src/utils"
)

// Handlers server configuration - Set port and listen server
func Handlers() {
	router := createRoutes()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = utils.Config.Server.Port
	}

	cors := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, cors))

}

// createRoutes Create routes for API
func createRoutes() *mux.Router {
	router := mux.NewRouter()

	// User Routes
	router.HandleFunc("/users/register", middleWares.CheckConnectionDataBase(user.RegisterUser)).Methods("POST")
	router.HandleFunc("/users/login", middleWares.CheckConnectionDataBase(user.Login)).Methods("POST")
	router.HandleFunc("/users/{userId}", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(user.Profile))).Methods("GET")
	router.HandleFunc("/users/{userId}", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(user.UpdateUser))).Methods("PUT")

	// Tweet Routes
	router.HandleFunc("/users/{userId}/tweets", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(tweet.CreateTwet))).Methods("POST")
	router.HandleFunc("/users/{userId}/tweets", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(tweet.GetTweetsUser))).Methods("GET")
	router.HandleFunc("/users/{userId}/tweets/{tweetId}", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(tweet.DeleteTweet))).Methods("DELETE")
	return router
}
