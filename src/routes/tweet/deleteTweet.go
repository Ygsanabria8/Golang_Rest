package tweet

import (
	"modules/src/infrastructure/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tweetId := params["tweetId"]
	if len(tweetId) < 1 {
		http.Error(w, "Tweet Id is requiered", http.StatusBadRequest)
		return
	}

	err := repository.DeleteTweet(tweetId, params["userId"])
	if err != nil {
		http.Error(w, "Error deleting tweet: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
