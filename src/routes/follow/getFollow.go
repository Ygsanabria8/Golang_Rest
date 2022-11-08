package follow

import (
	"encoding/json"
	finder "modules/src/infrastructure/finder"
	models "modules/src/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFollow(w http.ResponseWriter, r *http.Request) {
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

	status, err := finder.FindFollow(follow)
	if err != nil || !status {
		follow.IsFollowed = false
	} else {
		follow.IsFollowed = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&follow)
}
