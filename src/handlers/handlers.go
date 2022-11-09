package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	middleWares "modules/src/middleWares"
	follow "modules/src/routes/follow"
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
	router.HandleFunc("/users/{userId}/avatar", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(user.UploadAvatar))).Methods("POST")
	router.HandleFunc("/users/{userId}/avatar", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(user.UploadAvatar))).Methods("GET")
	router.HandleFunc("/users/{userId}/banner", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(user.UploadBanner))).Methods("POST")
	router.HandleFunc("/users/{userId}/banner", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(user.UploadBanner))).Methods("GET")
	router.HandleFunc("/users/{userId}/followed", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(user.GetFollowedUsers))).Methods("GET")

	// Follow
	router.HandleFunc("/users/{userId}/follow", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(follow.CreateFollow))).Methods("POST")
	router.HandleFunc("/users/{userId}/follow", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(follow.DeleteFollow))).Methods("DELETE")
	router.HandleFunc("/users/{userId}/follow", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(follow.GetFollow))).Methods("GET")
	router.HandleFunc("/users/{userId}/follow/tweets", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(follow.GetTweetsFollowers))).Methods("GET")

	// Tweet Routes
	router.HandleFunc("/users/{userId}/tweets", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(tweet.CreateTwet))).Methods("POST")
	router.HandleFunc("/users/{userId}/tweets", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(tweet.GetTweetsUser))).Methods("GET")
	router.HandleFunc("/users/{userId}/tweets/{tweetId}", middleWares.CheckConnectionDataBase(middleWares.JwtValidation(tweet.DeleteTweet))).Methods("DELETE")
	return router
}
