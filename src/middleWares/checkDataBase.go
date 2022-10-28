package middleWares

import (
	database "modules/src/dataBase"
	"net/http"
)

// CheckConnectionDataBase Middleware to check connection to data base
func CheckConnectionDataBase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() {
			http.Error(w, "Lost connection with database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
