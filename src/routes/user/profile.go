package user

import (
	"encoding/json"
	"net/http"

	finder "modules/src/infrastructure/finder"

	"github.com/gorilla/mux"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["userId"]
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
