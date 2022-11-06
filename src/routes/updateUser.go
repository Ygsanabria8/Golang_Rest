package routes

import (
	"encoding/json"
	repository "modules/src/infrastructure/repository"
	models "modules/src/models"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = repository.UpdateUser(user, UserId)
	if err != nil {
		http.Error(w, "Error updating user"+err.Error(), http.StatusInternalServerError)
	}

	if !status {
		http.Error(w, "Cannot update user"+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
