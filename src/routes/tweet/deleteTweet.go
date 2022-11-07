package tweet

import (
	"modules/src/infrastructure/repository"
	jwt "modules/src/jwt"
	"net/http"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	tweetId := r.URL.Query().Get("id")
	if len(tweetId) < 1 {
		http.Error(w, "Tweet Id is requiered", http.StatusBadRequest)
		return
	}

	err := repository.DeleteTweet(tweetId, jwt.UserId)
	if err != nil {
		http.Error(w, "Error deleting tweet: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
