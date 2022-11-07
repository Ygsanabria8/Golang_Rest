package user

import (
	"io"
	"modules/src/infrastructure/finder"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if len(params["userId"]) < 1 {
		http.Error(w, "User id is required", http.StatusBadRequest)
		return
	}

	profile, err := finder.FindUserById(params["userId"])
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Avatar not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copy avatar", http.StatusBadRequest)
	}

}
