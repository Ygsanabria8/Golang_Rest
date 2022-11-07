package tweet

import (
	"encoding/json"
	"modules/src/infrastructure/finder"
	"net/http"
	"strconv"
)

func GetTweetsUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user")

	if len(userId) < 1 {
		http.Error(w, "User Id is requiered", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page is requiered", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page is invalid", http.StatusBadRequest)
		return
	}

	tweets, status := finder.GetTweetByUserId(userId, int64(page))
	if !status {
		http.Error(w, "Error reading tweets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tweets)

}
