package routes

import (
	"encoding/json"
	"net/http"

	finder "modules/src/infrastructure/finder"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "User id is mandatory", http.StatusBadRequest)
		return
	}

	profile, err := finder.FindUserById(ID)
	if err != nil {
		http.Error(w, "Error ocurred"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
