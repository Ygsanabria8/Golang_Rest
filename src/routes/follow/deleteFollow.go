package follow

import (
	repository "modules/src/infrastructure/repository"
	models "modules/src/models"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteFollow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userFollowId := r.URL.Query().Get("id")

	if len(userFollowId) < 1 {
		http.Error(w, "Id is requiered", http.StatusBadRequest)
		return
	}

	var follow = &models.Follow{
		UserFollowId: userFollowId,
		UserId:       params["userId"],
	}

	status, err := repository.DeleteFollow(follow)
	if err != nil {
		http.Error(w, "Error deleting follow"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Cannot delete follow"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
