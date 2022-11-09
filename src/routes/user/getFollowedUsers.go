package user

import (
	"encoding/json"
	"modules/src/infrastructure/finder"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetFollowedUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")
	typeUser := r.URL.Query().Get("typeUser")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Page is invalid", http.StatusBadRequest)
		return
	}

	users, status := finder.GetFollowedUsers(userId, int64(pageInt), search, typeUser)
	if !status {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}
