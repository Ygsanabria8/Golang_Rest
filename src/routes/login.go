package routes

import (
	"encoding/json"
	"net/http"
	"time"

	dataBase "modules/src/dataBase"
	jwt "modules/src/jwt"
	models "modules/src/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid user or password: "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	userDb, exist := dataBase.Login(user.Email, user.Password)
	if !exist {
		http.Error(w, "Invalid user or password: ", 400)
		return
	}

	jwtKey, err := jwt.GenerateJwt(userDb)
	if err != nil {
		http.Error(w, "Error ocurrend generating tokent"+err.Error(), 500)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	})
}
