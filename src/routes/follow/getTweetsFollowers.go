package follow

import (
	"encoding/json"
	finder "modules/src/infrastructure/finder"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page is invalid", http.StatusBadRequest)
		return
	}

	result, status := finder.GetTweetsFollowers(userId, int64(page))
	if !status {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
