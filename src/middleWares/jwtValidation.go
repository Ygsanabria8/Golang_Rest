package middleWares

import (
	routes "modules/src/routes"
	"net/http"
)

func JwtValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
