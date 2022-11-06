package tweet

import (
	"encoding/json"
	"net/http"

	repository "modules/src/infrastructure/repository"
	jwt "modules/src/jwt"
	models "modules/src/models"
)

func CreateTwet(w http.ResponseWriter, r *http.Request) {
	var tweet *models.Tweet

	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		http.Error(w, "Error in body: "+err.Error(), http.StatusBadRequest)
		return
	}

	tweet.UserId = jwt.UserId

	tweetDB, status, err := repository.CreateTweet(tweet)
	if err != nil {
		http.Error(w, "Error saving tweet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "Error saving tweet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tweetDB)
}
