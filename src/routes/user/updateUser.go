package user

import (
	"encoding/json"
	repository "modules/src/infrastructure/repository"
	models "modules/src/models"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	var newUser *models.User
	params := mux.Vars(r)

	newUser, status, err = repository.UpdateUser(user, params["userId"])
	if err != nil {
		http.Error(w, "Error updating user"+err.Error(), http.StatusInternalServerError)
	}

	if !status {
		http.Error(w, "Cannot update user"+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newUser)
}
