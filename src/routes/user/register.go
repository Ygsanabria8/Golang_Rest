package user

import (
	"encoding/json"
	"net/http"

	finder "modules/src/infrastructure/finder"
	repository "modules/src/infrastructure/repository"
	models "modules/src/models"
)

// RegisterUser create a user in data base
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error in body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Password have to be more than six characters", http.StatusBadRequest)
		return
	}

	_, userExist, _ := finder.FindUserByEmail(user.Email)

	if userExist {
		http.Error(w, "The email already used", http.StatusBadRequest)
		return
	}

	userBD, status, err := repository.CreateUser(user)
	if err != nil {
		http.Error(w, "Error ocurred saving user: "+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "User do not save: "+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&userBD)

}