package user

import (
	"encoding/json"
	"io"
	"modules/src/infrastructure/repository"
	"modules/src/models"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Invalid data"+err.Error(), http.StatusBadRequest)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]
	var filePath string = "uploads/banners/" + params["userId"] + "." + extension

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copy data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user *models.User

	user.Banner = params["userId"] + "." + extension
	user, status, err := repository.UpdateUser(user, params["userId"])
	if err != nil || !status {
		http.Error(w, "Error saving data in bd: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}
